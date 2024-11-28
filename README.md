# Go Web Server Project

This project is a simple web server written in Go. It serves static files and has handlers for a basic HTML form and a "Hello" route.

## Features
- Serve static files from a `static` directory.
- Handle form submissions with `POST` requests to the `/form` endpoint.
- Respond to `GET` requests at the `/hello` endpoint.

## Prerequisites

Make sure you have the following installed:
- [Go](https://golang.org/dl/) (1.18 or later)
- A web browser for testing

## Project Structure
```
Project_Server_Golang/
|-- static/
|   |-- index.html
|   `-- form.html
|-- main.go
|-- README.md
```

## Usage

### Step 1: Clone the Repository
```bash
git clone [https://github.com/gaurikarkhile001/Project_Server_Golang.git](https://github.com/gaurikarkhile001/Project_Server_Golang.git)
cd Project_Server_Golang
```

### Step 2: Run the Server
To start the server, run:
```bash
go run main.go
```

The server will start on `http://localhost:8080` by default.

### Step 3: Access the Web Pages
1. Open your browser.
2. Visit `http://localhost:8080/index.html` to see the static home page.
3. Visit `http://localhost:8080/form.html` to submit a form.
4. Visit `http://localhost:8080/hello` for a simple "Hello" message.

## Endpoints

### `/`
Serves static files from the `static` directory. For example:
- `http://localhost:8080/index.html`
- `http://localhost:8080/form.html`

### `/form`
Handles `POST` requests for form submissions. Example:
- Submit a form with `name` and `address` fields.
- The server responds with the submitted data.

### `/hello`
Handles `GET` requests to display a simple greeting message.



## Contribution
Feel free to fork this repository and submit pull requests for improvements or additional features.

