# ITEMS API

## How to Run Locally

### Prerequisites
- Go 1.20 or superior
- MySQL
- Docker (optional)

### Configuration
1. Clone the repository.
2. Copy the environment variables example file:
   ```bash
   cp .env.example .env
   ```
3. Edit the `.env` file with your database credentials.

### Execution
#### Via Go Local:
```bash
go run main.go
```

#### Via Docker:
```bash
docker build -t items-api .
docker run -p 9000:9000 --env-file .env items-api
```

### Environment Variables
| Variable | Description | Default Value |
|----------|-----------|--------------|
| `API_PORT` | Port the API listens on | `9000` |
| `DB_USUARIO` | Database user | - |
| `DB_SENHA` | Database password | - |
| `DB_URL` | Database host | `localhost` |
| `DB_PORT` | Database port | `3306` |
| `DB_NOME` | Database name | `brasil` |
| `APP_NAME` | NewRelic application name | `api-items` |
| `NEW_RELIC_LICENSE_KEY` | NewRelic license key | - |
