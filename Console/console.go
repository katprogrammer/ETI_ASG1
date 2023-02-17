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

func main() {
outer:
	for {
		fmt.Println("\n")

		fmt.Println(" ==== Main Menu ==== \n",
			"\n",
			"<----Passenger Menu---->\n",
			"1.  Create New Passenger \n",       //create new passsenger
			"2.  Update Existing Passenger \n",  //update existing passenger
			"3.  View All Passenger - Admin \n", //view all passenger
			"4.  Request for New Trip \n",       //create new trip
			"5.  View Trip History \n",          //view passenger trip history
			"\n",
			"\n",
			"<----Driver Menu---->\n",
			"6.  Create New Driver\n",             //create new driver, driver status default set to busy
			"7.  Update Existing Driver\n",        //update all existing driver information, except for IcNo.
			"8.  View All Driver - Admin \n",      //view all driver
			"9.  Go Online  (Available Mode)\n",   //update status to available (eligible for trip assignment)
			"10. Go Offline (Busy Mode)\n",        //update status to buys (not eligible for trip assignment)
			"11. View Driver's Assigned Trips \n", //view trips assigned to driver, then able to start end or exit
			"0.  Exit")

		fmt.Print("Please select an option: ")

		var option int
		fmt.Scanf("%d\n", &option)

		switch option {

		case 0:
			break outer
		//Passenger
		case 1:
			getPassenger()
		default:
			fmt.Println("Sorry, we didn't catch that.")

		}
	}
}

func getPassenger() {

	fmt.Println("\n")
	fmt.Println("=== View All Passenger ===")

	//connect to client
	client := &http.Client{}

	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/passenger/view/", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {

				var res Passengers

				json.Unmarshal(body, &res)
				fmt.Println("=== Passenger Info ===")
				for k, v := range res.Passengers {
					fmt.Println("\n")
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
