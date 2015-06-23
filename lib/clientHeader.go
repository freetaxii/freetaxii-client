// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package lib

import (
	"net/http"
)

// --------------------------------------------------
// Add TAXII Values to HTTP Header
// --------------------------------------------------

func buildTaxiiHeaderJson(r *http.Request) {

	r.Header.Add("User-Agent", `freetaxii.httpClient`)

	// Specifies which HTTP Media Types the requestor accepts in response.
	r.Header.Add("Accept", `application/json`)

	// Specifies the HTTP Media Type in which the entity body is formatted.
	r.Header.Add("Content-Type", `application/json`)

	// Specifies which TAXII Message Bindings the requestor accepts in response.
	r.Header.Add("X-Taxii-Accept", `urn:taxii.mitre.org:message:json:1.1`)

	// Specifies the TAXII Message Binding in which the entity body is formatted.
	r.Header.Add("X-Taxii-Content-Type", `urn:taxii.mitre.org:message:json:1.1`)

	// Specifies which TAXII Protocol Binding is used for this message.
	r.Header.Add("X-Taxii-Protocol", `urn:taxii.mitre.org:protocol:http:1.0`)

	// Specifies the version of the TAXII Services Specification to which this message conforms.
	r.Header.Add("X-Taxii-Services", `urn:taxii.mitre.org:services:1.1`)
}
