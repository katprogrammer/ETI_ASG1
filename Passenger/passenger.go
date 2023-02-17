package main

// Import packages required
import (
	"database/sql"
	"encoding/json"
	"fmt"
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

type AllPassenger struct {
	Passengers map[string]Passenger `json:"Passenger"`
}

var passengerlist = map[string]Passenger{}
var triplist = map[string]Trip{}

func main() {
	// Connect to Router
	router := mux.NewRouter()

	// View all Passengers
	router.HandleFunc("/api/v1/passenger/view/", getPassenger).Methods("GET")

	// Create new Passenger Account
	router.HandleFunc("/api/v1/passenger/create/{passengerid}", createPassenger).Methods("POST")

	// Update Passenger Account
	router.HandleFunc("/api/v1/passenger/update/{passengerid}", updatePassenger).Methods("PUT")

	// Create new Trip
	router.HandleFunc("/api/v1/passenger/trip/{tripid}/{passengerid}/{driverid}", createTrip).Methods("GET", "POST")

	// List all Passenger Trips
	router.HandleFunc("/api/v1/trip/{passengerid}", getTrip).Methods("GET")

	//Port 5000
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

// View Passenger - Get all passengers
func getPassenger(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//Select all passengers from the passenger table
	results, err := db.Query("select * from Passengers")
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {

		var p Passenger

		err := results.Scan(&p.PassengerID, &p.FirstName, &p.LastName, &p.PhoneNum, &p.Email)
		if err != nil {
			fmt.Println("failed to scan")
		}

		passengerlist[p.PassengerID] = p
	}
	data, _ := json.Marshal(map[string]map[string]Passenger{"Passengers": passengerlist})
	fmt.Fprintf(w, "%s\n", data)
}

func createPassenger(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

}

func updatePassenger(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

}

func createTrip(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

}

// View Trips - Get all passenger's trips
func getTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trips_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//Select all trips based on passenger id
	results, err := db.Query("select * from Trips where PassengerId = ? ORDER BY StartTime Desc", params["passengerid"])
	if err != nil {
		fmt.Println(err)
	}
	for results.Next() {
		var t Trip
		var TripID string

		err := results.Scan(&TripID, &t.StartPostalCode, &t.EndPostalCode, &t.TripStatus, &t.StartTime, &t.EndTime, &t.PassengerID, &t.DriverID)
		if err != nil {
			fmt.Println("failed to scan")
			fmt.Println(err)
		}

		triplist[TripID] = t

	}
	data, _ := json.Marshal(map[string]map[string]Trip{"Trip": triplist})

	fmt.Fprintf(w, "%s\n", data)
}
