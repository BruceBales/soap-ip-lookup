package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	XMLName xml.Name
	Body    struct {
		XMLName               xml.Name
		GetIpLocationResponse struct {
			XMLName             xml.Name
			GetIpLocationResult string `xml:"GetIpLocationResult"`
		} `xml:"GetIpLocationResponse"`
	}
}

func main() {

	reqip := os.Args[1]

	var reqbod = fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	  <soap:Body>
		<GetIpLocation xmlns="http://lavasoft.com/">
		  <sIp>%s</sIp>
		</GetIpLocation>
	  </soap:Body>
	</soap:Envelope>`, reqip)
	//76.178.173.88
	soapyboi := new(Response)

	req, err := http.NewRequest("POST", "http://wsgeoip.lavasoft.com/ipservice.asmx", bytes.NewBuffer([]byte(reqbod)))
	if err != nil {
		fmt.Println("Big oof: ", err)
	}
	req.Header.Add("SOAPAction", "http://lavasoft.com/GetIpLocation")

	req.Header.Add("Content-Type", "text/xml")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Big off: ", err)
	}

	decoder := xml.NewDecoder(resp.Body)

	decoder.Decode(soapyboi)
	fmt.Println(soapyboi.Body.GetIpLocationResponse.GetIpLocationResult)
}
