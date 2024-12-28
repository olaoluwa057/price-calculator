package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputPath string
	OutputPath string
}


func (fm FileManager) ReadFile() ([]string, error ){
	file, err := os.Open(fm.InputPath)

	if err != nil { 
		
		return nil, errors.New("could not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines,  scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("error scanning")
	}

	return lines, nil


}

func (fm FileManager ) WriteJsonFile( data any) error {
	file, err := os.Create(fm.OutputPath)

	time.Sleep(3 * time.Second)

	if err!= nil {
        return errors.New("could not create file")
    }
	defer file.Close()

	err = json.NewEncoder(file).Encode(data)

	if err!= nil {
     
        return errors.New("error encoding")
    }

	return nil
}


func New(inputPath, outputPath string) FileManager{
	fm := FileManager{
        InputPath:  inputPath,
        OutputPath: outputPath,
    }
    return fm
}