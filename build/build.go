package main

import "github.com/goyek/goyek"

func main() {
	flow := &goyek.Taskflow{}

	msg := flow.RegisterStringParam(goyek.StringParam{
		Name:    "msg",
		Usage:   "message",
		Default: "Hello world",
	})

	hello := flow.Register(goyek.Task{
		Name:   "hello",
		Usage:  "I am printing a message",
		Params: goyek.Params{msg},
		Command: func(tf *goyek.TF) {
			tf.Log(msg.Get(tf))
		},
	})

	flow.Register(goyek.Task{
		Name:  "fail",
		Usage: "I am failing",
		Command: func(tf *goyek.TF) {
			tf.Error("I am always failing, never working like in reality")
		},
	})

	test := flow.Register(goyek.Task{
		Name:    "test",
		Usage:   "go test",
		Command: goyek.Exec("go", "test"),
	})

	lint := flow.Register(goyek.Task{
		Name:  "lint",
		Usage: "golangci-lint",
		Command: func(tf *goyek.TF) {
			if err := tf.Cmd("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0").Run(); err != nil {
				tf.Fatalf("install golangci-lint: %v", err)
			}
			if err := tf.Cmd("golangci-lint", "run").Run(); err != nil {
				tf.Fatal(err)
			}
		},
	})

	all := flow.Register(goyek.Task{
		Name:  "all",
		Usage: "build pipeline",
		Deps:  goyek.Deps{hello, test, lint},
	})

	flow.DefaultTask = all

	flow.Main()
}
