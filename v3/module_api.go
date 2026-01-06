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
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/bali-nebula/go-document-notation/wiki
*/
package module

import (
	ast "github.com/bali-nebula/go-document-notation/v3/ast"
	gra "github.com/bali-nebula/go-document-notation/v3/grammar"
	com "github.com/craterdog/go-essential-composites/v8"
)

// TYPE ALIASES

// Ast

type (
	AcceptClauseClassLike        = ast.AcceptClauseClassLike
	ArgumentClassLike            = ast.ArgumentClassLike
	ArithmeticClassLike          = ast.ArithmeticClassLike
	AssignClauseClassLike        = ast.AssignClauseClassLike
	AssignmentClassLike          = ast.AssignmentClassLike
	AssociationClassLike         = ast.AssociationClassLike
	AtLevelClassLike             = ast.AtLevelClassLike
	AttributesClassLike          = ast.AttributesClassLike
	BagClassLike                 = ast.BagClassLike
	BreakClauseClassLike         = ast.BreakClauseClassLike
	CheckoutClauseClassLike      = ast.CheckoutClauseClassLike
	CitationClassLike            = ast.CitationClassLike
	CollectionClassLike          = ast.CollectionClassLike
	ComparisonClassLike          = ast.ComparisonClassLike
	ComplementClassLike          = ast.ComplementClassLike
	ComponentClassLike           = ast.ComponentClassLike
	ConstantClassLike            = ast.ConstantClassLike
	ConstraintClassLike          = ast.ConstraintClassLike
	ContinueClauseClassLike      = ast.ContinueClauseClassLike
	DefineClauseClassLike        = ast.DefineClauseClassLike
	DiscardClauseClassLike       = ast.DiscardClauseClassLike
	DocumentClassLike            = ast.DocumentClassLike
	DraftClassLike               = ast.DraftClassLike
	ElementClassLike             = ast.ElementClassLike
	EmptyClassLike               = ast.EmptyClassLike
	EntityClassLike              = ast.EntityClassLike
	ExpressionClassLike          = ast.ExpressionClassLike
	FlowControlClassLike         = ast.FlowControlClassLike
	FunctionClassLike            = ast.FunctionClassLike
	GenericsClassLike            = ast.GenericsClassLike
	IfClauseClassLike            = ast.IfClauseClassLike
	IndexClassLike               = ast.IndexClassLike
	InspectClauseClassLike       = ast.InspectClauseClassLike
	InverseClassLike             = ast.InverseClassLike
	InversionClassLike           = ast.InversionClassLike
	InvocationClassLike          = ast.InvocationClassLike
	InvokeClauseClassLike        = ast.InvokeClauseClassLike
	ItemsClassLike               = ast.ItemsClassLike
	LeftClassLike                = ast.LeftClassLike
	LexicalClassLike             = ast.LexicalClassLike
	LiteralClassLike             = ast.LiteralClassLike
	LocalTransformationClassLike = ast.LocalTransformationClassLike
	LocationClassLike            = ast.LocationClassLike
	LogicalClassLike             = ast.LogicalClassLike
	MagnitudeClassLike           = ast.MagnitudeClassLike
	MainClauseClassLike          = ast.MainClauseClassLike
	MatchingClauseClassLike      = ast.MatchingClauseClassLike
	MessageClassLike             = ast.MessageClassLike
	MessageHandlingClassLike     = ast.MessageHandlingClassLike
	MethodClassLike              = ast.MethodClassLike
	NotarizeClauseClassLike      = ast.NotarizeClauseClassLike
	NumericalClassLike           = ast.NumericalClassLike
	OnClauseClassLike            = ast.OnClauseClassLike
	OperatorClassLike            = ast.OperatorClassLike
	ParameterClassLike           = ast.ParameterClassLike
	PrecedenceClassLike          = ast.PrecedenceClassLike
	PredicateClassLike           = ast.PredicateClassLike
	PrimitiveClassLike           = ast.PrimitiveClassLike
	ProcedureClassLike           = ast.ProcedureClassLike
	PublishClauseClassLike       = ast.PublishClauseClassLike
	RangeClassLike               = ast.RangeClassLike
	ReceiveClauseClassLike       = ast.ReceiveClauseClassLike
	RecipientClassLike           = ast.RecipientClassLike
	ReferenceClassLike           = ast.ReferenceClassLike
	ReferentClassLike            = ast.ReferentClassLike
	RejectClauseClassLike        = ast.RejectClauseClassLike
	RepositoryAccessClassLike    = ast.RepositoryAccessClassLike
	RetrieveClauseClassLike      = ast.RetrieveClauseClassLike
	ReturnClauseClassLike        = ast.ReturnClauseClassLike
	ReversibleClassLike          = ast.ReversibleClassLike
	RightClassLike               = ast.RightClassLike
	SaveClauseClassLike          = ast.SaveClauseClassLike
	SelectClauseClassLike        = ast.SelectClauseClassLike
	SendClauseClassLike          = ast.SendClauseClassLike
	SequenceClassLike            = ast.SequenceClassLike
	StatementClassLike           = ast.StatementClassLike
	SubcomponentClassLike        = ast.SubcomponentClassLike
	SubjectClassLike             = ast.SubjectClassLike
	ThrowClauseClassLike         = ast.ThrowClauseClassLike
	ValueClassLike               = ast.ValueClassLike
	VariableClassLike            = ast.VariableClassLike
	WhileClauseClassLike         = ast.WhileClauseClassLike
	WithClauseClassLike          = ast.WithClauseClassLike
)

