# Go Experiment

## Compile

You can easily compile this experiment using Rake, the Ruby version of Make.
If you have Ruby installed, Rake is available on your machine, just run the following command:

		$ rake compile

You can force the architecture to compile for by setting the ARCH environment.
Currently the following ARCH values are supported:

* x86-64
* 386
* arm

## App

The app doesn't do much yet. It's mainly a way for me to practice Go and try to build something
a bit more complex than a "Hello World" app.

At this point, the app proxies localhost requests to google.
Eventually, the app would proxy based on various route rules.

## Dependencies

There are no external dependencies outside of #golang obviously.
