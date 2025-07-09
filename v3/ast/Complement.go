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

func ComplementClass() ComplementClassLike {
	return complementClass()
}

// Constructor Methods

func (c *complementClass_) Complement(
	delimiter string,
	logical LogicalLike,
) ComplementLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(logical) {
		panic("The \"logical\" attribute is required by this class.")
	}
	var instance = &complement_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		logical_:   logical,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *complement_) GetClass() ComplementClassLike {
	return complementClass()
}

// Attribute Methods

func (v *complement_) GetDelimiter() string {
	return v.delimiter_
}

func (v *complement_) GetLogical() LogicalLike {
	return v.logical_
}

// PROTECTED INTERFACE

// Instance Structure

type complement_ struct {
	// Declare the instance attributes.
	delimiter_ string
	logical_   LogicalLike
}

// Class Structure

type complementClass_ struct {
	// Declare the class constants.
}

// Class Reference

func complementClass() *complementClass_ {
	return complementClassReference_
}

var complementClassReference_ = &complementClass_{
	// Initialize the class constants.
}
