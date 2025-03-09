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

// resents a Prefix expression e.g "-15; or !5;"
type PrefixExpression struct {
	Token token.Token	// the prefix token, e.g. !
	Operator string		// contains either "-" or "!"
	Right Expression	// contains the expression to the right of the expression
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

type AssignStatement struct {
	Token token.Token // the token.ASSIGN token
	Name  *Identifier
	Value Expression
}

type StringLiteral struct {
	Token token.Token // the token.STRING token
	Value string
}
type Boolean struct {
	Token token.Token // the token.BOOL token
	Value bool
}


// represents interger values
type IntegerLiteral struct {
	Token token.Token // The token.INT token
	Value int64
}
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}


func (vs *VarStatement) statementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }

// StatementNode implements Statement.
func (vs *VarStatement) StatementNode() {
	panic("unimplemented")
}


func (i *Identifier) ExpressionNode() {}
func (i *Identifier) TokenLiteral() string {
	if i.Token.Literal == "" {
		return "nil"
	}
	return i.Token.Literal
}



func (as *AssignStatement) statementNode()       {}
func (as *AssignStatement) TokenLiteral() string { return as.Token.Literal }



func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// is an interface for all expressions
type Expression interface {
	Node
	ExpressionNode()
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

// Prefix Expression functions
func (pe *PrefixExpression) expressionNOde() {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}