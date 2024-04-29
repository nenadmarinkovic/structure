### structure

A simple CLI program that adds tree like structure to your README.MD file, like:

### Tree structure of the project:

- go.mod
- helpers.go
- main.go
- structure

### How to run it?

To create executable, navigate to your Go code and run:

```bash
go build -o structure
```

Then move structure app into `/usr/local/bin` (Unix)

Now you can run the application with `structure .`.
It will append tree structure of the project to your README.md file (if it doesn't exist, it will create one).

### Don't have Go installed?

Go to [Go](https://go.dev) for more info
