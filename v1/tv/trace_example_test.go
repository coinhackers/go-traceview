// Copyright (C) 2016 AppNeta, Inc. All rights reserved.

package tv_test

import (
	"github.com/appneta/go-traceview/v1/tv"
	"golang.org/x/net/context"
)

func ExampleNewTrace() {
	f0 := func(ctx context.Context) { // example work function
		defer tv.BeginProfile(ctx, "f0").End()
		// ... something slow ...
	}

	// create a new trace, and a context to carry it around
	ctx := tv.NewContext(context.Background(), tv.NewTrace("myExample"))
	// do some work
	f0(ctx)
	// end the trace
	tv.EndTrace(ctx)
}

func ExampleBeginLayer() {
	// create trace and bind to context, reporting first event
	ctx := tv.NewContext(context.Background(), tv.NewTrace("baseLayer"))
	// ... do something ...

	// instrument a DB query
	l, _ := tv.BeginLayer(ctx, "DBx", "Query", "SELECT * FROM tbl")
	// .. execute query ..
	l.End()

	// end trace
	tv.EndTrace(ctx)
}
