package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, os.Interrupt)

	log.Println("starting...")
	defer log.Println("exit")
	closer, err := run(args())
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	exit := make(chan interface{})
	go func() {
		defer close(exit)
		ioutil.ReadAll(os.Stdin)
	}()
	select {
	case <-ctrlc:
	case <-exit:
	}
}

func args() Args {
	source, err := withext("db3")
	if err != nil {
		log.Fatal(err)
	}
	args := NewArgs()
	args.Set("source", getenv("AUTH_DB_SOURCE", source))
	args.Set("driver", getenv("AUTH_DB_DRIVER", "sqlite"))
	args.Set("endpoint", getenv("AUTH_ENDPOINT", "127.0.0.1:31651"))
	return args
}

func run(args Args) (func(), error) {
	dao, err := NewDao(args)
	if err != nil {
		return nil, err
	}
	args.Set("dao", dao)
	closer, err := rest(args)
	if err != nil {
		return nil, err
	}
	return func() {
		closer()
		err := dao.Close()
		if err != nil {
			log.Println(err)
		}
	}, nil
}
