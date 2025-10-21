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
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func LexicalClass() LexicalClassLike {
	return lexicalClass()
}

// Constructor Methods

func (c *lexicalClass_) Lexical(
	any_ any,
) LexicalLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &lexical_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *lexical_) GetClass() LexicalClassLike {
	return lexicalClass()
}

// Attribute Methods

func (v *lexical_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type lexical_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type lexicalClass_ struct {
	// Declare the class constants.
}

// Class Reference

func lexicalClass() *lexicalClass_ {
	return lexicalClassReference_
}

var lexicalClassReference_ = &lexicalClass_{
	// Initialize the class constants.
}
