from .common import DogBreed
from .database import Database

class DogBreedRepository:
    def __init__(self):
        self.db = Database()

    def get_all(self):
        dog_breeds = []
        select_query = "SELECT * FROM dog_breed;"
        try:
            conn = self.db.get_conn()
            cursor = conn.cursor()
            cursor.execute(select_query)
            rows = cursor.fetchall()
            for row in rows:
                dog_breeds.append(
                    DogBreed(
                        row[0],
                        row[1],
                        row[2],
                        row[3],
                        row[4],
                        row[5],
                        row[6],
                        row[7],
                        row[8],
                        row[9],
                        row[10],
                        row[11],
                        row[12],
                        row[13],
                        row[14],
                        row[15],
                        row[16],
                    )
                )
        except Exception as e:
            print(f"Failed to select data: {e}")
            raise
        return dog_breeds