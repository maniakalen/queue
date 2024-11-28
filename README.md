# go-queue

Package queue provides a simple queue implementation
Using channels and a mutex for synchronization.
It runs two parallel goroutines, one for adding elements to the queue and one for removing them.
This way the queue can be used in a thread-safe manner as a channel with a dynamic buffer size.