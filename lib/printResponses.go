// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package lib

import (
	"fmt"
	"github.com/freetaxii/libtaxii/collection"
	"github.com/freetaxii/libtaxii/discovery"
	"github.com/freetaxii/libtaxii/status"
	"log"
)

// --------------------------------------------------
// Print a header for all output
// --------------------------------------------------

func PrintOutputHeader() {

	if iDebug >= 3 {
		log.Println("DEBUG: Entering printOutputHeader")
	}

	fmt.Println("")
	fmt.Println("FreeTAXII Client")
	fmt.Println("Copyright, Bret Jordan")
	fmt.Println("Version:", sVersion)
	fmt.Println("")
}

// --------------------------------------------------
// Print the Status Message
// --------------------------------------------------

func printStatusMessage(reqid string, o *status.StatusMessageType) {

	if iDebug >= 3 {
		log.Println("DEBUG: Entering printStatusMessage")
	}

	PrintOutputHeader()
	fmt.Println("Request")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:   discovery_request")
	fmt.Println("Message ID:    ", reqid)
	fmt.Println("\n\n")

	fmt.Println("Response")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:   status_message")
	fmt.Println("Message ID:    ", o.Id)
	if o.InResponseTo != "" {
		fmt.Println("Response ID:   ", o.InResponseTo)
	}
	fmt.Println("Type:          ", o.Type)
	if o.Details != nil {
		fmt.Println("Details:")
		for k, v := range o.Details {
			fmt.Println("    ", k, " = ", v)
		}
	}
	if o.Message != "" {
		fmt.Println("Message:       ", o.Message)
	}
}

// --------------------------------------------------
// Print the Discovery Response
// --------------------------------------------------

func printDiscoveryResponse(reqid string, o *discovery.DiscoveryResponseType) {

	if iDebug >= 3 {
		log.Println("DEBUG: Entering printDiscoveryResponse")
	}

	PrintOutputHeader()
	fmt.Println("Request")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:   discovery_request")
	fmt.Println("Message ID:    ", reqid)
	fmt.Println("\n\n")

	fmt.Println("Response")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:   discovery_response")
	fmt.Println("Message ID:    ", o.Id)
	fmt.Println("Response ID:   ", o.InResponseTo)

	for i, _ := range o.Services {
		fmt.Println("")
		fmt.Println("    === Service Instance ===")
		fmt.Println("    Service Type:    ", o.Services[i].Type)
		fmt.Println("    Available:       ", o.Services[i].Available)
		fmt.Println("    Service Address: ", o.Services[i].Address)
		fmt.Println("    Service Version: ", o.Services[i].Version)
		fmt.Println("    Service Protocol:", o.Services[i].Protocol)
		fmt.Println("    Supported Encodings")
		for j, _ := range o.Services[i].Encodings {
			fmt.Println("        ", o.Services[i].Encodings[j])
		}
		if o.Services[i].SupportedQueries != nil {
			fmt.Println("    Supported Queries")
			for k, v := range o.Services[i].SupportedQueries {
				fmt.Println("Key", k, "Query", v)
			}
		}
		if o.Services[i].Message != "" {
			fmt.Println("    Message:         ", o.Services[i].Message)
		}
	}
}

// --------------------------------------------------
// Print the Collection Response
// --------------------------------------------------

func printCollectionResponse(reqid string, o *collection.CollectionResponseType) {

	if iDebug >= 3 {
		log.Println("DEBUG: Entering printCollectionResponse")
	}

	PrintOutputHeader()
	fmt.Println("Request")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:   collection_information_request")
	fmt.Println("Message ID:    ", reqid)
	fmt.Println("\n\n")

	fmt.Println("Response")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:   collection_information_response")
	fmt.Println("Message ID:    ", o.Id)
	fmt.Println("Response ID:   ", o.InResponseTo)

	for i, _ := range o.Collections {
		fmt.Println("")
		fmt.Println("    === Collections ===")

		fmt.Println("    Name:        ", o.Collections[i].Name)
		fmt.Println("    Type:        ", o.Collections[i].Type)
		fmt.Println("    Available:   ", o.Collections[i].Available)
		fmt.Println("    Description: ", o.Collections[i].Description)
		fmt.Println("    Volume:      ", o.Collections[i].Volume)

		for j, _ := range o.Collections[i].PushMethods {
			fmt.Println("    Push Methods:")
			fmt.Println("        Address:   ", o.Collections[i].PushMethods[j].Address)
			fmt.Println("        Protocol:  ", o.Collections[i].PushMethods[j].Protocol)
			for k, _ := range o.Collections[i].PushMethods[j].Encodings {
				fmt.Println("        Encodings: ", o.Collections[i].PushMethods[j].Encodings[k])
			}
		}
		for j, _ := range o.Collections[i].PollServices {
			fmt.Println("    Poll Services:")
			fmt.Println("        Address:   ", o.Collections[i].PollServices[j].Address)
			fmt.Println("        Protocol:  ", o.Collections[i].PollServices[j].Protocol)
			for k, _ := range o.Collections[i].PollServices[j].Encodings {
				fmt.Println("        Encodings: ", o.Collections[i].PollServices[j].Encodings[k])
			}
		}
		for j, _ := range o.Collections[i].SubscriptionServices {
			fmt.Println("    Subscription Services:")
			fmt.Println("        Address:   ", o.Collections[i].SubscriptionServices[j].Address)
			fmt.Println("        Protocol:  ", o.Collections[i].SubscriptionServices[j].Protocol)
			for k, _ := range o.Collections[i].SubscriptionServices[j].Encodings {
				fmt.Println("        Encodings: ", o.Collections[i].SubscriptionServices[j].Encodings[k])
			}

		}
		for j, _ := range o.Collections[i].InboxServices {
			fmt.Println("    Inbox Services:")
			fmt.Println("        Address:   ", o.Collections[i].InboxServices[j].Address)
			fmt.Println("        Protocol:  ", o.Collections[i].InboxServices[j].Protocol)
			for k, _ := range o.Collections[i].InboxServices[j].Encodings {
				fmt.Println("        Encodings: ", o.Collections[i].InboxServices[j].Encodings[k])
			}
		}
	}
}
