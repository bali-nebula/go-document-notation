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

func RefinementClass() RefinementClassLike {
	return refinementClass()
}

// Constructor Methods

func (c *refinementClass_) Refinement(
	modifier ModifierLike,
	subject SubjectLike,
) RefinementLike {
	if uti.IsUndefined(modifier) {
		panic("The \"modifier\" attribute is required by this class.")
	}
	if uti.IsUndefined(subject) {
		panic("The \"subject\" attribute is required by this class.")
	}
	var instance = &refinement_{
		// Initialize the instance attributes.
		modifier_: modifier,
		subject_:  subject,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *refinement_) GetClass() RefinementClassLike {
	return refinementClass()
}

// Attribute Methods

func (v *refinement_) GetModifier() ModifierLike {
	return v.modifier_
}

func (v *refinement_) GetSubject() SubjectLike {
	return v.subject_
}

// PROTECTED INTERFACE

// Instance Structure

type refinement_ struct {
	// Declare the instance attributes.
	modifier_ ModifierLike
	subject_  SubjectLike
}

// Class Structure

type refinementClass_ struct {
	// Declare the class constants.
}

// Class Reference

func refinementClass() *refinementClass_ {
	return refinementClassReference_
}

var refinementClassReference_ = &refinementClass_{
	// Initialize the class constants.
}
