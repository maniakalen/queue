// Copyright 2024 peter.georgiev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package queue provides a simple queue implementation
Using channels and a mutex for synchronization.
It runs two parallel goroutines, one for adding elements to the queue and one for removing them.
This way the queue can be used in a thread-safe manner as a channel with a dynamic buffer size.
*/
package queue
