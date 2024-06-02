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

package agent

const classTemplate_ = `<Notice>
<Header>
<Imports>
<Access>
<Class>
<Instance>`

const headerTemplate_ = `
package <Name>`

const importsTemplate_ = `
import (<Modules>)
`

const classAccessTemplate_ = `// CLASS ACCESS

// Reference
<Reference>

// Function
<Function>
`

const classReferenceTemplate_ = `
var <TargetName>Class = &<TargetName>Class_{
	// Any private class constants should be initialized here.
}`

const genericReferenceTemplate_ = `
var <TargetName>Class = map[string]any{}
var <TargetName>Mutex syn.Mutex`

const classFunctionTemplate_ = `
func <ClassName>() <ClassName>ClassLike {
	return <TargetName>Class
}`

const genericFunctionTemplate_ = `
func <ClassName>[<Parameters>]() <ClassName>ClassLike[<Arguments>] {
	// Generate the name of the bound class type.
	var result_ <ClassName>ClassLike[<Arguments>]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	<TargetName>Mutex.Lock()
	var value = <TargetName>Class[name]
	switch actual := value.(type) {
	case *<TargetName>Class_[<Arguments>]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &<TargetName>Class_[<Arguments>]{
			// Any private class constants should be initialized here.
		}
		<TargetName>Class[name] = result_
	}
	<TargetName>Mutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}`

const classMethodsTemplate_ = `// CLASS METHODS

// Target
<Target>
// Constants
<Constants>
// Constructors
<Constructors>
// Functions
<Functions>`

const classTargetTemplate_ = `
type <TargetName>Class_[<Parameters>] struct {<Constants>}`

const classConstantTemplate_ = `
	<ConstantName>_ <ConstantType>`

const classMethodTemplate_ = `
func (c *<TargetName>Class_[<Arguments>]) <MethodName>(<Parameters>)<ResultType> {<Body>}`

const constantBodyTemplate_ = `
	return c.<ConstantName>_
`

const constructorBodyTemplate_ = `
	return &<TargetName>_[<Arguments>]{<Assignments>}
`

const attributeAssignmentTemplate_ = `
		<AttributeName>_: <ParameterName>,`

const functionBodyTemplate_ = `
	var result_<ResultType>
	// TBA - Implement the function.
	return result_
`

const instanceMethodsTemplate_ = `// INSTANCE METHODS

// Target
<Target>
// Attributes
<Attributes>
<Abstractions>
// Public
<Methods>
// Private
`

const instanceTargetTemplate_ = `
type <TargetName>_[<Parameters>] struct {<Attributes>}`

const instanceAttributeTemplate_ = `
	<AttributeName>_ <AttributeType>`

const instanceAspectTemplate_ = `// <AspectName>
<Methods>`

const instanceMethodTemplate_ = `
func (v *<TargetName>_[<Arguments>]) <MethodName>(<Parameters>)<ResultType> {<Body>}`

const methodBodyTemplate_ = `
	// TBA - Implement the method.
`

const resultBodyTemplate_ = `
	var result_<ResultType>
	// TBA - Implement the method.
	return result_
`

const returnBodyTemplate_ = `
	// TBA - Implement the method.
	return
`

const getterBodyTemplate_ = `
	return v.<AttributeName>_
`

const setterBodyTemplate_ = `
	v.<AttributeName>_ = <ParameterName>
`

const modelTemplate_ = `/*
................................................................................
<Copyright>
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
Package "<name>" provides...

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package <name>

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
}`
