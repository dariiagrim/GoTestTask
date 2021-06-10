package api

import (
	"github.com/stretchr/testify/assert"
	"test_task/models"
	"testing"
)

func TestSecondApi_Get(t *testing.T) {
	api2 := SecondApi{}
	params1 := models.Api2Parameters{
		Consignor:    models.PersonInfo{
			FirstName: "Agnella",
			LastName:  "MacCallam",
		},
		Consignee:  models.PersonInfo{
			FirstName: "Pepita",
			LastName:  "Halegarth",
		},
		Cartons: models.PackageDimensions{
			Length: 10,
			Height: 10,
			Width:  10,
			Weight: 10,
		},
	}
	assert.Equal(t, api2.Get(params1), models.Api2Response{Amount: 150})

	params2 := models.Api2Parameters{
		Consignor:    models.PersonInfo{
			FirstName: "Eliza",
			LastName:  "Slyford",
		},
		Consignee:  models.PersonInfo{
			FirstName: "Mariejeanne",
			LastName:  "Hawthorne",
		},
		Cartons: models.PackageDimensions{
			Length: 20,
			Height: 30,
			Width:  10,
			Weight: 50,
		},
	}
	assert.Equal(t, api2.Get(params2), models.Api2Response{Amount: 200})
}