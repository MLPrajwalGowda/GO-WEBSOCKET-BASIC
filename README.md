This Project is a Basic Level of understanding about WEBSOCKETS which is written in GO Progamming Language.

- WEBSOCKETS??

1. WebSockets are used for real-time, event-driven communication between clients and servers.

2. They are particularly useful for building applications requiring instant updates, such as real-time chat, messaging, and multiplayer games.

3. (p1.png)

- How to the Program ??
  command:- go run main.go

After running the command, you will get a line printing in terminal saying connect to server with the websocket connection.
extension

- How to Know that websocket connection is working ??

Method 1:

Please install a chrome extension named "Simple Websocket Client"
open the extension and connect to the websocket server (ws://localhost:8080/socket) and send a message through socket and you can a status received from client.

Please go through the images attached in the folder for any confusion.

1. (s1.png)
2. (<s2 (2).png>)
3. (<s3 (2).png>)

Method 2 :

1. Go to website link attached below
   https://livepersoninc.github.io/ws-test-page/

2. connect to the server ws://localhost:8080/socket

3. Send the message through it and you will recevie a client message in terminal.Please find the image find below for clarification.
4. (<s4 (2).png>)

References:-

1. https://pkg.go.dev/github.com/gorilla/websocket
   Please go through the package listed above with the link.
