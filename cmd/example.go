package main

import (
	"fmt"

	"github.com/PaulioRandall/go-examples/webserver"
)

func main() {
	fmt.Println("Starting server on port 8080")
	webserver.ListenAndServe(":8080")
	//_ = webserver.ListenAndServe // stop import error

	// 1.
	// Assignment
	// Assignment, implicit typing (:=)
	// Assignment, var
	// Assignment, const

	// 2.
	// Global, const
	// Global, var

	// 3.
	// Basic types
	//
	// bool
	//
	// int
	// int8
	// int16
	// int32
	// int64
	//
	// uint
	// uint8
	// uint16
	// uint32
	// uint64
	// uintptr
	//
	// float32
	// float64
	//
	// string
	//
	// byte (uint8)
	//
	// rune (int32, Unicode)
	//
	// complex64
	// complex128

	// 4. Moved

	// 5.
	// Imports
	// Imports, pulls in from repository

	// 6.
	// If
	// If else
	// If else if

	// 7.
	// Switch
	// Switch, fallthrough
	// Switch, match block
	// Switch, type switch

	// 8.
	// For, infinite
	// For, while
	// For, i++
	// For, range (each)
	// For, do while

	// 9.
	// Strings
	// Strings, bytes
	// Strings, runes
	// Strings, for loop

	// 10.
	// Pointers
	// Casting, type(value), compile time
	// Assertions, .(type), runtime

	// 11.
	// Functions
	// Functions, parameters (pass by value)
	// Functions, multiple return (not the same as tuples)
	// Functions, named return parameters

	// 12.
	// Error handling
	// Error handling, use an if

	// 13.
	// Structs
	// Structs, fields
	// Structs, public vs private
	// Structs, function receivers
	// Structs, embedding

	// 14.
	// Interfaces
	// Interfaces, functions
	// Interfaces, empty interface
	// Interfaces, embedding
	// Interfaces, define at the site of use
	// Interfaces, use as parameters but return concrete types

	// 15.
	// Zero values
	//
	// bool:       false
	// any number: 0
	// string:     ""
	// function:   nil
	// struct:     instance with all fields being zero
	// interface:  nil
	// pointers:   nil

	// 16.
	// Testing 

	// 17.
	// Concise error handling, still being debated
	// Generics, on their way
}
