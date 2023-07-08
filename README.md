

# GO CRUD
The purpose of my web application is to provide a platform for performing CRUD operations and showcasing relevant statistics.
It allows users to create, read, update, and delete records in a user-friendly manner while providing insights through the displayed statistics.

## Installing / Getting started
To get started with the GO-CRUD web application, follow the steps below:

Prerequisites
- Go (version 1.20.5): [Download Go](https://golang.org/dl/)
- Docker: [Download Docker](https://www.docker.com/)
- PostgreSQL: [Download PostgreSQL](https://www.postgresql.org/download/)
- PgAdmin 4 or Beekeeper:[Download PgAdmin 4](https://www.pgadmin.org/download/)
- or
-  [Download Beekeeper([https://www.pgadmin.org/download/](https://www.beekeeperstudio.io/get)) 


                                       
## Clone the Repository
1. Clone the repository to your local machine:
```shell
git clone <repository-url>
```
2. cd into the directory's project
```shell
cd go-crud
```

## Set Up the Development Environment
1. Make sure you have Docker installed and running on your system.
2. Install Go by following the instructions on the official Go website.
3. Install PostgreSQL and PgAdmin 4 by downloading the respective installers from the provided links.

   
## Build and Run the Application
Use the following command to build and run the Docker containers:
```shell
docker-compose up
```
This command will build the necessary Docker images and start the containers defined in the docker-compose.yml file.

## Access the Web Application
Once the Docker containers are up and running,
open a web browser and navigate to http://localhost:3000/users/form to access the GO-CRUD web application.

## Developing

### Built With
- Go (version 1.20.5)
- Gin framework (version 1.9.1)
- GORM (version 1.25.1)
- PostgreSQL (version 15)
- HTML, CSS, and HTML5


## Setting up Dev
To set up the development environment:

1. Clone the repository to your local machine.
2. Install Go and Docker as mentioned in the prerequisites.
3. Configure the environment variables as specified above.
4. Build and run the Docker containers using docker-compose up.

## Testing
To run the tests, use the following command:
```shell
go test ./...
```
## Database
The GO-CRUD web application uses PostgreSQL as the database system.
Make sure you have PostgreSQL installed and configured with the correct connection details.

## Licensing
The project is licensed under the MIT license. You can find the text version of the license in the LICENSE file.

Feel free to replace <repository-url> with the actual URL of your repository.

