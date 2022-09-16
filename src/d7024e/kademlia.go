package d7024e

// Stores routing table
type Kademlia struct {
	routing *RoutingTable
}

func NewKademlia(table *RoutingTable) *Kademlia {
	kademlia := &Kademlia{}
	kademlia.routing = table
	return kademlia
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
}
