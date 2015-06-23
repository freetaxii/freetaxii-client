// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package lib

import (
	"bytes"
	"encoding/json"
	"github.com/freetaxii/libtaxii/messages/collectionMessage"
	"github.com/freetaxii/libtaxii/messages/discoveryMessage"
	"github.com/freetaxii/libtaxii/messages/pollMessage"
	"io/ioutil"
	"log"
	"net/http"
)

// --------------------------------------------------
// Send Discovery Request Message to Server
// --------------------------------------------------

func SendDiscoveryRequest(serverurl string) (string, string, []byte) {

	if DebugLevel >= 4 {
		log.Println("DEBUG-4: Entering SendDiscoveryRequest")
	}

	var err error

	// --------------------------------------------------
	// Create Discovery Request Message
	// --------------------------------------------------
	requestObject := discoveryMessage.NewRequest()
	msgToSend, err := json.Marshal(requestObject)
	if err != nil {
		// If we can not create a valid message then there is something
		// wrong with the APIs and nothing is going to work.
		log.Fatalln("Unable to create Discovery Request Message")
	}

	rawResponseData := sendTaxiiMessage(serverurl, msgToSend)
	return requestObject.Id, requestObject.MessageType, rawResponseData
}

// --------------------------------------------------
// Send Collection Request Message to Server
// --------------------------------------------------

func SendCollectionRequest(serverurl string) (string, string, []byte) {

	if DebugLevel >= 4 {
		log.Println("DEBUG-4: Entering SendCollectionRequest")
	}

	var err error

	// --------------------------------------------------
	// Create Discovery Request Message
	// --------------------------------------------------
	requestObject := collectionMessage.NewRequest()
	msgToSend, err := json.Marshal(requestObject)
	if err != nil {
		// If we can not create a valid message then there is something
		// wrong with the APIs and nothing is going to work.
		log.Fatalln("Unable to create Collection Request Message")
	}

	rawResponseData := sendTaxiiMessage(serverurl, msgToSend)
	return requestObject.Id, requestObject.MessageType, rawResponseData
}

// --------------------------------------------------
// Send Poll Request Message to Server
// --------------------------------------------------

func SendPollRequest(serverurl, collectionName string) (string, string, []byte) {

	if DebugLevel >= 4 {
		log.Println("DEBUG-4: Entering SendPollRequest")
	}

	var err error

	// --------------------------------------------------
	// Create Discovery Request Message
	// --------------------------------------------------
	requestObject := pollMessage.NewRequest()
	requestObject.AddCollectionName(collectionName)
	pp := requestObject.NewPollParameters()
	pp.SetContentEncodingToJson()

	msgToSend, err := json.Marshal(requestObject)
	if err != nil {
		// If we can not create a valid message then there is something
		// wrong with the APIs and nothing is going to work.
		log.Fatalln("Unable to create Collection Request Message")
	}

	rawResponseData := sendTaxiiMessage(serverurl, msgToSend)
	return requestObject.Id, requestObject.MessageType, rawResponseData
}

// --------------------------------------------------
// Send Actual TAXII Message to Server
// --------------------------------------------------

func sendTaxiiMessage(serverurl string, msgToSend []byte) []byte {

	if DebugLevel >= 4 {
		log.Println("DEBUG: Entering sendTaxiiMessage")
	}

	// --------------------------------------------------
	// Create http web client
	// --------------------------------------------------

	webclient := &http.Client{}

	httpRequest, err := http.NewRequest("POST", serverurl, bytes.NewBuffer(msgToSend))
	if err != nil {
		log.Fatalln("%s", err)
	}

	// --------------------------------------------------
	// Send HTTP Post and get and process HTTP response
	// --------------------------------------------------

	buildTaxiiHeaderJson(httpRequest)
	httpResponse, err := webclient.Do(httpRequest)
	if err != nil {
		log.Fatalf("%s", err)
	}

	defer httpResponse.Body.Close()

	// Read in data from http response
	rawInboundData, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if DebugLevel >= 5 {
		log.Println("DEBUG: Raw HTTP Response Data", string(rawInboundData))
	}

	return rawInboundData
}
