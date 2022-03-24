package massbank

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"time"
)

func (p StringProperty) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(p.string)
}

func (p SubtagProperty) MarshalBSONValue() (bsontype.Type, []byte, error) {
	m := map[string]string{p.subtag: p.string}
	return bson.MarshalValue(m)
}

func (p DatabaseProperty) MarshalBSONValue() (bsontype.Type, []byte, error) {
	m := map[string]string{p.Database: p.Identifier}
	return bson.MarshalValue(m)
}

func (p RecordDeprecated) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(struct {
		Date   time.Time
		reason string
	}{p.Date, p.Reason})
}

func (p RecordDate) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(struct {
		updated time.Time
		created time.Time
	}{p.Updated, p.Created})
}

func (p RecordAuthorNames) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(p.value)
}

func (p ChCompoundClasses) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(p.value)
}

func (p ChMass) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(p.value)
}

func (p SpLineage) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(p.value)
}

func (p PkPeak) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(struct {
		Header []string
		Values []PeakValue
	}{p.Header, p.Values})
}

func (p PeakValue) MarshalBSONValue() (bsontype.Type, []byte, error) {
	val := bson.A{p.Mz, p.Intensity, p.Rel}
	return bson.MarshalValue(val)
}

func (p PkNumPeak) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(p.Value)
}

func (p PkAnnotation) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(struct {
		Header []string
		Values []AnnotationValue
	}{p.Header, p.Values})
}
