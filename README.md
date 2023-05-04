# ☕ Moxie Chat App

Moxie is a chat app that uses a microservices architecture to provide scalability and modularity. This project is built with Golang and gRPC, and uses Protocol Buffers to define the API contract between microservices.

## ⚙️ Prerequisites

Before you can run Moxie, you will need the following tools installed on your system:

- [Golang](https://go.dev/) (version 1.16 or later)
- [Protocol Buffers](https://github.com/protocolbuffers/protobuf-go) (protoc binaries)
- [Docker](https://www.docker.com/)
- Windows System (anything else is untested)

## 🚀 Getting Started

To get started with Moxie, follow these steps:

1. Clone the Moxie repository to your local machine and navigate to the cloned repository:

   ```sh
   git clone https://github.com/xd-Abi/moxie
   cd moxie
   ```

2. Install the required dependencies by running the following command:

   ```sh
   go get
   ```

3. Create all needed docker containers using the following command:

   ```sh
   docker-compose up -d
   ```

4. Create all needed docker containers using:

   ```sh
   docker-compose up -d
   ```

5. Navigate to the [scripts](/scripts/) folder and execute [setup.bat](/scripts/setup.bat) to generate all the necessary gRPC protocol buffers.
6. Copy the [.env.dev](/.env.dev) file and rename it to .env. Modify the SendGrid API key and sender according to your SendGrid API configuration. This step can be skipped, although some microservices may thrown an exception due to invalid api keys.

7. Run all microservices one by one using:

   ```sh
    go run apis/<api-name>/main.go

    # Example
    go run apis/auth/main.go
   ```

## 🔑 License

This project is licensed under the terms of the [MIT license](/LICENSE).
