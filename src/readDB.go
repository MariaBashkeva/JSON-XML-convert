package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Ingredient struct {
	IngredientName  string  `json:"ingredient_name" xml:"itemname"`
	IngredientCount string `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  string  `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type BakeryData struct {
	Cakes []Cake `json:"cake" xml:"cake"`
}

type DBReader interface {
	Read() (BakeryData, error)
}

type JSONReader struct {
	file *os.File
}

func (r *JSONReader) Read() (BakeryData, error) {
	var data BakeryData
	decoder := json.NewDecoder(r.file)
	err := decoder.Decode(&data)
	return data, err
}

type XMLReader struct {
	file *os.File
}

func (r *XMLReader) Read() (BakeryData, error) {
	var data BakeryData
	decoder := xml.NewDecoder(r.file)
	err := decoder.Decode(&data)
	return data, err
}

func NewDBReader(filename string) (DBReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	if strings.HasSuffix(filename, ".json") {
		return &JSONReader{file: file}, nil
	} else if strings.HasSuffix(filename, ".xml") {
		return &XMLReader{file: file}, nil
	} else {
		file.Close() 
		return nil, fmt.Errorf("unsupported file format: %s", filename)
	}
}

func read(data BakeryData,fileName *string) {
	if strings.HasSuffix(*fileName, ".json") {
		xmlData, err := xml.MarshalIndent(data, "", "    ")
		if err != nil {
			fmt.Println("Error formatting as XML:", err)
			os.Exit(1)
		}
		fmt.Println(string(xmlData))
	} else if strings.HasSuffix(*fileName, ".xml") {
		jsonData, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			fmt.Println("Error formatting as JSON:", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	}
}
