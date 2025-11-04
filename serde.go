package serde

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

func EncodeJson[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return SerializingError(err)
	}
	return nil
}

func EncodeXml[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.WriteHeader(status)
	if err := xml.NewEncoder(w).Encode(v); err != nil {
		return SerializingError(err)
	}
	return nil
}

// deseriliazing a request body to a struct. In the case of a server we don'tneed to call r.Close(), its done automaically by the server
func Decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, DeserializingError(err)
	}
	return v, nil
}

// second version of deseriaizing that is more general. Use this one.
func DecodeV2[T any](r io.ReadCloser) (T, error) {
	defer r.Close()
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, DeserializingError(err)
	}
	return v, nil
}

func DecodeJsonFileToStruct[T any](path string) (T, error) {
	var result T

	file, err := os.Open(path)
	if err != nil {
		return result, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	// to  check that file implements the io.Reader interface
	// var _ io.Reader = (*os.File)(nil)

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&result); err != nil {
		return result, DeserializingError(err)
	}

	return result, nil
}

// pass an os.Open(path) object
func DecodeJsonFileToStructV2[T any](r io.ReadCloser) (T, error) {
	var result T

	defer r.Close()

	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&result); err != nil {
		return result, DeserializingError(err)
	}

	return result, nil
}

func WriteStructToFileAsJson[T any](p string, v T) error {
	// var f os.File
	file, err := os.Create(p)

	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer file.Close()

	if err = json.NewEncoder(file).Encode(v); err != nil {
		return SerializingError(err)
	}

	return nil
}

func WriteStructToFileAsJsonV2[T any](f io.WriteCloser, v T) error {

	defer f.Close()

	if err := json.NewEncoder(f).Encode(v); err != nil {
		return SerializingError(err)
	}

	return nil
}

func SerializingError(e error) error {
	return fmt.Errorf("at serializing(encode, marshall) %w", e)
}

func DeserializingError(e error) error {
	return fmt.Errorf("at deserializing(decode, unmarshal) %w", e)
}

func DecodeXml[T any](r io.ReadCloser) (T, error) {
	var v T
	defer r.Close()
	if err := xml.NewDecoder(r).Decode(&v); err != nil {
		return v, DeserializingError(err)
	}
	return v, nil
}
