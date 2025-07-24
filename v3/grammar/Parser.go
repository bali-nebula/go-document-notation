/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	ast "github.com/bali-nebula/go-document-notation/v3/ast"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	mat "math"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ParserClass() ParserClassLike {
	return parserClass()
}

// Constructor Methods

func (c *parserClass_) Parser() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserClass()
}

func (v *parser_) ParseSource(
	source string,
) ast.DocumentLike {
	v.source_ = sts.ReplaceAll(source, "\t", "    ")
	v.tokens_ = fra.Queue[TokenLike]()
	v.next_ = fra.Stack[TokenLike]()

	// The scanner runs in a separate Go routine.
	ScannerClass().Scanner(v.source_, v.tokens_)

	// Attempt to parse the document.
	var document, token, ok = v.parseDocument()
	if !ok || !v.tokens_.IsEmpty() {
		var message = v.formatError("$Document", token)
		panic(message)
	}
	return document
}

// PROTECTED INTERFACE

// Private Methods

func (v *parser_) parseAcceptClause() (
	acceptClause ast.AcceptClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "accept" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("accept")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AcceptClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AcceptClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Message rule.
	var message ast.MessageLike
	message, token, ok = v.parseMessage()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Message rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AcceptClause", token)
		panic(message)
	}

	// Found a single AcceptClause rule.
	ok = true
	v.remove(tokens)
	acceptClause = ast.AcceptClauseClass().AcceptClause(
		delimiter,
		message,
	)
	return
}

func (v *parser_) parseActionInduction() (
	actionInduction ast.ActionInductionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single DoClause ActionInduction.
	var doClause ast.DoClauseLike
	doClause, token, ok = v.parseDoClause()
	if ok {
		// Found a single DoClause ActionInduction.
		actionInduction = ast.ActionInductionClass().ActionInduction(doClause)
		return
	}

	// Attempt to parse a single LetClause ActionInduction.
	var letClause ast.LetClauseLike
	letClause, token, ok = v.parseLetClause()
	if ok {
		// Found a single LetClause ActionInduction.
		actionInduction = ast.ActionInductionClass().ActionInduction(letClause)
		return
	}

	// This is not a single ActionInduction rule.
	return
}

func (v *parser_) parseAnnotation() (
	annotation ast.AnnotationLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single note Annotation.
	var note string
	note, token, ok = v.parseToken(NoteToken)
	if ok {
		// Found a single note Annotation.
		annotation = ast.AnnotationClass().Annotation(note)
		return
	}

	// Attempt to parse a single comment Annotation.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if ok {
		// Found a single comment Annotation.
		annotation = ast.AnnotationClass().Annotation(comment)
		return
	}

	// This is not a single Annotation rule.
	return
}

func (v *parser_) parseArgument() (
	argument ast.ArgumentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Value Argument.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	if ok {
		// Found a single Value Argument.
		argument = ast.ArgumentClass().Argument(value)
		return
	}

	// Attempt to parse a single Primitive Argument.
	var primitive ast.PrimitiveLike
	primitive, token, ok = v.parsePrimitive()
	if ok {
		// Found a single Primitive Argument.
		argument = ast.ArgumentClass().Argument(primitive)
		return
	}

	// This is not a single Argument rule.
	return
}

func (v *parser_) parseArithmeticOperator() (
	arithmeticOperator ast.ArithmeticOperatorLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "+" delimiter.
	delimiter, token, ok = v.parseDelimiter("+")
	if ok {
		// Found a single "+" delimiter.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(delimiter)
		return
	}

	// Attempt to parse a single "-" delimiter.
	delimiter, token, ok = v.parseDelimiter("-")
	if ok {
		// Found a single "-" delimiter.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(delimiter)
		return
	}

	// Attempt to parse a single "*" delimiter.
	delimiter, token, ok = v.parseDelimiter("*")
	if ok {
		// Found a single "*" delimiter.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(delimiter)
		return
	}

	// Attempt to parse a single "/" delimiter.
	delimiter, token, ok = v.parseDelimiter("/")
	if ok {
		// Found a single "/" delimiter.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(delimiter)
		return
	}

	// Attempt to parse a single "%" delimiter.
	delimiter, token, ok = v.parseDelimiter("%")
	if ok {
		// Found a single "%" delimiter.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(delimiter)
		return
	}

	// Attempt to parse a single "^" delimiter.
	delimiter, token, ok = v.parseDelimiter("^")
	if ok {
		// Found a single "^" delimiter.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(delimiter)
		return
	}

	// This is not a single ArithmeticOperator rule.
	return
}

func (v *parser_) parseAssignment() (
	assignment ast.AssignmentLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single ":=" delimiter.
	delimiter, token, ok = v.parseDelimiter(":=")
	if ok {
		// Found a single ":=" delimiter.
		assignment = ast.AssignmentClass().Assignment(delimiter)
		return
	}

	// Attempt to parse a single "?=" delimiter.
	delimiter, token, ok = v.parseDelimiter("?=")
	if ok {
		// Found a single "?=" delimiter.
		assignment = ast.AssignmentClass().Assignment(delimiter)
		return
	}

	// Attempt to parse a single "+=" delimiter.
	delimiter, token, ok = v.parseDelimiter("+=")
	if ok {
		// Found a single "+=" delimiter.
		assignment = ast.AssignmentClass().Assignment(delimiter)
		return
	}

	// Attempt to parse a single "-=" delimiter.
	delimiter, token, ok = v.parseDelimiter("-=")
	if ok {
		// Found a single "-=" delimiter.
		assignment = ast.AssignmentClass().Assignment(delimiter)
		return
	}

	// Attempt to parse a single "*=" delimiter.
	delimiter, token, ok = v.parseDelimiter("*=")
	if ok {
		// Found a single "*=" delimiter.
		assignment = ast.AssignmentClass().Assignment(delimiter)
		return
	}

	// Attempt to parse a single "/=" delimiter.
	delimiter, token, ok = v.parseDelimiter("/=")
	if ok {
		// Found a single "/=" delimiter.
		assignment = ast.AssignmentClass().Assignment(delimiter)
		return
	}

	// This is not a single Assignment rule.
	return
}

func (v *parser_) parseAssociation() (
	association ast.AssociationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Primitive rule.
	var primitive ast.PrimitiveLike
	primitive, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Primitive rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Association", token)
		panic(message)
	}

	// Attempt to parse a single ":" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Association rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Association", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Document rule.
	var document ast.DocumentLike
	document, token, ok = v.parseDocument()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Document rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Association", token)
		panic(message)
	}

	// Found a single Association rule.
	ok = true
	v.remove(tokens)
	association = ast.AssociationClass().Association(
		primitive,
		delimiter,
		document,
	)
	return
}

func (v *parser_) parseAtLevel() (
	atLevel ast.AtLevelLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "at" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("at")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AtLevel rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AtLevel", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "level" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("level")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AtLevel rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AtLevel", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AtLevel", token)
		panic(message)
	}

	// Found a single AtLevel rule.
	ok = true
	v.remove(tokens)
	atLevel = ast.AtLevelClass().AtLevel(
		delimiter1,
		delimiter2,
		expression,
	)
	return
}

func (v *parser_) parseAttributes() (
	attributes ast.AttributesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Attributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Attributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Association rules.
	var associations = fra.List[ast.AssociationLike]()
associationsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var association ast.AssociationLike
		association, token, ok = v.parseAssociation()
		if !ok {
			switch {
			case count_ >= 1:
				break associationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Association rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Attributes", token)
				message += "1 or more Association rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		associations.AppendValue(association)
	}

	// Attempt to parse a single "]" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Attributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Attributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Attributes rule.
	ok = true
	v.remove(tokens)
	attributes = ast.AttributesClass().Attributes(
		delimiter1,
		associations,
		delimiter2,
	)
	return
}

func (v *parser_) parseBag() (
	bag ast.BagLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Bag", token)
		panic(message)
	}

	// Found a single Bag rule.
	ok = true
	v.remove(tokens)
	bag = ast.BagClass().Bag(expression)
	return
}

func (v *parser_) parseBra() (
	bra ast.BraLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "[" delimiter.
	delimiter, token, ok = v.parseDelimiter("[")
	if ok {
		// Found a single "[" delimiter.
		bra = ast.BraClass().Bra(delimiter)
		return
	}

	// Attempt to parse a single "(" delimiter.
	delimiter, token, ok = v.parseDelimiter("(")
	if ok {
		// Found a single "(" delimiter.
		bra = ast.BraClass().Bra(delimiter)
		return
	}

	// This is not a single Bra rule.
	return
}

func (v *parser_) parseBreakClause() (
	breakClause ast.BreakClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "break" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("break")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single BreakClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$BreakClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "loop" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("loop")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single BreakClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$BreakClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single BreakClause rule.
	ok = true
	v.remove(tokens)
	breakClause = ast.BreakClauseClass().BreakClause(
		delimiter1,
		delimiter2,
	)
	return
}

