# User Management System

*Duration:* Approximately 1 full day

*Concepts:*
- Struct Definition and Usage
- Function Design and Implementation
- File Input/Output (JSON Serialization/Deserialization)
- Control Flow (Menu system, conditional logic)
- Data Structures (Slices)

## Task:

Develop a command-line interface (CLI) user management tool in Go. This tool will allow users to register, view, search, and delete user 
information.

## Requirements:

1. **User Struct:** Define a `User` struct with the following fields:
    * `Name`: String (User's name)
    * `Age`: Integer (User's age)
    * `ID`: Integer (Unique user identifier)
    * `RegisteredAt`: String (Timestamp of registration - can be a simple string like "YYYY-MM-DD HH:MM:SS")

2. **Data Storage:**  Replace the previously used parallel slices for storing user data with a single slice of `User` structs.

3. **Menu System:** Implement a CLI menu system with the following options:
    * **Register a new user:**  Prompt the user for name, age, and ID, create a new `User` struct, and add it to the slice.
    * **View all users:** Display the information of all registered users.
    * **Search by name:**  Prompt the user for a name, then display all users whose names contain that search string (case-insensitive).
    * **Delete by ID:** Prompt the user for an ID, then remove the user with that ID from the slice.  Handle the case where the ID is not found.
    * **Save/Load users to/from a JSON file:**
        * **Save:**  Serialize the slice of `User` structs to a JSON file.
        * **Load:**  Deserialize a JSON file into the slice of `User` structs.  This should happen on startup if a filename is provided.

4. **Modular Code:** Use functions to organize your code.  Suggested functions:
    * `registerUser()`:  Handles user registration.
    * `printUsers()`:  Prints the information of all users.
    * `searchUsersByName()`: Searches and prints users by name.
    * `deleteUserByID()`: Deletes a user by ID.
    * `saveToFile()`: Saves the user data to a JSON file.
    * `loadFromFile()`: Loads user data from a JSON file.
    * `main()`: The main function, which drives the menu and calls the other functions.

## Stretch Goals:

* **Sort Users:** Sort the slice of users by name (alphabetically) before displaying them in the "View all users" option.
* **Login Simulation:** Implement a basic login simulation.  Prompt the user for a name and ID.  If a user with that name and ID exists in the system, display a "Login successful" message.
* **CLI Flag:** Use a command-line flag (e.g., using a library like `flag` there is a lab on this if you want to see an example!) to allow the user to specify the JSON file to load at startup.  If no file is provided, start with an empty user list.
* **Test Cases:** Provide a set of test cases that cover various scenarios (e.g., adding users, deleting users, searching by name) to ensure the program behaves as expected.



## Notes for Helping Students:

* **Design First:** Encourage the students to *sketch out* the overall structure of the program *before* writing code.  What functions will they need? How will the menu work? What 
data structures will they use?
* **Error Handling:**  Discuss the importance of error handling (e.g., what happens if the JSON file is invalid, or a user tries to delete a non-existent ID).
* **Modularity:** Emphasize the benefits of writing small, well-defined functions.  This makes the code easier to read, test, and maintain.
* **JSON Libraries:**  Point them to the appropriate JSON serialization/deserialization libraries for the chosen language.
* **String Comparison:**  Remind them that string comparisons are case-sensitive by default and how to perform case-insensitive comparisons.
* **Unique IDs:** Discuss how they might ensure that user IDs are unique.  (Simple solution: check if an ID already exists before adding a new user.)
* **Data persistence** Talk about the fact that if they don't save to file the data will be lost.
