# ShopSmart Microservices (Go + Fiber + PostgreSQL)

Proyecto de ejemplo con arquitectura de microservicios usando Go, Fiber, GORM y PostgreSQL.

Incluye orquestación completa con Docker Compose para levantar todas las APIs y la base de datos en un solo comando.

## Arquitectura

- users-service: gestión de usuarios
- inventory-service: gestión de productos e inventario
- orders-service: gestión de órdenes
- infra-postgres: base de datos PostgreSQL con inicialización automática de esquemas y tablas

Cada servicio se ejecuta de forma independiente y comparte la misma base de datos `shopdb`, separando datos por schema:

- `users.users`
- `inventory.products`
- `orders.orders`

## Stack

- Go 1.22.2
- Fiber v2
- GORM + driver PostgreSQL
- PostgreSQL 16 (Docker)

## Estructura del Proyecto

```text
.
├── docker-compose.yml
├── go.mod
├── infra-postgres/
│   ├── docker-compose.yml
│   └── init.sql
├── users-service/
│   ├── Dockerfile
│   ├── main.go
│   ├── database.go
│   └── handler.go
├── inventory-service/
│   ├── Dockerfile
│   ├── main.go
│   ├── database.go
│   └── handler.go
└── orders-service/
  ├── Dockerfile
    ├── main.go
    ├── database.go
    └── handler.go
```

## Requisitos

- Go 1.22+ instalado
- Docker y Docker Compose instalados
- Puerto `5432` libre para PostgreSQL
- Puertos `3001`, `3002`, `3003` libres para los microservicios

## Levantar el Proyecto

### Opción recomendada: Todo con Docker Compose

Desde la raíz del proyecto:

```bash
docker compose up -d --build
```

Esto levanta:

- PostgreSQL (`shop-postgres`)
- users-service (`http://localhost:3001`)
- inventory-service (`http://localhost:3002`)
- orders-service (`http://localhost:3003`)

Ver logs:

```bash
docker compose logs -f
```

Detener todo:

```bash
docker compose down
```

Detener y borrar datos de la DB:

```bash
docker compose down -v
```

### Opción alternativa: Servicios en local + DB en Docker

#### 1. Iniciar PostgreSQL

Desde la raíz del proyecto:

```bash
docker compose -f infra-postgres/docker-compose.yml up -d
```

Esto crea:

- Base de datos `shopdb`
- Extensión `uuid-ossp`
- Schemas `users`, `inventory`, `orders`
- Tablas iniciales definidas en `infra-postgres/init.sql`

#### 2. Descargar dependencias de Go

```bash
go mod tidy
```

### 3. Ejecutar cada microservicio (en terminales separadas)

Users service:

```bash
go run ./users-service
```

Inventory service:

```bash
go run ./inventory-service
```

Orders service:

```bash
go run ./orders-service
```

## Puertos

- users-service: `http://localhost:3001`
- inventory-service: `http://localhost:3002`
- orders-service: `http://localhost:3003`

## Conexión para API Gateway Externo

Si tu API Gateway está fuera de este compose:

- En host local (gateway corriendo en tu máquina):
  - `http://localhost:3001` (users)
  - `http://localhost:3002` (inventory)
  - `http://localhost:3003` (orders)

- En otro stack Docker:
  1. Conecta el contenedor del gateway a la red `shopsmart-network`.
  2. Usa estos upstreams por nombre de servicio:
     - `http://users-service:3001`
     - `http://inventory-service:3002`
     - `http://orders-service:3003`

Ejemplo para conectar un contenedor gateway ya creado a la red:

```bash
docker network connect shopsmart-network <nombre-contenedor-gateway>
```

## API por Servicio

### Users Service (`:3001`)

- `POST /users`
- `GET /users/:id`
- `GET /users`

Ejemplo crear usuario:

```bash
curl -X POST http://localhost:3001/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Eduardo","email":"eduardo@example.com"}'
```

Ejemplo listar usuarios:

```bash
curl http://localhost:3001/users
```

### Inventory Service (`:3002`)

- `POST /products`
- `GET /products`

Ejemplo crear producto:

```bash
curl -X POST http://localhost:3002/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","stock":15}'
```

Ejemplo listar productos:

```bash
curl http://localhost:3002/products
```

### Orders Service (`:3003`)

- `POST /orders`
- `GET /orders/:id`

Ejemplo crear orden:

```bash
curl -X POST http://localhost:3003/orders \
  -H "Content-Type: application/json" \
  -d '{"user_id":"00000000-0000-0000-0000-000000000001"}'
```

Ejemplo obtener orden por ID:

```bash
curl http://localhost:3003/orders/<order_id>
```

## Configuración de Base de Datos

Los tres servicios leen configuración por variables de entorno, con defaults para desarrollo local:

- `DB_HOST` (default: `localhost`)
- `DB_PORT` (default: `5432`)
- `DB_USER` (default: `postgres`)
- `DB_PASSWORD` (default: `postgres`)
- `DB_NAME` (default: `shopdb`)
- `DB_SSLMODE` (default: `disable`)

Si quieres usar otra configuración, define esas variables de entorno o ajusta los defaults en:

- `users-service/database.go`
- `inventory-service/database.go`
- `orders-service/database.go`

## Detener Servicios

Si los ejecutaste con `go run`, detén cada proceso con `Ctrl + C`.

Para detener PostgreSQL:

```bash
docker compose -f infra-postgres/docker-compose.yml down
```

Para detener y eliminar volumen de datos:

```bash
docker compose -f infra-postgres/docker-compose.yml down -v
```

## Notas

- El API Gateway externo se puede conectar por puertos publicados o por la red Docker `shopsmart-network`.
- No hay autenticación/autorización implementada.
- No hay validaciones avanzadas ni manejo de errores robusto en todos los handlers.
