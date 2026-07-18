package storage

import (
	"os"
	"path/filepath"
)

const (
	appName  = "TaskTracker"
	saveFile = "data.json"
)

func dataPath() (string, error) {
	config, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(
		config,
		appName,
	)
	err = os.MkdirAll(
		dir,
		0755,
	)
	if err != nil {
		return "", err
	}
	return filepath.Join(
		dir,
		saveFile,
	), nil
}
