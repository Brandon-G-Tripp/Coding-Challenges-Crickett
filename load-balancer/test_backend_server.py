#!/usr/bin/env python3

from http.server import HTTPServer, SimpleHTTPRequestHandler
import sys

class TestHandler(SimpleHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/plain')
        self.end_headers()
        response = f"Hello from backend server on port {self.server.server_port}"
        self.wfile.write(response.encode())


def run_server(port):
    print("Entering run_server function")
    server_address = ('', port)
    print(f"Server address: {server_address}")
    httpd = HTTPServer(server_address, TestHandler)
    print(f"HTTPServer created")
    print(f"Starting test backend server on port {port}")
    httpd.serve_forever()


if __name__ == "__main__":
    try:
        if len(sys.argv) > 1:
            port = int(sys.argv[1])
        else:
            port = 8000
        print(f"Attempting to start server on port {port}")
        run_server(port)
    except Exception as e: 
        print(f"An error occurred: {e}")
