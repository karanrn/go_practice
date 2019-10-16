package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"fmt"
)

// Recover
func recovery(){
	if r := recover(); r!= nil {
        fmt.Println("recovered from ", r)
    }
}

// Checking logging levels
func logger(){
	// recover method
	defer recovery()

	// Log to console
	log.Print("Console logging.")

	// Log to file
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
        log.Fatal(err)
	}
	
	// Use defer to close the file before exiting the program
	defer file.Close()

	/*
		Debug -> Info -> Warn -> Error -> Fatal -> Panic

		Prefer Panic over Fatal and Error because they exit : os.exit()
		Panic can be recovered. Be consious while using fatal and error.
		Panic will allow defer statements to be executed prior to exiting
		from the function.
	*/
	log.SetOutput(file)

	// Set the formatter so it can be moved to centralized logging
	log.SetFormatter(&log.JSONFormatter{})
	
	// Set log level
	log.SetLevel(log.InfoLevel)

	log.Debug("Debug the previous error.")
	log.Info("Logging to a file.")
	log.Warn("Just a warning.")
	//log.Error("Error logging.") // Calls os.exit(1) after logging	
	//log.Fatal("Calling panic()") // Calls os.exit(1) after logging
	log.Panic("Panicked!!!!")

}
// Entry point
func main(){
	logger()
	fmt.Println("Recovered from panic!!")
}