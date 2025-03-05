package ast

import (
	"bytes"

	"ksm/token"
)

/*A parser is a software component that takes input data (frequently text) and builds
a data structure – often some kind of parse tree, abstract syntax tree or other
hierarchical structure – giving a structural representation of the input, checking for
correct syntax in the process. */

// common interface for all AST nodes.
type Node interface {
	TokenLiteral() string
	String() string
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

// represents a variable declaration(var x = 5)
type VarStatement struct {
	Token token.Token // The token.VAR token
	Name  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }

// StatementNode implements Statement.
func (vs *VarStatement) StatementNode() {
	panic("unimplemented")
}

// represents variable names
type Identifier struct {
	Token token.Token // The token.IDENTIFIER token
	Value string
}

func (i *Identifier) ExpressionNode() {}
func (i *Identifier) TokenLiteral() string {
	if i.Token.Literal == "" {
		return "nil"
	}
	return i.Token.Literal
}

type AssignStatement struct {
	Token token.Token // the token.ASSIGN token
	Name  *Identifier
	Value Expression
}

func (as *AssignStatement) statementNode()       {}
func (as *AssignStatement) TokenLiteral() string { return as.Token.Literal }

type StringLiteral struct {
	Token token.Token // the token.STRING token
	Value string
}
type Boolean struct {
	Token token.Token // the token.BOOL token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// is an interface for all expressions
type Expression interface {
	Node
	ExpressionNode()
}

// represents interger values
type IntegerLiteral struct {
	Token token.Token // The token.INT token
	Value int64
}

func (il *IntegerLiteral) ExpressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (be *BinaryExpression) ExpressionNode() {}
func (be *BinaryExpression) TokenLiteral() string {
	return be.Operator
}

type AssignmentStatement struct {
	Name  string
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (ls *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString(rs.TokenLiteral()+ " ")
	
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (i *Identifier) String() string { return i.Value }
