package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var i int

func init()  {
	i = 0
}

func TestStudent(t *testing.T) {
	s1 := NewStudent("John")
	s2 := NewStudent("Kevin")
	s3 := NewStudent("Jerry")
	s4 := NewStudent("Judy")

	users := []*Student{s1, s2, s3, s4}

	si := NewStudentIterator(users)

	testAssertIndex(t, 0, si.Index)
	testAssertStudent(t, s1, si.Next())
	testAssertIndex(t, 1, si.Index)
	testAssertStudent(t, s2, si.Next())
	testAssertIndex(t, 2, si.Index)
	testAssertStudent(t, s3, si.Next())
	testAssertIndex(t, 3, si.Index)
	testAssertStudent(t, s4, si.Next())

}

func testAssertStudent(t *testing.T, expected, actual *Student, )  {
	assert.Equal(t, expected, actual)
}

func testAssertIndex(t *testing.T, expected, actual int, )  {
	assert.Equal(t, expected, actual)
}
