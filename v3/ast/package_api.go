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
│         This "package_api.go" file was automatically generated using:        │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.

For detailed documentation on this package refer to the wiki:
  - https://github.com/bali-nebula/go-document-notation/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AcceptClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete accept-clause-like class.
*/
type AcceptClauseClassLike interface {
	// Constructor Methods
	AcceptClause(
		delimiter string,
		message MessageLike,
	) AcceptClauseLike
}

/*
ActionInductionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete action-induction-like class.
*/
type ActionInductionClassLike interface {
	// Constructor Methods
	ActionInduction(
		any_ any,
	) ActionInductionLike
}

/*
AnnotationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete annotation-like class.
*/
type AnnotationClassLike interface {
	// Constructor Methods
	Annotation(
		any_ any,
	) AnnotationLike
}

/*
ArgumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete argument-like class.
*/
type ArgumentClassLike interface {
	// Constructor Methods
	Argument(
		any_ any,
	) ArgumentLike
}

/*
ArithmeticOperatorClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete arithmetic-operator-like class.
*/
type ArithmeticOperatorClassLike interface {
	// Constructor Methods
	ArithmeticOperator(
		any_ any,
	) ArithmeticOperatorLike
}

/*
AssignmentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete assignment-like class.
*/
type AssignmentClassLike interface {
	// Constructor Methods
	Assignment(
		any_ any,
	) AssignmentLike
}

/*
AssociationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete association-like class.
*/
type AssociationClassLike interface {
	// Constructor Methods
	Association(
		primitive PrimitiveLike,
		delimiter string,
		document DocumentLike,
	) AssociationLike
}

/*
AtLevelClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete at-level-like class.
*/
type AtLevelClassLike interface {
	// Constructor Methods
	AtLevel(
		delimiter1 string,
		delimiter2 string,
		expression ExpressionLike,
	) AtLevelLike
}

/*
AttributesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete attributes-like class.
*/
type AttributesClassLike interface {
	// Constructor Methods
	Attributes(
		delimiter1 string,
		associations fra.Sequential[AssociationLike],
		delimiter2 string,
	) AttributesLike
}

/*
BagClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete bag-like class.
*/
type BagClassLike interface {
	// Constructor Methods
	Bag(
		expression ExpressionLike,
	) BagLike
}

/*
BreakClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete break-clause-like class.
*/
type BreakClauseClassLike interface {
	// Constructor Methods
	BreakClause(
		delimiter1 string,
		delimiter2 string,
	) BreakClauseLike
}

/*
CheckoutClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete checkout-clause-like class.
*/
type CheckoutClauseClassLike interface {
	// Constructor Methods
	CheckoutClause(
		delimiter1 string,
		recipient RecipientLike,
		optionalAtLevel AtLevelLike,
		delimiter2 string,
		cited CitedLike,
	) CheckoutClauseLike
}

/*
CitedClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete cited-like class.
*/
type CitedClassLike interface {
	// Constructor Methods
	Cited(
		expression ExpressionLike,
	) CitedLike
}

/*
CollectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete collection-like class.
*/
type CollectionClassLike interface {
	// Constructor Methods
	Collection(
		any_ any,
	) CollectionLike
}

/*
ComparisonOperatorClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete comparison-operator-like class.
*/
type ComparisonOperatorClassLike interface {
	// Constructor Methods
	ComparisonOperator(
		any_ any,
	) ComparisonOperatorLike
}

/*
ComplementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete complement-like class.
*/
type ComplementClassLike interface {
	// Constructor Methods
	Complement(
		delimiter string,
		logical LogicalLike,
	) ComplementLike
}

/*
ComponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructor Methods
	Component(
		any_ any,
	) ComponentLike
}

/*
ConditionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete condition-like class.
*/
type ConditionClassLike interface {
	// Constructor Methods
	Condition(
		expression ExpressionLike,
	) ConditionLike
}

/*
ContinueClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete continue-clause-like class.
*/
type ContinueClauseClassLike interface {
	// Constructor Methods
	ContinueClause(
		delimiter1 string,
		delimiter2 string,
	) ContinueClauseLike
}

/*
DiscardClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete discard-clause-like class.
*/
type DiscardClauseClassLike interface {
	// Constructor Methods
	DiscardClause(
		delimiter string,
		draft DraftLike,
	) DiscardClauseLike
}

