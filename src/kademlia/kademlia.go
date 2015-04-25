package kademlia

// Contains the core kademlia type. In addition to core state, this type serves
// as a receiver for the RPC methods, which is required by that package.

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
  "strconv"
	"container/list"

)

const (
	alpha = 3
	b     = 8 * IDBytes
	bucketsize = 20 // 做了少许修改，因为把常量名由 k->bucketsize
)

// Kademlia type. You can put whatever state you need in this.
//我认为这是k桶的结构

type Kademlia struct {
	NodeID ID
  SelfContact Contact
	Buckets []*list.List          // 每个Kademlia object 里面放置160 个桶
}

//我们需要新建一个routing table， 根据现有的id， bucket size 以及 等等（待添加）
func NewKademlia(laddr string) *Kademlia {
	// TODO: Initialize other state here as you add functionality.
	k := new(Kademlia)
	k.NodeID = NewRandomID()
  k.Buckets = make([]*list.List,160)
	// Set up RPC server
	// NOTE: KademliaCore is just a wrapper around Kademlia. This type includes
	// the RPC functions.
	rpc.Register(&KademliaCore{k})
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", laddr)
	if err != nil {
		log.Fatal("Listen: ", err)
	}
	// Run RPC server forever.
	go http.Serve(l, nil)

    // Add self contact
    hostname, port, _ := net.SplitHostPort(l.Addr().String())
    port_int, _ := strconv.Atoi(port)
    ipAddrStrings, err := net.LookupHost(hostname)
    var host net.IP
    for i := 0; i < len(ipAddrStrings); i++ {
        host = net.ParseIP(ipAddrStrings[i])
        if host.To4() != nil {
            break
        }
    }
    k.SelfContact = Contact{k.NodeID, host, uint16(port_int)}
		k.Update(&k.SelfContact)


	return k
}

type NotFoundError struct {
	id  ID
	msg string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%x %s", e.id, e.msg)
}

func (k *Kademlia) FindContact(nodeId ID) (*Contact, error) {
	// TODO: Search through contacts, find specified ID
	// Find contact with provided ID

	return nil, &NotFoundError{nodeId, "Not found"}
}

// This is the function to perform the RPC
func (k *Kademlia) DoPing(host net.IP, port uint16) string {
	// TODO: Implement
	// If all goes well, return "OK: <output>", otherwise print "ERR: <messsage>"
	return "ERR: Not implemented"
}

func (k *Kademlia) DoStore(contact *Contact, key ID, value []byte) string {
	// TODO: Implement
	// If all goes well, return "OK: <output>", otherwise print "ERR: <messsage>"
	return "ERR: Not implemented"
}

func (k *Kademlia) DoFindNode(contact *Contact, searchKey ID) string {
	// TODO: Implement
	// If all goes well, return "OK: <output>", otherwise print "ERR: <messsage>"
	return "ERR: Not implemented"
}

func (k *Kademlia) DoFindValue(contact *Contact, searchKey ID) string {
	// TODO: Implement
	// If all goes well, return "OK: <output>", otherwise print "ERR: <messsage>"
	return "ERR: Not implemented"
}

func (k *Kademlia) LocalFindValue(searchKey ID) string {
	// TODO: Implement
	// If all goes well, return "OK: <output>", otherwise print "ERR: <messsage>"
	return "ERR: Not implemented"
}

func (k *Kademlia) DoIterativeFindNode(id ID) string {
	// For project 2!
	return "ERR: Not implemented"
}
func (k *Kademlia) DoIterativeStore(key ID, value []byte) string {
	// For project 2!
	return "ERR: Not implemented"
}
func (k *Kademlia) DoIterativeFindValue(key ID) string {
	// For project 2!
	return "ERR: Not implemented"
}



func (k *Kademlia) Update(contact *Contact) {
	distance := k.NodeID.Xor(contact.NodeID)
	index := 159 - distance.PrefixLen()
	node := k.Buckets[index]
	if node == nil {
		node = list.New()
		node.PushBack(contact)
	} else {
		for e := node.Front(); e != nil; e = e.Next() {
			if e.Value.(Contact).NodeID == contact.NodeID {
				node.MoveToBack(e)
			} else if node.Len() < 20 {
				node.PushBack(contact)
			} else {
				k.DoPing(node.Host, node.Port)
			}
		}
	}
	//fmt.Println(index)
	return
}
