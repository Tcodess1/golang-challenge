package main

// Go Constructs Practice — 15 Basic Coding Questions
//
// Instructions:
// - Each question is described above its function.
// - Complete the function body without changing the signature.
// - Keep solutions simple and idiomatic.

// 1) Variables + Arithmetic
// Question: Return the sum of two integers.
func Q1Sum(a int, b int) int {
	panic("TODO: implement")
}

// 2) Conditionals (if/else)
// Question: Return true if n is even, otherwise false.
func Q2IsEven(n int) bool {
	panic("TODO: implement")
}

// 3) For loop
// Question: Return the factorial of n (n!). Assume n >= 0.
func Q3Factorial(n int) int {
	panic("TODO: implement")
}

// 4) Slices
// Question: Return a new slice containing only even numbers from nums.
func Q4FilterEvens(nums []int) []int {
	panic("TODO: implement")
}

// 5) Maps
// Question: Count the frequency of each word in words.
func Q5WordCount(words []string) map[string]int {
	panic("TODO: implement")
}

// 6) Strings + Range (runes)
// Question: Count vowels (a, e, i, o, u — case-insensitive) in s.
func Q6CountVowels(s string) int {
	panic("TODO: implement")
}

// 7) Structs
// Question: Define total score by summing all Student scores.
type Student struct {
	Name   string
	Scores []int
}

func Q7TotalScore(st Student) int {
	panic("TODO: implement")
}

// 8) Methods
// Question: Add a method that returns true if Rectangle is a square.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) IsSquare() bool {
	panic("TODO: implement")
}

// 9) Pointers
// Question: Double the value pointed to by n.
func Q9DoubleInPlace(n *int) {
	panic("TODO: implement")
}

// 10) Multiple return values + error handling
// Question: Divide a by b. If b == 0, return an error.
func Q10Divide(a float64, b float64) (float64, error) {
	panic("TODO: implement")
}

// 11) Variadic functions
// Question: Return the largest number from nums.
// If nums is empty, return 0.
func Q11Max(nums ...int) int {
	panic("TODO: implement")
}

// 12) Defer
// Question: Append numbers from n down to 1 into a slice using defer behavior.
// Example: n=3 -> []int{3,2,1}
func Q12Countdown(n int) []int {
	panic("TODO: implement")
}

// 13) Interfaces
// Question: Implement Area for Circle and Rectangle2.
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle2 struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	panic("TODO: implement")
}

func (r Rectangle2) Area() float64 {
	panic("TODO: implement")
}

// 14) Goroutines + Channels
// Question: Send numbers 1..n into a channel in a goroutine,
// then return their sum after receiving all values.
func Q14SumWithChannel(n int) int {
	panic("TODO: implement")
}

// 15) Closures
// Question: Return a function that increments and returns an internal counter.
// Example:
//
//	next := Q15Counter()
//	next() -> 1
//	next() -> 2
func Q15Counter() func() int {
	panic("TODO: implement")
}
