package ast

import "ksm/token"

/*A parser is a software component that takes input data (frequently text) and builds
a data structure – often some kind of parse tree, abstract syntax tree or other
hierarchical structure – giving a structural representation of the input, checking for
correct syntax in the process. */

// common interface for all AST nodes.
type Node interface {
	TokenLiteral() string
}

// represents the whole program
type Program struct {
	Statements []Statement
}

// interface representing a statement
type Statement interface {
	Node
	statementNode()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// is an interface for all expressions
type Expression interface {
	Node
	expressionNode()
}

// represents a variable declaration(var x = 5)
type VarStatement struct {
	Token token.Token // The token.VAR token
	Name  *Identifier
	Value Expression
}

// represents variable names
type Identifier struct {
	Token token.Token // The token.IDENTIFIER token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	if i.Token.Literal == "" {
		return "nil"
	}
	return i.Token.Literal
}

// represents interger values
// type IntegerLiteral struct {
// 	Token token.Token // The token.INT token
// 	Value int64
// }

type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

type AssignmentStatement struct {
	Name  string
	Value Expression
}

func (ls *VarStatement) statementNode()       {}
func (ls *VarStatement) TokenLiteral() string { return ls.Token.Literal }

type ReturnStatment struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatment) statementNode()       {}
func (rs *ReturnStatment) TokenLiteral() string { return rs.Token.Literal }
