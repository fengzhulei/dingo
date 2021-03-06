package dingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskEqual(t *testing.T) {
	ass := assert.New(t)

	// same
	{
		m1 := map[string]string{
			"t1": "1",
			"t2": "2",
		}

		m2 := map[string]string{
			"t1": "1",
			"t2": "2",
		}

		t, err := composeTask("name#1", nil, []interface{}{1, "test123", m1})
		ass.Nil(err)

		o, err := composeTask("name#1", nil, []interface{}{1, "test123", m2})
		ass.Nil(err)

		o.H.I = t.H.ID()
		ass.Equal(o, t)
	}

	// diff map
	{
		t, err := composeTask("name#1", nil, []interface{}{1, "test123", map[string]string{
			"t1": "1",
			"t2": "2",
		}})
		ass.Nil(err)

		o, err := composeTask("name#1", nil, []interface{}{1, "test123", map[string]string{
			"t2": "2",
			"t3": "3",
		}})
		ass.Nil(err)

		o.H.I = t.H.ID()
		ass.NotEqual(t, o)
	}

	// only Name is different
	{
		t, err := composeTask("name#1", nil, []interface{}{1, "test#123"})
		ass.Nil(err)

		o, err := composeTask("name#2", nil, []interface{}{1, "test#123"})
		ass.Nil(err)

		o.H.I = t.H.ID()
		ass.NotEqual(t, o)
	}

	// sequence of args is different
	{
		t, err := composeTask("name#1", nil, []interface{}{1, "test#123"})
		ass.Nil(err)

		o, err := composeTask("name#1", nil, []interface{}{"test#123", 1})
		ass.Nil(err)

		o.H.I = t.H.ID()
		ass.NotEqual(t, o)
	}

	// different args
	{
		t, err := composeTask("name#1", nil, []interface{}{1, "test#123"})
		ass.Nil(err)

		o, err := composeTask("name#1", nil, []interface{}{2, "test#123"})
		ass.Nil(err)

		o.H.I = t.H.ID()
		ass.NotEqual(t, o)
	}

	// different option
	{
		t, err := composeTask("name#1", DefaultOption().IgnoreReport(true), []interface{}{1, "test#123"})
		ass.Nil(err)

		o, err := composeTask("name#1", DefaultOption().IgnoreReport(false), []interface{}{2, "test#123"})
		ass.Nil(err)

		o.H.I = t.H.ID()
		ass.NotEqual(t, o)
	}
}
