import socket
import threading
import argparse


def receive_messages(client_socket):
    while True:
        try:
            message = client_socket.recv(1024).decode('utf-8')
            print(message)
        except ConnectionResetError:
            print("Disconnected from the server.")
            break


def main():
    parser = argparse.ArgumentParser(description="GoChat Client")
    parser.add_argument(
        "-s", "--server", default="localhost", help="Server address")
    parser.add_argument("-p", "--port", type=int,
                        default=6969, help="Port number")
    args = parser.parse_args()

    host = args.server
    port = args.port

    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client_socket.connect((host, port))

    receive_thread = threading.Thread(
        target=receive_messages, args=(client_socket,))
    receive_thread.daemon = True
    receive_thread.start()

    while True:
        message = input()
        if message == "exit":
            break
        client_socket.send(message.encode('utf-8'))


if __name__ == "__main__":
    main()
