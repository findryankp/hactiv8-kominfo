# Tugas 2 Kominfo - Hactiv8 - Golang
API Orders.

**Development by:** 
- Findryankp

## Init Project
- go mod init orders

## Install Packages
* go get github.com/gin-gonic/gin
* go get github.com/go-sql-driver/mysql
* go get github.com/jinzhu/gorm
* go get github.com/joho/godotenv

**Swaggo documentation API:** 
* go get -u github.com/swaggo/swag/cmd/swag
* go get -u github.com/swaggo/files
* go get -u github.com/swaggo/gin-swagger

## Step By step (MAC OS)
- export PATH=$(go env GOPATH)/bin:$PATH
- swag init

## Fitur
* Auth
   - Register
   - Login
   - Logout
* Manajemen User
* Manajemen Pekerjaan
* Manajemen Lokasi
* Manajemen Transaksi Lokasi

## License
[MIT](https://choosealicense.com/licenses/mit/)