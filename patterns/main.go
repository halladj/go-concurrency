package main

import "fmt"


func main(){
	/*
	For-Select-Pattern:
	*/ 
	//Sending iteration vars out on a channel.
	for _, s := range []string{"a", "b", "c"}{
		select {
		case <- done:
			return
		case stringStream <- s:

		}
	}

	// Looping inf waiting to be preempted(stopped). 
	for {
		select {
		case <- done:
			return 
		default:
		}// do not non-preempetd.
	}

	/*
	Preventing Goroutines Leaks:
	*/

	doWork := func (strings <-chan string) <- interface{} {
		completed := make(chan interface{})
		go func () {
			defer fmt.Println("DoWork exited")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil) // since we are passing a nil channel, then
		    // this goroutine will never halt, and will be
		    // in-memory forever.
	fmt.Println("Done")


	// THE PROPER WAY TO HANDLE CANCELATION.
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	)<-chan interface{} {

		terminated := make (chan interface{})
		go func(){
			defer fmt.Println("DoWork exited.")
			defer close(terminated)

			for {
				select {
				case s:= <-strings:
					fmt.Println(s)

				case <- done:
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated = doWork()

	go func(){
		time.Sleep(2 * time.Second)
		fmt.Println("Canceling dowork goroutines ...")
		close(done)
	}()

	<- terminated
	fmt.Println("Done.")

}
