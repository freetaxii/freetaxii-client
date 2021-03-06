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
var sOptDiscoveryPath = getopt.StringLong("discovery-service", 0, "/services/discovery", "Discovery Service Directory (ex. /services/discovery)", "string")
var sOptCollectionPath = getopt.StringLong("collection-service", 0, "/services/collection", "Collection Service Directory (ex. /services/collection)", "string")
var sOptPollPath = getopt.StringLong("poll-service", 0, "/services/poll", "Poll Service Directory (ex. /services/poll)", "string")
var sOptPollCollectionName = getopt.StringLong("poll-collection-name", 0, "", "Poll Collection Name (required)", "string")

var bOptDiscovery = getopt.BoolLong("discovery", 0, "Send Discovery Reqeust")
var bOptCollection = getopt.BoolLong("collection", 0, "Send Collection Reqeust")
var bOptPoll = getopt.BoolLong("poll", 0, "Send Poll Reqeust")

var bOptHelp = getopt.BoolLong("help", 0, "Help")
var bOptVer = getopt.BoolLong("version", 0, "Version")

func main() {
	getopt.HelpColumn = 40
	getopt.DisplayWidth = 120
	getopt.SetParameters("")
	getopt.Parse()

	if *bOptVer {
		printVersion()
	}

	if *bOptHelp {
		printHelp()
	}

	if *bOptDiscovery {
		serverurl := lib.MakeServerUrl(*sOptURL, *sOptPort, *sOptDiscoveryPath)
		requestId, requestType, rawResponseData := lib.SendDiscoveryRequest(serverurl)
		lib.ProcessResponse(requestId, requestType, rawResponseData)
	}

	if *bOptCollection {
		serverurl := lib.MakeServerUrl(*sOptURL, *sOptPort, *sOptCollectionPath)
		requestId, requestType, rawResponseData := lib.SendCollectionRequest(serverurl)
		lib.ProcessResponse(requestId, requestType, rawResponseData)
	}

	if *bOptPoll {
		if *sOptPollCollectionName == "" {
			printHelp()
		}
		serverurl := lib.MakeServerUrl(*sOptURL, *sOptPort, *sOptPollPath)
		requestId, requestType, rawResponseData := lib.SendPollRequest(serverurl, *sOptPollCollectionName)
		lib.ProcessResponse(requestId, requestType, rawResponseData)
	}
}

func printVersion() {
	lib.PrintOutputHeader()
	os.Exit(0)
}

func printHelp() {
	lib.PrintOutputHeader()
	getopt.Usage()
	os.Exit(0)
}
