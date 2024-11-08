
**Context**

Welcome to AlphaTech, where we build simplified software solutions. You're a part of the backend team and we're eager to see how you adapt to our stack. Our primary language is Golang and we're currently working on a series of microservices.

**The Project**

We're in the process of developing a service that requires fast read and write operations. As a quick POC (Proof of Concept), you are tasked to build a REST API in Golang/Node Typescript that performs CRUD (Create, Read, Update, Delete) operations on a list of products. Additionally, implement a sorting algorithm that sorts these products based on price.

**Requirements**

REST API: Your API should have endpoints to:

- Create a product
- Read a product
- Update a product
- Delete a product
- List all products

Data Structure: Use a simple struct for a product, which includes:

- ID
- Name
- Price

Sorting Algorithm: Implement a sorting function that sorts the products based on price. You can use any sorting algorithm.

Test Cases: Write test cases to validate your sorting algorithm.

Bonus: Implement pagination on the list all products endpoint.

Documentation: Provide a README that explains how to run your code and use your API.

**Constraints**

Stick to Golang's(Node Typescript) standard library. You can use a package like Gorilla Mux(Expresss) for routing.
Keep your application as stateless as possible.
Please make sure to comment your code.
Success Criteria
All endpoints should work without any errors.
Sorting algorithm should correctly sort the products in ascending order based on price.
Test cases should cover basic scenarios and edge cases for the sorting algorithm.


**Documentation**

1 - To access the "deployments" folder and run the command "docker-compose up" to initiate the database container.

2 - To access the "cmd" folder and run the command "go run main.go" to initiate the web-server.

3 - To open the web browser and access the address "http://localhost:3333/docs" to view the docs and make requests through Swagger.

4 - At the root of the project run the command "go test ./..." to run the test for the sorting algorithm (quicksort).
