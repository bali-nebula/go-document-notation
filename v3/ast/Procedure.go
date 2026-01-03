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
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ProcedureClass() ProcedureClassLike {
	return procedureClass()
}

// Constructor Methods

func (c *procedureClass_) Procedure(
	delimiter1 string,
	statements com.Sequential[StatementLike],
	delimiter2 string,
) ProcedureLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(statements) {
		panic("The \"statements\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &procedure_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		statements_: statements,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *procedure_) GetClass() ProcedureClassLike {
	return procedureClass()
}

// Attribute Methods

func (v *procedure_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *procedure_) GetStatements() com.Sequential[StatementLike] {
	return v.statements_
}

func (v *procedure_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type procedure_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	statements_ com.Sequential[StatementLike]
	delimiter2_ string
}

// Class Structure

type procedureClass_ struct {
	// Declare the class constants.
}

// Class Reference

func procedureClass() *procedureClass_ {
	return procedureClassReference_
}

var procedureClassReference_ = &procedureClass_{
	// Initialize the class constants.
}
