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

func ExplanationClass() ExplanationClassLike {
	return explanationClass()
}

// Constructor Methods

func (c *explanationClass_) Explanation(
	any_ any,
) ExplanationLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &explanation_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *explanation_) GetClass() ExplanationClassLike {
	return explanationClass()
}

// Attribute Methods

func (v *explanation_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type explanation_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type explanationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func explanationClass() *explanationClass_ {
	return explanationClassReference_
}

var explanationClassReference_ = &explanationClass_{
	// Initialize the class constants.
}
