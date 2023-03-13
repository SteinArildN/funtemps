package main

import (
	"flag"
	"fmt"

	//local
	//"workspace/conv"
	//github
	"github.com/SteinArildN/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahrj float64
var outj string
var funfactsj string

// stein variabler
var fahr float64
var cels float64
var kelv float64

var numb float64
var tconv float64

var jippi string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahrj, "Fj", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&outj, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfactsj, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

	//stein kode
	// flag variablene
	flag.Float64Var(&fahr, "F", 15.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 30.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 45.0, "temperatur i grader kelvin")
	//
	flag.Float64Var(&numb, "N", 777.0, "custom number")
	flag.Float64Var(&tconv, "T", 1, "type konversjon")
}

func main() {
	var jtest bool
	jtest = false
	flag.Parse()
	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.
	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.
	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())
	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	if jtest {
		// Her er noen eksempler du kan bruke i den manuelle testingen
		fmt.Println(fahrj, outj, funfactsj)

		fmt.Println("len(flag.Args())", len(flag.Args()))
		fmt.Println("flag.NFlag()", flag.NFlag())

		fmt.Println(isFlagPassed("out"))

		// Eksempel på enkel logikk
		if outj == "C" && isFlagPassed("F") {
			// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
			// skal returnere °C
			fmt.Println("0°F er -17.78°C")
		}
	}
	if !jtest {
		fmt.Println(fahr, cels, kelv)
		if tconv <= 5 && tconv >= 0 {
			switch tconv {
			case 0:
				fmt.Println(numb, "farhenheit er", conv.FahrenheitToCelsius(numb), "celsius")
			case 1:
				fmt.Println(numb, "farhenheit er", conv.FahrenheitToKelvin(numb), "kelvin")
			case 2:
				fmt.Println(numb, "celsius er", conv.CelsiusToFahrenheit(numb), "fahrenheit")
			case 3:
				fmt.Println(numb, "celsius er", conv.CelsiusToKelvin(numb), "kelvin")
			case 4:
				fmt.Println(numb, "Kelvin er", conv.KelvintoFahrenheit(numb), "fahrenheit")
			case 5:
				fmt.Println(numb, "Kelvin er", conv.KelvintoCelsius(numb), "celsius")
			}

		} else {
			fmt.Println("error, select type from 0 to 5..")
		}
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
