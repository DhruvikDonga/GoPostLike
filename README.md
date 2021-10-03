## Golang gorilla based MVC template  

[TOC]

### About  
This project will help to divide the web application code to packages using MVC code structure .  
It uses gorilla/mux and gorm for mysql database  

### File structure:-  
- main.go *server configurations*  
- api  
-v1.go *contains the api routers*  
- DBs  
-Migrations.go *database configurations and migrations*  
- models  
-Users.go *sample table structure*  
- controllers  
-UserController.go *contains CRUD functions*  

### Further updates:-  
- JWT middleware  
- Swagger Docs  