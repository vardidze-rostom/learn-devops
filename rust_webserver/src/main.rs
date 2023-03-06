use std::io::{Write};
use std::net::TcpListener;

fn main() {
    let listener = TcpListener::bind("0.0.0.0:8080").unwrap();

    for stream in listener.incoming() {
        let mut stream = stream.unwrap();

        let response = "HTTP/1.1 200 OK\r\n\r\nHello World! v1";
        stream.write(response.as_bytes()).unwrap();
        stream.flush().unwrap();
    }
}
