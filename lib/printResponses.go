// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package lib

import (
	"fmt"
	"github.com/freetaxii/libtaxii/collectionMessage"
	"github.com/freetaxii/libtaxii/discoveryMessage"
	"github.com/freetaxii/libtaxii/pollMessage"
	"github.com/freetaxii/libtaxii/statusMessage"
	"log"
)

// --------------------------------------------------
// Print a header for all output
// --------------------------------------------------

func PrintOutputHeader() {

	if DebugLevel >= 4 {
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

func printStatusMessage(reqid, messageType string, o statusMessage.StatusMessageType) {

	if DebugLevel >= 4 {
		log.Println("DEBUG-4: Entering printStatusMessage")
	}

	PrintOutputHeader()
	fmt.Println("Request")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:  ", messageType)
	fmt.Println("Message ID:    ", reqid)
	fmt.Println("\n\n")

	fmt.Println("Response")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:  ", o.MessageType)
	fmt.Println("Message ID:    ", o.Id)
	if o.InResponseTo != "" {
		fmt.Println("Response ID:   ", o.InResponseTo)
	}
	fmt.Println("Type:          ", o.Type)
	if o.Details != nil {
		fmt.Println("Details:")
		for k, v := range o.Details {
			if v != "" {
				fmt.Println("    ", k, " = ", v)
			}
		}
	}
	if o.Message != "" {
		fmt.Println("Message:       ", o.Message)
	}
}

// --------------------------------------------------
// Print the Discovery Response
// --------------------------------------------------

func printDiscoveryResponse(reqid, messageType string, o discoveryMessage.DiscoveryResponseMessageType) {

	if DebugLevel >= 4 {
		log.Println("DEBUG: Entering printDiscoveryResponse")
	}

	PrintOutputHeader()
	fmt.Println("Request")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:  ", messageType)
	fmt.Println("Message ID:    ", reqid)
	fmt.Println("\n\n")

	fmt.Println("Response")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:  ", o.MessageType)
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

func printCollectionResponse(reqid, messageType string, o collectionMessage.CollectionResponseMessageType) {

	if DebugLevel >= 4 {
		log.Println("DEBUG: Entering printCollectionResponse")
	}

	PrintOutputHeader()
	fmt.Println("Request")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:  ", messageType)
	fmt.Println("Message ID:    ", reqid)
	fmt.Println("\n\n")

	fmt.Println("Response")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:  ", o.MessageType)
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

		for j, _ := range o.Collections[i].DeliveryParameters {
			fmt.Println("    Push Methods:")
			fmt.Println("        Address:   ", o.Collections[i].DeliveryParameters[j].Address)
			fmt.Println("        Protocol:  ", o.Collections[i].DeliveryParameters[j].Protocol)
			for k, _ := range o.Collections[i].DeliveryParameters[j].Encodings {
				fmt.Println("        Encodings: ", o.Collections[i].DeliveryParameters[j].Encodings[k])
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

// --------------------------------------------------
// Print the Poll Response
// --------------------------------------------------

func printPollResponse(reqid, messageType string, o pollMessage.PollResponseMessageType) {

	if DebugLevel >= 4 {
		log.Println("DEBUG: Entering printPollResponse")
	}

	PrintOutputHeader()
	fmt.Println("Request")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:  ", messageType)
	fmt.Println("Message ID:    ", reqid)
	fmt.Println("\n\n")

	fmt.Println("Response")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Message Type:    ", o.MessageType)
	fmt.Println("Message ID:      ", o.Id)
	fmt.Println("Response ID:     ", o.InResponseTo)
	fmt.Println("Collection Name: ", o.CollectionName)
	fmt.Println("More:            ", o.More)
	fmt.Println("Result ID:       ", o.ResultId)
	fmt.Println("Result Part #:   ", o.ResultPartNumber)
	if o.SubscriptionId != "" {
		fmt.Println("Subscription ID: ", o.SubscriptionId)
	}
	if o.BeginTimestamp != "" {
		fmt.Println("Begin Timestamp: ", o.BeginTimestamp)
	}
	if o.EndTimestamp != "" {
		fmt.Println("End Timestamp:   ", o.EndTimestamp)
	}
	fmt.Println("Record Count:    ", o.RecordCount)
	fmt.Println("Partial Count:   ", o.PartialCount)
	if o.Message != "" {
		fmt.Println("Message:         ", o.Message)
	}

	for i, _ := range o.ContentBlocks {
		fmt.Println("")
		fmt.Println("    === Content ===")
		fmt.Println("    Content Encoding: ", o.ContentBlocks[i].ContentEncoding)
		fmt.Println("    Content:\n", o.ContentBlocks[i].Content)
		if o.ContentBlocks[i].TimestampLabel != "" {
			fmt.Println("    Timestamp:        ", o.ContentBlocks[i].TimestampLabel)
		}
		if o.ContentBlocks[i].Message != "" {
			fmt.Println("    Message:          ", o.ContentBlocks[i].Message)
		}
	}
}
