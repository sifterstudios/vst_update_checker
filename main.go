package main

import (
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("C:/Users/thesifter/AppData/Roaming/Steinberg/Cubase 12_64/Cubase Pro VST3 Cache/vst3plugins.xml")

	if err != nil {
		log.Fatal(err)
	}

	println(string(content))
}
