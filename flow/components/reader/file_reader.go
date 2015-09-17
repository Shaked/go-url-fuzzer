package reader

import (
	"bufio"
	"log"
	"os"

	"github.com/mtojek/go-url-fuzzer/configuration"
)

type fileReader struct {
	inputFile *os.File
}

func newFileReader(configuration *configuration.Configuration) *fileReader {
	inputFile := configuration.FuzzSetFile()
	return &fileReader{inputFile: inputFile}
}

func (f *fileReader) pipe(dataOut chan string, done chan bool) {
	scanner := bufio.NewScanner(f.inputFile)
	f.pipeFileContent(scanner, dataOut)
	f.closeFile()
	done <- true
}

func (f *fileReader) pipeFileContent(scanner *bufio.Scanner, dataOut chan string) {
	for scanner.Scan() {
		line := scanner.Text()
		dataOut <- line
	}
}

func (f *fileReader) closeFile() {
	if error := f.inputFile.Close(); nil != error {
		log.Fatalf("Error occured while closing a file \"%v\": %v", f.inputFile.Name(), error)
	}
}
