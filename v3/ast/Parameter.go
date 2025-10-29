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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ParameterClass() ParameterClassLike {
	return parameterClass()
}

// Constructor Methods

func (c *parameterClass_) Parameter(
	symbol string,
	delimiter string,
	constraint ConstraintLike,
) ParameterLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(constraint) {
		panic("The \"constraint\" attribute is required by this class.")
	}
	var instance = &parameter_{
		// Initialize the instance attributes.
		symbol_:     symbol,
		delimiter_:  delimiter,
		constraint_: constraint,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parameter_) GetClass() ParameterClassLike {
	return parameterClass()
}

// Attribute Methods

func (v *parameter_) GetSymbol() string {
	return v.symbol_
}

func (v *parameter_) GetDelimiter() string {
	return v.delimiter_
}

func (v *parameter_) GetConstraint() ConstraintLike {
	return v.constraint_
}

// PROTECTED INTERFACE

// Instance Structure

type parameter_ struct {
	// Declare the instance attributes.
	symbol_     string
	delimiter_  string
	constraint_ ConstraintLike
}

// Class Structure

type parameterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func parameterClass() *parameterClass_ {
	return parameterClassReference_
}

var parameterClassReference_ = &parameterClass_{
	// Initialize the class constants.
}
