// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.

package modbus

// RtuViaTCPClientHandler implements Packager and Transporter interface.
type RtuViaTCPClientHandler struct {
	rtuPackager
	tcpTransporter
}

// NewRtuViaTCPClientHandler allocates a new RtuViaTCPClientHandler.
func NewRtuViaTCPClientHandler(address string) *RtuViaTCPClientHandler {
	h := &RtuViaTCPClientHandler{}
	h.Address = address
	h.Timeout = tcpTimeout
	h.IdleTimeout = tcpIdleTimeout
	return h
}

// RtuViaTCPClient creates TCP client with default handler and given connect string.
func RtuViaTCPClient(address string) Client {
	handler := NewRtuViaTCPClientHandler(address)
	return NewClient(handler)
}
