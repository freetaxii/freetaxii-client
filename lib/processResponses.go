// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package lib

import (
	"encoding/json"
	"github.com/freetaxii/libtaxii/collectionMessage"
	"github.com/freetaxii/libtaxii/discoveryMessage"
	"github.com/freetaxii/libtaxii/pollMessage"
	"github.com/freetaxii/libtaxii/statusMessage"
	"log"
)

type ResponseMessageType struct {
	MessageType string `json:"message_type,omitempty"`
}

// --------------------------------------------------
// Process Response Data from Server
// --------------------------------------------------

func ProcessResponse(requestId, requestType string, rawResponseData []byte) {

	if DebugLevel >= 4 {
		log.Println("DEBUG-4: Entering processResponse")
	}

	// --------------------------------------------------
	// Figure out which TAXII message
	// --------------------------------------------------
	// First thing we need to do is figure out what type of message we got back

	var err error

	//taxiiMessage := make(map[string]interface{})
	var taxiiMessage ResponseMessageType

	// Unmarshal the JSON to the map we just defined
	err = json.Unmarshal(rawResponseData, &taxiiMessage)
	if err != nil {
		log.Fatalln("Bad Response, unable to decode JSON message type")
	}

	if DebugLevel >= 1 {
		log.Println("DEBUG-1: Found a JSON message type of", taxiiMessage.MessageType)
	}

	// --------------------------------------------------
	// Process the TAXII message
	// --------------------------------------------------
	// Now that we know what type of message we got back, lets process it

	if taxiiMessage.MessageType == "discovery_response" {
		var responseObject discoveryMessage.DiscoveryResponseMessageType
		err = json.Unmarshal(rawResponseData, &responseObject)
		if err != nil {
			log.Fatalln("Bad Response, unable to decode JSON message contents")
		}
		printDiscoveryResponse(requestId, requestType, responseObject)

	} else if taxiiMessage.MessageType == "status_message" {
		var statusObject statusMessage.StatusMessageType
		err = json.Unmarshal(rawResponseData, &statusObject)
		if err != nil {
			log.Fatalln("Bad Response, unable to decode JSON message contents")
		}
		printStatusMessage(requestId, requestType, statusObject)

	} else if taxiiMessage.MessageType == "collection_information_response" {
		var responseObject collectionMessage.CollectionResponseMessageType
		err = json.Unmarshal(rawResponseData, &responseObject)
		if err != nil {
			log.Fatalln("Bad Response, unable to decode JSON message contents")
		}
		printCollectionResponse(requestId, requestType, responseObject)

	} else if taxiiMessage.MessageType == "poll_response" {
		var responseObject pollMessage.PollResponseMessageType
		err = json.Unmarshal(rawResponseData, &responseObject)
		if err != nil {
			log.Fatalln("Bad Response, unable to decode JSON message contents")
		}
		printPollResponse(requestId, requestType, responseObject)

	} else {
		log.Fatalln("Server did not respond with a valid TAXII Message, message type was", taxiiMessage.MessageType)
	}
}
