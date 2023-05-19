package gocache

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LFU_Set(t *testing.T) {
	lfu := NewLFU(1024)

	lfu.Set("key", "value")
	value, exist := lfu.Get("key")
	require.Equal(t, exist, true)
	require.Equal(t, "value", value)

	lfu.Set("key1", "value_two")
	value, exist = lfu.Get("key1")
	require.Equal(t, exist, true)
	require.Equal(t, "value_two", value)

	setTestData := make(map[string]string)

	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		setTestData[key] = value
		lfu.Set(key, value)
	}

	for key, value := range setTestData {
		v, exist := lfu.Get(key)
		require.Equal(t, exist, true)
		require.Equal(t, value, v)
	}
}

func Test_LFU_Get(t *testing.T) {
	lfu := NewLFU(1024)

	lfu.Set("key", "value")
	value, exist := lfu.Get("key")
	require.Equal(t, exist, true)
	require.Equal(t, "value", value)

	lfu.Set("key1", "value_two")
	value, exist = lfu.Get("key1")
	require.Equal(t, exist, true)
	require.Equal(t, "value_two", value)

	setTestData := make(map[string]string)

	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		setTestData[key] = value
		lfu.Set(key, value)
	}

	for key, value := range setTestData {
		v, exist := lfu.Get(key)
		require.Equal(t, exist, true)
		require.Equal(t, value, v)
	}

	_, exist = lfu.Get("not_exist_key")
	require.Equal(t, exist, false)

}

func Test_LFU_Contains(t *testing.T) {
	lfu := NewLFU(1024)

	lfu.Set("key", "value")
	exist := lfu.Contains("key")
	require.Equal(t, exist, true)

	lfu.Set("key1", "value_two")
	exist = lfu.Contains("key1")
	require.Equal(t, exist, true)

	setTestData := make(map[string]string)

	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		setTestData[key] = value
		lfu.Set(key, value)
	}

	for key := range setTestData {
		exist := lfu.Contains(key)
		require.Equal(t, exist, true)
	}

	exist = lfu.Contains("not_exist_key")
	require.Equal(t, exist, false)
}

func Test_LFU_Evite(t *testing.T) {
	lfu := NewLFU(3)

	lfu.Set("key", "value")
	value, exist := lfu.Get("key")
	require.Equal(t, exist, true)
	require.Equal(t, "value", value)

	lfu.Set("key1", "value_two")
	value, exist = lfu.Get("key1")
	require.Equal(t, exist, true)
	require.Equal(t, "value_two", value)

	lfu.Set("key3", "value_three")
	value, exist = lfu.Get("key3")
	require.Equal(t, exist, true)
	require.Equal(t, "value_three", value)

	lfu.Get("key")
	lfu.Get("key")
	lfu.Get("key1")

	lfu.Set("key4", "value")
	exist = lfu.Contains("key3")
	require.Equal(t, exist, false)

	lfu.Set("key5", "value")
	exist = lfu.Contains("key4")
	require.Equal(t, exist, false)
}

func Test_LFU_Delete(t *testing.T) {

}
