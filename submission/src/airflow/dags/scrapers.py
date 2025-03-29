from __future__ import annotations

import pendulum
from airflow.models.dag import DAG
from airflow.utils.dates import days_ago
from airflow import settings
from airflow.models import Connection
from airflow.operators.python import PythonOperator
from airflow.exceptions import AirflowException
import docker  # Use the Docker SDK directly
import os

def create_connection():
    session = settings.Session()

    # Check if the connection already exists
    existing_conn = session.query(Connection).filter(Connection.conn_id == "docker_default").first()

    if not existing_conn:
        try:
            conn = Connection(
                conn_id="docker_default",
                conn_type="docker",
                host="unix://var/run/docker.sock",  # Use the local Docker daemon
                extra='{}'  # No need for registry_url
            )
            session.add(conn)
            session.commit()
            print("Docker connection 'docker_default' created.")
        except AirflowException as e:
            session.rollback()
            print(f"Error creating connection: {e}")
            raise
    else:
        print("Docker connection 'docker_default' already exists.")

    session.close()

def run_docker_container(image_name: str):
    # Use the Docker SDK directly to interact with the Docker daemon
    client = docker.DockerClient(base_url="unix://var/run/docker.sock")

    # Run the Docker container
    container = client.containers.run(
        image=image_name,  # Use the local image
        auto_remove=True,  # Remove the container after it exits
        detach=True,  # Run the container in detached mode
        environment={
            "POSTGRES_HOST": "postgres",
            "POSTGRES_PORT": 5432,
            "POSTGRES_USER": "spcapetmatchmaker",
            "POSTGRES_PASSWORD": "spcapetmatchmaker",
            "POSTGRES_DB": "spcapetmatchmaker"
        },
        network="spca_pet_matchmaker_default"
    )

    # Wait for the container to finish and log its output
    for line in container.logs(stream=True):
        print(line.decode().strip())

    # Check the container's exit status
    container.reload()
    if container.status == "exited" and container.attrs["State"]["ExitCode"] != 0:
        raise AirflowException(f"Container failed with exit code {container.attrs['State']['ExitCode']}")

with DAG(
    dag_id="spca_scrapers",
    schedule=None,
    start_date=pendulum.datetime(2023, 10, 27, tz="UTC"),  # Updated start date
    catchup=False,
    dagrun_timeout=pendulum.duration(minutes=60),
    tags=["spca", "scraper", "docker"],
) as dag:
    acquire_docker_connection = PythonOperator(
        task_id="acquire_docker_connection",
        python_callable=create_connection,
    )

    amanzimtoti_scraper = PythonOperator(
        task_id="amanzimtoti_scraper",
        python_callable=run_docker_container,
        op_kwargs={"image_name": "spca_amanzimtoti_scraper:latest"}
    )

    # cape_scraper = PythonOperator(
    #     task_id="cape_scraper",
    #     python_callable=run_docker_container,
    #     op_kwargs={"image_name": "spca_cape_scraper:latest"}
    # )

    durban_and_coast_scraper = PythonOperator(
        task_id="durban_and_coast_scraper",
        python_callable=run_docker_container,
        op_kwargs={"image_name": "spca_durban_and_coast_scraper:latest"}
    )

    lower_south_coast_scraper = PythonOperator(
        task_id="lower_south_coast_scraper",
        python_callable=run_docker_container,
        op_kwargs={"image_name": "spca_lower_south_coast_scraper:latest"}
    )

    randburg_scraper = PythonOperator(
        task_id="randburg_scraper",
        python_callable=run_docker_container,
        op_kwargs={"image_name": "spca_randburg_scraper:latest"}
    )

    roodepoort_scraper = PythonOperator(
        task_id="roodepoort_scraper",
        python_callable=run_docker_container,
        op_kwargs={"image_name": "spca_roodepoort_scraper:latest"}
    )

    acquire_docker_connection >> [
        amanzimtoti_scraper, 
        durban_and_coast_scraper, 
        lower_south_coast_scraper, 
        randburg_scraper,
        roodepoort_scraper
    ]