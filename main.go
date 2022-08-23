package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Plugins struct {
	XMLName xml.Name `xml:"plugins"`
	Plugins []Plugin `xml:"plugin"`
}

type Timestamps struct {
	XMLName    xml.Name `xml:"timestamps"`
	Executable string   `xml:"executable"`
}

type Class struct {
	XMLName       xml.Name `xml:"class"`
	Cid           string   `xml:"cid"`
	Cardinality   int      `xml:"cardinality"`
	Category      string   `xml:"category"`
	Name          string   `xml:"name"`
	ClassFlags    int      `xml:"classFlags"`
	SubCategories string   `xml:"subCategories"`
	Vendor        string   `xml:"vendor"`
	Version       string   `xml:"version"`
	SdkVersion    string   `xml:"sdkVersion"`
}

type Plugin struct {
	XMLName    xml.Name   `xml:"plugin"`
	Path       string     `xml:"path"`
	Vendor     string     `xml:"vendor"`
	Url        string     `xml:"url"`
	Email      string     `xml:"email"`
	Flags      int        `xml:"flags"`
	Codesigned bool       `xml:"codesigned"`
	Timestamps Timestamps `xml:"timestamps"`
	Class      Class      `xml:"class"`
}

func main() {

	var UserConfigDir, _ = os.UserConfigDir()

	xmlFile, err := os.Open(UserConfigDir + "/Steinberg/Cubase 12_64/Cubase Pro VST3 Cache/vst3plugins.xml")

	if err != nil {
		log.Fatal(err) //
	}

	println("Successfully Opened 'vst3plugins.xml'")
	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {

		}
	}(xmlFile)

	byteValue, _ := io.ReadAll(xmlFile)

	var plugins Plugins
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
