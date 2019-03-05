-- +mig Up
CREATE TABLE provinces (
  id VARCHAR(2) PRIMARY KEY,
  name VARCHAR(50)
);

-- +mig Down
DROP TABLE provinces;
