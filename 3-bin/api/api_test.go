package api_test

import (
	"bytes"
	"demo/3-bin/api"
	"demo/3-bin/config"
	"demo/3-bin/storage"
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"strings"
	"testing"
)

type MockHTTPClient struct {
	Resp *http.Response
	Err  error
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.Resp, m.Err
}

// Сохраняем оригинальные аргументы и восстанавливаем после теста
func setupTest(command string, storageFile string, args []string) func() {
	oldArgs := os.Args
	os.Args = append([]string{"cmd", command}, args...)

	os.Setenv("STORAGE_FILENAME", storageFile)

	return func() {
		os.Args = oldArgs
		os.Remove(storageFile)
	}
}

func TestCreateBinOK(t *testing.T) {
	t.Parallel()
	storageFile := fmt.Sprintf("/tmp/storage_test_%d.json", rand.Int())
	cleanup := setupTest("create", storageFile, []string{"--file", "../testdata/data.json", "--name", "my-test-bin"})
	defer cleanup()

	data, err := os.ReadFile("../testdata/create_bin_ok_resp.json")

	if err != nil {
		t.Fatal("Failed to read test data:", err)
	}

	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service := api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	_, err = service.CreateBin()

	if err != nil {
		t.Error(err.Error())
	}
}

func TestCreateBinFAIL(t *testing.T) {
	t.Parallel()
	storageFile := fmt.Sprintf("/tmp/storage_test_%d.json", rand.Int())
	cleanup := setupTest("create", storageFile, []string{"--file", "../testdata/data.json", "--name", "my-test-bin"})
	defer cleanup()

	mockResp := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(strings.NewReader(`{"message": "Invalid JSON. Please try again"}`)),
	}

	service := api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	_, err := service.CreateBin()

	if !errors.Is(err, api.ErrCreateFail) {
		t.Errorf("Ожидалось: %v, получено: %v", api.ErrCreateFail, err)
	}
}

func TestUpdateBinOK(t *testing.T) {
	t.Parallel()
	storageFile := fmt.Sprintf("/tmp/storage_test_%d.json", rand.Int())
	cleanup := setupTest("create", storageFile, []string{"--file", "../testdata/data.json", "--name", "my-test-bin"})
	defer cleanup()

	data, err := os.ReadFile("../testdata/create_bin_ok_resp.json")

	if err != nil {
		t.Fatal("Failed to read test data:", err)
	}

	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service := api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	_, err = service.CreateBin()
	if err != nil {
		t.Error(err.Error())
	}

	cleanup = setupTest("update", storageFile, []string{"--file", "testdata/data.json", "--id", service.LatestMetadata.ID})

	data, err = os.ReadFile("../testdata/update_bin_ok_resp.json")
	mockResp = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service = api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	result, err := service.UpdateBin()

	if err != nil {
		t.Error(err.Error())
	}

	if !result {
		t.Error(err.Error())
	}
}

func TestUpdateBinFAIL(t *testing.T) {
	t.Parallel()
	storageFile := fmt.Sprintf("/tmp/storage_test_%d.json", rand.Int())
	cleanup := setupTest("create", storageFile, []string{"--file", "../testdata/data.json", "--name", "my-test-bin"})
	defer cleanup()

	data, err := os.ReadFile("../testdata/create_bin_ok_resp.json")

	if err != nil {
		t.Fatal("Failed to read test data:", err)
	}

	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service := api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	_, err = service.CreateBin()
	if err != nil {
		t.Error(err.Error())
	}

	cleanup = setupTest("update", storageFile, []string{"--file", "testdata/data.json", "--id", "1231231231231"})

	data, err = os.ReadFile("../testdata/update_bin_ok_resp.json")
	mockResp = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service = api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	result, err := service.UpdateBin()

	if result {
		t.Error(err.Error())
	}
}
func TestGetBinOK(t *testing.T) {
	t.Parallel()
	storageFile := fmt.Sprintf("/tmp/storage_test_%d.json", rand.Int())
	cleanup := setupTest("create", storageFile, []string{"--file", "../testdata/data.json", "--name", "my-test-bin"})
	defer cleanup()

	data, err := os.ReadFile("../testdata/create_bin_ok_resp.json")

	if err != nil {
		t.Fatal("Failed to read test data:", err)
	}

	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service := api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	_, err = service.CreateBin()
	if err != nil {
		t.Error(err.Error())
	}

	cleanup = setupTest("update", storageFile, []string{"--id", service.LatestMetadata.ID})

	data, err = os.ReadFile("../testdata/create_bin_ok_resp.json")
	mockResp = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service = api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	result, err := service.GetBin()

	if err != nil {
		t.Error(err.Error())
	}

	if !result {
		t.Error(err.Error())
	}
}

func TestGetBinFAIL(t *testing.T) {
	t.Parallel()
	storageFile := fmt.Sprintf("/tmp/storage_test_%d.json", rand.Int())
	cleanup := setupTest("create", storageFile, []string{"--file", "../testdata/data.json", "--name", "my-test-bin"})
	defer cleanup()

	data, err := os.ReadFile("../testdata/create_bin_ok_resp.json")

	if err != nil {
		t.Fatal("Failed to read test data:", err)
	}

	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service := api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	_, err = service.CreateBin()
	if err != nil {
		t.Error(err.Error())
	}

	cleanup = setupTest("update", storageFile, []string{"--id", "12312323123"})
	defer cleanup()

	data, err = os.ReadFile("../testdata/create_bin_ok_resp.json")
	mockResp = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service = api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	result, err := service.GetBin()

	if result {
		t.Error(err.Error())
	}
}

func TestDeleteBinOK(t *testing.T) {
	t.Parallel()
	storageFile := fmt.Sprintf("/tmp/storage_test_%d.json", rand.Int())
	cleanup := setupTest("create", storageFile, []string{"--file", "../testdata/data.json", "--name", "my-test-bin"})
	defer cleanup()

	data, err := os.ReadFile("../testdata/create_bin_ok_resp.json")

	if err != nil {
		t.Fatal("Failed to read test data:", err)
	}

	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service := api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	_, err = service.CreateBin()
	if err != nil {
		t.Error(err.Error())
	}

	cleanup = setupTest("update", storageFile, []string{"--id", service.LatestMetadata.ID})

	data, err = os.ReadFile("../testdata/create_bin_ok_resp.json")
	mockResp = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service = api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	result, err := service.DeleteBin()

	if err != nil {
		t.Error(err.Error())
	}

	if !result {
		t.Error(err.Error())
	}
}

func TestDeleteBinFAIL(t *testing.T) {
	t.Parallel()
	storageFile := fmt.Sprintf("/tmp/storage_test_%d.json", rand.Int())
	cleanup := setupTest("create", storageFile, []string{"--file", "../testdata/data.json", "--name", "my-test-bin"})
	defer cleanup()

	data, err := os.ReadFile("../testdata/create_bin_ok_resp.json")

	if err != nil {
		t.Fatal("Failed to read test data:", err)
	}

	mockResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service := api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	_, err = service.CreateBin()
	if err != nil {
		t.Error(err.Error())
	}

	cleanup = setupTest("update", storageFile, []string{"--id", "12312323123"})
	defer cleanup()

	data, err = os.ReadFile("../testdata/create_bin_ok_resp.json")
	mockResp = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(data)),
	}

	service = api.NewService(&MockHTTPClient{
		Resp: mockResp,
		Err:  nil,
	}, config.NewConfig(), storage.NewStorage())

	result, err := service.DeleteBin()

	if result {
		t.Error(err.Error())
	}
}
