/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent_test

import (
	age "github.com/craterdog/go-model-framework/v4/gcmn/agent"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestCreateModel(t *tes.T) {
	var generator = age.Generator().Make()
	var name = "example"
	var copyright string

	// Generate a new model template default copyright.
	var actual = generator.CreateModel(name, copyright)
	ass.Equal(t, defaultCopyright+template, actual)

	// Generate a new model template custom copyright.
	copyright = "Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved."
	actual = generator.CreateModel(name, copyright)
	ass.Equal(t, customCopyright+template, actual)
}

func TestGenerateClass(t *tes.T) {
	// Parse the source code for the class model.
	var parser = age.Parser().Make()
	var model = parser.ParseSource(gcmn)

	// Generate a new concrete class for the set class.
	var generator = age.Generator().Make()
	var actual = generator.GenerateClass(model, "set")
	ass.Equal(t, class, actual)
}

func TestRoundtrip(t *tes.T) {
	// Parse the source code for the class model.
	var parser = age.Parser().Make()
	var model = parser.ParseSource(gcmn)

	// Validate the class model.
	var validator = age.Validator().Make()
	validator.ValidateModel(model)

	// Format the class model.
	var formatter = age.Formatter().Make()
	var actual = formatter.FormatModel(model)
	ass.Equal(t, gcmn, actual)
}

const defaultCopyright = `/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................`

const customCopyright = `/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................`

const template = `
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
Package "example" provides...

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package example

// Types

/*
<ConstrainedType> is a constrained type representing...
*/
type <ConstrainedType> <primitiveType>

const (
	<1stValue> <ConstrainedType> = iota
	<2ndValue>
	<3rdValue>
	...
)
...

// Functionals

/*
<FunctionName>Function is a functional type that defines the signature for any
function that...
*/
type <FunctionName>Function func(<Parameters>) <AbstractType>
...

// Aspects

/*
<AspectName> is an aspect interface that defines a set of method signatures
that must be supported by each instance of a <aspect name> concrete class.
*/
type <AspectName> interface {
	// Methods
	<MethodName>(<Parameters>) <AbstractType>
	...
}
...

// Classes

/*
<ClassName>ClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
<class-name>-like concrete class.
*/
type <ClassName>ClassLike interface {
	// Constants
	<ConstantName>() <AbstractType>
	...

	// Constructors
	Make<FromContext>(<Parameters>) <ClassName>Like
	...

	// Functions
	<FunctionName>(<Parameters>) <AbstractType>
	...
}
...

// Instances

/*
<ClassName>Like is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a <class-name>-like concrete class.
*/
type <ClassName>Like interface {
	// Attributes
	Get<AttributeName>() <AttributeType>
	Set<AttributeName>(<attributeName> <AttributeType>)
	...

	// Abstractions
	<AspectName>
	...

	// Methods
	<MethodName>(<Parameters>) <AbstractType>
	...
}
...
`

const gcmn = `/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
Package "test" provides...

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package test

// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	CommentToken
	DelimiterToken
	EOFToken
	EOLToken
	IdentifierToken
	NoteToken
	SpaceToken
	TextToken
)

// Functionals

/*
RankingFunction[V any] is a functional type that defines the signature for any
function that can determine the relative ordering of two values. The result must
be one of the following:

	-1: The first value is less than the second value.
	 0: The first value is equal to the second value.
	 1: The first value is more than the second value.

The meaning of "less" and "more" is determined by the specific function that
implements this signature.
*/
type RankingFunction[V any] func(
	first V,
	second V,
) int

// Aspects

/*
Sequential[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of a sequential concrete class.
*/
type Sequential[V any] interface {
	// Methods
	AsArray() []V
	GetSize() int
	IsEmpty() bool
}

// Classes

/*
SetClassLike[V any] is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete set-like class.

The following functions are supported:

And() returns a new set containing the values that are both of the specified
sets.

Or() returns a new set containing the values that are in either of the specified
sets.

Sans() returns a new set containing the values that are in the first specified
set but not in the second specified set.

Xor() returns a new set containing the values that are in the first specified
set or the second specified set but not both.
*/
type SetClassLike[V any] interface {
	// Constants
	Ranker() RankingFunction

	// Constructors
	Make() SetLike[V]
	MakeFromArray(values []V) SetLike[V]
	MakeFromSequence(values Sequential[V]) SetLike[V]
	MakeFromSource(source string) SetLike[V]

	// Functions
	And(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
	Or(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
	Sans(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
	Xor(
		first SetLike[V],
		second SetLike[V],
	) SetLike[V]
}

// Instances

/*
SetLike[V any] is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete set-like class.  A set-like class maintains an ordered
sequence of values which can grow or shrink as needed.

This type is parameterized as follows:
  - V is any type of value.

The order of the values is determined by a configurable ranking function.
*/
type SetLike[V any] interface {
	// Attributes
	GetClass() SetClassLike[V]

	// Abstractions
	Sequential[V]

	// Methods
	AddValue(value V)
	AddValues(values Sequential[V])
	RemoveAll()
	RemoveValue(value V)
	RemoveValues(values Sequential[V])
}
`

const class = `/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package test

import (
	fmt "fmt"
	syn "sync"
)

// CLASS ACCESS

// Reference

var setClass = map[string]any{}
var setMutex syn.Mutex

// Function

func Set[V any]() SetClassLike[V] {
	// Generate the name of the bound class type.
	var result_ SetClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	setMutex.Lock()
	var value = setClass[name]
	switch actual := value.(type) {
	case *setClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &setClass_[V]{
			// Any private class constants should be initialized here.
		}
		setClass[name] = result_
	}
	setMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type setClass_[V any] struct {
	ranker_ RankingFunction
}

// Constants

func (c *setClass_[V]) Ranker() RankingFunction {
	return c.ranker_
}

// Constructors

func (c *setClass_[V]) Make() SetLike[V] {
	return &set_[V]{}
}

func (c *setClass_[V]) MakeFromArray(values []V) SetLike[V] {
	return &set_[V]{}
}

func (c *setClass_[V]) MakeFromSequence(values Sequential[V]) SetLike[V] {
	return &set_[V]{}
}

func (c *setClass_[V]) MakeFromSource(source string) SetLike[V] {
	return &set_[V]{}
}

// Functions

func (c *setClass_[V]) And(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

func (c *setClass_[V]) Or(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

func (c *setClass_[V]) Sans(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

func (c *setClass_[V]) Xor(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type set_[V any] struct {
	class_ SetClassLike[V]
}

// Attributes

func (v *set_[V]) GetClass() SetClassLike[V] {
	return v.class_
}

// Sequential[V]

func (v *set_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

func (v *set_[V]) AddValue(value V) {
	// TBA - Implement the method.
}

func (v *set_[V]) AddValues(values Sequential[V]) {
	// TBA - Implement the method.
}

func (v *set_[V]) RemoveAll() {
	// TBA - Implement the method.
}

func (v *set_[V]) RemoveValue(value V) {
	// TBA - Implement the method.
}

func (v *set_[V]) RemoveValues(values Sequential[V]) {
	// TBA - Implement the method.
}

// Private
`
