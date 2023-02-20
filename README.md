# ETI_ASG1 Wai Keat's Ride Sharing Microservices

## Overview
*** 
This ride-sharing platform connects passengers with drivers so passengers can reach their destination. The platform typically consists of multiple components, including a frontend interface for passengers to request rides and for drivers to accept them, a backend server for processing ride requests and managing drivers, and a database for storing ride and driver information.

In the case of using microservices architecture, each component of the ride-sharing platform can be designed and deployed as a separate microservice. For example, the frontend interface for passengers and drivers can be implemented as a separate microservice, and the backend server that processes ride requests can be another microservice. By splitting up the application into smaller, modular components, it becomes easier to maintain and scale the system as needed.

To enable persistent storage, MySQL can be used as the database management system for the ride-sharing platform. MySQL is a popular relational database system that provides features such as ACID compliance, data replication, and data partitioning for horizontal scaling.

## Architecture Diagram
*** 
![Architecture Diagram](https://user-images.githubusercontent.com/73064976/220102446-34597b37-1afd-49e1-ade8-1c9904c9c161.jpg)

Here is an overview of how the ride-sharing platform with passenger, driver, and interface microservices can be implemented using MySQL for persistent storage:

* Passenger Microservice:
The Passenger microservice will handle passenger-related functionalities such as ride requests, account creation and update account details. It will communicate with the driver microservice to find an available driver for a ride request. The Passenger microservice will store passenger information, ride requests, and rating data in the MySQL database.

* Driver Microservice:
The Driver microservice will handle driver-related functionalities such as accepting and completing ride requests. It will communicate with the Passenger microservice to receive ride requests and provide ride status updates. The Driver microservice will store driver information and ride history data in the MySQL database.

* Interface Microservice:
The Interface microservice will provide the frontend interface for passengers and drivers to interact with the ride-sharing platform. It will communicate with the Passenger and Driver microservices to display ride information and handle user input. The Interface microservice will not store any data in the MySQL database but will retrieve data from the other microservices as needed.

* MySQL Database:
The MySQL databases will store all data related to passenger, driver and trip history information. The database will be managed by a database management system (DBMS) such as MySQL Server. The Passenger and Driver microservices will interact with the database through a MySQL connector library using go programming language

## Design Considerations using Domain Driven Design (DDD)
*** 
Domain-driven design (DDD) is an approach to software development that emphasizes the importance of understanding and modeling the business domain of a system. DDD encourages developers to think about the domain in terms of distinct, cohesive concepts or "domains" that interact with each other.

For a ride-sharing platform, there are two primary domains: Passenger and Driver. Let's take a closer look at how these domains can be implemented using DDD principles.

Passenger Domain:
The Passenger domain represents the users of the platform who request rides. The Passenger domain should have the following characteristics:
* Identity: Each passenger should have a unique identifier that can be used to track their ride requests and history.
* Attributes: Passengers can have attributes such as first name, last name, email, phone number 
* Actions: Passengers can create account, update account details, request rides, view ride history.

Driver Domain:
The Driver domain represents the users of the platform who provide rides to passengers. The Driver domain should have the following characteristics:

* Identity: Each driver should have a unique identifier that can be used to track their availability and ride history.
* Attributes: Drivers can have attributes such as first name, last name, email, identification number, phone number, lisence number and driver status
* Actions: Drivers can start rides, end rides, and view ride history.

Now let's see how these domains can be implemented using DDD principles:

* Bounded Contexts:
In DDD, a bounded context represents a distinct area of the system where a certain domain concept is defined and used. For our ride-sharing platform, we can have two bounded contexts: Passenger context and Driver context. These bounded contexts will define the entities, services, and repositories related to Passenger and Driver domains.

* Entities:
Entities are objects with unique identities that have a lifecycle within a bounded context. In our ride-sharing platform, Passenger and Driver are entities. They will have unique identifiers, attributes, and actions associated with them.

* Services:
Services are operations that don't fit well within a single entity but have meaning in the domain. For example, finding an available driver for a ride request is a service that involves both Passenger and Driver domains. This service can be defined in the Passenger bounded context and called by the application layer.

* Repositories:
Repositories are objects that provide an abstraction layer over data storage. They allow us to work with entities without worrying about the underlying data storage. In our ride-sharing platform, we can have PassengerRepository and DriverRepository that handle the persistence of passenger and driver entities.

* Aggregate Roots:
Aggregate roots are entities that act as a gateway to access other entities within the same bounded context. In our ride-sharing platform, Passenger and Driver can be aggregate roots. They will encapsulate the behavior of the entities within their respective domains.

* Domain-driven design can be implemented for a ride-sharing platform by identifying the primary domains (Passenger and Driver), defining bounded contexts for each domain, defining entities, services, repositories, and aggregate roots within each bounded context, and using these elements to build the application layer of the system. By applying DDD principles, we can build a system that models the ride-sharing domain accurately and provides a better user experience.

## API Documentation
*** 
### Passenger Microservice

##### Endpoints:
* GET /api/v1/passenger/view: Returns a list of all the passengers.
* POST /api/v1/passenger/create/{passengerid}: Adds a new passenger record in the MySQL database with the specified passenger ID and their FirstName, LastName, PhoneNumber, and Email. Their new account will then be accessible from the menu.
* PUT /api/v1/passenger/update/{passengerid}: Updates passenger account details in the MySQL database based on the specified passengerID and updates their FirstName, LastName, MobileNumber, and Email. The menu will then display the updated details.
* POST /api/v1/passenger/trip/{tripid}/{passengerid}/{driverid}: Creates a new trip record for the specified passenger. A random available driver is assigned to the ride created and their status is set to 'Occupied.
* GET /api/v1/trip/{passengerid}: Returns a list of all the passenger's trips in reverse chronological order.

### Driver Microservice

##### Endpoints:
* GET /api/v1/driver/view: Returns a list of all the drivers.
* GET /api/v1/driver/trips/{driverid}: Returns a list of all the driver's trips.
* POST /api/v1/driver/create/{driverid}: Adds a new driver record in the MySQL database with the specified driver ID and their FirstName, LastName, PhoneNumber, NRIC, Vehicle lisence number and Email. Their new account will then be accessible from the menu.
* PUT /api/v1/driver/update/{driverid}: Updates driver account details in the MySQL database based on the specified driverID and updates their FirstName, LastName, NRIC, Vehicle lisence number and Email. The menu will then display the updated details.
* PUT /api/v1/driver/update/{driverid}/occupied: Updates driver status automatically to occupied after starting a trip.
* PUT /api/v1/driver/start/{tripid}: Updates trip status to "Started" after starting a trip.
* PUT /api/v1/driver/end/{tripid}/{driverid}: Updates trip status to "Ended" after ending a trip and updates the driver status to "Available"
* /api/v1/drivers: Auto assigns a driver to a trip after the passenger has requested a trip and input the pick up and drop off postal codes.

## Instructions to set-up and run microservices
*** 
1. Download / Clone the repo , and all the folders for the microservices and save it to your desktop / designated folder.
3. Open MySQL WorkBench
4. Connect to local instance, ensure that the localhost port is :3306 and enter your password.
5. After that, in the local instance, open the three different sql files ("passengerdb.sql", "driverdb.sql" and "tripdb.sql")
7. On your keyboard, press Ctrl + Shift + Enter to execute the entire Sql Query in each database.
8. After the tables have been created to the new database, open each of the .go files (In VSCode) for the Passenger and Driver folders.
9. Open Command Prompt and open three tabs for the passenger microservice, driver microservice and the interface microservice.
10. Then navigate to the location of where the microservice .go files are saved. (For this example: C:\Users\username\ -> type in "cd desktop/menu")
11. In the passenger tab, type "go run passenger.go" to run the passenger microservice.
12. In the driver tab, type "go run driver.go" to run the driver microservice.
13. After running the microservices, in the console tab, type "go run console.go" the menu will then be displayed for the user to use the Ride-Sharing Platform
