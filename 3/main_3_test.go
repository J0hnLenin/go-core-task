package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestNewStringMapNotNill() {
	have := NewStringIntMap()
	assert.NotNil(s.T(), have)
}

func (s *ServiceSuite) TestNewStringMapType() {
	want := StringIntMap{}
	have := NewStringIntMap()
	assert.IsType(s.T(), want, have)
}

func (s *ServiceSuite) TestAddOneValue() {
	newMap := NewStringIntMap()
	
	key := "k1"
	value := 100
	newMap.Add(key, value)
	
	want := value
	have := newMap[key]

	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestAddValueTwice() {
	newMap := NewStringIntMap()
	key := "k1"
	value1 := 100
	value2 := 200
	
	newMap.Add(key, value1)
	newMap.Add(key, value2)
	
	want := value2
	have := newMap[key]

	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestAddTwoValue() {
	newMap := NewStringIntMap()
	key1 := "k1"
	key2 := "k2"
	value1 := 100
	value2 := 200
	newMap.Add(key1, value1)
	newMap.Add(key2, value2)
	
	want1 := value1
	want2 := value2
	have1 := newMap[key1]
	have2 := newMap[key2]

	assert.Equal(s.T(), want1, have1)
	assert.Equal(s.T(), want2, have2)
}

func (s *ServiceSuite) TestRemoveExistedKey() {
	newMap := NewStringIntMap()
	key := "k1"
	value := 100
	newMap[key] = value

	newMap.Remove(key)

	have, exists := newMap[key]

	assert.Zero(s.T(), have)
	assert.False(s.T(), exists)
}

func (s *ServiceSuite) TestRemoveNotExistedKey() {
	newMap := NewStringIntMap()
	key := "k1"
	newMap.Remove(key)

	have, exists := newMap[key]

	assert.Zero(s.T(), have)
	assert.False(s.T(), exists)
}

func (s *ServiceSuite) TestCopyEmptyMap() {
	newMap := NewStringIntMap()
	
	want := newMap
	have := newMap.Copy()

	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestCopyMap() {
	newMap := NewStringIntMap()
	key := "k1"
	value := -100

	newMap[key] = value

	want := newMap
	have := newMap.Copy()

	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestCopyChangedMap() {
	newMap := NewStringIntMap()
	key := "k1"
	value1 := -100
	value2 := 100

	newMap[key] = value1

	have := newMap.Copy()
	newMap[key] = value2
	want := newMap
	
	assert.NotEqual(s.T(), want, have)
}

func (s *ServiceSuite) TestExistsTrue() {
	newMap := NewStringIntMap()
	key := "k1"
	value := 100
	newMap[key] = value

	have := newMap.Exists(key)
	
	assert.True(s.T(), have)
}

func (s *ServiceSuite) TestExistsFalse() {
	newMap := NewStringIntMap()
	key := "k1"

	have := newMap.Exists(key)
	
	assert.False(s.T(), have)
}

func (s *ServiceSuite) TestGetNotExists() {
	newMap := NewStringIntMap()
	key := "k1"

	value, exists := newMap.Get(key)
	
	assert.False(s.T(), exists)
	assert.Zero(s.T(), value)
}

func (s *ServiceSuite) TestGetExists() {
	newMap := NewStringIntMap()
	key := "k1"
	value := 333

	newMap[key] = value

	have, exists := newMap.Get(key)
	want := value

	assert.True(s.T(), exists)
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestGetChanged() {
	newMap := NewStringIntMap()
	key := "k1"
	value1 := 333
	value2 := 222

	newMap[key] = value1
	newMap[key] = value2

	have, exists := newMap.Get(key)
	want := value2

	assert.True(s.T(), exists)
	assert.Equal(s.T(), want, have)
}