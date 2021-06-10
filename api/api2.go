package api

import (
	"test_task/input"
	"test_task/models"
)

type SecondApi struct {
	Users []models.User
}

func (a *SecondApi) Get(parameters models.Api2Parameters) models.Api2Response {
	address1 := input.GetAddressByName(a.Users, parameters.Consignee.FirstName, parameters.Consignee.LastName)
	address2 := input.GetAddressByName(a.Users, parameters.Consignor.FirstName, parameters.Consignor.LastName)
	lat1, lon1 := input.GetCoordinatesByAddress(address1)
	lat2, lon2 := input.GetCoordinatesByAddress(address2)
	distance := input.HaversineDistance(lat1, lon1, lat2, lon2)
	amount := 50.0
	if distance > 10 && distance < 100 {
		amount += 30
	} else {
		amount += 100
	}
	if parameters.Cartons.Weight > 20 {
		amount += parameters.Cartons.Weight
	}
	return models.Api2Response{Amount: amount}
}
