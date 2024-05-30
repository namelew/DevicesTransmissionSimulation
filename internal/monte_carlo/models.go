package montecarlo

import (
	"github.com/namelew/DevicesTransmissionSimulation/packages/writer"
)

type GlobalOptions struct {
	R          uint8
	NPreambles uint8
	NDevices   uint16
	NRounds    uint32
}

type Simulation struct {
	options        *GlobalOptions
	resultsWritter *writer.Writer
}
