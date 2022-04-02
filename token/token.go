package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, sub, x, y
	INT   = "INT"   // 1 2 3 4

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"

	// Booleans
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// Two var
	EQ     = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TokenType{
	"fn":    FUNCTION,
	"let":   LET,
	"true":  TRUE,
	"false": FALSE,

	"if":   IF,
	"else": ELSE,

	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
