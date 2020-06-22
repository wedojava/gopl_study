package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	eval "gopl.io/ch7/exercise/ex7.14"
)

func main() {
	exitCode := 0
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("Expression: ")
	stdin.Scan()
	exprStr := stdin.Text()
	// exprStr = "1/x"
	fmt.Printf("Variables (<var>=<val>, eg: x=3): ")
	stdin.Scan()
	envStr := stdin.Text()
	// envStr = "x=3"
	if stdin.Err() != nil {
		fmt.Fprintln(os.Stderr, stdin.Err())
		os.Exit(1)
	}
	env := eval.Env{}
	assignments := strings.Fields(envStr)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad assignment: %s\n", a)
			exitCode = 2
		}
		ident, valStr := fields[0], fields[1]      // "x", "3"
		val, err := strconv.ParseFloat(valStr, 64) // 3
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad value for %s, using zero: %s\n", ident, err)
			exitCode = 2
		}
		env[eval.Var(ident)] = val // env[eval.Var("x") = 3]
		// env -> ["x": 3]
	}

	expr, err := eval.Parse(exprStr) // exprStr -> "1/x"
	// expr -> {op:47, x: 1, y: "x",}
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad expression: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(expr.Eval(env)) // env = ["x": 3] expr = {op:47, x:1, y:"x"}
	os.Exit(exitCode)
}
