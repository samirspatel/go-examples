# go-examples

### Getting started
Enable dep tracking by creating go.mod
`go mod init`

When you add a new package reference in `import`, to install it in go.mod
`go mod tidy`

In order to develop a module and use it locally, replace
`go mod edit -replace github.com/samirspatel/go-examples/greetings=../greetings`

Build and create local executable
`go build` 

Discover go install path
`go list -f '{{.Target}}'`

