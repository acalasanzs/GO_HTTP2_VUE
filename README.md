# Go + Vue HTTP v2
Fet amb carinyo per Albert Calasanz.


## How-To

## Implementation Guide for HTTP/2 Support

### 1. Setting up the Go Backend with HTTP/2

HTTP/2 requires TLS in browsers, so I've updated the Go server to use HTTPS:

1. Generate self-signed certificates for development:
```bash
cd backend
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host localhost
```

This will create `cert.pem` and `key.pem` files in your backend directory.

2. Update the server to use TLS and HTTP/2 (HTTP/2 is enabled by default when using TLS in Go's HTTP server):
   - Changed port to 8443 (standard for HTTPS)
   - Added TLS configuration
   - Added protocol logging to verify HTTP/2 usage
   - Updated CORS to allow HTTPS origins

3. Run the Go server:
```bash
go run main.go
```

### 2. Setting up the Vue Frontend for HTTP/2

1. Copy the generated certificates to the frontend folder as well:
```bash
cp backend/cert.pem frontend/
cp backend/key.pem frontend/
```

2. Create a `vite.config.js` file in the frontend directory with the code from the vite-config artifact to enable HTTPS

3. Update the Vue App component to:
   - Use HTTPS URL for API connections
   - Show the connection status and protocol version
   - Add better error handling for connection issues
   - Add timeout support for fetch requests

4. Install the necessary dependencies:
```bash
cd frontend
npm install
```

5. Run the Vue development server:
```bash
npm run dev
```

### Key HTTP/2 Benefits in This Implementation:

1. **Multiplexing**: HTTP/2 can send multiple requests and responses simultaneously over a single connection, reducing latency

2. **Binary Protocol**: More efficient than HTTP/1.1's text protocol

3. **Header Compression**: Reduces overhead with each request

4. **Server Push**: Though not implemented in this example, HTTP/2 allows servers to push resources to clients before they're requested

### Notes on Browser Access:

When you access the app at `https://localhost:5173`, you'll likely get a browser warning about the self-signed certificate. You'll need to:

1. Click "Advanced" or similar
2. Click "Proceed to localhost (unsafe)" or equivalent
3. Do the same for the backend if you directly access it at `https://localhost:8443`

### Security Note:

This implementation uses self-signed certificates for development purposes only. In a production environment, you should use proper certificates from a trusted certificate authority.

Would you like me to explain any specific aspect of the HTTP/2 implementation in more detail?