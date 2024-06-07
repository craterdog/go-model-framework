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

const typeBodyTemplate_ = `
	// TBA - Validate the value.
	return <TargetName>_[<Arguments>](value)
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

const typeTargetTemplate_ = `
type <TargetName>_[<Parameters>] <TargetType>
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

const typeMethodTemplate_ = `
func (v <TargetName>_[<Arguments>]) <MethodName>(<Parameters>)<ResultType> {<Body>}
`

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

const getterClassTemplate_ = `
	return <ClassName>[<Arguments>]()
`

const getterBodyTemplate_ = `
	return v.<AttributeName>_
`

const setterBodyTemplate_ = `
	v.<AttributeName>_ = <ParameterName>
`

const angleTemplate_ = `/*
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
	MakeFromValue(value float64) AngleLike
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

	// Abstractions
	Angular

	// Methods
	AsFloat() float64
	IsZero() bool
}`

const arrayTemplate_ = `/*
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
Accessible[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of an accessible concrete class.  The
values in an accessible class are accessed using indices. The indices of an
accessible class are ORDINAL rather than ZERO based—which never really made
sense except for pointer offsets.

This approach allows for positive indices starting at the beginning of the
sequence, and negative indices starting at the end of the sequence as follows:

	    1           2           3             N
	[value 1] . [value 2] . [value 3] ... [value N]
	   -N        -(N-1)      -(N-2)          -1

Notice that because the indices are ordinal based, the positive and negative
indices are symmetrical.
*/
type Accessible[V any] interface {
	// Methods
	GetValue(index int) V
	GetValues(
		first int,
		last int,
	) Sequential[V]
}

/*
Sequential[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of a sequential concrete class.

NOTE: that sizes should be of type "uint" but the Go language does not allow
arithmetic and comparison operations between "int" and "uint" so we us "int" for
the return type to make it easier to use.
*/
type Sequential[V any] interface {
	// Methods
	IsEmpty() bool
	GetSize() int
	AsArray() []V
}

/*
Updatable[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of an updatable concrete class.
*/
type Updatable[V any] interface {
	// Methods
	SetValue(
		index int,
		value V,
	)
	SetValues(
		index int,
		values Sequential[V],
	)
}

// Classes

/*
ArrayClassLike[V any] is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete array-like class.
*/
type ArrayClassLike[V any] interface {
	// Constants
	DefaultRanker() RankingFunction[V]

	// Constructors
	MakeWithSize(size uint) ArrayLike[V]
	MakeFromValue(value []V) ArrayLike[V]
	MakeFromSequence(values Sequential[V]) ArrayLike[V]
}

// Instances

/*
ArrayLike[V any] is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete array-like class.

An array-like class maintains a fixed length indexed sequence of values.  Each
value is associated with an implicit positive integer index. An array-like class
uses ORDINAL based indexing rather than the more common—and nonsensical—ZERO
based indexing scheme.
*/
type ArrayLike[V any] interface {
	// Attributes
	GetClass() ArrayClassLike[V]

	// Abstractions
	Accessible[V]
	Sequential[V]
	Updatable[V]

	// Methods
	SortValues()
	SortValuesWithRanker(ranker RankingFunction[V])
}`

const complexTemplate_ = `/*
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
Units is a constrained type representing the possible notational forms for the
complex number.
*/
type Form uint8

const (
	Polar Form = iota
	Rectangular
)

// Functionals

/*
NormFunction[V any] is a functional type that defines the signature for any
mathematical norm function.
*/
type NormFunction[V any] func(value V) float64

// Aspects

/*
Continuous is an aspect interface that defines a set of method signatures
that must be supported by each instance of a continuous concrete class.
*/
type Continuous interface {
	// Methods
	IsZero() bool
	IsDiscrete() bool
	IsInfinity() bool
}

// Classes

/*
ComplexClassLike is a class interface that defines the set of class constants,
constructors and functions that must be supported by each complex-like concrete
class.
*/
type ComplexClassLike interface {
	// Constants
	Zero() ComplexLike
	Infinity() ComplexLike

	// Constructors
	MakeWithAttributes(
		realPart float64,
		imaginaryPart float64,
		form Form,
	) ComplexLike
	MakeFromValue(value complex128) ComplexLike

	// Functions
	Inverse(value ComplexLike) ComplexLike
	Sum(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Difference(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Reciprocal(value ComplexLike) ComplexLike
	Product(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Quotient(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Norm(
		function NormFunction[ComplexLike],
		value ComplexLike,
	) float64
}

// Instances

/*
ComplexLike is an instance interface that defines the complete set of attributes,
abstractions and methods that must be supported by each instance of a concrete
complex-like class.
*/
type ComplexLike interface {
	// Attributes
	GetClass() ComplexClassLike
	GetRealPart() float64
	GetImaginaryPart() float64
	GetForm() Form
	SetForm(form Form)

	// Abstractions
	Continuous

	// Methods
	IsReal() bool
	IsImaginary() bool
}`

