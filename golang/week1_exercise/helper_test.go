package exercise

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	t.Run("1. with is empty | simple data case ", func(t *testing.T) {
		input := []interface{}{false, 0, "", nil}
		expect := true
		for _, v := range input {
			actual := IsEmpty(v)
			if actual != expect {
				t.Error("IsEmpty of \"", v, "\" is failed")
			}
		}
	})
	t.Run("2. with is not empty | simple data case", func(t *testing.T) {
		input := []interface{}{true, 1, "a"}
		expect := false
		for _, v := range input {
			actual := IsEmpty(v)
			if actual != expect {
				t.Error("IsEmpty of \"", v, "\" is success")
			}
		}
	})
	t.Run("3. with is empty | array or slice case", func(t *testing.T) {
		input := []float64{}
		expect := true
		actual := IsEmpty(input)
		if actual != expect {
			t.Error("IsEmpty of \"", input, "\" is failed")
		}
	})
	t.Run("4. with is not empty | array or slice case", func(t *testing.T) {
		input := []float64{1}
		expect := false
		actual := IsEmpty(input)
		if actual != expect {
			t.Error("IsEmpty of \"", input, "\" is failed")
		}
	})
	t.Run("5. with is empty | struct case", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}

		input := &User{}
		expect := true
		actual := IsEmpty(input)
		if actual != expect {
			t.Error("IsEmpty of \"", input, "\" is failed")
		}
	})
	t.Run("6. with is not empty | struct case", func(t *testing.T) {
		type User struct {
			Name string
			Age  uint16
		}

		input := &User{Name: "mario", Age: 33}
		expect := false
		actual := IsEmpty(input)
		if actual != expect {
			t.Error("IsEmpty of \"", input, "\" is failed")
		}
	})
}

//==============================
func TestMax(t *testing.T) {
	t.Run("1. with illegal input | not array or slice", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("should panic here")
			}
		}()
		Max("abc")
	})
	t.Run("2. with illegal input | empty array", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("should panic here")
			}
		}()
		Max([]int{})
	})
	t.Run("3. with legal input | array of int", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		var expect int = 4
		max := Max(input)
		if max != expect {
			t.Error("Max of \"", input, "\" is wrong")
		}
	})
	t.Run("4.1 with legal input | array of uint8", func(t *testing.T) {
		input := []uint8{1, 2, 3, 4}
		var expect uint8 = 4
		max := Max(input)
		if max != expect {
			t.Error("Max of \"", input, "\" is wrong")
		}
	})
	t.Run("4.2 with legal input | array of uint16", func(t *testing.T) {
		input := []uint16{1, 2, 3, 4}
		var expect uint16 = 4
		max := Max(input)
		if max != expect {
			t.Error("Max of \"", input, "\" is wrong")
		}
	})
	t.Run("4.3 with legal input | array of uint32", func(t *testing.T) {
		input := []uint32{1, 2, 3, 4}
		var expect uint32 = 4
		max := Max(input)
		if max != expect {
			t.Error("Max of \"", input, "\" is wrong")
		}
	})
	t.Run("4.4 with legal input | array of uint64", func(t *testing.T) {
		input := []uint64{1, 2, 3, 4}
		var expect uint64 = 4
		max := Max(input)
		if max != expect {
			t.Error("Max of \"", input, "\" is wrong")
		}
	})
	t.Run("5.1 with legal input | array of float32", func(t *testing.T) {
		input := []float32{1, 2, 3, 4}
		var expect float32 = 4
		max := Max(input)
		if max != expect {
			t.Error("Max of \"", input, "\" is wrong")
		}
	})
	t.Run("5.2 with legal input | array of float64", func(t *testing.T) {
		input := []float64{1, 2, 3, 4}
		var expect float64 = 4
		max := Max(input)
		if max != expect {
			t.Error("Max of \"", input, "\" is wrong")
		}
	})
}

