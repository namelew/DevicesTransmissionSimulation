package discreteevent

import (
	"github.com/fschuetz04/simgo"
	"github.com/namelew/DevicesTransmissionSimulation/packages/writer"
)

type Simulation struct {
	options        *GlobalOptions
	simulation     *simgo.Simulation
	resultsWritter *writer.Writer
}

type GlobalOptions struct {
	NPreambles        uint8
	NRounds           uint16
	DeviceArrivalRate float64
}
