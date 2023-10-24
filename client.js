const WebSocket = require('ws');
const fs = require('fs');
const path = require('path');

const socket = new WebSocket('ws://localhost:8080');

socket.on('open', function open() {
    console.log('Conectado al servidor WebSocket');

    const token = 'FG5D41G653DS14';

    // Enviar usuario y contraseña al servidor
    socket.send(JSON.stringify({ token }));

    const filePath = 'C:/Users/Suspect/Desktop/GOLang/servidor/RCV_CV.docx';
    const data = fs.readFileSync(filePath);
    console.log('Datos leídos del archivo:', data);

    // Enviar los datos binarios al servidor
    socket.send(data);

    console.log('Archivo enviado al servidor');
});

socket.on('message', function incoming(message) {
    console.log('Recibido del servidor:', message);
});

socket.on('close', function close(event) {
    if (event.wasClean) {
        console.log(`Conexión cerrada limpiamente, código=${event.code}, razón=${event.reason}`);
    } else {
        console.error('La conexión se perdió');
    }
});

socket.on('error', function error(err) {
    console.error('Error:', err.message);
});
