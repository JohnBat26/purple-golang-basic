package storage

import (
	"demo/3-bin/bins"
	"demo/3-bin/files"
	"encoding/json"

	"github.com/fatih/color"
)

type Storage struct {
	Bins bins.BinList `json:"bins"`
}

func (storage *Storage) Save() {
	data, err := storage.ToBytes()

	if err != nil {
		color.Red("Не удалось преобразовать: ", err.Error())
	}

	files.WriteFile(data, "storage.json")

	if err != nil {
		color.Red("Не удалось записать файл data.json: ", err.Error())
	}
}

func (storage *Storage) Load() {
	data, err := files.ReadFile("storage.json")
	if err != nil {
		storage.Bins = bins.BinList{}
	}

	err = json.Unmarshal(data, &storage)

	if err != nil {
		color.Red("Не удалось разобрать файл data.json: ", err.Error(), "\n")

		storage.Bins = bins.BinList{}
	}
}

func (storage *Storage) ToBytes() ([]byte, error) {
	data, err := json.Marshal(storage)
	if err != nil {
		return nil, err
	}

	return data, nil
}
