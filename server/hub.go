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
	"sync"
	"bytes"
	"encoding/gob"
	"crypto/sha256"

	"github.com/fengyfei/tcp-zero/interfaces"
)

type hub struct {
	mu       sync.Mutex
	sessions map[string]interfaces.Session
}

func newHub() *hub {
	return &hub{
		sessions: make(map[string]interfaces.Session, 0),
	}
}

func (h *hub) Put(session interfaces.Session) error {
	serialized, err := serializeSess(session)
	if err != nil {
		return err
	}
	key := string(serialized)

	h.mu.Lock()
	defer h.mu.Unlock()
	if h.sessions[key] == nil {
		h.sessions[key] = session
	}

	return nil
}

func (h *hub) Remove(session interfaces.Session) error {
	serialized, err := serializeSess(session)
	if err != nil {
		return err
	}
	key := string(serialized)

	h.mu.Lock()
	defer h.mu.Unlock()

	sess := h.sessions[key]
	err = sess.Close()
	if err != nil {
		return err
	}

	delete(h.sessions, key)

	return nil
}

func (h *hub) Destroy() error {
	var err error

	h.mu.Lock()
	defer h.mu.Unlock()

	for k, v := range h.sessions {
		err = v.Close()
		if err != nil {
			return err
		}
		delete(h.sessions, k)
	}
	h = &hub{}

	return nil
}

func serializeSess(session interfaces.Session) ([]byte, error) {
	var res bytes.Buffer

	enc := gob.NewEncoder(&res)
	err := enc.Encode(session)
	if err != nil {
		return []byte{}, err
	}

	hash := sha256.Sum256(res.Bytes())
	return hash[:], nil
}
