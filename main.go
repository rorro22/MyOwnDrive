package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error durante la actualización de la conexión:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Cliente conectado")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error al leer el mensaje:", err)
			return
		}

		if messageType == websocket.TextMessage {
			// El cliente envía un objeto JSON con el nombre, extensión y datos del archivo
			var fileData map[string]interface{}
			if err := json.Unmarshal(p, &fileData); err != nil {
				fmt.Println("Error al decodificar el objeto JSON:", err)
				return
			}

			fmt.Printf("Tipo de dato de fileData: %T\n", fileData)
			fmt.Printf("Tipo de dato de p: %T\n", p)

			name, _ := fileData["name"].(string)
			extension, _ := fileData["extension"].(string)
			data, _ := fileData["data"].([]byte)
			fmt.Printf("Tamaño de datos recibidos: %d bytes\n", len(data))
			fmt.Println(name, extension, data)

			if name != "" && extension != "" && data != nil {
				// Guardar el archivo en el servidor con el nombre y extensión originales
				filename := "Docs/" + name + "." + extension
				err := saveFile(filename, data)
				if err != nil {
					fmt.Println("Error al guardar el archivo:", err)
					return
				}
				fmt.Println("Archivo recibido y guardado como:", filename)
			} else {
				fmt.Println("Datos de archivo incompletos")
			}
		}
	}
}

func saveFile(filename string, data []byte) error {
	// Obtener el directorio donde se guardará el archivo
	directory := filepath.Dir(filename)

	// Crear el directorio si no existe
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return err
	}

	// Crear el archivo en el directorio especificado
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Escribir los datos en el archivo
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func extractFileNameAndExtension(contentDisposition string) (string, string) {
	parts := strings.Split(contentDisposition, "filename=")
	if len(parts) != 2 {
		return "", ""
	}
	filename := strings.Trim(parts[1], "\"")
	extParts := strings.Split(filename, ".")
	if len(extParts) < 2 {
		return "", ""
	}
	name := strings.Join(extParts[:len(extParts)-1], ".")
	extension := extParts[len(extParts)-1]
	return name, extension
}

func main() {
	http.HandleFunc("/", handleConnection)

	fmt.Println("Servidor WebSocket iniciado en :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
