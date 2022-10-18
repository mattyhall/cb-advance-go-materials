# Channel tasks

1. Let's model a heartbeat server. To check whether a connection is working it is common for applications to implement a "heartbeat". In this a client will send a message (often called PING) periodically (maybe every 30s) and the server will respond (PONG).

Create two goroutines - one a client and one a server. The client should send a PING message on a channel every 30s and wait for the server to reply PONG on a different channel. The server should respond to PINGs. Have some `fmt.Println`s so you know what is happening

2. Really we want the server to detect when the client hasn't sent a PING when we think it should have. Use a `select` statement to time out receiving on the channel after 31s and exit. Modify your client to occasionally send the PING too slowly

3. Add a buffer to the channel the client sends on. Why is this incorrect for heartbeats? Can you demonstrate this?

4. Remove the timeout and have your server listen to two different clients
