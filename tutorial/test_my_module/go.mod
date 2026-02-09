module "github.com/test_my_module"

go 1.21.4


require (
	github.com/go_tutorial/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

replace github.com/go_tutorial/greetings => ../greetings
