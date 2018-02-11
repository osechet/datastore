# go-datastore
An experiment around Go and Google datastore API.

[![Build Status](https://travis-ci.org/osechet/go-datastore.svg?branch=master)](https://travis-ci.org/osechet/go-datastore)
[![codecov](https://codecov.io/gh/osechet/go-datastore/branch/master/graph/badge.svg)](https://codecov.io/gh/osechet/go-datastore)

For the moment the project only provides ways to easily use datastore's [Query](https://godoc.org/google.golang.org/genproto/googleapis/datastore/v1#Query) in Go, such a `Match()` function to check if a database entry matches a [Filter](https://godoc.org/google.golang.org/genproto/googleapis/datastore/v1#Filter) or a `Comparator` interface that can be used to sort a slice of results.
