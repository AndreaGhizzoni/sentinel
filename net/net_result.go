package net

import "bytes"

type Result struct {
	Hosts []Host `xml:"host"`
}

func (r *Result) String() string {
	var buffer bytes.Buffer
	for i := 0; i < len(r.Hosts); i++ {
		buffer.WriteString(r.Hosts[i].String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

type Host struct {
	Status      *State    `xml:"status"`
	AddressList []Address `xml:"address"`
}

func (h *Host) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(h.Status.String() + "\n")
	for i := 0; i < len(h.AddressList); i++ {
		buffer.WriteString("  - " + h.AddressList[i].String() + "\n")
	}
	return buffer.String()
}

type State struct {
	Up     string `xml:"state,attr"`
	Reason string `xml:"reason,attr"`
}

func (s *State) String() string {
	return s.Up + " from " + s.Reason
}

type Address struct {
	Value  string `xml:"addr,attr"`
	Type   string `xml:"addrtype,attr"`
	Vendor string `xml:"vendor,attr"`
}

func (a *Address) String() string {
	return a.Type + " " + a.Value + " " + a.Vendor
}
