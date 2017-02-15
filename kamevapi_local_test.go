//+build integration

/*
Released under MIT License <http://www.opensource.org/licenses/mit-license.php
Copyright (C) ITsysCOM GmbH. All Rights Reserved.

*/

package kamevapi

import (
	"flag"
	"testing"
	"time"
)

var kamAddr = flag.String("kam_addr", "127.0.0.1:8448", "Address where to reach kamailio evapi")

func TestKamailioConn(t *testing.T) {
	var err error
	if kea, err = NewKamEvapi(*kamAddr, "", 3, nil); err != nil {
		t.Fatal("Could not create KamEvapi, error: ", err)
	}
	err = kea.Send("CGR-SM Connected!")
	time.Sleep(time.Duration(2) * time.Second) // Wait 5 mins for events to show up
}
