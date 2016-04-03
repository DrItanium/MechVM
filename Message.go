// Message passing interface
package main

type MessageType int32

type Messenger interface {
	send(*Messenger, MessageType) bool
	receive(*Messenger, MessageType)
}
