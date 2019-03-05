package prambanan_test

import (
	"github.com/helloproclub/prambanan"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PrambananSuite struct {
	suite.Suite
	NikDecoder *prambanan.Prambanan
}

func TestNikDecoderSuite(t *testing.T) {
	suite.Run(t, &PrambananSuite{})
}

func (n *PrambananSuite) SetupSuite() {

}
