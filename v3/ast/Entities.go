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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func EntitiesClass() EntitiesClassLike {
	return entitiesClass()
}

// Constructor Methods

func (c *entitiesClass_) Entities(
	delimiter1 string,
	items fra.Sequential[ItemLike],
	delimiter2 string,
) EntitiesLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(items) {
		panic("The \"items\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &entities_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		items_:      items,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *entities_) GetClass() EntitiesClassLike {
	return entitiesClass()
}

// Attribute Methods

func (v *entities_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *entities_) GetItems() fra.Sequential[ItemLike] {
	return v.items_
}

func (v *entities_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type entities_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	items_      fra.Sequential[ItemLike]
	delimiter2_ string
}

// Class Structure

type entitiesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func entitiesClass() *entitiesClass_ {
	return entitiesClassReference_
}

var entitiesClassReference_ = &entitiesClass_{
	// Initialize the class constants.
}
