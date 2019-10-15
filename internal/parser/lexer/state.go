package lexer

// state is a definition of a state. The state returned will be the one that is
// executed next.
type state func(*Lexer) state

// lexScript is the default entry point for the lexer. This may be overridden by
// tests. It delegates directly to lexToken, which is responsible for all
// tokens.
func lexScript(l *Lexer) state {
	return lexToken
}

func lexToken(l *Lexer) state {
	if l.eof() {
		return nil
	}

	switch r := l.peek(); r {
	default:
		// handle all cases that cannot be expressed in a switch block. for
		// optimization, matcher matches from here can be converted to cases if
		// necessary.

		if punctuatorStart.Matches(r) {
			return lexPunctuator
		}
	}
	return unexpectedToken
}