package fs

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ess/pylades"
)

type Warehouse struct {
	BaseDir string
}

func (warehouse *Warehouse) IDs() []string {
	warehouse.setup()

	keys := make([]string, 0)

	Walk(
		warehouse.BaseDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				keys = append(keys, strings.Replace(path, warehouse.BaseDir+"/", "", 1))
			}

			return nil
		},
	)

	return keys
}

func (warehouse *Warehouse) Store(id string, data []byte) error {
	warehouse.setup()

	abs := filepath.Join(warehouse.BaseDir, id)

	err := CreateDir(filepath.Dir(abs), 0755)
	if err != nil {
		return errors.New("could not write")
	}

	encoded := []byte(base64.StdEncoding.EncodeToString(data))

	err = ioutil.WriteFile(abs, encoded, 0644)
	if err != nil {
		return errors.New("could not write")
	}

	return nil
}

func (warehouse *Warehouse) Retrieve(id string) ([]byte, error) {
	warehouse.setup()

	abs := filepath.Join(warehouse.BaseDir, id)
	if !FileExists(abs) {
		return nil, errors.New("not found")
	}

	data, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, errors.New("could not read file")
	}

	decoded, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, errors.New("could not decode data")
	}

	return decoded, nil
}

func (warehouse *Warehouse) setup() {
	CreateDir(warehouse.BaseDir, 0755)
}

func NewWarehouse(baseDir string) pylades.Warehouse {
	abs, err := filepath.Abs(baseDir)
	if err == nil {
		baseDir = abs
	}

	return &Warehouse{BaseDir: baseDir}
}
