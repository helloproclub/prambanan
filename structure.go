package prambanan

// Prambanan main struct of the project
type Prambanan struct {
	Database Database
}

// Database interface of the database
type Database interface {
	GetProvinceByID(id string) (province string, err error)
	GetCityByID(provinceID, id string) (city string, err error)
	GetDistrictByID(cityID, id string) (district string, err error)
}

// Result hold the result of nik decoder
type Result struct {
	Province   string `json:"provinsi"`
	City       string `json:"kota"`
	District   string `json:"kecamatan"`
	Gender     string `json:"jenis_kelamin"`
	BirthDate  string `json:"tanggal_lahir"`
	UniqueCode string `json:"kode_unik"`
}
