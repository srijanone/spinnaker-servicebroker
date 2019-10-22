package broker

import (
	"flag"
)

// AddFlags is a hook called to initialize the CLI flags for broker options.
// It is called after the flags are added for the skeleton and before flag
// parse is called.
func AddFlags(o *Options) {
	flag.StringVar(&o.CatalogPath, "catalogPath", "", "The path to the catalog")
	flag.BoolVar(&o.Async, "async", false, "Indicates whether the broker is handling the requests asynchronously.")
	flag.StringVar(&o.GateUrl, "GateUrl", "", "Spinnaker Gate Url.")
}
