package main

//Import
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Passenger struct {
	PassengerID int    `json:"Passenger Id"`
	FirstName   string `json:"First Name"`
	LastName    string `json:"Last Name"`
	PhoneNum    string `json:"Mobile Number"`
	Email       string `json:"Email Address"`
}

type Passengers struct {
	Passengers map[string]Passenger `json:"Passengers"`
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

type Trips struct {
	Trips map[string]Trip `json:"Trips"`
}

func main() {
outer:
	for {
		fmt.Print("\n")
		fmt.Println(" ====== Main Menu ====== \n",
			"\n",
			"-----Passenger Menu-----\n",
			"1)  Create Passenger Account \n", //create new passsenger
			"2)  Update Account Details\n",    //update existing passenger
			"3)  Request Trip \n",             //create new trip
			"4)  View Trip History \n",        //view passenger trip history
			"\n",
			"\n",
			"-----Driver Menu-----\n",
			"5)  Create Driver Account\n",  //create new driver, driver status default set to busy
			"6)  Update Account Details\n", //update all existing driver information, except for IcNo.
			"7)  Start Trip\n",             //update status to unavailable
			"8) End Trip\n",                //update status to available
			"\n",
			"\n",
			"-----Admin Menu-----\n",
			"9)  View all Passengers\n", //view all passengers
			"10)  View all Drivers\n",   //view all drivers
			"0)  Exit",
			"\n",
			"\n")

		fmt.Print("Please select an option: ")

		var option int
		fmt.Scanf("%d\n", &option)

		switch option {

		case 0:
			break outer
		//Passenger
		case 4:
			getPassengerTrips()
		case 9:
			getPassenger()
		default:
			fmt.Println("Invalid Option!")

		}
	}
}

func getPassenger() {

	fmt.Print("\n")
	fmt.Println("=== View All Passenger ===")

	//connect to client
	client := &http.Client{}

	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/passenger/view/", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {

				var res Passengers

				json.Unmarshal(body, &res)
				fmt.Print("\n")
				fmt.Println("=== Passenger Info ===")
				for k, v := range res.Passengers {
					fmt.Println("Passenger ID : ", k, " ")
					fmt.Println("First Name : ", v.FirstName)
					fmt.Println("Last Name : ", v.LastName)
					fmt.Println("MobileNumber : ", v.PhoneNum)
					fmt.Println("Email : ", v.Email)
					fmt.Println("\n")

				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func getPassengerTrips() {
	fmt.Print("\n")
	fmt.Println("=== View Passenger Trips ===")

	//connect to client
	client := &http.Client{}
	var passengerid string
	fmt.Print("Please enter Passenger ID: ")
	fmt.Scanf("%v\n", &passengerid)
	fmt.Print("\n")
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/trip/"+passengerid, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {

				var res map[string]map[string]Trip

				json.Unmarshal(body, &res)
				fmt.Println("=== Passenger Trips ===")
				fmt.Print("\n")
				for k, v := range res["Trip"] {
					fmt.Println("Trip ID : ", k, " ")
					fmt.Println("Pickup Postal Code : ", v.StartPostalCode)
					fmt.Println("Dropoff Postal Code : ", v.EndPostalCode)
					fmt.Println("Trip Status : ", v.TripStatus)
					fmt.Println("Trip End Date : ", v.EndTime)
					fmt.Print("\n")

				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func getDrivers() {

	fmt.Println("\n")
	fmt.Println("=== View All Passenger ===")

	//connect to client
	client := &http.Client{}

	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/passenger/view/", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {

				var res Passengers

				json.Unmarshal(body, &res)
				fmt.Println("\n")
				fmt.Println("=== Passenger Info ===")
				for k, v := range res.Passengers {
					fmt.Println("Passenger ID : ", k, " ")
					fmt.Println("First Name : ", v.FirstName)
					fmt.Println("Last Name : ", v.LastName)
					fmt.Println("MobileNumber : ", v.PhoneNum)
					fmt.Println("Email : ", v.Email)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}
