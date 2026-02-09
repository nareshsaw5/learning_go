module learning_go/hello

go 1.22.5

// Use local greetings module during development
replace learning_go/greetings => ../greetings

require learning_go/greetings v0.0.0-00010101000000-000000000000
