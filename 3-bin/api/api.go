package api

import (
	"bytes"
	"demo/3-bin/bins"
	"demo/3-bin/config"
	"demo/3-bin/files"
	"demo/3-bin/storage"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Metadata struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Private   bool   `json:"private"`
}

type ResponseData struct {
	Metadata Metadata `json:"metadata"`
}

// Функционал:
//   - создает новый Bin
//   - получает данные из хранилища: https://api.jsonbin.io
//   - сохраняет локально информацию о нём и имя
func CreateBin(config *config.Config, storage *storage.Storage) (bool, error) {
	cmd := flag.NewFlagSet("create", flag.ExitOnError)
	fileName := cmd.String("file", "data.json", "Файл с JSON")
	binName := cmd.String("name", "my-bin", "Имя bin")
	cmd.Parse(os.Args[2:])
	fmt.Printf("Вызвана команда create с параметрами file: %s, и name: %s \n", *fileName, *binName)

	storage.Load(config.StorageFilename)
	dataJSON, err := files.ReadFile(*fileName)

	bodyResponse, err := doRequest("POST", config.StoreBaseURL, dataJSON, config)
	if err != nil {
		return false, err
	}

	var data ResponseData
	err = json.Unmarshal(bodyResponse, &data)

	if err != nil {
		return false, err
	}

	fmt.Println("Создание прошло успешно")
	fmt.Println(data.Metadata)

	storeMetadata(data, *binName, *storage, *config)

	return true, nil
}

// Функционал:
//   - обновляет bin из файла по id
func UpdateBin(config *config.Config, storage *storage.Storage) (bool, error) {
	cmd := flag.NewFlagSet("update", flag.ExitOnError)
	fileName := cmd.String("file", "data.json", "Файл с JSON")
	id := cmd.String("id", "", "Идентификатор bin")
	cmd.Parse(os.Args[2:])
	fmt.Printf("Вызвана команда update с параметрами file: %s, и id: %s \n", *fileName, *id)

	storage.Load(config.StorageFilename)
	bin := storage.FindBinByID(*id)

	if bin == nil {
		return false, errors.New("Такой bin отсутствует в локальном хранилище")
	}

	dataJSON, err := files.ReadFile(*fileName)

	targetURL := fmt.Sprint(config.StoreBaseURL, "/", *id)

	bodyResponse, err := doRequest("PUT", targetURL, dataJSON, config)
	if err != nil {
		return false, err
	}

	var data ResponseData
	err = json.Unmarshal(bodyResponse, &data)

	if err != nil {
		return false, err
	}

	fmt.Println("Обновление прошло успешно")
	fmt.Println(data.Metadata)

	return true, nil
}

// Функционал:
//   - удаляет bin по id из API хранилища: https://api.jsonbin.io и локального файла
func DeleteBin(config *config.Config, storage *storage.Storage) (bool, error) {
	cmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := cmd.String("id", "", "Идентификатор bin")
	cmd.Parse(os.Args[2:])
	fmt.Printf("Вызвана команда delete с параметром id: %s \n", *id)

	storage.Load(config.StorageFilename)
	bin := storage.FindBinByID(*id)

	if bin == nil {
		return false, errors.New("Такой bin отсутствует в локальном хранилище")
	}

	targetURL := fmt.Sprint(config.StoreBaseURL, "/", *id)

	bodyResponse, err := doRequest("DELETE", targetURL, []byte{}, config)
	if err != nil {
		return false, err
	}

	var data ResponseData
	err = json.Unmarshal(bodyResponse, &data)

	if err != nil {
		return false, err
	}

	storage.DeleteBinFromStorage(*bin, *config)

	fmt.Println("Удаление прошло успешно")
	fmt.Println(data.Metadata)

	return true, nil
}

// Функционал:
//   - выводит в stdout bin из хранилища: https://api.jsonbin.io по его id
func GetBin(c *config.Config, s *storage.Storage) (bool, error) {
	cmd := flag.NewFlagSet("get", flag.ExitOnError)
	id := cmd.String("id", "", "Идентификатор bin")
	cmd.Parse(os.Args[2:])
	fmt.Printf("Вызвана команда get с параметром id: %s \n", *id)

	s.Load(c.StorageFilename)

	bin := s.FindBinByID(*id)

	if bin == nil {
		return false, errors.New("Такой bin отсутствует в локальном хранилище")
	}

	targetURL := fmt.Sprint(c.StoreBaseURL, "/", *id)

	bodyResponse, err := doRequest("GET", targetURL, []byte{}, c)
	if err != nil {
		return false, err
	}

	var data ResponseData
	err = json.Unmarshal(bodyResponse, &data)

	if err != nil {
		return false, errors.New(err.Error())
	}

	fmt.Println("Поиск прошел успешно")
	fmt.Println(string(bodyResponse))

	return true, nil
}

// Функционал:
//   - выводит список id и name для bin из локального файла
func ListBins(c *config.Config, s *storage.Storage) (bool, error) {
	cmd := flag.NewFlagSet("list", flag.ExitOnError)
	cmd.Parse(os.Args[2:])
	fmt.Println("Вызвана команда list")

	s.Load(c.StorageFilename)

	for _, b := range s.Bins.Bins {
		fmt.Println(b)
	}

	return true, nil
}

func doRequest(method string, url string, data []byte, config *config.Config) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, errors.New(err.Error())
	}

	req.Header["X-Master-Key"] = []string{config.Key}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	fmt.Println(req)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to send request, status code: %d", resp.StatusCode)
		return nil, err
	}

	bodyResponse, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return bodyResponse, nil
}

func storeMetadata(data ResponseData, binName string, s storage.Storage, c config.Config) {
	// Сохраняем метаданные в файл
	bin := bins.NewBin(data.Metadata.ID, binName, data.Metadata.Private)

	var existBin bool

	for _, b := range s.Bins.Bins {
		if b.ID == bin.ID {
			existBin = true
			b.CreatedAt = bin.CreatedAt
			b.Name = bin.Name
			b.Private = bin.Private
		}
	}

	if !existBin {
		s.Bins.Bins = append(s.Bins.Bins, bin)
	}

	s.Save(c.StorageFilename)
}
