package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var r *Basic

func init() {
	r = new(Basic)
	r.resumeType = "basic"
	r.name = "王大明"
	r.age = 12
}

func TestCloneResume(t *testing.T) {
	rIT := r.clone()
	rIT.resumeType = "IT"
	rIT.age = 11

	assert.Equal(t, r, &Basic{
		name:       "王大明",
		age:        12,
		resumeType: "basic",
	})

	assert.Equal(t, rIT, &Basic{
		name:       "王大明",
		age:        11,
		resumeType: "IT",
	})
}

func TestCopyResume(t *testing.T) {
	rIT := r.copy()
	rIT.resumeType = "IT"
	rIT.age = 11

	assert.Equal(t, r, &Basic{
		name:       "王大明",
		age:        12,
		resumeType: "basic",
	})

	assert.Equal(t, rIT, Basic{
		name:       "王大明",
		age:        11,
		resumeType: "IT",
	})
}
