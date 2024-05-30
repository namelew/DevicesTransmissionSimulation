package montecarlo

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/namelew/DevicesTransmissionSimulation/packages/writer"
)

func New(opt *GlobalOptions) *Simulation {
	result_filename := "monte-carlo-simulation-results-%d_%d_%d.csv"

	result_filename = fmt.Sprintf(result_filename, opt.NDevices, opt.NPreambles, opt.R)

	return &Simulation{
		options:        opt,
		resultsWritter: writer.New(opt.NRounds, result_filename),
	}
}

func (s *Simulation) Start() {
	s.resultsWritter.Start()

	wg_rounds := sync.WaitGroup{}

	wg_rounds.Add(int(s.options.NRounds))

	for i := range s.options.NRounds {
		go s.execRound(int(i), &wg_rounds)
	}

	wg_rounds.Wait()

	s.resultsWritter.Close()
}

func (s *Simulation) execRound(rid int, wg *sync.WaitGroup) {
	preambles := make([]uint8, s.options.NPreambles)

	for range s.options.NDevices {
		selectedPreamble := rand.Int31n(int32(s.options.NPreambles))

		preambles[selectedPreamble] += 1
	}

	successTransmission := 0
	usedPreambles := 0

	for p := range preambles {
		if preambles[p] >= 1 {
			usedPreambles++
			if preambles[p] == 1 {
				successTransmission++
			}
		}
	}

	successProb := float32(successTransmission) / float32(s.options.NPreambles)
	utilization := float32(usedPreambles) / float32(s.options.NPreambles)

	s.resultsWritter.Write(&writer.WriterRegister{
		R:                    s.options.R,
		NPreambles:           s.options.NPreambles,
		NDevices:             s.options.NDevices,
		NSuccessTransmitions: uint8(successTransmission),
		CollisionProb:        1 - successProb,
		Utilization:          utilization,
	})

	log.Printf("Round %d:\n\tSuccess:%f\n\tUtilization:%f\n", rid, successProb, utilization)

	wg.Done()
}
