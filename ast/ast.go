package ast

import "monkey/token"

// Creates an abstract syntax tree which parses the token into the functionality that the code provides

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier // Holds the identifier of the binding
	Value Expression  // The expression which provides the value
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type Expression interface {
	Node
	expressionNode()
}

// Root node of the AST
// A program is a series of statements and these are contained in Program.statmenets which is just a slice of AST nodes that implement the statement interface
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
