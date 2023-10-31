package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type FilePayload struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var payload FilePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	data, err := base64.StdEncoding.DecodeString(payload.Data)
	if err != nil {
		http.Error(w, "Failed to decode base64 data", http.StatusBadRequest)
		return
	}

	// Ensure the "docs" directory exists
	if err := os.MkdirAll("docs", 0755); err != nil {
		http.Error(w, "Failed to create the docs directory", http.StatusInternalServerError)
		return
	}

	// Construct the full file path
	filePath := filepath.Join("docs", payload.Filename)

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		http.Error(w, "Failed to save the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "File uploaded and saved!")
}
