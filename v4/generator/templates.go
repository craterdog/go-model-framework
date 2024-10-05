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

package generator

// Templates

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
other interfaces and intrinsic types—and the class implementations only depend
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
type TrigonometricFunction func(
	angle AngleLike,
) float64

// Classes

/*
AngleClassLike is a class interface that defines the set of class constants,
constructors and functions that must be supported by each angle-like concrete
class.
*/
type AngleClassLike interface {
	// Constructor
	MakeFromValue(
		value float64,
	) AngleLike
	MakeFromString(
		value string,
	) AngleLike

	// Constants
	Pi() AngleLike
	Tau() AngleLike

	// Functions
	Apply(
		function TrigonometricFunction,
		angle AngleLike,
	) float64
	Sine(
		angle AngleLike,
	) float64
	Cosine(
		angle AngleLike,
	) float64
	Tangent(
		angle AngleLike,
	) float64
}

// Instances

/*
AngleLike is an instance interface that defines the complete set of attributes,
abstractions and methods that must be supported by each instance of a concrete
angle-like class.
*/
type AngleLike interface {
	// Public
	GetClass() AngleClassLike
	AsFloat() float64
	IsZero() bool

	// Aspect
	Angular
}

// Aspects

/*
Angular is an aspect interface that defines a set of method signatures that
must be supported by each instance of an angular concrete class.
*/
type Angular interface {
	AsNormalized() AngleLike
	InUnits(
		units Units,
	) float64
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
other interfaces and intrinsic types—and the class implementations only depend
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

// Classes

/*
ArrayClassLike[V any] is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete array-like class.
*/
type ArrayClassLike[V any] interface {
	// Constructor
	MakeWithSize(
		size uint,
	) ArrayLike[V]
	MakeFromValue(
		value []V,
	) ArrayLike[V]
	MakeFromSequence(
		values Sequential[V],
	) ArrayLike[V]

	// Constants
	DefaultRanker() RankingFunction[V]
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
	// Public
	GetClass() ArrayClassLike[V]
	SortValues()
	SortValuesWithRanker(
		ranker RankingFunction[V],
	)

	// Aspect
	Accessible[V]
	Sequential[V]
	Updatable[V]
}

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
	GetValue(
		index int,
	) V
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
	IsEmpty() bool
	GetSize() int
	AsArray() []V
}

/*
Updatable[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of an updatable concrete class.
*/
type Updatable[V any] interface {
	SetValue(
		index int,
		value V,
	)
	SetValues(
		index int,
		values Sequential[V],
	)
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
other interfaces and intrinsic types—and the class implementations only depend
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
type NormFunction[V any] func(
	value V,
) float64

// Classes

/*
ComplexClassLike is a class interface that defines the set of class constants,
constructors and functions that must be supported by each complex-like concrete
class.
*/
type ComplexClassLike interface {
	// Constructor
	Make(
		realPart float64,
		imaginaryPart float64,
		form Form,
	) ComplexLike
	MakeFromValue(
		value complex128,
	) ComplexLike

	// Constants
	Zero() ComplexLike
	Infinity() ComplexLike

	// Functions
	Inverse(
		value ComplexLike,
	) ComplexLike
	Sum(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Difference(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Reciprocal(
		value ComplexLike,
	) ComplexLike
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
	// Public
	GetClass() ComplexClassLike
	IsReal() bool
	IsImaginary() bool

	// Attribute
	GetRealPart() float64
	GetImaginaryPart() float64
	GetForm() Form
	SetForm(
		form Form,
	)

	// Aspect
	Continuous
}

// Aspects

/*
Continuous is an aspect interface that defines a set of method signatures
that must be supported by each instance of a continuous concrete class.
*/
type Continuous interface {
	IsZero() bool
	IsDiscrete() bool
	IsInfinity() bool
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
other interfaces and intrinsic types—and the class implementations only depend
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

// Classes

/*
AssociationClassLike[K comparable, V any] is a class interface that defines
the complete set of class constants, constructors and functions that must be
supported by each concrete association-like class.
*/
type AssociationClassLike[K comparable, V any] interface {
	// Constructor
	Make(
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
	// Constructor
	Make() CatalogLike[K, V]
	MakeFromArray(
		associations []AssociationLike[K, V],
	) CatalogLike[K, V]
	MakeFromMap(
		associations map[K]V,
	) CatalogLike[K, V]
	MakeFromSequence(
		associations Sequential[AssociationLike[K, V]],
	) CatalogLike[K, V]

	// Constants
	DefaultRanker() RankingFunction[AssociationLike[K, V]]

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
	// Public
	GetClass() AssociationClassLike[K, V]

	// Attribute
	GetKey() K
	GetValue() V
	SetValue(
		value V,
	)
}

/*
CatalogLike[K comparable, V any] is an instance interface that defines the
complete set of instance attributes, abstractions and methods that must be
supported by each instance of a concrete catalog-like class.
*/
type CatalogLike[K comparable, V any] interface {
	// Public
	GetClass() CatalogClassLike[K, V]
	SortValues()
	SortValuesWithRanker(
		ranker RankingFunction[AssociationLike[K, V]],
	)

	// Aspect
	Associative[K, V]
	Sequential[AssociationLike[K, V]]
}

// Aspects

/*
Associative[K comparable, V any] defines the set of method signatures that
must be supported by all sequences of key-value associations.
*/
type Associative[K comparable, V any] interface {
	GetKeys() Sequential[K]
	GetValue(
		key K,
	) V
	RemoveValue(
		key K,
	) V
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
	AsArray() []V
	GetSize() int
	IsEmpty() bool
}`
