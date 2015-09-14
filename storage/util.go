package storage

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	typeofstring            = reflect.TypeOf((*string)(nil)).Elem()
	typeoftime              = reflect.TypeOf((*time.Time)(nil)).Elem()
	typeofint64             = reflect.TypeOf((*int64)(nil)).Elem()
	typeofBlobPropertiesPtr = reflect.TypeOf((*BlobProperties)(nil))
	typeofBlobProperties    = typeofBlobPropertiesPtr.Elem()
	typeofMapStringString   = reflect.TypeOf((*map[string]string)(nil)).Elem()
	typeToPropertyHeaderMap = make(map[reflect.Type]propertyHeaderMaps)
)

type propertyHeaderMaps struct {
	HeaderToValueIndex                     map[string]int
	ValueIndexToHeaders                    map[int][]string
	ExtraHeadersValueIndex                 int
	PrefixedExtraHeadersPrefixToValueIndex map[string]int
	PrefixedExtraHeadersValueIndexToPrefix map[int]string
}

func init() {
	initHeaderPropertiesMapsForType(typeofBlobPropertiesPtr)
}

func (c Client) computeHmac256(message string) string {
	h := hmac.New(sha256.New, c.accountKey)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func currentTimeRfc1123Formatted() string {
	return timeRfc1123Formatted(time.Now().UTC())
}

func timeRfc1123Formatted(t time.Time) string {
	return t.Format(http.TimeFormat)
}

func mergeParams(v1, v2 url.Values) url.Values {
	out := url.Values{}
	for k, v := range v1 {
		out[k] = v
	}
	for k, v := range v2 {
		vals, ok := out[k]
		if ok {
			vals = append(vals, v...)
			out[k] = vals
		} else {
			out[k] = v
		}
	}
	return out
}

func prepareBlockListRequest(blocks []Block) string {
	s := `<?xml version="1.0" encoding="utf-8"?><BlockList>`
	for _, v := range blocks {
		s += fmt.Sprintf("<%s>%s</%s>", v.Status, v.ID, v.Status)
	}
	s += `</BlockList>`
	return s
}

func xmlUnmarshal(body io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(body)
	log.Println(string(data))
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, v)
}

func xmlMarshal(v interface{}) (io.Reader, int, error) {
	b, err := xml.Marshal(v)
	if err != nil {
		return nil, 0, err
	}
	return bytes.NewReader(b), len(b), nil
}

func setUntypedProperties(propsValue reflect.Value, vix int, header, headerValue string) {
	v := propsValue.Field(vix)
	var m map[string]string
	if v.IsNil() {
		m = make(map[string]string)
		v.Set(reflect.ValueOf(m))
	} else {
		m = v.Interface().(map[string]string)
	}
	m[header] = headerValue
}

func unmarshalProperties(headers map[string]string, props interface{}) error {
	propsType := reflect.TypeOf(props)
	if propsType.Kind() != reflect.Ptr || propsType.Elem().Kind() != reflect.Struct {
		return errors.New("props must be a pointer to a struct")
	}

	propsType = propsType.Elem()
	propertiesMaps, ok := typeToPropertyHeaderMap[propsType]
	if !ok {
		return fmt.Errorf("Type '%s' is not registered, did you forget calling initHeadersPropertyMapsForType(refelct.TypeOf(%s)) ?",
			propsType.Name(), propsType.Name())
	}

	bvalue := reflect.ValueOf(props).Elem()
	for header, headerValue := range headers {
		header = strings.ToLower(header)
		vix, ok := propertiesMaps.HeaderToValueIndex[header]
		if !ok {
			foundPrefixed := false
			for prefix, vix := range propertiesMaps.PrefixedExtraHeadersPrefixToValueIndex {
				if len(header) > len(prefix) && strings.HasPrefix(header, prefix) {
					foundPrefixed = true
					setUntypedProperties(bvalue, vix, strings.TrimPrefix(header, prefix), headerValue)
					break
				}
			}
			if !foundPrefixed {
				if propertiesMaps.ExtraHeadersValueIndex > -1 {
					setUntypedProperties(bvalue, propertiesMaps.ExtraHeadersValueIndex, header, headerValue)
				}
			}
			continue
		}
		v := bvalue.Field(vix)
		t := propsType.Field(vix).Type
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
			deref := reflect.New(t)
			v.Set(deref)
			v = deref.Elem()
		}
		if t.ConvertibleTo(typeofint64) {
			ival, _ := strconv.ParseInt(headerValue, 10, 64)
			v.SetInt(ival)
		} else if t.ConvertibleTo(typeofstring) {
			v.SetString(headerValue)
		}
	}

	return nil
}

