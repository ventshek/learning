/* Run go code by using 'go run xxx.go' */

go build = 'Compiles a bunch of Go source code.'
go run = 'Compiles and executes one or two files.'
go fmt = 'Formats all the code in each file in the current directory.'
go install = 'Compiles and "installs" a package.'
go get = "Downloads the raw source code of someone else's"
go test = 'Runs any tests asociated with the current project. '

There are two types of packages in Go, Executable and Reusable

Executable = Generates a file that we can run.
Reusable = Code used as helpers, good place to put reusable logic. 

package main = Defines a package that can be compiled and then *executed*. Must have a funx called main. 

list of golang packages = "https://pkg.go.dev/std"

main.go is always organised as package > import > functions + variable declarations. 