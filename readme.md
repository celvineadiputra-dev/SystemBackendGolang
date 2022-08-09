# Golang 🐱‍🚀
<em>version go1.18.4</em>
<hr/>

Required Package : 
1. gin-gonic
2. godotenv
3. mysql driver
4. gorm
5. jwt
6. crypto

<hr/>

**Setup Database :**
1. Open file .env and change the config by yourself
2. Open file main.go in root project
3. Uncomment (optional for migrate database)
```
bootstrap.MigrateDatabase()
```

**How to run with go :**
1. Open Project In Terminal and run
```
go run main.go
```

**How to run with npm :** 
1. Open Project in Terminal (optional)
```
npm install
```
2. Run (for auto reload)
```
npm run watch
```

**Test**
1. Open Api Test Ex : Postman
2. Import file api/api_collection.json for test
3. Open file documentation.yml for swagger documentation (optional)
4. Default url : **localhost:5120**

<hr/>