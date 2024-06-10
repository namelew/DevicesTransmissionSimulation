package writer

import "encoding/csv"

type WriterRegister struct {
	NPreambles           uint8
	NSuccessTransmitions uint8
	NUsedPreambles       uint8
	R                    uint16
	NDevices             uint16
	CollisionProb        float32
}

type Writer struct {
	maxBufferSize      uint32
	aggregationChannel chan *WriterRegister
	aggregationBuffer  []*WriterRegister
	filename           string
	csvWriter          *csv.Writer
}
