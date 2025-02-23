package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

var investmentAmount float64
var years float64
var expectedReturnRate float64


func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := futureValueCalculator(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	//fmt.Println("Future Value:", futureValue)
	//fmt.Println("Future Value After Inflation:", futureRealValue)

	//Using Printf
	fmt.Printf("Future Value: %.2f\nFuture Value After Inflation: %.2f", futureValue, futureRealValue)
}

//Passing parameters to functions
func outputText(text string) {
	fmt.Print(text)
}

func futureValueCalculator(investmentAmount float64, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv / math.Pow(1 + inflationRate/100, years)
	return fv, rfv
	//return
}