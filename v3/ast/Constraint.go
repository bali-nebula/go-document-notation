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

func ConstraintClass() ConstraintClassLike {
	return constraintClass()
}

// Constructor Methods

func (c *constraintClass_) Constraint(
	type_ TypeLike,
	optionalParameterization ParameterizationLike,
) ConstraintLike {
	if uti.IsUndefined(type_) {
		panic("The \"type\" attribute is required by this class.")
	}
	var instance = &constraint_{
		// Initialize the instance attributes.
		type_:                     type_,
		optionalParameterization_: optionalParameterization,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *constraint_) GetClass() ConstraintClassLike {
	return constraintClass()
}

// Attribute Methods

func (v *constraint_) GetType() TypeLike {
	return v.type_
}

func (v *constraint_) GetOptionalParameterization() ParameterizationLike {
	return v.optionalParameterization_
}

// PROTECTED INTERFACE

// Instance Structure

type constraint_ struct {
	// Declare the instance attributes.
	type_                     TypeLike
	optionalParameterization_ ParameterizationLike
}

// Class Structure

type constraintClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constraintClass() *constraintClass_ {
	return constraintClassReference_
}

var constraintClassReference_ = &constraintClass_{
	// Initialize the class constants.
}
