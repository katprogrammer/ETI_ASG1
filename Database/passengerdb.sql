DROP database if exists passenger_db;
CREATE database passenger_db;
USE passenger_db;

# Creating Table
CREATE TABLE Passengers (
 PassengerID VARCHAR(10) PRIMARY KEY UNIQUE,
 FirstName VARCHAR(50) NOT NULL,
 LastName VARCHAR(50) NOT NULL,
 PhoneNum varchar(8) NOT NULL unique,
 Email VARCHAR(100) NOT NULL unique
 );

#===============================[Creating Rows for Passenger]===============================#
INSERT INTO Passengers (PassengerID, FirstName, LastName, PhoneNum, Email) VALUES ("P1","Peter","Hung","93629930","peterhung@gmail.com");
INSERT INTO Passengers (PassengerID, FirstName, LastName, PhoneNum, Email) VALUES ("P2","Charis","Tang","91345690","charistang@hotmail.com");
INSERT INTO Passengers (PassengerID, FirstName, LastName, PhoneNum, Email) VALUES ("P3","Mike","Wee","98901234","mikewee@gmail.com");

select * from passengers

