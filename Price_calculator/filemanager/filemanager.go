package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

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
