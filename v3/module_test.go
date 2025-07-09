/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package module_test

import (
	fmt "fmt"
	doc "github.com/bali-nebula/go-document-notation/v3"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

const testDirectory = "./test/"

func TestParsingRoundtrips(t *tes.T) {
	var files, err = osx.ReadDir(testDirectory)
	if err != nil {
		panic("Could not find the ./test directory.")
	}

	for _, file := range files {
		var filename = testDirectory + file.Name()
		if sts.HasSuffix(filename, ".bali") {
			fmt.Println(filename)
			var source = uti.ReadFile(filename)
			var document = doc.ParseSource(source)
			var formatted = doc.FormatDocument(document)
			ass.Equal(t, source, formatted)
		}
	}
}
