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
	db, err := geoip2.Open("./GeoIP2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP("81.2.69.142")
	record, err := db.City(ip)
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
func TestGetCountryByIp(t *testing.T) {
	db, err := geoip2.Open("./GeoIP2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP("81.2.69.142")
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	countryInfo := record.Country
	countryName := countryInfo.Names
	countryIsoCode := countryInfo.IsoCode
	fmt.Println("country name:", countryName)
	fmt.Println("country iso code:", countryIsoCode)
}

// output:
// country name: map[de:Vereinigtes Königreich en:United Kingdom es:Reino Unido fr:Royaume-Uni ja:イギリス pt-BR:Reino Unido ru:Великобритания zh-CN:英国]
// country iso code: GB
