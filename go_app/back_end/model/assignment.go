package model

type Assignment struct {
  Id   string `gorethink:"id,omitempty"`
  Name string `gorethink:"name"`
  Date string `gorethink:"date"`
}
