// Copyright 2024 peter.georgiev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package queue provides a simple queue implementation
Using channels and a mutex for synchronization.
It runs two parallel goroutines, one for adding elements to the queue and one for removing them.
This way the queue can be used in a thread-safe manner as a channel with a dynamic buffer size.

You can send data into the queue by using the In channel and receive data from the queue by using the Out channel.
The queue can be closed by calling the Close method, which will empty the queue and close the In and Out channels.
Examples:

	q := queue.New(context.Background())
	go func() {
		for i := 0; i < 10; i++ {
			q.In <- i
		}
		q.Close()
	}
	for {
		select {
		case i := <-q.Out:
			fmt.Println(i)
		case <-q.Done():
			return
		}
	}
*/
package queue
