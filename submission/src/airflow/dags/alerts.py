from __future__ import annotations

import pendulum
from airflow.models.dag import DAG
from airflow.utils.dates import days_ago
from airflow import settings
from airflow.models import Connection
from airflow.operators.python import PythonOperator
from airflow.exceptions import AirflowException
import docker
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
                extra='{}'
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
    client = docker.DockerClient(base_url="unix://var/run/docker.sock")

    # Run the Docker container
    container = client.containers.run(
        image=image_name,
        auto_remove=True,
        detach=True,
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
    dag_id="spca_alerts",
    schedule=None,
    start_date=pendulum.datetime(2023, 10, 27, tz="UTC"),
    catchup=False,
    dagrun_timeout=pendulum.duration(minutes=60),
    tags=["spca", "alerts", "docker"],
) as dag:
    acquire_docker_connection = PythonOperator(
        task_id="acquire_docker_connection",
        python_callable=create_connection,
    )

    alerts_worker = PythonOperator(
        task_id="alerts_worker",
        python_callable=run_docker_container,
        op_kwargs={"image_name": "alerts_worker:latest"}
    )

    acquire_docker_connection >> alerts_worker