package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTranpsort(t *testing.T) {
	// Create a new TCPTransport
	listenAddress := ":4000"
	transport := NewTCPTransport(listenAddress)

	assert.Equal(t, listenAddress, transport.listenAddress)

	assert.Nil(t, transport.ListenAndAccept())
}