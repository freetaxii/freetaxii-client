// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package lib

// Debug Level 1 = basic debugging information
// Debug Level 3 = Method / Functional walk
// Debug Level 5 = RAW packet/message output
var iDebug = 0

var sVersion = "0.2"

func MakeServerUrl(url, port, path string) string {
	return url + ":" + port + path
}
