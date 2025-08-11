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

func LeftClass() LeftClassLike {
	return leftClass()
}

// Constructor Methods

func (c *leftClass_) Left(
	any_ any,
) LeftLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &left_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *left_) GetClass() LeftClassLike {
	return leftClass()
}

// Attribute Methods

func (v *left_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type left_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type leftClass_ struct {
	// Declare the class constants.
}

// Class Reference

func leftClass() *leftClass_ {
	return leftClassReference_
}

var leftClassReference_ = &leftClass_{
	// Initialize the class constants.
}
