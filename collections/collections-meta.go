package collections

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
)

const COLLECTIONS_FILE_NAME = "collections.json"

// FIXME this should be from the CONFIG
const APP_NAME = "apicli"

func GetConfigDirPath() (string, error) {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, APP_NAME), nil
}

func GetAppConfigDirectory() (string, error) {
	configDirPath, err := GetConfigDirPath()

	if err != nil {
		return "", err
	}

	_, err = os.Stat(configDirPath)

	if err != nil {
		return "", err
	}

	return configDirPath, nil
}

func CreateAppConfigDirectory() error {
	configDirPath, err := GetConfigDirPath()

	if err != nil {
		return err
	}

	err = os.Mkdir(configDirPath, 0755)

	if err != nil {
		return err
	}

	return nil
}

type MetaFileScheme struct {
	FieldTest string `json:"fieldTest"`
}

func LoadMetaFile(logger *slog.Logger) (MetaFileScheme, error) {
	configDirPath, err := GetAppConfigDirectory()

	if err != nil {
		logger.Error("cannot get meta file " + err.Error())
		return MetaFileScheme{}, err
	}

	content, err := os.ReadFile(filepath.Join(configDirPath, COLLECTIONS_FILE_NAME))

	if err != nil {
		logger.Error("cannot read meta file " + err.Error())
		return MetaFileScheme{}, err
	}

	m := MetaFileScheme{}
	if err = json.Unmarshal(content, &m); err != nil {
		logger.Error("cannot parse the meta json " + err.Error())
		return MetaFileScheme{}, err
	}

	return m, nil
}

func SaveCollectionsMetadata(logger *slog.Logger) error {
	logger.Debug("entering SaveCollectionsMetadata")
	configDirPath, err := GetAppConfigDirectory()

	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			logger.Info("config directory does not exists, creating")
			err = CreateAppConfigDirectory()
			if err != nil {
				errorMessage := fmt.Sprintf("tried to create config directory %s", err.Error())

				logger.Error(errorMessage)
				return err
			}

		} else {
			error := fmt.Sprintf("cannot get user config %s", err.Error())
			logger.Error(error)
			return fmt.Errorf("%s", error)
		}
	}

	// TODO:: add content from paramter
	content := []byte("{\"fieldTest\": \"test\"}")

	metaDataPath := filepath.Join(configDirPath, COLLECTIONS_FILE_NAME)
	logger.Debug("saving content file to " + COLLECTIONS_FILE_NAME)
	err = os.WriteFile(metaDataPath, content, 0755)

	if err != nil {
		error := fmt.Sprintf("cannot create metadata collection file %s", err.Error())
		logger.Error(error)
		return fmt.Errorf("%s", error)
	}
	logger.Info("file saved")
	return nil
}
