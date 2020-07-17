//Package Tars comment
// This file war generated by tars2go 1.1
// Generated from tarsregistry.tars
package Tars

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//AdapterConf strcut implement
type AdapterConf struct {
	Servant      string `json:"servant"`
	Endpoint     string `json:"endpoint"`
	Protocol     string `json:"protocol"`
	MaxConns     int32  `json:"maxConns"`
	ThreadNum    int32  `json:"threadNum"`
	QueueCap     int32  `json:"queueCap"`
	QueueTimeout int32  `json:"queueTimeout"`
}

func (st *AdapterConf) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *AdapterConf) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Servant, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Endpoint, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Protocol, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.MaxConns, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.ThreadNum, 4, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.QueueCap, 5, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.QueueTimeout, 6, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *AdapterConf) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require AdapterConf, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *AdapterConf) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Servant, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Endpoint, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Protocol, 2)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.MaxConns, 3)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.ThreadNum, 4)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.QueueCap, 5)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.QueueTimeout, 6)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *AdapterConf) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}