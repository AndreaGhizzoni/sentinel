package gitter

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Workspace struct {
	Base      string     `json:"base"`
	Languages []Language `json:"languages"`
}

const JsonSettings = "settings.json"

func NewWorkspace() (*Workspace, error) {
	jsonFileAbs, err := filepath.Abs(JsonSettings)
	if err != nil {
		return nil, err
	}
	jsonFile, err := os.Open(jsonFileAbs)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	workspace := &Workspace{}
	if err := json.Unmarshal(byteValue, &workspace); err != nil {
		return nil, err
	}
	return workspace, nil
}

func (w *Workspace) String() string {
	j, err := json.MarshalIndent(w, "  ", "  ")
	if err != nil {
		return err.Error()
	}
	return string(j)
}
