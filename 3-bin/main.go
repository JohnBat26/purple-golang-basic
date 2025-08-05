package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 1 {
		fmt.Println("Необходимо использовать хотя бы один параметр: --create, --update, --delete, --get или --list")
		return
	}

	switch os.Args[1] {
	case "--create":
		_, err = createBin()
	case "--update":
		_, err = updateBin()
	case "--delete":
		_, err = deleteBin()
	case "--get":
		_, err = getBin()
	case "--list":
		_, err = listBins()
	default:
		fmt.Println("Необходимо использовать хотя бы один параметр: --create, --update, --delete, --get или --list")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}

// Функционал:
//   - создает новый Bin
//   - получает данные из json.bin
//   - сохраняет локально информацию о нём и имя
func createBin() (bool, error) {
	cmd := flag.NewFlagSet("create", flag.ExitOnError)
	fileName := cmd.String("file", "bin.json", "Файл с JSON")
	binName := cmd.String("name", "my-bin", "Имя bin")
	cmd.Parse(os.Args[2:])
	fmt.Printf("Вызвана команда create с параметрами file: %s, и name: %s \n", *fileName, *binName)

	return true, nil
}

// Функционал:
//   - обновляет bin из файла по id
func updateBin() (bool, error) {
	cmd := flag.NewFlagSet("update", flag.ExitOnError)
	fileName := cmd.String("file", "bin.json", "Файл с JSON")
	id := cmd.String("id", "", "Идентификатор bin")
	cmd.Parse(os.Args[2:])
	fmt.Printf("Вызвана команда update с параметрами file: %s, и id: %s \n", *fileName, *id)

	return true, nil
}

// Функционал:
//   - удаляет bin по id из API и локального файла
func deleteBin() (bool, error) {
	cmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := cmd.String("id", "", "Идентификатор bin")
	cmd.Parse(os.Args[2:])
	fmt.Printf("Вызвана команда delete с параметром id: %s \n", *id)

	return true, nil
}

// Функционал:
//   - выводит в stdout bin из json.bin по его id
func getBin() (bool, error) {
	cmd := flag.NewFlagSet("get", flag.ExitOnError)
	id := cmd.String("id", "", "Идентификатор bin")
	cmd.Parse(os.Args[2:])
	fmt.Printf("Вызвана команда get с параметром id: %s \n", *id)

	return true, nil
}

// Функционал:
//   - выводит список id и name для bin из локального файла
func listBins() (bool, error) {
	cmd := flag.NewFlagSet("list", flag.ExitOnError)
	cmd.Parse(os.Args[2:])
	fmt.Println("Вызвана команда list")

	return true, nil
}
