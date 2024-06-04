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

const classTemplate_ = `<Notice><Header><Imports><Access><Class><Instance>`

const headerTemplate_ = `
package <Name>
`

const importsTemplate_ = `
import (<Modules>)
`

const classAccessTemplate_ = `
// CLASS ACCESS

// Reference
<Reference>
// Function
<Function>`

const classReferenceTemplate_ = `
var <TargetName>Class = &<TargetName>Class_{
	// Initialize class constants.
}
`

const genericReferenceTemplate_ = `
var <TargetName>Class = map[string]any{}
var <TargetName>Mutex syn.Mutex
`

const classFunctionTemplate_ = `
func <ClassName>() <ClassName>ClassLike {
	return <TargetName>Class
}
`

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
			// Initialize class constants.
		}
		<TargetName>Class[name] = result_
	}
	<TargetName>Mutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}
`

const classMethodsTemplate_ = `
// CLASS METHODS

// Target
<Target><Constants><Constructors><Functions>`

const classTargetTemplate_ = `
type <TargetName>Class_[<Parameters>] struct {
	// Define class constants.<Constants>
}
`

const classConstantTemplate_ = `
	<ConstantName>_ <ConstantType>`

const classMethodTemplate_ = `
func (c *<TargetName>Class_[<Arguments>]) <MethodName>(<Parameters>)<ResultType> {<Body>}
`

const constantBodyTemplate_ = `
	return c.<ConstantName>_
`

const constructorBodyTemplate_ = `
	return &<TargetName>_[<Arguments>]{
		// Initialize instance attributes.
		class_: c,<Assignments>
	}
`

const attributeAssignmentTemplate_ = `
		<AttributeName>_: <ParameterName>,`

const functionBodyTemplate_ = `
	var result_<ResultType>
	// TBA - Implement the function.
	return result_
`

const instanceMethodsTemplate_ = `
// INSTANCE METHODS

// Target
<Target><Attributes><Abstractions><Methods>
// Private
`

const instanceTargetTemplate_ = `
type <TargetName>_[<Parameters>] struct {
	// Define instance attributes.<Attributes>
}
`

const instanceAttributeTemplate_ = `
	<AttributeName>_ <AttributeType>`

const instanceAspectTemplate_ = `
// <AspectName>
<Methods>`

const instanceMethodTemplate_ = `
func (v *<TargetName>_[<Arguments>]) <MethodName>(<Parameters>)<ResultType> {<Body>}
`

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
Units is a constrained type representing the possible units for an angle.
*/
type Units uint8

const (
	Degrees Units = iota
	Radians
	Gradians
)

// Functionals

/*
TrigonometricFunction is a functional type that defines the signature for any
trigonometric function.
*/
type TrigonometricFunction func(angle AngleLike) float64

// Aspects

/*
Angular is an aspect interface that defines a set of method signatures that
must be supported by each instance of an angular concrete class.
*/
type Angular interface {
	// Methods
	AsNormalized() AngleLike
	InUnits(units Units) float64
}

// Classes

/*
AngleClassLike is a class interface that defines the set of class constants,
constructors and functions that must be supported by each angle-like concrete
class.
*/
type AngleClassLike interface {
	// Constants
	Pi() AngleLike
	Tau() AngleLike

	// Constructors
	MakeWithValue(value float64) AngleLike
	MakeFromString(value string) AngleLike

	// Functions
	Apply(function TrigonometricFunction, angle AngleLike) float64
	Sine(angle AngleLike) float64
	Cosine(angle AngleLike) float64
	Tangent(angle AngleLike) float64
}

// Instances

/*
AngleLike is an instance interface that defines the complete set of attributes,
abstractions and methods that must be supported by each instance of a concrete
angle-like class.
*/
type AngleLike interface {
	// Attributes
	GetClass() AngleClassLike
	GetValue() float64

	// Abstractions
	Angular

	// Methods
	IsZero() bool
	AsString() string
}`

const genericTemplate_ = `/*
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
Rank is a constrained type representing the possible rankings for two values.
*/
type Rank uint8

const (
	LesserRank Rank = iota
	EqualRank
	GreaterRank
)

// Functionals

/*
RankingFunction[V any] is a functional type that defines the signature for any
function that can determine the relative ranking of two values.
*/
type RankingFunction[V any] func(
	first V,
	second V,
) Rank

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
	DefaultRanker() RankingFunction[V]

	// Constructors
	Make() SetLike[V]
	MakeFromArray(values []V) SetLike[V]
	MakeFromSequence(values Sequential[V]) SetLike[V]
	MakeWithRanker(ranker RankingFunction[V]) SetLike[V]

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
	SetRanker(ranker RankingFunction[V])

	// Abstractions
	Sequential[V]

	// Methods
	AddValue(value V)
	RemoveValue(value V)
	RemoveAll()
}`
