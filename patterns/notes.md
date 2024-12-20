# Concurrency Patterns: 


## The for select Loop:

* This patterns will be seen all over Go projects, and it looks someting like this: 

```
for { // Either inf loop or over somthing (using range)
	select { // Do some work with channels.

	}
}
```

1. Sending iteration vars out on a channel.
2. Looping inf waiting to be preempted(stopped).

## Preventing Goroutines Leaks:

* we mentioned before that goroutines are more effiecnt compared to Green-threads, and run-time managed threads. But they are not **free**, they do cost resources, and they *are not collected by the garbege collected*.

* Goroutines has few paths to termination: 
	1. When it has comleted its work.
	2. When it cannot continue its work due to an unrecoverable error.
	3. When it's told to stop working.

we get the first two patterns for free, but the third one must managed by us(work cancelation).

* Goroutines work in collaboration. So we could represent the interdependece between them using *Graph*. The root of this graph, can be said or called the **main goroutine**.


* To prevent that from happening, we can mitigate to establish a signal between the parent goroutine and its children that allows the parent to signal cancellation to its children.	
