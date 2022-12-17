package main

import (
	"net"

	"github.com/IncSW/geoip2"
)

func main() {
	reader, err := geoip2.NewCityReaderFromFile("ip/GeoLite2-City.mmdb")
	if err != nil {
		panic(err)
	}

	record, err := reader.Lookup(net.ParseIP("120.229.72.77"))
	//record, err := reader.Lookup(net.ParseIP("103.170.4.145"))
	//record, err := reader.Lookup(net.ParseIP("2a02:2f0e:dc0b:4400:c71:7513:f404:b4aa"))

	if err != nil {
		panic(err)
	}

	println(record.Continent.Names["zh-CN"])
	println(record.Continent.Names["en"])
	println(record.Country.ISOCode)
	println(record.Country.Names["zh-CN"])
	println(record.Location.TimeZone)
	println(record.City.Names["zh-CN"])
}
