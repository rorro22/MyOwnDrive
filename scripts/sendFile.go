package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

func sendFileToAPI(filename string, endpoint string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	encodedData := base64.StdEncoding.EncodeToString(data)

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer([]byte(fmt.Sprintf(`{"filename": "%s", "data": "%s"}`, filename, encodedData))))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}

func main() {
	filename := "RCV_CV.docx"                  // replace with your file name
	endpoint := "http://localhost:8080/upload" // replace with your API endpoint
	err := sendFileToAPI(filename, endpoint)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
