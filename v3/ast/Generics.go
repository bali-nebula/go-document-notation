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

func GenericsClass() GenericsClassLike {
	return genericsClass()
}

// Constructor Methods

func (c *genericsClass_) Generics(
	delimiter1 string,
	parameters com.Sequential[ParameterLike],
	delimiter2 string,
) GenericsLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &generics_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		parameters_: parameters,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *generics_) GetClass() GenericsClassLike {
	return genericsClass()
}

// Attribute Methods

func (v *generics_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *generics_) GetParameters() com.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *generics_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type generics_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	parameters_ com.Sequential[ParameterLike]
	delimiter2_ string
}

// Class Structure

type genericsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func genericsClass() *genericsClass_ {
	return genericsClassReference_
}

var genericsClassReference_ = &genericsClass_{
	// Initialize the class constants.
}
