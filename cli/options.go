package cli

import (
	"flag"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/il9ue/brewdeep/config"
)

var optionUsage = map[string]string {
	"addr": `The host:port to bind for HTTP/RPC traffic`,
	"certs": `Directory containing certifications. if --secure=true, this option will be used.`
}

func initOpts() {
	
}

func init() {
	initOpts()

	cobra.OnIntialize(func() {
		// ToDo
	})
}