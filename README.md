# 🧰 Complete REST API in Go

## 🚧 Development Setup

> ⚠️ **Important:** Before running the application, make sure to set up the database by following the **[Database Migrations](#database-migrations)** section below.

---

## 🔄 Using Air for Live Reload

This project uses [**Air**](https://github.com/cosmtrek/air) for live reload during development.

### ➤ Install Air
```bash
go install github.com/cosmtrek/air@latest
```

### ➤ Run the application with Air
```bash
air
```

Air will automatically:
- Watch your Go files for changes
- Rebuild your application on file changes
- Restart your server automatically
- Show build errors with colorized output

> To stop Air, press `Ctrl+C` in your terminal.

### ✅ Air Configuration
The project includes a pre-configured `.air.toml` file that:
- Watches the `cmd/api` directory
- Excludes test files and common directories
- Includes Go files, templates, and HTML files
- Uses colorized output for better visibility

---

## 🌐 Environment Variables

You can optionally create a `.env` file for local development:

```env
BASE_URL=http://localhost:8080
PORT=8080
JWT_SECRET=your-secret-key
```

For **production**, configure these via your deployment platform’s environment settings.

---

## ▶️ Running Without Air

If you prefer not to use Air, run the application manually:

```bash
go run ./cmd/api
```

The server will start at: [http://localhost:8080](http://localhost:8080)

> You’ll need to restart manually when code changes.

---

## 📦 Dependencies

This project uses **Go modules** for dependency management. Dependencies are automatically handled when you build or run the app.

To download them explicitly:
```bash
go mod download
```

---

## 🗄️ Database Migrations

This project uses [**golang-migrate**](https://github.com/golang-migrate/migrate) for managing database schema.

### ➤ Install Migrate CLI
See the instructions here:  
[golang-migrate CLI install guide](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

### ⚠️ Run Migrations Before First Launch
```bash
go run ./cmd/migrate up
```

### ⬅️ Rollback Migrations
```bash
go run ./cmd/migrate down
```

### 🆕 Create New Migrations
```bash
migrate create -ext sql -dir ./cmd/migrate/migrations -seq name_of_migration
```

This creates two new files:
- `{timestamp}_name_of_migration.up.sql`
- `{timestamp}_name_of_migration.down.sql`

---

## 🛠️ Building the Application

To compile the application:

```bash
go build -o api ./cmd/api
```

This creates an `api` executable in the project root.

---

## 🚀 Running the Application

After building, run the app with:

```bash
./api
```

By default, it starts at: [http://localhost:8080](http://localhost:8080)

---

## 📘 API Documentation (Swagger)

Generate Swagger docs using:

```bash
swag init --dir cmd/api --parseDependency --parseInternal --parseDepth 1
```

Then view it at:  
[http://localhost:8080/swagger](http://localhost:8080/swagger)

Swagger UI provides:
- A full list of API endpoints
- Request and response schema documentation
- **Try-it-out** functionality
- Authentication and headers support