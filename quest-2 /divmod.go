package piscine 

func DivMod(a int, b int, div *int, mod *int) {

	divv := a / b 
	modd := a % b 

	*div = divv
	*mod = modd

}