type (
	AcceptClauseLike        = ast.AcceptClauseLike
	ArgumentLike            = ast.ArgumentLike
	ArithmeticLike          = ast.ArithmeticLike
	AssignClauseLike        = ast.AssignClauseLike
	AssignmentLike          = ast.AssignmentLike
	AssociationLike         = ast.AssociationLike
	AtLevelLike             = ast.AtLevelLike
	AttributesLike          = ast.AttributesLike
	BagLike                 = ast.BagLike
	BreakClauseLike         = ast.BreakClauseLike
	CheckoutClauseLike      = ast.CheckoutClauseLike
	CitationLike            = ast.CitationLike
	CollectionLike          = ast.CollectionLike
	ComparisonLike          = ast.ComparisonLike
	ComplementLike          = ast.ComplementLike
	ComponentLike           = ast.ComponentLike
	ConstantLike            = ast.ConstantLike
	ConstraintLike          = ast.ConstraintLike
	ContinueClauseLike      = ast.ContinueClauseLike
	DefineClauseLike        = ast.DefineClauseLike
	DiscardClauseLike       = ast.DiscardClauseLike
	DocumentLike            = ast.DocumentLike
	DraftLike               = ast.DraftLike
	ElementLike             = ast.ElementLike
	EmptyLike               = ast.EmptyLike
	EntityLike              = ast.EntityLike
	ExpressionLike          = ast.ExpressionLike
	FlowControlLike         = ast.FlowControlLike
	FunctionLike            = ast.FunctionLike
	GenericsLike            = ast.GenericsLike
	IfClauseLike            = ast.IfClauseLike
	IndexLike               = ast.IndexLike
	InspectClauseLike       = ast.InspectClauseLike
	InverseLike             = ast.InverseLike
	InversionLike           = ast.InversionLike
	InvocationLike          = ast.InvocationLike
	InvokeClauseLike        = ast.InvokeClauseLike
	ItemsLike               = ast.ItemsLike
	LeftLike                = ast.LeftLike
	LexicalLike             = ast.LexicalLike
	LiteralLike             = ast.LiteralLike
	LocalTransformationLike = ast.LocalTransformationLike
	LocationLike            = ast.LocationLike
	LogicalLike             = ast.LogicalLike
	MagnitudeLike           = ast.MagnitudeLike
	MainClauseLike          = ast.MainClauseLike
	MatchingClauseLike      = ast.MatchingClauseLike
	MessageLike             = ast.MessageLike
	MessageHandlingLike     = ast.MessageHandlingLike
	MethodLike              = ast.MethodLike
	NotarizeClauseLike      = ast.NotarizeClauseLike
	NumericalLike           = ast.NumericalLike
	OnClauseLike            = ast.OnClauseLike
	OperatorLike            = ast.OperatorLike
	ParameterLike           = ast.ParameterLike
	PrecedenceLike          = ast.PrecedenceLike
	PredicateLike           = ast.PredicateLike
	PrimitiveLike           = ast.PrimitiveLike
	ProcedureLike           = ast.ProcedureLike
	PublishClauseLike       = ast.PublishClauseLike
	RangeLike               = ast.RangeLike
	ReceiveClauseLike       = ast.ReceiveClauseLike
	RecipientLike           = ast.RecipientLike
	ReferenceLike           = ast.ReferenceLike
	ReferentLike            = ast.ReferentLike
	RejectClauseLike        = ast.RejectClauseLike
	RepositoryAccessLike    = ast.RepositoryAccessLike
	RetrieveClauseLike      = ast.RetrieveClauseLike
	ReturnClauseLike        = ast.ReturnClauseLike
	ReversibleLike          = ast.ReversibleLike
	RightLike               = ast.RightLike
	SaveClauseLike          = ast.SaveClauseLike
	SelectClauseLike        = ast.SelectClauseLike
	SendClauseLike          = ast.SendClauseLike
	SequenceLike            = ast.SequenceLike
	StatementLike           = ast.StatementLike
	SubcomponentLike        = ast.SubcomponentLike
	SubjectLike             = ast.SubjectLike
	ThrowClauseLike         = ast.ThrowClauseLike
	ValueLike               = ast.ValueLike
	VariableLike            = ast.VariableLike
	WhileClauseLike         = ast.WhileClauseLike
	WithClauseLike          = ast.WithClauseLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken       = gra.ErrorToken
	AngleToken       = gra.AngleToken
	BinaryToken      = gra.BinaryToken
	BooleanToken     = gra.BooleanToken
	BytecodeToken    = gra.BytecodeToken
	CommentToken     = gra.CommentToken
	DelimiterToken   = gra.DelimiterToken
	DurationToken    = gra.DurationToken
	GlyphToken       = gra.GlyphToken
	IdentifierToken  = gra.IdentifierToken
	MomentToken      = gra.MomentToken
	NameToken        = gra.NameToken
	NarrativeToken   = gra.NarrativeToken
	NewlineToken     = gra.NewlineToken
	NoteToken        = gra.NoteToken
	NumberToken      = gra.NumberToken
	PatternToken     = gra.PatternToken
	PercentageToken  = gra.PercentageToken
	ProbabilityToken = gra.ProbabilityToken
	QuoteToken       = gra.QuoteToken
	ResourceToken    = gra.ResourceToken
	SpaceToken       = gra.SpaceToken
	SymbolToken      = gra.SymbolToken
	TagToken         = gra.TagToken
	VersionToken     = gra.VersionToken
)

