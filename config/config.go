package config

import (
	"fmt"
	"os"

	"github.com/mrjerz/bookmarks/model"
)

// Reads a configuration file from path and returns all saved bookmarks
func Read(path string) (model.Bookmarks, error) {
	readfile, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		return model.Bookmarks{}, fmt.Errorf("config: couldn't read config file. %s", err)
	}
	defer readfile.Close()

	bms, err := model.Load(readfile)
	if err != nil {
		return model.Bookmarks{}, fmt.Errorf("config: couldn't load config. %s", err)
	}

	return bms, nil
}

// writes all bookmarks back to the configuration file located in path
func Write(path string, b model.Bookmarks) error {

	writefile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return fmt.Errorf("config: couldn't open config for write: %s", err)
	}
	defer writefile.Close()

	if err := b.Save(writefile); err != nil {
		return fmt.Errorf("config: couldn't save config: %s", err)
	}

	return nil
}

// StdConfigPath returns the default configuration path
func StdConfigPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("StdConfigPath: Couldn't retrieve default config path: %s", err)
	}

	return dir + "/.bookmarks", nil
}
