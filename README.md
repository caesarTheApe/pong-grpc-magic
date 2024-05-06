# Pong gRPC Ethereum Game

## Overview

This project implements a classic Pong game using gRPC for client-server communication and an Ethereum blockchain contract for secure player betting. Players can bet Ethereum cryptocurrency against each other, with the winner of the game claiming the pot.

## Prerequisites

Go 1.20+
Node.js 20.x+
Truffle Suite
Ganache (for local blockchain simulation)
Solidity 0.8.2

## Installation

Clone the repository and install the necessary dependencies.

```
git clone https://github.com/vctt94/pong-grpc-eth.git
cd pong-grpc-eth
```

## Setting up the Solidity contract

Navigate to the Solidity contract directory and install dependencies:

```
cd pong-contract
npm install
```

Compile the Solidity contracts:

```
truffle compile
```

## Deploying the contract

Start Ganache and deploy the contracts to your local blockchain:

```
truffle migrate --reset
```

## Generating Go contract bindings

Generate Go bindings for the Solidity contract using the provided script:

```
cd contract-go
./generate-contract-bindings.sh ../pong-contract/build/contracts/PongGame.json PongGame
```

## Setting up the gRPC server and client

Navigate to the gRPC directory and regenerate the gRPC client code:

```
cd pongrpc
./regen-clientrpc.sh
```

## Running the Game

Start the gRPC server

```
cd server
go build && ./server
```

Start a gRPC go client

Open another terminal and run:

```
cd client
go build && client
```

## Setting up the UI client

Navigate to the UI client directory and install Node.js dependencies:

```
cd ../ui-client
npm install
```

Start the Envoy proxy to route gRPC-web calls:

```
envoy -c envoy.yaml
```

Start the UI client:

```
npm run start
```

## Game Rules
Players login using browser wallet, bet by sending Ethereum to the contract before starting a game.

The game is played over gRPC.

The winner is declared in-game, and the contract disburses the pot to the winner's address.

## Contributing
Feel free to fork the project and submit pull requests. You can also open issues if you encounter any problems or have suggestions for improvements.
