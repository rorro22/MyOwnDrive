const WebSocket = require('ws');
const fs = require('fs');
const path = require('path'); // Agrega esta línea

const socket = new WebSocket('ws://localhost:8080');

socket.on('open', function open() {
    console.log('Conectado al servidor WebSocket');
    
    // Leer el archivo y obtener su nombre y extensión
    const filePath = 'C:/Users/Suspect/Desktop/GOLang/servidor/RCV_CV.docx'; // Reemplaza con la ruta de tu archivo
    const data = fs.readFileSync(filePath);
    console.log('Datos leídos del archivo:', data);
    const fileName = path.basename(filePath); // Obtiene el nombre del archivo
    const fileExtension = path.extname(filePath).slice(1); // Obtiene la extensión del archivo
    
    // Crear un objeto para enviar tanto los datos como la información del archivo
    const fileData = {
        name: fileName,
        extension: fileExtension,
        data: Array.from(data)
    };
    
    // Enviar los datos como un mensaje binario al servidor
    socket.send(JSON.stringify(fileData));
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
