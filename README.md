Instructions
Start the application by running the following command:
- docker-compose up
Alternatively, you can use the following command if you have a Makefile:

- make run
After starting the application, you can test the API using the following endpoints:

Create a user: Send a POST request to http://localhost:3000/simple-service/users
Login: Send a GET request to http://localhost:3000/simple-service/users/login
Logout: Send a POST request to http://localhost:3000/simple-service/users/logout
Please note that you need to replace http://localhost:3000 with the appropriate host and port if you are running the application on a different address.
