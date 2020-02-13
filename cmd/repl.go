package cmd

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kieranajp/monkey/pkg/lexer"
	"github.com/kieranajp/monkey/pkg/token"
)

// StartRepl begins running the REPL
func StartRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
