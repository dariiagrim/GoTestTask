package api

import (
	"github.com/stretchr/testify/assert"
	"test_task/models"
	"testing"
)

func TestFirstApi_Get(t *testing.T) {
	api1 := FirstApi{}
	params1 := models.Api1Parameters{
		ContactAddress:    "555 East Main St, Orange MA 1364",
		WarehouseAddress:  "1095 Industrial Pkwy, Saraland AL 36571",
		PackageDimensions: models.PackageDimensions{
			Length: 10,
			Height: 10,
			Width:  10,
			Weight: 10,
		},
	}
	assert.Equal(t, api1.Get(params1), models.Api1Response{Total:1019.5457566487138})

	params2 := models.Api1Parameters{
		ContactAddress:    "1300 Gilmer Ave, Tallassee AL 36078",
		WarehouseAddress:  "34301 Hwy 43, Thomasville AL 36784",
		PackageDimensions: models.PackageDimensions{
			Length: 25,
			Height: 20,
			Width:  30,
			Weight: 5,
		},
	}
	assert.Equal(t, api1.Get(params2), models.Api1Response{Total: 202.15976496931376})
}