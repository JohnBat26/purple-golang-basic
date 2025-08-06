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

	apiService := api.NewService(api.NewRealHTTPClient(), c, s)

	switch os.Args[1] {
	case "--create":
		_, err = apiService.CreateBin()
	case "--update":
		_, err = apiService.UpdateBin()
	case "--delete":
		_, err = apiService.DeleteBin()
	case "--get":
		_, err = apiService.GetBin()
	case "--list":
		_, err = apiService.ListBins()
	default:
		fmt.Println("Необходимо использовать хотя бы один параметр: --create, --update, --delete, --get или --list")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}
