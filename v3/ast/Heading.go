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

func HeadingClass() HeadingClassLike {
	return headingClass()
}

// Constructor Methods

func (c *headingClass_) Heading(
	comment string,
) HeadingLike {
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	var instance = &heading_{
		// Initialize the instance attributes.
		comment_: comment,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *heading_) GetClass() HeadingClassLike {
	return headingClass()
}

// Attribute Methods

func (v *heading_) GetComment() string {
	return v.comment_
}

// PROTECTED INTERFACE

// Instance Structure

type heading_ struct {
	// Declare the instance attributes.
	comment_ string
}

// Class Structure

type headingClass_ struct {
	// Declare the class constants.
}

// Class Reference

func headingClass() *headingClass_ {
	return headingClassReference_
}

var headingClassReference_ = &headingClass_{
	// Initialize the class constants.
}
