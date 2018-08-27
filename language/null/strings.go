package Null

import . "github.com/qlova/script/language"

//Returns a Symbol that the Go style literal represents ('').
func (l *language) LiteralSymbol(literal string) Symbol {
	panic("Error in "+Name+".LiteralSymbol("+literal+"): Unimplemented")
	return nil
}

//Returns a String that the Go style literal represents ("").
func (l *language) LiteralString(literal string) String {
	panic("Error in "+Name+".LiteralString("+literal+"): Unimplemented")
	return nil
}

//Returns a new String that concatenates 'a' and 'b'.
func (l *language) JoinString(a, b String) String {
	panic("Error in "+Name+".JoinString(String, String): Unimplemented")
	return nil
}
