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
	Bucketsize int         //我新建了一个Bucketsize 属性
	Other_Contacts []Contact	 //这是一个指向Contact数组的指针，即指向其他结点
}


//定义 Kademlia 结构的一个方法
func (route_table *Kademlia) Update(id ID) {
		fmt.Println("good enough") //测试代码
		route_table.FindContact(id)
}



//我们需要新建一个routing table， 根据现有的id， bucket size 以及 等等（待添加）
func NewKademlia(laddr string) *Kademlia {
	// TODO: Initialize other state here as you add functionality.
	k := new(Kademlia)
	k.NodeID = NewRandomID()
	k.Bucketsize = bucketsize	//设定每个bucket大小为20


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
		k.Other_Contacts = []Contact{}   // 初始化数组指针
		k.Update(k.NodeID)	//测试代码
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
		for i := 0; i<len(k.Other_Contacts); i++ {
			//用for 循环查找 Other_Contacts 中，是否含有 特定的id值
			if nodeId == k.Other_Contacts[i].NodeID {
				return &k.Other_Contacts[i], nil
			}
		}
    if nodeId == k.SelfContact.NodeID {
				fmt.Println("Find myself") //测试代码
        return &k.SelfContact, nil
    }
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
