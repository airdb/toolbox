package utilkit

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	FilePerm600 os.FileMode = 0o600 // For secret files.
	FilePerm644 os.FileMode = 0o644 // For normal files.
	FilePerm755 os.FileMode = 0o755 // For directory or execute files.
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func EnsureFolderExists(folder string) {
	if Exists(folder) {
		return
	}

	err := os.MkdirAll(folder, FilePerm755)
	if err != nil {
		log.Fatal("directory not exists, err: ", err)
	}
}

func EnsureFileExists(path string) {
	if Exists(path) {
		return
	}

	EnsureFolderExists(filepath.Dir(path))

	file, err := os.Create(path)
	if err != nil {
		log.Fatal("file not exists, err: ", err)
	}
	file.Close()
}

func ReadFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("read file fail, err: ", err)
		return nil, err
	}

	return content, nil
}

func WriteFile(path string, content string) error {
	EnsureFileExists(path)

	err := os.WriteFile(path, []byte(content), FilePerm644)
	if err != nil {
		return err
	}

	return nil
}

func WriteTarReaderToFile(filename string, tarReader *tar.Reader) {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
	}

	if _, err := io.Copy(outFile, tarReader); err != nil {
		log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
	}

	outFile.Close()
}
