package main

import (
	"fmt"
	"TPGolang/model"
)


	func main() {

		var cad string
		fmt.Print("Ingrese cadena: ")
		fmt.Scan(&cad)
		
		var res model.Result
		var err error
		var r *model.Result
		
		r, err = res.GetResult(cad)
		if err == nil {
			fmt.Println(r)
		} else {
			fmt.Println(err)
		}
	}
	