func (v *parser_) parseCheckoutClause() (
	checkoutClause ast.CheckoutClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "checkout" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("checkout")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single CheckoutClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$CheckoutClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Recipient rule.
	var recipient ast.RecipientLike
	recipient, token, ok = v.parseRecipient()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Recipient rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$CheckoutClause", token)
		panic(message)
	}

	// Attempt to parse an optional AtLevel rule.
	var optionalAtLevel ast.AtLevelLike
	optionalAtLevel, _, ok = v.parseAtLevel()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single "from" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("from")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single CheckoutClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$CheckoutClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Cited rule.
	var cited ast.CitedLike
	cited, token, ok = v.parseCited()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Cited rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$CheckoutClause", token)
		panic(message)
	}

	// Found a single CheckoutClause rule.
	ok = true
	v.remove(tokens)
	checkoutClause = ast.CheckoutClauseClass().CheckoutClause(
		delimiter1,
		recipient,
		optionalAtLevel,
		delimiter2,
		cited,
	)
	return
}

func (v *parser_) parseCited() (
	cited ast.CitedLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Cited", token)
		panic(message)
	}

	// Found a single Cited rule.
	ok = true
	v.remove(tokens)
	cited = ast.CitedClass().Cited(expression)
	return
}

func (v *parser_) parseCollection() (
	collection ast.CollectionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Attributes Collection.
	var attributes ast.AttributesLike
	attributes, token, ok = v.parseAttributes()
	if ok {
		// Found a single Attributes Collection.
		collection = ast.CollectionClass().Collection(attributes)
		return
	}

	// Attempt to parse a single Items Collection.
	var items ast.ItemsLike
	items, token, ok = v.parseItems()
	if ok {
		// Found a single Items Collection.
		collection = ast.CollectionClass().Collection(items)
		return
	}

	// This is not a single Collection rule.
	return
}

func (v *parser_) parseComparisonOperator() (
	comparisonOperator ast.ComparisonOperatorLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "<" delimiter.
	delimiter, token, ok = v.parseDelimiter("<")
	if ok {
		// Found a single "<" delimiter.
		comparisonOperator = ast.ComparisonOperatorClass().ComparisonOperator(delimiter)
		return
	}

	// Attempt to parse a single "=" delimiter.
	delimiter, token, ok = v.parseDelimiter("=")
	if ok {
		// Found a single "=" delimiter.
		comparisonOperator = ast.ComparisonOperatorClass().ComparisonOperator(delimiter)
		return
	}

	// Attempt to parse a single ">" delimiter.
	delimiter, token, ok = v.parseDelimiter(">")
	if ok {
		// Found a single ">" delimiter.
		comparisonOperator = ast.ComparisonOperatorClass().ComparisonOperator(delimiter)
		return
	}

	// Attempt to parse a single "is" delimiter.
	delimiter, token, ok = v.parseDelimiter("is")
	if ok {
		// Found a single "is" delimiter.
		comparisonOperator = ast.ComparisonOperatorClass().ComparisonOperator(delimiter)
		return
	}

	// Attempt to parse a single "matches" delimiter.
	delimiter, token, ok = v.parseDelimiter("matches")
	if ok {
		// Found a single "matches" delimiter.
		comparisonOperator = ast.ComparisonOperatorClass().ComparisonOperator(delimiter)
		return
	}

	// This is not a single ComparisonOperator rule.
	return
}

func (v *parser_) parseComplement() (
	complement ast.ComplementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "not" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("not")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Complement rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Complement", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Logical rule.
	var logical ast.LogicalLike
	logical, token, ok = v.parseLogical()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Logical rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Complement", token)
		panic(message)
	}

	// Found a single Complement rule.
	ok = true
	v.remove(tokens)
	complement = ast.ComplementClass().Complement(
		delimiter,
		logical,
	)
	return
}

func (v *parser_) parseComponent() (
	component ast.ComponentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Element Component.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	if ok {
		// Found a single Element Component.
		component = ast.ComponentClass().Component(element)
		return
	}

	// Attempt to parse a single String Component.
	var string_ ast.StringLike
	string_, token, ok = v.parseString()
	if ok {
		// Found a single String Component.
		component = ast.ComponentClass().Component(string_)
		return
	}

	// Attempt to parse a single Range Component.
	var range_ ast.RangeLike
	range_, token, ok = v.parseRange()
	if ok {
		// Found a single Range Component.
		component = ast.ComponentClass().Component(range_)
		return
	}

	// Attempt to parse a single Collection Component.
	var collection ast.CollectionLike
	collection, token, ok = v.parseCollection()
	if ok {
		// Found a single Collection Component.
		component = ast.ComponentClass().Component(collection)
		return
	}

	// Attempt to parse a single Procedure Component.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	if ok {
		// Found a single Procedure Component.
		component = ast.ComponentClass().Component(procedure)
		return
	}

	// This is not a single Component rule.
	return
}

func (v *parser_) parseCondition() (
	condition ast.ConditionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Condition", token)
		panic(message)
	}

	// Found a single Condition rule.
	ok = true
	v.remove(tokens)
	condition = ast.ConditionClass().Condition(expression)
	return
}

func (v *parser_) parseContinueClause() (
	continueClause ast.ContinueClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "continue" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("continue")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ContinueClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ContinueClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "loop" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("loop")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ContinueClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ContinueClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ContinueClause rule.
	ok = true
	v.remove(tokens)
	continueClause = ast.ContinueClauseClass().ContinueClause(
		delimiter1,
		delimiter2,
	)
	return
}

func (v *parser_) parseDiscardClause() (
	discardClause ast.DiscardClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "discard" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("discard")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single DiscardClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$DiscardClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Draft rule.
	var draft ast.DraftLike
	draft, token, ok = v.parseDraft()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Draft rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$DiscardClause", token)
		panic(message)
	}

	// Found a single DiscardClause rule.
	ok = true
	v.remove(tokens)
	discardClause = ast.DiscardClauseClass().DiscardClause(
		delimiter,
		draft,
	)
	return
}

func (v *parser_) parseDoClause() (
	doClause ast.DoClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "do" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single DoClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$DoClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Invocation rule.
	var invocation ast.InvocationLike
	invocation, token, ok = v.parseInvocation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Invocation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$DoClause", token)
		panic(message)
	}

	// Found a single DoClause rule.
	ok = true
	v.remove(tokens)
	doClause = ast.DoClauseClass().DoClause(
		delimiter,
		invocation,
	)
	return
}

func (v *parser_) parseDocument() (
	document ast.DocumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Document", token)
		panic(message)
	}

	// Attempt to parse an optional Parameters rule.
	var optionalParameters ast.ParametersLike
	optionalParameters, _, ok = v.parseParameters()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok {
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	} else {
		optionalNote = "" // Reset this to undefined.
	}

	// Found a single Document rule.
	ok = true
	v.remove(tokens)
	document = ast.DocumentClass().Document(
		component,
		optionalParameters,
		optionalNote,
	)
	return
}

func (v *parser_) parseDraft() (
	draft ast.DraftLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Draft", token)
		panic(message)
	}

	// Found a single Draft rule.
	ok = true
	v.remove(tokens)
	draft = ast.DraftClass().Draft(expression)
	return
}

func (v *parser_) parseElement() (
	element ast.ElementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single angle Element.
	var angle string
	angle, token, ok = v.parseToken(AngleToken)
	if ok {
		// Found a single angle Element.
		element = ast.ElementClass().Element(angle)
		return
	}

	// Attempt to parse a single boolean Element.
	var boolean string
	boolean, token, ok = v.parseToken(BooleanToken)
	if ok {
		// Found a single boolean Element.
		element = ast.ElementClass().Element(boolean)
		return
	}

	// Attempt to parse a single duration Element.
	var duration string
	duration, token, ok = v.parseToken(DurationToken)
	if ok {
		// Found a single duration Element.
		element = ast.ElementClass().Element(duration)
		return
	}

	// Attempt to parse a single glyph Element.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if ok {
		// Found a single glyph Element.
		element = ast.ElementClass().Element(glyph)
		return
	}

	// Attempt to parse a single moment Element.
	var moment string
	moment, token, ok = v.parseToken(MomentToken)
	if ok {
		// Found a single moment Element.
		element = ast.ElementClass().Element(moment)
		return
	}

	// Attempt to parse a single number Element.
	var number string
	number, token, ok = v.parseToken(NumberToken)
	if ok {
		// Found a single number Element.
		element = ast.ElementClass().Element(number)
		return
	}

	// Attempt to parse a single percentage Element.
	var percentage string
	percentage, token, ok = v.parseToken(PercentageToken)
	if ok {
		// Found a single percentage Element.
		element = ast.ElementClass().Element(percentage)
		return
	}

	// Attempt to parse a single probability Element.
	var probability string
	probability, token, ok = v.parseToken(ProbabilityToken)
	if ok {
		// Found a single probability Element.
		element = ast.ElementClass().Element(probability)
		return
	}

	// Attempt to parse a single resource Element.
	var resource string
	resource, token, ok = v.parseToken(ResourceToken)
	if ok {
		// Found a single resource Element.
		element = ast.ElementClass().Element(resource)
		return
	}

	// Attempt to parse a single symbol Element.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if ok {
		// Found a single symbol Element.
		element = ast.ElementClass().Element(symbol)
		return
	}

	// This is not a single Element rule.
	return
}

