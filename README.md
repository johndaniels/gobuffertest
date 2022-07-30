Simple http programs to show how network buffers behave in Go.

The client reads 128k once every second, while the server blocks in a loop writing 16 megabytes one megabyte at a time. If you run the server first, and then the client, the server pretty immediately writes 1 MB, which is being stored in kernel or userspace network buffers (unclear whether this is the client's or server's buffer. Testing this with remote machines and something to simulate network latency would be interesting). The client in this instance almost never blocks when reading, meaning that the 128k it asks for is almost always already available in a network buffer.

The client's slowness puts backpressure on the server sending down the 16 megabyte file, which is what causes the server to slowly send 1 megabyte at a time. 

Further experiments that would be interesting are messing with the server to cause the client to wait on reading instead, playing with buffer sizes in general, perhaps by having one side or the other write a single byte at a time, and interposing a proxy that simulates network latency.