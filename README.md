# 🚌 API Chombi

API REST para la aplicación móvil **Chombi**, diseñada para gestionar los turnos y registros de combis (autobuses) cuando llegan a terminales específicas. Esta API proporciona funcionalidades completas de autenticación, gestión de usuarios y administración de vehículos con históricos detallados.

## 📋 Descripción del Proyecto

Chombi es una solución integral para operadores de transporte que necesitan rastrear y gestionar eficientemente sus vehículos en terminales. La API permite registrar arrivals, documentar históricos de viajes, gestionar la información de vehículos y mantener un registro detallado de operaciones.

## ✨ Características Implementadas

### Autenticación y Usuarios
- ✅ **Registro de Usuarios** - Crear nuevas cuentas con email y contraseña encriptada
- ✅ **Login de Usuarios** - Autenticación segura con JWT
- ✅ **Middleware de Autenticación** - Protección de endpoints con tokens JWT

### Gestión de Vehículos
- ✅ **Registrar Vehículos** - Crear nuevos registros de combis
- ✅ **Listar Vehículos** - Obtener todos los vehículos registrados
- ✅ **Obtener Vehículo por Número de Unidad** - Búsqueda específica de combis
- ✅ **Actualizar Información de Vehículos** - Modificar datos de vehículos existentes
- ✅ **Eliminar Vehículos** - Remover registros de vehículos

### Gestión de Históricos y Turnos
- ✅ **Registrar Historiales de Turnos** - Documentar llegadas a terminales
- ✅ **Obtener Histórico de Vehículos** - Ver todos los turnos registrados
- ✅ **Obtener Histórico por Fecha** - Filtrar turnos por rango de fechas
- ✅ **Carga de Documentos PDF** - Subir archivos a Cloudinary

### Infraestructura
- ✅ **Base de Datos MySQL** - Persistencia de datos
- ✅ **CORS Configurado** - Soporte para múltiples orígenes
- ✅ **Cloudinary Integration** - Almacenamiento de archivos en la nube

## 🛠 Requisitos Previos

- **Go 1.25.0** o superior
- **MySQL 8.0** o superior
- **Cloudinary Account** (para subir PDFs)
- Terminal/CMD con acceso a comandos

## 📦 Instalación

### macOS

```bash
# 1. Descargar Go desde https://golang.org/dl/
# O usando Homebrew:
brew install go

# 2. Verificar instalación
go version

# 3. Clonar el proyecto
git clone https://github.com/bchanona/Api_chombi.git
cd Api_chombi

# 4. Descargar dependencias
go mod download

# 5. Instalar dependencias
go mod tidy
```

### Linux (Ubuntu/Debian)

```bash
# 1. Actualizar paquetes
sudo apt update

# 2. Instalar Go
sudo apt install golang-go

# 3. Verificar instalación
go version

# 4. Clonar el proyecto
git clone https://github.com/bchanona/Api_chombi.git
cd Api_chombi

# 5. Descargar dependencias
go mod download
go mod tidy
```

### Windows

```bash
# 1. Descargar Go desde https://golang.org/dl/
# Instalar el archivo .msi

# 2. Verificar instalación (en PowerShell o CMD)
go version

# 3. Clonar el proyecto
git clone https://github.com/bchanona/Api_chombi.git
cd Api_chombi

# 4. Descargar dependencias
go mod download
go mod tidy
```

## ⚙️ Configuración

### Configurar Base de Datos

Es **IMPRESCINDIBLE** crear la base de datos antes de ejecutar la aplicación. Ejecuta el script SQL ubicado en `database/schema.sql`:

```bash
# Conectarse a MySQL
mysql -u root -p

# Ejecutar el script dentro de MySQL
mysql -u root -p < database/schema.sql
```

O también puedes copiar y ejecutar el contenido del archivo [database/schema.sql](database/schema.sql) directamente en tu cliente MySQL (MySQL Workbench, phpMyAdmin, etc.)

> **Nota**: Este script creará la base de datos `chombidb` con todas las tablas necesarias (Roles, Users, Vehicles, vehicle_shift_history)

### Variables de Entorno

Configura las siguientes variables de entorno:

