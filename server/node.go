package server

import (
   "fmt"
   "net"
   "time"

   "golang.org/x/net/context"  
)

const (
   defaultGossipPeriod = 1 * time.Minute
)

/*
   node.go does the following:
      1. loaded network config info and get updated
      2. initiate network and local descriptor
      3. init monitoring agent that propagates change on network connections
      4. initiate store to aggregate statistics
      5. make net server to accept requests from remote host
      6. init client to send request
      7. initiate gossip group to update node's status periodically
*/
type Node struct {
   clusterID string
}

func startBootstrap(clusterID string) error {

}

