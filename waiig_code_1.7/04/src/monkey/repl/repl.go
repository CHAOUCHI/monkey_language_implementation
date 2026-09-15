package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

const PROMPT = ">> "

/**
* Start starts a REPL (Read-Eval-Print Loop) for the Monkey programming language.
* It takes an input reader and an output writer as parameters.
* The REPL continuously :
*   1. Prompts the user for  a line of code,
* 	2. lexes and parses it,
* 	3. evaluates the resulting AST,
*	4. prints the result.
* If there are any parser errors, they are printed to the output.
 */
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		// 1. Ask the user for one or many line of code
		line := scanner.Text()

		// 2. Make a lexical analyzing of the line of code
		// This creates a array of tokens from the line of code
		// each token is a struct with a type and a literal value
		// exemple : let a = 3
		// - let is a token of type LET and literal value "let"
		// - a is a token of type IDENT and literal value "a"
		// - = is a token of type ASSIGN and literal value "="
		// - 3 is a token of type INT and literal value "3"
		// *the goal is to extract the word(token separate by spaces) of the code into a data structure (variable) that we can read and execute function depending on the type of the token*
		l := lexer.New(line)

		// 3. Make a syntactic analyzing of the array of tokens
		// This creates an AST (Abstract Syntax Tree) from the array of tokens
		// The AST is a tree representation of the code
		// exemple :
		// let a = 3
		// a
		// AST :
		// Program
		// ├── LetStatement
		// │   ├── Identifier (a)
		// │   └── IntegerLiteral (3)
		// └── ExpressionStatement
		//     └── Identifier (a)
		// each token's type have a meaning but this is the work of the parser to understand the meaning of each token
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
