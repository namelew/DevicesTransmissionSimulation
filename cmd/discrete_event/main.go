package main

import (
	"flag"

	discreteevent "github.com/namelew/DevicesTransmissionSimulation/internal/discrete_event"
)

func main() {
	nPreambles := flag.Uint("M", 64, "Number of avaliable preambles or communication channels in the simulation")
	deviceArrivalRate := flag.Float64("drate", 1, "Device arrival rate from the simulation")
	rounds := flag.Uint("r", 10, "Number of simulation rounds")

	flag.Parse()

	sim := discreteevent.New(&discreteevent.GlobalOptions{
		NRounds:           uint8(*rounds),
		DeviceArrivalRate: *deviceArrivalRate,
		NPreambles:        uint8(*nPreambles),
	})

	sim.Start()
}
