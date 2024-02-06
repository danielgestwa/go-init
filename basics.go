package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Travel() string
}

type Dog struct {
	name  string
	breed string
}

func (dog Dog) bark() string {
	return dog.name + ": Woof!"
}

func (dogBarking Dog) barkOn(dogBarked Dog) string {
	return dogBarking.name + " is barking on " + dogBarked.name + "!"
}

func (dog Dog) Travel() string {
	return dog.name + " is running"
}

func showNumbers(numbers ...int) (int, int) {
	if len(numbers) >= 2 {
		return numbers[0], numbers[1]
	}
	return 0, 0
}

func (dog *Dog) changeName(name string) {
	dog.name = name
}

func basics() {
	var i int = 999
	var s string = "MyString"
	var b bool = true
	var u uint = 1
	const pepperPl = 200
	test, test2 := "Hello!", 1234
	var c64 complex64 = 10 + 2i
	var r rune = 'A'
	var bt byte = 'a'
	var f float64 = 0.000000001

	fmt.Println(i, s, b, u, test, test2, pepperPl, c64, bt, r, f)

	var testString string = "Test string 1234"
	var p1 string = "p2"
	p2 := p1
	p1 = "p1"
	tshp := strings.HasPrefix(testString, "Te")
	tsup := strings.ToUpper(testString)
	fmt.Println(len(testString), testString[0], testString[2:14], p1+" is not "+p2, tshp, tsup)

	var strArr [10]string = [10]string{"a", "b", "c"}
	var s2Arr = [...]string{"x", "y", "z"}
	fmt.Println(strArr, s2Arr, s2Arr[1], len(s2Arr))

	var slice = []string{"a", "b", "c"}
	slice = append(slice, "d", "e", "f")
	var slice2 []string
	slice2 = append(slice2, "x", "y", "z")

	var sfa []string = strArr[:2]
	sfaLong := strArr[:]

	fmt.Println(slice, slice2, sfa, sfaLong)

	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"
	var m2 = map[int]string{0: "a", 1: "c"}
	m2[1] = "b"
	delete(m2, 0)
	fmt.Println(m, m2)

	j := 0
	for i = 0; i < 10; i++ {
		j += i * 10
	}

	w := 0
	k := 0
	for k < 100 {
		w += k * 2
		k += 2
	}
	fmt.Println(j, w)

	for val, num := range sfa {
		fmt.Println(val, num)
	}

	// THIS WILL WORK WITH GO 1.22
	// for val, num := range 10 {
	// 	fmt.Println(val, num)
	// }

	if i < 2 {
		fmt.Println(i)
	} else {
		fmt.Println(i)
	}

	switch i {
	case 1:
		{
			fmt.Println("It's one!")
		}
	default:
		fmt.Println("None that I know")
	}

	green := Dog{name: "Green", breed: "Mixed"}
	pinkie := Dog{name: "Pinkie", breed: "Mixed"}
	fmt.Println(green.name, green.breed)
	fmt.Println(green.bark())
	fmt.Println(green.barkOn(pinkie))

	var emptyDog = Dog{}
	var emptyDogPtr = &emptyDog
	emptyDog.changeName("Lua")
	fmt.Println(emptyDog, emptyDogPtr, *emptyDogPtr)
	emptyDogPtr.changeName("Luna")
	fmt.Println(emptyDog, emptyDogPtr, *emptyDogPtr)

	var luna Animal = emptyDog
	fmt.Println(luna.Travel())
	fmt.Println(showNumbers(1000, 2000, 1, 2, 3, 4))
}
