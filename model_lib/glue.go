package model_lib

import (
	"code.google.com/p/tcgl/redis"
	"fmt"
	"github.com/dchest/uniuri"
	"reflect"
	"strconv"
	"time"
)

// File contains glue code that connects the Model
// abstraction to the database.

// writes the interface to the redis database
// throws an error if the type has not yet been
// registered
func Save(m ModelInterface) (ModelInterface, error) {
	fmt.Println("models.Save() was called")

	// get the type of the model we're trying to save
	// make sure it is registered
	typ := reflect.TypeOf(m)
	name, ok := typeToName[typ]
	if !ok {
		return nil, NewModelTypeNotRegisteredError(typ)
	}

	// prepare the arguments for redis driver
	id := generateRandomId()
	key := name + ":" + id
	args := toArgSlice(key, m)

	// invoke redis driver to commit to database
	result := db.Command("hmset", args...)
	if result.Error() != nil {
		return nil, result.Error()
	}

	m.SetId(id)

	return m, nil
}

// TODO: remove the record from the database
func Delete(ModelInterface) {
	fmt.Println("models.Delete() was called")
}

func FindById(modelName, id string) (ModelInterface, error) {

	typ, ok := nameToType[modelName]
	if !ok {
		return nil, NewModelNameNotRegisteredError(modelName)
	}

	key := modelName + ":" + id
	result := db.Command("hgetall", key)
	if result.Error() != nil {
		return nil, result.Error()
	}

	keyValues := result.KeyValues()
	resultMap := convertKeyValuesToMap(keyValues)
	model, err := fromMap(resultMap, typ)
	if err != nil {
		return nil, err
	}

	model.SetId(id)
	return model, nil
}

// Converts and returns slice of redis.KeyValues into a map
func convertKeyValuesToMap(slice []*redis.KeyValue) map[string]string {
	m := make(map[string]string)
	for _, elem := range slice {
		key := elem.Key
		val := elem.Value.String()
		m[key] = val
	}
	return m
}

// Uses reflect to dynamically convert a map of
// [string]string to a ModelInterface (a struct)
func fromMap(m map[string]string, typ reflect.Type) (ModelInterface, error) {
	fmt.Println("map: ", m)

	val := reflect.New(typ).Elem()
	numFields := val.NumField()

	fmt.Println("val: ", val)
	fmt.Println("numFields: ", numFields)

	for i := 0; i < numFields; i++ {
		field := typ.Field(i)
		mapVal, ok := m[field.Name]
		fmt.Println("fieldName: ", field.Name)
		fmt.Println("mapVal: ", mapVal)
		fmt.Println("can set? ", val.Field(i).CanSet())
		if ok {
			fmt.Println("Twas okay")
			fmt.Println("Kind: ", val.Field(i).Kind())
			switch val.Field(i).Kind() {
			case reflect.String:
				val.Field(i).SetString(mapVal)
				fmt.Println("Field ", i, val.String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valInt, err := strconv.Atoi(mapVal)
				if err != nil {
					return nil, err
				}
				val.Field(i).SetInt(int64(valInt))
			}
			// TODO: add more cases
		}
	}

	// Set and allocate the embedded *Model attribute
	// This is so we can call SetId() later
	val.FieldByName("Model").Set(reflect.ValueOf(new(Model)))

	// Typecast and return the result
	model := val.Interface().(ModelInterface)
	fmt.Println("model: ", model)
	return model, nil
}

// converts an interface and a given name to a slice of interface{}
// the slice can then be passed directly to the redis driver
func toArgSlice(key string, in ModelInterface) []interface{} {

	// get the number of fields
	sVal := reflect.ValueOf(in)
	sType := reflect.TypeOf(in)
	numFields := sVal.NumField()

	// init/allocate a slice of arguments
	args := make([]interface{}, 0, numFields*2+1)

	// the first arg is the key for the set
	args = append(args, key)

	// the remaining args are members of the redis set and their values.
	// iterate through fields and add each one to the slice
	for i := 0; i < numFields; i++ {
		fieldName := sType.Field(i).Name
		if fieldName == "Model" {
			continue
		}
		args = append(args, fieldName)
		value := sVal.Field(i)
		switch value.Kind() {
		case reflect.String:
			args = append(args, value.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valueString := strconv.Itoa(int(value.Int()))
			args = append(args, valueString)
		default:
			// fmt.Sprint method as a catch-all
			// it's a little bit slower
			valueString := fmt.Sprint(value.Interface())
			args = append(args, valueString)
		}
	}

	return args
}

// generates a random string that is more or less
// garunteed to be unique
func generateRandomId() string {
	timeInt := time.Now().Unix()
	timeString := strconv.FormatInt(timeInt, 36)
	randomString := uniuri.NewLen(16)
	return randomString + timeString
}