//==============================
func TestContain(t *testing.T) {
	t.Run("1.1 with illegal input "+
		"| obj is not array or slice "+
		"| obj is not string "+
		"| obj is not struct "+
		"-> obj is numberic ", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("should be panic here")
			}
		}()
		Contain(0, 0)
	})
	t.Run("1.2 with illegal input "+
		"| obj is not array or slice "+
		"| obj is not string "+
		"| obj is not struct "+
		"-> obj is nil", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("should be panic here")
			}
		}()
		Contain(nil, 0)
	})
	t.Run("2.1 with legal input | obj is array or slice but empty", func(t *testing.T) {
		obj := []int{}
		element := 0
		expect := false
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("3.1 with legal input | obj is string but empty", func(t *testing.T) {
		obj := ""
		element := "a"
		expect := false
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("3.2 with legal input | obj is string but element is not", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("should be panic here")
			}
		}()
		Contain("abc", 0)
	})
	t.Run("3.3 with legal input | obj is string simple case", func(t *testing.T) {
		obj := "abc"
		element := "a"
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("4.1 with legal input | obj is struct sample case fail", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		obj := User{Name: "mario", Age: 33}
		element := "maria"
		expect := false
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("4.2 with legal input | obj is struct sample case success", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		obj := User{Name: "mario", Age: 33}
		element := "mario"
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("4.3 with legal input | obj is struct sample case success", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		obj := User{Name: "mario", Age: 33}
		element := 33
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("4.4 with legal input | obj is struct sample case success", func(t *testing.T) {
		type User struct {
			Name string
			Age  uint
		}
		obj := User{Name: "mario", Age: 33}
		var element uint = 33
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("4.5 with legal input | obj is struct sample case success", func(t *testing.T) {
		type User struct {
			Name string
			Age  float32
		}
		obj := User{Name: "mario", Age: 33}
		var element float32 = 33
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
}
func TestCheckArrContainSimple(t *testing.T) {
	t.Run("2.2 with legal input | obj is array or slice and success"+
		"| simple case : numberic [uint, unit8, uint16, uni32, unit64, float32, float64])", func(t *testing.T) {
		obj := []int{1, 2, 3, 4}
		element := 1
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.3 with legal input | obj is array or slice and success"+
		"| simple case : numberic [uint, unit8, uint16, uni32, unit64, float32, float64])", func(t *testing.T) {
		obj := []uint{1, 2, 3, 4}
		var element uint = 1
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.4 with legal input | obj is array or slice and success"+
		"| simple case : numberic [uint, unit8, uint16, uni32, unit64, float32, float64])", func(t *testing.T) {
		obj := []uint8{1, 2, 3, 4}
		var element uint8 = 1
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.5 with legal input | obj is array or slice and success"+
		"| simple case : numberic [uint, unit8, uint16, uni32, unit64, float32, float64])", func(t *testing.T) {
		obj := []uint16{1, 2, 3, 4}
		var element uint16 = 1
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.6 with legal input | obj is array or slice and success"+
		"| simple case : numberic [uint, unit8, uint16, uni32, unit64, float32, float64])", func(t *testing.T) {
		obj := []uint32{1, 2, 3, 4}
		var element uint32 = 1
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.7 with legal input | obj is array or slice and success"+
		"| simple case : numberic [uint, unit8, uint16, uni32, unit64, float32, float64])", func(t *testing.T) {
		obj := []uint64{1, 2, 3, 4}
		var element uint64 = 1
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.8 with legal input | obj is array or slice and success"+
		"| simple case : numberic [uint, unit8, uint16, uni32, unit64, float32, float64])", func(t *testing.T) {
		obj := []float32{1, 2, 3, 4}
		var element float32 = 1
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.9 with legal input | obj is array or slice and success"+
		"| simple case : numberic [uint, unit8, uint16, uni32, unit64, float32, float64])", func(t *testing.T) {
		obj := []float64{1, 2, 3, 4}
		var element float64 = 1
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.10 with legal input | obj is array or slice"+
		"| simple case : array of int & element is string)", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("should be panic here")
			}
		}()
		Contain([]int{1, 2, 3, 4}, "a")
	})
	t.Run("2.11 with legal input | obj is array or slice and success"+
		"| simple case : array of int & element is string)", func(t *testing.T) {
		obj := []string{"a", "b", "c", "d"}
		element := "b"
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.12 with legal input | obj is array or slice and success"+
		"| complex case : array of struct and element is same type)", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		user1 := User{Name: "mario", Age: 33}
		user2 := User{Name: "maria", Age: 18}
		obj := []User{user1, user2}
		element := user2
		expect := false
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
}
func TestCheckArrContainComplex(t *testing.T) {
	t.Run("2.12 with legal input | obj is array or slice and success"+
		" | complex case : element is array or slice but emty", func(t *testing.T) {
		obj := []int{1, 2, 3, 4}
		element := []int{}
		expect := false
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.14 with legal input | obj is array or slice and success"+
		" | complex case : element is array or slice", func(t *testing.T) {
		obj := []int{1, 2, 3, 4}
		element := []int{5}
		expect := false
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
	t.Run("2.13 with legal input | obj is array or slice and success"+
		" | complex case : element is array or slice", func(t *testing.T) {
		obj := []int{1, 2, 3, 4}
		element := []int{1, 3}
		expect := true
		actual := Contain(obj, element)
		if actual != expect {
			t.Error("Check contain of \"", element, "\" in \"", obj, "\" is wrong")
		}
	})
}

//==============================
func TestLast(t *testing.T) {
	t.Run("1.1 with illegal input", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("should be panic here")
			}
		}()
		Last(0)
	})
	t.Run("1.2 with legal input but empty ", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("should be panic here")
			}
		}()
		Last([]int{})
	})
	t.Run("2.1 with legal input | simple case int success", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		expect := 4
		actual := Last(arr)
		if actual != expect {
			t.Error("Last of \"", arr, "\" is wrong")
		}
	})
	t.Run("2.2 with legal input | simple case string success", func(t *testing.T) {
		arr := []string{"a", "b", "c", "d"}
		expect := "d"
		actual := Last(arr)
		if actual != expect {
			t.Error("Last of \"", arr, "\" is wrong")
		}
	})
	t.Run("2.3 with legal input | simple case struct success", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		user1 := User{Name: "mario", Age: 33}
		user2 := User{Name: "maria", Age: 18}
		arr := []User{user1, user2}
		expect := user2
		actual := Last(arr)
		if actual != expect {
			t.Error("Last of \"", arr, "\" is wrong")
		}
	})
}
