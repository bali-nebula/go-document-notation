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

func ItemClass() ItemClassLike {
	return itemClass()
}

// Constructor Methods

func (c *itemClass_) Item(
	document DocumentLike,
) ItemLike {
	if uti.IsUndefined(document) {
		panic("The \"document\" attribute is required by this class.")
	}
	var instance = &item_{
		// Initialize the instance attributes.
		document_: document,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *item_) GetClass() ItemClassLike {
	return itemClass()
}

// Attribute Methods

func (v *item_) GetDocument() DocumentLike {
	return v.document_
}

// PROTECTED INTERFACE

// Instance Structure

type item_ struct {
	// Declare the instance attributes.
	document_ DocumentLike
}

// Class Structure

type itemClass_ struct {
	// Declare the class constants.
}

// Class Reference

func itemClass() *itemClass_ {
	return itemClassReference_
}

var itemClassReference_ = &itemClass_{
	// Initialize the class constants.
}
