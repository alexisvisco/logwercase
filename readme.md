## Logwercaser

Logwercase is a Golang linter that report use of uppercase for message in log printer functions. 

Example:

```go
logrus.Print("Hello world") // ✗ bad
logrus.Print("hello world") // ✓ good 
```

It supports logrus and stdlib logger. 

### Installation

```bash
GO111MODULE=on go get github.com/alexisvisco/logwercase/cmd/logwercaser@0.1
```
