package model

type Discussion struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
	Date string `gorethink:"date"`
}
