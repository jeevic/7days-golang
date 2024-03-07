package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"github.com/jeevic/7days-golang/gee-rpc/geerpc"
)

//func startServer(addr chan string) {
//	// pick a free port
//	l, err := net.Listen("tcp", ":0")
//	if err != nil {
//		log.Fatal("network error:", err)
//	}
//	log.Println("start rpc server on", l.Addr())
//	addr <- l.Addr().String()
//	geerpc.Accept(l)
//}
//
//func main() {
//	addr := make(chan string)
//	go startServer(addr)ls

//
//	// in fact, following code is like a simple geerpc client
//	// conn, _ := net.Dial("tcp", <-addr)
//
//	client, _ := geerpc.Dial("tcp", <- addr)
//	defer func() { _ = client.Close() }()
//
//	time.Sleep(time.Second)
//	// send options
//	//_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
//	//cc := codec.NewGobCodec(conn)
//
//	var wg sync.WaitGroup
//	// send request & receive response
//	for i := 0; i < 5; i++ {
//
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			args := fmt.Sprintf("geerpc req %d", i)
//			var reply string
//			if err := client.Call("Foo.Sum", args, &reply); err != nil {
//				log.Fatal("call Foo.Sum error:", err)
//			}
//			log.Println("reply:", reply)
//		}(i)
//		wg.Wait()
//
//
//		/*
//		h := &codec.Header{
//			ServiceMethod: "Foo.Sum",
//			Seq:           uint64(i),
//		}
//		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
//		_ = cc.ReadHeader(h)
//		var reply string
//		_ = cc.ReadBody(&reply)
//		log.Println("reply header:",  h)
//		log.Println("reply:", reply)
//		*/
//	}
//}

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addr chan string) {
	var foo Foo
	if err := geerpc.Register(&foo); err != nil {
		log.Fatal("register error:", err)
	}
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	geerpc.Accept(l)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)
	client, _ := geerpc.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	// send request & receive response
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			ctx := context.TODO()
			var reply int
			if err := client.Call(ctx, "Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}
