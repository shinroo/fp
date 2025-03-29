SELECT id, name, latitude, longitude, website, address
FROM spca
WHERE id IN (?);