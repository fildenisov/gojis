// Code generated by "aststring"; DO NOT EDIT.
package ast

import (
	"bytes"
	"fmt"
)

func (node *Literal) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("Literal (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "NullLiteral", node.NullLiteral), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "BooleanLiteral", node.BooleanLiteral), "  "))
	_, _ = buf.WriteString("\n")
	if node.NumericLiteral != nil {
		_, _ = buf.WriteString(PrefixToString(node.NumericLiteral.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "StringLiteral", node.StringLiteral), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *NumericLiteral) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("NumericLiteral (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "DecimalLiteral", node.DecimalLiteral), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "BinaryIntegerLiteral", node.BinaryIntegerLiteral), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "OctalIntegerLiteral", node.OctalIntegerLiteral), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "HexIntegerLiteral", node.HexIntegerLiteral), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ArrayLiteral) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ArrayLiteral (")
	_, _ = buf.WriteString("\n")
	if node.Elision != nil {
		_, _ = buf.WriteString(PrefixToString(node.Elision.String(), "  "))
	}
	if node.ElementList != nil {
		_, _ = buf.WriteString(PrefixToString(node.ElementList.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Comma", node.Comma), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ElementList) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ElementList (")
	_, _ = buf.WriteString("\n")
	if node.Elision != nil {
		_, _ = buf.WriteString(PrefixToString(node.Elision.String(), "  "))
	}
	if node.AssignmentExpression != nil {
		_, _ = buf.WriteString(PrefixToString(node.AssignmentExpression.String(), "  "))
	}
	if node.SpreadElement != nil {
		_, _ = buf.WriteString(PrefixToString(node.SpreadElement.String(), "  "))
	}
	if node.ElementList != nil {
		_, _ = buf.WriteString(PrefixToString(node.ElementList.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Comma", node.Comma), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *SpreadElement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SpreadElement (")
	_, _ = buf.WriteString("\n")
	if node.AssignmentExpression != nil {
		_, _ = buf.WriteString(PrefixToString(node.AssignmentExpression.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ObjectLiteral) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ObjectLiteral (")
	_, _ = buf.WriteString("\n")
	if node.PropertyDefinitionList != nil {
		_, _ = buf.WriteString(PrefixToString(node.PropertyDefinitionList.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Comma", node.Comma), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *PropertyDefinitionList) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("PropertyDefinitionList (")
	_, _ = buf.WriteString("\n")
	for _, elem := range node.PropertyDefinitions {
		_, _ = buf.WriteString(PrefixToString(elem.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *PropertyDefinition) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("PropertyDefinition (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Colon", node.Colon), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Ellipsis", node.Ellipsis), "  "))
	_, _ = buf.WriteString("\n")
	if node.IdentifierReference != nil {
		_, _ = buf.WriteString(PrefixToString(node.IdentifierReference.String(), "  "))
	}
	if node.CoverInitializedName != nil {
		_, _ = buf.WriteString(PrefixToString(node.CoverInitializedName.String(), "  "))
	}
	if node.PropertyName != nil {
		_, _ = buf.WriteString(PrefixToString(node.PropertyName.String(), "  "))
	}
	if node.AssignmentExpression != nil {
		_, _ = buf.WriteString(PrefixToString(node.AssignmentExpression.String(), "  "))
	}
	if node.MethodDefinition != nil {
		_, _ = buf.WriteString(PrefixToString(node.MethodDefinition.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *PropertyName) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("PropertyName (")
	_, _ = buf.WriteString("\n")
	if node.LiteralPropertyName != nil {
		_, _ = buf.WriteString(PrefixToString(node.LiteralPropertyName.String(), "  "))
	}
	if node.ComputedPropertyName != nil {
		_, _ = buf.WriteString(PrefixToString(node.ComputedPropertyName.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *LiteralPropertyName) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("LiteralPropertyName (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "IdentifierName", node.IdentifierName), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "StringLiteral", node.StringLiteral), "  "))
	_, _ = buf.WriteString("\n")
	if node.NumericLiteral != nil {
		_, _ = buf.WriteString(PrefixToString(node.NumericLiteral.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ComputedPropertyName) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ComputedPropertyName (")
	_, _ = buf.WriteString("\n")
	if node.AssignmentExpression != nil {
		_, _ = buf.WriteString(PrefixToString(node.AssignmentExpression.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *CoverInitializedName) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("CoverInitializedName (")
	_, _ = buf.WriteString("\n")
	if node.IdentifierReference != nil {
		_, _ = buf.WriteString(PrefixToString(node.IdentifierReference.String(), "  "))
	}
	if node.Initializer != nil {
		_, _ = buf.WriteString(PrefixToString(node.Initializer.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *TemplateLiteral) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("TemplateLiteral (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "NoSubstitutionTemplate", node.NoSubstitutionTemplate), "  "))
	_, _ = buf.WriteString("\n")
	if node.SubstitutionTemplate != nil {
		_, _ = buf.WriteString(PrefixToString(node.SubstitutionTemplate.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *SubstitutionTemplate) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SubstitutionTemplate (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "TemplateHead", node.TemplateHead), "  "))
	_, _ = buf.WriteString("\n")
	if node.Expression != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression.String(), "  "))
	}
	if node.TemplateSpans != nil {
		_, _ = buf.WriteString(PrefixToString(node.TemplateSpans.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *TemplateSpans) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("TemplateSpans (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "TemplateTail", node.TemplateTail), "  "))
	_, _ = buf.WriteString("\n")
	if node.TemplateMiddleList != nil {
		_, _ = buf.WriteString(PrefixToString(node.TemplateMiddleList.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *TemplateMiddleList) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("TemplateMiddleList (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "TemplateMiddle", node.TemplateMiddle), "  "))
	_, _ = buf.WriteString("\n")
	if node.Expression != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression.String(), "  "))
	}
	if node.TemplateMiddleList != nil {
		_, _ = buf.WriteString(PrefixToString(node.TemplateMiddleList.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *RegularExpressionLiteral) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("RegularExpressionLiteral (")
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "RegularExpressionBody", node.RegularExpressionBody), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "RegularExpressionFlags", node.RegularExpressionFlags), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}