func (v *parser_) parseEntity() (
	entity ast.EntityLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Document rule.
	var document ast.DocumentLike
	document, token, ok = v.parseDocument()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Document rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Entity", token)
		panic(message)
	}

	// Found a single Entity rule.
	ok = true
	v.remove(tokens)
	entity = ast.EntityClass().Entity(document)
	return
}

func (v *parser_) parseEvent() (
	event ast.EventLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Event", token)
		panic(message)
	}

	// Found a single Event rule.
	ok = true
	v.remove(tokens)
	event = ast.EventClass().Event(expression)
	return
}

func (v *parser_) parseException() (
	exception ast.ExceptionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Exception", token)
		panic(message)
	}

	// Found a single Exception rule.
	ok = true
	v.remove(tokens)
	exception = ast.ExceptionClass().Exception(expression)
	return
}

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Subject rule.
	var subject ast.SubjectLike
	subject, token, ok = v.parseSubject()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Subject rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Expression", token)
		panic(message)
	}

	// Attempt to parse multiple Predicate rules.
	var predicates = fra.List[ast.PredicateLike]()
predicatesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var predicate ast.PredicateLike
		predicate, token, ok = v.parsePredicate()
		if !ok {
			switch {
			case count_ >= 0:
				break predicatesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Predicate rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Expression", token)
				message += "0 or more Predicate rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		predicates.AppendValue(predicate)
	}

	// Found a single Expression rule.
	ok = true
	v.remove(tokens)
	expression = ast.ExpressionClass().Expression(
		subject,
		predicates,
	)
	return
}

func (v *parser_) parseFailure() (
	failure ast.FailureLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Failure", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Failure rule.
	ok = true
	v.remove(tokens)
	failure = ast.FailureClass().Failure(symbol)
	return
}

func (v *parser_) parseFlowControl() (
	flowControl ast.FlowControlLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single IfClause FlowControl.
	var ifClause ast.IfClauseLike
	ifClause, token, ok = v.parseIfClause()
	if ok {
		// Found a single IfClause FlowControl.
		flowControl = ast.FlowControlClass().FlowControl(ifClause)
		return
	}

	// Attempt to parse a single SelectClause FlowControl.
	var selectClause ast.SelectClauseLike
	selectClause, token, ok = v.parseSelectClause()
	if ok {
		// Found a single SelectClause FlowControl.
		flowControl = ast.FlowControlClass().FlowControl(selectClause)
		return
	}

	// Attempt to parse a single WhileClause FlowControl.
	var whileClause ast.WhileClauseLike
	whileClause, token, ok = v.parseWhileClause()
	if ok {
		// Found a single WhileClause FlowControl.
		flowControl = ast.FlowControlClass().FlowControl(whileClause)
		return
	}

	// Attempt to parse a single WithClause FlowControl.
	var withClause ast.WithClauseLike
	withClause, token, ok = v.parseWithClause()
	if ok {
		// Found a single WithClause FlowControl.
		flowControl = ast.FlowControlClass().FlowControl(withClause)
		return
	}

	// Attempt to parse a single ContinueClause FlowControl.
	var continueClause ast.ContinueClauseLike
	continueClause, token, ok = v.parseContinueClause()
	if ok {
		// Found a single ContinueClause FlowControl.
		flowControl = ast.FlowControlClass().FlowControl(continueClause)
		return
	}

	// Attempt to parse a single BreakClause FlowControl.
	var breakClause ast.BreakClauseLike
	breakClause, token, ok = v.parseBreakClause()
	if ok {
		// Found a single BreakClause FlowControl.
		flowControl = ast.FlowControlClass().FlowControl(breakClause)
		return
	}

	// Attempt to parse a single ReturnClause FlowControl.
	var returnClause ast.ReturnClauseLike
	returnClause, token, ok = v.parseReturnClause()
	if ok {
		// Found a single ReturnClause FlowControl.
		flowControl = ast.FlowControlClass().FlowControl(returnClause)
		return
	}

	// Attempt to parse a single ThrowClause FlowControl.
	var throwClause ast.ThrowClauseLike
	throwClause, token, ok = v.parseThrowClause()
	if ok {
		// Found a single ThrowClause FlowControl.
		flowControl = ast.FlowControlClass().FlowControl(throwClause)
		return
	}

	// This is not a single FlowControl rule.
	return
}

func (v *parser_) parseFunction() (
	function ast.FunctionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Function", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Function rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Function", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Argument rules.
	var arguments = fra.List[ast.ArgumentLike]()
argumentsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var argument ast.ArgumentLike
		argument, token, ok = v.parseArgument()
		if !ok {
			switch {
			case count_ >= 0:
				break argumentsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Argument rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Function", token)
				message += "0 or more Argument rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		arguments.AppendValue(argument)
	}

	// Attempt to parse a single ")" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Function rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Function", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Function rule.
	ok = true
	v.remove(tokens)
	function = ast.FunctionClass().Function(
		identifier,
		delimiter1,
		arguments,
		delimiter2,
	)
	return
}

func (v *parser_) parseIfClause() (
	ifClause ast.IfClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "if" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("if")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single IfClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$IfClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Condition rule.
	var condition ast.ConditionLike
	condition, token, ok = v.parseCondition()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Condition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$IfClause", token)
		panic(message)
	}

	// Attempt to parse a single "do" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single IfClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$IfClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Procedure rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$IfClause", token)
		panic(message)
	}

	// Found a single IfClause rule.
	ok = true
	v.remove(tokens)
	ifClause = ast.IfClauseClass().IfClause(
		delimiter1,
		condition,
		delimiter2,
		procedure,
	)
	return
}

func (v *parser_) parseIndex() (
	index ast.IndexLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Value Index.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	if ok {
		// Found a single Value Index.
		index = ast.IndexClass().Index(value)
		return
	}

	// Attempt to parse a single Primitive Index.
	var primitive ast.PrimitiveLike
	primitive, token, ok = v.parsePrimitive()
	if ok {
		// Found a single Primitive Index.
		index = ast.IndexClass().Index(primitive)
		return
	}

	// This is not a single Index rule.
	return
}

func (v *parser_) parseIndirect() (
	indirect ast.IndirectLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Document Indirect.
	var document ast.DocumentLike
	document, token, ok = v.parseDocument()
	if ok {
		// Found a single Document Indirect.
		indirect = ast.IndirectClass().Indirect(document)
		return
	}

	// Attempt to parse a single Subcomponent Indirect.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Indirect.
		indirect = ast.IndirectClass().Indirect(subcomponent)
		return
	}

	// Attempt to parse a single Referent Indirect.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Indirect.
		indirect = ast.IndirectClass().Indirect(referent)
		return
	}

	// Attempt to parse a single Function Indirect.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Indirect.
		indirect = ast.IndirectClass().Indirect(function)
		return
	}

	// Attempt to parse a single Method Indirect.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Indirect.
		indirect = ast.IndirectClass().Indirect(method)
		return
	}

	// Attempt to parse a single Value Indirect.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	if ok {
		// Found a single Value Indirect.
		indirect = ast.IndirectClass().Indirect(value)
		return
	}

	// This is not a single Indirect rule.
	return
}

func (v *parser_) parseInverse() (
	inverse ast.InverseLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "-" delimiter.
	delimiter, token, ok = v.parseDelimiter("-")
	if ok {
		// Found a single "-" delimiter.
		inverse = ast.InverseClass().Inverse(delimiter)
		return
	}

	// Attempt to parse a single "/" delimiter.
	delimiter, token, ok = v.parseDelimiter("/")
	if ok {
		// Found a single "/" delimiter.
		inverse = ast.InverseClass().Inverse(delimiter)
		return
	}

	// Attempt to parse a single "*" delimiter.
	delimiter, token, ok = v.parseDelimiter("*")
	if ok {
		// Found a single "*" delimiter.
		inverse = ast.InverseClass().Inverse(delimiter)
		return
	}

	// This is not a single Inverse rule.
	return
}

func (v *parser_) parseInversion() (
	inversion ast.InversionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Inverse rule.
	var inverse ast.InverseLike
	inverse, token, ok = v.parseInverse()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Inverse rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Inversion", token)
		panic(message)
	}

	// Attempt to parse a single Numerical rule.
	var numerical ast.NumericalLike
	numerical, token, ok = v.parseNumerical()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Numerical rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Inversion", token)
		panic(message)
	}

	// Found a single Inversion rule.
	ok = true
	v.remove(tokens)
	inversion = ast.InversionClass().Inversion(
		inverse,
		numerical,
	)
	return
}

