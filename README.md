# go-file-diff

> This is a simple golang application that allow you to compare 2 JSON files


## Requirements

- Golang 1.16+

[optional] - you can use Docker if you prefer a container execution

### Quickstart

```bash
git clone THIS REPO
cd THIS REPO
go mod verify
go build -v
go install -v
go-file-diff FILE1 FILE2
# output
Source file: FILE1
Target file: FILE2
result : Files Match 👍!
```

### Technical decisions 

#### Project Structure 

Followed https://github.com/golang-standards/project-layout since it is a well explained repository

Decided to use a simple iteration through both objects due to be the simplest,
not the performant, way to compare two inputs

Also decided to use https://github.com/stretchr/testify for unit testing due to be
the go-to way community uses to test

Finally used `pkg/errors` to raise any error that could occur during execution time

 
