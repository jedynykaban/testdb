package implTestElements

import (
	//"context"
	"reflect"
	"strconv"

	"cloud.google.com/go/civil"

	log "github.com/Sirupsen/logrus"
)

type TestEntity struct {
	TestDateTimeField civil.DateTime
	TestFloatField    float32
	TestIntField      int
	TestStringField   string
}

func TestEntityReflectionFun() {
	log.Info("TestEntityReflectionFun started")

	y := TestEntity{}
	x := &y
	log.Info("type: " + reflect.TypeOf(x).String())
	isPointer := reflect.TypeOf(x).Kind() == reflect.Ptr
	baseIsPointer := reflect.TypeOf(y).Kind() == reflect.Ptr
	log.Info("	* is pointer: " + strconv.FormatBool(isPointer))
	log.Info("	* base is pointer: " + strconv.FormatBool(baseIsPointer))

	log.Info("TestEntityReflectionFun finished")
}
