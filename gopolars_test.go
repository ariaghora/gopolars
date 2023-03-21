package gopolars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSV(t *testing.T) {
	_, err := ReadCSV("res/test.csv")
	assert.Nil(t, err)
	_, err = ReadCSV("res/test_nonexistent.csv")
	assert.NotNil(t, err)
}

func TestColumns(t *testing.T) {
	df, err := ReadCSV("res/test.csv")
	assert.Nil(t, err)

	_, err = df.Columns("a")
	assert.Nil(t, err)

	_, err = df.Columns("b")
	assert.Nil(t, err)

	_, err = df.Columns("z")
	assert.NotNil(t, err)
}

func TestFrameEquals(t *testing.T) {
	df1, err := ReadCSV("res/test_with_missing.csv")
	assert.Nil(t, err)

	df2, err := ReadCSV("res/test_with_missing.csv")
	assert.Nil(t, err)

	assert.True(t, df1.FrameEqual(df2, true))
	assert.False(t, df1.FrameEqual(df2, false))
}

func TestCollect(t *testing.T) {
	df, err := ReadCSV("res/test.csv")
	assert.Nil(t, err)

	dlf, err := df.Lazy().Collect()
	assert.Nil(t, err)

	assert.True(t, df.FrameEqual(dlf, true))
}

func TestNewSeries(t *testing.T) {
	s := NewIntSeries("a", []int{1, 2, 3})
	assert.NotNil(t, s.data)
}

func TestSelect(t *testing.T) {
	df, err := ReadCSV("res/test.csv")
	assert.Nil(t, err)

	df, err = df.Lazy().Select(
		Col("a"),
		Col("a").Eq(Int(2)).Alias("is two?"),
	).Collect()

	assert.Nil(t, err)
}
