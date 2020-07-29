package dbconnection

import (
	"log"
)

func init() {

	log.Println("Initializing service.")
	err := DummyData()
	if err != nil {
		log.Fatalln("Error while populating movie dummy data.")
	}
	err = AddUsers()
	if err != nil {
		log.Fatalln("Error while populating user dummy data.")
	}
}
