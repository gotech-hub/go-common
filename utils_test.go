package utils

import (
	"context"
	"testing"
)

func TestGetRequestIdByContext(t *testing.T) {
	ctx1 := context.Background()
	traceInfo1 := GetRequestIdByContext(ctx1)
	if traceInfo1 != nil {
		t.Errorf("traceInfo1 should be nil")
	}

	ctx2 := context.WithValue(context.Background(), KeyTraceInfo, TraceInfo{RequestID: "123"})
	traceInfo2 := GetRequestIdByContext(ctx2)
	if traceInfo2 == nil {
		t.Errorf("traceInfo2 should not be nil")
	}

	if traceInfo2.RequestID != "123" {
		t.Errorf("traceInfo2.RequestID should be 123")
	}
}

func TestNewContextWithRequestId(t *testing.T) {
	ctx1 := context.Background()
	ctx2, traceInfo2 := NewContextWithRequestId(ctx1)
	if traceInfo2.RequestID == "" {
		t.Errorf("traceInfo2.RequestID should not be empty")
	}

	traceInfo3 := GetRequestIdByContext(ctx2)
	if traceInfo3 == nil {
		t.Errorf("traceInfo3 should not be nil")
	}

	if traceInfo3.RequestID != traceInfo2.RequestID {
		t.Errorf("traceInfo3.RequestID should be equal to traceInfo2.RequestID")
	}
}

func TestJoin(t *testing.T) {
	type args struct {
		elems interface{}
		delim string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "join int slice",
			args: args{
				elems: []int{1, 2, 3},
				delim: ",",
			},
			want:    "1,2,3",
			wantErr: false,
		},
		{
			name: "join string slice",
			args: args{
				elems: []string{"a", "b", "c"},
				delim: ",",
			},
			want:    "'a','b','c'",
			wantErr: false,
		},
		{
			name: "join float slice",
			args: args{
				elems: []float64{1.1, 2.2, 3.3},
				delim: ",",
			},
			want:    "1.1,2.2,3.3",
			wantErr: false,
		},
		{
			name: "join int slice with empty delim",
			args: args{
				elems: []int{1, 2, 3},
				delim: "",
			},
			want:    "123",
			wantErr: false,
		},
		{
			name: "join string slice with empty delim",
			args: args{
				elems: []string{"a", "b", "c"},
				delim: "",
			},
			want:    "'a''b''c'",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := join(tt.args.elems, tt.args.delim)
			if (err != nil) != tt.wantErr {
				t.Errorf("join() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("join() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertSqlTemplate(t *testing.T) {
	type args struct {
		query string
		args  map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				query: `select * from local_taydu_raw_login where level in ({{join.levels}})`,
				args: map[string]interface{}{
					"levels": []int{1, 2, 3},
				},
			},
			want:    `select * from local_taydu_raw_login where level in (1,2,3)`,
			wantErr: false,
		},
		{
			name: "test 2",
			args: args{
				query: `select * from local_taydu_raw_login where level in ({{join.levels}}) and id = {{.id}}`,
				args: map[string]interface{}{
					"levels": []int{1, 2, 3},
					"id":     123,
				},
			},
			want:    `select * from local_taydu_raw_login where level in (1,2,3) and id = 123`,
			wantErr: false,
		},
		{
			name: "test 3",
			args: args{
				query: `select * from local_taydu_raw_login where level in ({{join.levels}}) and id = {{.id}} and name = {{.name}}`,
				args: map[string]interface{}{
					"levels": []int{1, 2, 3},
					"id":     123,
					"name":   "abc",
				},
			},
			want:    `select * from local_taydu_raw_login where level in (1,2,3) and id = 123 and name = 'abc'`,
			wantErr: false,
		},
		{
			name: "test 4",
			args: args{
				query: `select * from local_taydu_raw_login where level in ({{join.levels}}) and id = {{.id}} and name = {{.name}} and age = {{.age}}`,
				args: map[string]interface{}{
					"levels": []int{1, 2, 3},
					"id":     123,
					"name":   "abc",
					"age":    18,
					"test":   "test",
				},
			},
			want:    `select * from local_taydu_raw_login where level in (1,2,3) and id = 123 and name = 'abc' and age = 18`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertSqlTemplate(tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertSqlTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("ConvertSqlTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
