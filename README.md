# goLift 🏋🏻

An app for both lifters and coaches. Create workouts for your athletes and yourself; and keep all the data centralized in one place.

## Tech Stack

### Backend
- Go
- PostgreSQL
- Chi Router
- JWT Authentication

### Frontend (Coming Soon)
- Next.js
- Tailwind CSS
- shadcn/ui

## Getting Started

### Prerequisites
- Go 1.21+
- PostgreSQL 17
- Node.js (for frontend)

### Backend Setup

1. Clone the repository
```bash
git clone https://github.com/thefonzie-codes/goLift.git
```

2. Set up PostgreSQL database
```bash
Please configure your own Postgres DB. Once there is a live db, it will be connected to the live app, of course.
```

3. Configure environment
```bash
cp backend/.env.example backend/.env
# Edit .env with your database credentials
```

4. Run the server
```bash
cd backend
go run cmd/server/main.go
```

### Testing
```bash
cd backend
go test -v ./...
```

## Project Structure
```
backend/
├── cmd/
│   └── server/          # Application entrypoint
├── internal/
│   ├── config/          # Configuration management
│   ├── database/        # Database connection
│   ├── handlers/        # HTTP handlers
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models
│   └── routes/          # Route definitions
├── pkg/
│   └── testutils/       # Testing utilities
└── goLift.gg.sql       # Database schema
```

## API Documentation
Coming soon...
```