/*
DoClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete do-clause-like class.
*/
type DoClauseClassLike interface {
	// Constructor Methods
	DoClause(
		delimiter string,
		method MethodLike,
	) DoClauseLike
}

/*
DocumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructor Methods
	Document(
		component ComponentLike,
		optionalParameters ParametersLike,
		optionalNote string,
	) DocumentLike
}

/*
DraftClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete draft-like class.
*/
type DraftClassLike interface {
	// Constructor Methods
	Draft(
		expression ExpressionLike,
	) DraftLike
}

/*
ElementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete element-like class.
*/
type ElementClassLike interface {
	// Constructor Methods
	Element(
		any_ any,
	) ElementLike
}

/*
EntitiesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete entities-like class.
*/
type EntitiesClassLike interface {
	// Constructor Methods
	Entities(
		delimiter1 string,
		items fra.Sequential[ItemLike],
		delimiter2 string,
	) EntitiesLike
}

/*
EventClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete event-like class.
*/
type EventClassLike interface {
	// Constructor Methods
	Event(
		expression ExpressionLike,
	) EventLike
}

/*
ExceptionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete exception-like class.
*/
type ExceptionClassLike interface {
	// Constructor Methods
	Exception(
		expression ExpressionLike,
	) ExceptionLike
}

/*
ExpressionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete expression-like class.
*/
type ExpressionClassLike interface {
	// Constructor Methods
	Expression(
		subject SubjectLike,
		predicates fra.Sequential[PredicateLike],
	) ExpressionLike
}

/*
FailureClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete failure-like class.
*/
type FailureClassLike interface {
	// Constructor Methods
	Failure(
		symbol string,
	) FailureLike
}

/*
FlowControlClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete flow-control-like class.
*/
type FlowControlClassLike interface {
	// Constructor Methods
	FlowControl(
		any_ any,
	) FlowControlLike
}

/*
FunctionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete function-like class.
*/
type FunctionClassLike interface {
	// Constructor Methods
	Function(
		identifier string,
		delimiter1 string,
		arguments fra.Sequential[ArgumentLike],
		delimiter2 string,
	) FunctionLike
}

/*
IfClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete if-clause-like class.
*/
type IfClauseClassLike interface {
	// Constructor Methods
	IfClause(
		delimiter1 string,
		condition ConditionLike,
		delimiter2 string,
		procedure ProcedureLike,
	) IfClauseLike
}

/*
IndexClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete index-like class.
*/
type IndexClassLike interface {
	// Constructor Methods
	Index(
		any_ any,
	) IndexLike
}

/*
InverseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inverse-like class.
*/
type InverseClassLike interface {
	// Constructor Methods
	Inverse(
		any_ any,
	) InverseLike
}

/*
InversionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inversion-like class.
*/
type InversionClassLike interface {
	// Constructor Methods
	Inversion(
		inverse InverseLike,
		numerical NumericalLike,
	) InversionLike
}

/*
InvokeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete invoke-like class.
*/
type InvokeClassLike interface {
	// Constructor Methods
	Invoke(
		any_ any,
	) InvokeLike
}

/*
ItemClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete item-like class.
*/
type ItemClassLike interface {
	// Constructor Methods
	Item(
		document DocumentLike,
	) ItemLike
}

/*
LeftClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete left-like class.
*/
type LeftClassLike interface {
	// Constructor Methods
	Left(
		any_ any,
	) LeftLike
}

/*
LetClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete let-clause-like class.
*/
type LetClauseClassLike interface {
	// Constructor Methods
	LetClause(
		delimiter string,
		recipient RecipientLike,
		assignment AssignmentLike,
		expression ExpressionLike,
	) LetClauseLike
}

/*
LexicalOperatorClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete lexical-operator-like class.
*/
type LexicalOperatorClassLike interface {
	// Constructor Methods
	LexicalOperator(
		any_ any,
	) LexicalOperatorLike
}

/*
LineClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete line-like class.
*/
type LineClassLike interface {
	// Constructor Methods
	Line(
		any_ any,
	) LineLike
}

/*
LogicalClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete logical-like class.
*/
type LogicalClassLike interface {
	// Constructor Methods
	Logical(
		any_ any,
	) LogicalLike
}

/*
LogicalOperatorClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete logical-operator-like class.
*/
type LogicalOperatorClassLike interface {
	// Constructor Methods
	LogicalOperator(
		any_ any,
	) LogicalOperatorLike
}

