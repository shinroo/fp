import psycopg2
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