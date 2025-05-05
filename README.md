# Agnos Candidate Assignment

This project is an Agnos Candidate Assignment by using the Gin web framework in Go with NGINX, and PostgreSQL. It provides a minimal example to get me started with tracing HTTP requests in a Gin application. Unfortunately, I do not have enough time to develop the project with Docker yet. However, it works ok.

## Prerequisites

- Go 1.23 or later
- Git
- NGINX
- PostgreSQL

## Project Structure

- **controllers/**: Contains GO controller files.
  - `patientController.go`: The controller file of Patient.
  - `staffController.go`: The controller file of Staff.
- **initializers/**: Contains GO initailizing files.
  - `database.go`: The db connection file.
  - `loadEnvs.go`: The loading .env file.
- **middlewares/**: Contains GO middleware files.
  - `auth.go`: The procedure of authorization file.
- **models/**: Contains GO model files.
  - `authInput.go`: The Model of auth input.
  - `patient.go`: The Model of Patient.
  - `staff.go`: The Model of Staff.
- **.env**: The config of the project such as db_name, db_port, engine_port, etcs.
- **go.mod**: The Go module file, listing the dependencies required for the project.
- **go.sum**: The Go checksum file, ensuring the integrity of the dependencies.
- **main.go**: The main application file, setting up the web server and routes.
- **main_test.go**: The main testing file.
- **sql.sql**: SQL file for creating Tables and insert dummy datas.
- **nginx.conf**: Nginx configuration file for setting up proxy server.

## Configuration

Ensure to config value of database and engine in '.env' file correctly before run the application.

## Installation

### Clone the repository

```bash
git clone https://github.com/somkiatprogrammer/agnos-candidate-assignment.git
cd agnos-candidate-assignment

### Install the dependencies

```bash
go mod tidy
```

## Running the Example

To run the example application, use the following command:

```bash
go run main.go
```

This will start the Gin server on `http://localhost:8080`.

### Config NGINX

Copy nginx.conf to your nginx directory for setting up proxy server, 
in case of you do not have other projects. 

