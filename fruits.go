package main

import "fmt"

func main() {
	applecost := 5.99
	pearcost := 7.0
	wallet := 23.0
	
	fmt.Println("Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн. Ми маємо 23 грн.")
	fmt.Println("______________")
	fmt.Println("Питання перше:")
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println("Щоб купити 9 яблук та 8 груш нам потрібно", applecost*9+pearcost*8, "гривень")
	fmt.Println("______________")
	fmt.Println("Питання друге:")
	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Printf("В нас 23 гривні, отже за них ми можемо купити %.f груші", wallet/pearcost)
	fmt.Println("______________")
	fmt.Println("Питання третє:")
	fmt.Println("Скільки яблук ми можемо купити?")
	fmt.Printf("В нас 23 гривні, отже за них ми можемо купити %.f яблука", wallet/applecost)
	fmt.Println("______________")
	fmt.Println("Питання четверте:")
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?")
	fmt.Printf("Дві груші та 2 яблука, котують %.2f. А в нас 23 гривні, отже нам не вистачить грошей", 2*applecost+2*pearcost)

}
