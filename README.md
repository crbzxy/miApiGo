
# API en Go

Este proyecto es una API desarrollada en Go que utiliza Swaggo para generar documentación Swagger.

## Requisitos Previos

Asegúrate de tener Go instalado en tu máquina. Esta API requiere Go 1.15 o superior.

## Instalación

Clona el repositorio en tu máquina local usando:

```bash
git clone https://github.com/tu-usuario/tu-repositorio.git
cd tu-repositorio
```

Instala todas las dependencias necesarias:

```bash
go mod tidy
```

Genera la documentación Swagger:

```bash
swag init
```

## Ejecución

Para iniciar la API, ejecuta el siguiente comando en el directorio raíz del proyecto:

```bash
go run ./cmd/main.go
```

La API estará disponible en http://localhost:8080.

## Acceder a la Documentación Swagger

Una vez que la API esté ejecutándose, puedes acceder a la documentación Swagger en:

http://localhost:8080/swagger/index.html

## Despliegue

### En Raspberry Pi/Orange Pi con Docker y Portainer

**Docker** es una plataforma que permite empaquetar una aplicación y sus dependencias en un contenedor virtual que puede ejecutarse en cualquier sistema operativo Linux. **Portainer** es una herramienta de gestión para Docker que proporciona una interfaz gráfica para manejar contenedores fácilmente.

1. **Preparar el dispositivo:**
   - Asegúrate de que tu dispositivo esté configurado con un sistema operativo compatible que soporte Docker (como Raspbian para Raspberry Pi o Armbian para Orange Pi).
   - Instala Docker en el dispositivo. Puedes encontrar instrucciones específicas para Raspberry Pi en [Docker Hub](https://hub.docker.com/r/arm32v7/docker/).

2. **Crear un archivo `Dockerfile`:**
   - En el directorio raíz de tu proyecto, crea un `Dockerfile` que describa el entorno de tu aplicación. Aquí hay un ejemplo básico:

    ```dockerfile
    FROM golang:1.15

    WORKDIR /app

    COPY . .

    RUN go mod tidy
    RUN swag init

    CMD ["go", "run", "./cmd/main.go"]
    ```

3. **Construir la imagen Docker:**
   - Desde la terminal, ejecuta el siguiente comando para construir tu imagen Docker:

    ```bash
    docker build -t mi-api-go .
    ```

4. **Despliegue con Docker:**
   - Ejecuta el contenedor en tu Raspberry Pi o Orange Pi:

    ```bash
    docker run -d -p 8080:8080 mi-api-go
    ```

5. **Instalar y configurar Portainer (opcional):**
   - Si deseas una gestión más sencilla, puedes instalar Portainer para administrar tus contenedores Docker de forma gráfica:

    ```bash
    docker run -d -p 9000:9000 --name=portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock portainer/portainer-ce
    ```

   - Ahora puedes acceder a Portainer desde un navegador web en `http://<IP-de-tu-dispositivo>:9000` para gestionar tu aplicación.

## Contribuciones

Las contribuciones son bienvenidas. Por favor, lee CONTRIBUTING.md para más detalles sobre nuestro código de conducta, y el proceso para enviarnos pull requests.
