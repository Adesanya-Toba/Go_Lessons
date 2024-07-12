# Learning Go!
Using the amazing book by *Jon Bodner: Learning Go An Idiomatic Approach to Real-World Go Programming*

### Create a Go Module
A Go project is called a **module**. A module is not just source code. It is also an exact specification of the dependencies of the code within the module. Every module has a go.mod file in its root directory.
```bash
go mod init <program_name>
```

### Build a Go Program with
```bash
go build -o build/go_app
```

### Run a Go Program with
```bash
./executable_name
```

Alternatively, just run `make` using the `M### Build a Go Program withakefile` in the directory.