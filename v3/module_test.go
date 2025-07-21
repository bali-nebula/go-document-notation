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
	sts "strings"
	tes "testing"
)

const testDirectory = "./test/"

func TestParsingRoundtrips(t *tes.T) {
	var filenames = uti.ReadDirectory(testDirectory)
	for _, filename := range filenames {
		if sts.HasSuffix(filename, ".bali") {
			filename = testDirectory + filename
			fmt.Println(filename)
			var source = uti.ReadFile(filename)
			var document = doc.ParseSource(source)
			var formatted = doc.FormatDocument(document)
			ass.Equal(t, source, formatted)
		}
	}
}

func TestEntityAccess(t *tes.T) {
	var document doc.DocumentLike = nil
	var entity = doc.GetItem(document, 1)
	ass.Equal(t, document, entity)

	var source = `[ ]`
	document = doc.ParseSource(source)
	entity = doc.GetItem(document)
	ass.Equal(t, document, entity)

	source = `[ ]`
	document = doc.ParseSource(source)
	entity = doc.GetItem(document, 1)
	ass.Equal(t, nil, entity)

	source = `[
    $alpha
    $beta
    $gamma
]`
	document = doc.ParseSource(source)
	entity = doc.GetItem(document, 1)
	ass.Equal(t, "$alpha\n", doc.FormatDocument(entity))
	entity = doc.GetItem(document, 2)
	ass.Equal(t, "$beta\n", doc.FormatDocument(entity))
	entity = doc.GetItem(document, 3)
	ass.Equal(t, "$gamma\n", doc.FormatDocument(entity))

	source = `[
    $alpha: "1"
    $beta: "2"
    $gamma: "3"
]`
	document = doc.ParseSource(source)
	entity = doc.GetItem(document, "$alpha")
	ass.Equal(t, "\"1\"\n", doc.FormatDocument(entity))
	entity = doc.GetItem(document, "$beta")
	ass.Equal(t, "\"2\"\n", doc.FormatDocument(entity))
	entity = doc.GetItem(document, "$gamma")
	ass.Equal(t, "\"3\"\n", doc.FormatDocument(entity))

	source = `[
    $items: [
        1
        2
        3
    ]
    $attributes: [
        $alpha: "1"
        $beta: "2"
        $gamma: "3"
    ]
]`
	document = doc.ParseSource(source)
	entity = doc.GetItem(document, "$items", 2)
	ass.Equal(t, "2\n", doc.FormatDocument(entity))
	entity = doc.GetItem(document, "$attributes", "$gamma")
	ass.Equal(t, "\"3\"\n", doc.FormatDocument(entity))

	source = `[
    [
        1
        2
        [
            ~pi: 3.14
            ~tau: 6.28
        ]
    ]
    [
        $alpha: [
            'a'
            'b'
            'c'
        ]
        $beta: "2"
        $gamma: "3"
    ]
]`
	document = doc.ParseSource(source)
	entity = doc.GetItem(document, 1, 2)
	ass.Equal(t, "2\n", doc.FormatDocument(entity))
	entity = doc.GetItem(document, 2, "$beta")
	ass.Equal(t, "\"2\"\n", doc.FormatDocument(entity))
	entity = doc.GetItem(document, 1, 3, "~tau")
	ass.Equal(t, "6.28\n", doc.FormatDocument(entity))
	entity = doc.GetItem(document, 2, "$alpha", -1)
	ass.Equal(t, "'c'\n", doc.FormatDocument(entity))
	entity = doc.GetItem(document, 3)
	ass.Equal(t, nil, entity)
	entity = doc.GetItem(document, 2, "$delta")
	ass.Equal(t, nil, entity)
}
