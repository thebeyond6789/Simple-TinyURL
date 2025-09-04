# TinyURL Clone - Application Running Guide

## Project Structure

```
tinyurl-go/
├── backend/
│   ├── handlers/
│   │   └── url_handlers.go
│   ├── models/
│   │   └── url_mapping.go
│   ├── utils/
│   │   └── slug.go
│   ├── main.go
│   ├── go.mod
│   └── go.sum
└── frontend/
    ├── index.html
    ├── css/
    │   └── style.css
    └── js/
        └── script.js
```

## How to Run the Application

### System Requirements

- Go 1.16 or higher
- Internet connection to download dependencies

### Step 1: Install Dependencies

```bash
cd backend
go mod tidy
```

### Step 2: Run the Application

```bash
go run main.go
```

The application will run on `http://localhost:8080`.

### Step 3: Access the Application

Open your browser and go to `http://localhost:8080`.

## How to Use

1. Enter the URL you want to shorten in the input field.
2. Click the "Shorten" button.
3. The application will return a shortened URL.
4. When accessing the shortened URL, you will be redirected to the original URL.

## Notes

During development, I encountered some issues with the WSL environment when testing simple servers (Go, Node.js, Python) that couldn't be accessed from localhost. Possible reasons include:

1. **Firewall**: Check if the firewall is blocking port 8080.
2. **WSL Network Configuration**: Network configuration for WSL might be needed.
3. **Port Conflict**: Check if another application is using port 8080 with the command `netstat -tulpn | grep :8080`.

If you encounter similar issues, try:

- Changing the port in the `main.go` file (replace `:8080` with `:8081` or another port).
- Running the application with root privileges: `sudo go run main.go`.
- Checking the firewall and network configuration.

## Application Architecture

### Backend (Go)

- Uses Gorilla Mux for routing.
- Generates random slugs for shortened URLs.
- Stores the mapping between slugs and original URLs in memory (using a map).
- API endpoints:
  - `POST /api/shorten` to shorten a URL.
  - `GET /{slug}` to redirect to the original URL.

### Frontend (HTML/CSS/JS)

- Simple interface with an input field for entering URLs and a button to shorten them.
- Uses Fetch API to call backend APIs.
- Displays the shortened URL upon success.
- Displays errors when they occur.

## Features

- Shortens long URLs into short URLs.
- Redirects from short URLs to original URLs.
- Simple and intuitive user interface.

## Limitations

- Data is stored in memory and will be lost when the server restarts.
- No authentication or security.
- No rate limiting.
- No duplicate slug checking.

## Potential Improvements

- Use a database (SQLite, PostgreSQL, MongoDB) for persistent data storage.
- Add authentication and security.
- Add rate limiting to prevent abuse.
- Add duplicate slug checking and generate a new slug if needed.
- Add access statistics tracking.
- Add a URL management interface.
