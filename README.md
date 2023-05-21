Instructions<br>
Start the application by running the following command:<br>
- `docker-compose up`<br>
Alternatively, you can use the following command if you have a Makefile:<br>

- `make run`<br>
After starting the application, you can test the API using the following endpoints:<br>

Create a user: Send a POST request to http://localhost:3000/simple-service/users<br>
Login: Send a GET request to http://localhost:3000/simple-service/users/login<br>
Logout: Send a POST request to http://localhost:3000/simple-service/users/logout<br>
Please note that you need to replace http://localhost:3000 with the appropriate host and port if you are running the application on a different address.
