// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package artist

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	album "main/internal/pkg/album"
	playlist "main/internal/pkg/playlist"
	track "main/internal/pkg/track"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3e1fa5ecDecodeMainInternalPkgArtist(in *jlexer.Lexer, out *SearchResponse) {
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
		case "Playlists":
			if in.IsNull() {
				in.Skip()
				out.Playlists = nil
			} else {
				in.Delim('[')
				if out.Playlists == nil {
					if !in.IsDelim(']') {
						out.Playlists = make([]playlist.Base, 0, 1)
					} else {
						out.Playlists = []playlist.Base{}
					}
				} else {
					out.Playlists = (out.Playlists)[:0]
				}
				for !in.IsDelim(']') {
					var v1 playlist.Base
					(v1).UnmarshalEasyJSON(in)
					out.Playlists = append(out.Playlists, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Albums":
			if in.IsNull() {
				in.Skip()
				out.Albums = nil
			} else {
				in.Delim('[')
				if out.Albums == nil {
					if !in.IsDelim(']') {
						out.Albums = make([]album.Base, 0, 1)
					} else {
						out.Albums = []album.Base{}
					}
				} else {
					out.Albums = (out.Albums)[:0]
				}
				for !in.IsDelim(']') {
					var v2 album.Base
					(v2).UnmarshalEasyJSON(in)
					out.Albums = append(out.Albums, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
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
					var v3 track.Response
					(v3).UnmarshalEasyJSON(in)
					out.Tracks = append(out.Tracks, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Artists":
			if in.IsNull() {
				in.Skip()
				out.Artists = nil
			} else {
				in.Delim('[')
				if out.Artists == nil {
					if !in.IsDelim(']') {
						out.Artists = make([]Base, 0, 1)
					} else {
						out.Artists = []Base{}
					}
				} else {
					out.Artists = (out.Artists)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Base
					(v4).UnmarshalEasyJSON(in)
					out.Artists = append(out.Artists, v4)
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
func easyjson3e1fa5ecEncodeMainInternalPkgArtist(out *jwriter.Writer, in SearchResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Playlists\":"
		out.RawString(prefix[1:])
		if in.Playlists == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Playlists {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Albums\":"
		out.RawString(prefix)
		if in.Albums == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Albums {
				if v7 > 0 {
					out.RawByte(',')
				}
				(v8).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Tracks\":"
		out.RawString(prefix)
		if in.Tracks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Tracks {
				if v9 > 0 {
					out.RawByte(',')
				}
				(v10).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Artists\":"
		out.RawString(prefix)
		if in.Artists == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Artists {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e1fa5ecEncodeMainInternalPkgArtist(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e1fa5ecEncodeMainInternalPkgArtist(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e1fa5ecDecodeMainInternalPkgArtist(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e1fa5ecDecodeMainInternalPkgArtist(l, v)
}
func easyjson3e1fa5ecDecodeMainInternalPkgArtist1(in *jlexer.Lexer, out *Response) {
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
		case "Avatar":
			out.Avatar = string(in.String())
		case "Albums":
			if in.IsNull() {
				in.Skip()
				out.Albums = nil
			} else {
				in.Delim('[')
				if out.Albums == nil {
					if !in.IsDelim(']') {
						out.Albums = make([]album.Base, 0, 1)
					} else {
						out.Albums = []album.Base{}
					}
				} else {
					out.Albums = (out.Albums)[:0]
				}
				for !in.IsDelim(']') {
					var v13 album.Base
					(v13).UnmarshalEasyJSON(in)
					out.Albums = append(out.Albums, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
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
					var v14 track.Response
					(v14).UnmarshalEasyJSON(in)
					out.Tracks = append(out.Tracks, v14)
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
func easyjson3e1fa5ecEncodeMainInternalPkgArtist1(out *jwriter.Writer, in Response) {
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
		const prefix string = ",\"Avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"Albums\":"
		out.RawString(prefix)
		if in.Albums == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v15, v16 := range in.Albums {
				if v15 > 0 {
					out.RawByte(',')
				}
				(v16).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Tracks\":"
		out.RawString(prefix)
		if in.Tracks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Tracks {
				if v17 > 0 {
					out.RawByte(',')
				}
				(v18).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e1fa5ecEncodeMainInternalPkgArtist1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e1fa5ecEncodeMainInternalPkgArtist1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e1fa5ecDecodeMainInternalPkgArtist1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e1fa5ecDecodeMainInternalPkgArtist1(l, v)
}
func easyjson3e1fa5ecDecodeMainInternalPkgArtist2(in *jlexer.Lexer, out *Base) {
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
		case "Avatar":
			out.Avatar = string(in.String())
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
func easyjson3e1fa5ecEncodeMainInternalPkgArtist2(out *jwriter.Writer, in Base) {
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
		const prefix string = ",\"Avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Base) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e1fa5ecEncodeMainInternalPkgArtist2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Base) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e1fa5ecEncodeMainInternalPkgArtist2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Base) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e1fa5ecDecodeMainInternalPkgArtist2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Base) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e1fa5ecDecodeMainInternalPkgArtist2(l, v)
}
func easyjson3e1fa5ecDecodeMainInternalPkgArtist3(in *jlexer.Lexer, out *Artists) {
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
		case "Artists":
			if in.IsNull() {
				in.Skip()
				out.Artists = nil
			} else {
				in.Delim('[')
				if out.Artists == nil {
					if !in.IsDelim(']') {
						out.Artists = make([]Base, 0, 1)
					} else {
						out.Artists = []Base{}
					}
				} else {
					out.Artists = (out.Artists)[:0]
				}
				for !in.IsDelim(']') {
					var v19 Base
					(v19).UnmarshalEasyJSON(in)
					out.Artists = append(out.Artists, v19)
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
func easyjson3e1fa5ecEncodeMainInternalPkgArtist3(out *jwriter.Writer, in Artists) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Artists\":"
		out.RawString(prefix[1:])
		if in.Artists == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.Artists {
				if v20 > 0 {
					out.RawByte(',')
				}
				(v21).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Artists) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e1fa5ecEncodeMainInternalPkgArtist3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Artists) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e1fa5ecEncodeMainInternalPkgArtist3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Artists) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e1fa5ecDecodeMainInternalPkgArtist3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Artists) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e1fa5ecDecodeMainInternalPkgArtist3(l, v)
}
