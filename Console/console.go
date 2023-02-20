package main

//Import
import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Passenger struct {
	PassengerID int    `json:"Passenger Id"`
	FirstName   string `json:"First Name"`
	LastName    string `json:"Last Name"`
	PhoneNum    string `json:"Phone Number"`
	Email       string `json:"Email Address"`
}

type Passengers struct {
	Passengers map[string]Passenger `json:"Passengers"`
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

type Drivers struct {
	Drivers map[string]Driver `json:"Drivers"`
}

var randomdriver Driver
var randomdriverid string

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
			"8)  End Trip\n",               //update status to available
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
		case 1:
			createPassenger()
		case 2:
			updatePassenger()
		case 3:
			createTrip()
		case 4:
			getPassengerTrips()
		case 5:
			createDriver()
		case 6:
			updateDriver()
		case 7:
			startTrip()
		case 8:
			endTrip()
		case 9:
			getPassenger()
		case 10:
			getDrivers()
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
					fmt.Println("Phone Number : ", v.PhoneNum)
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

	fmt.Print("\n")
	fmt.Println("=== View All Drivers ===")

	//connect to client
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/driver/view/", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {

				var res Drivers

				json.Unmarshal(body, &res)
				fmt.Print("\n")
				fmt.Println("=== Driver Info ===")
				for k, v := range res.Drivers {
					fmt.Print("\n")
					fmt.Println("Driver ID : ", k, " ")
					fmt.Println("First Name : ", v.FirstName)
					fmt.Println("Last Name : ", v.LastName)
					fmt.Println("Phone Number : ", v.PhoneNum)
					fmt.Println("Email : ", v.Email)
					fmt.Println("NRIC : ", v.NRIC)
					fmt.Println("Car License Number : ", v.LisenceNum)
					fmt.Println("Status : ", v.DriverStatus)
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
func createPassenger() {

	reader := bufio.NewReader(os.Stdin)

	var np Passenger
	fmt.Print("\n")
	fmt.Println("=== New Passenger Creation ===")

	var randpid string
	var randit int
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randit = r.Intn(100)
	randpid = "P" + strconv.Itoa(randit)

	fmt.Print("Enter First Name: ")
	firstname, _ := reader.ReadString('\n')
	np.FirstName = strings.TrimSpace(firstname)

	fmt.Print("Enter Last Name: ")
	lastname, _ := reader.ReadString('\n')
	np.LastName = strings.TrimSpace(lastname)

	fmt.Print("Enter Phone Number: ")
	phonenum, _ := reader.ReadString('\n')
	np.PhoneNum = strings.TrimSpace(phonenum)

	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')
	np.Email = strings.TrimSpace(email)

	postBody, _ := json.Marshal(np)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/passenger/create/"+randpid, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Print("\n")
				fmt.Println("* New Passenger with ID : ", randpid, " created! *")
			} else if res.StatusCode == 409 {
				fmt.Println("* Error - Passenger", randpid, "already exists! *")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

func createDriver() {

	reader := bufio.NewReader(os.Stdin)

	var nd Driver

	fmt.Println("=== Create New Driver Account ===")
	var randdid string
	var randit int
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randit = r.Intn(5000)
	randdid = "D" + strconv.Itoa(randit)

	fmt.Print("Enter First Name: ")
	firstname, _ := reader.ReadString('\n')
	nd.FirstName = strings.TrimSpace(firstname)

	fmt.Print("Enter Last Name: ")
	lastname, _ := reader.ReadString('\n')
	nd.LastName = strings.TrimSpace(lastname)

	fmt.Print("Enter Phone Number: ")
	phonenum, _ := reader.ReadString('\n')
	nd.PhoneNum = strings.TrimSpace(phonenum)

	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')
	nd.Email = strings.TrimSpace(email)

	fmt.Print("Enter IC: ")
	nric, _ := reader.ReadString('\n')
	nd.NRIC = strings.TrimSpace(nric)

	fmt.Print("Enter Car License: ")
	lisencenum, _ := reader.ReadString('\n')
	nd.LisenceNum = strings.TrimSpace(lisencenum)

	nd.DriverStatus = "Available"

	postBody, _ := json.Marshal(nd)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/driver/create/"+randdid, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("* Driver with ID: ", randdid, "created! *")
			} else if res.StatusCode == 409 {
				fmt.Println("* Error - Driver", randdid, "already exists! *")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

// Function - Update Existing Passenger
func updatePassenger() {

	reader := bufio.NewReader(os.Stdin)

	var updatePassenger Passenger

	fmt.Print("\n")
	fmt.Println("=== Update Passenger Details ===")
	fmt.Print("\n")
	fmt.Print("Enter chosen Passenger ID:")
	var passengerid string
	fmt.Scanf("%v\n", &passengerid)

	fmt.Print("Enter First Name: ")
	firstname, _ := reader.ReadString('\n')
	updatePassenger.FirstName = strings.TrimSpace(firstname)

	fmt.Print("Enter Last Name: ")
	lastname, _ := reader.ReadString('\n')
	updatePassenger.LastName = strings.TrimSpace(lastname)

	fmt.Print("Enter Phone Number: ")
	phonenum, _ := reader.ReadString('\n')
	updatePassenger.PhoneNum = strings.TrimSpace(phonenum)

	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')
	updatePassenger.Email = strings.TrimSpace(email)

	postBody, _ := json.Marshal(updatePassenger)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/passenger/update/"+passengerid, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Print("\n")
				fmt.Println("* Passenger ", passengerid, " updated successfully! *")
			} else if res.StatusCode == 404 {
				fmt.Println("* Error - Passenger", passengerid, "does not exist! *")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

func updateDriver() {

	reader := bufio.NewReader(os.Stdin)

	var ud Driver

	fmt.Println("=== Update Driver Account Details ===")
	fmt.Print("\n")
	fmt.Print("Enter chosen Driver ID:")
	var driverid string
	fmt.Scanf("%v\n", &driverid)

	fmt.Print("Enter First Name: ")
	firstname, _ := reader.ReadString('\n')
	ud.FirstName = strings.TrimSpace(firstname)

	fmt.Print("Enter Last Name: ")
	lastname, _ := reader.ReadString('\n')
	ud.LastName = strings.TrimSpace(lastname)

	fmt.Print("Enter Phone Number: ")
	phonenum, _ := reader.ReadString('\n')
	ud.PhoneNum = strings.TrimSpace(phonenum)

	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')
	ud.Email = strings.TrimSpace(email)

	fmt.Print("Enter IC: ")
	nric, _ := reader.ReadString('\n')
	ud.NRIC = strings.TrimSpace(nric)

	fmt.Print("Enter Car License: ")
	lisencenum, _ := reader.ReadString('\n')
	ud.LisenceNum = strings.TrimSpace(lisencenum)

	ud.DriverStatus = "Available"

	postBody, _ := json.Marshal(ud)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/driver/update/"+driverid, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Print("\n")
				fmt.Println("* Driver", driverid, "updated successfully! *")
			} else if res.StatusCode == 404 {
				fmt.Println("* Error - Driver", driverid, "does not exist! *")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

// Function - Create Trip Request
func createTrip() {
	resp, err := http.Get("http://localhost:3000/api/v1/drivers/")
	if err != nil {
		fmt.Println("All Drivers are busy")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var assigndriver map[string]map[string]Driver
	json.Unmarshal([]byte(body), &assigndriver)

	for key, element := range assigndriver["Selected driver"] {
		randomdriverid = key
		randomdriver = element
	}
	reader := bufio.NewReader(os.Stdin)
	var newtrip Trip
	var randtripid string
	var randit int
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randit = r.Intn(10000)
	randtripid = "T" + strconv.Itoa(randit)

	fmt.Print("Enter Passenger ID : ")
	var passengerid string
	fmt.Scanf("%v\n", &passengerid)

	fmt.Print("Enter Pickup Code: ")
	pickupcode, _ := reader.ReadString('\n')
	newtrip.StartPostalCode = strings.TrimSpace(pickupcode)

	fmt.Print("Enter Dropoff Code: ")
	dropoffcode, _ := reader.ReadString('\n')
	newtrip.EndPostalCode = strings.TrimSpace(dropoffcode)

	newtrip.TripStatus = "Requested"

	postBody, _ := json.Marshal(newtrip)
	resBody := bytes.NewBuffer(postBody)
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/passenger/trip/"+randtripid+"/"+passengerid+"/"+randomdriverid, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("* New Trip:", randtripid, "created successfully! *")

				var updateBusy Driver

				updateBusy.DriverStatus = "Occupied"

				driverpostBody, _ := json.Marshal(updateBusy)

				if req, err := http.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/driver/update/"+randomdriverid+"/occupied", bytes.NewBuffer(driverpostBody)); err == nil {

					if res, err := client.Do(req); err == nil {
						if res.StatusCode == 202 {
							fmt.Println("* A new driver is assigned to your trip! * ")
						} else if res.StatusCode == 404 {
							fmt.Println("* Driver does not exist! * ")
						}
					} else {
						fmt.Println(2, err)
					}
				} else {
					fmt.Println(3, err)
				}
			} else if res.StatusCode == 409 {
				fmt.Println("*** Error - Passenger", passengerid, "already in Ongoing Trip! *** ")
			} else if res.StatusCode == 404 {
				fmt.Println("*** Error - Passenger", passengerid, "does not exist! *** ")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

// Function - Start a Trip
func startTrip() {

	var updateTrip Trip
	var driverid string
	var tripid string
	viewDriverTrips()
	fmt.Print("Enter Driver ID : ")
	fmt.Scanf("%v\n", &driverid)
	fmt.Print("Enter Trip ID : ")
	fmt.Scanf("%v\n", &tripid)

	fmt.Print("Changing Status...")

	updateTrip.TripStatus = "Started"

	postBody, _ := json.Marshal(updateTrip)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/driver/start/"+tripid, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("* Trip", tripid, "has started! *")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - Driver", driverid, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

// Function - End a Trip
func endTrip() {
	var updateTrip Trip

	var driverid string
	var tripid string

	viewDriverTrips()
	fmt.Print("Enter Driver ID : ")
	fmt.Scanf("%v\n", &driverid)
	fmt.Print("Enter Trip ID : ")
	fmt.Scanf("%v\n", &tripid)

	fmt.Print("Changing Status...")

	updateTrip.TripStatus = "Ended"

	postBody, _ := json.Marshal(updateTrip)

	//parse multiple strings as parameters
	url := fmt.Sprintf("http://localhost:3000/api/v1/driver/end/" + tripid + "/" + driverid)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("* Trip", tripid, "has ended! *")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - Driver", driverid, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}
func viewDriverTrips() {

	var driverid string
	fmt.Print("Please enter Driver ID to view trips: ")
	fmt.Scanf("%v\n", &driverid)

	client := &http.Client{}

	if req, err := http.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/driver/trips/"+driverid, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {

				var res map[string]map[string]Trip

				json.Unmarshal(body, &res)

				fmt.Println("=== Trip Info ===")
				for k, v := range res["Driver's Trips"] {

					fmt.Println("Trip ID : ", k, " ")
					fmt.Println("Pickup Code : ", v.StartPostalCode)
					fmt.Println("Dropoff Code : ", v.EndPostalCode)
					fmt.Println("Trip Status : ", v.TripStatus)
					fmt.Print("\n")
				}
			}
		}
	}

}
