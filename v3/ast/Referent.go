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

func ReferentClass() ReferentClassLike {
	return referentClass()
}

// Constructor Methods

func (c *referentClass_) Referent(
	delimiter string,
	indirect IndirectLike,
) ReferentLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(indirect) {
		panic("The \"indirect\" attribute is required by this class.")
	}
	var instance = &referent_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		indirect_:  indirect,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *referent_) GetClass() ReferentClassLike {
	return referentClass()
}

// Attribute Methods

func (v *referent_) GetDelimiter() string {
	return v.delimiter_
}

func (v *referent_) GetIndirect() IndirectLike {
	return v.indirect_
}

// PROTECTED INTERFACE

// Instance Structure

type referent_ struct {
	// Declare the instance attributes.
	delimiter_ string
	indirect_  IndirectLike
}

// Class Structure

type referentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func referentClass() *referentClass_ {
	return referentClassReference_
}

var referentClassReference_ = &referentClass_{
	// Initialize the class constants.
}
