// Code generated by "aststring"; DO NOT EDIT.
package ast

import (
	"bytes"
	"fmt"
)

func (node *IdentifierReference) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("IdentifierReference (")
	_, _ = buf.WriteString("\n")
	if node.Identifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.Identifier.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Yield", node.Yield), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Await", node.Await), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *BindingIdentifier) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("BindingIdentifier (")
	_, _ = buf.WriteString("\n")
	if node.Identifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.Identifier.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Yield", node.Yield), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Await", node.Await), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *LabelIdentifier) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("LabelIdentifier (")
	_, _ = buf.WriteString("\n")
	if node.Identifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.Identifier.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Yield", node.Yield), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Await", node.Await), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *Identifier) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("Identifier (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "IdentifierName", node.IdentifierName), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}
