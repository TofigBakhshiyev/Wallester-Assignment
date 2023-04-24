### Customer Platform

#### Usage
`Add postman collection for tests`

#### Database config
- `Install Docker` 
- `docker pull postgres`
- `docker run -itd -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -p 5432:5432 -v /data:/var/lib/postgresql/data --name postgresql postgres`

#### Run server
`go run server.go`