```bash
# Base de Datos MySQL
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=tu_contraseña
DB_NAME=chombidb
DB_PORT=3306

# Cloudinary (opcional, para subida de PDFs)
CLOUDINARY_CLOUD_NAME=tu_cloud_name
CLOUDINARY_API_KEY=tu_api_key
CLOUDINARY_API_SECRET=tu_api_secret
```

## 🚀 Cómo Correr el Programa

### Ejecución Local

```bash
# En la raíz del proyecto
go run main.go
```

El servidor iniciará en `http://localhost:8080`

### Compilar a Binario

```bash
# Compilar para tu sistema operativo actual
go build -o chombi-api main.go

# Ejecutar el binario
./chombi-api
```

### Compilar para Otros Sistemas Operativos

```bash
# Para macOS (desde Linux o Windows)
GOOS=darwin GOARCH=amd64 go build -o chombi-api-mac main.go

# Para Linux (desde macOS o Windows)
GOOS=linux GOARCH=amd64 go build -o chombi-api-linux main.go

# Para Windows (desde macOS o Linux)
GOOS=windows GOARCH=amd64 go build -o chombi-api.exe main.go
```

## 📡 Endpoints API

La API está organizada en la ruta base `/api/v2`

### Autenticación (Sin protección)

```
POST   /api/v2/auth/register          - Registrar nuevo usuario
POST   /api/v2/auth/login             - Login de usuario (retorna JWT)
```

### Vehículos (Requieren autenticación JWT)

```
POST   /api/v2/vehicles/              - Crear nuevo vehículo
GET    /api/v2/vehicles/              - Listar todos los vehículos del usuario
GET    /api/v2/vehicles/by-unit/:unitNumber  - Obtener vehículo por número de unidad
PUT    /api/v2/vehicles/:vehicleId    - Actualizar vehículo
DELETE /api/v2/vehicles/:vehicleId    - Eliminar vehículo
```

### Históricos de Turnos (Requieren autenticación JWT)

```
POST   /api/v2/vehicles/history       - Registrar nuevo turno/llegada a terminal
GET    /api/v2/vehicles/history       - Obtener histórico de vehículos
GET    /api/v2/vehicles/history/by-date/:date  - Obtener histórico por fecha
```

### Documentos (Requieren autenticación JWT)

```
POST   /api/v2/vehicles/upload-pdf    - Subir documento PDF
```

## 🏗 Estructura del Proyecto

```
Api_chombi/
├── main.go                 # Punto de entrada de la aplicación
├── go.mod                  # Definiciones de módulos
├── README.md              # Este archivo
├── database/
│   └── schema.sql         # Script SQL para crear la base de datos
└── src/
    ├── helper/            # Utilidades compartidas
    │   ├── config/        # Configuración (BD, CORS, Cloudinary)
    │   ├── middlewares/   # Middleware de autenticación
    │   └── services/      # Servicios (login, JWT, UUID)
    ├── users/             # Módulo de Usuarios
    │   ├── application/   # Casos de uso
    │   ├── domain/        # Entidades y repositorios
    │   ├── infrastructure/# Controladores, rutas, BD
    │   └── controllers/   # APIs endpoints
    └── vehicles/          # Módulo de Vehículos
        ├── application/   # Casos de uso
        ├── domain/        # Entidades y repositorios
        ├── infrastructure/# Controladores, rutas, BD
        └── controllers/   # APIs endpoints
```

## 🔐 Seguridad

- Las contraseñas se almacenan encriptadas
- Todos los endpoints de vehículos están protegidos con JWT
- CORS está configurado para controlar accesos desde diferentes orígenes
- Tokens JWT con expiración configurada

## 📚 Stack Tecnológico

- **Lenguaje**: Go 1.25.0
- **Framework Web**: Gin Gonic
- **Base de Datos**: MySQL
- **Autenticación**: JWT (JSON Web Tokens)
- **Almacenamiento en Nube**: Cloudinary
- **Seguridad**: Encriptación de contraseñas, CORS

## 👨‍💻 Autor

- **Briyan Chanona** - Desarrollador Principal

## 📝 Licencia

Este proyecto está bajo licencia privada.

---

**Nota**: Asegúrate de configurar correctamente la base de datos MySQL y las credenciales de Cloudinary antes de ejecutar la aplicación.
