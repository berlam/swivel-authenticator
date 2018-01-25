package pkg

import (
	"bytes"
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
)

type ServerId string

type ServerDetailsResponse struct {
	XMLName      xml.Name `xml:"getServerDetailsResponse"`
	ResponseCode int      `xml:"responseCode"`
	SSDetails    struct {
		Hostname       string `xml:"hostname"`
		UsingSSL       string `xml:"usingSSL"`
		PushSupport    string `xml:"pushSupport"`
		Local          string `xml:"local"`
		Pin            string `xml:"pin"`
		Oath           string `xml:"oath"`
		Port           string `xml:"port"`
		ConnectionType string `xml:"connectionType"`
	}
}

type verificationResponse struct {
	XMLName   xml.Name `xml:"SASResponse"`
	Version   string   `xml:"Version"`
	RequestID string   `xml:"RequestID"`
	Result    string   `xml:"Result"`
}

type manualSettingsVerificationPostPayload struct {
	XMLName xml.Name `xml:"SASRequest"`
	Version string   `xml:"Version"`
	Action  string   `xml:"Action"`
}

func getServerDetails(serverId ServerId) *ServerDetailsResponse {
	resp, _ := http.Get("https://ssd.swivelsecure.net/ssdserver/getServerDetails?id=" + string(serverId))
	defer resp.Body.Close()

	var result ServerDetailsResponse
	err := xml.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Print(err)
	}
	return &result
}

func postManualSettingsVerification(url *url.URL) *verificationResponse {
	b := new(bytes.Buffer)
	xml.NewEncoder(b).Encode(manualSettingsVerificationPostPayload{
		Version: "3.6",
		Action:  "ping",
	})
	resp, _ := http.Post(url.String()+SUFFIX_AGENT, APPLICATION_TYPE_XML, b)
	defer resp.Body.Close()

	var result verificationResponse
	xml.NewDecoder(resp.Body).Decode(&result)
	return &result
}

func buildServerUrl(serverDetails *ServerDetailsResponse) *url.URL {
	scheme := "http"
	if serverDetails.SSDetails.UsingSSL == "YES" {
		scheme = "https"
	}
	return &url.URL{
		Scheme: scheme,
		Host:   serverDetails.SSDetails.Hostname + ":" + serverDetails.SSDetails.Port,
		Path:   serverDetails.SSDetails.ConnectionType,
	}
}

func GetServerUrl(serverId ServerId) *url.URL {
	return buildServerUrl(getServerDetails(serverId))
}
