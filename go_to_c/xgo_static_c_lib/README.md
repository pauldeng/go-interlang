# Calling a C static library from Go

To build the C static library and Go executable, run `make`. Then run
`./go_to_c_static` to run the Go program.

## Another similar example

[github.com/shadowmint/go-static-linking](https://github.com/shadowmint/go-static-linking)

mkdir build
cmake -G 'Unix Makefiles' -DCMAKE_BUILD_TYPE=Debug ..