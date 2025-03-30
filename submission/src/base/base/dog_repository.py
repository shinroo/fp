from .common import Dog, Gender, LifeStage
from .database import Database

class DogRepository:
    def __init__(self):
        self.db = Database()

    def get_similar(self, account_id, threshold):
        matching_dogs = []

        select_query = """
        SELECT
        dog.name AS name,
        dog.breed AS breed,
        spca.name AS spca_name,
        dog.image_url AS image_url
        FROM dog
        LEFT JOIN spca ON dog.spca_id = spca.id
        LEFT JOIN dog_breed ON dog.breed = dog_breed.name
        WHERE 1=1
        AND ARRAY[
            COALESCE(dog_breed.affectionate_with_family,1),
            COALESCE(dog_breed.good_with_young_children,1),
            COALESCE(dog_breed.good_with_other_dogs,1), 
            COALESCE(dog_breed.shedding_level,5),
            COALESCE(dog_breed.coat_grooming_frequency,5),
            COALESCE(dog_breed.drooling_level,5),
            COALESCE(dog_breed.coat_length,5),
            COALESCE(dog_breed.openness_to_strangers,1),
            COALESCE(dog_breed.playfulness_level,1),
            COALESCE(dog_breed.watchdog_protective_nature,1),
            COALESCE(dog_breed.adaptability_level,1),
            COALESCE(dog_breed.trainability_level,1),
            COALESCE(dog_breed.energy_level,5),
            COALESCE(dog_breed.barking_level,5),
            COALESCE(dog_breed.mental_stimulation_needs,5)
        ]::vector <-> (
            SELECT
            ARRAY[
                COALESCE(profile.affectionate_with_family,1),
                COALESCE(profile.good_with_young_children,1),
                COALESCE(profile.good_with_other_dogs,1), 
                COALESCE(profile.shedding_level,5),
                COALESCE(profile.coat_grooming_frequency,5),
                COALESCE(profile.drooling_level,5),
                COALESCE(profile.coat_length,5),
                COALESCE(profile.openness_to_strangers,1),
                COALESCE(profile.playfulness_level,1),
                COALESCE(profile.watchdog_protective_nature,1),
                COALESCE(profile.adaptability_level,1),
                COALESCE(profile.trainability_level,1),
                COALESCE(profile.energy_level,5),
                COALESCE(profile.barking_level,5),
                COALESCE(profile.mental_stimulation_needs,5)
            ]::vector AS embedding
            FROM profile
            WHERE account_id = %s
        ) > %s
        LIMIT 1
        """
        try:
            conn = self.db.get_conn()
            cursor = conn.cursor()
            cursor.execute(select_query, (account_id, threshold))
            rows = cursor.fetchall()
            for row in rows:
                matching_dogs.append(
                    Dog(
                        row[0],
                        row[1],
                        row[2],
                        row[3]
                    )
                )

        except Exception as e:
            print(f"Failed to select data: {e}")
            raise
        return matching_dogs

    def get_specific(self, breed, gender: Gender, life_stage: LifeStage):
        matching_dogs = []

        select_query = """
        SELECT
        dog.name AS name,
        dog.breed AS breed,
        spca.name AS spca_name,
        dog.image_url AS image_url
        FROM dog
        LEFT JOIN spca ON dog.spca_id = spca.id
        WHERE 1=1
        AND gender = %s
        AND life_stage = %s
        AND breed = %s
        LIMIT 1
        """

        try:
            conn = self.db.get_conn()
            cursor = conn.cursor()
            cursor.execute(select_query, (gender, life_stage, breed))
            rows = cursor.fetchall()
            for row in rows:
                matching_dogs.append(
                    Dog(
                        row[0],
                        row[1],
                        row[2],
                        row[3]
                    )
                )

        except Exception as e:
            print(f"Failed to select data: {e}")
            raise
        return matching_dogs

    def create(self, identifier, name, gender, life_stage, image_url, spca_id, embedding, breed):
        if breed != "Unknown":
            insert_query = """
            INSERT INTO dog (identifier, name, gender, life_stage, image_url, spca_id, breed)
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
                breed,
            )
        else:
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
                spca_id,
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