package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"vst_update_checker/constants"
	"vst_update_checker/xml_parse"
)

func main() {
	var xmlFile, err = getInstallationLocation(constants.InstallationLocations)

	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {

		}
	}(xmlFile)

	byteValue, _ := io.ReadAll(xmlFile)

	var plugins xml_parse.Plugins
	err = xml.Unmarshal(byteValue, &plugins)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(plugins.Plugins); i++ {
		fmt.Println("Plugin Vendor: " + plugins.Plugins[i].Class.Vendor)
		fmt.Println("Plugin Name: " + plugins.Plugins[i].Class.Name)
		fmt.Println("Plugin Url: " + plugins.Plugins[i].Url)
		fmt.Println("Plugin Version: " + plugins.Plugins[i].Class.Version)
	}

}

func getInstallationLocation(installationLocations []string) (*os.File, error) {
	var UserConfigDir, _ = os.UserConfigDir()
	for _, v := range installationLocations {
		xmlFile, err := os.Open(filepath.Join(UserConfigDir, v))
		if err != nil {
			//println("XML File not found at " + v + ", continuing...")
			continue
		}

		return xmlFile, nil
	}

	return nil, errors.New("no installation locations found")
}
