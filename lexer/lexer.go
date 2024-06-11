package lexer

type Lexer struct {
	input		 string
	position	 int //cur pos in input
	readPosition int //cur reading position in input	
	ch 			 byte //current char 
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}