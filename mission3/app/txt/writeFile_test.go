package txt

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"

	"com.mission/mission3/formatinput"
)

func TestWriteFileByLine(t *testing.T) {
	dataDateInput := "2020-04-1"

	type args struct {
		fileName  string
		textInput string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		wantErr  bool
		wantFail bool
	}{
		{
			name: "WriteFile success.",
			args: args{
				fileName:  formatinput.FileName(dataDateInput),
				textInput: "D|600000000024|16268.58|36|18|Promo2|2020-04-01|2020-06-30",
			},
			want:     "D|600000000024|16268.58|36|18|Promo2|2020-04-01|2020-06-30",
			wantErr:  false,
			wantFail: false,
		},
		{
			name: "WriteFile Fail.",
			args: args{
				fileName:  formatinput.FileName(dataDateInput),
				textInput: "D|600000000024|16268.58|36|18|Promo2",
			},
			want:     "D|600000000024|16268.58|36|18|Promo2|2020-04-01|2020-06-30",
			wantErr:  false,
			wantFail: true,
		},
	}
	for _, tt := range tests {
		//want = got, wanterror = true
		t.Run(tt.name, func(t *testing.T) {
			err := WriteFileByLine(tt.args.fileName, tt.args.textInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFileByLine() error = %v, wantErr %v", err, tt.wantErr)
			}
			f, err := os.Open(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFileByLine() error = %v, wantErr %v", err, tt.wantErr)
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)
			var got string
			for scanner.Scan() {

				got = scanner.Text()
			}
			if err := scanner.Err(); (err != nil) != tt.wantErr {
				t.Errorf("WriteFileByLine() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(string(got), tt.want) != tt.wantFail {
				t.Errorf("WriteFileByLine() = %v, want %v", got, tt.want)
			}
			os.Remove(tt.args.fileName)
			fmt.Println("pass")
		})

	}
}
