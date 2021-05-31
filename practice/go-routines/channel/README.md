In Go language, a ```channel``` is a medium through which a goroutine communicates with another goroutine and this communication is lock-free.
```By default channel is bidirectional```

Create a channel

```channel_name:= make(chan Type)```

Channel sends a copy to targeted go-routines.

# Types

1) Unbuffered channel

   ```channel_name:= make(chan Type, buffer_size)```

2) Buffered channel

   ```channel_name:= make(chan Type)```

# Types

1) bidirectional channel

        func funcName(ch chan type) { // something}

2) send-only channel

        func funcName(ch chan<- type) { //something}

2) recieve-only channel

        func funcName(ch <-chan type) { //something}


# Important Points

1) Blocking Send and Receive: In the channel when the data sent to a channel the control is blocked in that send statement until other goroutine reads from that channel. Similarly, when a channel receives data from the goroutine the read statement block until another goroutine statement.

2) Zero Value Channel: The zero value of the channel is nil.

3) For loop in Channel: A for loop can iterate over the sequential values sent on the channel until it closed.
Syntax:

       for item := range Chnl { 
          // statements..
       }

# Important links

https://www.geeksforgeeks.org/channel-in-golang/

https://golang.org/pkg/sync/#RWMutex