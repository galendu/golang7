package transform

import (
	"strconv"
	"github.com/dgryski/go-farm"
	"github.com/spaolacci/murmur3"
)

type Transformer interface {
	Hash(string,int) uint64
}

type Murmur struct {}

func (Murmur) Hash(str string, featureId int) uint64 {
	ss := strconv.FormatInt(int64(featureId),10) + ":" +str
	hash := murmur3.New64()
	hash.Write([]byte(ss))
	z := hash.Sum64()
	hash.Reset()
	return z
}

type FarmHash struct {}

func (FarmHash) Hash(str string,featureId int) unit64 {
	return farm.Hase64WithSeed([]byte(str),unit64(featureId))
}
