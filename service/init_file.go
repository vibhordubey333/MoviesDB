package service

import (
	"log"
)

func init() {

	log.Println("Initializing service.")
	err := DummyData()
	if err != nil {
		log.Fatalln("Error while populating dummy data.")
	}

}
