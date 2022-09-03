
# Technical Test





## Prepare Database

Import db.sql to mariadb
    
## Run Locally

Clone the project

```bash
  git clone https://github.com/wahyugnc/technical-test.git
```

Go to the project directory

```bash
  cd technical-test
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run main.go
```


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`MARIADB_CONNECTION_STRING=root:root123@tcp(127.0.0.1:3306)/usedeall_db?charset=utf8mb4&parseTime=True&loc=Local`

`JWT_SECRET_KEY=yprxr6sXO8Maz1nicmwPxeX3TbKEKr7z`

`JWT_ACCESS_TOKEN_EXPIRE=15`

`JWT_REFRESH_TOKEN_EXPIRE=60`


## Documentation

Import collection to postman

[![Download Collection](https://heremaps.github.io/postman-collections/img/download.svg)](../../raw/master/UseDeall.postman_collection.json)


## Demo

Role Admin and User

```bash
  Admin 
  username : wahyugnc
  password : 123123

  User
  username : nabil123
  password : 123123
```


## Authors

- [@wahyugnc](https://www.github.com/wahyugnc)

