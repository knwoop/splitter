package main

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	fixtures = []string{
		`header1,header2
"1111","22222"
33333,44444
55555,66666
77777,88888`,
		`"1111","22222"
33333,44444
55555,66666
77777,88888`,
	}
)

func TestSplit(t *testing.T) {
	type args struct {
		reader    io.Reader
		hasHeader bool
		partSize1 int
		partSize2 int
	}
	type want struct {
		r1  string
		r2  string
		err error
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Success with Header",
			args: args{
				reader:    strings.NewReader(fixtures[0]),
				hasHeader: true,
				partSize1: 2,
				partSize2: 2,
			},
			want: want{
				r1: `header1,header2
"1111","22222"
33333,44444
`,
				r2: `header1,header2
55555,66666
77777,88888
`,
			},
			wantErr: false,
		},
		{
			name: "Success not match size with Header",
			args: args{
				reader:    strings.NewReader(fixtures[0]),
				hasHeader: true,
				partSize1: 1,
				partSize2: 3,
			},
			want: want{
				r1: `header1,header2
"1111","22222"
`,
				r2: `header1,header2
33333,44444
55555,66666
77777,88888
`,
			},
			wantErr: false,
		},
		{
			name: "Success with No header",
			args: args{
				reader:    strings.NewReader(fixtures[1]),
				hasHeader: false,
				partSize1: 2,
				partSize2: 2,
			},
			want: want{
				r1: `"1111","22222"
33333,44444
`,
				r2: `55555,66666
77777,88888
`,
			},
			wantErr: false,
		},
		{
			name: "Success not match size with No header",
			args: args{
				reader:    strings.NewReader(fixtures[1]),
				hasHeader: false,
				partSize1: 1,
				partSize2: 3,
			},
			want: want{
				r1: `"1111","22222"
`,
				r2: `33333,44444
55555,66666
77777,88888
`,
			},
			wantErr: false,
		},
		{
			name: "error not match sizes with No header",
			args: args{
				reader:    strings.NewReader(fixtures[1]),
				hasHeader: false,
				partSize1: 3,
				partSize2: 3,
			},
			want:    want{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, err := Split(tt.args.reader, tt.args.hasHeader, tt.args.partSize1, tt.args.partSize2)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Split(%v, %t, %d, %d) errpr = %v, wantErr %t",
						tt.args.reader, tt.args.hasHeader, tt.args.partSize1, tt.args.partSize2, err, tt.wantErr)
				}
				return
			}
			buf1 := new(bytes.Buffer)
			if _, err := buf1.ReadFrom(got1); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(buf1.String(), tt.want.r1); diff != "" {
				t.Errorf("(-got +want)\n%v", diff)
			}

			buf2 := new(bytes.Buffer)
			if _, err := buf2.ReadFrom(got2); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(buf2.String(), tt.want.r2); diff != "" {
				t.Errorf("(-got +want)\n%v", diff)
			}
		})
	}
}
