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
	PassengerID string `json:"Passenger Id"`
	FirstName   string `json:"First Name"`
	LastName    string `json:"Last Name"`
	PhoneNum    string `json:"Mobile Number"`
	Email       string `json:"Email Address"`
}

type AllPassenger struct {
	Passengers map[string]Passenger `json:"Passenger"`
}

var passengerlist = map[string]Passenger{}

func main() {
	// Connect to Router
	router := mux.NewRouter()

	// View all Passengers
	router.HandleFunc("/api/v1/passenger/view/", getPassenger).Methods("GET")

	// Create new Passenger Account
	router.HandleFunc("/api/v1/passenger/{passengerid}", createPassenger).Methods("POST", "PUT")

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
		fmt.Println("ID: ", p.PassengerID, "Passenger Name: ", p.FirstName, p.LastName, "Mobile Nomber:", p.PhoneNum, "Email:", p.Email)

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

func createTrip(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

}

func getTrip(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

}
