package models

type User struct {
	Id int
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string
}


type AddressData struct {
	Data []AddressInfo
}

type AddressInfo struct {
	Latitude float64
	Longitude float64
}

type Api1Parameters struct {
	ContactAddress string `json:"contact_address"`
	WarehouseAddress string `json:"warehouse_address"`
	PackageDimensions PackageDimensions `json:"package_dimensions"`
}

type PackageDimensions struct {
	Length float64
	Height float64
	Width float64
	Weight float64
}

type Api1Response struct {
	Total float64
}

type Api2Response struct {
	Amount float64
}

type StandardInput struct {
	ShippingSourceAddress string
	ShippingDestinationAddress string
	BoxDimensions PackageDimensions
}

type Api2Parameters struct {
	Consignee PersonInfo
	Consignor PersonInfo
	Cartons PackageDimensions `json:"package_dimensions"`
}

type PersonInfo struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}