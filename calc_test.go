package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testStructure struct {
	x        int
	y        int
	expected int
}

type testDivide struct {
	x        int
	y        int
	expected float64
}

type AddTestSuite struct {
	suite.Suite
}

type SubtractTestSuite struct {
	suite.Suite
}

type MultiplyTestSuite struct {
	suite.Suite
}

type DivideTestSuite struct {
	suite.Suite
}

func (suite *AddTestSuite) TestAdd() {
	testData := []testStructure{
		{5, 7, 12},
		{8, 99, 107},
		{33, 66, 99},
		{109, -200, -91},
		{-5, -4, -9},
	}
	for _, v := range testData {
		assert.Equal(suite.T(), Add(v.x, v.y), v.expected, "Adding %d and %d will give %d", v.x, v.y, v.expected)
	}
}

func (suite *SubtractTestSuite) TestSubtract() {
	testData := []testStructure{
		{12, 4, 8},
		{50, 50, 0},
		{30, 45, -15},
		{-20, -22, 2},
		{-23, 50, -73},
	}
	for _, v := range testData {
		assert.Equal(suite.T(), Subtract(v.x, v.y), v.expected, "%d subtract %d will give %d", v.x, v.y, v.expected)
	}
}
func (suite *MultiplyTestSuite) TestMultiply() {
	testData := []testStructure{
		{2, 3, 6},
		{4, 7, 28},
		{-5, 3, -15},
		{8, 0, 0},
	}
	for _, v := range testData {
		assert.Equal(suite.T(), Multiply(v.x, v.y), v.expected, "%d multiply by %d will give %d", v.x, v.y, v.expected)
	}
}

func (suite *DivideTestSuite) TestDivide() {
	testData := []testDivide{
		{5, -2, -2.5},
		{6, 8, 0.75},
		{-30, -10, 3.0},
		{100, 25, 4.0},
		{100, 0, math.Inf(1)},
	}
	for _, v := range testData {
		assert.Equal(suite.T(), Divide(v.x, v.y), v.expected, "%d divided by %d will give %d", v.x, v.y, v.expected)
	}
}
func TestSuites(t *testing.T) {
	suite.Run(t, new(AddTestSuite))
	suite.Run(t, new(SubtractTestSuite))
	suite.Run(t, new(MultiplyTestSuite))
	suite.Run(t, new(DivideTestSuite))
}
