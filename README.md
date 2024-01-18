# Campaign Management

A simple Full-Stack application for campaign management.

## Features

- Visualize your campaigns list
- Add a new campaign
- Update an existing campaign name or status whether it is ENABLED or PAUSED
- Delete a campaign

## Technical requirements

- Go version v1.21.5
- Node version v21.1.0

## Technologies

### Backend

- Go
- Echo
- viper
- In memory database

### Frontend

- TypeScript
- React
- Vite
- axios
- tailwindcss
- react-query

## Project setup and local execution

- Clone the repository to your local environment:

```bash
git clone https://github.com/buemura/campaign-management.git
```

- Install Go dependencies

```bash
go mod download

```

- Run Go API

```bash
go run cmd/http/main.go

```

- Install React dependencies

```bash
cd web && npm ci
```

- Run Go API

```bash
cd web && npm run dev

```
