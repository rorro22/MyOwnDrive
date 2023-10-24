package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/websocket"
	"github.com/mehanizm/airtable"
)

var validTokens = map[string]bool{
	"FG5D41G653DS14":     true,
	"15DSF64S3D1SFD5436": true,
}

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

	// Lee el token del cliente
	_, p, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Error al leer el mensaje:", err)
		return
	}

	token := string(p)
	if !validTokens[token] {
		fmt.Println("Token no válido")
		return
	}

	fmt.Println("Cliente conectado")

	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error al leer el mensaje:", err)
			return
		}

		if messageType == websocket.BinaryMessage {
			// Aquí se reciben los datos del archivo
			// Guardar el archivo en el servidor
			filename := "Docs/RCV_CV.docx" // Puedes definir un nombre fijo o extraerlo del mensaje del cliente
			err := saveFile(filename, data)
			if err != nil {
				fmt.Println("Error al guardar el archivo:", err)
				return
			}
			fmt.Println("Archivo recibido y guardado como:", filename)
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

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./Docs"))))
	http.HandleFunc("/", handleConnection)
	apiToken := "patEaDy8sH84VP66P.536412a0a3d4115c583997da1d404ad2d9f8fe9cb31f24942cf3875416294403"
	baseID := "appucocvVbrg17jpT"
	client := airtable.NewClient(apiToken)
	table := client.GetTable(baseID, "Usuarios")

	records, err := table.GetRecords().Do()
	if err != nil {
		// Maneja el error, por ejemplo, imprime un mensaje de error
		fmt.Println("Error al obtener registros:", err)
		return
	}

	for _, record := range records.Records {
		// Accede a los campos del registro
		field1 := record.Fields["Name"]
		field2 := record.Fields["Password"]

		// Realiza las operaciones que desees con los datos
		fmt.Printf("Field1: %v, Field2: %v\n", field1, field2)
	}

	recordsToSend := &airtable.Records{
		Records: []*airtable.Record{
			{
				Fields: map[string]any{
					"Name":     "lana",
					"Password": "LanaHuala",
				},
			},
		},
	}
	receivedRecords, err := table.AddRecords(recordsToSend)
	if err != nil {
		// Handle error
	}
	for _, record := range receivedRecords.Records {
		// Accede a los campos del registro
		field1 := record.Fields["Name"]
		field2 := record.Fields["Password"]

		// Realiza las operaciones que desees con los datos
		fmt.Printf("Field1: %v, Field2: %v\n", field1, field2)
	}

	fmt.Println("Servidor WebSocket iniciado en :8080")
	fallo := http.ListenAndServe(":8080", nil)
	if fallo != nil {
		panic("Error: " + fallo.Error())
	}
}
