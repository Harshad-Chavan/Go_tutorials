package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputpath string, outputpath string) FileManager {

	newfm := FileManager{
		InputFilePath:  inputpath,
		OutputFilePath: outputpath,
	}

	return newfm
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Error while opening the file")
	}

	// to read the content of the file
	scanner := bufio.NewScanner(file)

	//
	var lines []string
	//read line by line scanner.scan return bool value..if value is there it will return true
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())

	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error while reading the file")
	}

	file.Close()
	return lines, nil

}

func (fm FileManager) WriteJson(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Error in creating a file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to json")

	}

	file.Close()
	return nil

}
