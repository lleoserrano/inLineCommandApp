# inLineCommandApp

## Overview
inLineCommandApp is a Command Line Interface (CLI) application to search servers and IPs' information built with Go.

## Features
- Search servers based on url.
- Ping IP addresses based on url.
- Display servers or IPs.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)

## Installation
To install the application, download it and run go build or run on terminal.

## Usage
Run this command from on terminal.

### Servers
```go
go run main.go server --host amazon.com.br
```
##### Output
```
In Line Command App
ns2.amzndns.co.uk.
ns2.amzndns.com.
ns2.amzndns.org.
ns2.amzndns.net.
ns1.amzndns.net.
ns1.amzndns.co.uk.
ns1.amzndns.com.
ns1.amzndns.org
```

### Ips
```go
go run main.go ip --host amazon.com.br
```

##### Output
```
In Line Command App
54.239.26.87
52.94.225.243
72.21.203.171
```

