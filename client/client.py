import socket
import threading


def receive_messages(client_socket):
    while True:
        try:
            message = client_socket.recv(1024).decode('utf-8')
            print(message)
        except ConnectionResetError:
            print("Disconnected from the server.")
            break


def main():
    host = "localhost"
    port = 6969

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
