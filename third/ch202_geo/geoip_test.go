package ch202_geo

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/oschwald/geoip2-golang"
)

// geoip2解析
func TestGeoIpQuickStart(t *testing.T) {
	reader, err := geoip2.Open("./GeoIP2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	ip := net.ParseIP("81.2.69.142")
	record, err := reader.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	if b, err := json.Marshal(record); err == nil {
		fmt.Printf("record:\n%s\n", string(b))
	}
}

// output:
// record:
// {"City":{"GeoNameID":0,"Names":null},"Continent":{"Code":"EU","GeoNameID":6255148,"Names":{"de":"Europa","en":"Europe","es":"Europa","fr":"Europe","ja":"ヨーロッパ","pt-BR":"Europa","ru":"Европа","zh-CN":"欧洲"}},"Country":{"GeoNameID":2635167,"IsInEuropeanUnion":true,"IsoCode":"GB","Names":{"de":"Vereinigtes Königreich","en":"United Kingdom","es":"Reino Unido","fr":"Royaume-Uni","ja":"イギリス","pt-BR":"Reino Unido","ru":"Великобритания","zh-CN":"英国"}},"Location":{"AccuracyRadius":0,"Latitude":0,"Longitude":0,"MetroCode":0,"TimeZone":""},"Postal":{"Code":""},"RegisteredCountry":{"GeoNameID":2635167,"IsInEuropeanUnion":true,"IsoCode":"GB","Names":{"de":"Vereinigtes Königreich","en":"United Kingdom","es":"Reino Unido","fr":"Royaume-Uni","ja":"イギリス","pt-BR":"Reino Unido","ru":"Великобритания","zh-CN":"英国"}},"RepresentedCountry":{"GeoNameID":0,"IsInEuropeanUnion":false,"IsoCode":"","Names":null,"Type":""},"Subdivisions":null,"Traits":{"IsAnonymousProxy":false,"IsSatelliteProvider":false}}

// 根据ip查询国家信息
func TestGetCountryByIP(t *testing.T) {
	// 数据源: https://cdn.jsdelivr.net/npm/geolite2-country@1.0.2/GeoLite2-Country.mmdb.gz
	reader, err := geoip2.Open("./GeoIP2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	// ip := net.ParseIP("81.2.69.142")
	ip := net.ParseIP("102.221.64.25")
	country, err := reader.Country(ip)
	if err != nil {
		log.Fatal(err)
	}

	countryInfo := country.Country
	countryName := countryInfo.Names
	countryIsoCode := countryInfo.IsoCode
	fmt.Println("country name:", countryName)
	fmt.Println("country iso code:", countryIsoCode)
}

// output:
// country name: map[de:Vereinigtes Königreich en:United Kingdom es:Reino Unido fr:Royaume-Uni ja:イギリス pt-BR:Reino Unido ru:Великобритания zh-CN:英国]
// country iso code: GB

// 根据ip查询国家信息+城市信息
func TestGetCityByIP(t *testing.T) {
	// 数据源: https://cdn.jsdelivr.net/npm/geolite2-city@1.0.0/GeoLite2-City.mmdb.gz
	reader, err := geoip2.Open("./GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	ip := net.ParseIP("218.17.161.141")
	city, err := reader.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("city info:", city.City)
	fmt.Println("country info:", city.Country)
}

// output:
// city info: {1795565 map[de:Shenzhen en:Shenzhen es:Shenzhen fr:Shenzhen ja:深セン市 pt-BR:Shenzhen ru:Шэньчжэнь zh-CN:深圳市]}
// country info: {1814991 false CN map[de:China en:China es:China fr:Chine ja:中国 pt-BR:China ru:Китай zh-CN:中国]}
