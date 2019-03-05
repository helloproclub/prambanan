-- +mig Up
CREATE TABLE cities (
  id VARCHAR(2) PRIMARY KEY,
  name VARCHAR(50),
  province_id VARCHAR(2),
  CONSTRAINT fk_province_cities FOREIGN KEY (province_id) REFERENCES provinces(id)
);

-- +mig Down
DROP TABLE cities;
