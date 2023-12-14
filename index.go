package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

func generatePassword(length int) string {
    rand.Seed(time.Now().UnixNano())
    chars := "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM!@#$%^&*()_-+={}[]|:;',<.>/?"
    password := strings.Builder{}
    for i := 0; i < length; i++ {
        password.WriteByte(chars[rand.Intn(len(chars))])
    }
    return password.String()
}

func main() {
    for {
        fmt.Print("Enter the password length (8-75): ")
        var length int
        _, err := fmt.Scan(&length)
        if err != nil || length < 8 || length > 75 {
            fmt.Println("Invalid input, enter a number between 8 and 75.")
            continue
        }
        password := generatePassword(length)
        fmt.Println("Your new password:", password)
        break
    }
}
