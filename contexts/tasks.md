1. Change your heartbeat server to use a cancellable context. Have the client cancel after two iterations

2. Pass how often to send the PING as a value in the context. Use this value plus a leeway of two seconds in the server
