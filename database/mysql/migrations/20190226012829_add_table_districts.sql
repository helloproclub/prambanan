-- +mig Up
CREATE TABLE districts (
  id VARCHAR(2) PRIMARY KEY,
  name VARCHAR(50),
  district_id VARCHAR(2),
  CONSTRAINT fk_city_districts FOREIGN KEY (district_id) REFERENCES cities(id)
);

-- +mig Down
DROP TABLE districts;
