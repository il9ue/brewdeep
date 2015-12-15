
package config

import (
	"fmt",
	"os",
	"net",
	"syscall",
	"unsafe",
	"strings"
)

// interface needs to check network interfaces and make lists.
// some of tasks interface must to do are as follows :
// 	1. compare address and actual interface list
//		2. detect and fetch all the server connections

type connectionInfo struct {
	id					uint64
	remoteNode 		bool
	conn 				unsafe.Pointer

	selfNodes		(*syscall.IpAdapterInfo)
	selfAddrList	[]Addr
}

func newConnInfo() *connectionInfo {
	connInfo := &connectionInfo{
		id:			0,
		remoteNode: false,
		selfNodes:	nil
	}
	return connInfo
}

func (cInfo *connectionInfo) getAdapterList() (*syscall.IpAdapterInfo, error) {
	bTmp := make([]byte, 1000)
	length := uint32(len(bTmp))
	adtr := (*syscall.IpAdapterInfo)(unsafe.Pointer(&bTmp[0]))

	if err := syscall.GetAdaptersInfo(adtr, &length); err == syscall.ERROR_BUFFER_OVERFLOW {
		bTmp = make([]byte, length)
		adtr = (*syscall.IpAdapterInfo)(unsafe.Pointer(&bTmp[0]))
		if err = syscall.GetAdaptersInfo(adtr, &length); err != nil {
			return nil, os.NewSyscallError("GetAdaptersInfo", err)
		}
	}
	return adtr, nil
}

func (cInfo *connectionInfo) getLocalAddresses() (*syscall.IpAdapterInfo, error) {
	if ifs, err := net.Interfaces(); err != nil {
		return err
	}

	aList, err := cInfo.getAdapterList()
	if err != nil {
		return aList, err
	}

	for _, ifi := range ifs {
		for ai := aList; ai != nil; ai = ai.Next {
			index := ai.Index

			if ifi.Index == int(index) {
				ipl := &ai.IpAddressList
				/*
				for ; ipl != nil; ipl = ipl.Next {

					fmt.Printf("%s: %s (%s)\n", ifi.Name, ipl.IpAddress, ipl.IpMask)
				}*/
			}
		}
	}
   return aList, err
}


func (cInfo *connectionInfo) getNetworkIfs() error {
	ifs, err := net.Interfaces()
	if err != nil {
		return err
	}

	addrList := make([]Addr, 0, 1)
	for _, i := range ifs {
		if strings.Contains(i.Flags.String(), "up") && (strings.Contains(i.Flags.String(), "multicast") 
			|| strings.Contains(i.Flags.String(), "broadcast")) {
			
			aSlice, err := i.Addr()
			if err != nil {
				return nil, err
			}

			extendList(cInfo.selfAddrList, i)
		}
	}
 	cInfo.selfAddrList = Append(cInfo.selfAddrList, addrList...)		
}

func extendList(slice []Addr, elemnt Addr) []Addr {
	n := len(slice)
	if n == cap(slice) {
		newSlice := make([]Addr, len(slice), 2*len(slice) + 1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n + 1]
	slice[n] = element
	return slice
}
