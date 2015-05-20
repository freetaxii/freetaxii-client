// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package lib

import (
	"encoding/json"
	"github.com/freetaxii/libtaxii/collection"
	"github.com/freetaxii/libtaxii/discovery"
	"github.com/freetaxii/libtaxii/status"
	"log"
)

// --------------------------------------------------
// Process Response Data from Server
// --------------------------------------------------

func ProcessResponse(requestId string, rawResponseData []byte) {

	if iDebug >= 3 {
		log.Println("DEBUG: Entering processResponse")
	}

	// --------------------------------------------------
	// Figure out which TAXII message
	// --------------------------------------------------
	// First thing we need to do is figure out what type of message we got back
	// The two options are:
	// 		Discovery Response
	// 		Status Message (aka error message)

	var err error

	// Build a simple map for the first pass at looking at the message
	taxiiMessage := make(map[string]interface{})

	// Unmarshal the JSON to the map we just defined
	err = json.Unmarshal(rawResponseData, &taxiiMessage)
	if err != nil {
		log.Fatalln("Bad Response, unable to decode JSON message type")
	}

	if iDebug >= 1 {
		for k, _ := range taxiiMessage {
			log.Println("DEBUG: Found a JSON message type of", k)
		}
	}

	// --------------------------------------------------
	// Process the TAXII message
	// --------------------------------------------------
	// Now that we know what type of message we got back, lets process it

	if _, ok := taxiiMessage["discovery_response"]; ok {
		var responseObject discovery.TaxiiDiscoveryResponseType
		err = json.Unmarshal(rawResponseData, &responseObject)
		if err != nil {
			log.Fatalln("Bad Response, unable to decode JSON message contents")
		}
		printDiscoveryResponse(requestId, responseObject.TaxiiMessage)

	} else if _, ok := taxiiMessage["status_message"]; ok {
		var statusObject status.TaxiiStatusMessageType
		err = json.Unmarshal(rawResponseData, &statusObject)
		if err != nil {
			log.Fatalln("Bad Response, unable to decode JSON message contents")
		}
		printStatusMessage(requestId, statusObject.TaxiiMessage)

	} else if _, ok := taxiiMessage["collection_information_response"]; ok {
		var responseObject collection.TaxiiCollectionResponseType
		err = json.Unmarshal(rawResponseData, &responseObject)
		if err != nil {
			log.Fatalln("Bad Response, unable to decode JSON message contents")
		}
		printCollectionResponse(requestId, responseObject.TaxiiMessage)

	} else {
		log.Fatalln("Server did not respond with a valid TAXII Message")
	}
}
