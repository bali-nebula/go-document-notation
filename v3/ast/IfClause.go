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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func IfClauseClass() IfClauseClassLike {
	return ifClauseClass()
}

// Constructor Methods

func (c *ifClauseClass_) IfClause(
	delimiter1 string,
	condition ConditionLike,
	delimiter2 string,
	procedure ProcedureLike,
) IfClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(condition) {
		panic("The \"condition\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(procedure) {
		panic("The \"procedure\" attribute is required by this class.")
	}
	var instance = &ifClause_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		condition_:  condition,
		delimiter2_: delimiter2,
		procedure_:  procedure,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *ifClause_) GetClass() IfClauseClassLike {
	return ifClauseClass()
}

// Attribute Methods

func (v *ifClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *ifClause_) GetCondition() ConditionLike {
	return v.condition_
}

func (v *ifClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *ifClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// PROTECTED INTERFACE

// Instance Structure

type ifClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	condition_  ConditionLike
	delimiter2_ string
	procedure_  ProcedureLike
}

// Class Structure

type ifClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func ifClauseClass() *ifClauseClass_ {
	return ifClauseClassReference_
}

var ifClauseClassReference_ = &ifClauseClass_{
	// Initialize the class constants.
}
