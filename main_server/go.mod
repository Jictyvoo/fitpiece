module github.com/jictyvoo/fitpiece/main_server

go 1.17

require (
	github.com/jictyvoo/fitpiece/bumpingheart v0.0.0
)

replace (
	github.com/jictyvoo/fitpiece/bumpingheart => ../bumpingheart
)
