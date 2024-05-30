package main

import (
	"flag"
	"math"

	montecarlo "github.com/namelew/DevicesTransmissionSimulation/internal/monte_carlo"
)

func main() {
	nPreambles := flag.Uint("M", 64, "Number of avaliable preambles or communication channels in the simulation")
	nDevices := flag.Uint("n", 30, "Number of devices in racing from the preambles per simulation")
	powExecutions := flag.Uint("r", 4, "Expoent of the 10 pow that define the number of execution rounds")

	flag.Parse()

	maxOfSimulation := uint32(math.Pow10(int(*powExecutions)))

	monteCarloSimulations := montecarlo.New(&montecarlo.GlobalOptions{
		NRounds:    maxOfSimulation,
		NPreambles: uint8(*nPreambles),
		NDevices:   uint16(*nDevices),
		R:          uint8(*powExecutions),
	})

	monteCarloSimulations.Start()
}
