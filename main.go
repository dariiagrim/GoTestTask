package main

import (
	"fmt"
	"test_task/api"
	"test_task/input"
	"test_task/models"
)

func main() {
	in := input.Input{}
	in.LoadInput()
	standardInput := models.StandardInput{ShippingSourceAddress: in.Users[0].Address, ShippingDestinationAddress: in.Users[1].Address}
	standardInput.BoxDimensions = models.PackageDimensions{Length: 50, Height: 30, Width: 20}
	in.Info = standardInput
	minPrice := sendRequests(in, standardInput)
	fmt.Println(minPrice)
}

func sendRequests(in input.Input, standardInput models.StandardInput) float64 {
	firstApi := api.FirstApi{}
	firstApiResponse := firstApi.Get(models.Api1Parameters{
		ContactAddress:    standardInput.ShippingSourceAddress,
		WarehouseAddress:  standardInput.ShippingDestinationAddress,
		PackageDimensions: standardInput.BoxDimensions,
	})
	secondApi := api.SecondApi{}
	secondApiResponse := secondApi.Get(models.Api2Parameters{
		Consignee: input.GetNameByAddress(in.Users, standardInput.ShippingSourceAddress),
		Consignor: input.GetNameByAddress(in.Users, standardInput.ShippingDestinationAddress),
		Cartons:   standardInput.BoxDimensions,
	})
	if firstApiResponse.Total < secondApiResponse.Amount {
		return firstApiResponse.Total
	} else {
		return secondApiResponse.Amount
	}
}
