use std::net::{TcpListener, TcpStream};
use std::io::{Read, Write};

pub struct LoadBalancer {
    port: u16, 
    backend_url: String,
}

impl LoadBalancer {
    pub fn new(port: u16, backend_url: String) -> Self {
        LoadBalancer { port, backend_url } 
    }

    pub fn start(&self) -> Result<(), std::io::Error> {
        let listener = TcpListener::bind(format!("127.0.0.1:{}", self.port))?;
        println!("Load balancer listening on port {}", self.port);

        for stream in listener.incoming() {
            let stream = stream?;
            self.handle_connection(stream);
        }

        Ok(())
    }

    fn handle_connection(&self, mut client_stream: TcpStream) {
        let mut buffer = [0; 1024];
        client_stream.read(&mut buffer).unwrap();

        println!("Received request from {}", client_stream.peer_addr().unwrap());
        println!("{}", String::from_utf8_lossy(&buffer[..]));

        // Forward request to backend server
        match TcpStream::connect(&self.backend_url) {
            Ok(mut backend_stream) => {
                if let Err(e) = backend_stream.write_all(&buffer) {
                    eprintln!("Failed to write to backend: {}", e);
                    return;
                }

                // Read response from backend server
                let mut response = Vec::new();
                if let Err(e) = backend_stream.read_to_end(&mut response) {
                    eprintln!("Failed to read from backend: {}", e);
                    return;
                }

                println!("Received response form backend: {}", String::from_utf8_lossy(&response));

                if let Err(e) = client_stream.write_all(&response) {
                    eprintln!("Failed to write response to client: {}", e);
                }
                if let Err(e) = client_stream.flush() {
                    eprintln!("Failed to flush client stream: {}", e);
                }
            }
            Err(e) => {
                eprintln!("Failed to connect to backend: {}", e);
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::net::TcpStream;
    use std::io::{Read, Write};
    use std::thread;
    use std::time::Duration;
    use std::process::Command;
    use std::env;

    fn get_available_port() -> u16 {
        TcpListener::bind("127.0.0.1:0").unwrap().local_addr().unwrap().port()
    }

    fn start_backend_server(port: u16) -> std::process::Child {
        let current_dir = env::current_dir().unwrap();
        let parent_dir = current_dir.parent().unwrap();
        let python_script_path = parent_dir.join("test_backend_server.py");

        println!("attempting to start backend server with script: {:?}", python_script_path);

        let child = Command::new("python3")
            .arg(&python_script_path)
            .arg(port.to_string())
            .spawn()
            .expect("Failed to start backend server");

        thread::sleep(Duration::from_secs(1));

        child
    }


    #[test]
    fn test_load_balancer_accepts_connection() {
        let backend_port = get_available_port();
        let mut backend_server = start_backend_server(backend_port);

        let lb_port = get_available_port();
        let lb = LoadBalancer::new(lb_port, format!("127.0.0.1:{}", backend_port));

        let lb_thread = thread::spawn(move || {
            lb.start().unwrap();
        });

        // Give some time for the server to start 
        thread::sleep(std::time::Duration::from_secs(2));

        let mut stream = TcpStream::connect(format!("127.0.0.1:{}", lb_port)).unwrap();
        stream.write_all(b"GET / HTTP/1.1\r\nHost: localhost\r\n\r\n").unwrap();

        let mut response = String::new();
        stream.set_read_timeout(Some(Duration::from_secs(5))).unwrap();
        match stream.read_to_string(&mut response) {
            Ok(_) => {
                println!("Received response: {}", response);
                assert!(response.contains("HTTP/1.0 200 OK"));
                let response_content = format!("Hello from backend server on port {}", backend_port);
                assert!(response.contains(&response_content));
            },
            Err(e) => {
                panic!("Failed to read response: {}", e);
            }
        }

        backend_server.kill().expect("Failed to kill backend server");
    }
}
