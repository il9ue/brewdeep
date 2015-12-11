
// author : Daniel Q. Kim

package main
//package cli

import (
   "fmt"
   "os"
   "os/signal"
   "time"

   "github.com/spf13/cobra"
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

func runInit(_ *cobra.Command, args []string) {
   // more logic here
   initNode(args);
}


// init nodes
// 1. load default user setting
// 2. Load default node network setting
func initNode(args []string) {

}