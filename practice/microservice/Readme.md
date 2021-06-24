# Desciption

This application has 2 services : 

    1. auth service for authentication
    2. book service for crud operations

Usage : 

    1. User should login first by auth-server. 
       http://localhost:8000/login
       This will create a JWT token and will store it in cache.

    2. Once user is authorised then he can call book crud api's.
       http://localhost:5000/api/books

# Run this application:

    1) docker compose up               //this will run 3 containers, 1 for bookapi, 1 for authserver and 1 for mysql db

    2) Once mysql container is running : 

        2.1) Login into mysql container using:
                 docker exec -it container_id bash
        2.2) Login as root user : 
                 mysql -u root -p
        2.3) create schema and table using : 
                 create database restapi;
                 use restapi;
                 CREATE TABLE book (
                        id int primary key not null auto_increment,
                        isbn varchar(255),
                        title varchar(255),
                        category varchar(255),
                        description varchar(255),
                        author varchar(255)
                 );
                  CREATE TABLE user (
                              username varchar(255) primary key not null,
                              password varchar(255),
                              role varchar(255)
                  );
                  INSERT INTO USER VALUES("admin", "admin", "admin");


To stop running containers-  

      docker compose down

API's - 

Bookapi req/resp format-

    1) GET - /api/books

    2) GET - /api/books/{id}

    3) POST - /api/books

    4) PUT - /api/books/{id}

    5) DELETE - /api/books/{id}

BookApi server accessible on 

    http://localhost:5000/

Book json format : 

    {
        "id": 1,
        "isbn": "temporary",
        "title": "Physics",
        "category": "High School",
        "description": "by Elon",
        "author": "Elon"
    }

Authserver-

    1) POST - /login

    2) POST - /logout

Auth req format

    {
        "username: : "admin",
        "password" : "admin"
    }

Auth server accessible on 

    http://localhost:8000/


TroubleShoot : (Cannot connect to mysql database: Access denied)

    docker-compose down -v                     //this will stop and remove running containers with any volume or network

Additional Info : 

    .env file : contains db and server config  // do not commit this file when working on real project

[database persistence in docker](https://www.youtube.com/watch?v=G-5c25DYnfI)

[docker swarm](https://www.youtube.com/watch?v=m6WgX_LBtEk)