/*
MagnitudeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete magnitude-like class.
*/
type MagnitudeClassLike interface {
	// Constructor Methods
	Magnitude(
		delimiter1 string,
		expression ExpressionLike,
		delimiter2 string,
	) MagnitudeLike
}

/*
MainClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete main-clause-like class.
*/
type MainClauseClassLike interface {
	// Constructor Methods
	MainClause(
		any_ any,
	) MainClauseLike
}

/*
MatchingClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete matching-clause-like class.
*/
type MatchingClauseClassLike interface {
	// Constructor Methods
	MatchingClause(
		delimiter1 string,
		template TemplateLike,
		delimiter2 string,
		procedure ProcedureLike,
	) MatchingClauseLike
}

/*
MessageClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete message-like class.
*/
type MessageClassLike interface {
	// Constructor Methods
	Message(
		expression ExpressionLike,
	) MessageLike
}

/*
MessageHandlingClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete message-handling-like class.
*/
type MessageHandlingClassLike interface {
	// Constructor Methods
	MessageHandling(
		any_ any,
	) MessageHandlingLike
}

/*
MethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete method-like class.
*/
type MethodClassLike interface {
	// Constructor Methods
	Method(
		identifier1 string,
		invoke InvokeLike,
		identifier2 string,
		delimiter1 string,
		arguments fra.Sequential[ArgumentLike],
		delimiter2 string,
	) MethodLike
}

/*
NotarizeClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete notarize-clause-like class.
*/
type NotarizeClauseClassLike interface {
	// Constructor Methods
	NotarizeClause(
		delimiter1 string,
		draft DraftLike,
		delimiter2 string,
		cited CitedLike,
	) NotarizeClauseLike
}

/*
NumericalClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete numerical-like class.
*/
type NumericalClassLike interface {
	// Constructor Methods
	Numerical(
		any_ any,
	) NumericalLike
}

/*
OnClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete on-clause-like class.
*/
type OnClauseClassLike interface {
	// Constructor Methods
	OnClause(
		delimiter string,
		failure FailureLike,
		matchingClauses fra.Sequential[MatchingClauseLike],
	) OnClauseLike
}

/*
OperatorClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete operator-like class.
*/
type OperatorClassLike interface {
	// Constructor Methods
	Operator(
		any_ any,
	) OperatorLike
}

/*
ParametersClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete parameters-like class.
*/
type ParametersClassLike interface {
	// Constructor Methods
	Parameters(
		delimiter1 string,
		associations fra.Sequential[AssociationLike],
		delimiter2 string,
	) ParametersLike
}

/*
PostClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete post-clause-like class.
*/
type PostClauseClassLike interface {
	// Constructor Methods
	PostClause(
		delimiter1 string,
		message MessageLike,
		delimiter2 string,
		bag BagLike,
	) PostClauseLike
}

/*
PrecedenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete precedence-like class.
*/
type PrecedenceClassLike interface {
	// Constructor Methods
	Precedence(
		delimiter1 string,
		expression ExpressionLike,
		delimiter2 string,
	) PrecedenceLike
}

/*
PredicateClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete predicate-like class.
*/
type PredicateClassLike interface {
	// Constructor Methods
	Predicate(
		operator OperatorLike,
		expression ExpressionLike,
	) PredicateLike
}

/*
PrimitiveClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete primitive-like class.
*/
type PrimitiveClassLike interface {
	// Constructor Methods
	Primitive(
		any_ any,
	) PrimitiveLike
}

/*
ProcedureClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete procedure-like class.
*/
type ProcedureClassLike interface {
	// Constructor Methods
	Procedure(
		delimiter1 string,
		lines fra.Sequential[LineLike],
		delimiter2 string,
	) ProcedureLike
}

/*
PublishClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete publish-clause-like class.
*/
type PublishClauseClassLike interface {
	// Constructor Methods
	PublishClause(
		delimiter string,
		event EventLike,
	) PublishClauseLike
}

/*
RangeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete range-like class.
*/
type RangeClassLike interface {
	// Constructor Methods
	Range(
		left LeftLike,
		primitive1 PrimitiveLike,
		delimiter string,
		primitive2 PrimitiveLike,
		right RightLike,
	) RangeLike
}

