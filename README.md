
<h1 align="center">Syscom TCP</h1>
<p align = "center">A TCP server/client</p>

[]()
<p align="center">
  <a href="#-technology">Technology</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
    <a href="#-project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-run">How to Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-license">License</a>
</p>

<p align="center">
  <img alt="License" src="https://img.shields.io/static/v1?label=license&message=MIT&color=8257E5&labelColor=000000">
</p>

## Introduction
In this challenge you must create a communication system between client and server using Go, including both the server and the client. You can use any external lib if necessary. The protocol must be based on pure TCP, do not use existing application protocols such as HTTP or WebSockets.

The following challenge can be found [here](https://app.devgym.com.br/challenges/c8ea75f7-d8b7-4306-89c2-10d155248719).

## ✨ Technology

The Project was develop as using the following techs:
- [Go](https://go.dev/)
- [UUID](https://github.com/google/uuid)
- [GoDotEnv](https://github.com/joho/godotenv)



## 💻 Project

The project was written in Go, using the standard `net` library. This enabled us to create a TCP server and client, allowing for a connection to be established. We can then send commands such as `LIST` and `RELAY`, and the server will respond accordingly.

###  📓 Requirements 
The project has the following requirements

- Create a TCP server that listens on a port specified by an environment variable PORT.
- Create a TCP client that connects to a server. When a client connects to the server, a unique id is generated and sent back to the client.
- The client can send messages with three data:
    - id - the client's id (required),
    - action - an action expected by the server (required),
    - body - any data (json, plain text, etc.) that a client wants to send to the server (optional).
- When the client sends a message with action LIST, the server should return a list of all connected client ids.
- When the client sends a message with action RELAY, the server should send the message from the body field to all connected clients.



## 🚀 How to Run
To run this project, follow the following instructions:

1. Clone the repository
2. Start the server: 
```bash
go run -race .\server\server.go
```
3. Connect the client to the server: 
```bash
go run -race .\client\client.go
```

4. Once the client is connected, you will receive an ID that is necessary to send to the server in order to receive a response. Therefore, the commands are:
   
```bash
    id LIST # to list all the clients
    id RELAY message  #to send messages to all the clients connected to the application
```

> The server can handle multiple clients. To do this, open as many Bash windows as you want and execute the command: `go run -race .\client\client.go`.


## 📄 License
The projects is under the MIT license. See the file [LICENSE](LICENSE) fore more details

---
## Author

Made with ♥ by Rafael 👋🏻


[![Linkedin Badge](https://img.shields.io/badge/-Rafael-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/tgmarinho/)](https://www.linkedin.com/in/rafael-mgr/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-red?style=flat-square&link=mailto:nelsonsantosaraujo@hotmail.com)](mailto:ribeirorafaelmatehus@gmail.com)