func (v *parser_) parseInvocation() (
	invocation ast.InvocationLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Function Invocation.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Invocation.
		invocation = ast.InvocationClass().Invocation(function)
		return
	}

	// Attempt to parse a single Method Invocation.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Invocation.
		invocation = ast.InvocationClass().Invocation(method)
		return
	}

	// This is not a single Invocation rule.
	return
}

func (v *parser_) parseInvoke() (
	invoke ast.InvokeLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "<-" delimiter.
	delimiter, token, ok = v.parseDelimiter("<-")
	if ok {
		// Found a single "<-" delimiter.
		invoke = ast.InvokeClass().Invoke(delimiter)
		return
	}

	// Attempt to parse a single "<~" delimiter.
	delimiter, token, ok = v.parseDelimiter("<~")
	if ok {
		// Found a single "<~" delimiter.
		invoke = ast.InvokeClass().Invoke(delimiter)
		return
	}

	// This is not a single Invoke rule.
	return
}

func (v *parser_) parseItems() (
	items ast.ItemsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Items rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Items", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Entity rules.
	var entities = fra.List[ast.EntityLike]()
entitiesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var entity ast.EntityLike
		entity, token, ok = v.parseEntity()
		if !ok {
			switch {
			case count_ >= 0:
				break entitiesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Entity rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Items", token)
				message += "0 or more Entity rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		entities.AppendValue(entity)
	}

	// Attempt to parse a single "]" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Items rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Items", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Items rule.
	ok = true
	v.remove(tokens)
	items = ast.ItemsClass().Items(
		delimiter1,
		entities,
		delimiter2,
	)
	return
}

func (v *parser_) parseKet() (
	ket ast.KetLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "]" delimiter.
	delimiter, token, ok = v.parseDelimiter("]")
	if ok {
		// Found a single "]" delimiter.
		ket = ast.KetClass().Ket(delimiter)
		return
	}

	// Attempt to parse a single ")" delimiter.
	delimiter, token, ok = v.parseDelimiter(")")
	if ok {
		// Found a single ")" delimiter.
		ket = ast.KetClass().Ket(delimiter)
		return
	}

	// This is not a single Ket rule.
	return
}

func (v *parser_) parseLetClause() (
	letClause ast.LetClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "let" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("let")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single LetClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$LetClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Recipient rule.
	var recipient ast.RecipientLike
	recipient, token, ok = v.parseRecipient()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Recipient rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$LetClause", token)
		panic(message)
	}

	// Attempt to parse a single Assignment rule.
	var assignment ast.AssignmentLike
	assignment, token, ok = v.parseAssignment()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Assignment rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$LetClause", token)
		panic(message)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$LetClause", token)
		panic(message)
	}

	// Found a single LetClause rule.
	ok = true
	v.remove(tokens)
	letClause = ast.LetClauseClass().LetClause(
		delimiter,
		recipient,
		assignment,
		expression,
	)
	return
}

func (v *parser_) parseLexicalOperator() (
	lexicalOperator ast.LexicalOperatorLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "&" delimiter.
	delimiter, token, ok = v.parseDelimiter("&")
	if ok {
		// Found a single "&" delimiter.
		lexicalOperator = ast.LexicalOperatorClass().LexicalOperator(delimiter)
		return
	}

	// This is not a single LexicalOperator rule.
	return
}

func (v *parser_) parseLine() (
	line ast.LineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Annotation Line.
	var annotation ast.AnnotationLike
	annotation, token, ok = v.parseAnnotation()
	if ok {
		// Found a single Annotation Line.
		line = ast.LineClass().Line(annotation)
		return
	}

	// Attempt to parse a single Statement Line.
	var statement ast.StatementLike
	statement, token, ok = v.parseStatement()
	if ok {
		// Found a single Statement Line.
		line = ast.LineClass().Line(statement)
		return
	}

	// This is not a single Line rule.
	return
}

func (v *parser_) parseLogical() (
	logical ast.LogicalLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Document Logical.
	var document ast.DocumentLike
	document, token, ok = v.parseDocument()
	if ok {
		// Found a single Document Logical.
		logical = ast.LogicalClass().Logical(document)
		return
	}

	// Attempt to parse a single Subcomponent Logical.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Logical.
		logical = ast.LogicalClass().Logical(subcomponent)
		return
	}

	// Attempt to parse a single Precedence Logical.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Logical.
		logical = ast.LogicalClass().Logical(precedence)
		return
	}

	// Attempt to parse a single Referent Logical.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Logical.
		logical = ast.LogicalClass().Logical(referent)
		return
	}

	// Attempt to parse a single Complement Logical.
	var complement ast.ComplementLike
	complement, token, ok = v.parseComplement()
	if ok {
		// Found a single Complement Logical.
		logical = ast.LogicalClass().Logical(complement)
		return
	}

	// Attempt to parse a single Function Logical.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Logical.
		logical = ast.LogicalClass().Logical(function)
		return
	}

	// Attempt to parse a single Method Logical.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Logical.
		logical = ast.LogicalClass().Logical(method)
		return
	}

	// Attempt to parse a single Value Logical.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	if ok {
		// Found a single Value Logical.
		logical = ast.LogicalClass().Logical(value)
		return
	}

	// This is not a single Logical rule.
	return
}

func (v *parser_) parseLogicalOperator() (
	logicalOperator ast.LogicalOperatorLike,
	token TokenLike,
	ok bool,
) {
	var delimiter string

	// Attempt to parse a single "and" delimiter.
	delimiter, token, ok = v.parseDelimiter("and")
	if ok {
		// Found a single "and" delimiter.
		logicalOperator = ast.LogicalOperatorClass().LogicalOperator(delimiter)
		return
	}

	// Attempt to parse a single "san" delimiter.
	delimiter, token, ok = v.parseDelimiter("san")
	if ok {
		// Found a single "san" delimiter.
		logicalOperator = ast.LogicalOperatorClass().LogicalOperator(delimiter)
		return
	}

	// Attempt to parse a single "ior" delimiter.
	delimiter, token, ok = v.parseDelimiter("ior")
	if ok {
		// Found a single "ior" delimiter.
		logicalOperator = ast.LogicalOperatorClass().LogicalOperator(delimiter)
		return
	}

	// Attempt to parse a single "xor" delimiter.
	delimiter, token, ok = v.parseDelimiter("xor")
	if ok {
		// Found a single "xor" delimiter.
		logicalOperator = ast.LogicalOperatorClass().LogicalOperator(delimiter)
		return
	}

	// This is not a single LogicalOperator rule.
	return
}

func (v *parser_) parseMagnitude() (
	magnitude ast.MagnitudeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "|" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("|")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Magnitude rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Magnitude", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Magnitude", token)
		panic(message)
	}

	// Attempt to parse a single "|" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("|")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Magnitude rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Magnitude", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Magnitude rule.
	ok = true
	v.remove(tokens)
	magnitude = ast.MagnitudeClass().Magnitude(
		delimiter1,
		expression,
		delimiter2,
	)
	return
}

func (v *parser_) parseMainClause() (
	mainClause ast.MainClauseLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single FlowControl MainClause.
	var flowControl ast.FlowControlLike
	flowControl, token, ok = v.parseFlowControl()
	if ok {
		// Found a single FlowControl MainClause.
		mainClause = ast.MainClauseClass().MainClause(flowControl)
		return
	}

	// Attempt to parse a single ActionInduction MainClause.
	var actionInduction ast.ActionInductionLike
	actionInduction, token, ok = v.parseActionInduction()
	if ok {
		// Found a single ActionInduction MainClause.
		mainClause = ast.MainClauseClass().MainClause(actionInduction)
		return
	}

	// Attempt to parse a single MessageHandling MainClause.
	var messageHandling ast.MessageHandlingLike
	messageHandling, token, ok = v.parseMessageHandling()
	if ok {
		// Found a single MessageHandling MainClause.
		mainClause = ast.MainClauseClass().MainClause(messageHandling)
		return
	}

	// Attempt to parse a single RepositoryAccess MainClause.
	var repositoryAccess ast.RepositoryAccessLike
	repositoryAccess, token, ok = v.parseRepositoryAccess()
	if ok {
		// Found a single RepositoryAccess MainClause.
		mainClause = ast.MainClauseClass().MainClause(repositoryAccess)
		return
	}

	// This is not a single MainClause rule.
	return
}

func (v *parser_) parseMatchingClause() (
	matchingClause ast.MatchingClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "matching" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("matching")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MatchingClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MatchingClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Template rule.
	var template ast.TemplateLike
	template, token, ok = v.parseTemplate()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Template rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$MatchingClause", token)
		panic(message)
	}

	// Attempt to parse a single "do" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MatchingClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MatchingClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Procedure rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$MatchingClause", token)
		panic(message)
	}

	// Found a single MatchingClause rule.
	ok = true
	v.remove(tokens)
	matchingClause = ast.MatchingClauseClass().MatchingClause(
		delimiter1,
		template,
		delimiter2,
		procedure,
	)
	return
}

