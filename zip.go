package main

/*
import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func writeZip(zipFile string, data []byte) {
	handle, err := openFile(zipFile)
	if err != nil {
		fmt.Println("[ERROR] Opening file:", err)
	}

	zipWriter, err := gzip.NewWriterLevel(handle, 9)
	if err != nil {
		fmt.Println("[ERROR] New gzip writer:", err)
	}
	numberOfBytesWritten, err := zipWriter.Write(data)
	if err != nil {
		fmt.Println("[ERROR] Writing:", err)
	}
	err = zipWriter.Close()
	if err != nil {
		fmt.Println("[ERROR] Closing zip writer:", err)
	}
	fmt.Println("[INFO] Number of bytes written:", numberOfBytesWritten)

	closeFile(handle)
}

func readZip(zipFile string) []byte {
	handle, err := openFile(zipFile)
	if err != nil {
		fmt.Println("[ERROR] Opening file:", err)
	}

	zipReader, err := gzip.NewReader(handle)
	if err != nil {
		fmt.Println("[ERROR] New gzip reader:", err)
	}
	defer zipReader.Close()

	fileContents, err := io.ReadAll(zipReader)
	if err != nil {
		fmt.Println("[ERROR] ReadAll:", err)
	}

	closeFile(handle)

	return fileContents
}

func openFile(fileToOpen string) (*os.File, error) {
	return os.OpenFile(fileToOpen, openFileOptions, openFilePermissions)
}

func closeFile(handle *os.File) {
	if handle == nil {
		return
	}

	err := handle.Close()
	if err != nil {
		fmt.Println("[ERROR] Closing file:", err)
	}
}

const openFileOptions int = os.O_CREATE | os.O_RDWR
const openFilePermissions os.FileMode = 0660
*/
