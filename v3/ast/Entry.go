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

func EntryClass() EntryClassLike {
	return entryClass()
}

// Constructor Methods

func (c *entryClass_) Entry(
	component ComponentLike,
	optionalNote string,
) EntryLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &entry_{
		// Initialize the instance attributes.
		component_:    component,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *entry_) GetClass() EntryClassLike {
	return entryClass()
}

// Attribute Methods

func (v *entry_) GetComponent() ComponentLike {
	return v.component_
}

func (v *entry_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type entry_ struct {
	// Declare the instance attributes.
	component_    ComponentLike
	optionalNote_ string
}

// Class Structure

type entryClass_ struct {
	// Declare the class constants.
}

// Class Reference

func entryClass() *entryClass_ {
	return entryClassReference_
}

var entryClassReference_ = &entryClass_{
	// Initialize the class constants.
}
