# Contributing

We love every form of contribution. By participating to this project, you
agree to abide to the our [code of conduct](/CODE_OF_CONDUCT.md).

## Setup your machine

`tracker` is written in [Go](https://golang.org/).

Prerequisites are:

* Build:
  * `make`
  * [Go 1.8+](http://golang.org/doc/install) (of course! :))

Clone `tracker` from source:

```sh
$ git clone https://github.com/lucapette/tracker.git
$ cd tracker
```

Install the build and lint dependencies:

``` sh
$ make setup
```

A good way of making sure everything is all right is running the test suite:

``` sh
$ make test
```

## Test your change

You can create a branch for your changes and try to build from the source as
you go:

``` sh
$ make build
```

When you are satisfied with the changes, I suggest you run:

``` sh
$ make ci
```

Which runs all the linters and tests.

## Submit a pull request

Push your branch to your `tracker` fork and open a pull request against
the master branch.
