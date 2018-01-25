package pkg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type UserConfig struct {
	ProvisionCodeParam ProvisionCode
	ProvisionId        ProvisionId
	IndexSecString     int
	LastIndexUsed      int
	SecurityStrings    []SecurityString
}

func createConfPath(userHomeDir string, serverId ServerId) string {
	return filepath.Clean(userHomeDir + string(filepath.Separator) + "swivel-" + string(serverId))
}

func GetUserConfig(userHomeDir string, serverId ServerId) *UserConfig {
	path := createConfPath(userHomeDir, serverId)
	file, err := ioutil.ReadFile(path)
	var config UserConfig
	if err != nil {
		SaveUserConfig(userHomeDir, serverId, &config)
	}
	b := bytes.NewReader(file)
	json.NewDecoder(b).Decode(&config)
	return &config
}

func SaveUserConfig(userHomeDir string, serverId ServerId, config *UserConfig) error {
	path := createConfPath(userHomeDir, serverId)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(config)
	return ioutil.WriteFile(path, b.Bytes(), 0644)
}
