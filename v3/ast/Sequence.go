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

func SequenceClass() SequenceClassLike {
	return sequenceClass()
}

// Constructor Methods

func (c *sequenceClass_) Sequence(
	expression ExpressionLike,
) SequenceLike {
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &sequence_{
		// Initialize the instance attributes.
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *sequence_) GetClass() SequenceClassLike {
	return sequenceClass()
}

// Attribute Methods

func (v *sequence_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type sequence_ struct {
	// Declare the instance attributes.
	expression_ ExpressionLike
}

// Class Structure

type sequenceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func sequenceClass() *sequenceClass_ {
	return sequenceClassReference_
}

var sequenceClassReference_ = &sequenceClass_{
	// Initialize the class constants.
}
