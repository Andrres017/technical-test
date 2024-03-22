# technical-test

# Proyecto Go (Nombre del Proyecto)

Este proyecto es una aplicación escrita en Go que hace XYZ.

## Requisitos Previos

Asegúrate de tener instalado Go en tu sistema. Puedes descargarlo desde [aquí](https://golang.org/dl/).

Además, necesitarás tener Docker instalado si deseas utilizar la base de datos en contenedor. Puedes descargar Docker desde [aquí](https://www.docker.com/products/docker-desktop).

## Configuración

Antes de ejecutar la aplicación, asegúrate de configurar las variables de entorno en un archivo `.env` en la raíz del proyecto. Aquí hay un ejemplo de archivo `.env`:

APP_PORT=3000
DB_USER=root
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=3307
DB_NAME=golang-test
OPENAI_API_KEY=sk-xxxx


## Ejecución

1. **Construir la aplicación:**


go build
./technical-test 

Uso del Docker de la Base de Datos
Si deseas utilizar la base de datos en un contenedor Docker, sigue estos pasos:

Levantar el contenedor de la base de datos:

docker-compose up -d

Detener y eliminar el contenedor de la base de datos:

docker-compose down
