# golang-requester
Golang HTTP GET requester with channels and goroutines

# Usage

```bash
$ go run main.go --help
$ go run main.go -d https://google.com -m 500
```

# About

Can be used when you want to test your website to DDOS attacks, autoscaling, HA...

# Problems with max open files (to much sockets opened)

```sh
ulimit -n 1000000
```
