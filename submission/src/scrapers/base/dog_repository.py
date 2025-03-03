import psycopg2
from psycopg2 import sql
import os

class Database:
    def __init__(self):
        self.db_params = {
            'dbname': os.getenv('POSTGRES_DB'),
            'user': os.getenv('POSTGRES_USER'),
            'password': os.getenv('POSTGRES_PASSWORD'),
            'host': os.getenv('POSTGRES_HOST'),
            'port': os.getenv('POSTGRES_PORT')
        }
        try:
            self.conn = psycopg2.connect(**self.db_params)
        except Exception as e:
            print(f"Failed to connect to database: {e}")
            raise

    def get_conn(self):
        return self.conn

class DogRepository:
    def __init__(self):
        self.db = Database()

    def create(self, identifier, name, gender, life_stage, image_url, spca_id):
        insert_query = """
        INSERT INTO dog (identifier, name, gender, life_stage, image_url, spca_id)
        VALUES (%s, %s, %s, %s, %s, %s)
        ON CONFLICT (identifier) DO NOTHING;
        """
        dog_data = (
            identifier,
            name,
            gender,
            life_stage,
            image_url,
            spca_id
        )
        try:
            conn = self.db.get_conn()
            cursor = conn.cursor()
            cursor.execute(insert_query, dog_data)
            conn.commit()
            print("Data inserted successfully!")
        except Exception as e:
            print(f"Failed to insert data: {e}")
            raise