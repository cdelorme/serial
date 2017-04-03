package benchmarks

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Entity) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Health":
			var zhct uint32
			zhct, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if zhct != 2 {
				err = msgp.ArrayError{Wanted: 2, Got: zhct}
				return
			}
			for zxvk := range z.Health {
				z.Health[zxvk], err = dc.ReadUint16()
				if err != nil {
					return
				}
			}
		case "Mana":
			var zcua uint32
			zcua, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if zcua != 2 {
				err = msgp.ArrayError{Wanted: 2, Got: zcua}
				return
			}
			for zbzg := range z.Mana {
				z.Mana[zbzg], err = dc.ReadUint16()
				if err != nil {
					return
				}
			}
		case "Stamina":
			var zxhx uint32
			zxhx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if zxhx != 2 {
				err = msgp.ArrayError{Wanted: 2, Got: zxhx}
				return
			}
			for zbai := range z.Stamina {
				z.Stamina[zbai], err = dc.ReadUint16()
				if err != nil {
					return
				}
			}
		case "Friends":
			var zlqf uint32
			zlqf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Friends) >= int(zlqf) {
				z.Friends = (z.Friends)[:zlqf]
			} else {
				z.Friends = make([]int64, zlqf)
			}
			for zcmr := range z.Friends {
				z.Friends[zcmr], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "Statuses":
			var zdaf uint32
			zdaf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Statuses) >= int(zdaf) {
				z.Statuses = (z.Statuses)[:zdaf]
			} else {
				z.Statuses = make([]Status, zdaf)
			}
			for zajw := range z.Statuses {
				err = z.Statuses[zajw].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Dead":
			z.Dead, err = dc.ReadBool()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Entity) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "Name"
	err = en.Append(0x87, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Health"
	err = en.Append(0xa6, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(2)
	if err != nil {
		return
	}
	for zxvk := range z.Health {
		err = en.WriteUint16(z.Health[zxvk])
		if err != nil {
			return
		}
	}
	// write "Mana"
	err = en.Append(0xa4, 0x4d, 0x61, 0x6e, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(2)
	if err != nil {
		return
	}
	for zbzg := range z.Mana {
		err = en.WriteUint16(z.Mana[zbzg])
		if err != nil {
			return
		}
	}
	// write "Stamina"
	err = en.Append(0xa7, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(2)
	if err != nil {
		return
	}
	for zbai := range z.Stamina {
		err = en.WriteUint16(z.Stamina[zbai])
		if err != nil {
			return
		}
	}
	// write "Friends"
	err = en.Append(0xa7, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Friends)))
	if err != nil {
		return
	}
	for zcmr := range z.Friends {
		err = en.WriteInt64(z.Friends[zcmr])
		if err != nil {
			return
		}
	}
	// write "Statuses"
	err = en.Append(0xa8, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Statuses)))
	if err != nil {
		return
	}
	for zajw := range z.Statuses {
		err = z.Statuses[zajw].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Dead"
	err = en.Append(0xa4, 0x44, 0x65, 0x61, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Dead)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Entity) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Name"
	o = append(o, 0x87, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Health"
	o = append(o, 0xa6, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68)
	o = msgp.AppendArrayHeader(o, 2)
	for zxvk := range z.Health {
		o = msgp.AppendUint16(o, z.Health[zxvk])
	}
	// string "Mana"
	o = append(o, 0xa4, 0x4d, 0x61, 0x6e, 0x61)
	o = msgp.AppendArrayHeader(o, 2)
	for zbzg := range z.Mana {
		o = msgp.AppendUint16(o, z.Mana[zbzg])
	}
	// string "Stamina"
	o = append(o, 0xa7, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61)
	o = msgp.AppendArrayHeader(o, 2)
	for zbai := range z.Stamina {
		o = msgp.AppendUint16(o, z.Stamina[zbai])
	}
	// string "Friends"
	o = append(o, 0xa7, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Friends)))
	for zcmr := range z.Friends {
		o = msgp.AppendInt64(o, z.Friends[zcmr])
	}
	// string "Statuses"
	o = append(o, 0xa8, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Statuses)))
	for zajw := range z.Statuses {
		o, err = z.Statuses[zajw].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Dead"
	o = append(o, 0xa4, 0x44, 0x65, 0x61, 0x64)
	o = msgp.AppendBool(o, z.Dead)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Entity) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpks uint32
	zpks, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpks > 0 {
		zpks--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Health":
			var zjfb uint32
			zjfb, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if zjfb != 2 {
				err = msgp.ArrayError{Wanted: 2, Got: zjfb}
				return
			}
			for zxvk := range z.Health {
				z.Health[zxvk], bts, err = msgp.ReadUint16Bytes(bts)
				if err != nil {
					return
				}
			}
		case "Mana":
			var zcxo uint32
			zcxo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if zcxo != 2 {
				err = msgp.ArrayError{Wanted: 2, Got: zcxo}
				return
			}
			for zbzg := range z.Mana {
				z.Mana[zbzg], bts, err = msgp.ReadUint16Bytes(bts)
				if err != nil {
					return
				}
			}
		case "Stamina":
			var zeff uint32
			zeff, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if zeff != 2 {
				err = msgp.ArrayError{Wanted: 2, Got: zeff}
				return
			}
			for zbai := range z.Stamina {
				z.Stamina[zbai], bts, err = msgp.ReadUint16Bytes(bts)
				if err != nil {
					return
				}
			}
		case "Friends":
			var zrsw uint32
			zrsw, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Friends) >= int(zrsw) {
				z.Friends = (z.Friends)[:zrsw]
			} else {
				z.Friends = make([]int64, zrsw)
			}
			for zcmr := range z.Friends {
				z.Friends[zcmr], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "Statuses":
			var zxpk uint32
			zxpk, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Statuses) >= int(zxpk) {
				z.Statuses = (z.Statuses)[:zxpk]
			} else {
				z.Statuses = make([]Status, zxpk)
			}
			for zajw := range z.Statuses {
				bts, err = z.Statuses[zajw].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Dead":
			z.Dead, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Entity) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + msgp.ArrayHeaderSize + (2 * (msgp.Uint16Size)) + 5 + msgp.ArrayHeaderSize + (2 * (msgp.Uint16Size)) + 8 + msgp.ArrayHeaderSize + (2 * (msgp.Uint16Size)) + 8 + msgp.ArrayHeaderSize + (len(z.Friends) * (msgp.Int64Size)) + 9 + msgp.ArrayHeaderSize
	for zajw := range z.Statuses {
		s += z.Statuses[zajw].Msgsize()
	}
	s += 5 + msgp.BoolSize
	return
}
