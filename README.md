# http-echo-headers

A very simple webserver in Golang that prints all headers reaching it.

The code is mostly copy+paste from these repositories:
- [hashicorp/http-echo](https://github.com/hashicorp/http-echo)
- [bukowa/http-headers](https://github.com/bukowa/http-headers)

## Run

### Using Docker
The simplest way to run would be to use docker.

#### Docker Build
To build docker container image, from the root of this project, simply run:
```shell
docker build -t local/http-echo-headers .
```

#### Docker Run
Once the docker container image is build, simply run:
```shell
docker run --rm --name http-echo-headers -p 8080:8080 local/http-echo-headers:latest
```

### Directly from command line
To run this webserver, pre-requisite is to have go `1.23+` installed. Once that's met, from root of the project, simply run:

```shell
go run main.go
```

To build this webserver, run
```shell
go build .
```

This will generate `http-echo-headers` binary in current directory. Then to run the webserver, run:
```shell
./http-echo-headers
```

To check help, run:
```shell
$ ./http-echo-headers --help
Usage of ./http-echo-headers:
  -host string
    	interface to listen on (default "0.0.0.0")
  -port string
    	port to listen on (default "8080")
```

## Check Run
If the run is successful, the output should look something like this:
```text
2024/10/10 12:02:17 [INFO] server is listening on 0.0.0.0:8080

```

## Test
To test, run this from command line:
```shell
curl -H "Custom-Header: SomeValue" http://localhost:8080/
```

the output should look like this:
```text
{"Accept":"*/*","Custom-Header":"SomeValue","User-Agent":"curl/8.7.1"}
```
