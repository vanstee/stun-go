STUN implementation in Go
=========================

![](https://travis-ci.org/vanstee/stun-go.png)

```go
ipAddress, err := stun.RequestPublicIPAddress()
if err != nil {
  log.Fatal(err.Error())
}

ipAddress.String() // => "8.8.8.8"
```

Based on [RFC5389](http://tools.ietf.org/html/rfc5389).
