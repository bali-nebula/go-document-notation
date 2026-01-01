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

func LocalTransformationClass() LocalTransformationClassLike {
	return localTransformationClass()
}

// Constructor Methods

func (c *localTransformationClass_) LocalTransformation(
	any_ any,
) LocalTransformationLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &localTransformation_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *localTransformation_) GetClass() LocalTransformationClassLike {
	return localTransformationClass()
}

// Attribute Methods

func (v *localTransformation_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type localTransformation_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type localTransformationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func localTransformationClass() *localTransformationClass_ {
	return localTransformationClassReference_
}

var localTransformationClassReference_ = &localTransformationClass_{
	// Initialize the class constants.
}
