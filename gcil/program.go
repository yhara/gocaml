package gcil

import (
	"fmt"
	"github.com/rhysd/gocaml/typing"
	"io"
	"os"
	"strings"
)

// Program representation. Program can be obtained after closure transform because
// all functions must be at the top.
type Program struct {
	Toplevel map[string]*Fun     // Mapping from function name to its instruction
	Closures map[string][]string // Mapping from closure name to it free variables
	Entry    *Block
}

func (prog *Program) Println(out io.Writer, env *typing.Env) {
	fmt.Fprintln(out, "[TOPLEVELS]")
	p := printer{env, out, ""}
	for n, f := range prog.Toplevel {
		p.printlnInsn(NewInsn(n, f))
		fmt.Fprintln(out)
	}

	fmt.Fprintln(out, "[CLOSURES]")
	for c, fv := range prog.Closures {
		fmt.Fprintf(out, "%s:\t%s\n", c, strings.Join(fv, ","))
	}
	fmt.Fprintln(out)

	fmt.Fprintln(out, "[ENTRY]")
	prog.Entry.Println(out, env)
}

func (prog *Program) Dump(env *typing.Env) {
	prog.Println(os.Stdout, env)
}
