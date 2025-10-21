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

func ReceiveClauseClass() ReceiveClauseClassLike {
	return receiveClauseClass()
}

// Constructor Methods

func (c *receiveClauseClass_) ReceiveClause(
	delimiter1 string,
	recipient RecipientLike,
	delimiter2 string,
	bag BagLike,
) ReceiveClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(bag) {
		panic("The \"bag\" attribute is required by this class.")
	}
	var instance = &receiveClause_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		recipient_:  recipient,
		delimiter2_: delimiter2,
		bag_:        bag,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *receiveClause_) GetClass() ReceiveClauseClassLike {
	return receiveClauseClass()
}

// Attribute Methods

func (v *receiveClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *receiveClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *receiveClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *receiveClause_) GetBag() BagLike {
	return v.bag_
}

// PROTECTED INTERFACE

// Instance Structure

type receiveClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	recipient_  RecipientLike
	delimiter2_ string
	bag_        BagLike
}

// Class Structure

type receiveClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func receiveClauseClass() *receiveClauseClass_ {
	return receiveClauseClassReference_
}

var receiveClauseClassReference_ = &receiveClauseClass_{
	// Initialize the class constants.
}
