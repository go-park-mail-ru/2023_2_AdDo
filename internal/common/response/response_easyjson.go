// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package response

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6ff3ac1dDecodeMainInternalCommonResponse(in *jlexer.Lexer, out *IsLiked) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "IsLiked":
			out.IsLiked = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6ff3ac1dEncodeMainInternalCommonResponse(out *jwriter.Writer, in IsLiked) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"IsLiked\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.IsLiked))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IsLiked) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeMainInternalCommonResponse(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IsLiked) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeMainInternalCommonResponse(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IsLiked) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeMainInternalCommonResponse(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IsLiked) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeMainInternalCommonResponse(l, v)
}
