package Config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rabadaz.at/defaultfiles/Utils"
)

type FileParamenter struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Optional    bool `json:"optional"`
}

type FileConfig struct {
	Tag     string           `json:"tag"`
	Path    string           `json:"path"`
	Version int32            `json:"version"`
	Params  []FileParamenter `json:"params"`
	Technology string		 `json:"technology"`
	RenameTo *string		 `json:"rename_to,omitempty"`
}

type RemoteConfig struct {
	CVersion string       `json:"c_version"`
	Edited   int64        `json:"edited"`
	Files    []FileConfig `json:"files"`
}

func ParseConfig(data []byte) (*RemoteConfig, error) {
	var conf RemoteConfig
	err := json.Unmarshal(data, &conf)
	return &conf , err
}

func downloadConfig() ([]byte, error){
	resp, err := http.Get(Utils.ConstructURL("main_config.json"))
	if err != nil { return nil, err }
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil { return nil, err }
	return data, nil
}

func LoadRemoteConfig() (*RemoteConfig, error){
	raw, err := downloadConfig();

	if err != nil{ return nil, err }

	return ParseConfig(raw)
}

func FindFileInRemoteConfig(config *RemoteConfig, tag string, tech string) *FileConfig{
	for _, file := range config.Files{
		if file.Tag == tag && file.Technology == tech{
			return &file
		}
	}
	return nil
}