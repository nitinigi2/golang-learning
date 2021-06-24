# Desciption

This application has 3 services : 

    1. auth service for authentication
    2. book service for crud operations
    3. nginx service as api gateway

Usage : 

    1. User should login first
       http://localhost/authapi/login
    
    2. Once user is authorised then he can call book crud api's.
       http://localhost/bookapi/api/books

    3. Only Get request are accessible by anyone. For post, put and delete book api's user must logged in as 
       {
          "username": "admin",
          "password": "admin" 
       }
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

    1) GET - /bookapi/api/books

    2) GET - /bookapi/api/books/{id}

    3) POST - /bookapi/api/books

    4) PUT - /bookapi/api/books/{id}

    5) DELETE - /bookapi/api/books/{id}

Authserver-

    1) POST - /authapi/login

    2) POST - /authapi/logout

Book json format : 

    {
        "id": 1,
        "isbn": "temporary",
        "title": "Physics",
        "category": "High School",
        "description": "by Elon",
        "author": "Elon"
    }

Auth req format

    {
        "username: : "admin",
        "password" : "admin"
    }

TroubleShoot : (Cannot connect to mysql database: Access denied)

    docker-compose down -v                     //this will stop and remove running containers with any volume or network

Additional Info : 

    .env file : contains db and server config  // do not commit this file when working on real project

[database persistence in docker](https://www.youtube.com/watch?v=G-5c25DYnfI)

[docker swarm](https://www.youtube.com/watch?v=m6WgX_LBtEk)


Bookapi and authserver can be accessed by direct ip's as well.

BookApi server accessible on 

    http://localhost:5000/

Auth server accessible on 

    http://localhost:8000/

API's - 

Bookapi req/resp format-

    1) GET - /api/books

    2) GET - /api/books/{id}

    3) POST - /api/books

    4) PUT - /api/books/{id}

    5) DELETE - /api/books/{id}

Authserver-

    1) POST - /login

    2) POST - /logout

However, in general we should always access api's from gateway only not from individual ips.



