# Golang User System

User System created with Golang. It's really fun with Golang. Don't believe me? Go ahead & use it by yourself!! 👻🤝👀

# Updated version
Visit [this](https://github.com/yTakkar/Go-React-User-System) for single-page user-system with React & Golang!!

# Requirements
Following packages should be installed.
1. [httprouter](github.com/julienschmidt/httprouter)
2. [negroni](github.com/urfave/negroni)
3. [checkmail](github.com/badoux/checkmail)
4. [MySQL driver](github.com/go-sql-driver/mysql)
5. [bcrypt](golang.org/x/crypto/bcrypt)
6. [sessions](github.com/gorilla/sessions)

# Usage
1. Open PHPMyAdmin, create a db & import `db.sql`.
2. Go to `models/db.go` & change DB connection settings.
3. My root folder name is `Go-User-System`, if yours is different then go ahead & change it as it used for imports. It can be done easily by searching the whole project.
4. Now run it.
	```javascript
	go run main.go
	```
5. Check your result. Type..
	```javascript
	localhost:[PORT] PORT=1110 (By default)
	```
6. Enjoy!!