/*
RecipientClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete recipient-like class.
*/
type RecipientClassLike interface {
	// Constructor Methods
	Recipient(
		any_ any,
	) RecipientLike
}

/*
ReferenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete reference-like class.
*/
type ReferenceClassLike interface {
	// Constructor Methods
	Reference(
		any_ any,
	) ReferenceLike
}

/*
ReferentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete referent-like class.
*/
type ReferentClassLike interface {
	// Constructor Methods
	Referent(
		delimiter string,
		reference ReferenceLike,
	) ReferentLike
}

/*
RejectClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete reject-clause-like class.
*/
type RejectClauseClassLike interface {
	// Constructor Methods
	RejectClause(
		delimiter string,
		message MessageLike,
	) RejectClauseLike
}

/*
RepositoryAccessClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete repository-access-like class.
*/
type RepositoryAccessClassLike interface {
	// Constructor Methods
	RepositoryAccess(
		any_ any,
	) RepositoryAccessLike
}

/*
ResultClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete result-like class.
*/
type ResultClassLike interface {
	// Constructor Methods
	Result(
		expression ExpressionLike,
	) ResultLike
}

/*
RetrieveClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete retrieve-clause-like class.
*/
type RetrieveClauseClassLike interface {
	// Constructor Methods
	RetrieveClause(
		delimiter1 string,
		recipient RecipientLike,
		delimiter2 string,
		bag BagLike,
	) RetrieveClauseLike
}

/*
ReturnClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete return-clause-like class.
*/
type ReturnClauseClassLike interface {
	// Constructor Methods
	ReturnClause(
		delimiter string,
		result ResultLike,
	) ReturnClauseLike
}

/*
RightClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete right-like class.
*/
type RightClassLike interface {
	// Constructor Methods
	Right(
		any_ any,
	) RightLike
}

/*
SaveClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete save-clause-like class.
*/
type SaveClauseClassLike interface {
	// Constructor Methods
	SaveClause(
		delimiter1 string,
		draft DraftLike,
		delimiter2 string,
		cited CitedLike,
	) SaveClauseLike
}

/*
SelectClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete select-clause-like class.
*/
type SelectClauseClassLike interface {
	// Constructor Methods
	SelectClause(
		delimiter string,
		expression ExpressionLike,
		matchingClauses fra.Sequential[MatchingClauseLike],
	) SelectClauseLike
}

/*
SequenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete sequence-like class.
*/
type SequenceClassLike interface {
	// Constructor Methods
	Sequence(
		expression ExpressionLike,
	) SequenceLike
}

/*
StatementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete statement-like class.
*/
type StatementClassLike interface {
	// Constructor Methods
	Statement(
		mainClause MainClauseLike,
		optionalOnClause OnClauseLike,
	) StatementLike
}

/*
StringClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete string-like class.
*/
type StringClassLike interface {
	// Constructor Methods
	String(
		any_ any,
	) StringLike
}

/*
SubcomponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete subcomponent-like class.
*/
type SubcomponentClassLike interface {
	// Constructor Methods
	Subcomponent(
		identifier string,
		delimiter1 string,
		indexes fra.Sequential[IndexLike],
		delimiter2 string,
	) SubcomponentLike
}

/*
SubjectClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete subject-like class.
*/
type SubjectClassLike interface {
	// Constructor Methods
	Subject(
		any_ any,
	) SubjectLike
}

/*
TemplateClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete template-like class.
*/
type TemplateClassLike interface {
	// Constructor Methods
	Template(
		expression ExpressionLike,
	) TemplateLike
}

/*
ThrowClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete throw-clause-like class.
*/
type ThrowClauseClassLike interface {
	// Constructor Methods
	ThrowClause(
		delimiter string,
		exception ExceptionLike,
	) ThrowClauseLike
}

/*
ValueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete value-like class.
*/
type ValueClassLike interface {
	// Constructor Methods
	Value(
		identifier string,
	) ValueLike
}

/*
VariableClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete variable-like class.
*/
type VariableClassLike interface {
	// Constructor Methods
	Variable(
		symbol string,
	) VariableLike
}

/*
WhileClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete while-clause-like class.
*/
type WhileClauseClassLike interface {
	// Constructor Methods
	WhileClause(
		delimiter1 string,
		condition ConditionLike,
		delimiter2 string,
		procedure ProcedureLike,
	) WhileClauseLike
}

