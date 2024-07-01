# Chat Application

This is a simple chat application implemented in Go, using WebSockets for real-time communication and Redis for message storage. The application is dockerized using Docker Compose to facilitate easy setup and deployment.

## Features (or Requirements)

1. **Create CLI (command-line interface) for chat application:**  

    &#9745; The client side of the application is built as a command-line interface (CLI) using Go. This allows users to interact with the chat application through a terminal.

2. **Implement WebSocket connections for real-time communication between client and server:**

    &#9745; The application uses the Gorilla WebSocket package to establish WebSocket connections between the client and server

3. **Function to send messages to all users in a chat room:**

    &#9745; Clients can send messages to all users in a specific chat room

4. **DM (Direct message):**

    &#9745; Clients can send direct messages to specific users.

5. **User authentication:**

    &#9745; The application uses simple password-based authentication for chat rooms. Users must provide a password to create, join, or delete a room, ensuring that only authorized users can access or modify a room.

6. **Allow users to join and leave chat rooms:**

    &#9745; Users can join and leave chat rooms by sending commands via the CLI and provide password.

7. **Allow users to create chat rooms:**

    &#9745; Users can create chat rooms with a specified password using the CLI. 

8. **Store messages (Database, File, etc.):**

    &#9745; Messages are stored in Redis, which acts as the message store for the application.

9. **Implement error handling for scenarios such as connection error:**

    &#9745; The application includes error handling for various scenarios, including connection errors, incorrect passwords, and invalid commands.

10. **Docker/Docker-compose for deployment:**

    &#9745; The entire application, including the client, server, and Redis, is dockerized.

## Prerequisites

- Docker
- Docker Compose

## Setup

1. **Unzip the file and go to working directory:**
    ```sh
    cd chat-app
    ```

2. **Build and start the containers:**
    ```sh
    docker compose build
    docker compose up -d ## Run the server in background
    ```

3. **Start a new client:**

    To start a new client instance, run:
    ```sh
    docker compose run client
    ```

## Usage

1. **Create a room:**
    ```sh
    create [room] [password]
    ## example:
    ## create general pass
    ```

2. **Join a room:**
    ```sh
    join [room] [password]
    ## example:
    ## join general pass
    ```

3. **Send a message:**
    ```sh
    send [room] [message]
    ## example:
    ## send general Hello World!
    ```

4. **DM a user:**
    ```sh
    dm [user] [message]
    ## example:
    ## dm user1 Hello User1
    ```

5. **Leave a room:**
    ```sh
    leave [room]
    ## example:
    ## leave general
    ```

6. **Delete a room:**
    ```sh
    delete [room] [password]
    ## example:
    ## delete general pass
    ```

7. **Exit:**
    ```sh
    exit
    ```
