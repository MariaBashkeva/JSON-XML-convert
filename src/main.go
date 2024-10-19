package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main(){
	fileName := flag.String("f", "", "Path to the database file (JSON or XML)")
	oldFileName := flag.String("old", "", "Path to the old database file (JSON or XML)")
	newFileName := flag.String("new", "", "Path to the new database file (JSON or XML)")
	flag.Parse()
	if *oldFileName != "" && *newFileName != ""  && (strings.HasSuffix(*oldFileName, ".xml") || strings.HasSuffix(*oldFileName, ".json")){
		oldReader, err := NewDBReader(*oldFileName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		oldData, err := oldReader.Read()
		if err != nil {
			fmt.Println("Error reading old data:", err)
			os.Exit(1)
		}
		newReader, err := NewDBReader(*newFileName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		newData, err := newReader.Read()
		if err != nil {
			fmt.Println("Error reading new data:", err)
			os.Exit(1)
		}
		compareDB(oldData,newData)
	}else if *oldFileName != "" && *newFileName != ""  && (strings.HasSuffix(*oldFileName, ".txt") || strings.HasSuffix(*oldFileName, ".txt")){
		if *oldFileName == "" || *newFileName == "" {
			fmt.Println("Usage: ./compareFS --old snapshot1.txt --new snapshot2.txt")
			os.Exit(1)
		}
	
		if err := compareSnapshots(*oldFileName, *newFileName); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}else if *fileName != "" {
		reader, err := NewDBReader(*fileName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		data, err := reader.Read()
		if err != nil {
			fmt.Println("Error reading data:", err)
			os.Exit(1)
		}
		read(data,fileName)
	}
	

}