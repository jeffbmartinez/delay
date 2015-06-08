# delay
Middleware for adding delays (latency) to your endpoint. Intended to simulate latency for testing it's effects on faster networks.

Add a delay of a specified duration by adding an http header to your request.

The header name is `X-Add-Delay`. The value can be any value parsable by [golang's time.ParseDuration](http://golang.org/pkg/time/#ParseDuration).

If the value is not parsable it will not add any latency. It will behave as if the header wasn't there at all.

Sorry, adding negative delay won't speed up your endpoints :p

## Examples

`X-Add-Delay: 300ms`

`X-Add-Delay: 2.5s`
