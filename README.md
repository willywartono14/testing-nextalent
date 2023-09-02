# Test Nextalent Project

This project is intended to accomplish [Nextalent] test.

## Tech Stack
- Golang v1.21
- PostgreSQL

## Tools
- Docker
- Docker Compose
- Git
- Postman

## How to Run the Project
### Prerequisites
1. Docker has been installed.
2. Docker Compose has been installed.
3. Golang has been installed.
4. Git has been installed.

### Steps
1. Go to the project folder
```shell
cd go/to/your/project/path
```
2. Clone this project
```shell
git clone github.com/willywartono14/testing-nextalent.git
```
3. Change directory to `testing-nextalent`
```shell
cd testing-nextalent
```
4. Execute `docker-compose` and wait for both `postgres` and `api` containers are successfully ran
```shell
docker-compose up --build
```
5. Execute migration
```shell
go run migration/main.go migrate-up
```

## Features

### GET ALL
1. **GET** `/api/country?type=getcountry&value=`
### GET By Name
1. **GET** `/api/country?type=getcountry&value=Adam`
### GET Time Zone by other API
1. **GET** `/api/country?type=gettimezone&value=Asia/Jakarta`
### Script Insert
1. **POST** `/api/country`

