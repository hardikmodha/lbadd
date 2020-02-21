package command

import "github.com/tomarrell/lbadd/internal/parser/ast"

// Command is the intermediate representation (IR) of an SQL ast.
type Command struct {
}

// From converts the given (*ast.SQLStmt) to the IR, which is a
// (command.Command).
func From(stmt *ast.SQLStmt) (Command, error) {
	return Command{}, nil
}
