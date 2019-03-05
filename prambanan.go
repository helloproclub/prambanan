package prambanan

import (
	"github.com/helloproclub/prambanan/errors"
	"github.com/helloproclub/prambanan/utility"

	"context"
	"strconv"
)

// NewPrambanan constructor like of Prambanan
func NewPrambanan(database Database) (prambanan *Prambanan, err error) {
	prambanan = &Prambanan{
		Database: database,
	}

	return
}

func (n *Prambanan) DecodeNIK(ctx context.Context, nik string) (result Result, err error) {
	if _, err = strconv.ParseUint(nik, 10, 64); err != nil {
		err = errors.ErrNikInvalid
		return
	}

	provinceID := nik[:2]
	cityID := nik[:4]
	districtID := nik[:6]
	result.Province, err = n.Database.GetProvinceByID(provinceID)
	if err == nil {
		result.City, err = n.Database.GetCityByID(provinceID, cityID)
		if err == nil {
			result.District, err = n.Database.GetDistrictByID(cityID, districtID)
		}
	}
	if err != nil {
		return
	}

	result.Gender, result.BirthDate, err = utility.ParseBirthDateCode(nik[6:12])
	if err != nil {
		return
	}

	result.UniqueCode = nik[12:]

	return
}
