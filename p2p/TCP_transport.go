package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	
	listenAddress string //the address that the transport will listen to (IP + Port)
	listerner net.Listener //listener that will accept incoming connections
	
	//mu (mutex) is the lock that will be used to protect the peers map
	//This is needed because the peers map will be accessed by multiple goroutines and we need to make sure that the access is safe
	mu sync.RWMutex
	peers map[net.Addr]Peer //key: address of the peer / value: the peer itself

}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress : listenAddress,
	}
}

func (t* TCPTransport) ListenAndAccept() error {
	//Listen for incoming connections
	var err error
	t.listerner, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	
	//Accept incoming connections
	go t.AcceptLoop()
	return nil
}

func (t*TCPTransport) AcceptLoop(){
	for {
		conn, err := t.listerner.Accept()
		if(err != nil){
			fmt.Println("Error accepting connection: ", err)
		}
		go t.HandleConnection(conn)
	}
}

func (t*TCPTransport) HandleConnection(conn net.Conn){
	fmt.Printf("New connection: %+v\n", conn)
}
