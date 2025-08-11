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
	ast "github.com/bali-nebula/go-document-notation/v3/ast"
	uti "github.com/craterdog/go-missing-utilities/v7"
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
		1,
		1,
	)
	v.visitDocument(document)
	v.processor_.PostprocessDocument(
		document,
		1,
		1,
	)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAcceptClause(
	acceptClause ast.AcceptClauseLike,
) {
	var delimiter = acceptClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAcceptClauseSlot(
		acceptClause,
		1,
	)

	var message = acceptClause.GetMessage()
	v.processor_.PreprocessMessage(
		message,
		1,
		1,
	)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(
		message,
		1,
		1,
	)
}

func (v *visitor_) visitActionInduction(
	actionInduction ast.ActionInductionLike,
) {
	// Visit the possible actionInduction rule types.
	switch actual := actionInduction.GetAny().(type) {
	case ast.DoClauseLike:
		v.processor_.PreprocessDoClause(
			actual,
			1,
			1,
		)
		v.visitDoClause(actual)
		v.processor_.PostprocessDoClause(
			actual,
			1,
			1,
		)
	case ast.LetClauseLike:
		v.processor_.PreprocessLetClause(
			actual,
			1,
			1,
		)
		v.visitLetClause(actual)
		v.processor_.PostprocessLetClause(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitAnnotation(
	annotation ast.AnnotationLike,
) {
	// Visit the possible annotation expression types.
	var actual = annotation.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, NoteToken):
		v.processor_.ProcessNote(actual)
	case ScannerClass().MatchesType(actual, CommentToken):
		v.processor_.ProcessComment(actual)
	}
}

func (v *visitor_) visitArgument(
	argument ast.ArgumentLike,
) {
	// Visit the possible argument rule types.
	switch actual := argument.GetAny().(type) {
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			1,
			1,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			1,
			1,
		)
	case ast.PrimitiveLike:
		v.processor_.PreprocessPrimitive(
			actual,
			1,
			1,
		)
		v.visitPrimitive(actual)
		v.processor_.PostprocessPrimitive(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitArithmeticOperator(
	arithmeticOperator ast.ArithmeticOperatorLike,
) {
	// Visit the possible arithmeticOperator literal values.
	var actual = arithmeticOperator.GetAny().(string)
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

func (v *visitor_) visitAssignment(
	assignment ast.AssignmentLike,
) {
	// Visit the possible assignment literal values.
	var actual = assignment.GetAny().(string)
	switch actual {
	case ":=":
		v.processor_.ProcessDelimiter(":=")
	case "?=":
		v.processor_.ProcessDelimiter("?=")
	case "+=":
		v.processor_.ProcessDelimiter("+=")
	case "-=":
		v.processor_.ProcessDelimiter("-=")
	case "*=":
		v.processor_.ProcessDelimiter("*=")
	case "/=":
		v.processor_.ProcessDelimiter("/=")
	}
}

func (v *visitor_) visitAssociation(
	association ast.AssociationLike,
) {
	var primitive = association.GetPrimitive()
	v.processor_.PreprocessPrimitive(
		primitive,
		1,
		1,
	)
	v.visitPrimitive(primitive)
	v.processor_.PostprocessPrimitive(
		primitive,
		1,
		1,
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

	var document = association.GetDocument()
	v.processor_.PreprocessDocument(
		document,
		1,
		1,
	)
	v.visitDocument(document)
	v.processor_.PostprocessDocument(
		document,
		1,
		1,
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
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
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
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
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
		1,
		1,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		1,
		1,
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
			1,
			1,
		)
		v.visitAtLevel(optionalAtLevel)
		v.processor_.PostprocessAtLevel(
			optionalAtLevel,
			1,
			1,
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

	var cited = checkoutClause.GetCited()
	v.processor_.PreprocessCited(
		cited,
		1,
		1,
	)
	v.visitCited(cited)
	v.processor_.PostprocessCited(
		cited,
		1,
		1,
	)
}

func (v *visitor_) visitCited(
	cited ast.CitedLike,
) {
	var expression = cited.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
}

func (v *visitor_) visitCollection(
	collection ast.CollectionLike,
) {
	// Visit the possible collection rule types.
	switch actual := collection.GetAny().(type) {
	case ast.AttributesLike:
		v.processor_.PreprocessAttributes(
			actual,
			1,
			1,
		)
		v.visitAttributes(actual)
		v.processor_.PostprocessAttributes(
			actual,
			1,
			1,
		)
	case ast.EntitiesLike:
		v.processor_.PreprocessEntities(
			actual,
			1,
			1,
		)
		v.visitEntities(actual)
		v.processor_.PostprocessEntities(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitComparisonOperator(
	comparisonOperator ast.ComparisonOperatorLike,
) {
	// Visit the possible comparisonOperator literal values.
	var actual = comparisonOperator.GetAny().(string)
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

	var logical = complement.GetLogical()
	v.processor_.PreprocessLogical(
		logical,
		1,
		1,
	)
	v.visitLogical(logical)
	v.processor_.PostprocessLogical(
		logical,
		1,
		1,
	)
}

func (v *visitor_) visitComponent(
	component ast.ComponentLike,
) {
	// Visit the possible component rule types.
	switch actual := component.GetAny().(type) {
	case ast.ElementLike:
		v.processor_.PreprocessElement(
			actual,
			1,
			1,
		)
		v.visitElement(actual)
		v.processor_.PostprocessElement(
			actual,
			1,
			1,
		)
	case ast.StringLike:
		v.processor_.PreprocessString(
			actual,
			1,
			1,
		)
		v.visitString(actual)
		v.processor_.PostprocessString(
			actual,
			1,
			1,
		)
	case ast.RangeLike:
		v.processor_.PreprocessRange(
			actual,
			1,
			1,
		)
		v.visitRange(actual)
		v.processor_.PostprocessRange(
			actual,
			1,
			1,
		)
	case ast.CollectionLike:
		v.processor_.PreprocessCollection(
			actual,
			1,
			1,
		)
		v.visitCollection(actual)
		v.processor_.PostprocessCollection(
			actual,
			1,
			1,
		)
	case ast.ProcedureLike:
		v.processor_.PreprocessProcedure(
			actual,
			1,
			1,
		)
		v.visitProcedure(actual)
		v.processor_.PostprocessProcedure(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitCondition(
	condition ast.ConditionLike,
) {
	var expression = condition.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
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

	var draft = discardClause.GetDraft()
	v.processor_.PreprocessDraft(
		draft,
		1,
		1,
	)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(
		draft,
		1,
		1,
	)
}

func (v *visitor_) visitDoClause(
	doClause ast.DoClauseLike,
) {
	var delimiter = doClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessDoClauseSlot(
		doClause,
		1,
	)

	var method = doClause.GetMethod()
	v.processor_.PreprocessMethod(
		method,
		1,
		1,
	)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(
		method,
		1,
		1,
	)
}

func (v *visitor_) visitDocument(
	document ast.DocumentLike,
) {
	var component = document.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		1,
		1,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessDocumentSlot(
		document,
		1,
	)

	var optionalParameters = document.GetOptionalParameters()
	if uti.IsDefined(optionalParameters) {
		v.processor_.PreprocessParameters(
			optionalParameters,
			1,
			1,
		)
		v.visitParameters(optionalParameters)
		v.processor_.PostprocessParameters(
			optionalParameters,
			1,
			1,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessDocumentSlot(
		document,
		2,
	)

	var optionalNote = document.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitDraft(
	draft ast.DraftLike,
) {
	var expression = draft.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
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
	case ScannerClass().MatchesType(actual, NumberToken):
		v.processor_.ProcessNumber(actual)
	case ScannerClass().MatchesType(actual, PercentageToken):
		v.processor_.ProcessPercentage(actual)
	case ScannerClass().MatchesType(actual, ProbabilityToken):
		v.processor_.ProcessProbability(actual)
	case ScannerClass().MatchesType(actual, ResourceToken):
		v.processor_.ProcessResource(actual)
	case ScannerClass().MatchesType(actual, SymbolToken):
		v.processor_.ProcessSymbol(actual)
	}
}

func (v *visitor_) visitEntities(
	entities ast.EntitiesLike,
) {
	var delimiter1 = entities.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessEntitiesSlot(
		entities,
		1,
	)

	var itemsIndex uint
	var items = entities.GetItems().GetIterator()
	var itemsCount = uint(items.GetSize())
	for items.HasNext() {
		itemsIndex++
		var rule = items.GetNext()
		v.processor_.PreprocessItem(
			rule,
			itemsIndex,
			itemsCount,
		)
		v.visitItem(rule)
		v.processor_.PostprocessItem(
			rule,
			itemsIndex,
			itemsCount,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessEntitiesSlot(
		entities,
		2,
	)

	var delimiter2 = entities.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitEvent(
	event ast.EventLike,
) {
	var expression = event.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
}

func (v *visitor_) visitException(
	exception ast.ExceptionLike,
) {
	var expression = exception.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
}

func (v *visitor_) visitExpression(
	expression ast.ExpressionLike,
) {
	var subject = expression.GetSubject()
	v.processor_.PreprocessSubject(
		subject,
		1,
		1,
	)
	v.visitSubject(subject)
	v.processor_.PostprocessSubject(
		subject,
		1,
		1,
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

func (v *visitor_) visitFailure(
	failure ast.FailureLike,
) {
	var symbol = failure.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitFlowControl(
	flowControl ast.FlowControlLike,
) {
	// Visit the possible flowControl rule types.
	switch actual := flowControl.GetAny().(type) {
	case ast.IfClauseLike:
		v.processor_.PreprocessIfClause(
			actual,
			1,
			1,
		)
		v.visitIfClause(actual)
		v.processor_.PostprocessIfClause(
			actual,
			1,
			1,
		)
	case ast.SelectClauseLike:
		v.processor_.PreprocessSelectClause(
			actual,
			1,
			1,
		)
		v.visitSelectClause(actual)
		v.processor_.PostprocessSelectClause(
			actual,
			1,
			1,
		)
	case ast.WhileClauseLike:
		v.processor_.PreprocessWhileClause(
			actual,
			1,
			1,
		)
		v.visitWhileClause(actual)
		v.processor_.PostprocessWhileClause(
			actual,
			1,
			1,
		)
	case ast.WithClauseLike:
		v.processor_.PreprocessWithClause(
			actual,
			1,
			1,
		)
		v.visitWithClause(actual)
		v.processor_.PostprocessWithClause(
			actual,
			1,
			1,
		)
	case ast.ContinueClauseLike:
		v.processor_.PreprocessContinueClause(
			actual,
			1,
			1,
		)
		v.visitContinueClause(actual)
		v.processor_.PostprocessContinueClause(
			actual,
			1,
			1,
		)
	case ast.BreakClauseLike:
		v.processor_.PreprocessBreakClause(
			actual,
			1,
			1,
		)
		v.visitBreakClause(actual)
		v.processor_.PostprocessBreakClause(
			actual,
			1,
			1,
		)
	case ast.ReturnClauseLike:
		v.processor_.PreprocessReturnClause(
			actual,
			1,
			1,
		)
		v.visitReturnClause(actual)
		v.processor_.PostprocessReturnClause(
			actual,
			1,
			1,
		)
	case ast.ThrowClauseLike:
		v.processor_.PreprocessThrowClause(
			actual,
			1,
			1,
		)
		v.visitThrowClause(actual)
		v.processor_.PostprocessThrowClause(
			actual,
			1,
			1,
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

	var condition = ifClause.GetCondition()
	v.processor_.PreprocessCondition(
		condition,
		1,
		1,
	)
	v.visitCondition(condition)
	v.processor_.PostprocessCondition(
		condition,
		1,
		1,
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
		1,
		1,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		1,
		1,
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
			1,
			1,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			1,
			1,
		)
	case ast.PrimitiveLike:
		v.processor_.PreprocessPrimitive(
			actual,
			1,
			1,
		)
		v.visitPrimitive(actual)
		v.processor_.PostprocessPrimitive(
			actual,
			1,
			1,
		)
	}
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
		1,
		1,
	)
	v.visitInverse(inverse)
	v.processor_.PostprocessInverse(
		inverse,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessInversionSlot(
		inversion,
		1,
	)

	var numerical = inversion.GetNumerical()
	v.processor_.PreprocessNumerical(
		numerical,
		1,
		1,
	)
	v.visitNumerical(numerical)
	v.processor_.PostprocessNumerical(
		numerical,
		1,
		1,
	)
}

func (v *visitor_) visitInvoke(
	invoke ast.InvokeLike,
) {
	// Visit the possible invoke literal values.
	var actual = invoke.GetAny().(string)
	switch actual {
	case "<-":
		v.processor_.ProcessDelimiter("<-")
	case "<~":
		v.processor_.ProcessDelimiter("<~")
	}
}

func (v *visitor_) visitItem(
	item ast.ItemLike,
) {
	var document = item.GetDocument()
	v.processor_.PreprocessDocument(
		document,
		1,
		1,
	)
	v.visitDocument(document)
	v.processor_.PostprocessDocument(
		document,
		1,
		1,
	)
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

func (v *visitor_) visitLetClause(
	letClause ast.LetClauseLike,
) {
	var delimiter = letClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessLetClauseSlot(
		letClause,
		1,
	)

	var recipient = letClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		1,
		1,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessLetClauseSlot(
		letClause,
		2,
	)

	var assignment = letClause.GetAssignment()
	v.processor_.PreprocessAssignment(
		assignment,
		1,
		1,
	)
	v.visitAssignment(assignment)
	v.processor_.PostprocessAssignment(
		assignment,
		1,
		1,
	)
	// Visit slot 3 between terms.
	v.processor_.ProcessLetClauseSlot(
		letClause,
		3,
	)

	var expression = letClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
}

func (v *visitor_) visitLexicalOperator(
	lexicalOperator ast.LexicalOperatorLike,
) {
	// Visit the possible lexicalOperator literal values.
	var actual = lexicalOperator.GetAny().(string)
	switch actual {
	case "&":
		v.processor_.ProcessDelimiter("&")
	}
}

func (v *visitor_) visitLine(
	line ast.LineLike,
) {
	// Visit the possible line rule types.
	switch actual := line.GetAny().(type) {
	case ast.AnnotationLike:
		v.processor_.PreprocessAnnotation(
			actual,
			1,
			1,
		)
		v.visitAnnotation(actual)
		v.processor_.PostprocessAnnotation(
			actual,
			1,
			1,
		)
	case ast.StatementLike:
		v.processor_.PreprocessStatement(
			actual,
			1,
			1,
		)
		v.visitStatement(actual)
		v.processor_.PostprocessStatement(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitLogical(
	logical ast.LogicalLike,
) {
	// Visit the possible logical rule types.
	switch actual := logical.GetAny().(type) {
	case ast.DocumentLike:
		v.processor_.PreprocessDocument(
			actual,
			1,
			1,
		)
		v.visitDocument(actual)
		v.processor_.PostprocessDocument(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			1,
			1,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			1,
			1,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			1,
			1,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			1,
			1,
		)
	case ast.ComplementLike:
		v.processor_.PreprocessComplement(
			actual,
			1,
			1,
		)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(
			actual,
			1,
			1,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			1,
			1,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			1,
			1,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			1,
			1,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			1,
			1,
		)
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			1,
			1,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitLogicalOperator(
	logicalOperator ast.LogicalOperatorLike,
) {
	// Visit the possible logicalOperator literal values.
	var actual = logicalOperator.GetAny().(string)
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
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
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
			1,
			1,
		)
		v.visitFlowControl(actual)
		v.processor_.PostprocessFlowControl(
			actual,
			1,
			1,
		)
	case ast.ActionInductionLike:
		v.processor_.PreprocessActionInduction(
			actual,
			1,
			1,
		)
		v.visitActionInduction(actual)
		v.processor_.PostprocessActionInduction(
			actual,
			1,
			1,
		)
	case ast.MessageHandlingLike:
		v.processor_.PreprocessMessageHandling(
			actual,
			1,
			1,
		)
		v.visitMessageHandling(actual)
		v.processor_.PostprocessMessageHandling(
			actual,
			1,
			1,
		)
	case ast.RepositoryAccessLike:
		v.processor_.PreprocessRepositoryAccess(
			actual,
			1,
			1,
		)
		v.visitRepositoryAccess(actual)
		v.processor_.PostprocessRepositoryAccess(
			actual,
			1,
			1,
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

	var template = matchingClause.GetTemplate()
	v.processor_.PreprocessTemplate(
		template,
		1,
		1,
	)
	v.visitTemplate(template)
	v.processor_.PostprocessTemplate(
		template,
		1,
		1,
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
		1,
		1,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		1,
		1,
	)
}

func (v *visitor_) visitMessage(
	message ast.MessageLike,
) {
	var expression = message.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
}

func (v *visitor_) visitMessageHandling(
	messageHandling ast.MessageHandlingLike,
) {
	// Visit the possible messageHandling rule types.
	switch actual := messageHandling.GetAny().(type) {
	case ast.PostClauseLike:
		v.processor_.PreprocessPostClause(
			actual,
			1,
			1,
		)
		v.visitPostClause(actual)
		v.processor_.PostprocessPostClause(
			actual,
			1,
			1,
		)
	case ast.RetrieveClauseLike:
		v.processor_.PreprocessRetrieveClause(
			actual,
			1,
			1,
		)
		v.visitRetrieveClause(actual)
		v.processor_.PostprocessRetrieveClause(
			actual,
			1,
			1,
		)
	case ast.AcceptClauseLike:
		v.processor_.PreprocessAcceptClause(
			actual,
			1,
			1,
		)
		v.visitAcceptClause(actual)
		v.processor_.PostprocessAcceptClause(
			actual,
			1,
			1,
		)
	case ast.RejectClauseLike:
		v.processor_.PreprocessRejectClause(
			actual,
			1,
			1,
		)
		v.visitRejectClause(actual)
		v.processor_.PostprocessRejectClause(
			actual,
			1,
			1,
		)
	case ast.PublishClauseLike:
		v.processor_.PreprocessPublishClause(
			actual,
			1,
			1,
		)
		v.visitPublishClause(actual)
		v.processor_.PostprocessPublishClause(
			actual,
			1,
			1,
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

	var invoke = method.GetInvoke()
	v.processor_.PreprocessInvoke(
		invoke,
		1,
		1,
	)
	v.visitInvoke(invoke)
	v.processor_.PostprocessInvoke(
		invoke,
		1,
		1,
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
		1,
		1,
	)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(
		draft,
		1,
		1,
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

	var cited = notarizeClause.GetCited()
	v.processor_.PreprocessCited(
		cited,
		1,
		1,
	)
	v.visitCited(cited)
	v.processor_.PostprocessCited(
		cited,
		1,
		1,
	)
}

func (v *visitor_) visitNumerical(
	numerical ast.NumericalLike,
) {
	// Visit the possible numerical rule types.
	switch actual := numerical.GetAny().(type) {
	case ast.DocumentLike:
		v.processor_.PreprocessDocument(
			actual,
			1,
			1,
		)
		v.visitDocument(actual)
		v.processor_.PostprocessDocument(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			1,
			1,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			1,
			1,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			1,
			1,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			1,
			1,
		)
	case ast.InversionLike:
		v.processor_.PreprocessInversion(
			actual,
			1,
			1,
		)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(
			actual,
			1,
			1,
		)
	case ast.MagnitudeLike:
		v.processor_.PreprocessMagnitude(
			actual,
			1,
			1,
		)
		v.visitMagnitude(actual)
		v.processor_.PostprocessMagnitude(
			actual,
			1,
			1,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			1,
			1,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			1,
			1,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			1,
			1,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			1,
			1,
		)
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			1,
			1,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			1,
			1,
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

	var failure = onClause.GetFailure()
	v.processor_.PreprocessFailure(
		failure,
		1,
		1,
	)
	v.visitFailure(failure)
	v.processor_.PostprocessFailure(
		failure,
		1,
		1,
	)
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
	case ast.LexicalOperatorLike:
		v.processor_.PreprocessLexicalOperator(
			actual,
			1,
			1,
		)
		v.visitLexicalOperator(actual)
		v.processor_.PostprocessLexicalOperator(
			actual,
			1,
			1,
		)
	case ast.LogicalOperatorLike:
		v.processor_.PreprocessLogicalOperator(
			actual,
			1,
			1,
		)
		v.visitLogicalOperator(actual)
		v.processor_.PostprocessLogicalOperator(
			actual,
			1,
			1,
		)
	case ast.ArithmeticOperatorLike:
		v.processor_.PreprocessArithmeticOperator(
			actual,
			1,
			1,
		)
		v.visitArithmeticOperator(actual)
		v.processor_.PostprocessArithmeticOperator(
			actual,
			1,
			1,
		)
	case ast.ComparisonOperatorLike:
		v.processor_.PreprocessComparisonOperator(
			actual,
			1,
			1,
		)
		v.visitComparisonOperator(actual)
		v.processor_.PostprocessComparisonOperator(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitParameters(
	parameters ast.ParametersLike,
) {
	var delimiter1 = parameters.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessParametersSlot(
		parameters,
		1,
	)

	var associationsIndex uint
	var associations = parameters.GetAssociations().GetIterator()
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
	v.processor_.ProcessParametersSlot(
		parameters,
		2,
	)

	var delimiter2 = parameters.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitPostClause(
	postClause ast.PostClauseLike,
) {
	var delimiter1 = postClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessPostClauseSlot(
		postClause,
		1,
	)

	var message = postClause.GetMessage()
	v.processor_.PreprocessMessage(
		message,
		1,
		1,
	)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(
		message,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessPostClauseSlot(
		postClause,
		2,
	)

	var delimiter2 = postClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessPostClauseSlot(
		postClause,
		3,
	)

	var bag = postClause.GetBag()
	v.processor_.PreprocessBag(
		bag,
		1,
		1,
	)
	v.visitBag(bag)
	v.processor_.PostprocessBag(
		bag,
		1,
		1,
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
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
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
		1,
		1,
	)
	v.visitOperator(operator)
	v.processor_.PostprocessOperator(
		operator,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessPredicateSlot(
		predicate,
		1,
	)

	var expression = predicate.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
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
			1,
			1,
		)
		v.visitElement(actual)
		v.processor_.PostprocessElement(
			actual,
			1,
			1,
		)
	case ast.StringLike:
		v.processor_.PreprocessString(
			actual,
			1,
			1,
		)
		v.visitString(actual)
		v.processor_.PostprocessString(
			actual,
			1,
			1,
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

	var linesIndex uint
	var lines = procedure.GetLines().GetIterator()
	var linesCount = uint(lines.GetSize())
	for lines.HasNext() {
		linesIndex++
		var rule = lines.GetNext()
		v.processor_.PreprocessLine(
			rule,
			linesIndex,
			linesCount,
		)
		v.visitLine(rule)
		v.processor_.PostprocessLine(
			rule,
			linesIndex,
			linesCount,
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

	var event = publishClause.GetEvent()
	v.processor_.PreprocessEvent(
		event,
		1,
		1,
	)
	v.visitEvent(event)
	v.processor_.PostprocessEvent(
		event,
		1,
		1,
	)
}

func (v *visitor_) visitRange(
	range_ ast.RangeLike,
) {
	var left = range_.GetLeft()
	v.processor_.PreprocessLeft(
		left,
		1,
		1,
	)
	v.visitLeft(left)
	v.processor_.PostprocessLeft(
		left,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		1,
	)

	var primitive1 = range_.GetPrimitive1()
	v.processor_.PreprocessPrimitive(
		primitive1,
		1,
		1,
	)
	v.visitPrimitive(primitive1)
	v.processor_.PostprocessPrimitive(
		primitive1,
		1,
		1,
	)
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

	var primitive2 = range_.GetPrimitive2()
	v.processor_.PreprocessPrimitive(
		primitive2,
		1,
		1,
	)
	v.visitPrimitive(primitive2)
	v.processor_.PostprocessPrimitive(
		primitive2,
		1,
		1,
	)
	// Visit slot 4 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		4,
	)

	var right = range_.GetRight()
	v.processor_.PreprocessRight(
		right,
		1,
		1,
	)
	v.visitRight(right)
	v.processor_.PostprocessRight(
		right,
		1,
		1,
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
			1,
			1,
		)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitReference(
	reference ast.ReferenceLike,
) {
	// Visit the possible reference rule types.
	switch actual := reference.GetAny().(type) {
	case ast.DocumentLike:
		v.processor_.PreprocessDocument(
			actual,
			1,
			1,
		)
		v.visitDocument(actual)
		v.processor_.PostprocessDocument(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			1,
			1,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			1,
			1,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			1,
			1,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			1,
			1,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			1,
			1,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			1,
			1,
		)
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			1,
			1,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			1,
			1,
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
		1,
		1,
	)
	v.visitReference(reference)
	v.processor_.PostprocessReference(
		reference,
		1,
		1,
	)
}

func (v *visitor_) visitRejectClause(
	rejectClause ast.RejectClauseLike,
) {
	var delimiter = rejectClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessRejectClauseSlot(
		rejectClause,
		1,
	)

	var message = rejectClause.GetMessage()
	v.processor_.PreprocessMessage(
		message,
		1,
		1,
	)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(
		message,
		1,
		1,
	)
}

func (v *visitor_) visitRepositoryAccess(
	repositoryAccess ast.RepositoryAccessLike,
) {
	// Visit the possible repositoryAccess rule types.
	switch actual := repositoryAccess.GetAny().(type) {
	case ast.CheckoutClauseLike:
		v.processor_.PreprocessCheckoutClause(
			actual,
			1,
			1,
		)
		v.visitCheckoutClause(actual)
		v.processor_.PostprocessCheckoutClause(
			actual,
			1,
			1,
		)
	case ast.SaveClauseLike:
		v.processor_.PreprocessSaveClause(
			actual,
			1,
			1,
		)
		v.visitSaveClause(actual)
		v.processor_.PostprocessSaveClause(
			actual,
			1,
			1,
		)
	case ast.DiscardClauseLike:
		v.processor_.PreprocessDiscardClause(
			actual,
			1,
			1,
		)
		v.visitDiscardClause(actual)
		v.processor_.PostprocessDiscardClause(
			actual,
			1,
			1,
		)
	case ast.NotarizeClauseLike:
		v.processor_.PreprocessNotarizeClause(
			actual,
			1,
			1,
		)
		v.visitNotarizeClause(actual)
		v.processor_.PostprocessNotarizeClause(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitResult(
	result ast.ResultLike,
) {
	var expression = result.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
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
		1,
		1,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		1,
		1,
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

	var bag = retrieveClause.GetBag()
	v.processor_.PreprocessBag(
		bag,
		1,
		1,
	)
	v.visitBag(bag)
	v.processor_.PostprocessBag(
		bag,
		1,
		1,
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

	var result = returnClause.GetResult()
	v.processor_.PreprocessResult(
		result,
		1,
		1,
	)
	v.visitResult(result)
	v.processor_.PostprocessResult(
		result,
		1,
		1,
	)
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
		1,
		1,
	)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(
		draft,
		1,
		1,
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

	var cited = saveClause.GetCited()
	v.processor_.PreprocessCited(
		cited,
		1,
		1,
	)
	v.visitCited(cited)
	v.processor_.PostprocessCited(
		cited,
		1,
		1,
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
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
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

func (v *visitor_) visitSequence(
	sequence ast.SequenceLike,
) {
	var expression = sequence.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
}

func (v *visitor_) visitStatement(
	statement ast.StatementLike,
) {
	var mainClause = statement.GetMainClause()
	v.processor_.PreprocessMainClause(
		mainClause,
		1,
		1,
	)
	v.visitMainClause(mainClause)
	v.processor_.PostprocessMainClause(
		mainClause,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessStatementSlot(
		statement,
		1,
	)

	var optionalOnClause = statement.GetOptionalOnClause()
	if uti.IsDefined(optionalOnClause) {
		v.processor_.PreprocessOnClause(
			optionalOnClause,
			1,
			1,
		)
		v.visitOnClause(optionalOnClause)
		v.processor_.PostprocessOnClause(
			optionalOnClause,
			1,
			1,
		)
	}
}

func (v *visitor_) visitString(
	string_ ast.StringLike,
) {
	// Visit the possible string expression types.
	var actual = string_.GetAny().(string)
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
	case ScannerClass().MatchesType(actual, TagToken):
		v.processor_.ProcessTag(actual)
	case ScannerClass().MatchesType(actual, VersionToken):
		v.processor_.ProcessVersion(actual)
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
	case ast.DocumentLike:
		v.processor_.PreprocessDocument(
			actual,
			1,
			1,
		)
		v.visitDocument(actual)
		v.processor_.PostprocessDocument(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			1,
			1,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			1,
			1,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			1,
			1,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			1,
			1,
		)
	case ast.ComplementLike:
		v.processor_.PreprocessComplement(
			actual,
			1,
			1,
		)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(
			actual,
			1,
			1,
		)
	case ast.InversionLike:
		v.processor_.PreprocessInversion(
			actual,
			1,
			1,
		)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(
			actual,
			1,
			1,
		)
	case ast.MagnitudeLike:
		v.processor_.PreprocessMagnitude(
			actual,
			1,
			1,
		)
		v.visitMagnitude(actual)
		v.processor_.PostprocessMagnitude(
			actual,
			1,
			1,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			1,
			1,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			1,
			1,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			1,
			1,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			1,
			1,
		)
	case ast.ValueLike:
		v.processor_.PreprocessValue(
			actual,
			1,
			1,
		)
		v.visitValue(actual)
		v.processor_.PostprocessValue(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitTemplate(
	template ast.TemplateLike,
) {
	var expression = template.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
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

	var exception = throwClause.GetException()
	v.processor_.PreprocessException(
		exception,
		1,
		1,
	)
	v.visitException(exception)
	v.processor_.PostprocessException(
		exception,
		1,
		1,
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

	var condition = whileClause.GetCondition()
	v.processor_.PreprocessCondition(
		condition,
		1,
		1,
	)
	v.visitCondition(condition)
	v.processor_.PostprocessCondition(
		condition,
		1,
		1,
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
		1,
		1,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		1,
		1,
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

	var variable = withClause.GetVariable()
	v.processor_.PreprocessVariable(
		variable,
		1,
		1,
	)
	v.visitVariable(variable)
	v.processor_.PostprocessVariable(
		variable,
		1,
		1,
	)
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

	var sequence = withClause.GetSequence()
	v.processor_.PreprocessSequence(
		sequence,
		1,
		1,
	)
	v.visitSequence(sequence)
	v.processor_.PostprocessSequence(
		sequence,
		1,
		1,
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
		1,
		1,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		1,
		1,
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
