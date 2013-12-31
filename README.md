Example bot implementation for [checkers-bot](https://github.com/tleyden/checkers-bot) which makes random moves.

# Big Picture

![architecture png](http://cl.ly/image/3S0G0h2U0R2b/Screen%20Shot%202013-10-08%20at%2010.43.00%20PM.png)

# Install 

First you will need to [install Go 1.1](http://golang.org/doc/install) or later.

```
$ go get github.com/tleyden/checkers-bot-random
$ go get github.com/couchbaselabs/go.assert
```

It is also possible to download the latest build of the [OSX binary](http://cbfs-ext.hq.couchbase.com/projects/checkers-bot/checkers-bot-random.mac.gz) directly.

# Validate installation - run tests

```
$ cd $GOPATH/github.com/tleyden/checkers-bot-random
$ go test -v
```

# Running

```
$ go build
$ ./checkers-bot-random --team RED --syncGatewayUrl http://foo.com:4984/checkers
```

# Install backend system components

To get a fully working system, you'll need the following:

* [Couchbase Server](http://www.couchbase.com/download) or [Kurobase hosted Couchbase Server](http://www.kurobase.com)

* [Sync Gateway](https://github.com/couchbase/sync_gateway)

* [Checkers Overlord](https://github.com/apage43/checkers-overlord)

# Install a viewer

[Checkers-iOS](https://github.com/couchbaselabs/Checkers-iOS) is not strictly required, but very useful in order to view the game.


# Using a proxy

NOTE: if you want to use a proxy server (useful for debugging), use the following command:

```
$ http_proxy=http://10.0.0.3:8888 ./checkers-bot-random --team BLUE --syncGatewayUrl http://10.0.0.3:4984/checkers
```

where 10.0.0.3 is the address where your proxy server and sync gateway is running.  The real address must be used rather than localhost.
