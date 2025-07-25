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

func FlowControlClass() FlowControlClassLike {
	return flowControlClass()
}

// Constructor Methods

func (c *flowControlClass_) FlowControl(
	any_ any,
) FlowControlLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &flowControl_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *flowControl_) GetClass() FlowControlClassLike {
	return flowControlClass()
}

// Attribute Methods

func (v *flowControl_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type flowControl_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type flowControlClass_ struct {
	// Declare the class constants.
}

// Class Reference

func flowControlClass() *flowControlClass_ {
	return flowControlClassReference_
}

var flowControlClassReference_ = &flowControlClass_{
	// Initialize the class constants.
}
