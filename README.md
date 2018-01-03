# Spider

A simple web server using golang that will crawl for partner website and then get all current transaction/order history from user.


## Index

- [Endpoint](#endpoint)
- [Setup](#setup)
    - [Install Golang](#install-golang)
    - [Install go depedencies](#install-go-depedencies)
    - [Compile Protobuf](#compile-protobuf)
    - [Compile webserver](#compile-webserver)
    - [Run the server](#run-the-server)
    - [Test the server](#test-the-server)
- [Project Structure](#project-structure)
- [Coding Guidelines](#coding-guidelines)


## Endpoint

This project have some endpoint.

| Uri | Type | Description |
|-----|------|-------------|
| /bca | Post | Get user transaction history from [BCA](https://ibank.klikbca.com/) |
| /blibli | Post | Get user order history from [Blili](https://www.blibli.com) |
| /elevenia | Post | Get user order history from [Elevania](https://www.elevenia.co.id/) |
| /lazada | Post | Get user order history from [Lazada](https://www.lazada.co.id) |
| /tokopedia | Post | Get user orderls history from [Tokopedia](https://www.tokopedia.com) |


## Setup

To setup this service, we need to follow these step:


### Install golang

Run this command to install golang to your **mac**:

```bash
make install_golang_mac
```

the minimum requirement for this script to run is that your mac machine have:
- brew
- make


### Install go depedencies

Run this command to install all golang depedencies:

```bash
make dep
```

this will install all go depedencies needed for webserver & crawling library.


### Compile protobuf

Run this command to compile the protobuf:

```bash
make proto
```

this will compile the proto to `src/spider_data.pb`. The minimum requirement is that your machine has protobuf installed (if you install golang using make, then protobuf also installed to your machine).


### Compile webserver

For **mac**, run this command to compile the code:

```bash
make spider_server_mac
```

For **linux**, run this command to compile the code:

```bash
make spider_server
```


### Run the server

After compile the webserver, the we can run the binary.

For **mac**, run this command to run the binary:

```bash
./spider_server_mac
```

For **linux**, run this command to run the binary:

```bash
./spider_server
```

**Remember** to compile first before run the server

By default, your server will be run on `localhost:8006`


### Test the server

Once the server run on your machine, try to curl/postman to your localhost with the proper endpoint. Here is the example of the curl:

```bash
curl -H "Content-Type: application/json" -X POST -d '{ "username" : "", "password" : "" }' localhost:8006/crawl/lazada
```

## Project Structure

This project has this project structure:

```
.
├── Makefile
├── README.md
├── SpiderServer
│   ├── BCAHandler.go
│   ├── BCAHandler_test.go
│   ├── BlibliHandler.go
│   ├── BlibliHandler_test.go
│   ├── EleveniaHandler.go
│   ├── EleveniaHandler_test.go
│   ├── LazadaHandler.go
│   ├── LazadaHandler_test.go
│   ├── TokopediaHandler.go
│   ├── TokopediaHandler_test.go
│   ├── config.go
│   ├── debug.html
│   └── util.go
├── conf
│   └── spider_server.xml
├── proto
│   └── spider_data.proto
└── spider_server.go
```


## Coding Guidelines

- **Requirement**: Read [Clean Code](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship/dp/0132350882)
- **Pull Requests**: Make sure the code works. Assume the reviewer is dumb.
- **Business Logic**: If writing a fair amount of business logic that should be unit tested, you should write the code in a modular way that can easily be unit tested.
- **KISS**: always Keep It Simple Stupid
- **100x30**: 100 length maximum, with 30 line maximum (should be able to read from my Atom with 14 pt font and normal theme without scrolling)