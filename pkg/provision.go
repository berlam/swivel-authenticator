package pkg

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"net/url"
)

type ProvisionId string
type ProvisionCode string

type provisionMessagePostPayload struct {
	XMLName       xml.Name      `xml:"SASRequest"`
	Version       string        `xml:"Version"`
	Action        string        `xml:"Action"`
	Username      Username      `xml:"Username"`
	ProvisionCode ProvisionCode `xml:"ProvisionCode"`
	//PushId        string
	//DeviceOS      string
}

type provisionResponse struct {
	XMLName   xml.Name    `xml:"SASResponse"`
	Version   string      `xml:"Version"`
	RequestID string      `xml:"RequestID"`
	Result    string      `xml:"Result"`
	Error     string      `xml:"Error"`
	Id        ProvisionId `xml:"Id"`
}

func postProvisionMessage(url *url.URL, username Username, provisionCode ProvisionCode) (*provisionResponse, error) {
	b := new(bytes.Buffer)
	xml.NewEncoder(b).Encode(provisionMessagePostPayload{
		Version:       "3.6",
		Action:        "provision",
		Username:      username,
		ProvisionCode: provisionCode,
	})
	resp, err := http.Post(url.String()+SUFFIX_AGENT, APPLICATION_TYPE_XML, b)
	if err == nil {
		defer resp.Body.Close()
		var result provisionResponse
		xml.NewDecoder(resp.Body).Decode(&result)
		return &result, nil
	}
	return nil, err
}

func Provision(serverId ServerId, username Username, code ProvisionCode) error {
	// Get the server url from Swivel
	serverUrl, err := GetServerUrl(serverId)

	if err == nil {
		// Start provisioning
		provisionMessage, err := postProvisionMessage(serverUrl, username, code)

		if err == nil {
			// Save provision
			SaveUserConfig(UserHomeDir(), serverId, &UserConfig{
				ProvisionCodeParam: code,
				SecurityStrings:    []SecurityString{},
				LastIndexUsed:      0,
				IndexSecString:     0,
				ProvisionId:        provisionMessage.Id,
			})

			return nil
		}
	}

	return err
}
