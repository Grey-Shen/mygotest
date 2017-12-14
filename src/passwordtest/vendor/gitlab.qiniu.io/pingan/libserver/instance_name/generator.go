package instance_name

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"
	"net"
)

func GenerateInstanceName(source string) (string, error) {
	var (
		buf        = bytes.NewBuffer(make([]byte, 0, 1024))
		interfaces []net.Interface
		addrs      []net.Addr
		err        error
	)
	if _, err = buf.WriteString(source); err != nil {
		return "", err
	}

	if interfaces, err = net.Interfaces(); err != nil {
		return "", err
	} else {
		for _, interf := range interfaces {
			if _, err = buf.Write(interf.HardwareAddr); err != nil {
				return "", err
			}
		}
	}

	if addrs, err = net.InterfaceAddrs(); err != nil {
		return "", err
	} else {
		for _, addr := range addrs {
			if _, err = buf.WriteString(addr.Network()); err != nil {
				return "", err
			}
			if _, err = buf.WriteString(addr.String()); err != nil {
				return "", err
			}
		}
	}

	hash := crc32.ChecksumIEEE(buf.Bytes())
	buf.Reset()
	if err = binary.Write(buf, binary.LittleEndian, hash); err != nil {
		return "", err
	} else {
		return hex.EncodeToString(buf.Bytes()), nil
	}
}
