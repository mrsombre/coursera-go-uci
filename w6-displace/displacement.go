package main

import (
	"fmt"
	"math"
)

type DisplaceFn func(time float64) float64

func main() {
	var err error

	var acceleration float64
	fmt.Println("Enter acceleration:")
	_, err = fmt.Scanln(&acceleration)
	if err != nil {
		panic(err.Error())
	}

	var velocity float64
	fmt.Println("Enter velocity:")
	_, err = fmt.Scanln(&velocity)
	if err != nil {
		panic(err.Error())
	}

	var displacement float64
	fmt.Println("Enter displacement:")
	_, err = fmt.Scanln(&displacement)
	if err != nil {
		panic(err.Error())
	}

	calc := GenDisplaceFn(acceleration, velocity, displacement)
	var result float64
	for {
		var time float64
		fmt.Println("Enter time:")
		_, err = fmt.Scanln(&time)
		if err != nil {
			fmt.Printf("Thanks for choosing our calculator ;)")
			return
		}

		result = calc(time)
		fmt.Printf("Final displacement: %v\n", result)
		time = 0
	}
}

func GenDisplaceFn(acceleration float64, velocity float64, displacement float64) DisplaceFn {
	return func(time float64) float64 {
		// s =½ a t2 + vot + so
		return acceleration/2*math.Pow(time, 2) + velocity*time + displacement
	}
}
