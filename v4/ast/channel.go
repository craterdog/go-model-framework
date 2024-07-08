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

package ast

// CLASS ACCESS

// Reference

var channelClass = &channelClass_{
	// Initialize class constants.
}

// Function

func Channel() ChannelClassLike {
	return channelClass
}

// CLASS METHODS

// Target

type channelClass_ struct {
	// Define class constants.
}

// Constructors

func (c *channelClass_) Make() ChannelLike {
	// Validate the arguments.
	switch {
	default:
		return &channel_{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

// INSTANCE METHODS

// Target

type channel_ struct {
	// Define instance attributes.
	class_ ChannelClassLike
}

// Attributes

func (v *channel_) GetClass() ChannelClassLike {
	return v.class_
}

// Private
