# AES Encryption with Golang and Gorilla Mux

This is a simple example of an HTTP server in Go that implements AES encryption using CFB (Cipher Feedback) mode.

## Features

The server has two routes:

1. `/encrypt/{key}/{plainText}` - Receives a key and plain text as parameters and returns the encrypted text in hexadecimal format.
2. `/decrypt/{key}/{cipherText}` - Receives a key and a hexadecimal encrypted text as parameters and returns the original plain text.

## Requirements

Make sure you have Go installed on your machine before running this code.

## How to Use

1. Clone this repository to your local environment.
2. Navigate to the project directory using the terminal.
3. Run the following command to start the server:

   ```shell
   go run main.go
   ```

4. The server will run at http://localhost:3000 or at the port defined by the PORT environment variable.
5. Use an HTTP client (such as cURL or Postman) to send requests to the /encrypt and /decrypt routes as needed.

## Requests Examples

Encrypt Text
Send a GET request to `http://localhost:3000/encrypt/mykey/HelloWorld`:

```shell
curl http://localhost:3000/encrypt/mykey/HelloWorld
```

Response:
```
5f0e0a33c22786762fbde72e1c78e900
```

Decrypt Text
Send a GET request to `http://localhost:3000/decrypt/mykey/5f0e0a33c22786762fbde72e1c78e900`:

```shell
curl http://localhost:3000/decrypt/mykey/5f0e0a33c22786762fbde72e1c78e900
```

Response:
```
HelloWorld
```

## Notes
- Make sure to keep the encryption and decryption keys consistent when using the server.
- This is just a simple example for educational purposes and should not be used in production without considering additional security aspects.