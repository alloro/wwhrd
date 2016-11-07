///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
)

// Initialize and run wwhrd
func main() {
	parser := newCli()
	c, err := parser.Parse()
	if err != nil {
		if _, ok := err.(*flags.Error); ok {
			typ := err.(*flags.Error).Type
			switch {
			case typ == flags.ErrHelp:
				parser.WriteHelp(os.Stdout)
			case typ == flags.ErrCommandRequired && len(c[0]) == 0:
				parser.WriteHelp(os.Stdout)
			default:
				log.Info(err.Error() + string(typ))
				parser.WriteHelp(os.Stdout)
			}
		} else {
			log.Fatalf("Exiting: %s", err.Error())
		}
	}
}