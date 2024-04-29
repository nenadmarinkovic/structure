# structure tree 🌳

A simple CLI program that adds a tree-like structure to your README.MD file, like:

## Project Structure

- go.mod
- helpers.go
- main.go
- structure

## How to Run

1. Navigate to your Go code directory and run the following command to create an executable:

```bash
go build -o structure
```

2. Move the structure app into /usr/local/bin:

```bash
mv structure /usr/local/bin
```

3. Now, go to any folder and run the structure command:

```bash
structure .
```

This will append a tree-like structure of the project to your README.md file. If README.md doesn't exist, it will create one.

### Don't have Go installed?

Go to [Go](https://go.dev) for more info.
