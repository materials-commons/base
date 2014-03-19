package protocol

import (
	"bytes"
	"github.com/ugorji/go/codec"
)

/*
The following code encodes and decodes buffers of bytes using MessagePack. It uses the approach
found in github.com/hashicorp/serf for identifying the type of message. The buffer has
a message type prepended to it. In our case we also prepend a version so that multiple
protocol versions can be supported.
*/

// Encode encodes a message using MsgPack. It prepends the message type and the CurrentVersion
// to the returned buffer.
func Encode(msgType MessageType, in interface{}) (*bytes.Buffer, error) {
	return EncodeVersion(msgType, CurrentVersion, in)
}

// EncodeVersion encodes a message using MessagePack. It prepends the message type and the passed in
// version to the returned buffer.
func EncodeVersion(msgType MessageType, version uint8, in interface{}) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(uint8(msgType))
	buf.WriteByte(version)
	handle := codec.MsgpackHandle{}
	encoder := codec.NewEncoder(buf, &handle)
	err := encoder.Encode(in)
	return buf, err
}

// Decode decodes a buffer using MessagePack. The buffer passed in needs to have removed the
// message type and version that were passed in.
func Decode(buf []byte, out interface{}) error {
	reader := bytes.NewReader(buf)
	handle := codec.MsgpackHandle{}
	decoder := codec.NewDecoder(reader, &handle)
	return decoder.Decode(out)
}

// Prepare retrieves the message type and version, and a buffer that is ready to be
// sent to Decode.
func Prepare(buf []byte) (msgType MessageType, version uint8, b []byte) {
	msgType = MessageType(buf[0])
	version = uint8(buf[1])
	return msgType, version, buf[2:]
}
