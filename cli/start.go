
// author : Daniel Q. Kim

package main
//package cli

import (
   "fmt"
   "os"
   "os/signal"
   "time"

   "github.com/spf13/cobra"
   "github.com/il9ue/brewdeep/config"
)

var initCmd = &cobra.Command{
   Use:  "init --nodes=...",
   Short:   "init new Node Mapping",
   Long: `
Initialize a new node mapping topology using the --nodes flag to specify
one or more node location settings. In force, when only one node is specified
then this means to get statistics on file systems of current node. If two or more nodes
are initiated then it's meant to do drawing node balance evaluation by periodically
crawling traffic and transaction data among network nodes.
`,
   Example: ` brewdeep init --nodes=ssd=/mnt/ssd1`,
   Run:  runInit,
}

// runInit would initiate node and cluster loading, 
// while gathering user authentication, and what else..?
func runInit(_ *cobra.Command, args []string) {
   // more logic here
   initNode(args);
}


// init nodes
// 1. load default user setting
// 2. Load default node network setting
func initNode(args []string) {

}


var startCmd = &cobra.Command{
   Use: "start",
   Short: "start a node, either joining existing the membership or by start measuring local",
   Long: `
Start a brewdeep by joining the membership or make membership rule while waiting for others to connect
and gathering local device's status and statistics.`,
   Example: `  brewdeep start --certs=<dir> --membership=host1:port1[,...] --stores=ssd,...`,
   Run:  runStart.  
}

func runStart(_ *cobra.Command, _ []string) {
   log.Infof("build version: version 0.1")

   
}


// teardown command shuts down the node server and
// destroys all data held by the node.
var teardownCmd = &cobra.Command{
   Use:   "teardown",
   Short: "destroy all data held by the node",
   Long: `
First shuts down the system and then destroys all data held by the
node, cycling through each store specified by the --stores flag.
`,
   Run: runExterminate,
}


// quitCmd command shuts down the node server.
var shutdownCmd = &cobra.Command{
   Use:   "shutdown",
   Short: "drain and shutdown node\n",
   Long: `
Shutdown the server. The first stage is drain, where any new requests
will be ignored by the server. When all extant requests have been
completed, the server exits.
`,
   Run: runQuit,
}


