package collections

import (
	"errors"
	"os"
	"path/filepath"
)

// FIXME this file should contain some mutex to concurrent operations
type FileLocation struct {
	Path    string
	Content []byte
}

func (f FileLocation) GetName() string {
	return "FILE_PROVIDER"
}

func (f FileLocation) Save() error {
	if _, err := os.Stat(f.Path); err != nil {
		return err
	}
	return os.WriteFile(f.Path, []byte("{}"), 0755)
}

func (f FileLocation) Create(name string) (FileLocation, error) {
	// TODO we should be able to get create a collection in a different location
	configDirPath, err := GetAppConfigDirectory()

	if err != nil {
		return FileLocation{}, err
	}

	fileLocation := filepath.Join(configDirPath, name)

	if _, err = os.Stat(fileLocation); err == nil {
		return FileLocation{}, errors.New("file already exists")
	}

	contentInit := []byte("{}")
	if err = os.WriteFile(fileLocation, []byte(contentInit), 0755); err != nil {
		return FileLocation{}, err
	}

	return FileLocation{
		Path:    fileLocation,
		Content: []byte(contentInit),
	}, nil
}

func Get(filepath string) (FileLocation, error) {
	if _, err := os.Stat(filepath); err != nil {
		return FileLocation{}, err
	}

	content, err := os.ReadFile(filepath)

	if err != nil {
		return FileLocation{}, err
	}

	return FileLocation{
		Path:    filepath,
		Content: content,
	}, nil
}
