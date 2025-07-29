package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	if !strings.HasSuffix(name, ".json") {
		return nil, errors.New("Файл должен иметь расширение json")
	}

	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, err
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(content)

	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Запись успешна")
}
