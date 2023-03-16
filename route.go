package main

import "my_webserver/framework"

func registerRouter(core *framework.Core) {
	core.Get("fool", FooControllerHandler)
}
