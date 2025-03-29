SELECT
    dog.identifier AS identifier,
    dog.name AS name,
    dog.gender AS gender,
    dog.life_stage AS life_stage,
    dog.image_url AS image_url,
    spca.id AS spca_id,
    COALESCE(dog.breed, 'Unknown breed') AS breed
FROM dog
LEFT JOIN spca ON dog.spca_id = spca.id
WHERE (
    dog.name ILIKE '%' || ? || '%'
    OR spca.name ILIKE '%' || ? || '%'
)