# My Own Drive - Documentación del Proyecto

## Descripción

My Own Drive es una aplicación de almacenamiento en la nube que ofrece una amplia gama de características, incluyendo el procesamiento de archivos, autenticación y autorización, sistema de archivos jerárquico, recuperación de versiones anteriores, búsqueda de archivos y metadatos, sincronización de archivos, escalabilidad, rendimiento, redundancia y seguridad. La aplicación cuenta con una interfaz de usuario atractiva y proporciona una solución integral para el almacenamiento y gestión de archivos.

## Características Destacadas

### Lenguajes y Tecnologías Utilizados

- **GOLang:** Para el desarrollo del servidor.
- **JavaScript:** Para la programación del cliente.
- **WebSockets:** Utilizados para la comunicación en tiempo real.
- **Airtable:** Base de datos utilizada para almacenar datos de usuarios y permisos.
- **Otras tecnologías o bibliotecas:** Puedes listar aquí cualquier otra tecnología relevante utilizada en tu proyecto.

### Almacenamiento de Archivos
My Own Drive permite a los usuarios cargar, descargar y organizar archivos en un sistema de archivos jerárquico. Los archivos son almacenados de forma segura y se pueden acceder desde cualquier lugar en cualquier momento.

### Recuperación de Versiones Anteriores
La aplicación guarda versiones anteriores de los archivos, permitiendo a los usuarios restaurar o consultar versiones anteriores de un archivo específico.

### Búsqueda Avanzada
My Own Drive proporciona potentes capacidades de búsqueda para ayudar a los usuarios a encontrar rápidamente los archivos y metadatos que necesitan.

### Sincronización en Tiempo Real
Los cambios realizados en un dispositivo se sincronizan automáticamente con otros dispositivos, lo que garantiza que los archivos estén siempre actualizados.

### Seguridad y Autenticación
La aplicación utiliza un sistema de autenticación con tokens y autenticación de usuario y contraseña para garantizar que solo usuarios autorizados puedan acceder a sus datos.

### Escalabilidad y Redundancia
My Own Drive está diseñado para ser escalable y redundante, lo que garantiza un alto rendimiento y disponibilidad.

## Comunicación mediante WebSockets

La aplicación utiliza WebSockets para establecer una comunicación en tiempo real entre el servidor y el cliente. Esto permite la carga de archivos, sincronización de cambios y notificaciones de eventos importantes. La comunicación es segura y cifrada para proteger la privacidad y la integridad de los datos.

### Autenticación con Token
Cuando un cliente intenta conectarse al servidor, debe proporcionar un token de autenticación válido.

### Autenticación con Usuario y Contraseña
Después de la autenticación exitosa, el servidor solicita al cliente que proporcione su nombre de usuario y contraseña.

### Comunicación de Datos
La comunicación es bidireccional, permitiendo a los clientes cargar archivos, recibir notificaciones y mantener sus datos sincronizados en tiempo real.

## Contribuciones

Si deseas contribuir al proyecto, eres bienvenido. Puedes contribuir mediante la creación de issues, sugerencias de mejoras o enviando pull requests. Para más detalles, consulta el archivo [CONTRIBUTING.md](CONTRIBUTING.md).

## Licencia

Este proyecto está bajo la licencia [MIT](LICENSE). Puedes consultar el archivo [LICENSE](LICENSE) para obtener más detalles.

## Autores

- [Rodrigo V.C.](https://github.com/rorro22) - Desarrollador principal

## Contacto

Para obtener más información sobre el proyecto, puedes contactar al equipo en [rodrivaldesc@hotmail.com](mailto:rodrivaldesc@hotmail.com).

---
My Own Drive - 2023
