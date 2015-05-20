// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"code.google.com/p/getopt"
	"github.com/freetaxii/freetaxii-client/lib"
	"os"
)

var sOptURL = getopt.StringLong("url", 'u', "http://test.freetaxii.com", "URL Address (ex. http://test.freetaxii.com)", "string")
var sOptPort = getopt.StringLong("port", 'p', "8000", "Port Number (ex. 8000)", "string")
var sOptCollectionPath = getopt.StringLong("collection-service", 0, "/services/collection", "Collection Service Directory (ex. /services/collection)", "string")
var sOptDiscoveryPath = getopt.StringLong("discovery-service", 0, "/services/discovery", "Discovery Service Directory (ex. /services/discovery)", "string")
var bOptCollection = getopt.BoolLong("collection", 'c', "Send Collection Reqeust")
var bOptDiscovery = getopt.BoolLong("discovery", 'd', "Send Discovery Reqeust")
var bOptHelp = getopt.BoolLong("help", 0, "Help")
var bOptVer = getopt.BoolLong("version", 0, "Version")

func main() {
	getopt.HelpColumn = 35
	getopt.DisplayWidth = 120
	getopt.SetParameters("")
	getopt.Parse()

	if *bOptVer {
		lib.PrintOutputHeader()
		os.Exit(0)
	}

	if *bOptHelp {
		lib.PrintOutputHeader()
		getopt.Usage()
		os.Exit(0)
	}

	if *bOptDiscovery {
		serverurl := lib.MakeServerUrl(*sOptURL, *sOptPort, *sOptDiscoveryPath)
		requestId, rawResponseData := lib.SendDiscoveryRequest(serverurl)
		lib.ProcessResponse(requestId, rawResponseData)
	}

	if *bOptCollection {
		serverurl := lib.MakeServerUrl(*sOptURL, *sOptPort, *sOptCollectionPath)
		requestId, rawResponseData := lib.SendCollectionRequest(serverurl)
		lib.ProcessResponse(requestId, rawResponseData)
	}
}
