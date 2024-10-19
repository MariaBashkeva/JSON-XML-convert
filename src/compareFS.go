package main

import (
	"bufio"
	"fmt"
	"os"
)

func readSnapshot(filePath string) (map[string]bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	filesMap := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		filesMap[line] = true
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return filesMap, nil
}


func compareSnapshots(oldSnapshotPath, newSnapshotPath string) error {
	oldFiles, err := readSnapshot(oldSnapshotPath)
	if err != nil {
		return err
	}
	newFile, err := os.Open(newSnapshotPath)
	if err != nil {
		return err
	}
	defer newFile.Close()
	newFiles := make(map[string]bool)
	scanner := bufio.NewScanner(newFile)
	for scanner.Scan() {
		line := scanner.Text()
		newFiles[line] = true
		if !oldFiles[line] {
			fmt.Println("ADDED", line)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	for file := range oldFiles {
		if !newFiles[file] {
			fmt.Println("REMOVED", file)
		}
	}

	return nil
}

