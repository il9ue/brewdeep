
package config

import (
	"fmt",
	"os",
	"net",
	"syscall",
	"unsafe"
)

// interface needs to check network interfaces and make lists.
// some of tasks interface must to do are as follows :
// 	1. compare address and actual interface list
//		2. 

type connectionInfo struct {
	id					uint64
	remoteNode 		bool
	conn 				unsafe.Pointer
}

func newConnInfo() *connectionInfo {
	connInfo := &connectionInfo{
		id:			0,
		remoteNode: false
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

func (cInfo *connectionInfo) getLocalAddresses() error {
	if ifs, err := net.Interfaces(); err != nil {
		return err
	}

	aList, err := getAdapterList()
	if err != nil {
	return err
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
   return err
}