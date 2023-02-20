package main

// Import packages required
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type Passenger struct {
	PassengerID string `json:"Passenger ID"`
	FirstName   string `json:"First Name"`
	LastName    string `json:"Last Name"`
	PhoneNum    string `json:"Phone Number"`
	Email       string `json:"Email Address"`
}

type Driver struct {
	DriverID     string `json:"Driver ID"`
	FirstName    string `json:"First Name"`
	LastName     string `json:"Last Name"`
	PhoneNum     string `json:"Phone Number"`
	Email        string `json:"Email"`
	NRIC         string `json:"NRIC"`
	LisenceNum   string `json:"License Number"`
	DriverStatus string `json:"Driver Status"`
}

type Trip struct {
	TripID          string `json:"Trip ID"`
	StartPostalCode string `json:"Start Postal Code"`
	EndPostalCode   string `json:"End Postal Code"`
	TripStatus      string `json:"Trip Status"`
	StartTime       string `json:"Start Time"`
	EndTime         string `json:"End Time"`
	DriverID        string `json:"DriverID"`
	PassengerID     string `json:"PassengerID"`
}

type AllDriver struct {
	Driver map[string]Driver `json:"Driver"`
}

var driverlist = map[string]Driver{}
var drivertriplist = map[string]Trip{}
var newdriver = map[string]Driver{}

func main() {
	// Connect to Router
	// Connect to Router
	router := mux.NewRouter()

	// View all the drivers available in the database.
	router.HandleFunc("/api/v1/driver/view/", getDriver).Methods("GET")

	// View all the drivers trips in the database.
	router.HandleFunc("/api/v1/driver/trips/{driverid}", getDriverTrips).Methods("GET")

	//Create Driver Account
	router.HandleFunc("/api/v1/driver/create/{driverid}", createDriver).Methods("POST")

	//Update Driver Account
	router.HandleFunc("/api/v1/driver/update/{driverid}", updateDriver).Methods("PUT")

	//Update Driver Status
	router.HandleFunc("/api/v1/driver/update/{driverid}/occupied", updateDriverStatus).Methods("PUT")

	//Assign Driver
	router.HandleFunc("/api/v1/drivers/", autoassigndriver)

	//Start Trip
	router.HandleFunc("/api/v1/driver/start/{tripid}", startTrip).Methods("PUT")

	//End Trip
	router.HandleFunc("/api/v1/driver/end/{tripid}/{driverid}", endTrip).Methods("PUT")

	//Port 3000
	fmt.Println("Listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

// View Driver - Get all Drivers
func getDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//Select all drivers from the Drivers table
	results, err := db.Query("select * from Drivers")
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {

		var d Driver

		err := results.Scan(&d.DriverID, &d.FirstName, &d.LastName, &d.PhoneNum, &d.Email, &d.NRIC, &d.LisenceNum, &d.DriverStatus)
		if err != nil {
			fmt.Println("failed to scan")
		}

		driverlist[d.DriverID] = d
	}
	data, _ := json.Marshal(map[string]map[string]Driver{"Drivers": driverlist})
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "%s\n", data)
}

// Populate Driver Map
func populateDriver(db *sql.DB) {
	//Select all drivers from the passenger table
	results, err := db.Query("select * from Drivers")
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {

		var d Driver

		err := results.Scan(&d.DriverID, &d.FirstName, &d.LastName, &d.PhoneNum, &d.Email, &d.NRIC, &d.LisenceNum, &d.DriverStatus)
		if err != nil {
			fmt.Println("failed to scan")
		}

		driverlist[d.DriverID] = d
	}
}

func createDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	// Retrieves the driverid specified in the API Endpoint which will be used to either Get, Post or Put.
	var driverID = ""
	driverID = params["driverid"]

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	populateDriver(db)
	_, exists := driverlist[driverID]
	if !exists {
		if r.Method == "POST" {
			// Create New Driver
			nd := Driver{}
			reqBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, &nd)
			_, err := db.Exec("INSERT INTO Drivers (DriverID, FirstName, LastName, PhoneNum, Email, NRIC, LisenceNum, DriverStatus) values(?, ?, ?, ?, ?, ?, ?, ?)", driverID, nd.FirstName, nd.LastName, nd.PhoneNum, nd.Email, nd.NRIC, nd.LisenceNum, nd.DriverStatus)
			if err != nil {
				panic(err.Error())
			}
			//Insert Into Map
			driverlist[driverID] = nd
			w.WriteHeader(http.StatusAccepted)

		} else if r.Method == "PUT" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Error! No driver account found.")

		}
	}
}

func updateDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	// Retrieves the driverid specified in the API Endpoint which will be used to either Get, Post or Put.
	var driverID = ""
	driverID = params["driverid"]

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	populateDriver(db)
	_, exists := driverlist[driverID]
	if exists {
		if r.Method == "PUT" {
			// Create Updated Driver
			ud := Driver{}
			reqBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, &ud)
			// Update the DB
			_, err := db.Exec("UPDATE Drivers SET FirstName=?, LastName=?, PhoneNum=?, Email=?, NRIC=?, LisenceNum=?, DriverStatus=? WHERE DriverID=?", ud.FirstName, ud.LastName, ud.PhoneNum, ud.Email, ud.NRIC, ud.LisenceNum, ud.DriverStatus, driverID)
			if err != nil {
				panic(err.Error())
			}
			// Update Map
			driverlist[driverID] = ud
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

func updateDriverStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	// Retrieves the driverid specified in the API Endpoint which will be used to either Get, Post or Put.
	var driverID = ""
	driverID = params["driverid"]

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	populateDriver(db)
	_, exists := driverlist[driverID]
	if exists {
		if r.Method == "PUT" {
			// Update the DB
			var driverstat = "Occupied"
			_, err := db.Exec("UPDATE Drivers SET DriverStatus=? WHERE DriverID=?", driverstat, driverID)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

func startTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var tripID = ""
	tripID = params["tripid"]

	if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Trip

			if err := json.Unmarshal(body, &data); err == nil {

				db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trips_db")
				if err != nil {
					fmt.Println("Failed to connect to DB")
					fmt.Println(err)
				}
				defer db.Close()

				var tripstat string = "Started"

				_, err = db.Exec("UPDATE Trips SET TripStatus=? WHERE TripID=?", tripstat, tripID)

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Trip is currently ongoing")
				}

				w.WriteHeader(http.StatusAccepted) //202

			} else {
				w.WriteHeader(http.StatusNotFound) //404
			}

		}
	}
}
func endTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var tripID = ""
	tripID = params["tripid"]
	// var driverID = ""
	// driverID = params["driverid"]
	// db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	// if err != nil {
	// 	fmt.Println("Failed to connect to DB")
	// 	fmt.Println(err)
	// }
	// defer db.Close()
	// var driverstat string = "Available"
	// _, err = db.Exec("UPDATE Drivers SET DriverStatus=? WHERE DriverID=?", driverstat, driverID)
	if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Trip

			if err := json.Unmarshal(body, &data); err == nil {

				db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trips_db")
				if err != nil {
					fmt.Println("Failed to connect to DB")
					fmt.Println(err)
				}
				defer db.Close()

				// var tripstat string = "Ended"
				fmt.Print(data.TripStatus)

				_, err = db.Exec("UPDATE Trips SET TripStatus=? WHERE TripID=?", data.TripStatus, tripID)

				if err != nil {
					fmt.Println(err)
				}
				w.WriteHeader(http.StatusAccepted) //202

			} else {
				w.WriteHeader(http.StatusNotFound) //404
			}

		}
	}
}

// Get a Random Driver with Status Available
func autoassigndriver(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	if err != nil {
		fmt.Println("failed to connect to db")
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM Drivers Where DriverStatus = 'Available' ORDER BY RAND() LIMIT 1")
	if err != nil {
		fmt.Println("failed to connect to insert")
		fmt.Println(err)
	} else {
		fmt.Println("Connected succesfully")
	}
	for results.Next() {

		var d Driver
		var driverid string

		err := results.Scan(&d.DriverID, &d.FirstName, &d.LastName, &d.PhoneNum, &d.Email, &d.NRIC, &d.LisenceNum, &d.DriverStatus)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(&driverid)
		}
		newdriver[d.DriverID] = d
	}

	data, _ := json.Marshal(map[string]map[string]Driver{"Selected driver": newdriver})
	fmt.Fprintf(w, "%s\n", data)

}

// Get Trips assigned to driver
func getDriverTrips(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trips_db")
	if err != nil {
		fmt.Println("Failed to connect to DB")
		fmt.Println(err)
	}
	defer db.Close()

	results, err := db.Query("Select * from Trips where DriverID = ?", params["driverid"])
	//db exec will return rows and primary key

	if err != nil {
		fmt.Println(err)
	}
	drivertriplist = map[string]Trip{}

	for results.Next() {
		var t Trip
		var tripid string

		err := results.Scan(&tripid, &t.StartPostalCode, &t.EndPostalCode, &t.TripStatus, &t.StartTime, &t.EndTime, &t.PassengerID, &t.DriverID)
		if err != nil {
			fmt.Println("failed to scan")
		}
		fmt.Println("ID no.: ", tripid, "Pick-Up Code:", t.StartPostalCode, "Drop-Off Code:", t.EndPostalCode, "Trip Status:", t.TripStatus, "Driver ID:", &t.DriverID)

		drivertriplist[tripid] = t
	}

	data, _ := json.Marshal(map[string]map[string]Trip{"Driver's Trips": drivertriplist})
	fmt.Fprintf(w, "%s\n", data)

}
