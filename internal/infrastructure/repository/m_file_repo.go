package repository

import (
	"bytes"
	"fiber-crud-demo/internal/domain"
	"fiber-crud-demo/util"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
	"time"
)

type MFileRepositoryImpl struct {
	mutex sync.Mutex
}

func NewMFileRepository() domain.MFileRepository {
	return &MFileRepositoryImpl{}
}

// Stream implements domain.MFileRepository.
func (m *MFileRepositoryImpl) Stream(id uint) (io.ReadCloser, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%d", os.Getenv("FILE_MANAGEMENT_HOST"), "m-file/file/stream", id))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", resp.Status)
	}
	return resp.Body, nil
}

// Upload implements domain.MFileRepository.
func (m *MFileRepositoryImpl) Upload(file *multipart.FileHeader, id int64) error {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return err
	}

	fileStream, err := file.Open()
	if err != nil {
		return err
	}
	defer fileStream.Close()

	_, err = io.Copy(part, fileStream)
	if err != nil {
		return err
	}

	writer.WriteField("user_id", fmt.Sprintf("%d", 0))
	writer.WriteField("module_id", os.Getenv("MODULE_ID"))
	writer.WriteField("id", fmt.Sprintf("%d", id))

	err = writer.Close()
	if err != nil {
		return err
	}

	url := os.Getenv("FILE_MANAGEMENT_HOST")
	url = url + "/m-file/file"

	req, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		util.Log("ERROR", "repository", "upload", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload file: %s", resp.Status)
	}

	return nil
}

func (m *MFileRepositoryImpl) Delete(id string) error {
	url := os.Getenv("FILE_MANAGEMENT_HOST") + "/m-file/file/" + id

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Resource deleted successfully.")
	} else {
		fmt.Printf("Failed to delete resource. Status: %s\n", resp.Status)
	}
	return nil
}
