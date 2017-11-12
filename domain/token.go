package domain

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"strconv"
	"fmt"
)

type Pin string
type SecurityString string

type updateKeysMessagePostPayload struct {
	XMLName xml.Name    `xml:"SASRequest"`
	Version string      `xml:"Version"`
	Action  string      `xml:"Action"`
	Id      ProvisionId `xml:"Id"`
}

type updateKeysResponse struct {
	XMLName         xml.Name `xml:"SASResponse"`
	Version         string   `xml:"Version"`
	RequestID       string   `xml:"RequestId"`
	Result          string   `xml:"Result"`
	SecurityStrings string   `xml:"SecurityStrings"`
	Policies        struct {
		Policy []struct {
			Value   string `xml:",chardata"`
			Id      string `xml:"id,attr"`
			Section string `xml:"section,attr"`
			Show    string `xml:"show,attr"`
		}
	}
}

func getNextSecurityKeyIndex(url *url.URL, username Username) int {
	resp, _ := http.Get(url.String() + "/TokenIndex?username=" + string(username))
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	result, _ := strconv.Atoi(string(body))
	return result
}

func postUpdateKeys(url *url.URL, provisionId ProvisionId) []string {
	b := new(bytes.Buffer)
	xml.NewEncoder(b).Encode(updateKeysMessagePostPayload{
		Version: "3.6",
		Action:  "SecurityStrings",
		Id:      provisionId,
	})
	resp, _ := http.Post(url.String()+SUFFIX_AGENT, APPLICATION_TYPE_XML, b)
	defer resp.Body.Close()

	var result updateKeysResponse
	xml.NewDecoder(resp.Body).Decode(&result)
	return strings.Split(result.SecurityStrings, ";")
}

func calculateOTC(secString SecurityString, pin Pin) string {
	otc := ""
	for _, c := range pin {
		idx, _ := strconv.Atoi(string(c))
		if idx == 0 {
			idx = 10
		}
		otc = otc + string(secString[idx-1])
	}
	return otc
}

func Token(serverId ServerId, pin Pin) string {
	// Get user config from home directory
	userConfig := GetUserConfig(UserHomeDir(), serverId)

	// Get the server url from Swivel
	serverUrl := GetServerUrl(serverId)

	// Check for remaining unused local keys
	total := len(userConfig.SecurityStrings) - 1
	if userConfig.IndexSecString >= total-1 {
		updateKeys := postUpdateKeys(serverUrl, userConfig.ProvisionId)
		keys := make([]SecurityString, len(updateKeys))
		for i, k := range updateKeys {
			keys[i] = SecurityString(k)
		}
		userConfig.SecurityStrings = keys
		userConfig.LastIndexUsed = 0
		userConfig.IndexSecString = 0
		SaveUserConfig(UserHomeDir(), serverId, userConfig)
	}
	// Get the index of the next security key
	index := userConfig.IndexSecString
	index++
	if index > userConfig.LastIndexUsed {
		userConfig.IndexSecString = index
		userConfig.LastIndexUsed = index
		SaveUserConfig(UserHomeDir(), serverId, userConfig)
	}
	SaveUserConfig(UserHomeDir(), serverId, userConfig)

	// Get the index of the next security key
	// nextSecurityKeyIndex := getNextSecurityKeyIndex(serverUrl, username)

	// Get the security key
	securityKey := userConfig.SecurityStrings[index]

	// Calculate the otc
	return calculateOTC(securityKey, pin) + fmt.Sprintf("%02d", index)
}
