package db

import (
	"fmt"
	"strconv"
   	"log"
	"net/url"
   	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/mohamedelkashif/store-location-service/model"
)


var database = make(map[string]interface{})


func FindAll() []interface{} {
	items := make([]interface{}, 0, len(database))
	
	for _, v := range database {
		if v.(*model.Store).CountryCode == "DE" {
			gettingWeatherInfo(v.(*model.Store).CountryCode, v.(*model.Store).Location.Lat, v.(*model.Store).Location.Lng, v);
			
		}else if v.(*model.Store).CountryCode == "FR"{
			gettingWeatherInfo(v.(*model.Store).Country, v.(*model.Store).Location.Lat, v.(*model.Store).Location.Lng, v);
		}
		items = append(items, v)
	}	
	return items
}

func Save(key string, item interface{}) {
	database[key] = item
}

func FindAllByCountry(country_code string, max string) []interface{} {
	items := make([]interface{}, 0, len(database))
	page:=0
	
	for _, v := range database {
		if v.(*model.Store).CountryCode == country_code {
			page++
			items = append(items, v)
		}
		n, err := strconv.Atoi(max)
		if err == nil {
    		fmt.Printf("%d of type %T", n, n)
		}
		fmt.Println(n)
		if page >= n{
			fmt.Println(page)
			break
		}

	}
	fmt.Println(items)
	return items
}

func prepareWeatherInfo(body []byte) (model.Weather, error) {
    var s model.Weather
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return s, err
}

func gettingWeatherInfo(country string, lat float64, lng float64, v interface {}) {
	baseUrl, err := url.Parse("http://localhost:3000")

	if country == "DE"{
		baseUrl.Path += "accuweather/last_minute"
	} else if country == "france"{
		baseUrl.Path += "aerisweather/api/v1/current"
	}

	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
	}
	params := url.Values{}

	params.Add("country", country)
	params.Add("lat", strconv.FormatFloat(lat, 'E', -1, 64))
	params.Add("lng", strconv.FormatFloat(lng, 'E', -1, 64))

	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) 
	s, err := prepareWeatherInfo([]byte(body))
	
	if s.PrecipitationLevel == "HIGH" || s.PrecipitationLast24h > 40 {
		v.(*model.Store).SlowService = true
	} else if s.PrecipitationLevel == "LOW" || s.PrecipitationLast24h < 40{
		v.(*model.Store).SlowService = false
	}else{
		v.(*model.Store).SlowService = false
	}
}