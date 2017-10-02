package main

import (
	"fmt"
)

//Punto contiene los valores x, y
type Punto struct {
	x, y int
}

//Punto3D contiene los valores x,y,z
type Punto3D struct {
	x, y, z int
	*Punto3D // APUNTADOR A UN OBJETO DE LA MISMA CLASE
}


/*//OpPunto tiene los metodos para realizar operaciones con puntos
 type OpPunto struct {
}*/

//
func main() {
	p := Punto{}
	fmt.Println("p: ", p)

	p2 := Punto3D{
		5,
		6,
		4,
		&Punto3D{
			6,
			4,
			6,
			nil,  // RECURSIVAMENTE
		},
	}
	fmt.Println("p2: ", p2)

	//Comparando Estructuras, LOS CAMPOS DEBEN SER IGUALES

	a := Punto{5, 6}
	b := Punto{7, 4}
	fmt.Println("a == b : ", a == b)
	//
	c := Punto{7, 4}
	fmt.Println("b == c :", b == c)

	//  COMO LAS ESTRUCTURAS SON COMPARABLES SE PUEDEN USAR COMO INDICES EN LAS HASHMAPS
	figuras := make(map[Punto]string)
	figuras[a] = "Estructuras Comparables"
	fmt.Println("figuras[a] :", figuras[a])

}