# Go REST Template 🚀

A reusable Go REST API template with Chi, modular structure & best practices!

---

## Features ✨

- **Chi Router**: Lightweight and efficient HTTP routing.
- **MongoDB Integration**: Easily configurable MongoDB database support.
- **Concurrency Support**: Goroutines for optimal performance.
- **Modular Structure**: CMD architecture for scalable projects.
- **Task Management Example**: CRUD implementation to demonstrate usage.
- **Middleware**: Logging, error recovery, and more.
- **Customizable**: Designed to be a flexible starting point for your needs.

---

## Getting Started 🏁

### Prerequisites

- Go 1.23.3 or later
- MongoDB instance (local or cloud)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yigit433/go-rest-template.git
   cd go-rest-template
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up your MongoDB connection string in the environment variables or configuration file.

4. Run the server:

   ```bash
   go run ./cmd/app/main.go
   ```

5. The server will be available at `http://localhost:8080`.

---

## Project Structure 🗂️

```
.
├── cmd
│   └── app
│       └── main.go                       # Entry point of the application, starts the app
│
├── internal
│   ├── app                               # Core application logic and features
│   │   ├── routes                        # HTTP route definitions and routing setup
│   │   │   └── tasks                     # Tasks-specific router
│   │   │       ├── handler.go            # HTTP handlers for task-related requests
│   │   │       └── setup.go              # Route setup and middleware binding for tasks
│   │   │
│   │   ├── repository                    # Data access layer (interfacing with databases)
│   │   │   └── tasks_repository.go       # Data operations related to task
│   │   │
│   │   └── model                         # Data models used across the application
│   │       └── tasks_model.go            # Definition of the Task data model
│   │     
│   ├── configs                           # Configuration management (e.g., environment variables, settings)
│   │
│   └── database                          # Database-related logic and connection management
│       └── mongodb.go                    # MongoDB setup and configuration details
│                          
├── go.mod                                # Go module definition, dependency tracking
└── README.md                             # Project documentation (setup, usage, etc.)
```

---

## API Endpoints 📡

### Task Management

| Method | Endpoint      | Description              |
|--------|---------------|--------------------------|
| GET    | `/tasks`      | Retrieve all tasks       |
| POST   | `/tasks`      | Create a new task        |
| PUT    | `/tasks/{id}` | Update an existing task  |
| DELETE | `/tasks/{id}` | Delete a task            |

---

## Logging 📝

All HTTP requests and responses are logged to the console for debugging and monitoring.

Example log output:

```
2025/01/23 12:34:56 Started GET /tasks from 127.0.0.1
2025/01/23 12:34:56 Completed in 5ms
```

---

## Contributing 🤝

Contributions are welcome! Feel free to open an issue or submit a pull request.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch-name`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch-name`).
5. Open a pull request.

---

## License 📝

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgements 🙏

Thanks to the open-source community for inspiration and support!

---

Happy coding! 💻
