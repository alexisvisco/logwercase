## Logwercaser

Logwercase is a Golang linter that report:
 - use of uppercase for message in log printer functions
 - use of non lower kebab case for WithField key in logrus

Example:

```go
logrus.Print("Hello world") // ✗ bad
logrus.Print("hello world") // ✓ good 

logrus.WithField("Hello world", ...) // ✗ bad
logrus.WithField("hello-world", ...) // ✓ good 
```

It supports logrus and stdlib logger. 

### Installation (cli require go)

```bash
GO111MODULE=on go get github.com/alexisvisco/logwercase/cmd/logwercaser@0.3.1
```

### With golangci

```bash
cd  /tmp \
    && git clone https://github.com/alexisvisco/logwercase \
    && cd logwercase \
    && go build -o logwercase.so -buildmode=plugin plugin/plugin.go \
    && cp logwercase.so /home/$USER/go/bin/
    && echo "logwercase.so created and set in " /home/$USER/go/bin/logwercase.so
```

In your .golangci.yml add these lines:

```yaml
linters-settings:
  custom:
    logwercase:
      path: /home/$USER/go/bin/logwercase.so
      description: Analyze case of log message and WithField keys
      original-url: github.com/alexisvisco/logwercase
```

Change $USER with your user
