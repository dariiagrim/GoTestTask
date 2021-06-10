package input

import (
	"encoding/json"
	"fmt"
	"github.com/umahmood/haversine"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"test_task/models"
	"time"
)

type Input struct {
	Info models.StandardInput
	Users []models.User
}

func (i *Input) LoadInput()  {
	jsonFile, err := os.Open("user_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &i.Users)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetCoordinatesByAddress(address string)  (float64, float64) {
	info := new(models.AddressData)
	addressWithoutSpaces := strings.ReplaceAll(address, " ", "%20")
	_ = getJson(fmt.Sprintf("http://api.positionstack.com/v1/forward?access_key=9e84eb298fcef482cd1f143b9278bc90&query=%s", addressWithoutSpaces), info)
	lat, lon := info.Data[0].Latitude, info.Data[0].Longitude
	return lat, lon
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func HaversineDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	address1 := haversine.Coord{Lat: lat1, Lon: lon1}
	address2  := haversine.Coord{Lat: lat2, Lon: lon2}
	_, km := haversine.Distance(address1, address2)
	return km
}

func GetAddressByName(users []models.User, firstName string, lastName string) string {
	for _, user := range users {
		if user.FirstName == firstName && user.LastName == lastName {
			return user.Address
		}
	}
	return "Not found"
}

func GetNameByAddress(users []models.User, address string) models.PersonInfo {
	for _, user := range users {
		if user.Address == address{
			return models.PersonInfo{
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}
		}
	}
	return models.PersonInfo{
		FirstName: "Not found",
		LastName:  "Not found",
	}
}