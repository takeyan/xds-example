// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
package main

import (
	"context"
	"flag"
	"os"
	"time"
	"fmt"

	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	testv3 "github.com/envoyproxy/go-control-plane/pkg/test/v3"

        example "bg-deploy/pkg/xdshelper"
)

var (
	l example.Logger

	port     uint
	basePort uint
	mode     string

	nodeID string

        upstreamHostname string = "www.ibm.com"
        snapshotVersion string = "1"
)

func init() {
	l = example.Logger{}

	flag.BoolVar(&l.Debug, "debug", false, "Enable xDS server debug logging")

	// The port that this xDS server listens on
	flag.UintVar(&port, "port", 18000, "xDS management server port")

	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")
}

func main() {
	flag.Parse()

	// Create a cache
	cache := cachev3.NewSnapshotCache(false, cachev3.IDHash{}, l)

	// Create the snapshot that we'll serve to Envoy
	snapshot := example.GenerateSnapshot2(upstreamHostname, snapshotVersion)
	if err := snapshot.Consistent(); err != nil {
		l.Errorf("snapshot inconsistency: %+v\n%+v", snapshot, err)
		os.Exit(1)
	}
	l.Debugf("will serve snapshot %+v", snapshot)

	// Add the snapshot to the cache
	if err := cache.SetSnapshot(nodeID, snapshot); err != nil {
		l.Errorf("snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}

	// Run the xDS server
	ctx := context.Background()
	cb := &testv3.Callbacks{Debug: l.Debug}
	srv := serverv3.NewServer(ctx, cache, cb)
	go example.RunServer(ctx, srv, port)


        fmt.Print("### before sleep 60 sec")
        time.Sleep(time.Second * 60)
        fmt.Print("### after sleep 60 sec")

  // Create the snapshot that we'll serve to Envoy
        snapshot = example.GenerateSnapshot2("www.redhat.com", "2")
        if err := snapshot.Consistent(); err != nil {
            l.Errorf("snapshot inconsistency: %+v\n%+v", snapshot, err)
            os.Exit(1)
        }
        l.Debugf("will serve snapshot %+v", snapshot)


    // Add the snapshot to the cache
        if err := cache.SetSnapshot(nodeID, snapshot); err != nil {
            l.Errorf("snapshot error %q for %+v", err, snapshot)
            os.Exit(1)
        }



        fmt.Print("### before sleep 3600 sec")
        time.Sleep(time.Second * 3600)
        fmt.Print("### after sleep 3600 sec")


}
