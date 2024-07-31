# Go REST API with Beego

This repository contains a RESTful API built using the Beego framework.

## Features

- CRUD operations
- RESTful API structure
- Postgresql database integration

## Requirements

- Go 1.15 or higher
- Go Beego Web Framework
- Postgresql

## Getting Started

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/muthukumar89uk/go-restapi-beego.git
   ```
Click here to directly [download it](https://github.com/muthukumar89uk/go-restapi-beego/zipball/master).

## Install dependencies:

          go mod tidy

## Run the Application
  1. Run the Server
   
       ```
          go run .
       ```   
  2. The server will start on `http://localhost:9000`.

## API Endpoints
 
- `POST /v1/api/create/employee`  - Create a new employee details
- `GET /v1/api/get/employees`     - Retrieve all employees details
- `GET /v1/api/getById/:id`       - Retrieve an employee details by ID
- `PUT /v1/api/updateById/:id`    - Update an existing employee details
- `DELETE /v1/api/deleteById/:id` - Delete an employee details

## Refer
  - [Beego Web Framework](https://github.com/beego/beego) 
  - [GORM](https://gorm.io/)