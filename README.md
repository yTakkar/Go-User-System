# Golang User System

User System created with Golang. It's really fun with Golang. Don't believe me? Go ahead & use it by yourself!! 👻🤝👀

# Updated version
Visit [this](https://github.com/yTakkar/Go-React-User-System) for single-page user-system with React & Golang!!

# Requirements
1. Make sure you keep this project inside `src/` of your Golang project folder.
2. Following packages should be installed.
    1. [httprouter](https://github.com/julienschmidt/httprouter)
    2. [negroni](https://github.com/urfave/negroni)
    3. [checkmail](https://github.com/badoux/checkmail)
    4. [MySQL driver](https://github.com/go-sql-driver/mysql)
    5. [bcrypt](https://golang.org/x/crypto/bcrypt)
    6. [sessions](https://github.com/gorilla/sessions)
    7. [godotenv](https://github.com/joho/godotenv)

# Usage
1. First install all the dependencies.
```javascript
npm install 
OR 
yarn
```

1. Open PHPMyAdmin, create a db & import `db.sql`.
2. Create a `.env` file & insert the following code. Repalce values with yours.
```javascript
PORT=PORT (default= 1100)  [STRING]
DB_USER=DB USER [STRING]
DB_PASSWORD=DB PASSWORD [STRING]
DB=DB YOU JUST CREATED [STRING]
```

3. My root folder name is `Go-User-System`, if yours is different then go ahead & change it as it used for imports. It can be done easily by searching the whole project.

4. Now run it.
	```javascript
	go run main.go
	```

5. Run the app in browser.
	```javascript
	localhost:[PORT] PORT=1110 (By default)
	```

6. Enjoy!!
