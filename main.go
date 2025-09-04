package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Struct to represent our chatbot
type ChatBot struct {
	name    string
	version string
}

// Method to make the bot introduce itself
func (bot *ChatBot) introduce() {
	fmt.Printf("🤖 Hello! I'm %s v%s\n", bot.name, bot.version)
	fmt.Println("🤖 I'm a simple chatbot built with Go!")
	fmt.Println("🤖 Type 'quit' or 'exit' to end our conversation.")
	fmt.Println("🤖 Let's chat!")
	fmt.Println(strings.Repeat("-", 50))
}

// Method to process user input and generate responses
func (bot *ChatBot) respond(input string) string {
	// Convert input to lowercase for easier matching
	input = strings.ToLower(strings.TrimSpace(input))
	
	// Use switch statement to handle different inputs
	switch {
	case strings.Contains(input, "hello") || strings.Contains(input, "hi"):
		return "🤖 Hello there! Nice to meet you!"
		
	case strings.Contains(input, "how are you"):
		return "🤖 I'm doing great! Thanks for asking. How are you?"
		
	case strings.Contains(input, "what's your name") || strings.Contains(input, "who are you"):
		return fmt.Sprintf("🤖 I'm %s, your friendly Go chatbot!", bot.name)
		
	case strings.Contains(input, "time"):
		currentTime := time.Now().Format("15:04:05")
		return fmt.Sprintf("🤖 The current time is: %s", currentTime)
		
	case strings.Contains(input, "date"):
		currentDate := time.Now().Format("January 2, 2006")
		return fmt.Sprintf("🤖 Today's date is: %s", currentDate)
		
	case strings.Contains(input, "go") || strings.Contains(input, "golang"):
		return "🤖 Go is awesome! It's fast, simple, and great for building applications!"
		
	case strings.Contains(input, "help"):
		return "🤖 You can ask me about:\n   - My name\n   - Current time/date\n   - How I'm doing\n   - About Go language\n   - Or just say hello!"
		
	case input == "":
		return "🤖 You didn't say anything. Try typing something!"
		
	default:
		// Array of default responses
		responses := []string{
			"🤖 That's interesting! Tell me more.",
			"🤖 I see! What else would you like to know?",
			"🤖 Hmm, I'm still learning. Can you ask me something else?",
			"🤖 That's cool! I'm here to chat whenever you want.",
		}
		// Return a random response (simplified - just use first one for now)
		return responses[0]
	}
}

// Function to check if user wants to quit
func shouldQuit(input string) bool {
	quitWords := []string{"quit", "exit", "bye", "goodbye"}
	input = strings.ToLower(strings.TrimSpace(input))
	
	for _, word := range quitWords {
		if input == word {
			return true
		}
	}
	return false
}

func main() {
	// Create a new chatbot instance
	bot := ChatBot{
		name:    "GoBot",
		version: "1.0",
	}
	
	// Bot introduces itself
	bot.introduce()
	
	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)
	
	// Main chat loop
	for {
		// Prompt for user input
		fmt.Print("You: ")
		
		// Read user input
		if !scanner.Scan() {
			break // Exit if there's an error reading input
		}
		
		userInput := scanner.Text()
		
		// Check if user wants to quit
		if shouldQuit(userInput) {
			fmt.Println("🤖 Goodbye! Thanks for chatting with me!")
			break
		}
		
		// Get bot response and print it
		response := bot.respond(userInput)
		fmt.Println(response)
		fmt.Println() // Add blank line for readability
	}
	
	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}