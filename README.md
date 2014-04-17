[![Build Status](https://drone.io/github.com/tleyden/checkers-bot-minimax/status.png)](https://drone.io/github.com/tleyden/checkers-bot-minimax/latest)

Bot implementation for [checkers-bot](https://github.com/tleyden/checkers-bot) which uses [minimax](http://en.wikipedia.org/wiki/Minimax) to choose the best move.  

The underlying data structures and algorithms are in [checkers-core](https://github.com/tleyden/checkers-core).

# Big Picture

![architecture png](http://cl.ly/image/3i010f3K1Z13/Screen%20Shot%202014-01-04%20at%205.51.28%20PM.png)

# Live Demo

Download the [Checkers with Crowds](https://itunes.apple.com/us/app/id698034787) iOS app and play against Red, which is currently (1/4/14) primarly being powered by the code in this project.

# Install 

First you will need to [install Go 1.1](http://golang.org/doc/install) or later.

```
$ go get -u -v github.com/tleyden/checkers-bot-minimax
$ go get github.com/couchbaselabs/go.assert
```

# Validate installation - run tests

```
$ cd $GOPATH/github.com/tleyden/checkers-bot-minimax
$ go test -v
```

# Running

```
$ go build
$ ./checkers-bot-minimax --team RED --syncGatewayUrl http://foo.com:4984/checkers
```

# Documentation

* [Checkers-bot-random]() has more information regarding getting the rest of the system components working (Couchbase Server, etc)
