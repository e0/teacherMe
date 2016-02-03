package model

type Course struct {
  Id            string       `gorethink:"id,omitempty"`
  Title         string       `gorethink:"title"`
  Description   string       `gorethink:"description"`
  Discussions   []Discussion `gorethink:"discussions"`
  Downloads     []Download   `gorethink:"downloads"`
  Assignments   []Assignment `gorethink:"assignments"`
  TeacherName   string       `gorethink:"teacherName"`
  StudentsCount int          `gorethink:"studentsCount"`
}