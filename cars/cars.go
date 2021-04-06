package main

import "fmt"

func main() {
	var c Car
	fmt.Println(c)
	fmt.Printf("+v: %+v\n", c)
	fmt.Printf("#v: %#v\n", c)
}

type Car struct {
	ID  string
	Lat float64
	Lng float64
}

func NewCar(id string, lat, lng float64) (*Car, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id")
	}

	car := Car{
		ID:  id,
		Lat: lat,
		Lng: lng,
	}
	return &car, nil
}
