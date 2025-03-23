from .database import Database

class DogRepository:
    def __init__(self):
        self.db = Database()

    def create(self, identifier, name, gender, life_stage, image_url, spca_id, embedding):
        insert_query = """
        INSERT INTO dog (identifier, name, gender, life_stage, image_url, spca_id, embedding)
        VALUES (%s, %s, %s, %s, %s, %s, %s)
        ON CONFLICT (identifier) DO NOTHING;
        """
        dog_data = (
            identifier,
            name,
            gender,
            life_stage,
            image_url,
            spca_id,
            embedding
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