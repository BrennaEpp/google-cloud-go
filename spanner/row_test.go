/*
Copyright 2017 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spanner

import (
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"cloud.google.com/go/internal/testutil"
	sppb "cloud.google.com/go/spanner/apiv1/spannerpb"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	proto "google.golang.org/protobuf/proto"
	proto3 "google.golang.org/protobuf/types/known/structpb"
)

var (
	tm    = time.Date(2016, 11, 15, 0, 0, 0, 0, time.UTC)
	dt, _ = civil.ParseDate("2016-11-15")
	// row contains a column for each unique Cloud Spanner type.
	row = Row{
		[]*sppb.StructType_Field{
			// STRING / STRING ARRAY
			{Name: "STRING", Type: stringType()},
			{Name: "NULL_STRING", Type: stringType()},
			{Name: "STRING_ARRAY", Type: listType(stringType())},
			{Name: "NULL_STRING_ARRAY", Type: listType(stringType())},
			// BYTES / BYTES ARRAY
			{Name: "BYTES", Type: bytesType()},
			{Name: "NULL_BYTES", Type: bytesType()},
			{Name: "BYTES_ARRAY", Type: listType(bytesType())},
			{Name: "NULL_BYTES_ARRAY", Type: listType(bytesType())},
			// INT64 / INT64 ARRAY
			{Name: "INT64", Type: intType()},
			{Name: "NULL_INT64", Type: intType()},
			{Name: "INT64_ARRAY", Type: listType(intType())},
			{Name: "NULL_INT64_ARRAY", Type: listType(intType())},
			// BOOL / BOOL ARRAY
			{Name: "BOOL", Type: boolType()},
			{Name: "NULL_BOOL", Type: boolType()},
			{Name: "BOOL_ARRAY", Type: listType(boolType())},
			{Name: "NULL_BOOL_ARRAY", Type: listType(boolType())},
			// FLOAT64 / FLOAT64 ARRAY
			{Name: "FLOAT64", Type: floatType()},
			{Name: "NULL_FLOAT64", Type: floatType()},
			{Name: "FLOAT64_ARRAY", Type: listType(floatType())},
			{Name: "NULL_FLOAT64_ARRAY", Type: listType(floatType())},
			// FLOAT32 / FLOAT32 ARRAY
			{Name: "FLOAT32", Type: float32Type()},
			{Name: "NULL_FLOAT32", Type: float32Type()},
			{Name: "FLOAT32_ARRAY", Type: listType(float32Type())},
			{Name: "NULL_FLOAT32_ARRAY", Type: listType(float32Type())},
			// TIMESTAMP / TIMESTAMP ARRAY
			{Name: "TIMESTAMP", Type: timeType()},
			{Name: "NULL_TIMESTAMP", Type: timeType()},
			{Name: "TIMESTAMP_ARRAY", Type: listType(timeType())},
			{Name: "NULL_TIMESTAMP_ARRAY", Type: listType(timeType())},
			// DATE / DATE ARRAY
			{Name: "DATE", Type: dateType()},
			{Name: "NULL_DATE", Type: dateType()},
			{Name: "DATE_ARRAY", Type: listType(dateType())},
			{Name: "NULL_DATE_ARRAY", Type: listType(dateType())},
			// UUID / UUID ARRAY
			{Name: "UUID", Type: uuidType()},
			{Name: "NULL_UUID", Type: uuidType()},
			{Name: "UUID_ARRAY", Type: listType(uuidType())},
			{Name: "NULL_UUID_ARRAY", Type: listType(uuidType())},

			// STRUCT ARRAY
			{
				Name: "STRUCT_ARRAY",
				Type: listType(
					structType(
						mkField("Col1", intType()),
						mkField("Col2", floatType()),
						mkField("Col3", float32Type()),
						mkField("Col4", stringType()),
						mkField("Col5", uuidType()),
					),
				),
			},
			{
				Name: "NULL_STRUCT_ARRAY",
				Type: listType(
					structType(
						mkField("Col1", intType()),
						mkField("Col2", floatType()),
						mkField("Col3", float32Type()),
						mkField("Col4", stringType()),
						mkField("Col5", uuidType()),
					),
				),
			},
		},
		[]*proto3.Value{
			// STRING / STRING ARRAY
			stringProto("value"),
			nullProto(),
			listProto(stringProto("value1"), nullProto(), stringProto("value3")),
			nullProto(),
			// BYTES / BYTES ARRAY
			bytesProto([]byte("value")),
			nullProto(),
			listProto(bytesProto([]byte("value1")), nullProto(), bytesProto([]byte("value3"))),
			nullProto(),
			// INT64 / INT64 ARRAY
			intProto(17),
			nullProto(),
			listProto(intProto(1), intProto(2), nullProto()),
			nullProto(),
			// BOOL / BOOL ARRAY
			boolProto(true),
			nullProto(),
			listProto(nullProto(), boolProto(true), boolProto(false)),
			nullProto(),
			// FLOAT64 / FLOAT64 ARRAY
			floatProto(1.7),
			nullProto(),
			listProto(nullProto(), nullProto(), floatProto(1.7)),
			nullProto(),
			// FLOAT32 / FLOAT32 ARRAY
			float32Proto(0.3),
			nullProto(),
			listProto(nullProto(), nullProto(), float32Proto(0.3)),
			nullProto(),
			// TIMESTAMP / TIMESTAMP ARRAY
			timeProto(tm),
			nullProto(),
			listProto(nullProto(), timeProto(tm)),
			nullProto(),
			// DATE / DATE ARRAY
			dateProto(dt),
			nullProto(),
			listProto(nullProto(), dateProto(dt)),
			nullProto(),
			// UUID / UUID ARRAY
			uuidProto(uuid1),
			nullProto(),
			listProto(nullProto(), uuidProto(uuid1)),
			nullProto(),
			// STRUCT ARRAY
			listProto(
				nullProto(),
				listProto(intProto(3), floatProto(33.3), float32Proto(0.3), stringProto("three"), uuidProto(uuid1)),
				nullProto(),
			),
			nullProto(),
		},
	}
)

// Test helpers for getting column values.
func TestColumnValues(t *testing.T) {
	vals := []interface{}{}
	wantVals := []interface{}{}
	// Test getting column values.
	for i, wants := range [][]interface{}{
		// STRING / STRING ARRAY
		{"value", NullString{"value", true}},
		{NullString{}},
		{[]NullString{{"value1", true}, {}, {"value3", true}}},
		{[]NullString(nil)},
		// BYTES / BYTES ARRAY
		{[]byte("value")},
		{[]byte(nil)},
		{[][]byte{[]byte("value1"), nil, []byte("value3")}},
		{[][]byte(nil)},
		// INT64 / INT64 ARRAY
		{int64(17), NullInt64{17, true}},
		{NullInt64{}},
		{[]NullInt64{{1, true}, {2, true}, {}}},
		{[]NullInt64(nil)},
		// BOOL / BOOL ARRAY
		{true, NullBool{true, true}},
		{NullBool{}},
		{[]NullBool{{}, {true, true}, {false, true}}},
		{[]NullBool(nil)},
		// FLOAT64 / FLOAT64 ARRAY
		{1.7, NullFloat64{1.7, true}},
		{NullFloat64{}},
		{[]NullFloat64{{}, {}, {1.7, true}}},
		{[]NullFloat64(nil)},
		// FLOAT32 / FLOAT64 ARRAY
		{float32(0.3), NullFloat32{0.3, true}},
		{NullFloat32{}},
		{[]NullFloat32{{}, {}, {float32(0.3), true}}},
		{[]NullFloat32(nil)},
		// TIMESTAMP / TIMESTAMP ARRAY
		{tm, NullTime{tm, true}},
		{NullTime{}},
		{[]NullTime{{}, {tm, true}}},
		{[]NullTime(nil)},
		// DATE / DATE ARRAY
		{dt, NullDate{dt, true}},
		{NullDate{}},
		{[]NullDate{{}, {dt, true}}},
		{[]NullDate(nil)},
		// UUID / UUID ARRAY
		{uuid1, NullUUID{uuid1, true}},
		{NullUUID{}},
		{[]NullUUID{{}, {uuid1, true}}},
		{[]NullUUID(nil)},
		// STRUCT ARRAY
		{
			[]*struct {
				Col1 NullInt64
				Col2 NullFloat64
				Col3 NullFloat32
				Col4 string
				Col5 NullUUID
			}{
				nil,

				{
					NullInt64{3, true},
					NullFloat64{33.3, true},
					NullFloat32{0.3, true},
					"three",
					NullUUID{uuid1, true},
				},
				nil,
			},
			[]NullRow{
				{},
				{
					Row: Row{
						fields: []*sppb.StructType_Field{
							mkField("Col1", intType()),
							mkField("Col2", floatType()),
							mkField("Col3", float32Type()),
							mkField("Col4", stringType()),
							mkField("Col5", uuidType()),
						},
						vals: []*proto3.Value{
							intProto(3),
							floatProto(33.3),
							float32Proto(0.3),
							stringProto("three"),
							uuidProto(uuid1),
						},
					},
					Valid: true,
				},
				{},
			},
		},
		{
			[]*struct {
				Col1 NullInt64
				Col2 NullFloat64
				Col3 NullFloat32
				Col4 string
				Col5 NullUUID
			}(nil),
			[]NullRow(nil),
		},
	} {
		for j, want := range wants {
			// Prepare Value vector to test Row.Columns.
			if j == 0 {
				vals = append(vals, reflect.New(reflect.TypeOf(want)).Interface())
				wantVals = append(wantVals, want)
			}
			// Column
			gotp := reflect.New(reflect.TypeOf(want))
			err := row.Column(i, gotp.Interface())
			if err != nil {
				t.Errorf("\t row.Column(%v, %T) returns error: %v, want nil", i, gotp.Interface(), err)
			}
			if got := reflect.Indirect(gotp).Interface(); !testEqual(got, want) {
				t.Errorf("\t row.Column(%v, %T) retrives %v, want %v", i, gotp.Interface(), got, want)
			}
			// ColumnByName
			gotp = reflect.New(reflect.TypeOf(want))
			err = row.ColumnByName(row.fields[i].Name, gotp.Interface())
			if err != nil {
				t.Errorf("\t row.ColumnByName(%v, %T) returns error: %v, want nil", row.fields[i].Name, gotp.Interface(), err)
			}
			if got := reflect.Indirect(gotp).Interface(); !testEqual(got, want) {
				t.Errorf("\t row.ColumnByName(%v, %T) retrives %v, want %v", row.fields[i].Name, gotp.Interface(), got, want)
			}
		}
	}
	// Test Row.Columns.
	if err := row.Columns(vals...); err != nil {
		t.Errorf("row.Columns() returns error: %v, want nil", err)
	}
	for i, want := range wantVals {
		if got := reflect.Indirect(reflect.ValueOf(vals[i])).Interface(); !testEqual(got, want) {
			t.Errorf("\t got %v(%T) for column[%v], want %v(%T)", got, got, row.fields[i].Name, want, want)
		}
	}
}

// Test decoding into nil destination.
func TestNilDst(t *testing.T) {
	for i, test := range []struct {
		r               *Row
		dst             interface{}
		wantErr         error
		structDst       interface{}
		wantToStructErr error
	}{
		{
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: stringType()},
				},
				[]*proto3.Value{stringProto("value")},
			},
			nil,
			errDecodeColumn(0, errNilDst(nil)),
			nil,
			errToStructArgType(nil),
		},
		{
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: stringType()},
				},
				[]*proto3.Value{stringProto("value")},
			},
			(*string)(nil),
			errDecodeColumn(0, errNilDst((*string)(nil))),
			(*struct{ STRING string })(nil),
			errNilDst((*struct{ STRING string })(nil)),
		},
		{
			&Row{
				[]*sppb.StructType_Field{
					{
						Name: "Col0",
						Type: listType(
							structType(
								mkField("Col1", intType()),
								mkField("Col2", floatType()),
								mkField("Col3", float32Type()),
								mkField("Col4", uuidType()),
							),
						),
					},
				},
				[]*proto3.Value{listProto(
					listProto(intProto(3), floatProto(33.3), float32Proto(0.3), uuidProto(uuid1)),
				)},
			},
			(*[]*struct {
				Col1 int
				Col2 float64
				Col3 float32
				Col4 uuid.UUID
			})(nil),
			errDecodeColumn(0, errNilDst((*[]*struct {
				Col1 int
				Col2 float64
				Col3 float32
				Col4 uuid.UUID
			})(nil))),
			(*struct {
				StructArray []*struct {
					Col1 int
					Col2 float64
					Col3 float32
					Col4 uuid.UUID
				} `spanner:"STRUCT_ARRAY"`
			})(nil),
			errNilDst((*struct {
				StructArray []*struct {
					Col1 int
					Col2 float64
					Col3 float32
					Col4 uuid.UUID
				} `spanner:"STRUCT_ARRAY"`
			})(nil)),
		},
	} {
		for j, toStuct := range []func(ptr interface{}) error{test.r.ToStruct, test.r.ToStructLenient} {
			if gotErr := test.r.Column(0, test.dst); !testEqual(gotErr, test.wantErr) {
				t.Errorf("%v: test.r.Column() returns error %v, want %v", i, gotErr, test.wantErr)
			}
			if gotErr := test.r.ColumnByName("Col0", test.dst); !testEqual(gotErr, test.wantErr) {
				t.Errorf("%v: test.r.ColumnByName() returns error %v, want %v", i, gotErr, test.wantErr)
			}
			// Row.Columns(T) should return nil on T == nil, otherwise, it should return test.wantErr.
			wantColumnsErr := test.wantErr
			if test.dst == nil {
				wantColumnsErr = nil
			}
			if gotErr := test.r.Columns(test.dst); !testEqual(gotErr, wantColumnsErr) {
				t.Errorf("%v: test.r.Columns() returns error %v, want %v", i, gotErr, wantColumnsErr)
			}
			if gotErr := toStuct(test.structDst); !testEqual(gotErr, test.wantToStructErr) {
				if j == 0 {
					t.Errorf("%v: test.r.ToStruct() returns error %v, want %v", i, gotErr, test.wantToStructErr)
				} else {
					t.Errorf("%v: test.r.ToStructLenient() returns error %v, want %v", i, gotErr, test.wantToStructErr)
				}
			}
		}
	}
}

// Test decoding NULL columns using Go types that don't support NULL.
func TestNullTypeErr(t *testing.T) {
	var tm time.Time
	ntoi := func(n string) int {
		for i, f := range row.fields {
			if f.Name == n {
				return i
			}
		}
		t.Errorf("cannot find column name %q in row", n)
		return 0
	}
	for _, test := range []struct {
		colName string
		dst     interface{}
	}{
		{
			"NULL_STRING",
			proto.String(""),
		},
		{
			"NULL_INT64",
			proto.Int64(0),
		},
		{
			"NULL_BOOL",
			proto.Bool(false),
		},
		{
			"NULL_FLOAT64",
			proto.Float64(0.0),
		},
		{
			"NULL_FLOAT32",
			proto.Float32(0.0),
		},
		{
			"NULL_TIMESTAMP",
			&tm,
		},
		{
			"NULL_DATE",
			&dt,
		},
		{
			"NULL_UUID",
			&uuid1,
		},
	} {
		wantErr := errDecodeColumn(ntoi(test.colName), errDstNotForNull(test.dst))
		if gotErr := row.ColumnByName(test.colName, test.dst); !testEqual(gotErr, wantErr) {
			t.Errorf("row.ColumnByName(%v) returns error %v, want %v", test.colName, gotErr, wantErr)
		}
	}
}

// Test using wrong destination type in column decoders.
func TestColumnTypeErr(t *testing.T) {
	// badDst cannot hold any of the column values.
	badDst := &struct{}{}
	for i, f := range row.fields { // For each of the columns, try to decode it into badDst.
		tc := f.Type.Code
		var etc sppb.TypeCode
		if strings.Contains(f.Name, "ARRAY") {
			etc = f.Type.ArrayElementType.Code
		}
		wantErr := errDecodeColumn(i, errTypeMismatch(tc, etc, badDst))
		if strings.Contains(f.Name, "STRUCT_ARRAY") {
			wantErr = errDecodeColumn(i, fmt.Errorf("the container is not a slice of struct pointers: %v", errTypeMismatch(tc, etc, badDst)))
		}
		if gotErr := row.Column(i, badDst); !testEqual(gotErr, wantErr) {
			t.Errorf("Column(%v): decoding into destination with wrong type %T returns error %v, want %v",
				i, badDst, gotErr, wantErr)
		}
		if gotErr := row.ColumnByName(f.Name, badDst); !testEqual(gotErr, wantErr) {
			t.Errorf("ColumnByName(%v): decoding into destination with wrong type %T returns error %v, want %v",
				f.Name, badDst, gotErr, wantErr)
		}
	}
	wantErr := errDecodeColumn(1, errTypeMismatch(sppb.TypeCode_STRING, sppb.TypeCode_TYPE_CODE_UNSPECIFIED, badDst))
	// badDst is used to receive column 1.
	vals := []interface{}{nil, badDst} // Row.Column() is expected to fail at column 1.
	// Skip decoding the rest columns by providing nils as the destinations.
	for i := 2; i < len(row.fields); i++ {
		vals = append(vals, nil)
	}
	if gotErr := row.Columns(vals...); !testEqual(gotErr, wantErr) {
		t.Errorf("Columns(): decoding column 1 with wrong type %T returns error %v, want %v",
			badDst, gotErr, wantErr)
	}
}

// Test the handling of invalid column decoding requests which cannot be mapped to correct column(s).
func TestInvalidColumnRequest(t *testing.T) {
	for _, test := range []struct {
		desc    string
		f       func() error
		wantErr error
	}{
		{
			"Request column index is out of range",
			func() error {
				return row.Column(10000, &struct{}{})
			},
			errColIdxOutOfRange(10000, &row),
		},
		{
			"Cannot find the named column",
			func() error {
				return row.ColumnByName("string", &struct{}{})
			},
			errColNotFound("string"),
		},
		{
			"Not enough arguments to call row.Columns()",
			func() error {
				return row.Columns(nil, nil)
			},
			errNumOfColValue(2, &row),
		},
		{
			"Call ColumnByName on row with duplicated column names",
			func() error {
				var s string
				r := &Row{
					[]*sppb.StructType_Field{
						{Name: "Val", Type: stringType()},
						{Name: "Val", Type: stringType()},
					},
					[]*proto3.Value{stringProto("value1"), stringProto("value2")},
				}
				return r.ColumnByName("Val", &s)
			},
			errDupColName("Val"),
		},
		{
			"Call ToStruct on row with duplicated column names",
			func() error {
				s := &struct {
					Val string
				}{}
				r := &Row{
					[]*sppb.StructType_Field{
						{Name: "Val", Type: stringType()},
						{Name: "Val", Type: stringType()},
					},
					[]*proto3.Value{stringProto("value1"), stringProto("value2")},
				}
				return r.ToStruct(s)
			},
			errDupSpannerField("Val", &sppb.StructType{
				Fields: []*sppb.StructType_Field{
					{Name: "Val", Type: stringType()},
					{Name: "Val", Type: stringType()},
				},
			}),
		},
		{
			"Call ToStruct on a row with unnamed field",
			func() error {
				s := &struct {
					Val string
				}{}
				r := &Row{
					[]*sppb.StructType_Field{
						{Name: "", Type: stringType()},
					},
					[]*proto3.Value{stringProto("value1")},
				}
				return r.ToStruct(s)
			},
			errUnnamedField(&sppb.StructType{Fields: []*sppb.StructType_Field{
				{Name: "", Type: stringType()},
			}}, 0),
		},
		{
			"Call ToStructLenient on row with duplicated column names",
			func() error {
				s := &struct {
					Val string
				}{}
				r := &Row{
					[]*sppb.StructType_Field{
						{Name: "Val", Type: stringType()},
						{Name: "Val", Type: stringType()},
					},
					[]*proto3.Value{stringProto("value1"), stringProto("value2")},
				}
				return r.ToStructLenient(s)
			},
			errDupSpannerField("Val", &sppb.StructType{
				Fields: []*sppb.StructType_Field{
					{Name: "Val", Type: stringType()},
					{Name: "Val", Type: stringType()},
				},
			}),
		},
		{
			"Call ToStructLenient on a row with unnamed field",
			func() error {
				s := &struct {
					Val string
				}{}
				r := &Row{
					[]*sppb.StructType_Field{
						{Name: "", Type: stringType()},
					},
					[]*proto3.Value{stringProto("value1")},
				}
				return r.ToStructLenient(s)
			},
			errUnnamedField(&sppb.StructType{Fields: []*sppb.StructType_Field{
				{Name: "", Type: stringType()},
			}}, 0),
		},
	} {
		if gotErr := test.f(); !testEqual(gotErr, test.wantErr) {
			t.Errorf("%v: test.f() returns error %v, want %v", test.desc, gotErr, test.wantErr)
		}
	}
}

// Test decoding the row with row.ToStruct into an invalid destination.
func TestToStructInvalidDst(t *testing.T) {
	for _, test := range []struct {
		desc     string
		dst      interface{}
		wantErr  error
		toStruct func(ptr interface{}) error
	}{
		{
			"row.ToStruct(): Decode row as STRUCT into int32",
			proto.Int32(1),
			errToStructArgType(proto.Int32(1)),
			row.ToStruct,
		},
		{
			"ToStructLenient(): Decode row as STRUCT into int32",
			proto.Int32(1),
			errToStructArgType(proto.Int32(1)),
			row.ToStructLenient,
		},
		{
			"row.ToStruct(): Decode row as STRUCT to nil Go struct",
			(*struct{})(nil),
			errNilDst((*struct{})(nil)),
			row.ToStruct,
		},
		{
			"row.ToStructLenient(): Decode row as STRUCT to nil Go struct",
			(*struct{})(nil),
			errNilDst((*struct{})(nil)),
			row.ToStructLenient,
		},
		{
			"row.ToStruct(): Decode row as STRUCT to Go struct with duplicated fields for the PK column",
			&struct {
				PK1 string `spanner:"STRING"`
				PK2 string `spanner:"STRING"`
			}{},
			errNoOrDupGoField(&struct {
				PK1 string `spanner:"STRING"`
				PK2 string `spanner:"STRING"`
			}{}, "STRING"),
			row.ToStruct,
		},
		{
			"row.ToStructLenient(): Decode row as STRUCT to Go struct with duplicated fields for the PK column",
			&struct {
				PK1 string `spanner:"STRING"`
				PK2 string `spanner:"STRING"`
			}{},
			errDupGoField(&struct {
				PK1 string `spanner:"STRING"`
				PK2 string `spanner:"STRING"`
			}{}, "STRING"),
			row.ToStructLenient,
		},
		{
			"row.ToStruct(): Decode row as STRUCT to Go struct with no field for the PK column",
			&struct {
				PK1 string `spanner:"_STRING"`
			}{},
			errNoOrDupGoField(&struct {
				PK1 string `spanner:"_STRING"`
			}{}, "STRING"),
			row.ToStruct,
		},
		{
			"row.ToStructLenient(): Decode row as STRUCT to Go struct with no field for the PK column",
			&struct {
				PK1 string `spanner:"_STRING"`
			}{},
			nil,
			row.ToStructLenient,
		},
		{
			"row.ToStruct(): Decode row as STRUCT to Go struct with wrong type for the PK column",
			&struct {
				PK1 int64 `spanner:"STRING"`
			}{},
			errDecodeStructField(&sppb.StructType{Fields: row.fields}, "STRING",
				errTypeMismatch(sppb.TypeCode_STRING, sppb.TypeCode_TYPE_CODE_UNSPECIFIED, proto.Int64(0))),
			row.ToStruct,
		},
		{
			"row.ToStructLenient(): Decode row as STRUCT to Go struct with wrong type for the PK column",
			&struct {
				PK1 int64 `spanner:"STRING"`
			}{},
			errDecodeStructField(&sppb.StructType{Fields: row.fields}, "STRING",
				errTypeMismatch(sppb.TypeCode_STRING, sppb.TypeCode_TYPE_CODE_UNSPECIFIED, proto.Int64(0))),
			row.ToStructLenient,
		},
	} {
		if gotErr := test.toStruct(test.dst); !testEqual(gotErr, test.wantErr) {
			t.Errorf("%v: decoding:\ngot  %v\nwant %v", test.desc, gotErr, test.wantErr)
		}
	}
}

// Test decoding a broken row.
func TestBrokenRow(t *testing.T) {
	for i, test := range []struct {
		row     *Row
		dst     interface{}
		wantErr error
	}{
		{
			// A row with no field.
			&Row{
				[]*sppb.StructType_Field{},
				[]*proto3.Value{stringProto("value")},
			},
			&NullString{"value", true},
			errFieldsMismatchVals(&Row{
				[]*sppb.StructType_Field{},
				[]*proto3.Value{stringProto("value")},
			}),
		},
		{
			// A row with nil field.
			&Row{
				[]*sppb.StructType_Field{nil},
				[]*proto3.Value{stringProto("value")},
			},
			&NullString{"value", true},
			errNilColType(0),
		},
		{
			// Field is not nil, but its type is nil.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: nil},
				},
				[]*proto3.Value{listProto(stringProto("value1"), stringProto("value2"))},
			},
			&[]NullString{},
			errDecodeColumn(0, errNilSpannerType()),
		},
		{
			// Field is not nil, field type is not nil, but it is an array and its array element type is nil.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: &sppb.Type{Code: sppb.TypeCode_ARRAY}},
				},
				[]*proto3.Value{listProto(stringProto("value1"), stringProto("value2"))},
			},
			&[]NullString{},
			errDecodeColumn(0, errNilArrElemType(&sppb.Type{Code: sppb.TypeCode_ARRAY})),
		},
		{
			// Field specifies valid type, value is nil.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: intType()},
				},
				[]*proto3.Value{nil},
			},
			&NullInt64{1, true},
			errDecodeColumn(0, errNilSrc()),
		},
		{
			// Field specifies INT64 type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: intType()},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_StringValue)(nil)}},
			},
			&NullInt64{1, true},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_StringValue)(nil)}, "String")),
		},
		{
			// Field specifies INT64 type, but value is for Number type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: intType()},
				},
				[]*proto3.Value{floatProto(1.0)},
			},
			&NullInt64{1, true},
			errDecodeColumn(0, errSrcVal(floatProto(1.0), "String")),
		},
		{
			// Field specifies INT64 type, but value is wrongly encoded.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: intType()},
				},
				[]*proto3.Value{stringProto("&1")},
			},
			proto.Int64(0),
			errDecodeColumn(0, errBadEncoding(stringProto("&1"), func() error {
				_, err := strconv.ParseInt("&1", 10, 64)
				return err
			}())),
		},
		{
			// Field specifies INT64 type, but value is wrongly encoded.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: intType()},
				},
				[]*proto3.Value{stringProto("&1")},
			},
			&NullInt64{},
			errDecodeColumn(0, errBadEncoding(stringProto("&1"), func() error {
				_, err := strconv.ParseInt("&1", 10, 64)
				return err
			}())),
		},
		{
			// Field specifies STRING type, but value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: stringType()},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_StringValue)(nil)}},
			},
			&NullString{"value", true},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_StringValue)(nil)}, "String")),
		},
		{
			// Field specifies STRING type, but value is for ARRAY type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: stringType()},
				},
				[]*proto3.Value{listProto(stringProto("value"))},
			},
			&NullString{"value", true},
			errDecodeColumn(0, errSrcVal(listProto(stringProto("value")), "String")),
		},
		{
			// Field specifies FLOAT64 type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: floatType()},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_NumberValue)(nil)}},
			},
			&NullFloat64{1.0, true},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_NumberValue)(nil)}, "Number")),
		},
		{
			// Field specifies FLOAT64 type, but value is for BOOL type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: floatType()},
				},
				[]*proto3.Value{boolProto(true)},
			},
			&NullFloat64{1.0, true},
			errDecodeColumn(0, errSrcVal(boolProto(true), "Number")),
		},
		{
			// Field specifies FLOAT64 type, but value is wrongly encoded.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: floatType()},
				},
				[]*proto3.Value{stringProto("nan")},
			},
			&NullFloat64{},
			errDecodeColumn(0, errUnexpectedFloat64Str("nan")),
		},
		{
			// Field specifies FLOAT64 type, but value is wrongly encoded.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: floatType()},
				},
				[]*proto3.Value{stringProto("nan")},
			},
			proto.Float64(0),
			errDecodeColumn(0, errUnexpectedFloat64Str("nan")),
		},
		{
			// Field specifies FLOAT32 type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: float32Type()},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_NumberValue)(nil)}},
			},
			&NullFloat32{1.0, true},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_NumberValue)(nil)}, "Number")),
		},
		{
			// Field specifies FLOAT32 type, but value is for BOOL type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: float32Type()},
				},
				[]*proto3.Value{boolProto(true)},
			},
			&NullFloat32{1.0, true},
			errDecodeColumn(0, errSrcVal(boolProto(true), "Number")),
		},
		{
			// Field specifies FLOAT32 type, but value is wrongly encoded.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: float32Type()},
				},
				[]*proto3.Value{stringProto("nan")},
			},
			&NullFloat32{},
			errDecodeColumn(0, errUnexpectedFloat32Str("nan")),
		},
		{
			// Field specifies FLOAT32 type, but value is wrongly encoded.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: float32Type()},
				},
				[]*proto3.Value{stringProto("nan")},
			},
			proto.Float32(0),
			errDecodeColumn(0, errUnexpectedFloat32Str("nan")),
		},
		{
			// Field specifies BYTES type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: bytesType()},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_StringValue)(nil)}},
			},
			&[]byte{},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_StringValue)(nil)}, "String")),
		},
		{
			// Field specifies BYTES type, but value is for BOOL type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: bytesType()},
				},
				[]*proto3.Value{boolProto(false)},
			},
			&[]byte{},
			errDecodeColumn(0, errSrcVal(boolProto(false), "String")),
		},
		{
			// Field specifies BYTES type, but value is wrongly encoded.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: bytesType()},
				},
				[]*proto3.Value{stringProto("&&")},
			},
			&[]byte{},
			errDecodeColumn(0, errBadEncoding(stringProto("&&"), func() error {
				_, err := base64.StdEncoding.DecodeString("&&")
				return err
			}())),
		},
		{
			// Field specifies BOOL type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: boolType()},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_BoolValue)(nil)}},
			},
			&NullBool{false, true},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_BoolValue)(nil)}, "Bool")),
		},
		{
			// Field specifies BOOL type, but value is for STRING type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: boolType()},
				},
				[]*proto3.Value{stringProto("false")},
			},
			&NullBool{false, true},
			errDecodeColumn(0, errSrcVal(stringProto("false"), "Bool")),
		},
		{
			// Field specifies TIMESTAMP type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: timeType()},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_StringValue)(nil)}},
			},
			&NullTime{time.Now(), true},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_StringValue)(nil)}, "String")),
		},
		{
			// Field specifies TIMESTAMP type, but value is for BOOL type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: timeType()},
				},
				[]*proto3.Value{boolProto(false)},
			},
			&NullTime{time.Now(), true},
			errDecodeColumn(0, errSrcVal(boolProto(false), "String")),
		},
		{
			// Field specifies TIMESTAMP type, but value is invalid timestamp.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: timeType()},
				},
				[]*proto3.Value{stringProto("junk")},
			},
			&NullTime{time.Now(), true},
			errDecodeColumn(0, errBadEncoding(stringProto("junk"), func() error {
				_, err := time.Parse(time.RFC3339Nano, "junk")
				return err
			}())),
		},
		{
			// Field specifies DATE type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: dateType()},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_StringValue)(nil)}},
			},
			&NullDate{civil.Date{}, true},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_StringValue)(nil)}, "String")),
		},
		{
			// Field specifies DATE type, but value is for BOOL type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: dateType()},
				},
				[]*proto3.Value{boolProto(false)},
			},
			&NullDate{civil.Date{}, true},
			errDecodeColumn(0, errSrcVal(boolProto(false), "String")),
		},
		{
			// Field specifies DATE type, but value is invalid timestamp.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: dateType()},
				},
				[]*proto3.Value{stringProto("junk")},
			},
			&NullDate{civil.Date{}, true},
			errDecodeColumn(0, errBadEncoding(stringProto("junk"), func() error {
				_, err := civil.ParseDate("junk")
				return err
			}())),
		},
		{
			// Field specifies UUID type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: uuidType()},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_StringValue)(nil)}},
			},
			&NullUUID{uuid1, true},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_StringValue)(nil)}, "String")),
		},
		{
			// Field specifies UUID type, but value is for BOOL type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: uuidType()},
				},
				[]*proto3.Value{boolProto(true)},
			},
			&NullUUID{uuid1, true},
			errDecodeColumn(0, errSrcVal(boolProto(true), "String")),
		},
		{
			// Field specifies UUID type, but value is incorrect UUID.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: uuidType()},
				},
				[]*proto3.Value{stringProto("xyz")},
			},
			&NullUUID{},
			errDecodeColumn(0, errBadEncoding(stringProto("xyz"), func() error {
				_, err := uuid.Parse("xyz")
				return err
			}())),
		},

		{
			// Field specifies ARRAY<INT64> type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(intType())},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_ListValue)(nil)}},
			},
			&[]NullInt64{},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_ListValue)(nil)}, "List")),
		},
		{
			// Field specifies ARRAY<INT64> type, value is having a nil ListValue.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(intType())},
				},
				[]*proto3.Value{{Kind: &proto3.Value_ListValue{}}},
			},
			&[]NullInt64{},
			errDecodeColumn(0, errNilListValue("INT64")),
		},
		{
			// Field specifies ARRAY<INT64> type, but value is for BYTES type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(intType())},
				},
				[]*proto3.Value{bytesProto([]byte("value"))},
			},
			&[]NullInt64{},
			errDecodeColumn(0, errSrcVal(bytesProto([]byte("value")), "List")),
		},
		{
			// Field specifies ARRAY<INT64> type, but value is for ARRAY<BOOL> type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(intType())},
				},
				[]*proto3.Value{listProto(boolProto(true))},
			},
			&[]NullInt64{},
			errDecodeColumn(0, errDecodeArrayElement(0, boolProto(true),
				"INT64", errSrcVal(boolProto(true), "String"))),
		},
		{
			// Field specifies ARRAY<STRING> type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(stringType())},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_ListValue)(nil)}},
			},
			&[]NullString{},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_ListValue)(nil)}, "List")),
		},
		{
			// Field specifies ARRAY<STRING> type, value is having a nil ListValue.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(stringType())},
				},
				[]*proto3.Value{{Kind: &proto3.Value_ListValue{}}},
			},
			&[]NullString{},
			errDecodeColumn(0, errNilListValue("STRING")),
		},
		{
			// Field specifies ARRAY<STRING> type, but value is for BOOL type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(stringType())},
				},
				[]*proto3.Value{boolProto(true)},
			},
			&[]NullString{},
			errDecodeColumn(0, errSrcVal(boolProto(true), "List")),
		},
		{
			// Field specifies ARRAY<STRING> type, but value is for ARRAY<BOOL> type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(stringType())},
				},
				[]*proto3.Value{listProto(boolProto(true))},
			},
			&[]NullString{},
			errDecodeColumn(0, errDecodeArrayElement(0, boolProto(true),
				"STRING", errSrcVal(boolProto(true), "String"))),
		},
		{
			// Field specifies ARRAY<FLOAT64> type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(floatType())},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_ListValue)(nil)}},
			},
			&[]NullFloat64{},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_ListValue)(nil)}, "List")),
		},
		{
			// Field specifies ARRAY<FLOAT64> type, value is having a nil ListValue.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(floatType())},
				},
				[]*proto3.Value{{Kind: &proto3.Value_ListValue{}}},
			},
			&[]NullFloat64{},
			errDecodeColumn(0, errNilListValue("FLOAT64")),
		},
		{
			// Field specifies ARRAY<FLOAT64> type, but value is for STRING type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(floatType())},
				},
				[]*proto3.Value{stringProto("value")},
			},
			&[]NullFloat64{},
			errDecodeColumn(0, errSrcVal(stringProto("value"), "List")),
		},
		{
			// Field specifies ARRAY<FLOAT64> type, but value is for ARRAY<BOOL> type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(floatType())},
				},
				[]*proto3.Value{listProto(boolProto(true))},
			},
			&[]NullFloat64{},
			errDecodeColumn(0, errDecodeArrayElement(0, boolProto(true),
				"FLOAT64", errSrcVal(boolProto(true), "Number"))),
		},
		{
			// Field specifies ARRAY<BYTES> type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(bytesType())},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_ListValue)(nil)}},
			},
			&[][]byte{},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_ListValue)(nil)}, "List")),
		},
		{
			// Field specifies ARRAY<BYTES> type, value is having a nil ListValue.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(bytesType())},
				},
				[]*proto3.Value{{Kind: &proto3.Value_ListValue{}}},
			},
			&[][]byte{},
			errDecodeColumn(0, errNilListValue("BYTES")),
		},
		{
			// Field specifies ARRAY<BYTES> type, but value is for FLOAT64 type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(bytesType())},
				},
				[]*proto3.Value{floatProto(1.0)},
			},
			&[][]byte{},
			errDecodeColumn(0, errSrcVal(floatProto(1.0), "List")),
		},
		{
			// Field specifies ARRAY<BYTES> type, but value is for ARRAY<FLOAT64> type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(bytesType())},
				},
				[]*proto3.Value{listProto(floatProto(1.0))},
			},
			&[][]byte{},
			errDecodeColumn(0, errDecodeArrayElement(0, floatProto(1.0),
				"BYTES", errSrcVal(floatProto(1.0), "String"))),
		},
		{
			// Field specifies ARRAY<BOOL> type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(boolType())},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_ListValue)(nil)}},
			},
			&[]NullBool{},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_ListValue)(nil)}, "List")),
		},
		{
			// Field specifies ARRAY<BOOL> type, value is having a nil ListValue.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(boolType())},
				},
				[]*proto3.Value{{Kind: &proto3.Value_ListValue{}}},
			},
			&[]NullBool{},
			errDecodeColumn(0, errNilListValue("BOOL")),
		},
		{
			// Field specifies ARRAY<BOOL> type, but value is for FLOAT64 type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(boolType())},
				},
				[]*proto3.Value{floatProto(1.0)},
			},
			&[]NullBool{},
			errDecodeColumn(0, errSrcVal(floatProto(1.0), "List")),
		},
		{
			// Field specifies ARRAY<BOOL> type, but value is for ARRAY<FLOAT64> type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(boolType())},
				},
				[]*proto3.Value{listProto(floatProto(1.0))},
			},
			&[]NullBool{},
			errDecodeColumn(0, errDecodeArrayElement(0, floatProto(1.0),
				"BOOL", errSrcVal(floatProto(1.0), "Bool"))),
		},
		{
			// Field specifies ARRAY<TIMESTAMP> type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(timeType())},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_ListValue)(nil)}},
			},
			&[]NullTime{},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_ListValue)(nil)}, "List")),
		},
		{
			// Field specifies ARRAY<TIMESTAMP> type, value is having a nil ListValue.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(timeType())},
				},
				[]*proto3.Value{{Kind: &proto3.Value_ListValue{}}},
			},
			&[]NullTime{},
			errDecodeColumn(0, errNilListValue("TIMESTAMP")),
		},
		{
			// Field specifies ARRAY<TIMESTAMP> type, but value is for FLOAT64 type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(timeType())},
				},
				[]*proto3.Value{floatProto(1.0)},
			},
			&[]NullTime{},
			errDecodeColumn(0, errSrcVal(floatProto(1.0), "List")),
		},
		{
			// Field specifies ARRAY<TIMESTAMP> type, but value is for ARRAY<FLOAT64> type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(timeType())},
				},
				[]*proto3.Value{listProto(floatProto(1.0))},
			},
			&[]NullTime{},
			errDecodeColumn(0, errDecodeArrayElement(0, floatProto(1.0),
				"TIMESTAMP", errSrcVal(floatProto(1.0), "String"))),
		},
		{
			// Field specifies ARRAY<DATE> type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(dateType())},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_ListValue)(nil)}},
			},
			&[]NullDate{},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_ListValue)(nil)}, "List")),
		},
		{
			// Field specifies ARRAY<DATE> type, value is having a nil ListValue.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(dateType())},
				},
				[]*proto3.Value{{Kind: &proto3.Value_ListValue{}}},
			},
			&[]NullDate{},
			errDecodeColumn(0, errNilListValue("DATE")),
		},
		{
			// Field specifies ARRAY<DATE> type, but value is for FLOAT64 type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(dateType())},
				},
				[]*proto3.Value{floatProto(1.0)},
			},
			&[]NullDate{},
			errDecodeColumn(0, errSrcVal(floatProto(1.0), "List")),
		},
		{
			// Field specifies ARRAY<DATE> type, but value is for ARRAY<FLOAT64> type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(dateType())},
				},
				[]*proto3.Value{listProto(floatProto(1.0))},
			},
			&[]NullDate{},
			errDecodeColumn(0, errDecodeArrayElement(0, floatProto(1.0),
				"DATE", errSrcVal(floatProto(1.0), "String"))),
		},
		{
			// Field specifies ARRAY<UUID> type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(uuidType())},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_ListValue)(nil)}},
			},
			&[]NullUUID{},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_ListValue)(nil)}, "List")),
		},
		{
			// Field specifies ARRAY<UUID> type, value is having a nil ListValue.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(uuidType())},
				},
				[]*proto3.Value{{Kind: &proto3.Value_ListValue{}}},
			},
			&[]NullUUID{},
			errDecodeColumn(0, errNilListValue("UUID")),
		},
		{
			// Field specifies ARRAY<UUID> type, but value is for FLOAT64 type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(uuidType())},
				},
				[]*proto3.Value{floatProto(1.0)},
			},
			&[]NullUUID{},
			errDecodeColumn(0, errSrcVal(floatProto(1.0), "List")),
		},
		{
			// Field specifies ARRAY<UUID> type, but value is for ARRAY<FLOAT64> type.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(uuidType())},
				},
				[]*proto3.Value{listProto(floatProto(1.0))},
			},
			&[]NullUUID{},
			errDecodeColumn(0, errDecodeArrayElement(0, floatProto(1.0),
				"UUID", errSrcVal(floatProto(1.0), "String"))),
		},
		{
			// Field specifies ARRAY<STRUCT> type, value is having a nil Kind.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(structType(
						mkField("Col1", intType()),
						mkField("Col2", floatType()),
						mkField("Col3", stringType()),
					))},
				},
				[]*proto3.Value{{Kind: (*proto3.Value_ListValue)(nil)}},
			},
			&[]*struct {
				Col1 int64
				Col2 float64
				Col3 string
			}{},
			errDecodeColumn(0, errSrcVal(&proto3.Value{Kind: (*proto3.Value_ListValue)(nil)}, "List")),
		},
		{
			// Field specifies ARRAY<STRUCT> type, value is having a nil ListValue.
			&Row{
				[]*sppb.StructType_Field{
					{Name: "Col0", Type: listType(structType(
						mkField("Col1", intType()),
						mkField("Col2", floatType()),
						mkField("Col3", stringType()),
					))},
				},
				[]*proto3.Value{{Kind: &proto3.Value_ListValue{}}},
			},
			&[]*struct {
				Col1 int64
				Col2 float64
				Col3 string
			}{},
			errDecodeColumn(0, errNilListValue("STRUCT")),
		},
		{
			// Field specifies ARRAY<STRUCT> type, value is having a nil ListValue.
			&Row{
				[]*sppb.StructType_Field{
					{
						Name: "Col0",
						Type: listType(
							structType(
								mkField("Col1", intType()),
								mkField("Col2", floatType()),
								mkField("Col3", stringType()),
							),
						),
					},
				},
				[]*proto3.Value{{Kind: &proto3.Value_ListValue{}}},
			},
			&[]NullRow{},
			errDecodeColumn(0, errNilListValue("STRUCT")),
		},
		{
			// Field specifies ARRAY<STRUCT> type, value is for BYTES type.
			&Row{
				[]*sppb.StructType_Field{
					{
						Name: "Col0",
						Type: listType(
							structType(
								mkField("Col1", intType()),
								mkField("Col2", floatType()),
								mkField("Col3", stringType()),
							),
						),
					},
				},
				[]*proto3.Value{bytesProto([]byte("value"))},
			},
			&[]*struct {
				Col1 int64
				Col2 float64
				Col3 string
			}{},
			errDecodeColumn(0, errSrcVal(bytesProto([]byte("value")), "List")),
		},
		{
			// Field specifies ARRAY<STRUCT> type, value is for BYTES type.
			&Row{
				[]*sppb.StructType_Field{
					{
						Name: "Col0",
						Type: listType(
							structType(
								mkField("Col1", intType()),
								mkField("Col2", floatType()),
								mkField("Col3", stringType()),
							),
						),
					},
				},
				[]*proto3.Value{listProto(bytesProto([]byte("value")))},
			},
			&[]NullRow{},
			errDecodeColumn(0, errNotStructElement(0, bytesProto([]byte("value")))),
		},
		{
			// Field specifies ARRAY<STRUCT> type, value is for ARRAY<BYTES> type.
			&Row{
				[]*sppb.StructType_Field{
					{
						Name: "Col0",
						Type: listType(
							structType(
								mkField("Col1", intType()),
								mkField("Col2", floatType()),
								mkField("Col3", stringType()),
							),
						),
					},
				},
				[]*proto3.Value{listProto(bytesProto([]byte("value")))},
			},
			&[]*struct {
				Col1 int64
				Col2 float64
				Col3 string
			}{},
			errDecodeColumn(0, errDecodeArrayElement(0, bytesProto([]byte("value")),
				"STRUCT", errSrcVal(bytesProto([]byte("value")), "List"))),
		},
		{
			// Field specifies ARRAY<STRUCT>, but is having nil StructType.
			&Row{
				[]*sppb.StructType_Field{
					{
						Name: "Col0", Type: listType(&sppb.Type{Code: sppb.TypeCode_STRUCT}),
					},
				},
				[]*proto3.Value{listProto(listProto(intProto(1), floatProto(2.0), stringProto("3")))},
			},
			&[]*struct {
				Col1 int64
				Col2 float64
				Col3 string
			}{},
			errDecodeColumn(0, errDecodeArrayElement(0, listProto(intProto(1), floatProto(2.0), stringProto("3")),
				"STRUCT", errNilSpannerStructType())),
		},
		{
			// Field specifies ARRAY<STRUCT>, but the second struct value is for BOOL type instead of FLOAT64.
			&Row{
				[]*sppb.StructType_Field{
					{
						Name: "Col0",
						Type: listType(
							structType(
								mkField("Col1", intType()),
								mkField("Col2", floatType()),
								mkField("Col3", stringType()),
							),
						),
					},
				},
				[]*proto3.Value{listProto(listProto(intProto(1), boolProto(true), stringProto("3")))},
			},
			&[]*struct {
				Col1 int64
				Col2 float64
				Col3 string
			}{},
			errDecodeColumn(
				0,
				errDecodeArrayElement(
					0, listProto(intProto(1), boolProto(true), stringProto("3")), "STRUCT",
					errDecodeStructField(
						&sppb.StructType{
							Fields: []*sppb.StructType_Field{
								mkField("Col1", intType()),
								mkField("Col2", floatType()),
								mkField("Col3", stringType()),
							},
						},
						"Col2",
						errSrcVal(boolProto(true), "Number"),
					),
				),
			),
		},
	} {
		if gotErr := test.row.Column(0, test.dst); !testEqual(gotErr, test.wantErr) {
			t.Errorf("%v: test.row.Column(0) got error %v, want %v", i, gotErr, test.wantErr)
		}
		if gotErr := test.row.ColumnByName("Col0", test.dst); !testEqual(gotErr, test.wantErr) {
			t.Errorf("%v: test.row.ColumnByName(%q) got error %v, want %v", i, "Col0", gotErr, test.wantErr)
		}
		if gotErr := test.row.Columns(test.dst); !testEqual(gotErr, test.wantErr) {
			t.Errorf("%v: test.row.Columns(%T) got error %v, want %v", i, test.dst, gotErr, test.wantErr)
		}
	}
}

// Test Row.ToStruct().
func TestToStruct(t *testing.T) {

	for i, toStuct := range []func(ptr interface{}) error{row.ToStruct, row.ToStructLenient} {
		s := []struct {
			// STRING / STRING ARRAY
			PrimaryKey      string       `spanner:"STRING"`
			NullString      NullString   `spanner:"NULL_STRING"`
			StringArray     []NullString `spanner:"STRING_ARRAY"`
			NullStringArray []NullString `spanner:"NULL_STRING_ARRAY"`
			// BYTES / BYTES ARRAY
			Bytes          []byte   `spanner:"BYTES"`
			NullBytes      []byte   `spanner:"NULL_BYTES"`
			BytesArray     [][]byte `spanner:"BYTES_ARRAY"`
			NullBytesArray [][]byte `spanner:"NULL_BYTES_ARRAY"`
			// INT64 / INT64 ARRAY
			Int64          int64       `spanner:"INT64"`
			NullInt64      NullInt64   `spanner:"NULL_INT64"`
			Int64Array     []NullInt64 `spanner:"INT64_ARRAY"`
			NullInt64Array []NullInt64 `spanner:"NULL_INT64_ARRAY"`
			// BOOL / BOOL ARRAY
			Bool          bool       `spanner:"BOOL"`
			NullBool      NullBool   `spanner:"NULL_BOOL"`
			BoolArray     []NullBool `spanner:"BOOL_ARRAY"`
			NullBoolArray []NullBool `spanner:"NULL_BOOL_ARRAY"`
			// FLOAT64 / FLOAT64 ARRAY
			Float64          float64       `spanner:"FLOAT64"`
			NullFloat64      NullFloat64   `spanner:"NULL_FLOAT64"`
			Float64Array     []NullFloat64 `spanner:"FLOAT64_ARRAY"`
			NullFloat64Array []NullFloat64 `spanner:"NULL_FLOAT64_ARRAY"`
			// FLOAT32 / FLOAT32 ARRAY
			Float32          float32       `spanner:"FLOAT32"`
			NullFloat32      NullFloat32   `spanner:"NULL_FLOAT32"`
			Float32Array     []NullFloat32 `spanner:"FLOAT32_ARRAY"`
			NullFloat32Array []NullFloat32 `spanner:"NULL_FLOAT32_ARRAY"`
			// TIMESTAMP / TIMESTAMP ARRAY
			Timestamp          time.Time  `spanner:"TIMESTAMP"`
			NullTimestamp      NullTime   `spanner:"NULL_TIMESTAMP"`
			TimestampArray     []NullTime `spanner:"TIMESTAMP_ARRAY"`
			NullTimestampArray []NullTime `spanner:"NULL_TIMESTAMP_ARRAY"`
			// DATE / DATE ARRAY
			Date          civil.Date `spanner:"DATE"`
			NullDate      NullDate   `spanner:"NULL_DATE"`
			DateArray     []NullDate `spanner:"DATE_ARRAY"`
			NullDateArray []NullDate `spanner:"NULL_DATE_ARRAY"`
			// UUID / UUID ARRAY
			UUID          uuid.UUID  `spanner:"UUID"`
			NullUUID      NullUUID   `spanner:"NULL_UUID"`
			UUIDArray     []NullUUID `spanner:"UUID_ARRAY"`
			NullUUIDArray []NullUUID `spanner:"NULL_UUID_ARRAY"`

			// STRUCT ARRAY
			StructArray []*struct {
				Col1 int64
				Col2 float64
				Col3 float32
				Col4 string
				Col5 uuid.UUID
			} `spanner:"STRUCT_ARRAY"`
			NullStructArray []*struct {
				Col1 int64
				Col2 float64
				Col3 float32
				Col4 string
				Col5 uuid.UUID
			} `spanner:"NULL_STRUCT_ARRAY"`
		}{
			{}, // got
			{
				// STRING / STRING ARRAY
				"value",
				NullString{},
				[]NullString{{"value1", true}, {}, {"value3", true}},
				[]NullString(nil),
				// BYTES / BYTES ARRAY
				[]byte("value"),
				[]byte(nil),
				[][]byte{[]byte("value1"), nil, []byte("value3")},
				[][]byte(nil),
				// INT64 / INT64 ARRAY
				int64(17),
				NullInt64{},
				[]NullInt64{{int64(1), true}, {int64(2), true}, {}},
				[]NullInt64(nil),
				// BOOL / BOOL ARRAY
				true,
				NullBool{},
				[]NullBool{{}, {true, true}, {false, true}},
				[]NullBool(nil),
				// FLOAT64 / FLOAT64 ARRAY
				1.7,
				NullFloat64{},
				[]NullFloat64{{}, {}, {1.7, true}},
				[]NullFloat64(nil),
				// FLOAT32 / FLOAT32 ARRAY
				float32(0.3),
				NullFloat32{},
				[]NullFloat32{{}, {}, {float32(0.3), true}},
				[]NullFloat32(nil),
				// TIMESTAMP / TIMESTAMP ARRAY
				tm,
				NullTime{},
				[]NullTime{{}, {tm, true}},
				[]NullTime(nil),
				// DATE / DATE ARRAY
				dt,
				NullDate{},
				[]NullDate{{}, {dt, true}},
				[]NullDate(nil),
				// UUID / UUID ARRAY
				uuid1,
				NullUUID{},
				[]NullUUID{{}, {uuid1, true}},
				[]NullUUID(nil),
				// STRUCT ARRAY
				[]*struct {
					Col1 int64
					Col2 float64
					Col3 float32
					Col4 string
					Col5 uuid.UUID
				}{
					nil,
					{3, 33.3, float32(0.3), "three", uuid1},
					nil,
				},
				[]*struct {
					Col1 int64
					Col2 float64
					Col3 float32
					Col4 string
					Col5 uuid.UUID
				}(nil),
			}, // want
		}
		funcName := "row.ToStruct()"
		if i != 0 {
			funcName = "row.ToStructLenient"
		}
		err := toStuct(&s[0])
		if err != nil {
			t.Errorf("%s returns error: %v, want nil", funcName, err)
		} else if !testEqual(s[0], s[1]) {
			t.Errorf("%s fetches struct %v, want %v", funcName, s[0], s[1])
		}
	}
}

// Test Row.ToStruct() with custom types.
func TestToStructWithCustomTypes(t *testing.T) {
	type CustomString string
	type CustomNullString NullString
	type CustomBytes []byte
	type CustomInt64 int64
	type CustomNullInt64 NullInt64
	type CustomBool bool
	type CustomNullBool NullBool
	type CustomFloat64 float64
	type CustomFloat32 float32
	type CustomNullFloat64 NullFloat64
	type CustomNullFloat32 NullFloat32
	type CustomTime time.Time
	type CustomNullTime NullTime
	type CustomDate civil.Date
	type CustomNullDate NullDate
	type CustomUUID uuid.UUID
	type CustomNullUUID NullUUID

	for i, toStuct := range []func(ptr interface{}) error{row.ToStruct, row.ToStructLenient} {
		s := []struct {
			// STRING / STRING ARRAY
			PrimaryKey      CustomString       `spanner:"STRING"`
			NullString      CustomNullString   `spanner:"NULL_STRING"`
			StringArray     []CustomNullString `spanner:"STRING_ARRAY"`
			NullStringArray []CustomNullString `spanner:"NULL_STRING_ARRAY"`
			// BYTES / BYTES ARRAY
			Bytes          CustomBytes   `spanner:"BYTES"`
			NullBytes      CustomBytes   `spanner:"NULL_BYTES"`
			BytesArray     []CustomBytes `spanner:"BYTES_ARRAY"`
			NullBytesArray []CustomBytes `spanner:"NULL_BYTES_ARRAY"`
			// INT64 / INT64 ARRAY
			Int64          CustomInt64       `spanner:"INT64"`
			NullInt64      CustomNullInt64   `spanner:"NULL_INT64"`
			Int64Array     []CustomNullInt64 `spanner:"INT64_ARRAY"`
			NullInt64Array []CustomNullInt64 `spanner:"NULL_INT64_ARRAY"`
			// BOOL / BOOL ARRAY
			Bool          CustomBool       `spanner:"BOOL"`
			NullBool      CustomNullBool   `spanner:"NULL_BOOL"`
			BoolArray     []CustomNullBool `spanner:"BOOL_ARRAY"`
			NullBoolArray []CustomNullBool `spanner:"NULL_BOOL_ARRAY"`
			// FLOAT64 / FLOAT64 ARRAY
			Float64          CustomFloat64       `spanner:"FLOAT64"`
			NullFloat64      CustomNullFloat64   `spanner:"NULL_FLOAT64"`
			Float64Array     []CustomNullFloat64 `spanner:"FLOAT64_ARRAY"`
			NullFloat64Array []CustomNullFloat64 `spanner:"NULL_FLOAT64_ARRAY"`
			// FLOAT32 / FLOAT32 ARRAY
			Float32          CustomFloat32       `spanner:"FLOAT32"`
			NullFloat32      CustomNullFloat32   `spanner:"NULL_FLOAT32"`
			Float32Array     []CustomNullFloat32 `spanner:"FLOAT32_ARRAY"`
			NullFloat32Array []CustomNullFloat32 `spanner:"NULL_FLOAT32_ARRAY"`
			// TIMESTAMP / TIMESTAMP ARRAY
			Timestamp          CustomTime       `spanner:"TIMESTAMP"`
			NullTimestamp      CustomNullTime   `spanner:"NULL_TIMESTAMP"`
			TimestampArray     []CustomNullTime `spanner:"TIMESTAMP_ARRAY"`
			NullTimestampArray []CustomNullTime `spanner:"NULL_TIMESTAMP_ARRAY"`
			// DATE / DATE ARRAY
			Date          CustomDate       `spanner:"DATE"`
			NullDate      CustomNullDate   `spanner:"NULL_DATE"`
			DateArray     []CustomNullDate `spanner:"DATE_ARRAY"`
			NullDateArray []CustomNullDate `spanner:"NULL_DATE_ARRAY"`

			// UUID / UUID ARRAY
			UUID          CustomUUID       `spanner:"UUID"`
			NullUUID      CustomNullUUID   `spanner:"NULL_UUID"`
			UUIDArray     []CustomNullUUID `spanner:"UUID_ARRAY"`
			NullUUIDArray []CustomNullUUID `spanner:"NULL_UUID_ARRAY"`

			// STRUCT ARRAY
			StructArray []*struct {
				Col1 CustomInt64
				Col2 CustomFloat64
				Col3 CustomFloat32
				Col4 CustomString
				Col5 CustomUUID
			} `spanner:"STRUCT_ARRAY"`
			NullStructArray []*struct {
				Col1 CustomInt64
				Col2 CustomFloat64
				Col3 CustomFloat32
				Col4 CustomString
				Col5 CustomUUID
			} `spanner:"NULL_STRUCT_ARRAY"`
		}{
			{}, // got
			{
				// STRING / STRING ARRAY
				"value",
				CustomNullString{},
				[]CustomNullString{{"value1", true}, {}, {"value3", true}},
				[]CustomNullString(nil),
				// BYTES / BYTES ARRAY
				CustomBytes("value"),
				CustomBytes(nil),
				[]CustomBytes{[]byte("value1"), nil, []byte("value3")},
				[]CustomBytes(nil),
				// INT64 / INT64 ARRAY
				CustomInt64(17),
				CustomNullInt64{},
				[]CustomNullInt64{{int64(1), true}, {int64(2), true}, {}},
				[]CustomNullInt64(nil),
				// BOOL / BOOL ARRAY
				true,
				CustomNullBool{},
				[]CustomNullBool{{}, {true, true}, {false, true}},
				[]CustomNullBool(nil),
				// FLOAT64 / FLOAT64 ARRAY
				1.7,
				CustomNullFloat64{},
				[]CustomNullFloat64{{}, {}, {1.7, true}},
				[]CustomNullFloat64(nil),
				// FLOAT32 / FLOAT32 ARRAY
				0.3,
				CustomNullFloat32{},
				[]CustomNullFloat32{{}, {}, {0.3, true}},
				[]CustomNullFloat32(nil),
				// TIMESTAMP / TIMESTAMP ARRAY
				CustomTime(tm),
				CustomNullTime{},
				[]CustomNullTime{{}, {tm, true}},
				[]CustomNullTime(nil),
				// DATE / DATE ARRAY
				CustomDate(dt),
				CustomNullDate{},
				[]CustomNullDate{{}, {dt, true}},
				[]CustomNullDate(nil),
				// UUID / UUID ARRAY
				CustomUUID(uuid1),
				CustomNullUUID{},
				[]CustomNullUUID{{}, {uuid1, true}},
				[]CustomNullUUID(nil),
				// STRUCT ARRAY
				[]*struct {
					Col1 CustomInt64
					Col2 CustomFloat64
					Col3 CustomFloat32
					Col4 CustomString
					Col5 CustomUUID
				}{
					nil,
					{3, 33.3, 0.3, "three", CustomUUID(uuid1)},
					nil,
				},
				[]*struct {
					Col1 CustomInt64
					Col2 CustomFloat64
					Col3 CustomFloat32
					Col4 CustomString
					Col5 CustomUUID
				}(nil),
			}, // want
		}
		funcName := "row.ToStruct()"
		if i != 0 {
			funcName = "row.ToStructLenient()"
		}
		err := toStuct(&s[0])
		if err != nil {
			t.Errorf("%s returns error: %v, want nil", funcName, err)
		} else if !testutil.Equal(s[0], s[1], cmp.AllowUnexported(CustomTime{})) {
			t.Errorf("%s fetches struct %v, want %v", funcName, s[0], s[1])
		}
	}
}

func TestToStructEmbedded(t *testing.T) {
	type (
		S1 struct{ F1 string }
		S2 struct {
			S1
			F2 string
		}
	)
	r := Row{
		[]*sppb.StructType_Field{
			{Name: "F1", Type: stringType()},
			{Name: "F2", Type: stringType()},
		},
		[]*proto3.Value{
			stringProto("v1"),
			stringProto("v2"),
		},
	}
	var got S2
	if err := r.ToStruct(&got); err != nil {
		t.Fatal(err)
	}
	want := S2{S1: S1{F1: "v1"}, F2: "v2"}
	if !testEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	var gotLenient S2
	if err := r.ToStructLenient(&gotLenient); err != nil {
		t.Fatal(err)
	}
	if !testEqual(gotLenient, want) {
		t.Errorf("gotLenient %+v, want %+v", got, want)
	}
}

func TestToStructWithUnEqualFields(t *testing.T) {
	type (
		extraField struct {
			F1 string
			F2 string
			F3 string
		}
		lessField struct {
			F1 string
			F3 string
		}
	)
	testCases := []struct {
		name string
		dst  interface{}
		want interface{}
		row  Row
	}{
		{
			name: "destination struct has extra field",
			dst:  &extraField{},
			row: Row{
				[]*sppb.StructType_Field{
					{Name: "F1", Type: stringType()},
					{Name: "F2", Type: stringType()},
				},
				[]*proto3.Value{
					stringProto("v1"),
					stringProto("v2"),
				},
			},
			want: &extraField{F1: "v1", F2: "v2", F3: ""},
		},
		{
			name: "destination struct has less field",
			dst:  &lessField{},
			row: Row{
				[]*sppb.StructType_Field{
					{Name: "F1", Type: stringType()},
					{Name: "F2", Type: stringType()},
					{Name: "F3", Type: stringType()},
				},
				[]*proto3.Value{
					stringProto("v1"),
					stringProto("v2"),
					stringProto("v3"),
				},
			},
			want: &lessField{F1: "v1", F3: "v3"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.row.ToStructLenient(tc.dst); err != nil {
				t.Fatal(err)
			}
			if !testEqual(tc.dst, tc.want) {
				t.Errorf("got %+v, want %+v", tc.dst, tc.want)
			}
		})
	}
}

func TestRowToString(t *testing.T) {
	r := Row{
		[]*sppb.StructType_Field{
			{Name: "F1", Type: stringType()},
			{Name: "F2", Type: stringType()},
		},
		[]*proto3.Value{
			stringProto("v1"),
			stringProto("v2"),
		},
	}
	got := r.String()
	want := `{fields: [name:"F1" type:{code:STRING} name:"F2" type:{code:STRING}], values: [string_value:"v1" string_value:"v2"]}`
	// In protobuf-go, the encoder will add an additional space based on a
	// deterministically random boolean value.
	wantWithTwoSpaces := `{fields: [name:"F1"  type:{code:STRING} name:"F2"  type:{code:STRING}], values: [string_value:"v1" string_value:"v2"]}`

	if !testEqual(r.String(), want) && !testEqual(r.String(), wantWithTwoSpaces) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

// Test helpers for getting column names.
func TestColumnNameAndIndex(t *testing.T) {
	// Test Row.Size().
	if rs := row.Size(); rs != len(row.fields) {
		t.Errorf("row.Size() returns %v, want %v", rs, len(row.fields))
	}
	// Test Row.Size() on empty Row.
	if rs := (&Row{}).Size(); rs != 0 {
		t.Errorf("empty_row.Size() returns %v, want %v", rs, 0)
	}
	// Test Row.ColumnName()
	for i, col := range row.fields {
		if cn := row.ColumnName(i); cn != col.Name {
			t.Errorf("row.ColumnName(%v) returns %q, want %q", i, cn, col.Name)
		}
		goti, err := row.ColumnIndex(col.Name)
		if err != nil {
			t.Errorf("ColumnIndex(%q) error %v", col.Name, err)
			continue
		}
		if goti != i {
			t.Errorf("ColumnIndex(%q) = %d, want %d", col.Name, goti, i)
		}
	}
	// Test Row.ColumnName on empty Row.
	if cn := (&Row{}).ColumnName(0); cn != "" {
		t.Errorf("empty_row.ColumnName(%v) returns %q, want %q", 0, cn, "")
	}
	// Test Row.ColumnIndex on empty Row.
	if _, err := (&Row{}).ColumnIndex(""); err == nil {
		t.Error("empty_row.ColumnIndex returns nil, want error")
	}
}

// Test helpers for getting column type and value.
func TestColumnTypeAndValue(t *testing.T) {
	// Test Row.ColumnType()
	for i, col := range row.fields {
		if ct := row.ColumnType(i); ct != col.Type {
			t.Errorf("row.ColumnType(%v) returns %q, want %q", i, ct, col.Type)
		}
	}
	// Test Row.ColumnValue()
	for i, val := range row.vals {
		if cv := row.ColumnValue(i); cv != val {
			t.Errorf("row.ColumnValue(%v) returns %q, want %q", i, cv, val)
		}
	}
	// Test Row.ColumnType on empty Row.
	if ct := (&Row{}).ColumnType(0); ct != nil {
		t.Errorf("empty_row.ColumnType(%v) returns %v, want %v", 0, ct, nil)
	}
	// Test Row.ColumnValue on empty Row.
	if cv := (&Row{}).ColumnValue(0); cv != nil {
		t.Errorf("empty_row.ColumnValue(%v) returns %v, want %v", 0, cv, nil)
	}
}

func TestNewRow(t *testing.T) {
	for _, test := range []struct {
		names   []string
		values  []interface{}
		want    *Row
		wantErr error
	}{
		{
			want: &Row{fields: []*sppb.StructType_Field{}, vals: []*proto3.Value{}},
		},
		{
			names:  []string{},
			values: []interface{}{},
			want:   &Row{fields: []*sppb.StructType_Field{}, vals: []*proto3.Value{}},
		},
		{
			names:   []string{"a", "b"},
			values:  []interface{}{},
			want:    nil,
			wantErr: errNamesValuesMismatch([]string{"a", "b"}, []interface{}{}),
		},
		{
			names:  []string{"a", "b", "c"},
			values: []interface{}{5, "abc", GenericColumnValue{listType(intType()), listProto(intProto(91), nullProto(), intProto(87))}},
			want: &Row{
				[]*sppb.StructType_Field{
					{Name: "a", Type: intType()},
					{Name: "b", Type: stringType()},
					{Name: "c", Type: listType(intType())},
				},
				[]*proto3.Value{
					intProto(5),
					stringProto("abc"),
					listProto(intProto(91), nullProto(), intProto(87)),
				},
			},
		},
	} {
		got, err := NewRow(test.names, test.values)
		if !testEqual(err, test.wantErr) {
			t.Errorf("NewRow(%v,%v).err = %s, want %s", test.names, test.values, err, test.wantErr)
			continue
		}
		if !testEqual(got, test.want) {
			t.Errorf("NewRow(%v,%v) = %s, want %s", test.names, test.values, got, test.want)
			continue
		}
	}
}

func BenchmarkColumn(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		if err := row.Column(0, &s); err != nil {
			b.Fatal(err)
		}
	}
}

type CustomType[T any] struct {
	Val T
}

func (n *CustomType[T]) DecodeSpanner(input any) error {
	switch val := input.(type) {
	case T:
		n.Val = val
		return nil
	case *T:
		if val == nil {
			return nil
		}
		n.Val = *val
		return nil
	}
	panic("n/a")
}

func TestSelectAll(t *testing.T) {
	skipUnsupportedPGTest(t)
	type args struct {
		destination interface{}
		options     []DecodeOptions
		mock        *mockRowIterator
	}
	type testStruct struct {
		Col1 int64
		// declaring second column in upper case here to verify SelectAll does case-insensitive matching
		COL2 float64
		Col3 string
		Col4 time.Time
	}

	type testStructWithCustom struct {
		Col1 int64
		COL2 float64
		Col3 CustomType[string]
		Col4 *CustomType[string]
	}

	type testStructWithTag struct {
		Col1 int64     `spanner:"tag1"`
		Col2 float64   `spanner:"Tag2"`
		Col3 string    `spanner:"taG3"`
		Col4 time.Time `spanner:"TAG4"`
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    interface{}
	}{
		{
			name: "success: using slice of primitives",
			args: args{
				destination: &[]string{},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col0", Type: stringType()},
						},
						[]*proto3.Value{stringProto("value")},
					},
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col0", Type: stringType()},
						},
						[]*proto3.Value{stringProto("value2")},
					},
					iterator.Done,
				),
			},
			want: &[]string{"value", "value2"},
		},
		{
			name: "success: using slice of pointer to primitives",
			args: args{
				destination: &[]*string{},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col0", Type: stringType()},
						},
						[]*proto3.Value{stringProto("value")},
					},
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col0", Type: stringType()},
						},
						[]*proto3.Value{stringProto("value2")},
					},
					iterator.Done,
				),
			},
			want: &[]*string{stringPointer("value"), stringPointer("value2")},
		},
		{
			name: "success: using slice of structs",
			args: args{
				destination: &[]testStruct{},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
							{Name: "Col4", Type: timeType()},
						},
						[]*proto3.Value{intProto(1), floatProto(1.1), stringProto("value"), timeProto(tm)},
					},
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
							{Name: "Col4", Type: timeType()},
						},
						[]*proto3.Value{intProto(2), floatProto(2.2), stringProto("value2"), timeProto(tm.Add(24 * time.Hour))},
					},
					iterator.Done,
				),
			},
			want: &[]testStruct{
				{Col1: 1, COL2: 1.1, Col3: "value", Col4: tm},
				{Col1: 2, COL2: 2.2, Col3: "value2", Col4: tm.Add(24 * time.Hour)},
			},
		},
		{
			name: "success: using slice of pointer to structs",
			args: args{
				destination: &[]*testStruct{},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
						},
						[]*proto3.Value{intProto(1), floatProto(1.1), stringProto("value")},
					},
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
						},
						[]*proto3.Value{intProto(2), floatProto(2.2), stringProto("value2")},
					},
					iterator.Done,
				),
			},
			want: &[]*testStruct{
				{Col1: 1, COL2: 1.1, Col3: "value"},
				{Col1: 2, COL2: 2.2, Col3: "value2"},
			}},
		{
			name: "success: when spanner row contains more columns than declared in Go struct but called WithLenient",
			args: args{
				destination: &[]*testStruct{},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
							{Name: "Col4", Type: timeType()},
							{Name: "Col5", Type: stringType()},
						},
						[]*proto3.Value{intProto(1), floatProto(1.1), stringProto("value"), timeProto(tm), stringProto("value2")},
					},
					// failure case
					iterator.Done,
				),
				options: []DecodeOptions{WithLenient()},
			},
			want: &[]*testStruct{
				{Col1: 1, COL2: 1.1, Col3: "value", Col4: tm},
			},
		},
		{
			name: "success: using prefilled destination should append to the destination",
			args: args{
				destination: &[]*testStruct{{Col1: 3, COL2: 3.3, Col3: "value3"}},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
						},
						[]*proto3.Value{intProto(1), floatProto(1.1), stringProto("value")},
					},
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
						},
						[]*proto3.Value{intProto(2), floatProto(2.2), stringProto("value2")},
					},
					iterator.Done,
				),
			},
			want: &[]*testStruct{
				{Col1: 3, COL2: 3.3, Col3: "value3"},
				{Col1: 1, COL2: 1.1, Col3: "value"},
				{Col1: 2, COL2: 2.2, Col3: "value2"},
			},
		},
		{
			name: "success: using destination with custom type with custom decoder with some null columns and pointer decoder",
			args: args{
				destination: &[]*testStructWithCustom{},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
							{Name: "Col4", Type: stringType()},
						},
						[]*proto3.Value{intProto(3), floatProto(3.3), stringProto("value3"), stringProto("test3")},
					},
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
							{Name: "Col4", Type: stringType()},
						},
						[]*proto3.Value{intProto(1), floatProto(1.1), nullProto(), nullProto()},
					},
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
							{Name: "Col4", Type: stringType()},
						},
						[]*proto3.Value{intProto(2), floatProto(2.2), stringProto("value2"), stringProto("test2")},
					},
					iterator.Done,
				),
			},
			want: &[]*testStructWithCustom{
				{Col1: 3, COL2: 3.3, Col3: CustomType[string]{"value3"}, Col4: &CustomType[string]{"test3"}},
				{Col1: 1, COL2: 1.1, Col3: CustomType[string]{}, Col4: nil},
				{Col1: 2, COL2: 2.2, Col3: CustomType[string]{"value2"}, Col4: &CustomType[string]{"test2"}},
			},
		},
		{
			name: "success: using slice of structs with spanner tag annotations",
			args: args{
				destination: &[]testStructWithTag{},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Tag1", Type: intType()},
							{Name: "Tag2", Type: floatType()},
							{Name: "Tag3", Type: stringType()},
							{Name: "Tag4", Type: timeType()},
						},
						[]*proto3.Value{intProto(1), floatProto(1.1), stringProto("value"), timeProto(tm)},
					},
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Tag1", Type: intType()},
							{Name: "Tag2", Type: floatType()},
							{Name: "Tag3", Type: stringType()},
							{Name: "Tag4", Type: timeType()},
						},
						[]*proto3.Value{intProto(2), floatProto(2.2), stringProto("value2"), timeProto(tm.Add(24 * time.Hour))},
					},
					iterator.Done,
				),
			},
			want: &[]testStructWithTag{
				{Col1: 1, Col2: 1.1, Col3: "value", Col4: tm},
				{Col1: 2, Col2: 2.2, Col3: "value2", Col4: tm.Add(24 * time.Hour)},
			},
		},
		{
			name: "failure: in case of error destination will have the partial result",
			args: args{
				destination: &[]*testStruct{{Col1: 3, COL2: 3.3, Col3: "value3"}},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
						},
						[]*proto3.Value{intProto(1), floatProto(1.1), stringProto("value")},
					},
					// failure case
					errors.New("some error"),
				),
			},
			want: &[]*testStruct{
				{Col1: 3, COL2: 3.3, Col3: "value3"},
				{Col1: 1, COL2: 1.1, Col3: "value"},
			},
			wantErr: true,
		},
		{
			name: "failure: when spanner row contains more columns than declared in Go struct",
			args: args{
				destination: &[]*testStruct{{Col1: 3, COL2: 3.3, Col3: "value3"}},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
							{Name: "Col4", Type: stringType()},
						},
						[]*proto3.Value{intProto(1), floatProto(1.1), stringProto("value")},
					},
					// failure case
					iterator.Done,
				),
			},
			want: &[]*testStruct{
				{Col1: 3, COL2: 3.3, Col3: "value3"},
			},
			wantErr: true,
		},
		{
			name: "failure: when spanner row contains more columns and destination is primitive slice",
			args: args{
				destination: &[]int64{},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
							{Name: "Col4", Type: stringType()},
						},
						[]*proto3.Value{intProto(1), floatProto(1.1), stringProto("value")},
					},
					// failure case
					iterator.Done,
				),
			},
			want:    &[]int64{},
			wantErr: true,
		},
		{
			name: "failure: when spanner row contains more columns and destination is primitive slice using WithLenient",
			args: args{
				destination: &[]int64{},
				mock: newMockIterator(
					&Row{
						[]*sppb.StructType_Field{
							{Name: "Col1", Type: intType()},
							{Name: "Col2", Type: floatType()},
							{Name: "Col3", Type: stringType()},
							{Name: "Col4", Type: stringType()},
						},
						[]*proto3.Value{intProto(1), floatProto(1.1), stringProto("value")},
					},
					// failure case
					iterator.Done,
				),
				options: []DecodeOptions{WithLenient()},
			},
			want:    &[]int64{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockIterator := tt.args.mock
			if err := SelectAll(mockIterator, tt.args.destination, tt.args.options...); (err != nil) != tt.wantErr {
				t.Errorf("SelectAll() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !testEqual(tt.args.destination, tt.want) {
				t.Errorf("SelectAll() = %v, want %v", tt.args.destination, tt.want)
			}
		})
	}
}

func stringPointer(s string) *string {
	return &s
}
