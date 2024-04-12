# Golang REST API with Gin and Swagger

This project demonstrates the creation of a RESTful API using the Go programming language and the Gin framework. It also showcases the integration of Swagger for API documentation, making it easier to understand and interact with the API endpoints.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project is a simple yet powerful example of a RESTful API that manages person data. It includes CRUD (Create, Read, Update, Delete) operations for managing person data, all documented using Swagger.

## Features

- **CRUD Operations**: Full CRUD operations for managing person data.
- **Swagger Documentation**: Utilizes Swagger for API documentation, providing a clear and interactive way to understand and test the API endpoints.

## Prerequisites

- Go 1.16 or later
- Gin framework
- Swag for Swagger documentation

## Installation

1. Clone the repository: git clone https://github.com/Animesh-roy100/REST-API-Using-Go.git

2. Navigate to the project directory: cd rest-api-using-go

3. Install dependencies: go mod download

4. Generate Swagger documentation: swag init

5. Run the server: go run main.go

## Usage

Once the server is running, you can access the Swagger UI at `http://localhost:<your_port>/swagger/index.html` to interact with the API endpoints.

## API Documentation

The API documentation is generated using Swagger and can be accessed at the `/swagger/index.html` endpoint. This documentation provides detailed information about each API endpoint, including the expected request format and response structure.

## Contributing

Contributions are welcome! Please read the [contributing guidelines](CONTRIBUTING.md) before getting started.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
