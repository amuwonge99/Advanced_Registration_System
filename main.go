package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Name         string
	Age          int
	Id           string
	RegisteredAt time.Time
}

var userSlice []User

func main() {
	fmt.Println("Welcome to the registration menu!")

	var command string

	for {
		fmt.Println("\nEnter one of the following commands: register, viewusers, searchbyname, delete, save, load. To leave, enter exit ")
		fmt.Scanln(&command)
		command = strings.ToLower(command)

		switch command {
		default:
			fmt.Println("Invalid command! Available commands: register, viewusers, searchbyname, load, save, delete, help, exit")
		case "help":
			fmt.Println("Available commands: register, viewusers, searchbyname, load, save, delete, help, exit")
		case "register":
			registerUser()
		case "viewusers":
			viewUsers()
		case "searchbyname":
			searchUsersByName()
		case "delete":
			deleteUserByID()
		case "save":
			saveToFile()
		case "load":
			loadFromFile()
		case "exit":
			fmt.Println("Goodbye!")
			return
		}
	}
}

func registerUser() {
	var name, ageStr string
	var age int

	fmt.Print("Enter your name (or 'cancel'): ")
	fmt.Scanln(&name)
	if strings.ToLower(name) == "cancel" {
		return
	}

	namePattern := regexp.MustCompile(`^[A-Za-z-]+$`)
	if !namePattern.MatchString(name) {
		fmt.Println("Invalid name. Only letters and hyphens are allowed.")
		return
	}

	fmt.Print("Enter your age (or 'cancel'): ")
	fmt.Scanln(&ageStr)
	if strings.ToLower(ageStr) == "cancel" {
		return
	}

	var err error
	age, err = strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Invalid age input. Please enter a valid number.")
		return
	}
	if age < 0 || age > 150 {
		fmt.Println("Invalid age. Please enter an age between 0 and 150.")
		return
	}

	newUser := User{
		Name:         name,
		Age:          age,
		Id:           uuid.New().String(),
		RegisteredAt: time.Now(),
	}
	userSlice = append(userSlice, newUser)
	fmt.Println("New user", name, "has registered at age", age, "ID:", newUser.Id)
}

func viewUsers() {
	if len(userSlice) == 0 {
		fmt.Println("No users found.")
		return
	}
	fmt.Println("Registered Users:")
	for _, user := range userSlice {
		fmt.Printf("Name: %s | Age: %d | ID: %s | Registered At: %s\n", user.Name, user.Age, user.Id, user.RegisteredAt.Format(time.RFC1123))
	}
}

func searchUsersByName() {
	var query string
	fmt.Print("Enter name to search: ")
	fmt.Scanln(&query)
	query = strings.ToLower(query)

	found := false
	for _, user := range userSlice {
		if strings.ToLower(user.Name) == query {
			fmt.Printf("User %s found with ID: %s\n", user.Name, user.Id)
			found = true
		}
	}
	if !found {
		fmt.Println("No user found with that name.")
	}
}

func deleteUserByID() {
	var deleteID string
	fmt.Print("Enter ID of the user to delete: ")
	fmt.Scanln(&deleteID)

	userIDFound := false
	for i, user := range userSlice {
		if user.Id == deleteID {
			userSlice = append(userSlice[:i], userSlice[i+1:]...)
			userIDFound = true
			fmt.Println("User with ID", deleteID, "has been deleted.")
			break
		}
	}

	if !userIDFound {
		fmt.Println("User", deleteID, "not found.")
	}
}

func saveToFile() error {
	var fileName string
	fmt.Print("Enter file name to save (e.g., users.json): ")
	fmt.Scanln(&fileName)

	if _, err := os.Stat(fileName); err == nil {
		fmt.Println("File already exists. Overwrite? (yes/no): ")
		var confirm string
		fmt.Scanln(&confirm)
		if strings.ToLower(confirm) != "yes" {
			fmt.Println("Aborted saving.")
			return nil
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file (%s): %s\n", fileName, err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(userSlice)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return err
	}
	fmt.Println("Data saved successfully.")
	return nil
}

func loadFromFile() error {
	var fileName string
	fmt.Print("Enter file name to load from (e.g., users.json): ")
	fmt.Scanln(&fileName)

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("ERROR: File does not exist or can't be read:", fileName)
		return err
	}

	var newData []User
	err = json.Unmarshal(file, &newData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return err
	}

	userSlice = append(userSlice, newData...)
	fmt.Printf("Loaded %d users.\n", len(newData))
	return nil
}
