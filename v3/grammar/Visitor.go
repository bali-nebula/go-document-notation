/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	ast "github.com/bali-nebula/go-document-notation/v3/ast"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClass()
}

// Constructor Methods

func (c *visitorClass_) Visitor(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		processor_: processor,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClass()
}

func (v *visitor_) VisitDocument(
	document ast.DocumentLike,
) {
	v.processor_.PreprocessDocument(
		document,
		0,
		0,
	)
	v.visitDocument(document)
	v.processor_.PostprocessDocument(
		document,
		0,
		0,
	)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAcceptClause(
	acceptClause ast.AcceptClauseLike,
) {
	var delimiter1 = acceptClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessAcceptClauseSlot(
		acceptClause,
		1,
	)

	var message = acceptClause.GetMessage()
	v.processor_.PreprocessMessage(
		message,
		0,
		0,
	)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(
		message,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessAcceptClauseSlot(
		acceptClause,
		2,
	)

	var delimiter2 = acceptClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessAcceptClauseSlot(
		acceptClause,
		3,
	)

	var bag = acceptClause.GetBag()
	v.processor_.PreprocessBag(
		bag,
		0,
		0,
	)
	v.visitBag(bag)
	v.processor_.PostprocessBag(
		bag,
		0,
		0,
	)
}

func (v *visitor_) visitArgument(
	argument ast.ArgumentLike,
) {
	// Visit the possible argument rule types.
	switch actual := argument.GetAny().(type) {
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			0,
			0,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			0,
			0,
		)
	case ast.EntityLike:
		v.processor_.PreprocessEntity(
			actual,
			0,
			0,
		)
		v.visitEntity(actual)
		v.processor_.PostprocessEntity(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitArithmetic(
	arithmetic ast.ArithmeticLike,
) {
	// Visit the possible arithmetic literal values.
	var actual = arithmetic.GetAny().(string)
	switch actual {
	case "+":
		v.processor_.ProcessDelimiter("+")
	case "-":
		v.processor_.ProcessDelimiter("-")
	case "*":
		v.processor_.ProcessDelimiter("*")
	case "/":
		v.processor_.ProcessDelimiter("/")
	case "%":
		v.processor_.ProcessDelimiter("%")
	case "^":
		v.processor_.ProcessDelimiter("^")
	}
}

func (v *visitor_) visitAssignClause(
	assignClause ast.AssignClauseLike,
) {
	var delimiter = assignClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAssignClauseSlot(
		assignClause,
		1,
	)

	var recipient = assignClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessAssignClauseSlot(
		assignClause,
		2,
	)

	var assignment = assignClause.GetAssignment()
	v.processor_.PreprocessAssignment(
		assignment,
		0,
		0,
	)
	v.visitAssignment(assignment)
	v.processor_.PostprocessAssignment(
		assignment,
		0,
		0,
	)
	// Visit slot 3 between terms.
	v.processor_.ProcessAssignClauseSlot(
		assignClause,
		3,
	)

	var expression = assignClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitAssignment(
	assignment ast.AssignmentLike,
) {
	// Visit the possible assignment literal values.
	var actual = assignment.GetAny().(string)
	switch actual {
	case "?=":
		v.processor_.ProcessDelimiter("?=")
	case ":=":
		v.processor_.ProcessDelimiter(":=")
	case "+=":
		v.processor_.ProcessDelimiter("+=")
	case "-=":
		v.processor_.ProcessDelimiter("-=")
	case "*=":
		v.processor_.ProcessDelimiter("*=")
	case "/=":
		v.processor_.ProcessDelimiter("/=")
	case "&=":
		v.processor_.ProcessDelimiter("&=")
	}
}

func (v *visitor_) visitAssociation(
	association ast.AssociationLike,
) {
	var primitive = association.GetPrimitive()
	v.processor_.PreprocessPrimitive(
		primitive,
		0,
		0,
	)
	v.visitPrimitive(primitive)
	v.processor_.PostprocessPrimitive(
		primitive,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessAssociationSlot(
		association,
		1,
	)

	var delimiter = association.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 2 between terms.
	v.processor_.ProcessAssociationSlot(
		association,
		2,
	)

	var component = association.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		0,
		0,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		0,
		0,
	)
}

func (v *visitor_) visitAtLevel(
	atLevel ast.AtLevelLike,
) {
	var delimiter1 = atLevel.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessAtLevelSlot(
		atLevel,
		1,
	)

	var delimiter2 = atLevel.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 2 between terms.
	v.processor_.ProcessAtLevelSlot(
		atLevel,
		2,
	)

	var expression = atLevel.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitAttributes(
	attributes ast.AttributesLike,
) {
	var delimiter1 = attributes.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessAttributesSlot(
		attributes,
		1,
	)

	var associationsIndex uint
	var associations = attributes.GetAssociations().GetIterator()
	var associationsCount = uint(associations.GetSize())
	for associations.HasNext() {
		associationsIndex++
		var rule = associations.GetNext()
		v.processor_.PreprocessAssociation(
			rule,
			associationsIndex,
			associationsCount,
		)
		v.visitAssociation(rule)
		v.processor_.PostprocessAssociation(
			rule,
			associationsIndex,
			associationsCount,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessAttributesSlot(
		attributes,
		2,
	)

	var delimiter2 = attributes.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitBag(
	bag ast.BagLike,
) {
	var expression = bag.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitBreakClause(
	breakClause ast.BreakClauseLike,
) {
	var delimiter1 = breakClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessBreakClauseSlot(
		breakClause,
		1,
	)

	var delimiter2 = breakClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
) {
	var delimiter1 = checkoutClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessCheckoutClauseSlot(
		checkoutClause,
		1,
	)

	var recipient = checkoutClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessCheckoutClauseSlot(
		checkoutClause,
		2,
	)

	var optionalAtLevel = checkoutClause.GetOptionalAtLevel()
	if uti.IsDefined(optionalAtLevel) {
		v.processor_.PreprocessAtLevel(
			optionalAtLevel,
			0,
			0,
		)
		v.visitAtLevel(optionalAtLevel)
		v.processor_.PostprocessAtLevel(
			optionalAtLevel,
			0,
			0,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessCheckoutClauseSlot(
		checkoutClause,
		3,
	)

	var delimiter2 = checkoutClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 4 between terms.
	v.processor_.ProcessCheckoutClauseSlot(
		checkoutClause,
		4,
	)

	var location = checkoutClause.GetLocation()
	v.processor_.PreprocessLocation(
		location,
		0,
		0,
	)
	v.visitLocation(location)
	v.processor_.PostprocessLocation(
		location,
		0,
		0,
	)
}

func (v *visitor_) visitCitation(
	citation ast.CitationLike,
) {
	var expression = citation.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitCollection(
	collection ast.CollectionLike,
) {
	// Visit the possible collection rule types.
	switch actual := collection.GetAny().(type) {
	case ast.EmptyLike:
		v.processor_.PreprocessEmpty(
			actual,
			0,
			0,
		)
		v.visitEmpty(actual)
		v.processor_.PostprocessEmpty(
			actual,
			0,
			0,
		)
	case ast.AttributesLike:
		v.processor_.PreprocessAttributes(
			actual,
			0,
			0,
		)
		v.visitAttributes(actual)
		v.processor_.PostprocessAttributes(
			actual,
			0,
			0,
		)
	case ast.ItemsLike:
		v.processor_.PreprocessItems(
			actual,
			0,
			0,
		)
		v.visitItems(actual)
		v.processor_.PostprocessItems(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitComparison(
	comparison ast.ComparisonLike,
) {
	// Visit the possible comparison literal values.
	var actual = comparison.GetAny().(string)
	switch actual {
	case "<":
		v.processor_.ProcessDelimiter("<")
	case "=":
		v.processor_.ProcessDelimiter("=")
	case ">":
		v.processor_.ProcessDelimiter(">")
	case "is":
		v.processor_.ProcessDelimiter("is")
	case "matches":
		v.processor_.ProcessDelimiter("matches")
	}
}

func (v *visitor_) visitComplement(
	complement ast.ComplementLike,
) {
	var delimiter = complement.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessComplementSlot(
		complement,
		1,
	)

	var reversible = complement.GetReversible()
	v.processor_.PreprocessReversible(
		reversible,
		0,
		0,
	)
	v.visitReversible(reversible)
	v.processor_.PostprocessReversible(
		reversible,
		0,
		0,
	)
}

func (v *visitor_) visitComponent(
	component ast.ComponentLike,
) {
	var literal = component.GetLiteral()
	v.processor_.PreprocessLiteral(
		literal,
		0,
		0,
	)
	v.visitLiteral(literal)
	v.processor_.PostprocessLiteral(
		literal,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessComponentSlot(
		component,
		1,
	)

	var optionalGenerics = component.GetOptionalGenerics()
	if uti.IsDefined(optionalGenerics) {
		v.processor_.PreprocessGenerics(
			optionalGenerics,
			0,
			0,
		)
		v.visitGenerics(optionalGenerics)
		v.processor_.PostprocessGenerics(
			optionalGenerics,
			0,
			0,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessComponentSlot(
		component,
		2,
	)

	var optionalNote = component.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitConstant(
	constant ast.ConstantLike,
) {
	var symbol = constant.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitConstraint(
	constraint ast.ConstraintLike,
) {
	var entity = constraint.GetEntity()
	v.processor_.PreprocessEntity(
		entity,
		0,
		0,
	)
	v.visitEntity(entity)
	v.processor_.PostprocessEntity(
		entity,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessConstraintSlot(
		constraint,
		1,
	)

	var optionalGenerics = constraint.GetOptionalGenerics()
	if uti.IsDefined(optionalGenerics) {
		v.processor_.PreprocessGenerics(
			optionalGenerics,
			0,
			0,
		)
		v.visitGenerics(optionalGenerics)
		v.processor_.PostprocessGenerics(
			optionalGenerics,
			0,
			0,
		)
	}
}

func (v *visitor_) visitContinueClause(
	continueClause ast.ContinueClauseLike,
) {
	var delimiter1 = continueClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessContinueClauseSlot(
		continueClause,
		1,
	)

	var delimiter2 = continueClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitDefineClause(
	defineClause ast.DefineClauseLike,
) {
	var delimiter1 = defineClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessDefineClauseSlot(
		defineClause,
		1,
	)

	var constant = defineClause.GetConstant()
	v.processor_.PreprocessConstant(
		constant,
		0,
		0,
	)
	v.visitConstant(constant)
	v.processor_.PostprocessConstant(
		constant,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessDefineClauseSlot(
		defineClause,
		2,
	)

	var delimiter2 = defineClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessDefineClauseSlot(
		defineClause,
		3,
	)

	var expression = defineClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitDiscardClause(
	discardClause ast.DiscardClauseLike,
) {
	var delimiter = discardClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessDiscardClauseSlot(
		discardClause,
		1,
	)

	var citation = discardClause.GetCitation()
	v.processor_.PreprocessCitation(
		citation,
		0,
		0,
	)
	v.visitCitation(citation)
	v.processor_.PostprocessCitation(
		citation,
		0,
		0,
	)
}

func (v *visitor_) visitDocument(
	document ast.DocumentLike,
) {
	var optionalComment = document.GetOptionalComment()
	if uti.IsDefined(optionalComment) {
		v.processor_.ProcessComment(optionalComment)
	}
	// Visit slot 1 between terms.
	v.processor_.ProcessDocumentSlot(
		document,
		1,
	)

	var component = document.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		0,
		0,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		0,
		0,
	)
}

func (v *visitor_) visitDraft(
	draft ast.DraftLike,
) {
	var expression = draft.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitElement(
	element ast.ElementLike,
) {
	// Visit the possible element expression types.
	var actual = element.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, AngleToken):
		v.processor_.ProcessAngle(actual)
	case ScannerClass().MatchesType(actual, BooleanToken):
		v.processor_.ProcessBoolean(actual)
	case ScannerClass().MatchesType(actual, DurationToken):
		v.processor_.ProcessDuration(actual)
	case ScannerClass().MatchesType(actual, GlyphToken):
		v.processor_.ProcessGlyph(actual)
	case ScannerClass().MatchesType(actual, MomentToken):
		v.processor_.ProcessMoment(actual)
	case ScannerClass().MatchesType(actual, PercentageToken):
		v.processor_.ProcessPercentage(actual)
	case ScannerClass().MatchesType(actual, NumberToken):
		v.processor_.ProcessNumber(actual)
	case ScannerClass().MatchesType(actual, ProbabilityToken):
		v.processor_.ProcessProbability(actual)
	case ScannerClass().MatchesType(actual, ResourceToken):
		v.processor_.ProcessResource(actual)
	}
}

func (v *visitor_) visitEmpty(
	empty ast.EmptyLike,
) {
	var delimiter1 = empty.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessEmptySlot(
		empty,
		1,
	)

	var optionalDelimiter = empty.GetOptionalDelimiter()
	if uti.IsDefined(optionalDelimiter) {
		v.processor_.ProcessDelimiter(optionalDelimiter)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessEmptySlot(
		empty,
		2,
	)

	var delimiter2 = empty.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitEntity(
	entity ast.EntityLike,
) {
	// Visit the possible entity rule types.
	switch actual := entity.GetAny().(type) {
	case ast.ElementLike:
		v.processor_.PreprocessElement(
			actual,
			0,
			0,
		)
		v.visitElement(actual)
		v.processor_.PostprocessElement(
			actual,
			0,
			0,
		)
	case ast.SequenceLike:
		v.processor_.PreprocessSequence(
			actual,
			0,
			0,
		)
		v.visitSequence(actual)
		v.processor_.PostprocessSequence(
			actual,
			0,
			0,
		)
	case ast.RangeLike:
		v.processor_.PreprocessRange(
			actual,
			0,
			0,
		)
		v.visitRange(actual)
		v.processor_.PostprocessRange(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitExpression(
	expression ast.ExpressionLike,
) {
	var subject = expression.GetSubject()
	v.processor_.PreprocessSubject(
		subject,
		0,
		0,
	)
	v.visitSubject(subject)
	v.processor_.PostprocessSubject(
		subject,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessExpressionSlot(
		expression,
		1,
	)

	var predicatesIndex uint
	var predicates = expression.GetPredicates().GetIterator()
	var predicatesCount = uint(predicates.GetSize())
	for predicates.HasNext() {
		predicatesIndex++
		var rule = predicates.GetNext()
		v.processor_.PreprocessPredicate(
			rule,
			predicatesIndex,
			predicatesCount,
		)
		v.visitPredicate(rule)
		v.processor_.PostprocessPredicate(
			rule,
			predicatesIndex,
			predicatesCount,
		)
	}
}

func (v *visitor_) visitFlowControl(
	flowControl ast.FlowControlLike,
) {
	// Visit the possible flowControl rule types.
	switch actual := flowControl.GetAny().(type) {
	case ast.IfClauseLike:
		v.processor_.PreprocessIfClause(
			actual,
			0,
			0,
		)
		v.visitIfClause(actual)
		v.processor_.PostprocessIfClause(
			actual,
			0,
			0,
		)
	case ast.SelectClauseLike:
		v.processor_.PreprocessSelectClause(
			actual,
			0,
			0,
		)
		v.visitSelectClause(actual)
		v.processor_.PostprocessSelectClause(
			actual,
			0,
			0,
		)
	case ast.WhileClauseLike:
		v.processor_.PreprocessWhileClause(
			actual,
			0,
			0,
		)
		v.visitWhileClause(actual)
		v.processor_.PostprocessWhileClause(
			actual,
			0,
			0,
		)
	case ast.WithClauseLike:
		v.processor_.PreprocessWithClause(
			actual,
			0,
			0,
		)
		v.visitWithClause(actual)
		v.processor_.PostprocessWithClause(
			actual,
			0,
			0,
		)
	case ast.ContinueClauseLike:
		v.processor_.PreprocessContinueClause(
			actual,
			0,
			0,
		)
		v.visitContinueClause(actual)
		v.processor_.PostprocessContinueClause(
			actual,
			0,
			0,
		)
	case ast.BreakClauseLike:
		v.processor_.PreprocessBreakClause(
			actual,
			0,
			0,
		)
		v.visitBreakClause(actual)
		v.processor_.PostprocessBreakClause(
			actual,
			0,
			0,
		)
	case ast.ReturnClauseLike:
		v.processor_.PreprocessReturnClause(
			actual,
			0,
			0,
		)
		v.visitReturnClause(actual)
		v.processor_.PostprocessReturnClause(
			actual,
			0,
			0,
		)
	case ast.ThrowClauseLike:
		v.processor_.PreprocessThrowClause(
			actual,
			0,
			0,
		)
		v.visitThrowClause(actual)
		v.processor_.PostprocessThrowClause(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitFunction(
	function ast.FunctionLike,
) {
	var identifier = function.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
	// Visit slot 1 between terms.
	v.processor_.ProcessFunctionSlot(
		function,
		1,
	)

	var delimiter1 = function.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessFunctionSlot(
		function,
		2,
	)

	var argumentsIndex uint
	var arguments = function.GetArguments().GetIterator()
	var argumentsCount = uint(arguments.GetSize())
	for arguments.HasNext() {
		argumentsIndex++
		var rule = arguments.GetNext()
		v.processor_.PreprocessArgument(
			rule,
			argumentsIndex,
			argumentsCount,
		)
		v.visitArgument(rule)
		v.processor_.PostprocessArgument(
			rule,
			argumentsIndex,
			argumentsCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessFunctionSlot(
		function,
		3,
	)

	var delimiter2 = function.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitGenerics(
	generics ast.GenericsLike,
) {
	var delimiter1 = generics.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessGenericsSlot(
		generics,
		1,
	)

	var parametersIndex uint
	var parameters = generics.GetParameters().GetIterator()
	var parametersCount = uint(parameters.GetSize())
	for parameters.HasNext() {
		parametersIndex++
		var rule = parameters.GetNext()
		v.processor_.PreprocessParameter(
			rule,
			parametersIndex,
			parametersCount,
		)
		v.visitParameter(rule)
		v.processor_.PostprocessParameter(
			rule,
			parametersIndex,
			parametersCount,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessGenericsSlot(
		generics,
		2,
	)

	var delimiter2 = generics.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitIfClause(
	ifClause ast.IfClauseLike,
) {
	var delimiter1 = ifClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessIfClauseSlot(
		ifClause,
		1,
	)

	var expression = ifClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessIfClauseSlot(
		ifClause,
		2,
	)

	var delimiter2 = ifClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessIfClauseSlot(
		ifClause,
		3,
	)

	var procedure = ifClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		0,
		0,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		0,
		0,
	)
}

func (v *visitor_) visitIndex(
	index ast.IndexLike,
) {
	// Visit the possible index rule types.
	switch actual := index.GetAny().(type) {
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			0,
			0,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			0,
			0,
		)
	case ast.EntityLike:
		v.processor_.PreprocessEntity(
			actual,
			0,
			0,
		)
		v.visitEntity(actual)
		v.processor_.PostprocessEntity(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitInspectClause(
	inspectClause ast.InspectClauseLike,
) {
	var delimiter1 = inspectClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessInspectClauseSlot(
		inspectClause,
		1,
	)

	var recipient = inspectClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessInspectClauseSlot(
		inspectClause,
		2,
	)

	var delimiter2 = inspectClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessInspectClauseSlot(
		inspectClause,
		3,
	)

	var location = inspectClause.GetLocation()
	v.processor_.PreprocessLocation(
		location,
		0,
		0,
	)
	v.visitLocation(location)
	v.processor_.PostprocessLocation(
		location,
		0,
		0,
	)
}

func (v *visitor_) visitInverse(
	inverse ast.InverseLike,
) {
	// Visit the possible inverse literal values.
	var actual = inverse.GetAny().(string)
	switch actual {
	case "-":
		v.processor_.ProcessDelimiter("-")
	case "/":
		v.processor_.ProcessDelimiter("/")
	case "*":
		v.processor_.ProcessDelimiter("*")
	}
}

func (v *visitor_) visitInversion(
	inversion ast.InversionLike,
) {
	var inverse = inversion.GetInverse()
	v.processor_.PreprocessInverse(
		inverse,
		0,
		0,
	)
	v.visitInverse(inverse)
	v.processor_.PostprocessInverse(
		inverse,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessInversionSlot(
		inversion,
		1,
	)

	var numerical = inversion.GetNumerical()
	v.processor_.PreprocessNumerical(
		numerical,
		0,
		0,
	)
	v.visitNumerical(numerical)
	v.processor_.PostprocessNumerical(
		numerical,
		0,
		0,
	)
}

func (v *visitor_) visitInvocation(
	invocation ast.InvocationLike,
) {
	// Visit the possible invocation literal values.
	var actual = invocation.GetAny().(string)
	switch actual {
	case "<-":
		v.processor_.ProcessDelimiter("<-")
	case "<~":
		v.processor_.ProcessDelimiter("<~")
	}
}

func (v *visitor_) visitInvokeClause(
	invokeClause ast.InvokeClauseLike,
) {
	var delimiter = invokeClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessInvokeClauseSlot(
		invokeClause,
		1,
	)

	var method = invokeClause.GetMethod()
	v.processor_.PreprocessMethod(
		method,
		0,
		0,
	)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(
		method,
		0,
		0,
	)
}

func (v *visitor_) visitItems(
	items ast.ItemsLike,
) {
	var delimiter1 = items.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessItemsSlot(
		items,
		1,
	)

	var componentsIndex uint
	var components = items.GetComponents().GetIterator()
	var componentsCount = uint(components.GetSize())
	for components.HasNext() {
		componentsIndex++
		var rule = components.GetNext()
		v.processor_.PreprocessComponent(
			rule,
			componentsIndex,
			componentsCount,
		)
		v.visitComponent(rule)
		v.processor_.PostprocessComponent(
			rule,
			componentsIndex,
			componentsCount,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessItemsSlot(
		items,
		2,
	)

	var delimiter2 = items.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitLeft(
	left ast.LeftLike,
) {
	// Visit the possible left literal values.
	var actual = left.GetAny().(string)
	switch actual {
	case "[":
		v.processor_.ProcessDelimiter("[")
	case "(":
		v.processor_.ProcessDelimiter("(")
	}
}

func (v *visitor_) visitLexical(
	lexical ast.LexicalLike,
) {
	// Visit the possible lexical literal values.
	var actual = lexical.GetAny().(string)
	switch actual {
	case "&":
		v.processor_.ProcessDelimiter("&")
	}
}

func (v *visitor_) visitLiteral(
	literal ast.LiteralLike,
) {
	// Visit the possible literal rule types.
	switch actual := literal.GetAny().(type) {
	case ast.ElementLike:
		v.processor_.PreprocessElement(
			actual,
			0,
			0,
		)
		v.visitElement(actual)
		v.processor_.PostprocessElement(
			actual,
			0,
			0,
		)
	case ast.SequenceLike:
		v.processor_.PreprocessSequence(
			actual,
			0,
			0,
		)
		v.visitSequence(actual)
		v.processor_.PostprocessSequence(
			actual,
			0,
			0,
		)
	case ast.RangeLike:
		v.processor_.PreprocessRange(
			actual,
			0,
			0,
		)
		v.visitRange(actual)
		v.processor_.PostprocessRange(
			actual,
			0,
			0,
		)
	case ast.CollectionLike:
		v.processor_.PreprocessCollection(
			actual,
			0,
			0,
		)
		v.visitCollection(actual)
		v.processor_.PostprocessCollection(
			actual,
			0,
			0,
		)
	case ast.ProcedureLike:
		v.processor_.PreprocessProcedure(
			actual,
			0,
			0,
		)
		v.visitProcedure(actual)
		v.processor_.PostprocessProcedure(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitLocalTransformation(
	localTransformation ast.LocalTransformationLike,
) {
	// Visit the possible localTransformation rule types.
	switch actual := localTransformation.GetAny().(type) {
	case ast.DefineClauseLike:
		v.processor_.PreprocessDefineClause(
			actual,
			0,
			0,
		)
		v.visitDefineClause(actual)
		v.processor_.PostprocessDefineClause(
			actual,
			0,
			0,
		)
	case ast.AssignClauseLike:
		v.processor_.PreprocessAssignClause(
			actual,
			0,
			0,
		)
		v.visitAssignClause(actual)
		v.processor_.PostprocessAssignClause(
			actual,
			0,
			0,
		)
	case ast.InvokeClauseLike:
		v.processor_.PreprocessInvokeClause(
			actual,
			0,
			0,
		)
		v.visitInvokeClause(actual)
		v.processor_.PostprocessInvokeClause(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitLocation(
	location ast.LocationLike,
) {
	var expression = location.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitLogical(
	logical ast.LogicalLike,
) {
	// Visit the possible logical literal values.
	var actual = logical.GetAny().(string)
	switch actual {
	case "and":
		v.processor_.ProcessDelimiter("and")
	case "san":
		v.processor_.ProcessDelimiter("san")
	case "ior":
		v.processor_.ProcessDelimiter("ior")
	case "xor":
		v.processor_.ProcessDelimiter("xor")
	}
}

func (v *visitor_) visitMagnitude(
	magnitude ast.MagnitudeLike,
) {
	var delimiter1 = magnitude.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMagnitudeSlot(
		magnitude,
		1,
	)

	var expression = magnitude.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessMagnitudeSlot(
		magnitude,
		2,
	)

	var delimiter2 = magnitude.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitMainClause(
	mainClause ast.MainClauseLike,
) {
	// Visit the possible mainClause rule types.
	switch actual := mainClause.GetAny().(type) {
	case ast.FlowControlLike:
		v.processor_.PreprocessFlowControl(
			actual,
			0,
			0,
		)
		v.visitFlowControl(actual)
		v.processor_.PostprocessFlowControl(
			actual,
			0,
			0,
		)
	case ast.LocalTransformationLike:
		v.processor_.PreprocessLocalTransformation(
			actual,
			0,
			0,
		)
		v.visitLocalTransformation(actual)
		v.processor_.PostprocessLocalTransformation(
			actual,
			0,
			0,
		)
	case ast.MessageHandlingLike:
		v.processor_.PreprocessMessageHandling(
			actual,
			0,
			0,
		)
		v.visitMessageHandling(actual)
		v.processor_.PostprocessMessageHandling(
			actual,
			0,
			0,
		)
	case ast.RepositoryAccessLike:
		v.processor_.PreprocessRepositoryAccess(
			actual,
			0,
			0,
		)
		v.visitRepositoryAccess(actual)
		v.processor_.PostprocessRepositoryAccess(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitMatchingClause(
	matchingClause ast.MatchingClauseLike,
) {
	var delimiter1 = matchingClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMatchingClauseSlot(
		matchingClause,
		1,
	)

	var expression = matchingClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessMatchingClauseSlot(
		matchingClause,
		2,
	)

	var delimiter2 = matchingClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessMatchingClauseSlot(
		matchingClause,
		3,
	)

	var procedure = matchingClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		0,
		0,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		0,
		0,
	)
}

func (v *visitor_) visitMessage(
	message ast.MessageLike,
) {
	var expression = message.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitMessageHandling(
	messageHandling ast.MessageHandlingLike,
) {
	// Visit the possible messageHandling rule types.
	switch actual := messageHandling.GetAny().(type) {
	case ast.SendClauseLike:
		v.processor_.PreprocessSendClause(
			actual,
			0,
			0,
		)
		v.visitSendClause(actual)
		v.processor_.PostprocessSendClause(
			actual,
			0,
			0,
		)
	case ast.ReceiveClauseLike:
		v.processor_.PreprocessReceiveClause(
			actual,
			0,
			0,
		)
		v.visitReceiveClause(actual)
		v.processor_.PostprocessReceiveClause(
			actual,
			0,
			0,
		)
	case ast.AcceptClauseLike:
		v.processor_.PreprocessAcceptClause(
			actual,
			0,
			0,
		)
		v.visitAcceptClause(actual)
		v.processor_.PostprocessAcceptClause(
			actual,
			0,
			0,
		)
	case ast.RejectClauseLike:
		v.processor_.PreprocessRejectClause(
			actual,
			0,
			0,
		)
		v.visitRejectClause(actual)
		v.processor_.PostprocessRejectClause(
			actual,
			0,
			0,
		)
	case ast.PublishClauseLike:
		v.processor_.PreprocessPublishClause(
			actual,
			0,
			0,
		)
		v.visitPublishClause(actual)
		v.processor_.PostprocessPublishClause(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitMethod(
	method ast.MethodLike,
) {
	var identifier1 = method.GetIdentifier1()
	v.processor_.ProcessIdentifier(identifier1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		1,
	)

	var invocation = method.GetInvocation()
	v.processor_.PreprocessInvocation(
		invocation,
		0,
		0,
	)
	v.visitInvocation(invocation)
	v.processor_.PostprocessInvocation(
		invocation,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		2,
	)

	var identifier2 = method.GetIdentifier2()
	v.processor_.ProcessIdentifier(identifier2)
	// Visit slot 3 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		3,
	)

	var delimiter1 = method.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 4 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		4,
	)

	var argumentsIndex uint
	var arguments = method.GetArguments().GetIterator()
	var argumentsCount = uint(arguments.GetSize())
	for arguments.HasNext() {
		argumentsIndex++
		var rule = arguments.GetNext()
		v.processor_.PreprocessArgument(
			rule,
			argumentsIndex,
			argumentsCount,
		)
		v.visitArgument(rule)
		v.processor_.PostprocessArgument(
			rule,
			argumentsIndex,
			argumentsCount,
		)
	}
	// Visit slot 5 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		5,
	)

	var delimiter2 = method.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitNotarizeClause(
	notarizeClause ast.NotarizeClauseLike,
) {
	var delimiter1 = notarizeClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessNotarizeClauseSlot(
		notarizeClause,
		1,
	)

	var draft = notarizeClause.GetDraft()
	v.processor_.PreprocessDraft(
		draft,
		0,
		0,
	)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(
		draft,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessNotarizeClauseSlot(
		notarizeClause,
		2,
	)

	var delimiter2 = notarizeClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessNotarizeClauseSlot(
		notarizeClause,
		3,
	)

	var location = notarizeClause.GetLocation()
	v.processor_.PreprocessLocation(
		location,
		0,
		0,
	)
	v.visitLocation(location)
	v.processor_.PostprocessLocation(
		location,
		0,
		0,
	)
}

func (v *visitor_) visitNumerical(
	numerical ast.NumericalLike,
) {
	// Visit the possible numerical rule types.
	switch actual := numerical.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(
			actual,
			0,
			0,
		)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(
			actual,
			0,
			0,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			0,
			0,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			0,
			0,
		)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			0,
			0,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			0,
			0,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			0,
			0,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			0,
			0,
		)
	case ast.InversionLike:
		v.processor_.PreprocessInversion(
			actual,
			0,
			0,
		)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(
			actual,
			0,
			0,
		)
	case ast.MagnitudeLike:
		v.processor_.PreprocessMagnitude(
			actual,
			0,
			0,
		)
		v.visitMagnitude(actual)
		v.processor_.PostprocessMagnitude(
			actual,
			0,
			0,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			0,
			0,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			0,
			0,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			0,
			0,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			0,
			0,
		)
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			0,
			0,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitOnClause(
	onClause ast.OnClauseLike,
) {
	var delimiter = onClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessOnClauseSlot(
		onClause,
		1,
	)

	var symbol = onClause.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
	// Visit slot 2 between terms.
	v.processor_.ProcessOnClauseSlot(
		onClause,
		2,
	)

	var matchingClausesIndex uint
	var matchingClauses = onClause.GetMatchingClauses().GetIterator()
	var matchingClausesCount = uint(matchingClauses.GetSize())
	for matchingClauses.HasNext() {
		matchingClausesIndex++
		var rule = matchingClauses.GetNext()
		v.processor_.PreprocessMatchingClause(
			rule,
			matchingClausesIndex,
			matchingClausesCount,
		)
		v.visitMatchingClause(rule)
		v.processor_.PostprocessMatchingClause(
			rule,
			matchingClausesIndex,
			matchingClausesCount,
		)
	}
}

func (v *visitor_) visitOperator(
	operator ast.OperatorLike,
) {
	// Visit the possible operator rule types.
	switch actual := operator.GetAny().(type) {
	case ast.ComparisonLike:
		v.processor_.PreprocessComparison(
			actual,
			0,
			0,
		)
		v.visitComparison(actual)
		v.processor_.PostprocessComparison(
			actual,
			0,
			0,
		)
	case ast.LogicalLike:
		v.processor_.PreprocessLogical(
			actual,
			0,
			0,
		)
		v.visitLogical(actual)
		v.processor_.PostprocessLogical(
			actual,
			0,
			0,
		)
	case ast.ArithmeticLike:
		v.processor_.PreprocessArithmetic(
			actual,
			0,
			0,
		)
		v.visitArithmetic(actual)
		v.processor_.PostprocessArithmetic(
			actual,
			0,
			0,
		)
	case ast.LexicalLike:
		v.processor_.PreprocessLexical(
			actual,
			0,
			0,
		)
		v.visitLexical(actual)
		v.processor_.PostprocessLexical(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitParameter(
	parameter ast.ParameterLike,
) {
	var symbol = parameter.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
	// Visit slot 1 between terms.
	v.processor_.ProcessParameterSlot(
		parameter,
		1,
	)

	var delimiter = parameter.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 2 between terms.
	v.processor_.ProcessParameterSlot(
		parameter,
		2,
	)

	var constraint = parameter.GetConstraint()
	v.processor_.PreprocessConstraint(
		constraint,
		0,
		0,
	)
	v.visitConstraint(constraint)
	v.processor_.PostprocessConstraint(
		constraint,
		0,
		0,
	)
}

func (v *visitor_) visitPrecedence(
	precedence ast.PrecedenceLike,
) {
	var delimiter1 = precedence.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessPrecedenceSlot(
		precedence,
		1,
	)

	var expression = precedence.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessPrecedenceSlot(
		precedence,
		2,
	)

	var delimiter2 = precedence.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitPredicate(
	predicate ast.PredicateLike,
) {
	var operator = predicate.GetOperator()
	v.processor_.PreprocessOperator(
		operator,
		0,
		0,
	)
	v.visitOperator(operator)
	v.processor_.PostprocessOperator(
		operator,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessPredicateSlot(
		predicate,
		1,
	)

	var expression = predicate.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitPrimitive(
	primitive ast.PrimitiveLike,
) {
	// Visit the possible primitive rule types.
	switch actual := primitive.GetAny().(type) {
	case ast.ElementLike:
		v.processor_.PreprocessElement(
			actual,
			0,
			0,
		)
		v.visitElement(actual)
		v.processor_.PostprocessElement(
			actual,
			0,
			0,
		)
	case ast.SequenceLike:
		v.processor_.PreprocessSequence(
			actual,
			0,
			0,
		)
		v.visitSequence(actual)
		v.processor_.PostprocessSequence(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitProcedure(
	procedure ast.ProcedureLike,
) {
	var delimiter1 = procedure.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessProcedureSlot(
		procedure,
		1,
	)

	var statementsIndex uint
	var statements = procedure.GetStatements().GetIterator()
	var statementsCount = uint(statements.GetSize())
	for statements.HasNext() {
		statementsIndex++
		var rule = statements.GetNext()
		v.processor_.PreprocessStatement(
			rule,
			statementsIndex,
			statementsCount,
		)
		v.visitStatement(rule)
		v.processor_.PostprocessStatement(
			rule,
			statementsIndex,
			statementsCount,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessProcedureSlot(
		procedure,
		2,
	)

	var delimiter2 = procedure.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitPublishClause(
	publishClause ast.PublishClauseLike,
) {
	var delimiter = publishClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessPublishClauseSlot(
		publishClause,
		1,
	)

	var message = publishClause.GetMessage()
	v.processor_.PreprocessMessage(
		message,
		0,
		0,
	)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(
		message,
		0,
		0,
	)
}

func (v *visitor_) visitRange(
	range_ ast.RangeLike,
) {
	var left = range_.GetLeft()
	v.processor_.PreprocessLeft(
		left,
		0,
		0,
	)
	v.visitLeft(left)
	v.processor_.PostprocessLeft(
		left,
		0,
		0,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		1,
	)

	var optionalPrimitive1 = range_.GetOptionalPrimitive1()
	if uti.IsDefined(optionalPrimitive1) {
		v.processor_.PreprocessPrimitive(
			optionalPrimitive1,
			0,
			0,
		)
		v.visitPrimitive(optionalPrimitive1)
		v.processor_.PostprocessPrimitive(
			optionalPrimitive1,
			0,
			0,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		2,
	)

	var delimiter = range_.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 3 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		3,
	)

	var optionalPrimitive2 = range_.GetOptionalPrimitive2()
	if uti.IsDefined(optionalPrimitive2) {
		v.processor_.PreprocessPrimitive(
			optionalPrimitive2,
			0,
			0,
		)
		v.visitPrimitive(optionalPrimitive2)
		v.processor_.PostprocessPrimitive(
			optionalPrimitive2,
			0,
			0,
		)
	}
	// Visit slot 4 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		4,
	)

	var right = range_.GetRight()
	v.processor_.PreprocessRight(
		right,
		0,
		0,
	)
	v.visitRight(right)
	v.processor_.PostprocessRight(
		right,
		0,
		0,
	)
}

func (v *visitor_) visitReceiveClause(
	receiveClause ast.ReceiveClauseLike,
) {
	var delimiter1 = receiveClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessReceiveClauseSlot(
		receiveClause,
		1,
	)

	var recipient = receiveClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessReceiveClauseSlot(
		receiveClause,
		2,
	)

	var delimiter2 = receiveClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessReceiveClauseSlot(
		receiveClause,
		3,
	)

	var bag = receiveClause.GetBag()
	v.processor_.PreprocessBag(
		bag,
		0,
		0,
	)
	v.visitBag(bag)
	v.processor_.PostprocessBag(
		bag,
		0,
		0,
	)
}

func (v *visitor_) visitRecipient(
	recipient ast.RecipientLike,
) {
	// Visit the possible recipient rule types.
	switch actual := recipient.GetAny().(type) {
	case ast.VariableLike:
		v.processor_.PreprocessVariable(
			actual,
			0,
			0,
		)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(
			actual,
			0,
			0,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			0,
			0,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitReference(
	reference ast.ReferenceLike,
) {
	// Visit the possible reference rule types.
	switch actual := reference.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(
			actual,
			0,
			0,
		)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(
			actual,
			0,
			0,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			0,
			0,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			0,
			0,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			0,
			0,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			0,
			0,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			0,
			0,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			0,
			0,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			0,
			0,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			0,
			0,
		)
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			0,
			0,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitReferent(
	referent ast.ReferentLike,
) {
	var delimiter = referent.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessReferentSlot(
		referent,
		1,
	)

	var reference = referent.GetReference()
	v.processor_.PreprocessReference(
		reference,
		0,
		0,
	)
	v.visitReference(reference)
	v.processor_.PostprocessReference(
		reference,
		0,
		0,
	)
}

func (v *visitor_) visitRejectClause(
	rejectClause ast.RejectClauseLike,
) {
	var delimiter1 = rejectClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessRejectClauseSlot(
		rejectClause,
		1,
	)

	var message = rejectClause.GetMessage()
	v.processor_.PreprocessMessage(
		message,
		0,
		0,
	)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(
		message,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessRejectClauseSlot(
		rejectClause,
		2,
	)

	var delimiter2 = rejectClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessRejectClauseSlot(
		rejectClause,
		3,
	)

	var bag = rejectClause.GetBag()
	v.processor_.PreprocessBag(
		bag,
		0,
		0,
	)
	v.visitBag(bag)
	v.processor_.PostprocessBag(
		bag,
		0,
		0,
	)
}

func (v *visitor_) visitRepositoryAccess(
	repositoryAccess ast.RepositoryAccessLike,
) {
	// Visit the possible repositoryAccess rule types.
	switch actual := repositoryAccess.GetAny().(type) {
	case ast.SaveClauseLike:
		v.processor_.PreprocessSaveClause(
			actual,
			0,
			0,
		)
		v.visitSaveClause(actual)
		v.processor_.PostprocessSaveClause(
			actual,
			0,
			0,
		)
	case ast.RetrieveClauseLike:
		v.processor_.PreprocessRetrieveClause(
			actual,
			0,
			0,
		)
		v.visitRetrieveClause(actual)
		v.processor_.PostprocessRetrieveClause(
			actual,
			0,
			0,
		)
	case ast.DiscardClauseLike:
		v.processor_.PreprocessDiscardClause(
			actual,
			0,
			0,
		)
		v.visitDiscardClause(actual)
		v.processor_.PostprocessDiscardClause(
			actual,
			0,
			0,
		)
	case ast.NotarizeClauseLike:
		v.processor_.PreprocessNotarizeClause(
			actual,
			0,
			0,
		)
		v.visitNotarizeClause(actual)
		v.processor_.PostprocessNotarizeClause(
			actual,
			0,
			0,
		)
	case ast.InspectClauseLike:
		v.processor_.PreprocessInspectClause(
			actual,
			0,
			0,
		)
		v.visitInspectClause(actual)
		v.processor_.PostprocessInspectClause(
			actual,
			0,
			0,
		)
	case ast.CheckoutClauseLike:
		v.processor_.PreprocessCheckoutClause(
			actual,
			0,
			0,
		)
		v.visitCheckoutClause(actual)
		v.processor_.PostprocessCheckoutClause(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitRetrieveClause(
	retrieveClause ast.RetrieveClauseLike,
) {
	var delimiter1 = retrieveClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessRetrieveClauseSlot(
		retrieveClause,
		1,
	)

	var recipient = retrieveClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessRetrieveClauseSlot(
		retrieveClause,
		2,
	)

	var delimiter2 = retrieveClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessRetrieveClauseSlot(
		retrieveClause,
		3,
	)

	var citation = retrieveClause.GetCitation()
	v.processor_.PreprocessCitation(
		citation,
		0,
		0,
	)
	v.visitCitation(citation)
	v.processor_.PostprocessCitation(
		citation,
		0,
		0,
	)
}

func (v *visitor_) visitReturnClause(
	returnClause ast.ReturnClauseLike,
) {
	var delimiter = returnClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessReturnClauseSlot(
		returnClause,
		1,
	)

	var expression = returnClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitReversible(
	reversible ast.ReversibleLike,
) {
	// Visit the possible reversible rule types.
	switch actual := reversible.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(
			actual,
			0,
			0,
		)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(
			actual,
			0,
			0,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			0,
			0,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			0,
			0,
		)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			0,
			0,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			0,
			0,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			0,
			0,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			0,
			0,
		)
	case ast.ComplementLike:
		v.processor_.PreprocessComplement(
			actual,
			0,
			0,
		)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(
			actual,
			0,
			0,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			0,
			0,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			0,
			0,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			0,
			0,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			0,
			0,
		)
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			0,
			0,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitRight(
	right ast.RightLike,
) {
	// Visit the possible right literal values.
	var actual = right.GetAny().(string)
	switch actual {
	case "]":
		v.processor_.ProcessDelimiter("]")
	case ")":
		v.processor_.ProcessDelimiter(")")
	}
}

func (v *visitor_) visitSaveClause(
	saveClause ast.SaveClauseLike,
) {
	var delimiter1 = saveClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessSaveClauseSlot(
		saveClause,
		1,
	)

	var draft = saveClause.GetDraft()
	v.processor_.PreprocessDraft(
		draft,
		0,
		0,
	)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(
		draft,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessSaveClauseSlot(
		saveClause,
		2,
	)

	var delimiter2 = saveClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessSaveClauseSlot(
		saveClause,
		3,
	)

	var recipient = saveClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)
}

func (v *visitor_) visitSelectClause(
	selectClause ast.SelectClauseLike,
) {
	var delimiter = selectClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessSelectClauseSlot(
		selectClause,
		1,
	)

	var expression = selectClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessSelectClauseSlot(
		selectClause,
		2,
	)

	var matchingClausesIndex uint
	var matchingClauses = selectClause.GetMatchingClauses().GetIterator()
	var matchingClausesCount = uint(matchingClauses.GetSize())
	for matchingClauses.HasNext() {
		matchingClausesIndex++
		var rule = matchingClauses.GetNext()
		v.processor_.PreprocessMatchingClause(
			rule,
			matchingClausesIndex,
			matchingClausesCount,
		)
		v.visitMatchingClause(rule)
		v.processor_.PostprocessMatchingClause(
			rule,
			matchingClausesIndex,
			matchingClausesCount,
		)
	}
}

func (v *visitor_) visitSendClause(
	sendClause ast.SendClauseLike,
) {
	var delimiter1 = sendClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessSendClauseSlot(
		sendClause,
		1,
	)

	var message = sendClause.GetMessage()
	v.processor_.PreprocessMessage(
		message,
		0,
		0,
	)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(
		message,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessSendClauseSlot(
		sendClause,
		2,
	)

	var delimiter2 = sendClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessSendClauseSlot(
		sendClause,
		3,
	)

	var bag = sendClause.GetBag()
	v.processor_.PreprocessBag(
		bag,
		0,
		0,
	)
	v.visitBag(bag)
	v.processor_.PostprocessBag(
		bag,
		0,
		0,
	)
}

func (v *visitor_) visitSequence(
	sequence ast.SequenceLike,
) {
	// Visit the possible sequence expression types.
	var actual = sequence.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, BinaryToken):
		v.processor_.ProcessBinary(actual)
	case ScannerClass().MatchesType(actual, BytecodeToken):
		v.processor_.ProcessBytecode(actual)
	case ScannerClass().MatchesType(actual, NameToken):
		v.processor_.ProcessName(actual)
	case ScannerClass().MatchesType(actual, NarrativeToken):
		v.processor_.ProcessNarrative(actual)
	case ScannerClass().MatchesType(actual, PatternToken):
		v.processor_.ProcessPattern(actual)
	case ScannerClass().MatchesType(actual, QuoteToken):
		v.processor_.ProcessQuote(actual)
	case ScannerClass().MatchesType(actual, SymbolToken):
		v.processor_.ProcessSymbol(actual)
	case ScannerClass().MatchesType(actual, TagToken):
		v.processor_.ProcessTag(actual)
	case ScannerClass().MatchesType(actual, VersionToken):
		v.processor_.ProcessVersion(actual)
	}
}

func (v *visitor_) visitStatement(
	statement ast.StatementLike,
) {
	var optionalComment = statement.GetOptionalComment()
	if uti.IsDefined(optionalComment) {
		v.processor_.ProcessComment(optionalComment)
	}
	// Visit slot 1 between terms.
	v.processor_.ProcessStatementSlot(
		statement,
		1,
	)

	var mainClause = statement.GetMainClause()
	v.processor_.PreprocessMainClause(
		mainClause,
		0,
		0,
	)
	v.visitMainClause(mainClause)
	v.processor_.PostprocessMainClause(
		mainClause,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessStatementSlot(
		statement,
		2,
	)

	var optionalOnClause = statement.GetOptionalOnClause()
	if uti.IsDefined(optionalOnClause) {
		v.processor_.PreprocessOnClause(
			optionalOnClause,
			0,
			0,
		)
		v.visitOnClause(optionalOnClause)
		v.processor_.PostprocessOnClause(
			optionalOnClause,
			0,
			0,
		)
	}
}

func (v *visitor_) visitSubcomponent(
	subcomponent ast.SubcomponentLike,
) {
	var identifier = subcomponent.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
	// Visit slot 1 between terms.
	v.processor_.ProcessSubcomponentSlot(
		subcomponent,
		1,
	)

	var delimiter1 = subcomponent.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessSubcomponentSlot(
		subcomponent,
		2,
	)

	var indexesIndex uint
	var indexes = subcomponent.GetIndexes().GetIterator()
	var indexesCount = uint(indexes.GetSize())
	for indexes.HasNext() {
		indexesIndex++
		var rule = indexes.GetNext()
		v.processor_.PreprocessIndex(
			rule,
			indexesIndex,
			indexesCount,
		)
		v.visitIndex(rule)
		v.processor_.PostprocessIndex(
			rule,
			indexesIndex,
			indexesCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessSubcomponentSlot(
		subcomponent,
		3,
	)

	var delimiter2 = subcomponent.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitSubject(
	subject ast.SubjectLike,
) {
	// Visit the possible subject rule types.
	switch actual := subject.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(
			actual,
			0,
			0,
		)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(
			actual,
			0,
			0,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			0,
			0,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			0,
			0,
		)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			0,
			0,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			0,
			0,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			0,
			0,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			0,
			0,
		)
	case ast.ComplementLike:
		v.processor_.PreprocessComplement(
			actual,
			0,
			0,
		)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(
			actual,
			0,
			0,
		)
	case ast.InversionLike:
		v.processor_.PreprocessInversion(
			actual,
			0,
			0,
		)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(
			actual,
			0,
			0,
		)
	case ast.MagnitudeLike:
		v.processor_.PreprocessMagnitude(
			actual,
			0,
			0,
		)
		v.visitMagnitude(actual)
		v.processor_.PostprocessMagnitude(
			actual,
			0,
			0,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			0,
			0,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			0,
			0,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			0,
			0,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			0,
			0,
		)
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			0,
			0,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitThrowClause(
	throwClause ast.ThrowClauseLike,
) {
	var delimiter = throwClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessThrowClauseSlot(
		throwClause,
		1,
	)

	var expression = throwClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitValue(
	value ast.ValueLike,
) {
	var identifier = value.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
}

func (v *visitor_) visitVariable(
	variable ast.VariableLike,
) {
	var symbol = variable.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitWhileClause(
	whileClause ast.WhileClauseLike,
) {
	var delimiter1 = whileClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessWhileClauseSlot(
		whileClause,
		1,
	)

	var expression = whileClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessWhileClauseSlot(
		whileClause,
		2,
	)

	var delimiter2 = whileClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessWhileClauseSlot(
		whileClause,
		3,
	)

	var procedure = whileClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		0,
		0,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		0,
		0,
	)
}

func (v *visitor_) visitWithClause(
	withClause ast.WithClauseLike,
) {
	var delimiter1 = withClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		1,
	)

	var delimiter2 = withClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 2 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		2,
	)

	var symbol = withClause.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
	// Visit slot 3 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		3,
	)

	var delimiter3 = withClause.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
	// Visit slot 4 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		4,
	)

	var expression = withClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
	// Visit slot 5 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		5,
	)

	var delimiter4 = withClause.GetDelimiter4()
	v.processor_.ProcessDelimiter(delimiter4)
	// Visit slot 6 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		6,
	)

	var procedure = withClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		0,
		0,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		0,
		0,
	)
}

// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}

// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func visitorClass() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
