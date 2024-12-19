package main

import "fmt"

func main() {
	// var firstName, lastName string = "John", "Doe"
	// lastName := "Doe" //Bunu kısaltılmış halinde yazsak bile Go dili bunun string olduğunu anlar o yüzden  var firstname string = "John" şeklinde yazmamıza gerek yok
	// firstName := "John"
	// var name = "John Doe"
	// var isMarried bool = true

	// var (
	// 	age    int     = 40
	// 	height float32 = 1.75
	// 	isOkey bool    = true
	// )
	// var weight, isStudent = 75, false
	// isWorking, surname, name := true, "Doe", "John"
	// fmt.Println(firstName, lastName, isMarried, name, age, height, isOkey, weight, isStudent, isWorking, surname)

	//ZERO VALUES
	var a int
	fmt.Println(a) //a değerine bir değer atamadık o yüzden go otomatik zeroa value olarak 0 döner
	// The zero value is:

	// 0 for numeric types,
	// false for the boolean type, and
	// "" (the empty string) for strings.

	studentName := "John"
	// studentName:= "Doe" //Bu şekilde tekrar tanımlama yapamayız çünkü ":" işareti atama anlamına gelir ikincisinide sadece = kullan yani;
	studentName = "Doe"
	fmt.Println(studentName)
	// fmt.Printf("%T\n", firstName)
	// fmt.Printf("%T\n", lastName)
	// fmt.Printf("%T\n", isMarried)
}