func (v *parser_) parseMessage() (
	message ast.MessageLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Message", token)
		panic(message)
	}

	// Found a single Message rule.
	ok = true
	v.remove(tokens)
	message = ast.MessageClass().Message(expression)
	return
}

func (v *parser_) parseMessageHandling() (
	messageHandling ast.MessageHandlingLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single PostClause MessageHandling.
	var postClause ast.PostClauseLike
	postClause, token, ok = v.parsePostClause()
	if ok {
		// Found a single PostClause MessageHandling.
		messageHandling = ast.MessageHandlingClass().MessageHandling(postClause)
		return
	}

	// Attempt to parse a single RetrieveClause MessageHandling.
	var retrieveClause ast.RetrieveClauseLike
	retrieveClause, token, ok = v.parseRetrieveClause()
	if ok {
		// Found a single RetrieveClause MessageHandling.
		messageHandling = ast.MessageHandlingClass().MessageHandling(retrieveClause)
		return
	}

	// Attempt to parse a single AcceptClause MessageHandling.
	var acceptClause ast.AcceptClauseLike
	acceptClause, token, ok = v.parseAcceptClause()
	if ok {
		// Found a single AcceptClause MessageHandling.
		messageHandling = ast.MessageHandlingClass().MessageHandling(acceptClause)
		return
	}

	// Attempt to parse a single RejectClause MessageHandling.
	var rejectClause ast.RejectClauseLike
	rejectClause, token, ok = v.parseRejectClause()
	if ok {
		// Found a single RejectClause MessageHandling.
		messageHandling = ast.MessageHandlingClass().MessageHandling(rejectClause)
		return
	}

	// Attempt to parse a single PublishClause MessageHandling.
	var publishClause ast.PublishClauseLike
	publishClause, token, ok = v.parsePublishClause()
	if ok {
		// Found a single PublishClause MessageHandling.
		messageHandling = ast.MessageHandlingClass().MessageHandling(publishClause)
		return
	}

	// This is not a single MessageHandling rule.
	return
}

func (v *parser_) parseMethod() (
	method ast.MethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier1 string
	identifier1, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Invoke rule.
	var invoke ast.InvokeLike
	invoke, token, ok = v.parseInvoke()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Invoke rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Method", token)
		panic(message)
	}

	// Attempt to parse a single identifier token.
	var identifier2 string
	identifier2, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Method rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Argument rules.
	var arguments = fra.List[ast.ArgumentLike]()
argumentsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var argument ast.ArgumentLike
		argument, token, ok = v.parseArgument()
		if !ok {
			switch {
			case count_ >= 0:
				break argumentsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Argument rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Method", token)
				message += "0 or more Argument rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		arguments.AppendValue(argument)
	}

	// Attempt to parse a single ")" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Method rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Method rule.
	ok = true
	v.remove(tokens)
	method = ast.MethodClass().Method(
		identifier1,
		invoke,
		identifier2,
		delimiter1,
		arguments,
		delimiter2,
	)
	return
}

func (v *parser_) parseNotarizeClause() (
	notarizeClause ast.NotarizeClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "notarize" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("notarize")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NotarizeClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NotarizeClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Draft rule.
	var draft ast.DraftLike
	draft, token, ok = v.parseDraft()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Draft rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$NotarizeClause", token)
		panic(message)
	}

	// Attempt to parse a single "as" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("as")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NotarizeClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NotarizeClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Cited rule.
	var cited ast.CitedLike
	cited, token, ok = v.parseCited()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Cited rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$NotarizeClause", token)
		panic(message)
	}

	// Found a single NotarizeClause rule.
	ok = true
	v.remove(tokens)
	notarizeClause = ast.NotarizeClauseClass().NotarizeClause(
		delimiter1,
		draft,
		delimiter2,
		cited,
	)
	return
}

func (v *parser_) parseNumerical() (
	numerical ast.NumericalLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Document Numerical.
	var document ast.DocumentLike
	document, token, ok = v.parseDocument()
	if ok {
		// Found a single Document Numerical.
		numerical = ast.NumericalClass().Numerical(document)
		return
	}

	// Attempt to parse a single Subcomponent Numerical.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Numerical.
		numerical = ast.NumericalClass().Numerical(subcomponent)
		return
	}

	// Attempt to parse a single Precedence Numerical.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Numerical.
		numerical = ast.NumericalClass().Numerical(precedence)
		return
	}

	// Attempt to parse a single Referent Numerical.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Numerical.
		numerical = ast.NumericalClass().Numerical(referent)
		return
	}

	// Attempt to parse a single Inversion Numerical.
	var inversion ast.InversionLike
	inversion, token, ok = v.parseInversion()
	if ok {
		// Found a single Inversion Numerical.
		numerical = ast.NumericalClass().Numerical(inversion)
		return
	}

	// Attempt to parse a single Magnitude Numerical.
	var magnitude ast.MagnitudeLike
	magnitude, token, ok = v.parseMagnitude()
	if ok {
		// Found a single Magnitude Numerical.
		numerical = ast.NumericalClass().Numerical(magnitude)
		return
	}

	// Attempt to parse a single Function Numerical.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Numerical.
		numerical = ast.NumericalClass().Numerical(function)
		return
	}

	// Attempt to parse a single Method Numerical.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Numerical.
		numerical = ast.NumericalClass().Numerical(method)
		return
	}

	// Attempt to parse a single Value Numerical.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	if ok {
		// Found a single Value Numerical.
		numerical = ast.NumericalClass().Numerical(value)
		return
	}

	// This is not a single Numerical rule.
	return
}

func (v *parser_) parseOnClause() (
	onClause ast.OnClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "on" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("on")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single OnClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$OnClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Failure rule.
	var failure ast.FailureLike
	failure, token, ok = v.parseFailure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Failure rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$OnClause", token)
		panic(message)
	}

	// Attempt to parse multiple MatchingClause rules.
	var matchingClauses = fra.List[ast.MatchingClauseLike]()
matchingClausesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var matchingClause ast.MatchingClauseLike
		matchingClause, token, ok = v.parseMatchingClause()
		if !ok {
			switch {
			case count_ >= 1:
				break matchingClausesLoop
			case uti.IsDefined(tokens):
				// This is not multiple MatchingClause rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$OnClause", token)
				message += "1 or more MatchingClause rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		matchingClauses.AppendValue(matchingClause)
	}

	// Found a single OnClause rule.
	ok = true
	v.remove(tokens)
	onClause = ast.OnClauseClass().OnClause(
		delimiter,
		failure,
		matchingClauses,
	)
	return
}

func (v *parser_) parseOperation() (
	operation ast.OperationLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single LexicalOperator Operation.
	var lexicalOperator ast.LexicalOperatorLike
	lexicalOperator, token, ok = v.parseLexicalOperator()
	if ok {
		// Found a single LexicalOperator Operation.
		operation = ast.OperationClass().Operation(lexicalOperator)
		return
	}

	// Attempt to parse a single LogicalOperator Operation.
	var logicalOperator ast.LogicalOperatorLike
	logicalOperator, token, ok = v.parseLogicalOperator()
	if ok {
		// Found a single LogicalOperator Operation.
		operation = ast.OperationClass().Operation(logicalOperator)
		return
	}

	// Attempt to parse a single ArithmeticOperator Operation.
	var arithmeticOperator ast.ArithmeticOperatorLike
	arithmeticOperator, token, ok = v.parseArithmeticOperator()
	if ok {
		// Found a single ArithmeticOperator Operation.
		operation = ast.OperationClass().Operation(arithmeticOperator)
		return
	}

	// Attempt to parse a single ComparisonOperator Operation.
	var comparisonOperator ast.ComparisonOperatorLike
	comparisonOperator, token, ok = v.parseComparisonOperator()
	if ok {
		// Found a single ComparisonOperator Operation.
		operation = ast.OperationClass().Operation(comparisonOperator)
		return
	}

	// This is not a single Operation rule.
	return
}

func (v *parser_) parseParameters() (
	parameters ast.ParametersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Parameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Association rules.
	var associations = fra.List[ast.AssociationLike]()
associationsLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var association ast.AssociationLike
		association, token, ok = v.parseAssociation()
		if !ok {
			switch {
			case count_ >= 1:
				break associationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Association rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Parameters", token)
				message += "1 or more Association rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		associations.AppendValue(association)
	}

	// Attempt to parse a single ")" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Parameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Parameters rule.
	ok = true
	v.remove(tokens)
	parameters = ast.ParametersClass().Parameters(
		delimiter1,
		associations,
		delimiter2,
	)
	return
}

func (v *parser_) parsePostClause() (
	postClause ast.PostClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "post" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("post")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PostClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PostClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Message rule.
	var message ast.MessageLike
	message, token, ok = v.parseMessage()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Message rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PostClause", token)
		panic(message)
	}

	// Attempt to parse a single "to" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("to")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PostClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PostClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Bag rule.
	var bag ast.BagLike
	bag, token, ok = v.parseBag()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Bag rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PostClause", token)
		panic(message)
	}

	// Found a single PostClause rule.
	ok = true
	v.remove(tokens)
	postClause = ast.PostClauseClass().PostClause(
		delimiter1,
		message,
		delimiter2,
		bag,
	)
	return
}

