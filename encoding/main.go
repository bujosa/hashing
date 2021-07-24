package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Juan Pablo Duarte (January 26, 1813 – July 15, 1877)[1] was a Dominican military leader, writer, activist, and nationalist politician who was the foremost of the founding fathers of the Dominican Republic. As one of the most celebrated figures in Dominican history, Duarte is considered a folk hero and revolutionary visionary in the modern Dominican Republic, who along with military general Ramón Matías Mella and Francisco del Rosario Sánchez, organized and promoted the Trinitario movement that eventually led to the Dominican revolt and independence from Haitian rule in 1844 and the start of a Dominican War of Independence."

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
