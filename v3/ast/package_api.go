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
	com "github.com/craterdog/go-essential-composites/v8"
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
		delimiter1 string,
		message MessageLike,
		delimiter2 string,
		bag BagLike,
	) AcceptClauseLike
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
ArithmeticClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete arithmetic-like class.
*/
type ArithmeticClassLike interface {
	// Constructor Methods
	Arithmetic(
		any_ any,
	) ArithmeticLike
}

/*
AssignClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete assign-clause-like class.
*/
type AssignClauseClassLike interface {
	// Constructor Methods
	AssignClause(
		delimiter string,
		recipient RecipientLike,
		assignment AssignmentLike,
		expression ExpressionLike,
	) AssignClauseLike
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
		content ContentLike,
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
		associations com.Sequential[AssociationLike],
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
		location LocationLike,
	) CheckoutClauseLike
}

/*
CitationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete citation-like class.
*/
type CitationClassLike interface {
	// Constructor Methods
	Citation(
		expression ExpressionLike,
	) CitationLike
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
ComparisonClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete comparison-like class.
*/
type ComparisonClassLike interface {
	// Constructor Methods
	Comparison(
		any_ any,
	) ComparisonLike
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
		reversible ReversibleLike,
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
		entity EntityLike,
		optionalGenerics GenericsLike,
	) ComponentLike
}

/*
ConstantClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constant-like class.
*/
type ConstantClassLike interface {
	// Constructor Methods
	Constant(
		symbol string,
	) ConstantLike
}

/*
ConstraintClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constraint-like class.
*/
type ConstraintClassLike interface {
	// Constructor Methods
	Constraint(
		metadata MetadataLike,
		optionalGenerics GenericsLike,
	) ConstraintLike
}

/*
ContentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete content-like class.
*/
type ContentClassLike interface {
	// Constructor Methods
	Content(
		component ComponentLike,
		optionalNote string,
	) ContentLike
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
DefineClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete define-clause-like class.
*/
type DefineClauseClassLike interface {
	// Constructor Methods
	DefineClause(
		delimiter1 string,
		constant ConstantLike,
		delimiter2 string,
		expression ExpressionLike,
	) DefineClauseLike
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
		citation CitationLike,
	) DiscardClauseLike
}

/*
DocumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructor Methods
	Document(
		optionalHeading HeadingLike,
		component ComponentLike,
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
EmptyClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete empty-like class.
*/
type EmptyClassLike interface {
	// Constructor Methods
	Empty(
		delimiter string,
	) EmptyLike
}

/*
EntityClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete entity-like class.
*/
type EntityClassLike interface {
	// Constructor Methods
	Entity(
		any_ any,
	) EntityLike
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
		predicates com.Sequential[PredicateLike],
	) ExpressionLike
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
		arguments com.Sequential[ArgumentLike],
		delimiter2 string,
	) FunctionLike
}

/*
GenericsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete generics-like class.
*/
type GenericsClassLike interface {
	// Constructor Methods
	Generics(
		delimiter1 string,
		parameters com.Sequential[ParameterLike],
		delimiter2 string,
	) GenericsLike
}

