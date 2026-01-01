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

func AssignClauseClass() AssignClauseClassLike {
	return assignClauseClass()
}

// Constructor Methods

func (c *assignClauseClass_) AssignClause(
	delimiter string,
	recipient RecipientLike,
	assignment AssignmentLike,
	expression ExpressionLike,
) AssignClauseLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(assignment) {
		panic("The \"assignment\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &assignClause_{
		// Initialize the instance attributes.
		delimiter_:  delimiter,
		recipient_:  recipient,
		assignment_: assignment,
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *assignClause_) GetClass() AssignClauseClassLike {
	return assignClauseClass()
}

// Attribute Methods

func (v *assignClause_) GetDelimiter() string {
	return v.delimiter_
}

func (v *assignClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *assignClause_) GetAssignment() AssignmentLike {
	return v.assignment_
}

func (v *assignClause_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type assignClause_ struct {
	// Declare the instance attributes.
	delimiter_  string
	recipient_  RecipientLike
	assignment_ AssignmentLike
	expression_ ExpressionLike
}

// Class Structure

type assignClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func assignClauseClass() *assignClauseClass_ {
	return assignClauseClassReference_
}

var assignClauseClassReference_ = &assignClauseClass_{
	// Initialize the class constants.
}
