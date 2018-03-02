/*
 * MIT License
 *
 * Copyright (c)  ShiChao
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/03/02        ShiChao
 */

package server

import (
	"encoding/binary"
	"bytes"
)

type message struct {
	msg []byte
}

func NewMsg(msg string) *message {
	return &message{
		[]byte(msg),
	}
}

func (m *message) Encode() ([]byte, error) {
	var err error

	m.msg, err = encode(m.msg)
	if err != nil {
		return []byte{}, err
	}

	return m.msg, nil
}

func (m *message) Decode(b []byte) error {
	var err error

	m.msg, err = decode(b)
	if err != nil {
		return err
	}

	return nil
}

func (m *message) Msg() []byte {
	return m.msg
}

func encode(b []byte) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, b)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func decode(b []byte) ([]byte, error) {
	var ret []byte
	buf := bytes.NewReader(b)

	err := binary.Read(buf, binary.BigEndian, ret)
	if err != nil {
		return []byte{}, err
	}
	return ret, nil
}
