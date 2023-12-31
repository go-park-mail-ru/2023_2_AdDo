// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package album

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	track "main/internal/pkg/track"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3e1fa5ecDecodeMainInternalPkgAlbum(in *jlexer.Lexer, out *Response) {
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
		case "Id":
			out.Id = uint64(in.Uint64())
		case "Name":
			out.Name = string(in.String())
		case "Preview":
			out.Preview = string(in.String())
		case "ArtistId":
			out.ArtistId = uint64(in.Uint64())
		case "ArtistName":
			out.ArtistName = string(in.String())
		case "Tracks":
			if in.IsNull() {
				in.Skip()
				out.Tracks = nil
			} else {
				in.Delim('[')
				if out.Tracks == nil {
					if !in.IsDelim(']') {
						out.Tracks = make([]track.Response, 0, 0)
					} else {
						out.Tracks = []track.Response{}
					}
				} else {
					out.Tracks = (out.Tracks)[:0]
				}
				for !in.IsDelim(']') {
					var v1 track.Response
					(v1).UnmarshalEasyJSON(in)
					out.Tracks = append(out.Tracks, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson3e1fa5ecEncodeMainInternalPkgAlbum(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Preview\":"
		out.RawString(prefix)
		out.String(string(in.Preview))
	}
	{
		const prefix string = ",\"ArtistId\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ArtistId))
	}
	{
		const prefix string = ",\"ArtistName\":"
		out.RawString(prefix)
		out.String(string(in.ArtistName))
	}
	{
		const prefix string = ",\"Tracks\":"
		out.RawString(prefix)
		if in.Tracks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Tracks {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e1fa5ecEncodeMainInternalPkgAlbum(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e1fa5ecEncodeMainInternalPkgAlbum(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e1fa5ecDecodeMainInternalPkgAlbum(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e1fa5ecDecodeMainInternalPkgAlbum(l, v)
}
func easyjson3e1fa5ecDecodeMainInternalPkgAlbum1(in *jlexer.Lexer, out *LikedAlbums) {
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
		case "Albums":
			if in.IsNull() {
				in.Skip()
				out.Albums = nil
			} else {
				in.Delim('[')
				if out.Albums == nil {
					if !in.IsDelim(']') {
						out.Albums = make([]Base, 0, 1)
					} else {
						out.Albums = []Base{}
					}
				} else {
					out.Albums = (out.Albums)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Base
					(v4).UnmarshalEasyJSON(in)
					out.Albums = append(out.Albums, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson3e1fa5ecEncodeMainInternalPkgAlbum1(out *jwriter.Writer, in LikedAlbums) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Albums\":"
		out.RawString(prefix[1:])
		if in.Albums == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Albums {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LikedAlbums) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e1fa5ecEncodeMainInternalPkgAlbum1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LikedAlbums) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e1fa5ecEncodeMainInternalPkgAlbum1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LikedAlbums) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e1fa5ecDecodeMainInternalPkgAlbum1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LikedAlbums) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e1fa5ecDecodeMainInternalPkgAlbum1(l, v)
}
func easyjson3e1fa5ecDecodeMainInternalPkgAlbum2(in *jlexer.Lexer, out *Id) {
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
		case "Id":
			out.Id = uint64(in.Uint64())
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
func easyjson3e1fa5ecEncodeMainInternalPkgAlbum2(out *jwriter.Writer, in Id) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Id) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e1fa5ecEncodeMainInternalPkgAlbum2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Id) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e1fa5ecEncodeMainInternalPkgAlbum2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Id) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e1fa5ecDecodeMainInternalPkgAlbum2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Id) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e1fa5ecDecodeMainInternalPkgAlbum2(l, v)
}
func easyjson3e1fa5ecDecodeMainInternalPkgAlbum3(in *jlexer.Lexer, out *Base) {
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
		case "Id":
			out.Id = uint64(in.Uint64())
		case "Name":
			out.Name = string(in.String())
		case "Preview":
			out.Preview = string(in.String())
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
func easyjson3e1fa5ecEncodeMainInternalPkgAlbum3(out *jwriter.Writer, in Base) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Preview\":"
		out.RawString(prefix)
		out.String(string(in.Preview))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Base) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e1fa5ecEncodeMainInternalPkgAlbum3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Base) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e1fa5ecEncodeMainInternalPkgAlbum3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Base) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e1fa5ecDecodeMainInternalPkgAlbum3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Base) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e1fa5ecDecodeMainInternalPkgAlbum3(l, v)
}
func easyjson3e1fa5ecDecodeMainInternalPkgAlbum4(in *jlexer.Lexer, out *Albums) {
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
		case "Albums":
			if in.IsNull() {
				in.Skip()
				out.Albums = nil
			} else {
				in.Delim('[')
				if out.Albums == nil {
					if !in.IsDelim(']') {
						out.Albums = make([]Response, 0, 0)
					} else {
						out.Albums = []Response{}
					}
				} else {
					out.Albums = (out.Albums)[:0]
				}
				for !in.IsDelim(']') {
					var v7 Response
					(v7).UnmarshalEasyJSON(in)
					out.Albums = append(out.Albums, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson3e1fa5ecEncodeMainInternalPkgAlbum4(out *jwriter.Writer, in Albums) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Albums\":"
		out.RawString(prefix[1:])
		if in.Albums == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Albums {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Albums) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e1fa5ecEncodeMainInternalPkgAlbum4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Albums) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e1fa5ecEncodeMainInternalPkgAlbum4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Albums) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e1fa5ecDecodeMainInternalPkgAlbum4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Albums) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e1fa5ecDecodeMainInternalPkgAlbum4(l, v)
}
