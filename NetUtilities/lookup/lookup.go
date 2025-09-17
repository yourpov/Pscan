package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

type IPInfo struct {
	IP             string  `json:"ip"`
	ISP            string  `json:"isp"`
	Organization   string  `json:"org"`
	Hostname       string  `json:"hostname"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	PostalCode     string  `json:"postal_code"`
	City           string  `json:"city"`
	CountryCode    string  `json:"country_code"`
	CountryName    string  `json:"country_name"`
	ContinentCode  string  `json:"continent_code"`
	ContinentName  string  `json:"continent_name"`
	Region         string  `json:"region"`
	District       string  `json:"district"`
	TimezoneName   string  `json:"timezone_name"`
	ConnectionType string  `json:"connection_type"`
	ASNNumber      int     `json:"asn_number"`
	ASNOrg         string  `json:"asn_org"`
	ASN            string  `json:"asn"`
	CurrencyCode   string  `json:"currency_code"`
	CurrencyName   string  `json:"currency_name"`
	Success        bool    `json:"success"`
	Premium        bool    `json:"premium"`
}

func main() {
	Lookup()
	os.Exit(0)
}

func Lookup() {
	if len(os.Args) < 2 {
		fmt.Println("lookup <ip>")
		return
	}

	ip := os.Args[1]
	hostnames, err := net.LookupAddr(ip)
	if err != nil {
		fmt.Println("\x1b[93mError\x1b[97m: \x1b[93m", err)
		return
	}

	fmt.Println("\x1b[93mIP\x1b[97m: ", ip)
	for _, hostname := range hostnames {
		fmt.Println("\x1b[93mHostname\x1b[97m: ", hostname)
	}

	info, err := getIPInfo(ip)
	if err != nil {
		fmt.Println("\x1b[93mError\x1b[91m:", err)
		return
	}

	fmt.Println("\x1b[93mOrganization\x1b[97m: ", info.Organization)
	fmt.Println("\x1b[93mCity\x1b[97m: ", info.City)
	fmt.Println("\x1b[93mRegion\x1b[97m: ", info.Region)
	fmt.Println("\x1b[93mCountry Code\x1b[97m: ", info.CountryCode)
	fmt.Println("\x1b[93mCountry Name\x1b[97m: ", info.CountryName)
	fmt.Println("\x1b[93mPostal Code\x1b[97m: ", info.PostalCode)
	fmt.Println("\x1b[93mLatitude\x1b[97m: ", info.Latitude)
	fmt.Println("\x1b[93mLongitude\x1b[97m: ", info.Longitude)
	fmt.Println("\x1b[93mContinent Code\x1b[97m: ", info.ContinentCode)
	fmt.Println("\x1b[93mContinent Name\x1b[97m: ", info.ContinentName)
	fmt.Println("\x1b[93mDistrict\x1b[97m: ", info.District)
	fmt.Println("\x1b[93mTimezone Name\x1b[97m: ", info.TimezoneName)
	fmt.Println("\x1b[93mConnection Type\x1b[97m: ", info.ConnectionType)
	fmt.Println("\x1b[93mASN Number\x1b[97m: ", info.ASNNumber)
	fmt.Println("\x1b[93mASN Organization\x1b[97m: ", info.ASNOrg)
	fmt.Println("\x1b[93mASN\x1b[97m: ", info.ASN)
	fmt.Println("\x1b[93mCurrency Code\x1b[97m: ", info.CurrencyCode)
	fmt.Println("\x1b[93mCurrency Name\x1b[97m: ", info.CurrencyName)
	fmt.Println("\x1b[93mSuccess\x1b[97m: ", info.Success)
	fmt.Println("\x1b[93mPremium\x1b[97m: ", info.Premium)
}

func getIPInfo(ip string) (*IPInfo, error) {
	url := fmt.Sprintf("https://json.geoiplookup.io/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var info IPInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
