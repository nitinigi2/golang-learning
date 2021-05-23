Sending to a  nil channel blocks forever

Receiving from nil channel blocks forever

you will get ```fatal error: all goroutines are asleep - deadlock``` if you do so

If you try to send message from a closed channel it will ```panic```

If you try to recieve from a closed channel, it will give you ```default value``` i.e ```0``` if channel type is int