/*
WithClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete with-clause-like class.
*/
type WithClauseClassLike interface {
	// Constructor Methods
	WithClause(
		delimiter1 string,
		delimiter2 string,
		variable VariableLike,
		delimiter3 string,
		sequence SequenceLike,
		delimiter4 string,
		procedure ProcedureLike,
	) WithClauseLike
}

// INSTANCE DECLARATIONS

/*
AcceptClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete accept-clause-like class.
*/
type AcceptClauseLike interface {
	// Principal Methods
	GetClass() AcceptClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetMessage() MessageLike
}

/*
ActionInductionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete action-induction-like class.
*/
type ActionInductionLike interface {
	// Principal Methods
	GetClass() ActionInductionClassLike

	// Attribute Methods
	GetAny() any
}

/*
AnnotationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete annotation-like class.
*/
type AnnotationLike interface {
	// Principal Methods
	GetClass() AnnotationClassLike

	// Attribute Methods
	GetAny() any
}

/*
ArgumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Principal Methods
	GetClass() ArgumentClassLike

	// Attribute Methods
	GetAny() any
}

/*
ArithmeticOperatorLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete arithmetic-operator-like class.
*/
type ArithmeticOperatorLike interface {
	// Principal Methods
	GetClass() ArithmeticOperatorClassLike

	// Attribute Methods
	GetAny() any
}

/*
AssignmentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete assignment-like class.
*/
type AssignmentLike interface {
	// Principal Methods
	GetClass() AssignmentClassLike

	// Attribute Methods
	GetAny() any
}

/*
AssociationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete association-like class.
*/
type AssociationLike interface {
	// Principal Methods
	GetClass() AssociationClassLike

	// Attribute Methods
	GetPrimitive() PrimitiveLike
	GetDelimiter() string
	GetDocument() DocumentLike
}

/*
AtLevelLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete at-level-like class.
*/
type AtLevelLike interface {
	// Principal Methods
	GetClass() AtLevelClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
	GetExpression() ExpressionLike
}

/*
AttributesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete attributes-like class.
*/
type AttributesLike interface {
	// Principal Methods
	GetClass() AttributesClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetAssociations() fra.Sequential[AssociationLike]
	GetDelimiter2() string
}

