// Build the project
go build

// Start the project
go ./rssagg.exe

// Create a table in postgress

goose postgres postgres://postgres:@localhost:5432/rssagg up

// Creates a new module, initializing the go.mod file that describes it.
go mod init

go build, go test, and other package-building commands add new dependencies to go.mod as needed.

// Prints the current module’s dependencies.
go list -m all

// Changes the required version of a dependency (or adds a new dependency).
go get

// Removes unused dependencies.
go mod tidy