# JSON Validation Service

Welcome to the JSON Validation Service project! This web service is designed to validate JSON data against a predefined structure using the Go programming language and the Gin web framework. Validate JSON data structure against a predefined `Book` struct and receive immediate feedback on whether the JSON structure matches the expected format.

## Getting Started

1. **Clone the repository:** `git clone https://github.com/nakulbh/gp-json-validation.git`, then `cd go-json-validation`
2. **Install dependencies:** `go get -u github.com/gin-gonic/gin`
3. **Run the application:** `go run main.go`
4. **Access the validation endpoint:** Open your web browser or use an API tool to send POST requests to `http://localhost:9000/validate-book`.

## How It Works

The `ValidateBook` endpoint checks the provided JSON data against the structure of a predefined `Book` struct. It ensures that the JSON keys match the expected fields of the `Book` struct.

## Usage

1. **Send a POST request:** Send a POST request to `http://localhost:9000/validate-book` with your JSON data in the request body.
2. **Immediate response:** Receive an immediate response indicating whether the JSON structure matches the `Book` struct.

## Limitations

This project is intended to showcase JSON validation against a struct using the Gin framework. It's not intended for production use without further enhancements.

## Contributions

Contributions to enhance this project are highly welcome! If you encounter a bug or have an idea for improvement, feel free to open an issue or submit a pull request.


## Contact

For questions or inquiries, please contact [nakulbhardwaj37@gmail.com.com]. Enjoy validating JSON structures with ease!
