package mysql

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type MySql struct {
	Database *sql.DB
}

func NewMySql(dataSourceName string) (mySql *MySql, err error) {
	mySql.Database, err = sql.Open("mysql", dataSourceName)
	return
}

func (m *MySql) GetProvinceByID(id string) (province string, err error) {
	return
}

func (m *MySql) GetCityByID(provinceID, id string) (city string, err error) {
	return
}

func (m *MySql) GetDistrictByID(cityID, id string) (district string, err error) {
	return
}