func (v *parser_) parsePrecedence() (
	precedence ast.PrecedenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Precedence rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Precedence", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Precedence", token)
		panic(message)
	}

	// Attempt to parse a single ")" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Precedence rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Precedence", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Precedence rule.
	ok = true
	v.remove(tokens)
	precedence = ast.PrecedenceClass().Precedence(
		delimiter1,
		expression,
		delimiter2,
	)
	return
}

func (v *parser_) parsePredicate() (
	predicate ast.PredicateLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Operation rule.
	var operation ast.OperationLike
	operation, token, ok = v.parseOperation()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Operation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Predicate", token)
		panic(message)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Predicate", token)
		panic(message)
	}

	// Found a single Predicate rule.
	ok = true
	v.remove(tokens)
	predicate = ast.PredicateClass().Predicate(
		operation,
		expression,
	)
	return
}

func (v *parser_) parsePrimitive() (
	primitive ast.PrimitiveLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Element Primitive.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	if ok {
		// Found a single Element Primitive.
		primitive = ast.PrimitiveClass().Primitive(element)
		return
	}

	// Attempt to parse a single String Primitive.
	var string_ ast.StringLike
	string_, token, ok = v.parseString()
	if ok {
		// Found a single String Primitive.
		primitive = ast.PrimitiveClass().Primitive(string_)
		return
	}

	// This is not a single Primitive rule.
	return
}

func (v *parser_) parseProcedure() (
	procedure ast.ProcedureLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "{" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Procedure rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Procedure", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Line rules.
	var lines = fra.List[ast.LineLike]()
linesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var line ast.LineLike
		line, token, ok = v.parseLine()
		if !ok {
			switch {
			case count_ >= 0:
				break linesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Line rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Procedure", token)
				message += "0 or more Line rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		lines.AppendValue(line)
	}

	// Attempt to parse a single "}" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Procedure rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Procedure", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Procedure rule.
	ok = true
	v.remove(tokens)
	procedure = ast.ProcedureClass().Procedure(
		delimiter1,
		lines,
		delimiter2,
	)
	return
}

func (v *parser_) parsePublishClause() (
	publishClause ast.PublishClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "publish" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("publish")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PublishClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PublishClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Event rule.
	var event ast.EventLike
	event, token, ok = v.parseEvent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Event rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PublishClause", token)
		panic(message)
	}

	// Found a single PublishClause rule.
	ok = true
	v.remove(tokens)
	publishClause = ast.PublishClauseClass().PublishClause(
		delimiter,
		event,
	)
	return
}

func (v *parser_) parseRange() (
	range_ ast.RangeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Bra rule.
	var bra ast.BraLike
	bra, token, ok = v.parseBra()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Bra rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Range", token)
		panic(message)
	}

	// Attempt to parse a single Primitive rule.
	var primitive1 ast.PrimitiveLike
	primitive1, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Primitive rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Range", token)
		panic(message)
	}

	// Attempt to parse a single ".." literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Range rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Range", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Primitive rule.
	var primitive2 ast.PrimitiveLike
	primitive2, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Primitive rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Range", token)
		panic(message)
	}

	// Attempt to parse a single Ket rule.
	var ket ast.KetLike
	ket, token, ok = v.parseKet()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Ket rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Range", token)
		panic(message)
	}

	// Found a single Range rule.
	ok = true
	v.remove(tokens)
	range_ = ast.RangeClass().Range(
		bra,
		primitive1,
		delimiter,
		primitive2,
		ket,
	)
	return
}

func (v *parser_) parseRecipient() (
	recipient ast.RecipientLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Variable Recipient.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Recipient.
		recipient = ast.RecipientClass().Recipient(variable)
		return
	}

	// Attempt to parse a single Subcomponent Recipient.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Recipient.
		recipient = ast.RecipientClass().Recipient(subcomponent)
		return
	}

	// This is not a single Recipient rule.
	return
}

func (v *parser_) parseReferent() (
	referent ast.ReferentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "@" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("@")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Referent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Referent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Indirect rule.
	var indirect ast.IndirectLike
	indirect, token, ok = v.parseIndirect()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Indirect rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Referent", token)
		panic(message)
	}

	// Found a single Referent rule.
	ok = true
	v.remove(tokens)
	referent = ast.ReferentClass().Referent(
		delimiter,
		indirect,
	)
	return
}

func (v *parser_) parseRejectClause() (
	rejectClause ast.RejectClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "reject" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("reject")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RejectClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RejectClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Message rule.
	var message ast.MessageLike
	message, token, ok = v.parseMessage()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Message rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$RejectClause", token)
		panic(message)
	}

	// Found a single RejectClause rule.
	ok = true
	v.remove(tokens)
	rejectClause = ast.RejectClauseClass().RejectClause(
		delimiter,
		message,
	)
	return
}

func (v *parser_) parseRepositoryAccess() (
	repositoryAccess ast.RepositoryAccessLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single CheckoutClause RepositoryAccess.
	var checkoutClause ast.CheckoutClauseLike
	checkoutClause, token, ok = v.parseCheckoutClause()
	if ok {
		// Found a single CheckoutClause RepositoryAccess.
		repositoryAccess = ast.RepositoryAccessClass().RepositoryAccess(checkoutClause)
		return
	}

	// Attempt to parse a single SaveClause RepositoryAccess.
	var saveClause ast.SaveClauseLike
	saveClause, token, ok = v.parseSaveClause()
	if ok {
		// Found a single SaveClause RepositoryAccess.
		repositoryAccess = ast.RepositoryAccessClass().RepositoryAccess(saveClause)
		return
	}

	// Attempt to parse a single DiscardClause RepositoryAccess.
	var discardClause ast.DiscardClauseLike
	discardClause, token, ok = v.parseDiscardClause()
	if ok {
		// Found a single DiscardClause RepositoryAccess.
		repositoryAccess = ast.RepositoryAccessClass().RepositoryAccess(discardClause)
		return
	}

	// Attempt to parse a single NotarizeClause RepositoryAccess.
	var notarizeClause ast.NotarizeClauseLike
	notarizeClause, token, ok = v.parseNotarizeClause()
	if ok {
		// Found a single NotarizeClause RepositoryAccess.
		repositoryAccess = ast.RepositoryAccessClass().RepositoryAccess(notarizeClause)
		return
	}

	// This is not a single RepositoryAccess rule.
	return
}

func (v *parser_) parseResult() (
	result ast.ResultLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Result", token)
		panic(message)
	}

	// Found a single Result rule.
	ok = true
	v.remove(tokens)
	result = ast.ResultClass().Result(expression)
	return
}

func (v *parser_) parseRetrieveClause() (
	retrieveClause ast.RetrieveClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "retrieve" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("retrieve")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RetrieveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RetrieveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Recipient rule.
	var recipient ast.RecipientLike
	recipient, token, ok = v.parseRecipient()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Recipient rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$RetrieveClause", token)
		panic(message)
	}

	// Attempt to parse a single "from" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("from")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RetrieveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RetrieveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Bag rule.
	var bag ast.BagLike
	bag, token, ok = v.parseBag()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Bag rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$RetrieveClause", token)
		panic(message)
	}

	// Found a single RetrieveClause rule.
	ok = true
	v.remove(tokens)
	retrieveClause = ast.RetrieveClauseClass().RetrieveClause(
		delimiter1,
		recipient,
		delimiter2,
		bag,
	)
	return
}

func (v *parser_) parseReturnClause() (
	returnClause ast.ReturnClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "return" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("return")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ReturnClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ReturnClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Result rule.
	var result ast.ResultLike
	result, token, ok = v.parseResult()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Result rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ReturnClause", token)
		panic(message)
	}

	// Found a single ReturnClause rule.
	ok = true
	v.remove(tokens)
	returnClause = ast.ReturnClauseClass().ReturnClause(
		delimiter,
		result,
	)
	return
}

func (v *parser_) parseSaveClause() (
	saveClause ast.SaveClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "save" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("save")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SaveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SaveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Draft rule.
	var draft ast.DraftLike
	draft, token, ok = v.parseDraft()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Draft rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$SaveClause", token)
		panic(message)
	}

	// Attempt to parse a single "as" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("as")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SaveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SaveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Cited rule.
	var cited ast.CitedLike
	cited, token, ok = v.parseCited()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Cited rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$SaveClause", token)
		panic(message)
	}

	// Found a single SaveClause rule.
	ok = true
	v.remove(tokens)
	saveClause = ast.SaveClauseClass().SaveClause(
		delimiter1,
		draft,
		delimiter2,
		cited,
	)
	return
}

