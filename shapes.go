package main

type initialShape struct {
	shape [tSize][tSize]bool
	value byte
}

// long boi
var shape1 = initialShape{[4][4]bool{
	{false, false, true, false},
	{false, false, true, false},
	{false, false, true, false},
	{false, false, true, false},
}, 'L'}

// t boi
var shape2 = initialShape{[4][4]bool{
	{false, true, false, false},
	{true, true, true, false},
	{false, false, false, false},
	{false, false, false, false},
}, 'T'}

// block boi
var shape3 = initialShape{[4][4]bool{
	{true, true, false, false},
	{true, true, false, false},
	{false, false, false, false},
	{false, false, false, false},
}, 'B'}

// annoying boi
var shape4 = initialShape{[4][4]bool{
	{true, true, false, false},
	{false, true, true, false},
	{false, false, false, false},
	{false, false, false, false},
}, 'Z'}

// difficult boi
var shape5 = initialShape{[4][4]bool{
	{false, true, true, false},
	{true, true, false, false},
	{false, false, false, false},
	{false, false, false, false},
}, 'Q'}

var shapes = [...]initialShape{shape1, shape2, shape3, shape4, shape5}
