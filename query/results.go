package query

import "github.com/golang/protobuf/descriptor"

type ResultSet interface {
	Len() int
	At(int) interface{}
	Insert(interface{}, int)
	Append(interface{})
}

type SliceResultSet struct {
	Items []descriptor.Message
}

func NewSliceResultSet() *SliceResultSet {
	return &SliceResultSet{
		make([]descriptor.Message, 0),
	}
}

func (rs *SliceResultSet) Len() int {
	return len(rs.Items)
}

func (rs *SliceResultSet) At(index int) interface{} {
	return rs.Items[index]
}

func (rs *SliceResultSet) Insert(item interface{}, index int) {
	rs.Items = append(rs.Items, nil)
	copy(rs.Items[index+1:], rs.Items[index:])
	rs.Items[index] = item.(descriptor.Message)
}

func (rs *SliceResultSet) Append(item interface{}) {
	rs.Items = append(rs.Items, item.(descriptor.Message))
}
