package basic_data_types

// TODO separate by funcs with output
func dummy() {
	// 1. Boolean Type
	var isActive bool = true
	var isAdmin = false // Type inference

	// 2. Numeric Types
	// Integer Types
	var age int = 30
	var population uint64 = 8000000000

	// Floating-Point Types
	var pi float64 = 3.1415926535
	temperature := 25.5 // float64 inferred

	// Complex Numbers
	var z complex128 = complex(3, 4) // 3 + 4i
	uin
	// 3. String Type
	var name string = "Alice"
	greeting := "Hello, 世界" // Supports Unicode

	// 3.1 String Operations
	s1 := "Hello"
	s2 := "World"
	result := s1 + " " + s2 // Concatenation
	length := len(s1)       // Length in bytes
	char := s1[0]           // Get byte at index (not a rune)

	// 4. Rune Type
	var r rune = 'A'
	unicode := '世' // Unicode character

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	/* 5. Zero Values
	Default values when variables are declared but not initialized:

	Type		Zero-Value
	bool		false
	Numeric		0
	string		""
	Pointers	nil
	*/

	// 6. Type Conversion
	// Numeric to string
	str := string(65) // "A" (ASCII code)

	_, _, _, _ = isActive, isAdmin, age, population
}
