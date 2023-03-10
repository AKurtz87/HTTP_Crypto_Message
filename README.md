# Text Message via GET Req 🔏

⚠️⚠️⚠️ ##**CHECK PATHS AND OTHER ELEMENTS IN SCRIPS AND ADJUST TO YOUR NEEDS** ⚠️⚠️⚠️

This GitHub repository provides a solution for secure text message sharing among multiple users.
When a user sends a GET request, the data is transmitted to the server via an argument in the HTTP request. The server receives the GET request and           parses the encrypted argument, which contains the message sent by the user. The server then decrypts the message using the shared encryption key before displaying it to the master user.

**It's important to note that HTTP is an unencrypted protocol, and data transmitted via HTTP can be intercepted by a third-party. Therefore, it's important to establish a secure communication channel and share the encryption key between the master user and client users before exchanging messages.**

It's also important to note that **only the master user can read the messages**, while client users can only send messages. This repository is designed for secure text message sharing but some implementations could be used for data exfiltration.

Overall, by using HTTP GET requests to transmit messages to the server, this repository provides a simple and convenient way for multiple users to securely exchange text messages. However, to ensure the security and confidentiality of messages exchanged, it's important to implement appropriate security measures, such as **sharing the encryption key through a secure channel.**

