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

// CLASS INTERFACE

// Access Function

func Channel() ChannelClassLike {
	return channelReference()
}

// Constructor Methods

func (c *channelClass_) Make() ChannelLike {
	var instance = &channel_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Public Methods

func (v *channel_) GetClass() ChannelClassLike {
	return v.getClass()
}

// Private Methods

func (v *channel_) getClass() *channelClass_ {
	return channelReference()
}

// PRIVATE INTERFACE

// Instance Structure

type channel_ struct {
	// Declare the instance attributes.
}

// Class Structure

type channelClass_ struct {
	// Declare the class constants.
}

// Class Reference

func channelReference() *channelClass_ {
	return channelReference_
}

var channelReference_ = &channelClass_{
	// Initialize the class constants.
}