func marshalProperties(props interface{}, headers map[string]string) error {
	propsType := reflect.TypeOf(props)
	if propsType.Kind() != reflect.Ptr || propsType.Elem().Kind() != reflect.Struct {
		return errors.New("props must be a pointer to a struct")
	}

	propsType = propsType.Elem()

	propertiesMaps, ok := typeToPropertyHeaderMap[propsType]
	if !ok {
		return fmt.Errorf("Type '%s' is not registered, did you forget calling initHeadersPropertyMapsForType(refelct.TypeOf(%s)) ?",
			propsType.Name(), propsType.Name())
	}

	bvalue := reflect.ValueOf(props).Elem()
	for fix, headerNames := range propertiesMaps.ValueIndexToHeaders {
		v := bvalue.Field(fix)
		var val string
		t := v.Type()
		if t.Kind() == reflect.Ptr {
			if v.IsNil() {
				continue
			} else {
				t = t.Elem()
				v = v.Elem()
			}
		}
		if t.ConvertibleTo(typeofint64) {
			ival := v.Convert(typeofint64).Int()
			val = strconv.FormatInt(ival, 10)
		} else if t.ConvertibleTo(typeofstring) {
			val = v.Convert(typeofstring).String()
		} else {
			return fmt.Errorf("Property field '%s' of type '%s' cannot be marshaled to headers. Only 'string' and 'int64' are currently supported",
				propsType.Field(fix).Name, t.Name())
		}
		if val == "" {
			continue
		}
		for _, headerName := range headerNames {
			headers[headerName] = val
		}
	}

	for fix, prefix := range propertiesMaps.PrefixedExtraHeadersValueIndexToPrefix {
		v := bvalue.Field(fix).Interface().(map[string]string)
		for k, v := range v {
			headers[prefix+k] = v
		}
	}

	if propertiesMaps.ExtraHeadersValueIndex > -1 {
		v := bvalue.Field(propertiesMaps.ExtraHeadersValueIndex).Interface().(map[string]string)
		for k, v := range v {
			headers[k] = v
		}
	}

	return nil
}

func initHeaderPropertiesMapsForType(t reflect.Type) {
	t = t.Elem()
	var propertiesMaps propertyHeaderMaps
	propertiesMaps.HeaderToValueIndex = make(map[string]int)
	propertiesMaps.ValueIndexToHeaders = make(map[int][]string)
	propertiesMaps.ExtraHeadersValueIndex = -1
	propertiesMaps.PrefixedExtraHeadersPrefixToValueIndex = make(map[string]int)
	propertiesMaps.PrefixedExtraHeadersValueIndexToPrefix = make(map[int]string)

	for fix := 0; fix < t.NumField(); fix++ {
		f := t.Field(fix)
		headerTag := f.Tag.Get("header")
		if headerTag == "" {
			continue
		}
		if headerTag[0] == '@' {
			if f.Type != typeofMapStringString {
				panic(fmt.Sprintf("Map properties must be of type '%s', while '%s' is '%s'",
					typeofMapStringString, f.Name, f.Type))
			}
			if len(headerTag) > 1 {
				prefix := headerTag[1:]
				prefixLower := strings.ToLower(prefix)
				if _, alreadyThere := propertiesMaps.PrefixedExtraHeadersPrefixToValueIndex[prefixLower]; alreadyThere {
					panic(fmt.Sprintf("Only a single map property can be mapped to a given pefix '%s'", prefix))
				}
				propertiesMaps.PrefixedExtraHeadersPrefixToValueIndex[prefixLower] = fix
				propertiesMaps.PrefixedExtraHeadersValueIndexToPrefix[fix] = prefix
			} else {
				if propertiesMaps.ExtraHeadersValueIndex > -1 {
					panic("Only a single non prefixed map property can be used")
				}

				propertiesMaps.ExtraHeadersValueIndex = fix
			}
			continue
		}
		headers := strings.Split(headerTag, ",")
		trimmedHeaders := make([]string, len(headers), len(headers))
		for hix := 0; hix < len(headers); hix++ {
			trimmed := strings.TrimSpace(headers[hix])
			trimmedHeaders[hix] = trimmed
			propertiesMaps.HeaderToValueIndex[strings.ToLower(trimmed)] = fix
		}
		propertiesMaps.ValueIndexToHeaders[fix] = headers
	}

	typeToPropertyHeaderMap[t] = propertiesMaps
}
