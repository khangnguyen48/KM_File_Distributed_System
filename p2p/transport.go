package p2p


//Peer is the interface that represents the remote node
type Peer interface {

}

//Transport is the interface that represents the transport layer
//It is anything that handle the communication between nodes in the network
//This can be of the form of TCP, UDP, Websocket, ...
type Transport interface {
}