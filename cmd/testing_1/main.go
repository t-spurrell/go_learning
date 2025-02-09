package main
import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println("Hello, World!")

	// numbers
	var uint_num uint = 10
	fmt.Println("uint_num:", uint_num)

	var float_num float64 = 10.2
	fmt.Println("float_num:", float_num)

	var num1 int16 = 10
	var num2 float32 = 10.2
	fmt.Println("num1:", num1)
	fmt.Println("num1 + num2:", float32(num1) + num2)

	var num3 = float32(num1)
	fmt.Println("num3:", num3)

	// strings
	var str1 string = "Hello"
	var str2 = `World`
	var rune1 = 'a'
	fmt.Println(str1, str2)
	fmt.Println(rune1, "here is a rune")

	fmt.Println(len("example")) // this is byte length not char length
	fmt.Println(utf8.RuneCountInString("example")) // this is char length

	// bools
	var bool1 bool = true
	var bool2 = false
	fmt.Println(bool1, bool2)
	var bool3 bool
	fmt.Println(bool3)

	cool_bool := true
	fmt.Println(cool_bool)

	// unpack
	t1, t2, b1 := 1, 2, true
	fmt.Println(t1, t2, b1)

}
