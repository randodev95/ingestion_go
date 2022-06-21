package utils

import (
	"fmt"
	"ingestion_api/configs"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	gl "github.com/ip2location/ip2location-go/v9"
	ua "github.com/mileusna/useragent"
)

var requestHeaders = []string{"X-Client-Ip", "X-Forwarded-For", "Cf-Connecting-Ip", "Fastly-Client-Ip", "True-Client-Ip", "X-Real-Ip", "X-Cluster-Client-Ip", "X-Forwarded", "Forwarded-For", "Forwarded"}
var validate = validator.New()

func GetIpFromRequest(req *http.Request) string {

	// Filtering through the diffenet headers for the IP Adress
	for _, header := range requestHeaders {
		switch header {
		case "X-Forwarded-For":
			header_list := req.Header.Get(header)
			if header_list == "" {
				return ""
			}
			forwardedIp := strings.Split(header_list, ",")
			for _, ip := range forwardedIp {
				ip = strings.TrimSpace(ip)
				if splitted := strings.Split(ip, ":"); len(splitted) == 2 {
					ip = splitted[0]
				}
				if net.ParseIP(ip) != nil {
					return ip
				}
			}

			return ""
		default:
			if ip := req.Header.Get(header); net.ParseIP(ip) != nil {
				return ip
			}
		}
	}

	// Checking For Remote Address

	ip, _, splitHostErr := net.SplitHostPort(req.RemoteAddr)

	if splitHostErr != nil && net.ParseIP(ip) != nil {
		return ip
	}

	return ""

}

func GetRequestDeviceInfo(r *http.Request) UserDeviceInfo {
	dev_info := UserDeviceInfo{}

	client := ua.Parse(r.UserAgent())

	if client.Device == "" && client.OS != "" {
		dev_info.Agent.Device.Name = fmt.Sprintf("%s System", client.OS)
	} else {
		dev_info.Agent.Device.Name = client.Device
	}

	// set struct values of agent in "UserAgentAndIPDetails" struct
	dev_info.Agent.Browser.Name = client.Name
	dev_info.Agent.Browser.Version = client.Version
	dev_info.Agent.Os.Name = client.OS
	dev_info.Agent.Os.Version = client.OSVersion
	dev_info.IsMobile = client.Mobile
	dev_info.IsTablet = client.Tablet
	dev_info.IsDesktop = client.Desktop
	dev_info.IsBot = client.Bot
	dev_info.IsIOS = client.IsIOS()

	return dev_info
}

func GetRequestAdress(ip string) Address {
	db, err := gl.OpenDB(configs.FetchEnv("GEODBPATH"))

	out_address := Address{}

	if err != nil {
		log.Fatal(err)
	}

	results, err := db.Get_all(ip)

	if err != nil {
		log.Fatal(err)
	}

	out_address.Country = results.Country_long
	out_address.CountryCode = results.Country_short
	out_address.City = results.City
	out_address.Region = results.Region
	out_address.ZipCode = results.Zipcode
	out_address.LatLong = [2]float32{results.Latitude, results.Longitude}

	return out_address
}
