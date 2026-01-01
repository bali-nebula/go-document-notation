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

package ast

import (
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func DefineClauseClass() DefineClauseClassLike {
	return defineClauseClass()
}

// Constructor Methods

func (c *defineClauseClass_) DefineClause(
	delimiter1 string,
	constant ConstantLike,
	delimiter2 string,
	expression ExpressionLike,
) DefineClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(constant) {
		panic("The \"constant\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &defineClause_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		constant_:   constant,
		delimiter2_: delimiter2,
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *defineClause_) GetClass() DefineClauseClassLike {
	return defineClauseClass()
}

// Attribute Methods

func (v *defineClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *defineClause_) GetConstant() ConstantLike {
	return v.constant_
}

func (v *defineClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *defineClause_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type defineClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	constant_   ConstantLike
	delimiter2_ string
	expression_ ExpressionLike
}

// Class Structure

type defineClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func defineClauseClass() *defineClauseClass_ {
	return defineClauseClassReference_
}

var defineClauseClassReference_ = &defineClauseClass_{
	// Initialize the class constants.
}
