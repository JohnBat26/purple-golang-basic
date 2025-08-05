package main

import (
	"demo/3-bin/api"
	"demo/3-bin/config"
	"demo/3-bin/storage"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	c := config.NewConfig()

	if len(os.Args) == 1 {
		fmt.Println("Необходимо использовать хотя бы один параметр: --create, --update, --delete, --get или --list")
		return
	}

	s := storage.NewStorage()

	switch os.Args[1] {
	case "--create":
		_, err = api.CreateBin(c, s)
	case "--update":
		_, err = api.UpdateBin(c, s)
	case "--delete":
		_, err = api.DeleteBin(c, s)
	case "--get":
		_, err = api.GetBin(c, s)
	case "--list":
		_, err = api.ListBins(c, s)
	default:
		fmt.Println("Необходимо использовать хотя бы один параметр: --create, --update, --delete, --get или --list")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}
