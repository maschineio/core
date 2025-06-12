package token_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core/token"
)

func TestKV(t *testing.T) {
	kv := token.KV{Key: "testKey", Value: 123}
	assert.Equal(t, "testKey", kv.Key)
	assert.Equal(t, 123, kv.Value)
	assert.Contains(t, kv.String(), "Key(testKey)=Val(int, '123')")
}

func TestReplacementKVBool(t *testing.T) {
	rKey := token.NewReplacementKey("boolKey.$")
	dVal := token.NewDataValue(true, token.Bool)
	rKV := token.ReplacementKV[bool]{
		Key:   rKey,
		Value: dVal,
	}
	assert.Equal(t, "boolKey", rKV.Key.Key())
	assert.Equal(t, "boolKey.$", rKV.Key.OriginKey())
	assert.Equal(t, true, rKV.Value.Value())
	assert.Equal(t, token.Bool, rKV.Value.Type())
}

func TestReplacementKeyWithoutSuffix(t *testing.T) {
	rKey := token.NewReplacementKey("plainKey")
	assert.Equal(t, "plainKey", rKey.Key())
	assert.Equal(t, "plainKey", rKey.OriginKey())
}

func TestReplacementKVFloat(t *testing.T) {
	rKey := token.NewReplacementKey("floatKey.$")
	dVal := token.NewDataValue(3.14, token.Float64) // Korrigiert von DataTypeNumber zu Float64
	rKV := token.ReplacementKV[float64]{
		Key:   rKey,
		Value: dVal,
	}
	assert.Equal(t, "floatKey", rKV.Key.Key())
	assert.Equal(t, "floatKey.$", rKV.Key.OriginKey())
	assert.Equal(t, 3.14, rKV.Value.Value())
	assert.Equal(t, token.Float64, rKV.Value.Type()) // Korrigiert von DataTypeNumber zu Float64
}

func TestKVWithNilValue(t *testing.T) {
	kv := token.KV{Key: "nilKey", Value: nil}
	assert.Equal(t, "nilKey", kv.Key)
	assert.Nil(t, kv.Value)
	assert.Contains(t, kv.String(), "Key(nilKey)=Val(<nil>")
}
