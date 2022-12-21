package plerrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetNestedNew(t *testing.T) {
	jsn := MetaData{
		"a": MetaData{
			"b": MetaData{"c": "d"},
			"l": "m",
			"n": []MetaData{{"o": "p"}, {"q": "r"}},
			"x": []string{"y", "z"},
		},
		"b": "l",
		"d": []string{"f", "z"},
	}

	jsnExpected := MetaData{
		"a": MetaData{
			"b": MetaData{"c": "e"},
			"l": "m",
			"n": []MetaData{{"o": "p"}, {"q": "r"}},
			"x": []string{"y", "z"},
		},
		"b": "l",
		"d": []string{"f", "z"},
	}
	assert.NoError(t, jsn.SetNested("a.b.c", "e"))
	assert.Equal(t, jsnExpected.String(), jsn.String())

	jsnExpected = MetaData{
		"a": MetaData{
			"b": MetaData{"c": "e"},
			"l": "m",
			"n": []MetaData{{"o": "e"}, {"q": "r"}},
			"x": []string{"y", "z"},
		},
		"b": "l",
		"d": []string{"f", "z"},
	}
	assert.NoError(t, jsn.SetNested("a.n[0].o", "e"))
	assert.Equal(t, jsnExpected.String(), jsn.String())

	jsnExpected = MetaData{
		"a": MetaData{
			"b": MetaData{"c": "e"},
			"l": "m",
			"n": []MetaData{{"o": "e"}, {"q": "r"}},
			"x": []string{"e", "z"},
		},
		"b": "l",
		"d": []string{"f", "z"},
	}
	assert.NoError(t, jsn.SetNested("a.x[0]", "e"))
	assert.Equal(t, jsnExpected.String(), jsn.String())

	jsnExpected = MetaData{
		"a": MetaData{
			"b": MetaData{"c": "e"},
			"l": "m",
			"n": []MetaData{{"o": "e"}, {"q": "r"}},
			"x": []string{"e", "z"},
		},
		"b": "e",
		"d": []string{"f", "z"},
	}
	assert.NoError(t, jsn.SetNested("b", "e"))
	assert.Equal(t, jsnExpected.String(), jsn.String())

	jsnExpected = MetaData{
		"a": MetaData{
			"b": MetaData{"c": "e"},
			"l": "m",
			"n": []MetaData{{"o": "e"}, {"q": "r"}},
			"x": []string{"e", "z"},
		},
		"b": "e",
		"d": []string{"f", "z"},
	}
	assert.NoError(t, jsn.SetNested("d[0]", "f"))
	assert.Equal(t, jsnExpected.String(), jsn.String())

	jsnExpected = MetaData{
		"a": MetaData{
			"b": MetaData{"c": "e"},
			"l": "m",
			"n": []MetaData{{"o": "e"}, {"q": "r"}},
			"x": []string{"e", "z"},
		},
		"b": "e",
		"d": []string{"f", "z"},
		"c": "f",
	}
	assert.NoError(t, jsn.SetNested("c", "f"))
	assert.Equal(t, jsnExpected.String(), jsn.String())

	jsnExpected = MetaData{
		"a": MetaData{
			"b": MetaData{"c": "e"},
			"l": "m",
			"n": []MetaData{{"o": "e"}, {"q": "r"}},
			"x": []string{"e", "z"},
		},
		"b": "e",
		"d": []string{"f", "z"},
		"c": "f",
		"k": []string{"hi", "my", "name"},
	}
	assert.NoError(t, jsn.SetNested("k", []string{"hi", "my", "name"}))
	assert.Equal(t, jsnExpected.String(), jsn.String())

	jsnExpected = MetaData{
		"a": MetaData{
			"b": MetaData{"c": "e"},
			"l": MetaData{"ll": "mm"},
			"n": []MetaData{{"o": "e"}, {"q": "r"}},
			"x": []string{"e", "z"},
		},
		"b": "e",
		"d": []string{"f", "z"},
		"c": "f",
		"k": []string{"hi", "my", "name"},
	}
	assert.NoError(t, jsn.SetNested("a.l", map[string]interface{}{"ll": "mm"}))
	assert.Equal(t, jsnExpected.String(), jsn.String())
}

func TestNewAppError(t *testing.T) {
	actual := NewAppError("A", "B", "C", 0, "E", nil)
	expected := &AppError{
		Code:          "B",
		Message:       "C",
		DetailedError: "E",
		RequestId:     "",
		StatusCode:    0,
		Where:         "A",
		IsOAuth:       false,
		Params:        nil,
		cerr:          nil,
	}
	assert.Equal(t, expected.Message, actual.Message)
	assert.Equal(t, expected.DetailedError, actual.DetailedError)
	assert.Equal(t, expected.StatusCode, actual.StatusCode)

	errString := actual.Error()
	assert.Equal(t, "A: C, E", errString)
}

func TestIsInternalServiceError(t *testing.T) {
	result := IsInternalServiceError(nil)
	assert.False(t, result)

	result = IsInternalServiceError(&AppError{})
	assert.False(t, result)

	result = IsInternalServiceError(&AppError{
		StatusCode: BadInput,
	})
	assert.False(t, result)

	result = IsInternalServiceError(&AppError{
		StatusCode: InternalServiceError,
	})
	assert.True(t, result)
}

func TestIsNoResultFoundError(t *testing.T) {
	result := IsNoResultFoundError(nil)
	assert.False(t, result)

	result = IsNoResultFoundError(&AppError{})
	assert.False(t, result)

	result = IsNoResultFoundError(&AppError{
		StatusCode: BadInput,
	})
	assert.False(t, result)

	result = IsNoResultFoundError(&AppError{
		StatusCode: NoResultFound,
	})
	assert.True(t, result)
}

func TestAppError_AppendParams(t *testing.T) {
	err := &AppError{}
	p := map[string]interface{}{
		"error": "testing",
	}

	err.AppendParams(p)
	assert.Equal(t, len(err.Params), 1)
	assert.Equal(t, err.Params["error"], "testing")
}
