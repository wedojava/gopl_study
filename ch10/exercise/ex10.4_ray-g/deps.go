// go run deps.go golang.org/x/net/html
// golang.org/x/net/html/charset
// golang.org/x/tools/cmd/html2article
// gopl.io/ch5/exercise/ex5.01
// gopl.io/ch5/exercise/ex5.02
// gopl.io/ch5/exercise/ex5.03
// gopl.io/ch5/exercise/ex5.04
// gopl.io/ch5/exercise/ex5.05
// gopl.io/ch5/exercise/ex5.07
// gopl.io/ch5/exercise/ex5.07_kdama
// gopl.io/ch5/exercise/ex5.07_kdama/prettyhtml
// gopl.io/ch5/exercise/ex5.08
// gopl.io/ch5/exercise/ex5.12
// gopl.io/ch5/exercise/ex5.13
// gopl.io/ch5/exercise/ex5.17
// gopl.io/ch5/findlinks1
// gopl.io/ch5/findlinks2
// gopl.io/ch5/findlinks3
// gopl.io/ch5/links
// gopl.io/ch5/outline
// gopl.io/ch5/outline2
// gopl.io/ch5/title1
// gopl.io/ch5/title2
// gopl.io/ch5/title3

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func logCommandError(context string, err error) {
	ee, ok := err.(*exec.ExitError)
	if !ok {
		log.Fatalf("%s: %s", context, err)
	}
	log.Printf("%s: %s", context, err)
	os.Stderr.Write(ee.Stderr)
	os.Exit(1)
}

// packages returns a slice of package import paths corresponding to slice of
// package patterns.
// See `go help packages` for different ways of specifying packages.
func packages(patterns []string) []string {
	args := []string{"list", "-f={{.ImportPath}}"}
	// go list -f='{{.ImportPath}}' compress/...
	// compress/bzip2
	// compress/flate
	// compress/gzip
	// compress/lzw
	// compress/zlib
	for _, pkg := range patterns {
		args = append(args, pkg)
	}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		logCommandError("resolve packages", err)
	}
	return strings.Fields(string(out))
}

func ancestors(packageNames []string) []string {
	targets := make(map[string]bool)
	for _, pkg := range packageNames {
		targets[pkg] = true
	}
	args := []string{"list", `-f={{.ImportPath}} {{join .Deps " "}}`, "..."}
	// go list -f='{{.ImportPath}} {{join .Deps " "}}' compress/...
	// compress/bzip2 bufio bytes errors internal/bytealg internal/cpu internal/race internal/reflectlite io runtime runtime/internal/atomic runtime/internal/math runtime/internal/sys sort sync sync/atomic unicode unicode/utf8 unsafe
	// compress/flate bufio bytes errors fmt internal/bytealg internal/cpu internal/fmtsort internal/oserror internal/poll internal/race internal/reflectlite internal/syscall/execenv internal/syscall/unix internal/testlog io math math/bits os reflect runtime runtime/internal/atomic runtime/internal/math runtime/internal/sys sort strconv sync sync/atomic syscall time unicode unicode/utf8 unsafe
	// compress/gzip bufio bytes compress/flate encoding/binary errors fmt hash hash/crc32 internal/bytealg internal/cpu internal/fmtsort internal/oserror internal/poll internal/race internal/reflectlite internal/syscall/execenv internal/syscall/unix internal/testlog io math math/bits os reflect runtime runtime/internal/atomic runtime/internal/math runtime/internal/sys sort strconv sync sync/atomic syscall time unicode unicode/utf8 unsafe
	// compress/lzw bufio bytes errors fmt internal/bytealg internal/cpu internal/fmtsort internal/oserror internal/poll internal/race internal/reflectlite internal/syscall/execenv internal/syscall/unix internal/testlog io math math/bits os reflect runtime runtime/internal/atomic runtime/internal/math runtime/internal/sys sort strconv sync sync/atomic syscall time unicode unicode/utf8 unsafe
	// compress/zlib bufio bytes compress/flate encoding/binary errors fmt hash hash/adler32 internal/bytealg internal/cpu internal/fmtsort internal/oserror internal/poll internal/race internal/reflectlite internal/syscall/execenv internal/syscall/unix internal/testlog io math math/bits os reflect runtime runtime/internal/atomic runtime/internal/math runtime/internal/sys sort strconv sync sync/atomic syscall time unicode unicode/utf8 unsafe
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		logCommandError("find ancestors", err)
	}
	var pkgs []string
	s := bufio.NewScanner(bytes.NewBuffer(out))
	for s.Scan() {
		fields := strings.Fields(s.Text())
		pkg := fields[0]
		deps := fields[1:]
		for _, dep := range deps {
			if targets[dep] {
				pkgs = append(pkgs, pkg)
				break
			}
		}
	}
	return pkgs
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	pkgs := ancestors(packages(os.Args[1:]))
	sort.StringSlice(pkgs).Sort()
	for _, pkg := range pkgs {
		fmt.Println(pkg)
	}
}
