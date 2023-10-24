package handlers

import (
	"encoding/json"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Aquí puedes validar el nombre de usuario y la contraseña contra una base de datos o cualquier otro mecanismo de tu elección.
	// ACTUALIZAR FUNCIONALIDAD DE AUTENTICACIÓN
	//		CONECTANDOSE A AIRTABLE Y COMPARANDO CREDENCIALES.
	//		DE SER POSIBLE ESCOGER UN ALGORITMO DE ENCRIPTACION PARA
	//			COMPARTIR CON EL CLIENTE Y QUE MANDEN LAS CONTRASEÑAS ENCRIPTADAS.
	//		TAL VEZ REGRESAR UN TOKEN DE AUTENTICACION PARA QUE USE EL CLIENTE
	//			EN EL RESTO DE PETICIONES AL SERVER COMO FORMA DE AUTENTICACION.
	if creds.Username == "test" && creds.Password == "password123" {
		w.Write([]byte("Login successful"))
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
