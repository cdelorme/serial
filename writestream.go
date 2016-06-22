package transport

import "encoding/binary"

type WriteStream struct {
	Data []byte
}

func (self *WriteStream) SerializeString(out *string) error {
	binary.Write(self, ByteOrder, int64(len(*out)))
	self.Data = append(self.Data, []byte(*out)...)
	return nil
}

func (self *WriteStream) SerializeInt(out *int) error {
	return binary.Write(self, ByteOrder, int64(*out))
}

func (self *WriteStream) Write(p []byte) (int, error) {
	self.Data = append(self.Data, p...)
	return len(p), nil
}
