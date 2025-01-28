// package main includes tests for particle engine covering (at least) the CLI packages.
// You can run the tests on this package and get the coverage of these tests
// across the entire sources with the commands:
//
// $ go test -v -coverpkg=./... -coverprofile=profile.cov ./cmd/particle engine
// $ go tool cover -html profile.cov
//
// To get the coverage of all the tests across the entire sources:
// $ go test -v -coverpkg=./... -coverprofile=profile.cov ./cmd/particle engine ./pkg/...
// $ go tool cover -html profile.cov
package main
