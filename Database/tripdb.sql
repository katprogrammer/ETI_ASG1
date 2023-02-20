DROP database if exists trips_db;
CREATE database trips_db;
USE trips_db;
 
 # Creating Trips Table
CREATE TABLE Trips (
 TripID VARCHAR(10) PRIMARY KEY UNIQUE,
 StartPostalCode BIGINT NOT NULL ,
 EndPostalCode BIGINT NOT NULL,
 TripStatus VARCHAR(10) NOT NULL,
 StartTime DATETIME NULL,
 EndTime DATETIME NULL,
 PassengerID VARCHAR(10) NOT NULL,
 DriverID VARCHAR(10) NOT NULL
 );

 #===============================[Creating Rows for Trips]===============================#
INSERT INTO Trips(TripID, StartPostalCode, EndPostalCode, TripStatus, StartTime, EndTime, PassengerID, DriverID)
VALUES("T1",680210,556798,"Started",NOW(),'2022-12-3 13:15:00',"P3","D2");

INSERT INTO Trips(TripID, StartPostalCode, EndPostalCode, TripStatus, StartTime, EndTime, PassengerID, DriverID)
VALUES("T2",456891,789450,"Started",NOW(),'2022-12-3 19:15:00',"P3","D2");

INSERT INTO Trips(TripID, StartPostalCode, EndPostalCode, TripStatus, StartTime, EndTime, PassengerID, DriverID)
VALUES("T3",123789,331250,'Ended','2022-12-17 15:46:33','2022-12-17 15:47:14',"P1","D2");

INSERT INTO Trips(TripID, StartPostalCode, EndPostalCode, TripStatus, StartTime, EndTime, PassengerID, DriverID)
VALUES("T4",782347,231653,'Ended','2022-12-17 19:31:46','2022-12-17 19:31:52',"P1","D1");

INSERT INTO Trips(TripID, StartPostalCode, EndPostalCode, TripStatus, StartTime, EndTime, PassengerID, DriverID)
VALUES("T5",246815,912367,'Ended','2022-12-17 19:32:04','2022-12-17 19:33:00',"P1","D1");

INSERT INTO Trips(TripID, StartPostalCode, EndPostalCode, TripStatus, StartTime, EndTime, PassengerID, DriverID)
VALUES("T6",344553,909091,'Ended','2022-12-18 10:59:39','2022-12-18 10:59:54',"P2","D1");

INSERT INTO Trips(TripID, StartPostalCode, EndPostalCode, TripStatus, StartTime, EndTime, PassengerID, DriverID)
VALUES("T7",567813,736489,'Requested','2022-12-15 19:32:04','2022-12-15 21:33:00',"P1","D1");

INSERT INTO Trips(TripID, StartPostalCode, EndPostalCode, TripStatus, StartTime, EndTime, PassengerID, DriverID)
VALUES("T8",991283,348929,'Requested','2022-12-16 10:59:39','2022-12-16 13:59:54',"P2","D2");

SELECT * FROM Trips

