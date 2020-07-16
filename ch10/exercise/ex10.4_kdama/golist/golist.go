// Package golist provides an API for the go list tool.
package golist

// Package represents the package metadata output by the go list tool.
type Package struct {
	ImportPath string
	Deps       []string
}