/*
BagLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete bag-like class.
*/
type BagLike interface {
	// Principal Methods
	GetClass() BagClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
BreakClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete break-clause-like class.
*/
type BreakClauseLike interface {
	// Principal Methods
	GetClass() BreakClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
}

/*
CheckoutClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete checkout-clause-like class.
*/
type CheckoutClauseLike interface {
	// Principal Methods
	GetClass() CheckoutClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetRecipient() RecipientLike
	GetOptionalAtLevel() AtLevelLike
	GetDelimiter2() string
	GetCited() CitedLike
}

/*
CitedLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete cited-like class.
*/
type CitedLike interface {
	// Principal Methods
	GetClass() CitedClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
CollectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete collection-like class.
*/
type CollectionLike interface {
	// Principal Methods
	GetClass() CollectionClassLike

	// Attribute Methods
	GetAny() any
}

/*
ComparisonOperatorLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete comparison-operator-like class.
*/
type ComparisonOperatorLike interface {
	// Principal Methods
	GetClass() ComparisonOperatorClassLike

	// Attribute Methods
	GetAny() any
}

/*
ComplementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete complement-like class.
*/
type ComplementLike interface {
	// Principal Methods
	GetClass() ComplementClassLike

	// Attribute Methods
	GetDelimiter() string
	GetLogical() LogicalLike
}

/*
ComponentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Principal Methods
	GetClass() ComponentClassLike

	// Attribute Methods
	GetAny() any
}

/*
ConditionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete condition-like class.
*/
type ConditionLike interface {
	// Principal Methods
	GetClass() ConditionClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ContinueClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete continue-clause-like class.
*/
type ContinueClauseLike interface {
	// Principal Methods
	GetClass() ContinueClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
}

/*
DiscardClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete discard-clause-like class.
*/
type DiscardClauseLike interface {
	// Principal Methods
	GetClass() DiscardClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetDraft() DraftLike
}

/*
DoClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete do-clause-like class.
*/
type DoClauseLike interface {
	// Principal Methods
	GetClass() DoClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetMethod() MethodLike
}

/*
DocumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete document-like class.
*/
type DocumentLike interface {
	// Principal Methods
	GetClass() DocumentClassLike

	// Attribute Methods
	GetComponent() ComponentLike
	GetOptionalParameters() ParametersLike
	GetOptionalNote() string
}

/*
DraftLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete draft-like class.
*/
type DraftLike interface {
	// Principal Methods
	GetClass() DraftClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ElementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete element-like class.
*/
type ElementLike interface {
	// Principal Methods
	GetClass() ElementClassLike

	// Attribute Methods
	GetAny() any
}

/*
EntitiesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete entities-like class.
*/
type EntitiesLike interface {
	// Principal Methods
	GetClass() EntitiesClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetItems() fra.Sequential[ItemLike]
	GetDelimiter2() string
}

/*
EventLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete event-like class.
*/
type EventLike interface {
	// Principal Methods
	GetClass() EventClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ExceptionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete exception-like class.
*/
type ExceptionLike interface {
	// Principal Methods
	GetClass() ExceptionClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ExpressionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete expression-like class.
*/
type ExpressionLike interface {
	// Principal Methods
	GetClass() ExpressionClassLike

	// Attribute Methods
	GetSubject() SubjectLike
	GetPredicates() fra.Sequential[PredicateLike]
}

/*
FailureLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete failure-like class.
*/
type FailureLike interface {
	// Principal Methods
	GetClass() FailureClassLike

	// Attribute Methods
	GetSymbol() string
}

/*
FlowControlLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete flow-control-like class.
*/
type FlowControlLike interface {
	// Principal Methods
	GetClass() FlowControlClassLike

	// Attribute Methods
	GetAny() any
}

/*
FunctionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete function-like class.
*/
type FunctionLike interface {
	// Principal Methods
	GetClass() FunctionClassLike

	// Attribute Methods
	GetIdentifier() string
	GetDelimiter1() string
	GetArguments() fra.Sequential[ArgumentLike]
	GetDelimiter2() string
}

/*
IfClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete if-clause-like class.
*/
type IfClauseLike interface {
	// Principal Methods
	GetClass() IfClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetCondition() ConditionLike
	GetDelimiter2() string
	GetProcedure() ProcedureLike
}

/*
IndexLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete index-like class.
*/
type IndexLike interface {
	// Principal Methods
	GetClass() IndexClassLike

	// Attribute Methods
	GetAny() any
}

/*
InverseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inverse-like class.
*/
type InverseLike interface {
	// Principal Methods
	GetClass() InverseClassLike

	// Attribute Methods
	GetAny() any
}

/*
InversionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inversion-like class.
*/
type InversionLike interface {
	// Principal Methods
	GetClass() InversionClassLike

	// Attribute Methods
	GetInverse() InverseLike
	GetNumerical() NumericalLike
}

/*
InvokeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete invoke-like class.
*/
type InvokeLike interface {
	// Principal Methods
	GetClass() InvokeClassLike

	// Attribute Methods
	GetAny() any
}

/*
ItemLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete item-like class.
*/
type ItemLike interface {
	// Principal Methods
	GetClass() ItemClassLike

	// Attribute Methods
	GetDocument() DocumentLike
}

/*
LeftLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete left-like class.
*/
type LeftLike interface {
	// Principal Methods
	GetClass() LeftClassLike

	// Attribute Methods
	GetAny() any
}

/*
LetClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete let-clause-like class.
*/
type LetClauseLike interface {
	// Principal Methods
	GetClass() LetClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetRecipient() RecipientLike
	GetAssignment() AssignmentLike
	GetExpression() ExpressionLike
}

/*
LexicalOperatorLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete lexical-operator-like class.
*/
type LexicalOperatorLike interface {
	// Principal Methods
	GetClass() LexicalOperatorClassLike

	// Attribute Methods
	GetAny() any
}

/*
LineLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete line-like class.
*/
type LineLike interface {
	// Principal Methods
	GetClass() LineClassLike

	// Attribute Methods
	GetAny() any
}

/*
LogicalLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete logical-like class.
*/
type LogicalLike interface {
	// Principal Methods
	GetClass() LogicalClassLike

	// Attribute Methods
	GetAny() any
}

/*
LogicalOperatorLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete logical-operator-like class.
*/
type LogicalOperatorLike interface {
	// Principal Methods
	GetClass() LogicalOperatorClassLike

	// Attribute Methods
	GetAny() any
}

/*
MagnitudeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete magnitude-like class.
*/
type MagnitudeLike interface {
	// Principal Methods
	GetClass() MagnitudeClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetExpression() ExpressionLike
	GetDelimiter2() string
}

/*
MainClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete main-clause-like class.
*/
type MainClauseLike interface {
	// Principal Methods
	GetClass() MainClauseClassLike

	// Attribute Methods
	GetAny() any
}

/*
MatchingClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete matching-clause-like class.
*/
type MatchingClauseLike interface {
	// Principal Methods
	GetClass() MatchingClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetTemplate() TemplateLike
	GetDelimiter2() string
	GetProcedure() ProcedureLike
}

/*
MessageLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete message-like class.
*/
type MessageLike interface {
	// Principal Methods
	GetClass() MessageClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
MessageHandlingLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete message-handling-like class.
*/
type MessageHandlingLike interface {
	// Principal Methods
	GetClass() MessageHandlingClassLike

	// Attribute Methods
	GetAny() any
}

/*
MethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete method-like class.
*/
type MethodLike interface {
	// Principal Methods
	GetClass() MethodClassLike

	// Attribute Methods
	GetIdentifier1() string
	GetInvoke() InvokeLike
	GetIdentifier2() string
	GetDelimiter1() string
	GetArguments() fra.Sequential[ArgumentLike]
	GetDelimiter2() string
}

/*
NotarizeClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete notarize-clause-like class.
*/
type NotarizeClauseLike interface {
	// Principal Methods
	GetClass() NotarizeClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDraft() DraftLike
	GetDelimiter2() string
	GetCited() CitedLike
}

/*
NumericalLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete numerical-like class.
*/
type NumericalLike interface {
	// Principal Methods
	GetClass() NumericalClassLike

	// Attribute Methods
	GetAny() any
}

/*
OnClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete on-clause-like class.
*/
type OnClauseLike interface {
	// Principal Methods
	GetClass() OnClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetFailure() FailureLike
	GetMatchingClauses() fra.Sequential[MatchingClauseLike]
}

/*
OperatorLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete operator-like class.
*/
type OperatorLike interface {
	// Principal Methods
	GetClass() OperatorClassLike

	// Attribute Methods
	GetAny() any
}

/*
ParametersLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete parameters-like class.
*/
type ParametersLike interface {
	// Principal Methods
	GetClass() ParametersClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetAssociations() fra.Sequential[AssociationLike]
	GetDelimiter2() string
}

/*
PostClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete post-clause-like class.
*/
type PostClauseLike interface {
	// Principal Methods
	GetClass() PostClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetMessage() MessageLike
	GetDelimiter2() string
	GetBag() BagLike
}

/*
PrecedenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete precedence-like class.
*/
type PrecedenceLike interface {
	// Principal Methods
	GetClass() PrecedenceClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetExpression() ExpressionLike
	GetDelimiter2() string
}

/*
PredicateLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete predicate-like class.
*/
type PredicateLike interface {
	// Principal Methods
	GetClass() PredicateClassLike

	// Attribute Methods
	GetOperator() OperatorLike
	GetExpression() ExpressionLike
}

/*
PrimitiveLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete primitive-like class.
*/
type PrimitiveLike interface {
	// Principal Methods
	GetClass() PrimitiveClassLike

	// Attribute Methods
	GetAny() any
}

/*
ProcedureLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete procedure-like class.
*/
type ProcedureLike interface {
	// Principal Methods
	GetClass() ProcedureClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetLines() fra.Sequential[LineLike]
	GetDelimiter2() string
}

/*
PublishClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete publish-clause-like class.
*/
type PublishClauseLike interface {
	// Principal Methods
	GetClass() PublishClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetEvent() EventLike
}

/*
RangeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete range-like class.
*/
type RangeLike interface {
	// Principal Methods
	GetClass() RangeClassLike

	// Attribute Methods
	GetLeft() LeftLike
	GetPrimitive1() PrimitiveLike
	GetDelimiter() string
	GetPrimitive2() PrimitiveLike
	GetRight() RightLike
}

/*
RecipientLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete recipient-like class.
*/
type RecipientLike interface {
	// Principal Methods
	GetClass() RecipientClassLike

	// Attribute Methods
	GetAny() any
}

/*
ReferenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete reference-like class.
*/
type ReferenceLike interface {
	// Principal Methods
	GetClass() ReferenceClassLike

	// Attribute Methods
	GetAny() any
}

/*
ReferentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete referent-like class.
*/
type ReferentLike interface {
	// Principal Methods
	GetClass() ReferentClassLike

	// Attribute Methods
	GetDelimiter() string
	GetReference() ReferenceLike
}

/*
RejectClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete reject-clause-like class.
*/
type RejectClauseLike interface {
	// Principal Methods
	GetClass() RejectClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetMessage() MessageLike
}

/*
RepositoryAccessLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete repository-access-like class.
*/
type RepositoryAccessLike interface {
	// Principal Methods
	GetClass() RepositoryAccessClassLike

	// Attribute Methods
	GetAny() any
}

/*
ResultLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete result-like class.
*/
type ResultLike interface {
	// Principal Methods
	GetClass() ResultClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
RetrieveClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete retrieve-clause-like class.
*/
type RetrieveClauseLike interface {
	// Principal Methods
	GetClass() RetrieveClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetRecipient() RecipientLike
	GetDelimiter2() string
	GetBag() BagLike
}

/*
ReturnClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete return-clause-like class.
*/
type ReturnClauseLike interface {
	// Principal Methods
	GetClass() ReturnClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetResult() ResultLike
}

/*
RightLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete right-like class.
*/
type RightLike interface {
	// Principal Methods
	GetClass() RightClassLike

	// Attribute Methods
	GetAny() any
}

/*
SaveClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete save-clause-like class.
*/
type SaveClauseLike interface {
	// Principal Methods
	GetClass() SaveClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDraft() DraftLike
	GetDelimiter2() string
	GetCited() CitedLike
}

/*
SelectClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete select-clause-like class.
*/
type SelectClauseLike interface {
	// Principal Methods
	GetClass() SelectClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetExpression() ExpressionLike
	GetMatchingClauses() fra.Sequential[MatchingClauseLike]
}

/*
SequenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete sequence-like class.
*/
type SequenceLike interface {
	// Principal Methods
	GetClass() SequenceClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
StatementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete statement-like class.
*/
type StatementLike interface {
	// Principal Methods
	GetClass() StatementClassLike

	// Attribute Methods
	GetMainClause() MainClauseLike
	GetOptionalOnClause() OnClauseLike
}

/*
StringLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete string-like class.
*/
type StringLike interface {
	// Principal Methods
	GetClass() StringClassLike

	// Attribute Methods
	GetAny() any
}

/*
SubcomponentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete subcomponent-like class.
*/
type SubcomponentLike interface {
	// Principal Methods
	GetClass() SubcomponentClassLike

	// Attribute Methods
	GetIdentifier() string
	GetDelimiter1() string
	GetIndexes() fra.Sequential[IndexLike]
	GetDelimiter2() string
}

/*
SubjectLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete subject-like class.
*/
type SubjectLike interface {
	// Principal Methods
	GetClass() SubjectClassLike

	// Attribute Methods
	GetAny() any
}

/*
TemplateLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete template-like class.
*/
type TemplateLike interface {
	// Principal Methods
	GetClass() TemplateClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ThrowClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete throw-clause-like class.
*/
type ThrowClauseLike interface {
	// Principal Methods
	GetClass() ThrowClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetException() ExceptionLike
}

/*
ValueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete value-like class.
*/
type ValueLike interface {
	// Principal Methods
	GetClass() ValueClassLike

	// Attribute Methods
	GetIdentifier() string
}

/*
VariableLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete variable-like class.
*/
type VariableLike interface {
	// Principal Methods
	GetClass() VariableClassLike

	// Attribute Methods
	GetSymbol() string
}

/*
WhileClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete while-clause-like class.
*/
type WhileClauseLike interface {
	// Principal Methods
	GetClass() WhileClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetCondition() ConditionLike
	GetDelimiter2() string
	GetProcedure() ProcedureLike
}

/*
WithClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete with-clause-like class.
*/
type WithClauseLike interface {
	// Principal Methods
	GetClass() WithClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
	GetVariable() VariableLike
	GetDelimiter3() string
	GetSequence() SequenceLike
	GetDelimiter4() string
	GetProcedure() ProcedureLike
}

// ASPECT DECLARATIONS
