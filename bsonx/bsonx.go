package bsonx

import (
	"bytes"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func M(key string, value any) bson.M {
	return bson.M{key: value}
}

func E(key string, value any) bson.E {
	return bson.E{Key: key, Value: value}
}

func A(values ...any) bson.A {
	value := make(bson.A, 0, len(values))
	for _, v := range values {
		value = append(value, v)
	}
	return value
}

func D(key string, value any) bson.D {
	return bson.D{bson.E{Key: key, Value: value}}
}

func Id(value any) bson.M {
	return M("_id", value)
}

func ToBsonM(data any) bson.M {
	if data == nil {
		return nil
	}
	if d, ok := data.(bson.M); ok {
		return d
	}

	if d, ok := data.(bson.D); ok {
		return dToM(d)
	}

	if d, ok := data.(map[string]any); ok {
		return MapToBsonM(d)
	}

	if d, ok := data.(*map[string]any); ok && d != nil {
		return MapToBsonM(*d)
	}

	return nil
}

func MapToBsonM(data map[string]any) bson.M {
	m := bson.M{}
	for k, v := range data {
		m[k] = v
	}
	return m
}

func dToM(d bson.D) bson.M {
	marshal, err := bson.Marshal(d)
	if err != nil {
		return nil
	}
	var m bson.M
	decoder := bson.NewDecoder(bson.NewDocumentReader(bytes.NewReader(marshal)))
	decoder.DefaultDocumentM()
	err = decoder.Decode(&m)
	if err != nil {
		return nil
	}
	return m
}
