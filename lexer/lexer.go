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

func (l *Lexer) readChar(){
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}