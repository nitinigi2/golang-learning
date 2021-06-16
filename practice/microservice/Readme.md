# Run this application run in terminal:

Create a folder named mysqldb : this folder will be used for database persistance

      1)docker build -t book-service .  //this will create a docker image with current directory

      2) docker compose up               //this will run 2 containers, 1 for webapp and 1 for mysql db

      3) Once mysql container is running : 

          3.1) Login into mysql container using:
                 docker exec -it container_id bash
          3.2) Login as root user : 
                 mysql -u root -p
          3.3) create schema and table using : 
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
                              password varchar(255)
                  );
                  INSERT INTO USER VALUES("admin", "admin");


To stop running containers-  

      docker compose down

API's - 

      1) GET - /api/books

      2) GET - /api/books/{id}

      3) POST - /api/books

      4) PUT - /api/books/{id}

      5) DELETE - /api/books/{id}


book json format : 

        {
                "id": 1,
                "isbn": "temporary",
                "title": "Physics",
                "category": "High School",
                "description": "by Elon",
                "author": "Elon"
        }

TroubleShoot : (Cannot connect to mysql database: Access denied)

      docker-compose down -v                     //this will stop and remove running containers with any volume or network

Additional Info : 

      .env file : contains db and server config  // do not commit this file when working on real project

[database persistence in docker](https://www.youtube.com/watch?v=G-5c25DYnfI)

[docker swarm](https://www.youtube.com/watch?v=m6WgX_LBtEk)
