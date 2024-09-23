// Copyright 2021-2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package comid

const MAX_TCB = 16

type TcbElement struct {
	Svn      uint   `cbor:"-1,keyasint,omitempty" json:"svn,omitempty"`
	Category string `cbor:"-2,keyasint,omitempty" json:"category,omitempty"`
	TcbType  string `cbor:"-3,keyasint,omitempty" json:"tcbType,omitempty"`
}

type TcbComponent struct {
	Element [MAX_TCB]TcbElement
	Pcesvn  *uint
}

type TcbLevel struct {
	Tcb       []TcbComponent `cbor:"-5,keyasint,omitempty" json:"tcb,omitempty"`
	TcbDate   string
	TcbStatus string
}

type TcbLevels []TcbLevel
