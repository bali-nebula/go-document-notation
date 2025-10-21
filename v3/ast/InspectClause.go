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

func InspectClauseClass() InspectClauseClassLike {
	return inspectClauseClass()
}

// Constructor Methods

func (c *inspectClauseClass_) InspectClause(
	delimiter1 string,
	recipient RecipientLike,
	delimiter2 string,
	location LocationLike,
) InspectClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(location) {
		panic("The \"location\" attribute is required by this class.")
	}
	var instance = &inspectClause_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		recipient_:  recipient,
		delimiter2_: delimiter2,
		location_:   location,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inspectClause_) GetClass() InspectClauseClassLike {
	return inspectClauseClass()
}

// Attribute Methods

func (v *inspectClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *inspectClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *inspectClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *inspectClause_) GetLocation() LocationLike {
	return v.location_
}

// PROTECTED INTERFACE

// Instance Structure

type inspectClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	recipient_  RecipientLike
	delimiter2_ string
	location_   LocationLike
}

// Class Structure

type inspectClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inspectClauseClass() *inspectClauseClass_ {
	return inspectClauseClassReference_
}

var inspectClauseClassReference_ = &inspectClauseClass_{
	// Initialize the class constants.
}
