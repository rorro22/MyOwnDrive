package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type FileDetail struct {
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
}

func FetchFilesHandler(w http.ResponseWriter, r *http.Request) {
	// ACTUALIZAR FUNCIONALIDAD DE REGRESAR DOCUMENTOS PARA
	//		USAR NUEVO LIBRERIA GO.
	//		REGRESAR CORRECTAMENTE LA EXTENSION DEL DOCUMENTO.
	//		OBTENER Y REGRESAR LOS DATOS BINARIOS DEL DOCUMENTO (revisar si en GO podemos usar archivos Blob para facilitar esto).
	files, err := ioutil.ReadDir("./docs")
	if err != nil {
		http.Error(w, "Error reading files", http.StatusInternalServerError)
		return
	}

	var fileDetails []FileDetail
	for _, file := range files {
		fileDetails = append(fileDetails, FileDetail{
			FileName: file.Name(),
			FileType: filepath.Ext(file.Name()),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fileDetails)
}
