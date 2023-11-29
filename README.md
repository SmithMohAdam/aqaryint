# aqaryint

## Student CRUD API
This is a simple MVC CRUD API built using the GIN and GORM frameworks for managing student data.

## Requirements
Before running this project, make sure you have the following installed:

Git
Go
Postgres
Installation
Clone the repository:


Install the project dependencies:
go get ./... 
## gin groam 

The following endpoints are available for managing student data:

POST /users
GET /users/id
UPDATE /users/id
DELETE /users/id



Send a JSON payload in the request body with the following format:
## POST  /users
{
  "name": "John Doe",
  "age": 20,
  "city": "Dubai",
  "class": "A"
}
This endpoint creates a new student with the provided details.

## Update a Student
PUT /users/:id

Replace  :id with the ID of the student you want to update. Send a JSON payload in the request body with the updated data. This endpoint updates the student with the specified ID.

## Delete a Student
DELETE /users/:id

Replace :id with the ID of the student you want to delete. This endpoint deletes the student with the specified ID.