type (
	FormatterClassLike = gra.FormatterClassLike
	ParserClassLike    = gra.ParserClassLike
	ProcessorClassLike = gra.ProcessorClassLike
	ScannerClassLike   = gra.ScannerClassLike
	TokenClassLike     = gra.TokenClassLike
	ValidatorClassLike = gra.ValidatorClassLike
	VisitorClassLike   = gra.VisitorClassLike
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// CLASS ACCESSORS

// Ast

func AcceptClauseClass() AcceptClauseClassLike {
	return ast.AcceptClauseClass()
}

func AcceptClause(
	delimiter1 string,
	message ast.MessageLike,
	delimiter2 string,
	bag ast.BagLike,
) AcceptClauseLike {
	return AcceptClauseClass().AcceptClause(
		delimiter1,
		message,
		delimiter2,
		bag,
	)
}

func ArgumentClass() ArgumentClassLike {
	return ast.ArgumentClass()
}

func Argument(
	any_ any,
) ArgumentLike {
	return ArgumentClass().Argument(
		any_,
	)
}

func ArithmeticClass() ArithmeticClassLike {
	return ast.ArithmeticClass()
}

func Arithmetic(
	any_ any,
) ArithmeticLike {
	return ArithmeticClass().Arithmetic(
		any_,
	)
}

func AssignClauseClass() AssignClauseClassLike {
	return ast.AssignClauseClass()
}

func AssignClause(
	delimiter string,
	recipient ast.RecipientLike,
	assignment ast.AssignmentLike,
	expression ast.ExpressionLike,
) AssignClauseLike {
	return AssignClauseClass().AssignClause(
		delimiter,
		recipient,
		assignment,
		expression,
	)
}

func AssignmentClass() AssignmentClassLike {
	return ast.AssignmentClass()
}

func Assignment(
	any_ any,
) AssignmentLike {
	return AssignmentClass().Assignment(
		any_,
	)
}

func AssociationClass() AssociationClassLike {
	return ast.AssociationClass()
}

func Association(
	primitive ast.PrimitiveLike,
	delimiter string,
	component ast.ComponentLike,
) AssociationLike {
	return AssociationClass().Association(
		primitive,
		delimiter,
		component,
	)
}

func AtLevelClass() AtLevelClassLike {
	return ast.AtLevelClass()
}

func AtLevel(
	delimiter1 string,
	delimiter2 string,
	expression ast.ExpressionLike,
) AtLevelLike {
	return AtLevelClass().AtLevel(
		delimiter1,
		delimiter2,
		expression,
	)
}

func AttributesClass() AttributesClassLike {
	return ast.AttributesClass()
}

func Attributes(
	delimiter1 string,
	associations com.Sequential[ast.AssociationLike],
	delimiter2 string,
) AttributesLike {
	return AttributesClass().Attributes(
		delimiter1,
		associations,
		delimiter2,
	)
}

func BagClass() BagClassLike {
	return ast.BagClass()
}

func Bag(
	expression ast.ExpressionLike,
) BagLike {
	return BagClass().Bag(
		expression,
	)
}

func BreakClauseClass() BreakClauseClassLike {
	return ast.BreakClauseClass()
}

func BreakClause(
	delimiter1 string,
	delimiter2 string,
) BreakClauseLike {
	return BreakClauseClass().BreakClause(
		delimiter1,
		delimiter2,
	)
}

func CheckoutClauseClass() CheckoutClauseClassLike {
	return ast.CheckoutClauseClass()
}

func CheckoutClause(
	delimiter1 string,
	recipient ast.RecipientLike,
	optionalAtLevel ast.AtLevelLike,
	delimiter2 string,
	location ast.LocationLike,
) CheckoutClauseLike {
	return CheckoutClauseClass().CheckoutClause(
		delimiter1,
		recipient,
		optionalAtLevel,
		delimiter2,
		location,
	)
}

func CitationClass() CitationClassLike {
	return ast.CitationClass()
}

func Citation(
	expression ast.ExpressionLike,
) CitationLike {
	return CitationClass().Citation(
		expression,
	)
}

func CollectionClass() CollectionClassLike {
	return ast.CollectionClass()
}

func Collection(
	any_ any,
) CollectionLike {
	return CollectionClass().Collection(
		any_,
	)
}

func ComparisonClass() ComparisonClassLike {
	return ast.ComparisonClass()
}

func Comparison(
	any_ any,
) ComparisonLike {
	return ComparisonClass().Comparison(
		any_,
	)
}

func ComplementClass() ComplementClassLike {
	return ast.ComplementClass()
}

func Complement(
	delimiter string,
	reversible ast.ReversibleLike,
) ComplementLike {
	return ComplementClass().Complement(
		delimiter,
		reversible,
	)
}

func ComponentClass() ComponentClassLike {
	return ast.ComponentClass()
}

func Component(
	literal ast.LiteralLike,
	optionalGenerics ast.GenericsLike,
	optionalNote string,
) ComponentLike {
	return ComponentClass().Component(
		literal,
		optionalGenerics,
		optionalNote,
	)
}

func ConstantClass() ConstantClassLike {
	return ast.ConstantClass()
}

func Constant(
	symbol string,
) ConstantLike {
	return ConstantClass().Constant(
		symbol,
	)
}

func ConstraintClass() ConstraintClassLike {
	return ast.ConstraintClass()
}

func Constraint(
	entity ast.EntityLike,
	optionalGenerics ast.GenericsLike,
) ConstraintLike {
	return ConstraintClass().Constraint(
		entity,
		optionalGenerics,
	)
}

func ContinueClauseClass() ContinueClauseClassLike {
	return ast.ContinueClauseClass()
}

func ContinueClause(
	delimiter1 string,
	delimiter2 string,
) ContinueClauseLike {
	return ContinueClauseClass().ContinueClause(
		delimiter1,
		delimiter2,
	)
}

func DefineClauseClass() DefineClauseClassLike {
	return ast.DefineClauseClass()
}

func DefineClause(
	delimiter1 string,
	constant ast.ConstantLike,
	delimiter2 string,
	expression ast.ExpressionLike,
) DefineClauseLike {
	return DefineClauseClass().DefineClause(
		delimiter1,
		constant,
		delimiter2,
		expression,
	)
}

func DiscardClauseClass() DiscardClauseClassLike {
	return ast.DiscardClauseClass()
}

func DiscardClause(
	delimiter string,
	citation ast.CitationLike,
) DiscardClauseLike {
	return DiscardClauseClass().DiscardClause(
		delimiter,
		citation,
	)
}

func DocumentClass() DocumentClassLike {
	return ast.DocumentClass()
}

func Document(
	optionalComment string,
	component ast.ComponentLike,
) DocumentLike {
	return DocumentClass().Document(
		optionalComment,
		component,
	)
}

func DraftClass() DraftClassLike {
	return ast.DraftClass()
}

func Draft(
	expression ast.ExpressionLike,
) DraftLike {
	return DraftClass().Draft(
		expression,
	)
}

func ElementClass() ElementClassLike {
	return ast.ElementClass()
}

func Element(
	any_ any,
) ElementLike {
	return ElementClass().Element(
		any_,
	)
}

func EmptyClass() EmptyClassLike {
	return ast.EmptyClass()
}

func Empty(
	delimiter string,
) EmptyLike {
	return EmptyClass().Empty(
		delimiter,
	)
}

func EntityClass() EntityClassLike {
	return ast.EntityClass()
}

func Entity(
	any_ any,
) EntityLike {
	return EntityClass().Entity(
		any_,
	)
}

func ExpressionClass() ExpressionClassLike {
	return ast.ExpressionClass()
}

func Expression(
	subject ast.SubjectLike,
	predicates com.Sequential[ast.PredicateLike],
) ExpressionLike {
	return ExpressionClass().Expression(
		subject,
		predicates,
	)
}

func FlowControlClass() FlowControlClassLike {
	return ast.FlowControlClass()
}

func FlowControl(
	any_ any,
) FlowControlLike {
	return FlowControlClass().FlowControl(
		any_,
	)
}

func FunctionClass() FunctionClassLike {
	return ast.FunctionClass()
}

func Function(
	identifier string,
	delimiter1 string,
	arguments com.Sequential[ast.ArgumentLike],
	delimiter2 string,
) FunctionLike {
	return FunctionClass().Function(
		identifier,
		delimiter1,
		arguments,
		delimiter2,
	)
}

func GenericsClass() GenericsClassLike {
	return ast.GenericsClass()
}

func Generics(
	delimiter1 string,
	parameters com.Sequential[ast.ParameterLike],
	delimiter2 string,
) GenericsLike {
	return GenericsClass().Generics(
		delimiter1,
		parameters,
		delimiter2,
	)
}

func IfClauseClass() IfClauseClassLike {
	return ast.IfClauseClass()
}

func IfClause(
	delimiter1 string,
	expression ast.ExpressionLike,
	delimiter2 string,
	procedure ast.ProcedureLike,
) IfClauseLike {
	return IfClauseClass().IfClause(
		delimiter1,
		expression,
		delimiter2,
		procedure,
	)
}

func IndexClass() IndexClassLike {
	return ast.IndexClass()
}

func Index(
	any_ any,
) IndexLike {
	return IndexClass().Index(
		any_,
	)
}

func InspectClauseClass() InspectClauseClassLike {
	return ast.InspectClauseClass()
}

func InspectClause(
	delimiter1 string,
	recipient ast.RecipientLike,
	delimiter2 string,
	location ast.LocationLike,
) InspectClauseLike {
	return InspectClauseClass().InspectClause(
		delimiter1,
		recipient,
		delimiter2,
		location,
	)
}

func InverseClass() InverseClassLike {
	return ast.InverseClass()
}

func Inverse(
	any_ any,
) InverseLike {
	return InverseClass().Inverse(
		any_,
	)
}

func InversionClass() InversionClassLike {
	return ast.InversionClass()
}

func Inversion(
	inverse ast.InverseLike,
	numerical ast.NumericalLike,
) InversionLike {
	return InversionClass().Inversion(
		inverse,
		numerical,
	)
}

func InvocationClass() InvocationClassLike {
	return ast.InvocationClass()
}

func Invocation(
	any_ any,
) InvocationLike {
	return InvocationClass().Invocation(
		any_,
	)
}

func InvokeClauseClass() InvokeClauseClassLike {
	return ast.InvokeClauseClass()
}

func InvokeClause(
	delimiter string,
	method ast.MethodLike,
) InvokeClauseLike {
	return InvokeClauseClass().InvokeClause(
		delimiter,
		method,
	)
}

func ItemsClass() ItemsClassLike {
	return ast.ItemsClass()
}

func Items(
	delimiter1 string,
	components com.Sequential[ast.ComponentLike],
	delimiter2 string,
) ItemsLike {
	return ItemsClass().Items(
		delimiter1,
		components,
		delimiter2,
	)
}

func LeftClass() LeftClassLike {
	return ast.LeftClass()
}

func Left(
	any_ any,
) LeftLike {
	return LeftClass().Left(
		any_,
	)
}

func LexicalClass() LexicalClassLike {
	return ast.LexicalClass()
}

func Lexical(
	any_ any,
) LexicalLike {
	return LexicalClass().Lexical(
		any_,
	)
}

func LiteralClass() LiteralClassLike {
	return ast.LiteralClass()
}

func Literal(
	any_ any,
) LiteralLike {
	return LiteralClass().Literal(
		any_,
	)
}

func LocalTransformationClass() LocalTransformationClassLike {
	return ast.LocalTransformationClass()
}

func LocalTransformation(
	any_ any,
) LocalTransformationLike {
	return LocalTransformationClass().LocalTransformation(
		any_,
	)
}

func LocationClass() LocationClassLike {
	return ast.LocationClass()
}

func Location(
	expression ast.ExpressionLike,
) LocationLike {
	return LocationClass().Location(
		expression,
	)
}

func LogicalClass() LogicalClassLike {
	return ast.LogicalClass()
}

func Logical(
	any_ any,
) LogicalLike {
	return LogicalClass().Logical(
		any_,
	)
}

func MagnitudeClass() MagnitudeClassLike {
	return ast.MagnitudeClass()
}

func Magnitude(
	delimiter1 string,
	expression ast.ExpressionLike,
	delimiter2 string,
) MagnitudeLike {
	return MagnitudeClass().Magnitude(
		delimiter1,
		expression,
		delimiter2,
	)
}

func MainClauseClass() MainClauseClassLike {
	return ast.MainClauseClass()
}

func MainClause(
	any_ any,
) MainClauseLike {
	return MainClauseClass().MainClause(
		any_,
	)
}

func MatchingClauseClass() MatchingClauseClassLike {
	return ast.MatchingClauseClass()
}

func MatchingClause(
	delimiter1 string,
	expression ast.ExpressionLike,
	delimiter2 string,
	procedure ast.ProcedureLike,
) MatchingClauseLike {
	return MatchingClauseClass().MatchingClause(
		delimiter1,
		expression,
		delimiter2,
		procedure,
	)
}

func MessageClass() MessageClassLike {
	return ast.MessageClass()
}

func Message(
	expression ast.ExpressionLike,
) MessageLike {
	return MessageClass().Message(
		expression,
	)
}

func MessageHandlingClass() MessageHandlingClassLike {
	return ast.MessageHandlingClass()
}

func MessageHandling(
	any_ any,
) MessageHandlingLike {
	return MessageHandlingClass().MessageHandling(
		any_,
	)
}

func MethodClass() MethodClassLike {
	return ast.MethodClass()
}

func Method(
	identifier1 string,
	invocation ast.InvocationLike,
	identifier2 string,
	delimiter1 string,
	arguments com.Sequential[ast.ArgumentLike],
	delimiter2 string,
) MethodLike {
	return MethodClass().Method(
		identifier1,
		invocation,
		identifier2,
		delimiter1,
		arguments,
		delimiter2,
	)
}

func NotarizeClauseClass() NotarizeClauseClassLike {
	return ast.NotarizeClauseClass()
}

func NotarizeClause(
	delimiter1 string,
	draft ast.DraftLike,
	delimiter2 string,
	location ast.LocationLike,
) NotarizeClauseLike {
	return NotarizeClauseClass().NotarizeClause(
		delimiter1,
		draft,
		delimiter2,
		location,
	)
}

func NumericalClass() NumericalClassLike {
	return ast.NumericalClass()
}

func Numerical(
	any_ any,
) NumericalLike {
	return NumericalClass().Numerical(
		any_,
	)
}

func OnClauseClass() OnClauseClassLike {
	return ast.OnClauseClass()
}

func OnClause(
	delimiter string,
	symbol string,
	matchingClauses com.Sequential[ast.MatchingClauseLike],
) OnClauseLike {
	return OnClauseClass().OnClause(
		delimiter,
		symbol,
		matchingClauses,
	)
}

func OperatorClass() OperatorClassLike {
	return ast.OperatorClass()
}

func Operator(
	any_ any,
) OperatorLike {
	return OperatorClass().Operator(
		any_,
	)
}

func ParameterClass() ParameterClassLike {
	return ast.ParameterClass()
}

func Parameter(
	symbol string,
	delimiter string,
	constraint ast.ConstraintLike,
) ParameterLike {
	return ParameterClass().Parameter(
		symbol,
		delimiter,
		constraint,
	)
}

func PrecedenceClass() PrecedenceClassLike {
	return ast.PrecedenceClass()
}

func Precedence(
	delimiter1 string,
	expression ast.ExpressionLike,
	delimiter2 string,
) PrecedenceLike {
	return PrecedenceClass().Precedence(
		delimiter1,
		expression,
		delimiter2,
	)
}

func PredicateClass() PredicateClassLike {
	return ast.PredicateClass()
}

func Predicate(
	operator ast.OperatorLike,
	expression ast.ExpressionLike,
) PredicateLike {
	return PredicateClass().Predicate(
		operator,
		expression,
	)
}

func PrimitiveClass() PrimitiveClassLike {
	return ast.PrimitiveClass()
}

func Primitive(
	any_ any,
) PrimitiveLike {
	return PrimitiveClass().Primitive(
		any_,
	)
}

func ProcedureClass() ProcedureClassLike {
	return ast.ProcedureClass()
}

func Procedure(
	delimiter1 string,
	statements com.Sequential[ast.StatementLike],
	delimiter2 string,
) ProcedureLike {
	return ProcedureClass().Procedure(
		delimiter1,
		statements,
		delimiter2,
	)
}

func PublishClauseClass() PublishClauseClassLike {
	return ast.PublishClauseClass()
}

func PublishClause(
	delimiter string,
	message ast.MessageLike,
) PublishClauseLike {
	return PublishClauseClass().PublishClause(
		delimiter,
		message,
	)
}

func RangeClass() RangeClassLike {
	return ast.RangeClass()
}

func Range(
	left ast.LeftLike,
	optionalPrimitive1 ast.PrimitiveLike,
	delimiter string,
	optionalPrimitive2 ast.PrimitiveLike,
	right ast.RightLike,
) RangeLike {
	return RangeClass().Range(
		left,
		optionalPrimitive1,
		delimiter,
		optionalPrimitive2,
		right,
	)
}

func ReceiveClauseClass() ReceiveClauseClassLike {
	return ast.ReceiveClauseClass()
}

func ReceiveClause(
	delimiter1 string,
	recipient ast.RecipientLike,
	delimiter2 string,
	bag ast.BagLike,
) ReceiveClauseLike {
	return ReceiveClauseClass().ReceiveClause(
		delimiter1,
		recipient,
		delimiter2,
		bag,
	)
}

func RecipientClass() RecipientClassLike {
	return ast.RecipientClass()
}

func Recipient(
	any_ any,
) RecipientLike {
	return RecipientClass().Recipient(
		any_,
	)
}

func ReferenceClass() ReferenceClassLike {
	return ast.ReferenceClass()
}

func Reference(
	any_ any,
) ReferenceLike {
	return ReferenceClass().Reference(
		any_,
	)
}

func ReferentClass() ReferentClassLike {
	return ast.ReferentClass()
}

func Referent(
	delimiter string,
	reference ast.ReferenceLike,
) ReferentLike {
	return ReferentClass().Referent(
		delimiter,
		reference,
	)
}

func RejectClauseClass() RejectClauseClassLike {
	return ast.RejectClauseClass()
}

func RejectClause(
	delimiter1 string,
	message ast.MessageLike,
	delimiter2 string,
	bag ast.BagLike,
) RejectClauseLike {
	return RejectClauseClass().RejectClause(
		delimiter1,
		message,
		delimiter2,
		bag,
	)
}

func RepositoryAccessClass() RepositoryAccessClassLike {
	return ast.RepositoryAccessClass()
}

func RepositoryAccess(
	any_ any,
) RepositoryAccessLike {
	return RepositoryAccessClass().RepositoryAccess(
		any_,
	)
}

func RetrieveClauseClass() RetrieveClauseClassLike {
	return ast.RetrieveClauseClass()
}

func RetrieveClause(
	delimiter1 string,
	recipient ast.RecipientLike,
	delimiter2 string,
	citation ast.CitationLike,
) RetrieveClauseLike {
	return RetrieveClauseClass().RetrieveClause(
		delimiter1,
		recipient,
		delimiter2,
		citation,
	)
}

func ReturnClauseClass() ReturnClauseClassLike {
	return ast.ReturnClauseClass()
}

func ReturnClause(
	delimiter string,
	expression ast.ExpressionLike,
) ReturnClauseLike {
	return ReturnClauseClass().ReturnClause(
		delimiter,
		expression,
	)
}

func ReversibleClass() ReversibleClassLike {
	return ast.ReversibleClass()
}

func Reversible(
	any_ any,
) ReversibleLike {
	return ReversibleClass().Reversible(
		any_,
	)
}

func RightClass() RightClassLike {
	return ast.RightClass()
}

func Right(
	any_ any,
) RightLike {
	return RightClass().Right(
		any_,
	)
}

func SaveClauseClass() SaveClauseClassLike {
	return ast.SaveClauseClass()
}

func SaveClause(
	delimiter1 string,
	draft ast.DraftLike,
	delimiter2 string,
	recipient ast.RecipientLike,
) SaveClauseLike {
	return SaveClauseClass().SaveClause(
		delimiter1,
		draft,
		delimiter2,
		recipient,
	)
}

func SelectClauseClass() SelectClauseClassLike {
	return ast.SelectClauseClass()
}

func SelectClause(
	delimiter string,
	expression ast.ExpressionLike,
	matchingClauses com.Sequential[ast.MatchingClauseLike],
) SelectClauseLike {
	return SelectClauseClass().SelectClause(
		delimiter,
		expression,
		matchingClauses,
	)
}

func SendClauseClass() SendClauseClassLike {
	return ast.SendClauseClass()
}

func SendClause(
	delimiter1 string,
	message ast.MessageLike,
	delimiter2 string,
	bag ast.BagLike,
) SendClauseLike {
	return SendClauseClass().SendClause(
		delimiter1,
		message,
		delimiter2,
		bag,
	)
}

func SequenceClass() SequenceClassLike {
	return ast.SequenceClass()
}

func Sequence(
	any_ any,
) SequenceLike {
	return SequenceClass().Sequence(
		any_,
	)
}

func StatementClass() StatementClassLike {
	return ast.StatementClass()
}

func Statement(
	optionalComment string,
	mainClause ast.MainClauseLike,
	optionalOnClause ast.OnClauseLike,
) StatementLike {
	return StatementClass().Statement(
		optionalComment,
		mainClause,
		optionalOnClause,
	)
}

func SubcomponentClass() SubcomponentClassLike {
	return ast.SubcomponentClass()
}

func Subcomponent(
	identifier string,
	delimiter1 string,
	indexes com.Sequential[ast.IndexLike],
	delimiter2 string,
) SubcomponentLike {
	return SubcomponentClass().Subcomponent(
		identifier,
		delimiter1,
		indexes,
		delimiter2,
	)
}

func SubjectClass() SubjectClassLike {
	return ast.SubjectClass()
}

func Subject(
	any_ any,
) SubjectLike {
	return SubjectClass().Subject(
		any_,
	)
}

func ThrowClauseClass() ThrowClauseClassLike {
	return ast.ThrowClauseClass()
}

func ThrowClause(
	delimiter string,
	expression ast.ExpressionLike,
) ThrowClauseLike {
	return ThrowClauseClass().ThrowClause(
		delimiter,
		expression,
	)
}

func ValueClass() ValueClassLike {
	return ast.ValueClass()
}

func Value(
	identifier string,
) ValueLike {
	return ValueClass().Value(
		identifier,
	)
}

func VariableClass() VariableClassLike {
	return ast.VariableClass()
}

func Variable(
	symbol string,
) VariableLike {
	return VariableClass().Variable(
		symbol,
	)
}

func WhileClauseClass() WhileClauseClassLike {
	return ast.WhileClauseClass()
}

func WhileClause(
	delimiter1 string,
	expression ast.ExpressionLike,
	delimiter2 string,
	procedure ast.ProcedureLike,
) WhileClauseLike {
	return WhileClauseClass().WhileClause(
		delimiter1,
		expression,
		delimiter2,
		procedure,
	)
}

func WithClauseClass() WithClauseClassLike {
	return ast.WithClauseClass()
}

func WithClause(
	delimiter1 string,
	delimiter2 string,
	symbol string,
	delimiter3 string,
	expression ast.ExpressionLike,
	delimiter4 string,
	procedure ast.ProcedureLike,
) WithClauseLike {
	return WithClauseClass().WithClause(
		delimiter1,
		delimiter2,
		symbol,
		delimiter3,
		expression,
		delimiter4,
		procedure,
	)
}

// Grammar

func FormatterClass() FormatterClassLike {
	return gra.FormatterClass()
}

func Formatter() FormatterLike {
	return FormatterClass().Formatter()
}

func ParserClass() ParserClassLike {
	return gra.ParserClass()
}

func Parser() ParserLike {
	return ParserClass().Parser()
}

func ProcessorClass() ProcessorClassLike {
	return gra.ProcessorClass()
}

func Processor() ProcessorLike {
	return ProcessorClass().Processor()
}

func ScannerClass() ScannerClassLike {
	return gra.ScannerClass()
}

func Scanner(
	source string,
	tokens com.QueueLike[gra.TokenLike],
) ScannerLike {
	return ScannerClass().Scanner(
		source,
		tokens,
	)
}

func TokenClass() TokenClassLike {
	return gra.TokenClass()
}

func Token(
	line uint,
	position uint,
	type_ gra.TokenType,
	value string,
) TokenLike {
	return TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

func ValidatorClass() ValidatorClassLike {
	return gra.ValidatorClass()
}

func Validator() ValidatorLike {
	return ValidatorClass().Validator()
}

func VisitorClass() VisitorClassLike {
	return gra.VisitorClass()
}

func Visitor(
	processor gra.Methodical,
) VisitorLike {
	return VisitorClass().Visitor(
		processor,
	)
}

// GLOBAL FUNCTIONS

func FormatDocument(
	document DocumentLike,
) string {
	var formatter = Formatter()
	return formatter.FormatDocument(document)
}

func ParseSource(
	source string,
) DocumentLike {
	var parser = Parser()
	return parser.ParseSource(source)
}
