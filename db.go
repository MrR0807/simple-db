package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {

	err := SaveData2("hello.txt", []byte("Hello\n"))
	if err != nil {
		panic(err)
	}
}

func SaveData2(path string, data []byte) error {
	tmpFile, err := os.CreateTemp(filepath.Dir(path), filepath.Base(path)+".tmp")
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}
	defer func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}()

	originalFile, err := os.Open(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error opening temp file: %w", err)
	}

	if originalFile != nil {
		if _, err := io.Copy(tmpFile, originalFile); err != nil {
			return fmt.Errorf("error copying temp file: %w", err)
		}
		originalFile.Close()
	}

	if _, err := tmpFile.Write(data); err != nil {
		return fmt.Errorf("error writing temp file: %w", err)
	}

	if err := tmpFile.Sync(); err != nil {
		return fmt.Errorf("error syncing temp file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("error closing temp file: %w", err)
	}

	if err := os.Rename(tmpFile.Name(), path); err != nil {
		return fmt.Errorf("error renaming temp file: %w", err)
	}

	return nil
}
