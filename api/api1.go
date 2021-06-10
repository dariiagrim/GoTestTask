package api

import (
	"test_task/input"
	"test_task/models"
)

type FirstApi struct {}

func (a *FirstApi) Get(parameters models.Api1Parameters) models.Api1Response {
	lat1, lon1 := input.GetCoordinatesByAddress(parameters.ContactAddress)
	lat2, lon2 := input.GetCoordinatesByAddress(parameters.WarehouseAddress)
	distance := input.HaversineDistance(lat1, lon1, lat2, lon2)
	boxSize := parameters.PackageDimensions
	total := distance / 2 + boxSize.Length + (boxSize.Height + boxSize.Width) * 1.5 + boxSize.Weight * 2
	return models.Api1Response{Total: total}
}
