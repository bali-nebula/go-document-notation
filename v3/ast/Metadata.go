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

func MetadataClass() MetadataClassLike {
	return metadataClass()
}

// Constructor Methods

func (c *metadataClass_) Metadata(
	any_ any,
) MetadataLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &metadata_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *metadata_) GetClass() MetadataClassLike {
	return metadataClass()
}

// Attribute Methods

func (v *metadata_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type metadata_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type metadataClass_ struct {
	// Declare the class constants.
}

// Class Reference

func metadataClass() *metadataClass_ {
	return metadataClassReference_
}

var metadataClassReference_ = &metadataClass_{
	// Initialize the class constants.
}
