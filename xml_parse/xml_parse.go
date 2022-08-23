package xml_parse

import "encoding/xml"

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
