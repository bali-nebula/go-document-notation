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

func LogicalOperatorClass() LogicalOperatorClassLike {
	return logicalOperatorClass()
}

// Constructor Methods

func (c *logicalOperatorClass_) LogicalOperator(
	any_ any,
) LogicalOperatorLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &logicalOperator_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *logicalOperator_) GetClass() LogicalOperatorClassLike {
	return logicalOperatorClass()
}

// Attribute Methods

func (v *logicalOperator_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type logicalOperator_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type logicalOperatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func logicalOperatorClass() *logicalOperatorClass_ {
	return logicalOperatorClassReference_
}

var logicalOperatorClassReference_ = &logicalOperatorClass_{
	// Initialize the class constants.
}
