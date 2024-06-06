package discreteevent

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/fschuetz04/simgo"
	"github.com/namelew/DevicesTransmissionSimulation/packages/writer"
	"gonum.org/v1/gonum/stat/distuv"
)

func New(opt *GlobalOptions) *Simulation {
	result_filename := "discrete-events-simulation-results-%f_%d_%d.csv"

	result_filename = fmt.Sprintf(result_filename, opt.DeviceArrivalRate, opt.NPreambles, opt.NRounds)

	return &Simulation{
		options:        opt,
		simulation:     &simgo.Simulation{},
		resultsWritter: writer.New(uint32(opt.NRounds), result_filename),
	}
}

func (s *Simulation) preamblesSelectionProcess(proc simgo.Process, preambles []uint8, mutex *sync.Mutex, rid uint8, did uint8) {
	chosedPreambles := rand.Intn(int(s.options.NPreambles))

	mutex.Lock()
	defer mutex.Unlock()

	preambles[chosedPreambles]++

	if preambles[chosedPreambles] == 1 {
		log.Printf("Round %d: Device %d manage to transmit in preamble %d", rid, did, chosedPreambles)
	} else {
		log.Printf("Round %d: Device %d failed to transmit in preamble %d, colission with %d devices", rid, did, chosedPreambles, preambles[chosedPreambles]-1)
	}
}

func (s *Simulation) transmissionsProcess(proc simgo.Process) {
	poissonRNG := distuv.Poisson{Lambda: s.options.DeviceArrivalRate}
	var roundID uint8 = 0

	for {
		numberOfDevices := int(poissonRNG.Rand())
		preambles := make([]uint8, s.options.NPreambles)
		pmutex := sync.Mutex{}
		childProcs := make([]simgo.Awaitable, numberOfDevices)

		log.Println("Starting Round", roundID)

		for i := range numberOfDevices {
			childProcs[i] = s.simulation.ProcessReflect(s.preamblesSelectionProcess, preambles, &pmutex, roundID, i)
		}

		for i := range childProcs {
			proc.Wait(childProcs[i])
		}

		var successTransmissions uint64 = 0
		var usedPreambles uint16 = 0

		for i := range preambles {
			if preambles[i] >= 1 {
				if preambles[i] == 1 {
					successTransmissions++
				}
				usedPreambles++
			}
		}

		successProb := float32(successTransmissions) / float32(s.options.NPreambles)

		log.Printf("Round %d:\n\tSuccess:%f\n\tTransmitted:%d\n", roundID, successProb, usedPreambles)

		s.resultsWritter.Write(&writer.WriterRegister{
			R:                    roundID,
			NPreambles:           s.options.NPreambles,
			NDevices:             uint16(numberOfDevices),
			NSuccessTransmitions: uint8(successTransmissions),
			CollisionProb:        1 - successProb,
			NUsedPreambles:       uint8(usedPreambles),
		})

		roundID++

		if roundID == s.options.NRounds {
			return
		}
	}
}

func (s *Simulation) Start() {
	s.resultsWritter.Start()

	s.simulation.Process(s.transmissionsProcess)

	s.simulation.RunUntil(float64(s.options.NRounds))

	s.resultsWritter.Close()
}
