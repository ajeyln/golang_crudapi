<h1 align="center"> CRUD Operation Using Golang </h1>

This repository consist CRUD Operation files using Golang(Gorilla/Mux).

The following prerequisites are required:

1. Install MySQL Workbench, configure database and connect (These infomations are available in db-connection.go)
2. Install Postman to do CRUD Operation 

The following Specifications has to be done: 

1. MySQL based webserver listening on localhost at certain specified port for user requests <Br />
2. Endpoints <Br />
    * /customer - <br>
        - POST Method - Create a Customer <br>
    * /customers - <br>
        - GET Method - Get All Customer information in database <br >
    * /customer/{id} <br />
        -  GET Method - Get Information about customer using ID <br />
        -  POST Method - Update Customer Information <br />
        -  DELETE Method - Delete the Customer

Persistence storage
   * The Registered Customer information will be stored in a database

## Tasks

User needs to compile packages using the command **Go Build**, then **.exe** file will be generated. These file needs to be executed to start CRUD operations<p>

### 1. Creating Database<br />
Database will be created during migration and do the manipulations (POSTMAN) for the Endpoints <br />

### 2. Registration <br />

![Customer Registration](./screenshots/registation.png) <br />

Customer can be register on POSTMAN "/customer" Endpoint with POST Method in json format and these Customer details will be updated in Database <br />
![Customer Registration1](./screenshots/registation1.png) <br />
user can View all customers information on POSTMAN "/customers" Endpoint with GET Method.
![View Registered](./screenshots/View_data.png) <br />
user can View customer information by ID on POSTMAN "/customer/id" Endpoint with GET Method.
![View Registered1](./screenshots/View_data1.png) <br />
###3. Update user <br />

User can update or change customer information at "/customer/id" Endpoint with PUT Method. <br />

![Update Customer](./screenshots/update.png) <br />

### 4. Delete user <br />

![Delete User1](./screenshots/delete1.png) <br />
User can delete any Customer by ID, "/customer/id" Endpoint with DELETE Method. <br /> 

![Delete User2](./screenshots/delete2.png) <br />


<h2>Useful Links :</h2>

| **Sl. No.** | **Link** | **Remarks** |
----------|--------------|--------------
1| [Go Tutorial Basic](https://www.youtube.com/watch?v=ty49_v1tV44) | Golang Tutorial for Beginners (Navin Reddy) |
2| [Golang REST API With Mux ](https://www.youtube.com/watch?v=SonwZ6MF5BE) | Golang REST API With Mux (Traversy Media
) |


