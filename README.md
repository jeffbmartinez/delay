# delay
Middleware for adding delays (latency) to your endpoint. Intended to simulate latency for testing its effects on faster networks.

Add a delay of a specified duration by adding an http header to your request.

The header name is `X-Add-Delay`. The value can be any value parsable by [golang's time.ParseDuration](http://golang.org/pkg/time/#ParseDuration). For example, `300ms`, `2.5s`, and `0.5m` are all valid values. Technically `-1.5s` is a valid value, but it won't actually speed up your end point :p

If the value is not parsable it will not add any latency. It will behave as if the header wasn't there at all.

## Header Examples

`X-Add-Delay: 300ms`

`X-Add-Delay: 2.5s`

## How to use it with negroni

There is no configuration.

Just slip it in as negroni middleware:

	import (
	    "github.com/codegangsta/negroni"
	    "github.com/jeffbmartinez/delay"
	)
	
    n := negroni.New(...)
	n.Use(delay.Middleware{})
	// ...