func (v *parser_) parseSelectClause() (
	selectClause ast.SelectClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "select" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("select")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SelectClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SelectClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Target rule.
	var target ast.TargetLike
	target, token, ok = v.parseTarget()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Target rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$SelectClause", token)
		panic(message)
	}

	// Attempt to parse multiple MatchingClause rules.
	var matchingClauses = fra.List[ast.MatchingClauseLike]()
matchingClausesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var matchingClause ast.MatchingClauseLike
		matchingClause, token, ok = v.parseMatchingClause()
		if !ok {
			switch {
			case count_ >= 1:
				break matchingClausesLoop
			case uti.IsDefined(tokens):
				// This is not multiple MatchingClause rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$SelectClause", token)
				message += "1 or more MatchingClause rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		matchingClauses.AppendValue(matchingClause)
	}

	// Found a single SelectClause rule.
	ok = true
	v.remove(tokens)
	selectClause = ast.SelectClauseClass().SelectClause(
		delimiter,
		target,
		matchingClauses,
	)
	return
}

func (v *parser_) parseSequence() (
	sequence ast.SequenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Sequence", token)
		panic(message)
	}

	// Found a single Sequence rule.
	ok = true
	v.remove(tokens)
	sequence = ast.SequenceClass().Sequence(expression)
	return
}

func (v *parser_) parseStatement() (
	statement ast.StatementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single MainClause rule.
	var mainClause ast.MainClauseLike
	mainClause, token, ok = v.parseMainClause()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single MainClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Statement", token)
		panic(message)
	}

	// Attempt to parse an optional OnClause rule.
	var optionalOnClause ast.OnClauseLike
	optionalOnClause, _, ok = v.parseOnClause()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Statement rule.
	ok = true
	v.remove(tokens)
	statement = ast.StatementClass().Statement(
		mainClause,
		optionalOnClause,
	)
	return
}

func (v *parser_) parseString() (
	string_ ast.StringLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single binary String.
	var binary string
	binary, token, ok = v.parseToken(BinaryToken)
	if ok {
		// Found a single binary String.
		string_ = ast.StringClass().String(binary)
		return
	}

	// Attempt to parse a single bytecode String.
	var bytecode string
	bytecode, token, ok = v.parseToken(BytecodeToken)
	if ok {
		// Found a single bytecode String.
		string_ = ast.StringClass().String(bytecode)
		return
	}

	// Attempt to parse a single name String.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if ok {
		// Found a single name String.
		string_ = ast.StringClass().String(name)
		return
	}

	// Attempt to parse a single narrative String.
	var narrative string
	narrative, token, ok = v.parseToken(NarrativeToken)
	if ok {
		// Found a single narrative String.
		string_ = ast.StringClass().String(narrative)
		return
	}

	// Attempt to parse a single pattern String.
	var pattern string
	pattern, token, ok = v.parseToken(PatternToken)
	if ok {
		// Found a single pattern String.
		string_ = ast.StringClass().String(pattern)
		return
	}

	// Attempt to parse a single quote String.
	var quote string
	quote, token, ok = v.parseToken(QuoteToken)
	if ok {
		// Found a single quote String.
		string_ = ast.StringClass().String(quote)
		return
	}

	// Attempt to parse a single tag String.
	var tag string
	tag, token, ok = v.parseToken(TagToken)
	if ok {
		// Found a single tag String.
		string_ = ast.StringClass().String(tag)
		return
	}

	// Attempt to parse a single version String.
	var version string
	version, token, ok = v.parseToken(VersionToken)
	if ok {
		// Found a single version String.
		string_ = ast.StringClass().String(version)
		return
	}

	// This is not a single String rule.
	return
}

func (v *parser_) parseSubcomponent() (
	subcomponent ast.SubcomponentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Subcomponent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "[" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Subcomponent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Subcomponent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Index rules.
	var indexes = fra.List[ast.IndexLike]()
indexesLoop:
	for count_ := 0; count_ < mat.MaxInt; count_++ {
		var index ast.IndexLike
		index, token, ok = v.parseIndex()
		if !ok {
			switch {
			case count_ >= 1:
				break indexesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Index rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Subcomponent", token)
				message += "1 or more Index rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		indexes.AppendValue(index)
	}

	// Attempt to parse a single "]" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Subcomponent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Subcomponent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Subcomponent rule.
	ok = true
	v.remove(tokens)
	subcomponent = ast.SubcomponentClass().Subcomponent(
		identifier,
		delimiter1,
		indexes,
		delimiter2,
	)
	return
}

func (v *parser_) parseSubject() (
	subject ast.SubjectLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Document Subject.
	var document ast.DocumentLike
	document, token, ok = v.parseDocument()
	if ok {
		// Found a single Document Subject.
		subject = ast.SubjectClass().Subject(document)
		return
	}

	// Attempt to parse a single Subcomponent Subject.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Subject.
		subject = ast.SubjectClass().Subject(subcomponent)
		return
	}

	// Attempt to parse a single Precedence Subject.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Subject.
		subject = ast.SubjectClass().Subject(precedence)
		return
	}

	// Attempt to parse a single Referent Subject.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Subject.
		subject = ast.SubjectClass().Subject(referent)
		return
	}

	// Attempt to parse a single Complement Subject.
	var complement ast.ComplementLike
	complement, token, ok = v.parseComplement()
	if ok {
		// Found a single Complement Subject.
		subject = ast.SubjectClass().Subject(complement)
		return
	}

	// Attempt to parse a single Inversion Subject.
	var inversion ast.InversionLike
	inversion, token, ok = v.parseInversion()
	if ok {
		// Found a single Inversion Subject.
		subject = ast.SubjectClass().Subject(inversion)
		return
	}

	// Attempt to parse a single Magnitude Subject.
	var magnitude ast.MagnitudeLike
	magnitude, token, ok = v.parseMagnitude()
	if ok {
		// Found a single Magnitude Subject.
		subject = ast.SubjectClass().Subject(magnitude)
		return
	}

	// Attempt to parse a single Function Subject.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Subject.
		subject = ast.SubjectClass().Subject(function)
		return
	}

	// Attempt to parse a single Method Subject.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Subject.
		subject = ast.SubjectClass().Subject(method)
		return
	}

	// Attempt to parse a single Value Subject.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	if ok {
		// Found a single Value Subject.
		subject = ast.SubjectClass().Subject(value)
		return
	}

	// This is not a single Subject rule.
	return
}

func (v *parser_) parseTarget() (
	target ast.TargetLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Function Target.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Target.
		target = ast.TargetClass().Target(function)
		return
	}

	// Attempt to parse a single Method Target.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Target.
		target = ast.TargetClass().Target(method)
		return
	}

	// Attempt to parse a single Subcomponent Target.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Target.
		target = ast.TargetClass().Target(subcomponent)
		return
	}

	// Attempt to parse a single Value Target.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	if ok {
		// Found a single Value Target.
		target = ast.TargetClass().Target(value)
		return
	}

	// This is not a single Target rule.
	return
}

func (v *parser_) parseTemplate() (
	template ast.TemplateLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Template", token)
		panic(message)
	}

	// Found a single Template rule.
	ok = true
	v.remove(tokens)
	template = ast.TemplateClass().Template(expression)
	return
}

func (v *parser_) parseThrowClause() (
	throwClause ast.ThrowClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "throw" literal.
	var delimiter string
	delimiter, token, ok = v.parseDelimiter("throw")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ThrowClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ThrowClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Exception rule.
	var exception ast.ExceptionLike
	exception, token, ok = v.parseException()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Exception rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ThrowClause", token)
		panic(message)
	}

	// Found a single ThrowClause rule.
	ok = true
	v.remove(tokens)
	throwClause = ast.ThrowClauseClass().ThrowClause(
		delimiter,
		exception,
	)
	return
}

func (v *parser_) parseValue() (
	value ast.ValueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Value", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Value rule.
	ok = true
	v.remove(tokens)
	value = ast.ValueClass().Value(identifier)
	return
}

func (v *parser_) parseVariable() (
	variable ast.VariableLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Variable", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Variable rule.
	ok = true
	v.remove(tokens)
	variable = ast.VariableClass().Variable(symbol)
	return
}

func (v *parser_) parseWhileClause() (
	whileClause ast.WhileClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "while" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("while")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WhileClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WhileClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Condition rule.
	var condition ast.ConditionLike
	condition, token, ok = v.parseCondition()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Condition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WhileClause", token)
		panic(message)
	}

	// Attempt to parse a single "do" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WhileClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WhileClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Procedure rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WhileClause", token)
		panic(message)
	}

	// Found a single WhileClause rule.
	ok = true
	v.remove(tokens)
	whileClause = ast.WhileClauseClass().WhileClause(
		delimiter1,
		condition,
		delimiter2,
		procedure,
	)
	return
}

