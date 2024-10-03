# Phost - Image Host System

This project provides a sample image hosting service using golang (database, backend) and nextjs (frontend)

## Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Test](#test)
6. [Contributing](#contributing)
7. [License](#license)

### Introduction

Phost started with my goal to develop a sample database using Golang. At the moment it has a very simple storage logic, but it is a database with its own syntax and its own business logic to understand how popular sql servers work. I aim to develop this database in the future to make it more complex.

After developing the database section, I wanted to make the whole system an interface instead of just leaving it with virtual tests. I prepared a backend service that handles CRUD operations and a simple interface that users can interact with.

### Features

- Basic Database
- CRUD operations
- Unit tests
- CI/CD using Github Actions

### Installation

1. Clone this project: `git clone https://github.com/afurgapil/phost.git`
2. Navigate to the project directory: `cd phost`
3. Check `.env` files
4. Install database dependencies: `cd database || go mod tidy`
5. Install backend dependencies: `cd backend || go mod tidy`
6. Install frontend dependencies: `cd frontend || npm install`

### Usage

1. Clone this project: `git clone https://github.com/afurgapil/phost.git`
2. Navigate to the project directory: `cd phost`
3. Check `.env` files
4. Run database: `cd database || go run cmd/phost/main.go`
5. Run backend: `cd backend || go run cmd/phost-backend/main.go`
6. Run frontend: `cd frontend || npm run dev`

### Test

- Database tests: `cd database || go test ./... -v`
- Backend tests: `cd backend || go test ./... -v`
- Frontend tests:`cd frontend || npm test`

### Contributing

If you encounter any issues or have suggestions for improvements, please feel free to contribute. Your feedback is highly appreciated and contributes to the learning experience.

### License

This project is licensed under the [MIT License](LICENSE). For more information, please refer to the LICENSE file.
