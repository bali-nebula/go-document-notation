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

func ComponentClass() ComponentClassLike {
	return componentClass()
}

// Constructor Methods

func (c *componentClass_) Component(
	literal LiteralLike,
	optionalGenerics GenericsLike,
	optionalNote string,
) ComponentLike {
	if uti.IsUndefined(literal) {
		panic("The \"literal\" attribute is required by this class.")
	}
	var instance = &component_{
		// Initialize the instance attributes.
		literal_:          literal,
		optionalGenerics_: optionalGenerics,
		optionalNote_:     optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *component_) GetClass() ComponentClassLike {
	return componentClass()
}

// Attribute Methods

func (v *component_) GetLiteral() LiteralLike {
	return v.literal_
}

func (v *component_) GetOptionalGenerics() GenericsLike {
	return v.optionalGenerics_
}

func (v *component_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type component_ struct {
	// Declare the instance attributes.
	literal_          LiteralLike
	optionalGenerics_ GenericsLike
	optionalNote_     string
}

// Class Structure

type componentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func componentClass() *componentClass_ {
	return componentClassReference_
}

var componentClassReference_ = &componentClass_{
	// Initialize the class constants.
}
