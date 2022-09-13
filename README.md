# golangEx
golang Exercises organized into modules

## Setup

go mod init https://github.com/damientl/golangEx/equivTree

go mod init github.com/damientl/golangEx/ex1 

go mod edit -replace github.com/damientl/golangEx/equivTree=../equivTree 

go mod tidy

go run .