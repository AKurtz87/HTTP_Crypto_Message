# Text Message via GET Req

🧟🧟🧟   CHECK PATHS AND OTHER ELEMENTS IN SCRIPS AND ADJUST TO YOUR NEEDS   🧟🧟🧟  

This GitHub repository provides a solution for secure text message sharing among multiple users. The master user runs the server, while other users can send messages to the master user by making HTTP GET requests via the terminal.

The server receives the GET request, parses the encrypted argument, and decrypts the message before displaying it to the master user. The encryption key is shared between the master user and clients through a secure channel before data is exchanged.

Please note that only the master user can read the messages, while client users can only send messages. This repository can also be used for data exfiltration by sending GET requests to a remote server.

To use this repository, it is important to establish a secure communication channel and share the encryption key between the master user and client users before exchanging messages.
