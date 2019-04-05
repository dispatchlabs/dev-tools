
<img src="https://www.dispatchlabs.io/wp-content/uploads/2018/12/Dispatch_Logo.png" width="250">
 
![Go Version 1.10.3](http://b.repl.ca/v1/Go_Version-1.10.3-brightgreen.png)

# Welcome to the Dispatch Protocol dev-tools!

This is the toolset developed for working on the [disgo](https://github.com/dispatchlabs/disgo) implementation of the Dispatch Protocol. This is the place to start if you're looking to run your own Dispatch network, or contribute to the open-source development of the protocol.

### Questions? Opinions? Talk to us:
If you have any questions or just want to get to know us better, come say hi [in our discord](https://Dispatchlabs.io/discord) (https://Dispatchlabs.io/discord)


## Table of Contents
 * [Getting Started](#prerequisites)
    + [Prerequisites](#prerequisites)
    + [Download](#download)
  * [Local Cluster](#local-cluster)
      + [newLocalCluster](#-newlocalcluster)
      + [reset](#-reset)
      + [Configuring your local cluster](#-configuring-your-local-cluster)
  * [Sending Transactions](#sending-transactions)
    + [Pick a seed](#-pick-a-seed)
    + [transfer](#-transfer)
    + [deployContractFromFile](#-deploycontractfromfile)
    + [executeWrite](#-executewrite)
    + [executeRead](#-executeread)
  * [Contributing!](#contributing)

## Getting Started

### Prerequisites:

#### Golang

Since both disgo and the dev-tools are written in go, it'll help if you [installed golang](https://golang.org/doc/install#install) on your machine

#### Branches

We try to keep our branch names connected between the disgo <--> dev-tools repos. If you're having a hard time building one or the other, make sure they're both on the same branch (dev and dev, or master and master).

### Download:
Go get the dev-tools
```bat
go get github.com/dispatchlabs/dev-tools
```
Then go get the dependencies:
```bat
cd ~/go/src/github.com/dispatchlabs/dev-tools
go get ./...
```

## Local Cluster

To do any development on the Dispatch protocol, it really helps to be able to run your own local cluster of Seeds and Delegates. 

### 🆕 newLocalCluster 
```bat
go run main.go newLocalCluster [numberOfDelegates (default=4)]
```
The **newLocalCluster** command will build Seed and Delegate nodes in your `~/disgo_cluster` directory. Then go to these directories and start up the multiple nodes in multiple terminal windows, starting with the seed:
```bat
cd ~/disgo_cluster/seed-0
./disgo
```
Then from another teminal window:
```bat
cd ~/disgo_cluster/delegate-0
./disgo
```
Rinse and repeat for as many Delegates as you'd like to run. 

### 🔥 reset
```bat
go run main.go reset
```

The **reset** command will re-build the disgo binaries and empties the `db` folder in `~/disgo_cluster` directories. This is the command you usually run when you've edited the disgo code and want to "re-compile" the cluster. 

*note: Reset will not empty or update the `config` folder.

### 🔧 Configuring your local cluster:
|        |grpc Port          |http Port        |local Api         |
|--------|-------------------|-----------------|------------------|
|Seed	      |localhost:1973    |  Ø  | localhost:1975
|Delegate 0   |localhost:3502  |localhost:3503     | localhost:3504|
|Delegate 1   |localhost:3505  |localhost:3506     | localhost:3507|
|Delegate 2   |localhost:3508  |localhost:3509     | localhost:3510|
|Delegate 3   |localhost:3511  |localhost:3512     | localhost:3513|

After creating your local cluster, each "node" in `~/disgo_cluster/` will come with a `config/config.json` file that specifies what ports that node is using. Here are the **default ports** for the nodes^^^ 

You can **test** the seed by going to [localhost:1975/v1/delegates](localhost:1975/v1/delegates), and you can test the rest of the delegates with their http API ([localhost:3503/v1/transactions](localhost:3503/v1/transactions)). For more documentation on the http API, check out [api.dispatchlabs.io](api.dispatchlabs.io).

The **genesis account** that contains all of the tokens at the instantiation of the network can be found in `dev-tools/transactions/constants.go`

## Sending Transactions

The dev-tools repo comes with tools that can send any of the 4 transaction types to any network: 

|Transaction  Type | What it does |
|--------|----------------------|
|0 |Divvy Token transfer |
|1 |Deploy Smart-Contract |
|2 |Execute Write |
|3 |Execute Read |

### 🌰 Pick a seed
```bat
go run main.go -seed=devseed.dispatchlabs.io:1975 transfer [to-address] [amount]
```
The dev-tools can send to **any network** with an active seed, but it will default to sending transactions to the local cluster. Just put the `-seed=*ip*:*port*` flag in front of the function name you want to call. The official Dispatch mainnet seed domain is static at seed.dispatchlabs.io and the devnet seed domain is devseed.dispatchlabs.io 

### 💸 transfer
```bat
go run main.go transfer [toAddress] [amount]
```
You can **transfer** tokens out of the genesis account defined in `dev-tools/transactions/constants.go`, you you could transfer tokens from another address with this command:
```bat
go run main.go transfer [fromPrivKey] [fromAddress] [toAddress] [amount]
```

### 📑 deployContractFromFile
```bat
go run main.go deployContractFromFile [contractName]
```
This command will search the `dev-tools/test_contracts/` directory for a compiled [contractName].abi and [contractName].bin  and  **deploy** it from the genesis account.

### 🖋 executeWrite
```bat
go run main.go executeWrite [smartContractAddress] [function] [parameter0Type]:[parameter0Value] [parameter1Type]:[parameter1Value]...
```
**executeWrite** will execute a function on a deployed smart-contract, and write the result to the ledger. Functions can have any number of parameters that are submitted as the parameter type and value separated by a colon (:). ***note**: The "address" parameter type in some smart-contracts is submitted as a "string" parameterType.

Because executeWrite is needs to be gossipped and costs bandwidth, it is recommended to use the executeRead function whenever possible.

### 📚 executeRead
```bat
go run main.go executeRead [smartContractAddress] [function] [parameter0Type]:[parameter0Value] [parameter1Type]:[parameter1Value]...
```
**executeRead** will return the result of a function called on a deployed smart-contract, without writing the transaction to the ledger. executeRead costs no bandwidth and returns a much faster result than executeWrite.

# Contributing!

We need your help. We are a small team, building an open-source technology that we really believe can give people sovereignty over their data and the value it creates. 

If you see something in the code that could use some improvement, or think of a feature that would add value. We'd love it if you made a Pull Request against the [dev branch of the disgo repo](https://github.com/dispatchlabs/disgo/tree/dev) 🙏.

Thank you so much for your support ❣️
-Zane W.
