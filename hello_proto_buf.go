package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {

	loc := &Point{
		X: 5,
		Y: 30,
	}

	origin := &Pixel{
		Location: loc,
		Colour: "FFFFFF",
	}

	data, err := proto.Marshal(origin)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)

	newOrigin := &Pixel{}

	err = proto.Unmarshal(data, newOrigin)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(newOrigin.Location.GetY())
	fmt.Println(newOrigin.Location.GetX())
	fmt.Println(newOrigin.GetColour())


}
