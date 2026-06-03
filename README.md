# Empleados REST API

Gin + GORM REST API for employee management with SQLite.

## Endpoints

- `GET /empleados` — List all employees
- `GET /empleados/:id` — Get an employee
- `POST /empleados` — Create an employee
- `PATCH /empleados/:id` — Update an employee
- `DELETE /empleados/:id` — Delete an employee

## Run

```bash
go run main.go
```

## Test

```bash
go test ./...
```
