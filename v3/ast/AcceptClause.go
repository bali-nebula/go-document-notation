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

func AcceptClauseClass() AcceptClauseClassLike {
	return acceptClauseClass()
}

// Constructor Methods

func (c *acceptClauseClass_) AcceptClause(
	delimiter string,
	message MessageLike,
) AcceptClauseLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(message) {
		panic("The \"message\" attribute is required by this class.")
	}
	var instance = &acceptClause_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		message_:   message,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *acceptClause_) GetClass() AcceptClauseClassLike {
	return acceptClauseClass()
}

// Attribute Methods

func (v *acceptClause_) GetDelimiter() string {
	return v.delimiter_
}

func (v *acceptClause_) GetMessage() MessageLike {
	return v.message_
}

// PROTECTED INTERFACE

// Instance Structure

type acceptClause_ struct {
	// Declare the instance attributes.
	delimiter_ string
	message_   MessageLike
}

// Class Structure

type acceptClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func acceptClauseClass() *acceptClauseClass_ {
	return acceptClauseClassReference_
}

var acceptClauseClassReference_ = &acceptClauseClass_{
	// Initialize the class constants.
}
