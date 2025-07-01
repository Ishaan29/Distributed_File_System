package main

import (
	"fmt"
	"log"

	"github.com/ishaan29/Distributed_File_System/p2p"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	fmt.Println("logic with the peer outside of the transport")
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("message received: %+v\n", msg)
		}
	}()
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
	//fmt.Println("Hello!, welcome to the world of go.")
}

// sandbox code // TODO: remove this

// type Task struct {
// 	ID   int
// 	msg  string
// 	done chan bool
// }

// func scheduler() chan Task {
// 	queue := make(chan Task)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Printf("Scheduler: sending task %d\n", i)
// 			q := Task{
// 				ID:   i,
// 				msg:  fmt.Sprintf("Task %d", i),
// 				done: make(chan bool),
// 			}
// 			queue <- q
// 			// time.Sleep(time.Second * 2)
// 		}
// 		close(queue)
// 	}()
// 	return queue
// }

// func worker(id int, queue <-chan Task, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for task := range queue {
// 		fmt.Printf("Worker %d received task %d: %s\n", id, task.ID, task.msg)
// 		time.Sleep(time.Second * 2)
// 		task.done <- true
// 		close(task.done)
// 	}
// 	fmt.Printf("Worker %d finished\n", id)
// }
// func main() {
// 	queue := scheduler()
// 	var wg sync.WaitGroup
// 	for i := 0; i < 1; i++ {
// 		wg.Add(1)
// 		go worker(i, queue, &wg)
// 	}

// 	wg.Wait()
// }
