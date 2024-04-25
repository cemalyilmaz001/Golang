package main

import ("fmt")


const(
	first = "name1"
    second = "name2"
    third = "name3"
)

func Primitive()  {
    name := "cezeri"

    var j int 
    j = 1

    var k float32
    k = 2.11

    fmt.Println(name)
    fmt.Println(j)
    fmt.Println(k)
}


func Pointer() {
	var number *int32 = new(int32)
	*number = 1337

	fmt.Println(number) //lastName'in adresini yazar.

	fmt.Println(*number) //1337 değerini yazar.
}

func Ornek3() {
	name := "MDI"

	var lastName *string = &name

	fmt.Println(*lastName, lastName, &name) //Çıktısı: MDI, lastname'in adresi, name'in adresi olur.
}

func main() {

	// Primitive Veri Yapısı.
	// Primitive() 

	// Pointer'lar
	// Pointer()

	// 
	// Ornek3()

	// Const 
	/*const pi = 3.14159
    println("Pi:", pi)
    //println(first,second,third) // iota --> Çıktısı: 0 1 2 */


    // Array
    var arr [3]int
	arr[0] = 1
	arr[1] = 0
	arr[2] = 0

	/*arr2 := [3]int{1,2,3}
	fmt.Println(arr, arr2)*/


	// Slice
	/*slice := []string { 
		"x","y","z","a",
	}

	fmt.Println(slice)*/


	// Slice
	/*slice := []string { 
		"x","y","z","a",
	}

	fmt.Println(slice)

	slice2 := arr[0:3]
	fmt.Println(slice2)*/


	/*m := map[string]int{"orhun",42}
	fmt.Println(m["orhun"]) //Çıktı: 42*/


}



