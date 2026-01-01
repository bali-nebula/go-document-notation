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

func ItemsClass() ItemsClassLike {
	return itemsClass()
}

// Constructor Methods

func (c *itemsClass_) Items(
	delimiter1 string,
	contents com.Sequential[ContentLike],
	delimiter2 string,
) ItemsLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(contents) {
		panic("The \"contents\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &items_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		contents_:   contents,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *items_) GetClass() ItemsClassLike {
	return itemsClass()
}

// Attribute Methods

func (v *items_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *items_) GetContents() com.Sequential[ContentLike] {
	return v.contents_
}

func (v *items_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type items_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	contents_   com.Sequential[ContentLike]
	delimiter2_ string
}

// Class Structure

type itemsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func itemsClass() *itemsClass_ {
	return itemsClassReference_
}

var itemsClassReference_ = &itemsClass_{
	// Initialize the class constants.
}
