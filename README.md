# Go Chatbot - Simple Hello World Project

A beginner-friendly command-line chatbot built with Go to demonstrate basic Go programming concepts and syntax.

## 🤖 About This Project

This is a simple interactive chatbot that runs in your terminal. It's designed as a learning project for developers new to Go (Golang), showcasing fundamental concepts like:

- Package structure and imports
- Structs and methods
- String manipulation
- User input handling
- Control flow (switch statements, loops)
- Basic error handling

## ✨ Features

- **Interactive Conversations**: Responds to various user inputs
- **Time & Date**: Can tell you the current time and date
- **Contextual Responses**: Recognizes greetings, questions about itself, and Go-related topics
- **Graceful Exit**: Type 'quit', 'exit', 'bye', or 'goodbye' to end the conversation
- **Help Command**: Type 'help' to see available commands
- **Error Handling**: Handles empty inputs and unexpected errors

## 🛠️ Prerequisites

- **Go 1.19+** installed on your system
- **Text editor** (VS Code recommended) 
- **Terminal/Command line** access

### For Ubuntu 24.04.2 Users:

```bash
# Install Go
sudo apt update
sudo rm -rf /usr/local/go
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verify installation
go version
```

## 🚀 Installation & Setup

1. **Clone or create the project directory:**
   ```bash
   mkdir go-chatbot
   cd go-chatbot
   ```

2. **Initialize Go module:**
   ```bash
   go mod init chatbot
   ```

3. **Create the main.go file** and copy the source code

4. **Run the chatbot:**
   ```bash
   go run main.go
   ```

## 📁 Project Structure

```
go-chatbot/
├── README.md
├── go.mod
├── go.sum         # (auto-generated)
└── main.go        # Main application file
```

## 🎮 Usage

1. **Start the chatbot:**
   ```bash
   go run main.go
   ```

2. **Interact with the bot:**
   ```
   🤖 Hello! I'm GoBot v1.0
   🤖 I'm a simple chatbot built with Go!
   🤖 Type 'quit' or 'exit' to end our conversation.
   🤖 Let's chat!
   --------------------------------------------------
   You: hello
   🤖 Hello there! Nice to meet you!

   You: what's your name
   🤖 I'm GoBot, your friendly Go chatbot!

   You: time
   🤖 The current time is: 14:30:25
   ```

3. **End the conversation:**
   ```
   You: quit
   🤖 Goodbye! Thanks for chatting with me!
   ```

## 💬 Supported Commands

| Input | Response |
|-------|----------|
| `hello`, `hi` | Greeting response |
| `how are you` | Status check response |
| `what's your name`, `who are you` | Bot introduction |
| `time` | Current time |
| `date` | Current date |
| `go`, `golang` | Information about Go language |
| `help` | List of available commands |
| `quit`, `exit`, `bye`, `goodbye` | Exit the program |

## 🎓 Learning Objectives

This project teaches the following Go concepts:

### 1. **Package Structure**
```go
package main  // Entry point package
```

### 2. **Imports**
```go
import (
    "bufio"    // Buffered I/O
    "fmt"      // Formatting
    "os"       // Operating system interface
    "strings"  // String manipulation
    "time"     // Time functions
)
```

### 3. **Structs & Methods**
```go
type ChatBot struct {
    name    string
    version string
}

func (bot *ChatBot) introduce() {
    // Method implementation
}
```

### 4. **Control Flow**
```go
// Switch statements
switch {
case strings.Contains(input, "hello"):
    return "Hello there!"
default:
    return "Default response"
}

// For loops
for {
    // Infinite loop with break conditions
}
```

### 5. **Error Handling**
```go
if err := scanner.Err(); err != nil {
    fmt.Printf("Error reading input: %v\n", err)
}
```

## 🔧 Development Setup (VS Code + Ubuntu)

### Install VS Code:
```bash
sudo snap install --classic code
```

### Install Go Extension:
1. Open VS Code
2. Press `Ctrl+Shift+X`
3. Search for "Go" 
4. Install the official Go extension by Google

### Install Go Tools:
1. Press `Ctrl+Shift+P`
2. Type "Go: Install/Update Tools"
3. Install all recommended tools

## 🚀 Building and Running

### Development Mode:
```bash
go run main.go
```

### Build Executable:
```bash
# Build for current platform
go build -o chatbot

# Run the executable
./chatbot
```

### Cross-Platform Build:
```bash
# Build for Windows
GOOS=windows GOARCH=amd64 go build -o chatbot.exe

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o chatbot-mac

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o chatbot-linux
```

## 🎨 Customization Ideas

Extend this chatbot by adding:

- **More conversation topics** and responses
- **Persistent memory** using files or databases
- **Web interface** using `net/http` package
- **Natural Language Processing** with external APIs
- **Configuration files** for responses
- **Logging system** for conversation history
- **Unit tests** for bot responses
- **Command-line flags** for different modes

## 🔍 Code Walkthrough

### Main Components:

1. **ChatBot Struct**: Holds bot properties (name, version)
2. **introduce()**: Bot's greeting method
3. **respond()**: Core logic for processing user input
4. **shouldQuit()**: Checks for exit commands  
5. **main()**: Program entry point with chat loop

### Key Go Features Demonstrated:

- **Struct methods** with receivers
- **String manipulation** with `strings` package
- **User input** with `bufio.Scanner`
- **Time formatting** with `time` package
- **Switch statements** for pattern matching
- **Slices** for storing multiple responses

## 📚 Next Steps

After mastering this project, consider learning:

1. **Web Development**: Build a web-based chatbot with `net/http`
2. **Concurrency**: Add goroutines for handling multiple users
3. **Databases**: Store conversations in SQLite or PostgreSQL
4. **APIs**: Integrate with external services
5. **Testing**: Write unit tests with Go's testing package
6. **Deployment**: Deploy to cloud platforms

## 🤝 Contributing

This is a learning project! Feel free to:
- Add new conversation topics
- Improve response logic
- Add new features
- Fix bugs
- Enhance documentation

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 👥 Authors

- **Your Name** - Initial work and documentation

## 🙏 Acknowledgments

- Go team for creating an amazing language
- VS Code team for excellent Go support  
- Ubuntu community for great development environment

---

**Happy Coding! 🚀**

*This project is part of learning Go programming language. Perfect for beginners who want to understand Go syntax through a practical, interactive example.*