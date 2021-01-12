# MC
MC - short for Memory Cache - is a simple distributed caching server for learning purpose.

## Features

- In memory LRU cache
- RPC transport
- Cluster
- Command line client
- Go client

So far, MC has only `get`, `add` and `remove` of caching and a Consistent Hash Algorithm will be considered to use for the scalability in the future.

## Usage

### Install

```bash
go get github.com/mivinci/mc
```

Or (not available yet)

```bash
docker pull mivinci/mc
```

Or [download](https://github.com/Mivinci/mc/tags) the pre-build version for macOS or Windows.

### Run

`mc` can be run either a server or a client.

#### Run an MC server.

```
mc -server
```

By default, MC chooses `127.0.0.1:8000` to serve on, or you can decide the address by providing a `-address` parameter.

```
mc -server -address 127.0.0.1:8001
```

for short

```
mc -server -address :8001
```

You can customize the maximum caching size for a single MC server.

```
mc -server -c 128
```

#### Run an MC client.

```
mc -client
```

Also, the MC client chooses `127.0.0.1:8000` to dial during usage. You can provide an `-address` parameter to the MC client to communicate with a different address.

#### Use Cluster

Once you have run multiple MC servers, you can communicate with them simply by using the MC client with multiple addresses seperated by commas.

```
mc -client -address 1.1.1.1:8000,1.1.1.2:8001,1.1.1.3:8002
```

The MC command line  client will lead you to an interface like this.

```
$ mc -client
MC v1.0.0, press [ctrl c] to exit.
127.0.0.1:8000> get name
XJJ
127.0.0.1:8000> remove name
127.0.0.1:8000> get name
nil
```

### Go Client

```go
import "github.com/mivinci/mc"
```

#### Add

```go
r := mc.Add("name", []byte("XJJ"))
```

#### Get

```go
r := mc.Get("name")
```

#### Remove

```go
r := mc.Remove("name")
```

#### Cluster

```go
client := mc.NewClient("1.1.1.1:8000", "1.1.1.2:8001", "1.1.1.3:8000")
r := client.Get("name")
...
```

## About

Thanks to the projects and videos below for the inspiration.

- [https://github.com/asim/mq](https://github.com/asim/mq)

- [https://www.youtube.com/watch?v=iuqZvajTOyA](https://www.youtube.com/watch?v=iuqZvajTOyA)

