# go-demo

## Compile the coded into executable:

Open a Terminal and run the command:

```sh
cd ./hello
go build
```

The `go build` command compiles the packages, along with their dependencies, but it doesn't install the results.

To run hello executable:

```sh
./hello
```

## Add the Go install directory to your system's shell path.

To see the path to your install directory, run the following command:

```sh
go list -f '{{.Target}}'
```

To be able to run the program's executable _without specifying where_ the executable is, run the command, where **/path/to/your/install/directory** is the result from the previous command:

```sh
export PATH=$PATH:/path/to/your/install/directory
```

To compile and install the package, run:

```sh
go install
```

The go install command compiles and installs the packages.

Now you are able to run the application only by typing its name - `hello`.
