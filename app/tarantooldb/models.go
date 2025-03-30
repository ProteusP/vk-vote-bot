package tarantooldb

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/vmihailenco/msgpack/v5"
)

type Vote struct {
	ID       string            `msgpack:"id"`
	Creator  string            `msgpack:"creator"`
	Question string            `msgpack:"question"`
	Options  map[string]uint64 `msgpack:"options,omitempty"`
	Votes    map[string]string `msgpack:"votes,omitempty"`
	Status   string            `msgpack:"status"`
}

type ConvertationError struct {
	Message string
}

func NewConvertationError() ConvertationError {
	return ConvertationError{
		Message: "Ошибка преобразования ответа к нужному формату",
	}
}

func (e *ConvertationError) Error() string {
	return e.Message
}

func (v *Vote) Results() string {
	var result strings.Builder
	result.WriteString("📊 Результаты:\n")

	for option := range v.Options {
		count := v.Options[option]

		result.WriteString(fmt.Sprintf("- %s: %d\n", option, count))
	}
	return result.String()
}

func (v *Vote) LoadFromResponse(responseData []interface{}) error {

	data := responseData[0]
	err := NewConvertationError()
	tuple, ok := data.([]interface{})
	if !ok {
		log.Printf("[DEBUG] Проблема в tuple: %v", data)
		return &err
	}

	id, ok := tuple[0].(string)
	if !ok {
		log.Printf("[DEBUG] Проблема в id: %v", tuple[0])
		return &err
	}

	v.ID = id

	creator, ok := tuple[1].(string)
	if !ok {
		log.Printf("[DEBUG] Проблема в creator: %v", tuple[1])
		return &err
	}

	v.Creator = creator

	question, ok := tuple[2].(string)
	if !ok {
		log.Printf("[DEBUG] Проблема в question: %v", tuple[2])
		return &err
	}

	v.Question = question

	rawOptions, ok := tuple[3].(map[interface{}]interface{})
	if !ok {
		log.Printf("[DEBUG] Проблема в rawOptions: %v", tuple[3])
		return &err
	}

	options := make(map[string]uint64)
	for key, val := range rawOptions {
		strKey, ok := key.(string)
		if !ok {
			log.Printf("[DEBUG] Проблема в options key: %v ; %v", key, rawOptions)
			return &err
		}
		log.Printf("%v, %v", key, val)

		intVal, ok := val.(uint64)
		if !ok {
			log.Printf("[DEBUG] Проблема в options val: %v ; %v", val, rawOptions)
			return &err
		}

		options[strKey] = intVal
	}

	v.Options = options

	rawVotes, ok := tuple[4].(map[interface{}]interface{})
	if !ok {
		log.Printf("[DEBUG] Проблема в rawVotes: %v", tuple[4])
		return &err
	}

	votes := make(map[string]string)
	for key, val := range rawVotes {
		strKey, ok := key.(string)
		if !ok {
			log.Printf("[DEBUG] Проблема в votes key: %v", key)
			return &err
		}
		strVal, ok := val.(string)
		if !ok {
			log.Printf("[DEBUG] Проблема в votes val: %v", val)
			return &err
		}
		votes[strKey] = strVal
	}

	v.Votes = votes

	status, ok := tuple[5].(string)
	if !ok {
		log.Printf("[DEBUG] Проблема в status: %v", tuple[5])
		return &err
	}

	v.Status = status

	return nil
}
