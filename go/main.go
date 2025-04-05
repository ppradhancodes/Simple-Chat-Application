package main

import (
    "bufio"
    "chat-app/handlers"
    "chat-app/models"
    "chat-app/utils"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println("Simple Chat Application")
    chatHandler := handlers.NewChatHandler()
    scanner := bufio.NewScanner(os.Stdin)
    var currentUser *models.User

    for {
        if currentUser == nil {
            fmt.Print("Enter username (or 'exit' to quit): ")
            if !scanner.Scan() {
                break
            }
            username := strings.TrimSpace(scanner.Text())

            if username == "exit" {
                break
            }

            user, err := chatHandler.RegisterOrLogin(username)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                continue
            }

            currentUser = &user
            fmt.Printf("Welcome, %s!\n", username)
        } else {
            fmt.Print("\nCommands:\n1. Send message\n2. View messages\n3. Search messages\n4. List users\n5. Logout\nChoice: ")
            if !scanner.Scan() {
                break
            }
            choice := strings.TrimSpace(scanner.Text())

            switch choice {
            case "1":
                fmt.Print("Enter recipient username: ")
                if !scanner.Scan() {
                    break
                }
                recipient := strings.TrimSpace(scanner.Text())

                fmt.Print("Enter message: ")
                if !scanner.Scan() {
                    break
                }
                content := strings.TrimSpace(scanner.Text())

                err := chatHandler.SendMessage(currentUser.ID, recipient, content)
                if err != nil {
                    fmt.Printf("Error: %v\n", err)
                } else {
                    fmt.Println("Message sent!")
                }

            case "2":
                messages := chatHandler.GetMessages(currentUser.ID)
                if len(messages) == 0 {
                    fmt.Println("No messages.")
                } else {
                    for _, msg := range messages {
                        sender, exists := chatHandler.GetUser(msg.SenderID)
                        senderName := "Unknown"
                        if exists {
                            senderName = sender.Username
                        }
                        utils.PrintMessage(senderName, msg.Content, msg.Timestamp)
                    }
                }

            case "3":
                fmt.Print("Enter search keyword: ")
                if !scanner.Scan() {
                    break
                }
                keyword := strings.TrimSpace(scanner.Text())

                messages := chatHandler.SearchMessages(keyword)
                if len(messages) == 0 {
                    fmt.Println("No messages found.")
                } else {
                    for _, msg := range messages {
                        sender, exists := chatHandler.GetUser(msg.SenderID)
                        senderName := "Unknown"
                        if exists {
                            senderName = sender.Username
                        }
                        utils.PrintMessage(senderName, msg.Content, msg.Timestamp)
                    }
                }

            case "4":
                users := chatHandler.ListUsers()
                fmt.Println("\nRegistered users:")
                for _, user := range users {
                    fmt.Printf("- %s\n", user.Username)
                }

            case "5":
                fmt.Printf("Goodbye, %s!\n", currentUser.Username)
                currentUser = nil

            default:
                fmt.Println("Invalid choice")
            }
        }
    }
} 