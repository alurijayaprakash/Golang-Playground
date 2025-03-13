# Redis-Cache Example with Go
This repository provides a simple example of using Redis as a caching mechanism in a Go application. It demonstrates how to store and retrieve student data using Redis to improve performance and reduce the load on the database.


## Features
**Redis caching**: The code utilizes Redis, an in-memory data structure store, as a cache to store student data. This helps reduce the need for repeated database queries and improves the overall application performance.

**HTTP server:** The code sets up an HTTP server using the Gorilla Mux router library, allowing clients to retrieve student data and store new student records via HTTP GET and POST requests, respectively.

**JSON serialization:** The student data is serialized to JSON format before being stored in Redis and deserialized back into a Student object when retrieved from the cache.

**Redis client:** The code uses the Go-Redis library, a Redis client for Go, to interact with the Redis server. It establishes a connection to Redis, sets and retrieves student data, and handles any errors that may occur.

## Usage
**Install Redis:** Make sure you have Redis installed and running on your local machine. You can download Redis from the official website or use a package manager.

**Install Go:** If you haven't already, install the Go programming language on your system. You can download it from the official website and follow the installation instructions.

**Clone the repository:** Clone this repository to your local machine using the command:
```
git clone https://github.com/alurijayaprakash/Golang-Playground.git
```
**Navigate to the project directory:**

```
cd Redis_REST_API
```
**Run the application:**
```
go run main.go
```

### Interact with the API: 
Once the server is running, you can make HTTP GET requests to retrieve student data by accessing the /getstudent/{id} endpoint, where {id} is the ID of the student. You can also store new student records by sending a JSON payload via an HTTP POST request to the /setstudent endpoint.

**Retrieve Student Data:**

Endpoint: GET /getstudent/{id}

Example: http://localhost:8080/getstudent/1

Description: This endpoint allows you to retrieve student data by providing the student's ID as a path parameter. Replace {id} with the desired student ID.

Please note that for the GET request, the payload is not required as it is a GET request to retrieve data based on the provided ID.

**Store Student Data:**

Endpoint: POST /setstudent

Example: http://localhost:8080/setstudent

Payoad :
```
{
  "id": "2",
  "name": "John Doe",
  "gpa": 3.5,
  "isEligible": true
}
```
This payload represents a new student record with the following details:

    ID: "2"

    Name: "John Doe"

    GPA: 3.5

    Eligibility: true

You can modify the values in the payload to store different student records.


**Description:** This endpoint allows you to store new student records by sending a JSON payload via an HTTP POST request. The payload should include the student's ID, name, GPA, and eligibility status.

Make sure to replace localhost:8080 with the appropriate host and port if you're running the server on a different address.

**Note:**

Make sure to adjust the Redis connection details in the InitializeRedis function according to your Redis server configuration.
This example assumes a single Redis instance running on the local machine, but you can modify the connection settings to connect to a remote Redis server.

**Contributing**
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

**License**
This project is licensed under the MIT License.
