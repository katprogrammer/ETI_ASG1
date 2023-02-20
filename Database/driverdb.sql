DROP database if exists driver_db;
CREATE database driver_db;
USE driver_db;
 
 # Creating Driver Table
CREATE TABLE Drivers (
 DriverID VARCHAR(10) PRIMARY KEY UNIQUE,
 FirstName VARCHAR(50) NOT NULL,
 LastName VARCHAR(50) NOT NULL,
 PhoneNum varchar(8) NOT NULL unique,
 Email VARCHAR(100) NOT NULL unique,
 NRIC VARCHAR(9) NOT NULL unique,
 LisenceNum VARCHAR(15) NOT NULL UNIQUE,
 DriverStatus VARCHAR(10) DEFAULT NULL
 );
 
 #===============================[Creating Rows for Driver]===============================#
INSERT INTO Drivers (DriverID, FirstName, LastName, PhoneNum, Email, NRIC, LisenceNum, DriverStatus) VALUES ("D1","Anita","Ryanto","92235566","anita@gmail.com","T1234567E","SLP1020G", "Available");
INSERT INTO Drivers (DriverID, FirstName, LastName, PhoneNum, Email, NRIC, LisenceNum, DriverStatus) VALUES ("D2","Jackie","Chan","93381234","jackie@hotmail.com","T9876543E","NLB2020F", "Occupied");

SELECT * FROM Drivers
