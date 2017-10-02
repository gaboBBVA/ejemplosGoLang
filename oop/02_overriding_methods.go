package main

import "fmt"

// EL TIPO Item ES LA FORMA GENERAL DE UN PROCUCTO
// EL TIPO SpecialItem EXTIENDE LA FORMA DEL PRIMER TIPO
// EL TIPO LuxuryItem TAMBIEN EXTIENDE EL PRIMER TIPO

type Item struct {
	id       string  // Named field (aggregation)
	price    float64 // Named field (aggregation)
	quantity int     // Named field (aggregation)
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item          // Anonymous field (embedding)
	catalogId int // Named field (aggregation)
}

func main() {
	special := SpecialItem{Item{"Green", 3, 5}, 207}
	fmt.Println(special.id, special.price, special.quantity, special.catalogId)
	fmt.Println("Costo Especial:", special.Cost())

	luxury := LuxuryItem{Item{"Blue", 30, 5}, 207}
	fmt.Println("Objeto de lujo:", luxury.id, " Precio:", luxury.price)
	fmt.Println("Costo de artículo de lujo:", luxury.Cost())
}

type LuxuryItem struct {
	Item           // Anonymous field (embedding)
	duty float64 // Named field (aggregation)
}

// func que sobre pone el metodo Cost de la clase que extiende
func (item *LuxuryItem) Cost() float64 {
	return item.Item.Cost() * item.duty
	//return item.price * item.markup
}