/*
HeadingClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete heading-like class.
*/
type HeadingClassLike interface {
	// Constructor Methods
	Heading(
		comment string,
	) HeadingLike
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
		expression ExpressionLike,
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
InspectClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inspect-clause-like class.
*/
type InspectClauseClassLike interface {
	// Constructor Methods
	InspectClause(
		delimiter1 string,
		recipient RecipientLike,
		delimiter2 string,
		location LocationLike,
	) InspectClauseLike
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
InvocationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete invocation-like class.
*/
type InvocationClassLike interface {
	// Constructor Methods
	Invocation(
		any_ any,
	) InvocationLike
}

/*
InvokeClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete invoke-clause-like class.
*/
type InvokeClauseClassLike interface {
	// Constructor Methods
	InvokeClause(
		delimiter string,
		method MethodLike,
	) InvokeClauseLike
}

/*
ItemsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete items-like class.
*/
type ItemsClassLike interface {
	// Constructor Methods
	Items(
		delimiter1 string,
		contents com.Sequential[ContentLike],
		delimiter2 string,
	) ItemsLike
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
LexicalClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete lexical-like class.
*/
type LexicalClassLike interface {
	// Constructor Methods
	Lexical(
		any_ any,
	) LexicalLike
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
LocalTransformationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete local-transformation-like class.
*/
type LocalTransformationClassLike interface {
	// Constructor Methods
	LocalTransformation(
		any_ any,
	) LocalTransformationLike
}

/*
LocationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete location-like class.
*/
type LocationClassLike interface {
	// Constructor Methods
	Location(
		expression ExpressionLike,
	) LocationLike
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
		expression ExpressionLike,
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
MetadataClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete metadata-like class.
*/
type MetadataClassLike interface {
	// Constructor Methods
	Metadata(
		any_ any,
	) MetadataLike
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
		invocation InvocationLike,
		identifier2 string,
		delimiter1 string,
		arguments com.Sequential[ArgumentLike],
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
		location LocationLike,
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
		symbol string,
		matchingClauses com.Sequential[MatchingClauseLike],
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
ParameterClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete parameter-like class.
*/
type ParameterClassLike interface {
	// Constructor Methods
	Parameter(
		symbol string,
		delimiter string,
		constraint ConstraintLike,
	) ParameterLike
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
		lines com.Sequential[LineLike],
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
		message MessageLike,
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
		optionalPrimitive1 PrimitiveLike,
		delimiter string,
		optionalPrimitive2 PrimitiveLike,
		right RightLike,
	) RangeLike
}

/*
ReceiveClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete receive-clause-like class.
*/
type ReceiveClauseClassLike interface {
	// Constructor Methods
	ReceiveClause(
		delimiter1 string,
		recipient RecipientLike,
		delimiter2 string,
		bag BagLike,
	) ReceiveClauseLike
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
		delimiter1 string,
		message MessageLike,
		delimiter2 string,
		bag BagLike,
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
		citation CitationLike,
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
		expression ExpressionLike,
	) ReturnClauseLike
}

/*
ReversibleClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete reversible-like class.
*/
type ReversibleClassLike interface {
	// Constructor Methods
	Reversible(
		any_ any,
	) ReversibleLike
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
		recipient RecipientLike,
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
		matchingClauses com.Sequential[MatchingClauseLike],
	) SelectClauseLike
}

/*
SendClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete send-clause-like class.
*/
type SendClauseClassLike interface {
	// Constructor Methods
	SendClause(
		delimiter1 string,
		message MessageLike,
		delimiter2 string,
		bag BagLike,
	) SendClauseLike
}

/*
SequenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete sequence-like class.
*/
type SequenceClassLike interface {
	// Constructor Methods
	Sequence(
		any_ any,
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
SubcomponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete subcomponent-like class.
*/
type SubcomponentClassLike interface {
	// Constructor Methods
	Subcomponent(
		identifier string,
		delimiter1 string,
		indexes com.Sequential[IndexLike],
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
ThrowClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete throw-clause-like class.
*/
type ThrowClauseClassLike interface {
	// Constructor Methods
	ThrowClause(
		delimiter string,
		expression ExpressionLike,
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
		expression ExpressionLike,
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
		symbol string,
		delimiter3 string,
		expression ExpressionLike,
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
	GetDelimiter1() string
	GetMessage() MessageLike
	GetDelimiter2() string
	GetBag() BagLike
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
ArithmeticLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete arithmetic-like class.
*/
type ArithmeticLike interface {
	// Principal Methods
	GetClass() ArithmeticClassLike

	// Attribute Methods
	GetAny() any
}

/*
AssignClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete assign-clause-like class.
*/
type AssignClauseLike interface {
	// Principal Methods
	GetClass() AssignClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetRecipient() RecipientLike
	GetAssignment() AssignmentLike
	GetExpression() ExpressionLike
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
	GetContent() ContentLike
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
	GetAssociations() com.Sequential[AssociationLike]
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
	GetLocation() LocationLike
}

/*
CitationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete citation-like class.
*/
type CitationLike interface {
	// Principal Methods
	GetClass() CitationClassLike

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
ComparisonLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete comparison-like class.
*/
type ComparisonLike interface {
	// Principal Methods
	GetClass() ComparisonClassLike

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
	GetReversible() ReversibleLike
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
	GetEntity() EntityLike
	GetOptionalGenerics() GenericsLike
}

/*
ConstantLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constant-like class.
*/
type ConstantLike interface {
	// Principal Methods
	GetClass() ConstantClassLike

	// Attribute Methods
	GetSymbol() string
}

/*
ConstraintLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constraint-like class.
*/
type ConstraintLike interface {
	// Principal Methods
	GetClass() ConstraintClassLike

	// Attribute Methods
	GetMetadata() MetadataLike
	GetOptionalGenerics() GenericsLike
}

/*
ContentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete content-like class.
*/
type ContentLike interface {
	// Principal Methods
	GetClass() ContentClassLike

	// Attribute Methods
	GetComponent() ComponentLike
	GetOptionalNote() string
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
DefineClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete define-clause-like class.
*/
type DefineClauseLike interface {
	// Principal Methods
	GetClass() DefineClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetConstant() ConstantLike
	GetDelimiter2() string
	GetExpression() ExpressionLike
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
	GetCitation() CitationLike
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
	GetOptionalHeading() HeadingLike
	GetComponent() ComponentLike
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
EmptyLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete empty-like class.
*/
type EmptyLike interface {
	// Principal Methods
	GetClass() EmptyClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
EntityLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete entity-like class.
*/
type EntityLike interface {
	// Principal Methods
	GetClass() EntityClassLike

	// Attribute Methods
	GetAny() any
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
	GetPredicates() com.Sequential[PredicateLike]
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
	GetArguments() com.Sequential[ArgumentLike]
	GetDelimiter2() string
}

/*
GenericsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete generics-like class.
*/
type GenericsLike interface {
	// Principal Methods
	GetClass() GenericsClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetParameters() com.Sequential[ParameterLike]
	GetDelimiter2() string
}

/*
HeadingLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete heading-like class.
*/
type HeadingLike interface {
	// Principal Methods
	GetClass() HeadingClassLike

	// Attribute Methods
	GetComment() string
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
	GetExpression() ExpressionLike
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
InspectClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inspect-clause-like class.
*/
type InspectClauseLike interface {
	// Principal Methods
	GetClass() InspectClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetRecipient() RecipientLike
	GetDelimiter2() string
	GetLocation() LocationLike
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
InvocationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete invocation-like class.
*/
type InvocationLike interface {
	// Principal Methods
	GetClass() InvocationClassLike

	// Attribute Methods
	GetAny() any
}

/*
InvokeClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete invoke-clause-like class.
*/
type InvokeClauseLike interface {
	// Principal Methods
	GetClass() InvokeClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetMethod() MethodLike
}

/*
ItemsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete items-like class.
*/
type ItemsLike interface {
	// Principal Methods
	GetClass() ItemsClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetContents() com.Sequential[ContentLike]
	GetDelimiter2() string
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
LexicalLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete lexical-like class.
*/
type LexicalLike interface {
	// Principal Methods
	GetClass() LexicalClassLike

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
LocalTransformationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete local-transformation-like class.
*/
type LocalTransformationLike interface {
	// Principal Methods
	GetClass() LocalTransformationClassLike

	// Attribute Methods
	GetAny() any
}

/*
LocationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete location-like class.
*/
type LocationLike interface {
	// Principal Methods
	GetClass() LocationClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
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
	GetExpression() ExpressionLike
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
MetadataLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete metadata-like class.
*/
type MetadataLike interface {
	// Principal Methods
	GetClass() MetadataClassLike

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
	GetInvocation() InvocationLike
	GetIdentifier2() string
	GetDelimiter1() string
	GetArguments() com.Sequential[ArgumentLike]
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
	GetLocation() LocationLike
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
	GetSymbol() string
	GetMatchingClauses() com.Sequential[MatchingClauseLike]
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
ParameterLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Principal Methods
	GetClass() ParameterClassLike

	// Attribute Methods
	GetSymbol() string
	GetDelimiter() string
	GetConstraint() ConstraintLike
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
	GetLines() com.Sequential[LineLike]
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
	GetMessage() MessageLike
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
	GetOptionalPrimitive1() PrimitiveLike
	GetDelimiter() string
	GetOptionalPrimitive2() PrimitiveLike
	GetRight() RightLike
}

/*
ReceiveClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete receive-clause-like class.
*/
type ReceiveClauseLike interface {
	// Principal Methods
	GetClass() ReceiveClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetRecipient() RecipientLike
	GetDelimiter2() string
	GetBag() BagLike
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
	GetDelimiter1() string
	GetMessage() MessageLike
	GetDelimiter2() string
	GetBag() BagLike
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
	GetCitation() CitationLike
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
	GetExpression() ExpressionLike
}

/*
ReversibleLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete reversible-like class.
*/
type ReversibleLike interface {
	// Principal Methods
	GetClass() ReversibleClassLike

	// Attribute Methods
	GetAny() any
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
	GetRecipient() RecipientLike
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
	GetMatchingClauses() com.Sequential[MatchingClauseLike]
}

/*
SendClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete send-clause-like class.
*/
type SendClauseLike interface {
	// Principal Methods
	GetClass() SendClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetMessage() MessageLike
	GetDelimiter2() string
	GetBag() BagLike
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
	GetAny() any
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
	GetIndexes() com.Sequential[IndexLike]
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
ThrowClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete throw-clause-like class.
*/
type ThrowClauseLike interface {
	// Principal Methods
	GetClass() ThrowClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetExpression() ExpressionLike
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
	GetExpression() ExpressionLike
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
	GetSymbol() string
	GetDelimiter3() string
	GetExpression() ExpressionLike
	GetDelimiter4() string
	GetProcedure() ProcedureLike
}

// ASPECT DECLARATIONS
