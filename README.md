# Income Tracker API - GraphQL Backend

A clean, well-structured GraphQL API for income tracking built with Go, Gin, gqlgen, MySQL, and GORM.

## 🛠️ Tech Stack

- **Go 1.21+** - Backend language
- **Gin** - HTTP routing framework
- **gqlgen** - GraphQL code generation
- **MySQL** - Primary database
- **GORM** - ORM for database operations
- **dotenv** - Environment configuration

## 📁 Project Structure

```
income_tracker_api/
├── cmd/                    # Application entrypoints
│   ├── migrate/           # Database migration CLI
│   └── server/            # HTTP server
├── pkg/                    # Public packages (can be imported by external apps)
│   ├── api/
│   │   └── graph/         # GraphQL implementation
│   │       ├── generated/ # Generated GraphQL code
│   │       ├── resolver/  # GraphQL resolvers
│   │       └── schema/    # GraphQL schema definitions
│   ├── config/            # Configuration management
│   ├── database/          # Database connection & migrations
│   ├── models/            # GORM models
│   └── utils/             # Shared utilities
├── internal/               # Private application code
│   └── migration/         # Internal migration scripts
├── scripts/                # Build and deployment scripts
├── test/
│   └── integration/       # Integration tests
├── docs/                   # Documentation
├── .env.example           # Environment variables template
├── gqlgen.yml             # GraphQL code generation config
├── go.mod                 # Go module definition
├── Makefile               # Build automation
└── README.md              # This file
```

## 🚀 Quickstart

### Prerequisites
- Go 1.21+
- MySQL 8.0+
- Git

### 1. Clone and Setup
```bash
git clone <repository-url>
cd income_tracker_api
```

### 2. Environment Configuration
```bash
cp .env.example .env
# Edit .env with your MySQL credentials
```

**Required environment variables:**
```env
PORT=8080
ENABLE_PLAYGROUND=true
DATABASE_DSN=root:your_password@tcp(127.0.0.1:3306)/income_tracker?charset=utf8mb4&parseTime=True&loc=Local
```

### 3. Database Setup
```bash
# Create database
mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS income_tracker CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# Run migrations
go run ./cmd/migrate
```

### 4. Start Server
```bash
go run ./cmd/server
```

**API Endpoints:**
- GraphQL: `POST /query`
- Playground: `GET /` (if enabled)

## 📊 API Examples

### Create User
```graphql
mutation CreateUser($input: NewUserInput!) {
  createUser(input: $input) {
    id
    email
    name
    createdAt
    updatedAt
  }
}
```
**Variables:**
```json
{ "input": { "email": "john@example.com", "name": "John Doe" } }
```

### Query Users
```graphql
query {
  users {
    id
    email
    name
    createdAt
    updatedAt
  }
}
```

### Get User by ID
```graphql
query GetUser($id: ID!) {
  user(id: $id) {
    id
    email
    name
    accounts {
      id
      name
      balance
    }
  }
}
```

## 🏗️ Development

### Adding New Features

1. **Update GraphQL Schema**
   ```bash
   # Edit: pkg/api/graph/schema/schema.graphqls
   ```

2. **Regenerate Code**
   ```bash
   go run github.com/99designs/gqlgen generate
   ```

3. **Implement Resolvers**
   ```bash
   # Edit: pkg/api/graph/resolver/schema.resolvers.go
   ```

4. **Update Models** (if needed)
   ```bash
   # Edit: pkg/models/*.go
   ```

5. **Run Migrations**
   ```bash
   go run ./cmd/migrate
   ```

### Code Quality

- **Consistent naming**: Use PascalCase for exported types/functions, camelCase for unexported
- **Error handling**: Always return descriptive errors
- **Imports**: Group standard library, then third-party, then internal packages
- **Documentation**: Add comments for exported functions/types

## 🗄️ Database

### Migration Strategy
- Uses GORM AutoMigrate for schema updates
- Safe for production (doesn't drop columns automatically)
- Manual SQL migrations for destructive changes

### Available Models
- **User** - System users
- **Currency** - Supported currencies
- **Account** - User accounts
- **Category** - Transaction categories
- **Transaction** - Financial transactions
- **Expense** - Expense records
- **Income** - Income records
- **Exchange** - Currency exchanges
- **Transfer** - Account transfers

## 📋 Commands

```bash
# Development
go mod tidy                    # Clean dependencies
go run ./cmd/server           # Start server
go run ./cmd/migrate          # Run migrations

# GraphQL
go run github.com/99designs/gqlgen generate  # Regenerate GraphQL code

# Testing
go test ./...                 # Run all tests

# Build
go build ./cmd/server         # Build server binary
```

## 🔧 Configuration

### Environment Variables
- `PORT` - Server port (default: 8080)
- `ENABLE_PLAYGROUND` - Enable GraphQL playground (default: true)
- `DATABASE_DSN` - MySQL connection string

### MySQL DSN Format
```
user:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.