func (v *parser_) parseWithClause() (
	withClause ast.WithClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "with" literal.
	var delimiter1 string
	delimiter1, token, ok = v.parseDelimiter("with")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "each" literal.
	var delimiter2 string
	delimiter2, token, ok = v.parseDelimiter("each")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Variable rule.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Variable rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WithClause", token)
		panic(message)
	}

	// Attempt to parse a single "in" literal.
	var delimiter3 string
	delimiter3, token, ok = v.parseDelimiter("in")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Sequence rule.
	var sequence ast.SequenceLike
	sequence, token, ok = v.parseSequence()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Sequence rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WithClause", token)
		panic(message)
	}

	// Attempt to parse a single "do" literal.
	var delimiter4 string
	delimiter4, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Procedure rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WithClause", token)
		panic(message)
	}

	// Found a single WithClause rule.
	ok = true
	v.remove(tokens)
	withClause = ast.WithClauseClass().WithClause(
		delimiter1,
		delimiter2,
		variable,
		delimiter3,
		sequence,
		delimiter4,
		procedure,
	)
	return
}

func (v *parser_) parseDelimiter(
	literal string,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == literal {
			// Found the desired delimiter.
			return
		}
		v.next_.AddValue(token)
		ok = false
	}

	// This is not the desired delimiter.
	return
}

func (v *parser_) parseToken(
	tokenType TokenType,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = fra.List[TokenLike]()
	token = v.getNextToken()
	for token != nil {
		tokens.AppendValue(token)
		switch token.GetType() {
		case tokenType:
			// Found the desired token type.
			value = token.GetValue()
			ok = true
			return
		case SpaceToken, NewlineToken:
			// Ignore any unrequested whitespace.
			token = v.getNextToken()
		default:
			// This is not the desired token type.
			v.putBack(tokens)
			return
		}
	}

	// We are at the end-of-file marker.
	return
}

func (v *parser_) formatError(
	ruleName string,
	token TokenLike,
) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		ScannerClass().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source lines with the error in it.
	message += "\033[36m"
	for index := line - 3; index < line; index++ {
		if index > 1 {
			message += fmt.Sprintf("%04d: ", index) + string(lines[index-1]) + "\n"
		}
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if uti.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(
	ruleName string,
) string {
	var syntax = parserClass().syntax_
	var definition = syntax.GetValue(ruleName)
	return definition
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveLast()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveFirst() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError("", token)
		panic(message)
	}

	return token
}

func (v *parser_) putBack(
	tokens fra.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

// NOTE:
// This method does nothing but must exist to satisfy the lint check on the
// generated parser code.  The generated code must call this method is some
// cases to make it look that the tokens variable is being used somewhere.
func (v *parser_) remove(
	tokens fra.Sequential[TokenLike],
) {
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ fra.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   fra.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ fra.CatalogLike[string, string]
}

// Class Reference

func parserClass() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: fra.CatalogFromMap[string, string](
		map[string]string{
			"$Document": `Component Parameters? note?`,
			"$Component": `
    Element
    String
    Range
    Collection
    Procedure`,
			"$Parameters":  `"(" Association+ ")"`,
			"$Association": `Primitive ":" Document`,
			"$Primitive": `
    Element
    String`,
			"$Element": `
    angle
    boolean
    duration
    glyph
    moment
    number
    percentage
    probability
    resource
    symbol`,
			"$String": `
    binary
    bytecode
    name
    narrative
    pattern
    quote
    tag
    version`,
			"$Range": `Bra Primitive ".." Primitive Ket`,
			"$Bra": `
    "["
    "("`,
			"$Ket": `
    "]"
    ")"`,
			"$Collection": `
    Attributes
    Items  ! Must be after range and attributes.`,
			"$Attributes": `"[" Association+ "]"`,
			"$Items":      `"[" Entity* "]"`,
			"$Entity":     `Document`,
			"$Procedure":  `"{" Line* "}"`,
			"$Line": `
    Annotation
    Statement`,
			"$Annotation": `
    note
    comment`,
			"$Statement": `MainClause OnClause?`,
			"$MainClause": `
    FlowControl
    ActionInduction
    MessageHandling
    RepositoryAccess`,
			"$OnClause":       `"on" Failure MatchingClause+`,
			"$Failure":        `symbol`,
			"$MatchingClause": `"matching" Template "do" Procedure`,
			"$Template":       `Expression`,
			"$FlowControl": `
    IfClause
    SelectClause
    WhileClause
    WithClause
    ContinueClause
    BreakClause
    ReturnClause
    ThrowClause`,
			"$ActionInduction": `
    DoClause
    LetClause`,
			"$MessageHandling": `
    PostClause
    RetrieveClause
    AcceptClause
    RejectClause
    PublishClause`,
			"$RepositoryAccess": `
    CheckoutClause
    SaveClause
    DiscardClause
    NotarizeClause`,
			"$IfClause":     `"if" Condition "do" Procedure`,
			"$Condition":    `Expression`,
			"$SelectClause": `"select" Target MatchingClause+`,
			"$Target": `
    Function
    Method
    Subcomponent
    Value  ! This must be last since others also begin with an identifier.`,
			"$Function": `identifier "(" Argument* ")"`,
			"$Method":   `identifier Invoke identifier "(" Argument* ")"`,
			"$Invoke": `
    "<-"
    "<~"`,
			"$Argument": `
    Value
    Primitive`,
			"$Value":        `identifier`,
			"$Subcomponent": `identifier "[" Index+ "]"`,
			"$Index": `
    Value
    Primitive`,
			"$WhileClause":    `"while" Condition "do" Procedure`,
			"$WithClause":     `"with" "each" Variable "in" Sequence "do" Procedure`,
			"$Variable":       `symbol`,
			"$Sequence":       `Expression`,
			"$ContinueClause": `"continue" "loop"`,
			"$BreakClause":    `"break" "loop"`,
			"$ReturnClause":   `"return" Result`,
			"$Result":         `Expression`,
			"$ThrowClause":    `"throw" Exception`,
			"$Exception":      `Expression`,
			"$DoClause":       `"do" Invocation`,
			"$Invocation": `
    Function
    Method`,
			"$LetClause": `"let" Recipient Assignment Expression`,
			"$Assignment": `
    ":="
    "?="
    "+="
    "-="
    "*="
    "/="`,
			"$Recipient": `
    Variable
    Subcomponent`,
			"$PostClause":     `"post" Message "to" Bag`,
			"$Message":        `Expression`,
			"$Bag":            `Expression`,
			"$RetrieveClause": `"retrieve" Recipient "from" Bag`,
			"$AcceptClause":   `"accept" Message`,
			"$RejectClause":   `"reject" Message`,
			"$PublishClause":  `"publish" Event`,
			"$Event":          `Expression`,
			"$CheckoutClause": `"checkout" Recipient AtLevel? "from" Cited`,
			"$AtLevel":        `"at" "level" Expression`,
			"$Cited":          `Expression`,
			"$SaveClause":     `"save" Draft "as" Cited`,
			"$Draft":          `Expression`,
			"$DiscardClause":  `"discard" Draft`,
			"$NotarizeClause": `"notarize" Draft "as" Cited`,
			"$Expression":     `Subject Predicate*`,
			"$Subject": `
    Document
    Subcomponent
    Precedence
    Referent
    Complement
    Inversion
    Magnitude
    Function
    Method
    Value  ! This must be last since others also begin with an identifier.`,
			"$Predicate": `Operation Expression`,
			"$Operation": `
    LexicalOperator
    LogicalOperator
    ArithmeticOperator
    ComparisonOperator`,
			"$LexicalOperator": `
    "&"`,
			"$LogicalOperator": `
    "and"
    "san"
    "ior"
    "xor"`,
			"$ArithmeticOperator": `
    "+"
    "-"
    "*"
    "/"
    "%"
    "^"`,
			"$ComparisonOperator": `
    "<"
    "="
    ">"
    "is"
    "matches"`,
			"$Precedence": `"(" Expression ")"`,
			"$Referent":   `"@" Indirect`,
			"$Indirect": `
    Document
    Subcomponent
    Referent
    Function
    Method
    Value  ! This must be last since others also begin with an identifier.`,
			"$Complement": `"not" Logical`,
			"$Logical": `
    Document
    Subcomponent
    Precedence
    Referent
    Complement
    Function
    Method
    Value  ! This must be last since others also begin with an identifier.`,
			"$Inversion": `Inverse Numerical`,
			"$Inverse": `
    "-"
    "/"
    "*"`,
			"$Numerical": `
    Document
    Subcomponent
    Precedence
    Referent
    Inversion
    Magnitude
    Function
    Method
    Value  ! This must be last since others also begin with an identifier.`,
			"$Magnitude": `"|" Expression "|"`,
		},
	),
}
