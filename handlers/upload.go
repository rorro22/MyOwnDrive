package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type FileUpload struct {
	FileName    string `json:"fileName"`
	FileType    string `json:"fileType"`
	FileContent []byte `json:"fileContent"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var fileUpload FileUpload
	err := json.NewDecoder(r.Body).Decode(&fileUpload)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// ACTUALIZAR FUNCIONALIDAD DE GUARDAR DOCUMENTOS PARA
	//		USAR NUEVO LIBRERIA GO.
	//		GUARDAR CORRECTAMENTE LOS DOCUMENTOS (hacer un cliente de prueba o peticiones postman).
	//		GUARDAR LOS DOCUMENTOS DENTRO DEL FOLDER DEL USUARIO "./docs/JorgeNava/".
	err = ioutil.WriteFile("./docs/"+fileUpload.FileName, fileUpload.FileContent, 0644)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File uploaded successfully"))
}
