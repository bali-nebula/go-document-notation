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

func SendClauseClass() SendClauseClassLike {
	return sendClauseClass()
}

// Constructor Methods

func (c *sendClauseClass_) SendClause(
	delimiter1 string,
	message MessageLike,
	delimiter2 string,
	location LocationLike,
) SendClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(message) {
		panic("The \"message\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(location) {
		panic("The \"location\" attribute is required by this class.")
	}
	var instance = &sendClause_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		message_:    message,
		delimiter2_: delimiter2,
		location_:   location,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *sendClause_) GetClass() SendClauseClassLike {
	return sendClauseClass()
}

// Attribute Methods

func (v *sendClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *sendClause_) GetMessage() MessageLike {
	return v.message_
}

func (v *sendClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *sendClause_) GetLocation() LocationLike {
	return v.location_
}

// PROTECTED INTERFACE

// Instance Structure

type sendClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	message_    MessageLike
	delimiter2_ string
	location_   LocationLike
}

// Class Structure

type sendClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func sendClauseClass() *sendClauseClass_ {
	return sendClauseClassReference_
}

var sendClauseClassReference_ = &sendClauseClass_{
	// Initialize the class constants.
}
