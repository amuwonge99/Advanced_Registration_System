# Basic User Registration

*Duration:* ~2 hours (for 2 students)

*Concepts:*
- Variables
- `if` statements
- `for` loops
- Slices
- Using an external package (`uuid`)
- Basic Input/Output
- Error Handling (minimal)

## Task:

Create a command-line program that allows users to register up to 5 usernames with age validation. The program should prompt the user for a name and age.  Valid users will have their 
information stored, and a unique ID generated for them.

## Requirements:

1. **Prompt for Input:** Ask the user to enter a name and age.  Use appropriate prompts (e.g., "Enter your name: ", "Enter your age: ").
2. **Age Validation:** If the entered age is less than 13, print a message like "You must be 13 or older to register." and skip registering that user.  Move to the next iteration of 
the loop.
3. **Unique ID Generation:** For each *valid* user (age 13 or older), generate a unique ID using the `uuid` package.
4. **Data Storage:** Store the user data in parallel slices:
    * `names []string` (to store usernames)
    * `ages []int` (to store ages)
    * `ids []string` (to store unique IDs)
5. **Limited Registrations:**  The program should allow a maximum of 5 valid registrations.
6. **Output Registered Users:** After the registration process (or after 5 valid users are registered), print all registered users in a clear format.  For example:
   ```
   Registered Users:
   Name: Alice, Age: 30, ID: a1b2c3d4-e5f6-7890-1234-567890abcdef
   Name: Bob, Age: 18, ID: fedcba09-8765-4321-0987-654321fedcba
   ```

## Stretch Goals:

* **Prevent Duplicate Names:** Before registering a user, check if the name already exists in the `names` slice. If it does, print a message like "Name already registered. Please 
choose a different name." and ask for input again for that iteration.
* **Early Exit:** Allow the user to type "exit" (case-insensitive) at the name or age prompt to terminate the registration process early.
* **Error Handling (Improved):**  Handle potential errors when converting the age input from a string to an integer.  If the input is not a valid number, print an error message and 
ask for the age again.



**Suggestions for Students:**

1. **Divide and Conquer:** One student can focus on the input/validation/looping, while the other focuses on the `uuid` integration and slice management. They should collaborate to 
integrate the pieces.
2. **Incremental Development:**  Start with the basic requirements. Get the input, age validation, and slice storage working before attempting the stretch goals.
3. **Test Frequently:**  After each small piece of code is added, test it to ensure it's working correctly.
4. **Comments:**  Encourage them to write comments in their code to explain what each section does.  This will help them understand their code later and will also make it easier for 
you to assess their work.
5. **Package Import:** Remind them that they need to import the `uuid` package: `import "github.com/google/uuid"`
6. **String to Int Conversion:** Remind them that the age input will be a string and they need to convert it to an integer using `strconv.Atoi()`. They should handle the potential 
error from this conversion.


**Example Code Snippet (to get them started - not a complete solution):**

```go
package main

import (
        "fmt"
        "github.com/google/uuid"
        "strconv"
)

func main() {
        names := []string{}
        ages := []int{}
        ids := []string{}

        for i := 0; i < 5; i++ {
                var name string
                var ageStr string

                fmt.Print("Enter your name (or 'exit'): ")
                fmt.Scanln(&name)

                if name == "exit" {
                        break
                }

                fmt.Print("Enter your age: ")
                fmt.Scanln(&ageStr)

                age, err := strconv.Atoi(ageStr)
                if err != nil {
                        fmt.Println("Invalid age. Please enter a number.")
                        continue // Skip to the next iteration
                }

                if age < 13 {
                        fmt.Println("You must be 13 or older to register.")
                        continue
                }

                id := uuid.New().String()

                names = append(names, name)
                ages = append(ages, age)
                ids = append(ids, id)
        }

        fmt.Println("\nRegistered Users:")
        for i := 0; i < len(names); i++ {
                fmt.Printf("Name: %s, Age: %d, ID: %s\n", names[i], ages[i], ids[i])
        }
}
```