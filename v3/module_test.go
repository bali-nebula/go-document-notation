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

func TestParameterAccess(t *tes.T) {
	var document doc.DocumentLike
	var key = doc.Primitive(doc.Element("$type"))
	var parameter = doc.GetParameter(document, key)
	ass.Equal(t, nil, parameter)

	var source = `[ ]`
	document = doc.ParseSource(source)
	parameter = doc.GetParameter(document, key)
	ass.Equal(t, nil, parameter)

	source = `[ ]($type: "foo")`
	document = doc.ParseSource(source)
	parameter = doc.GetParameter(document, key)
	ass.Equal(t, "\"foo\"\n", doc.FormatDocument(parameter))

	source = `[ ]($type: "foo" $hype: /bar)`
	document = doc.ParseSource(source)
	key = doc.Primitive(doc.Element("$hype"))
	parameter = doc.GetParameter(document, key)
	ass.Equal(t, "/bar\n", doc.FormatDocument(parameter))
	parameter = doc.ParseSource("'A'")
	ass.True(t, doc.SetParameter(document, key, parameter))
	parameter = doc.GetParameter(document, key)
	ass.Equal(t, "'A'\n", doc.FormatDocument(parameter))
	key = doc.Primitive(doc.Element("$new"))
	parameter = doc.ParseSource("none")
	ass.True(t, doc.SetParameter(document, key, parameter))
	parameter = doc.GetParameter(document, key)
	ass.Equal(t, "none\n", doc.FormatDocument(parameter))
}

func TestAttributeAccess(t *tes.T) {
	var document doc.DocumentLike
	var index uti.Index = 1
	var attribute = doc.GetAttribute(document, index)
	ass.Equal(t, document, attribute)

	var source = `[ ]`
	document = doc.ParseSource(source)
	attribute = doc.GetAttribute(document)
	ass.Equal(t, nil, attribute)

	source = `[ ]`
	document = doc.ParseSource(source)
	attribute = doc.GetAttribute(document, index)
	ass.Equal(t, nil, attribute)
	attribute = doc.ParseSource("$new")
	index = 0 // Append a new attribute.
	ass.True(t, doc.SetAttribute(document, attribute, index))
	index = 1
	attribute = doc.GetAttribute(document, index)
	ass.Equal(t, "$new\n", doc.FormatDocument(attribute))
	doc.RemoveAttribute(document, index)
	attribute = doc.GetAttribute(document, index)
	ass.Equal(t, nil, attribute)

	source = `[
    $alpha
    $beta
    $gamma
]`
	document = doc.ParseSource(source)
	index = 1
	attribute = doc.GetAttribute(document, index)
	ass.Equal(t, "$alpha\n", doc.FormatDocument(attribute))
	index = 2
	attribute = doc.GetAttribute(document, index)
	ass.Equal(t, "$beta\n", doc.FormatDocument(attribute))
	index = 3
	attribute = doc.GetAttribute(document, index)
	ass.Equal(t, "$gamma\n", doc.FormatDocument(attribute))
	attribute = doc.ParseSource("$delta")
	ass.True(t, doc.SetAttribute(document, attribute, index))
	attribute = doc.GetAttribute(document, index)
	ass.Equal(t, "$delta\n", doc.FormatDocument(attribute))
	index = 0
	attribute = doc.ParseSource("$epsilon")
	ass.True(t, doc.SetAttribute(document, attribute, index))
	index = 4
	attribute = doc.GetAttribute(document, index)
	ass.Equal(t, "$epsilon\n", doc.FormatDocument(attribute))
	doc.RemoveAttribute(document, index)
	index = -1
	attribute = doc.GetAttribute(document, index)
	ass.Equal(t, "$delta\n", doc.FormatDocument(attribute))

	source = `[
    $alpha: "1"
    $beta: "2"
    $gamma: "3"
]`
	document = doc.ParseSource(source)
	var key = doc.Primitive(doc.Element("$alpha"))
	attribute = doc.GetAttribute(document, key)
	ass.Equal(t, "\"1\"\n", doc.FormatDocument(attribute))
	key = doc.Primitive(doc.Element("$beta"))
	attribute = doc.GetAttribute(document, key)
	ass.Equal(t, "\"2\"\n", doc.FormatDocument(attribute))
	key = doc.Primitive(doc.Element("$gamma"))
	attribute = doc.GetAttribute(document, key)
	ass.Equal(t, "\"3\"\n", doc.FormatDocument(attribute))
	attribute = doc.ParseSource("\"5\"")
	ass.True(t, doc.SetAttribute(document, attribute, key))
	attribute = doc.GetAttribute(document, key)
	ass.Equal(t, "\"5\"\n", doc.FormatDocument(attribute))
	doc.RemoveAttribute(document, key)
	key = doc.Primitive(doc.Element("$beta"))
	attribute = doc.GetAttribute(document, key)
	ass.Equal(t, "\"2\"\n", doc.FormatDocument(attribute))

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
	key = doc.Primitive(doc.Element("$items"))
	index = 2
	attribute = doc.GetAttribute(document, key, index)
	ass.Equal(t, "2\n", doc.FormatDocument(attribute))
	key = doc.Primitive(doc.Element("$attributes"))
	var key2 = doc.Primitive(doc.Element("$gamma"))
	attribute = doc.GetAttribute(document, key, key2)
	ass.Equal(t, "\"3\"\n", doc.FormatDocument(attribute))
	attribute = doc.ParseSource("\"5\"")
	ass.True(t, doc.SetAttribute(document, attribute, key, key2))
	attribute = doc.GetAttribute(document, key, key2)
	ass.Equal(t, "\"5\"\n", doc.FormatDocument(attribute))
	doc.RemoveAttribute(document, key, key2)
	key2 = doc.Primitive(doc.Element("$beta"))
	attribute = doc.GetAttribute(document, key, key2)
	ass.Equal(t, "\"2\"\n", doc.FormatDocument(attribute))

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
	index = 1
	var index2 uti.Index = 2
	attribute = doc.GetAttribute(document, index, index2)
	ass.Equal(t, "2\n", doc.FormatDocument(attribute))
	index = 2
	key2 = doc.Primitive(doc.Element("$beta"))
	attribute = doc.GetAttribute(document, index, key2)
	ass.Equal(t, "\"2\"\n", doc.FormatDocument(attribute))
	index = 1
	index2 = 3
	var key3 = doc.Primitive(doc.Element("~tau"))
	attribute = doc.GetAttribute(document, index, index2, key3)
	ass.Equal(t, "6.28\n", doc.FormatDocument(attribute))
	attribute = doc.ParseSource("~τ")
	ass.True(t, doc.SetAttribute(document, attribute, index, index2, key3))
	attribute = doc.GetAttribute(document, index, index2, key3)
	ass.Equal(t, "~τ\n", doc.FormatDocument(attribute))
	index = 2
	key2 = doc.Primitive(doc.Element("$alpha"))
	var index3 uti.Index = -1
	doc.RemoveAttribute(document, index, key2, index3)
	attribute = doc.GetAttribute(document, index, key2, index3)
	ass.Equal(t, "'b'\n", doc.FormatDocument(attribute))
	index = 3
	attribute = doc.GetAttribute(document, index)
	ass.Equal(t, nil, attribute)
	index = 2
	key2 = doc.Primitive(doc.Element("$delta"))
	attribute = doc.GetAttribute(document, index, key2)
	ass.Equal(t, nil, attribute)
}