const catalogTemplate_ = `/*
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
Associative[K comparable, V any] defines the set of method signatures that
must be supported by all sequences of key-value associations.
*/
type Associative[K comparable, V any] interface {
	// Methods
	GetKeys() Sequential[K]
	GetValue(key K) V
	RemoveValue(key K) V
	SetValue(
		key K,
		value V,
	)
}

/*
Sequential[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of a sequential concrete class.

NOTE: that sizes should be of type "uint" but the Go language does not allow
arithmetic and comparison operations between "int" and "uint" so we us "int" for
the return type to make it easier to use.
*/
type Sequential[V any] interface {
	// Methods
	AsArray() []V
	GetSize() int
	IsEmpty() bool
}

// Classes

/*
AssociationClassLike[K comparable, V any] is a class interface that defines
the complete set of class constants, constructors and functions that must be
supported by each concrete association-like class.
*/
type AssociationClassLike[K comparable, V any] interface {
	// Constructors
	MakeWithAttributes(
		key K,
		value V,
	) AssociationLike[K, V]
}

/*
CatalogClassLike[K comparable, V any] is a class interface that defines the
complete set of class constants, constructors and functions that must be
supported by each concrete catalog-like class.

The following functions are supported:

Extract() returns a new catalog containing only the associations that are in
the specified catalog that have the specified keys.  The associations in the
resulting catalog will be in the same order as the specified keys.

Merge() returns a new catalog containing all of the associations that are in
the specified catalogs in the order that they appear in each catalog.  If a
key is present in both catalogs, the value of the key from the second
catalog takes precedence.
*/
type CatalogClassLike[K comparable, V any] interface {
	// Constants
	DefaultRanker() RankingFunction[AssociationLike[K, V]]

	// Constructors
	Make() CatalogLike[K, V]
	MakeFromArray(associations []AssociationLike[K, V]) CatalogLike[K, V]
	MakeFromMap(associations map[K]V) CatalogLike[K, V]
	MakeFromSequence(associations Sequential[AssociationLike[K, V]]) CatalogLike[K, V]

	// Functions
	Extract(
		catalog CatalogLike[K, V],
		keys Sequential[K],
	) CatalogLike[K, V]
	Merge(
		first CatalogLike[K, V],
		second CatalogLike[K, V],
	) CatalogLike[K, V]
}

// Instances

/*
AssociationLike[K comparable, V any] is an instance interface that defines the
complete set of instance attributes, abstractions and methods that must be
supported by each instance of a concrete association-like class.
*/
type AssociationLike[K comparable, V any] interface {
	// Attributes
	GetClass() AssociationClassLike[K, V]
	GetKey() K
	GetValue() V
	SetValue(value V)
}

/*
CatalogLike[K comparable, V any] is an instance interface that defines the
complete set of instance attributes, abstractions and methods that must be
supported by each instance of a concrete catalog-like class.
*/
type CatalogLike[K comparable, V any] interface {
	// Attributes
	GetClass() CatalogClassLike[K, V]

	// Abstractions
	Associative[K, V]
	Sequential[AssociationLike[K, V]]

	// Methods
	SortValues()
	SortValuesWithRanker(ranker RankingFunction[AssociationLike[K, V]])
}`
