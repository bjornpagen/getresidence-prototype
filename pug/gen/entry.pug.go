// Code generated by "jade.go"; DO NOT EDIT.

package gen

import (
	"io"
)

const (
	entry__0  = `<div id="entry" class="`
	entry__1  = `"><label for="`
	entry__2  = `">`
	entry__3  = `</label><input id="`
	entry__4  = `" type="`
	entry__5  = `" name="`
	entry__6  = `" value="`
	entry__7  = `" required="required" hx-put="`
	entry__8  = `" hx-swap="outerHTML" hx-sync="input:queue" hx-target="closest #entry"/>`
	entry__9  = `</div>`
	entry__10 = `<small>`
	entry__11 = `</small>`
)

func Jade_entry(s struct {
	Name     string
	Value    string
	Schema   string
	State    string
	Small    string
	Endpoint string
}, wr io.Writer) {
	buffer := &WriterAsBuffer{wr}

	buffer.WriteString(entry__0)
	WriteEscString(s.State, buffer)
	buffer.WriteString(entry__1)
	WriteEscString(s.Name, buffer)
	buffer.WriteString(entry__2)
	WriteEscString(s.Name, buffer)
	buffer.WriteString(entry__3)
	WriteEscString(s.Name, buffer)
	buffer.WriteString(entry__4)
	WriteEscString(s.Schema, buffer)
	buffer.WriteString(entry__5)
	WriteEscString(s.Name, buffer)
	buffer.WriteString(entry__6)
	WriteEscString(s.Value, buffer)
	buffer.WriteString(entry__7)
	WriteEscString(s.Endpoint, buffer)
	buffer.WriteString(entry__8)
	if len(s.Small) > 0 {
		buffer.WriteString(entry__10)
		WriteEscString(s.Small, buffer)
		buffer.WriteString(entry__11)
	}
	buffer.WriteString(entry__9)

}