package storage

import (
	"demo/3-bin/bins"
	"demo/3-bin/config"
	"demo/3-bin/files"
	"encoding/json"

	"github.com/fatih/color"
)

type Storage struct {
	Bins bins.BinList `json:"bins"`
}

func NewStorage() *Storage {
	return &Storage{
		Bins: bins.BinList{},
	}
}

func (storage *Storage) Save(filename string) {
	data, err := storage.ToBytes()

	if err != nil {
		color.Red("Не удалось преобразовать: ", err.Error())
	}

	files.WriteFile(data, filename)

	if err != nil {
		color.Red("Не удалось записать файл: ", err.Error())
	}
}

func (storage *Storage) Load(filename string) {
	data, err := files.ReadFile(filename)
	if err != nil {
		storage.Bins = bins.BinList{}
	}

	err = json.Unmarshal(data, &storage)

	if err != nil {
		color.Red("Не удалось разобрать файл: ", err.Error(), "\n")

		storage.Bins = bins.BinList{}
	}
}

func (storage *Storage) FindBinByID(id string) *bins.Bin {
	for _, b := range storage.Bins.Bins {
		if b.ID == id {
			return &b
		}
	}

	return nil
}

func (storage *Storage) DeleteBinFromStorage(bin bins.Bin, c config.Config) {
	for i, b := range storage.Bins.Bins {
		if b.ID == bin.ID {
			storage.Bins.Bins = append(storage.Bins.Bins[:i], storage.Bins.Bins[i+1:]...)
		}
	}

	storage.Save(c.StorageFilename)
}

func (storage *Storage) ToBytes() ([]byte, error) {
	data, err := json.Marshal(storage)
	if err != nil {
		return nil, err
	}

	return data, nil
}
