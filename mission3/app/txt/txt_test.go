package txt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"com.mission/mission3/entity"
	"com.mission/mission3/formatinput"
)

func TestWriteHeader(t *testing.T) {
	dataDateInput := "2020-03-20"
	dataDate := strings.Replace(dataDateInput, "-", "", -1)

	type args struct {
		fileName string
		dataDate string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		wantErr  bool
		wantFail bool
	}{
		{
			name: "Write Header Success",
			args: args{
				fileName: formatinput.FileName(dataDateInput),
				dataDate: dataDate,
			},
			want:     fmt.Sprintf("H|%s", dataDate),
			wantErr:  false,
			wantFail: false,
		},
		{
			name: "Write Header User Fail",
			args: args{
				fileName: formatinput.FileName(dataDateInput),
				dataDate: dataDate,
			},
			want:     fmt.Sprintf("H|%s%s", dataDate, dataDate),
			wantErr:  false,
			wantFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteHeader(tt.args.fileName, tt.args.dataDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteHeader() error = %v, wantErr %v", err, tt.wantErr)
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
			os.Remove(formatinput.FileName(dataDateInput))
		})
	}
}

func TestWriteRecord(t *testing.T) {
	dataDateInput := "2020-01-01"
	var startDate entity.Date
	json.Unmarshal([]byte(`"2020-01-01"`), &startDate)
	var endDate entity.Date
	json.Unmarshal([]byte(`"2020-03-31"`), &endDate)

	type args struct {
		fileName  string
		dataInput entity.WriteRecordInput
	}
	tests := []struct {
		name     string
		args     args
		want     string
		wantErr  bool
		wantFail bool
	}{
		{
			name: "Write record Success",
			args: args{
				fileName: formatinput.FileName(dataDateInput),
				dataInput: entity.WriteRecordInput{
					AccountNumber:     "600000000016",
					InstallmentAmount: 29563.14,
					NumberOfPayment:   12,
					InterestRate:      2.5,
					PromotionName:     "Promo1",
					StartDate:         startDate,
					EndDate:           endDate,
				},
			},
			want:     "D|600000000016|29563.14|12|2.5|Promo1|2020-01-01|2020-03-31",
			wantErr:  false,
			wantFail: false,
		},
		{
			name: "Write record User Fail",
			args: args{
				fileName: formatinput.FileName(dataDateInput),
				dataInput: entity.WriteRecordInput{
					AccountNumber:     "600000000016",
					InstallmentAmount: 295645,
					NumberOfPayment:   12,
					InterestRate:      2.5,
					PromotionName:     "Promo1",
					StartDate:         startDate,
					EndDate:           endDate,
				},
			},
			want:     "D|600000000016|29563.14|12|2.5|Promo1|2020-01-01|2020-03-31",
			wantErr:  false,
			wantFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteRecord(tt.args.fileName, tt.args.dataInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteRecord() error = %v, wantErr %v", err, tt.wantErr)
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
			os.Remove(formatinput.FileName(dataDateInput))
		})
	}
}

func TestWriteFooter(t *testing.T) {
	dataDateInput := "2020-01-01"

	type args struct {
		fileName    string
		totalRecord int
		totalAmount float64
	}
	tests := []struct {
		name     string
		args     args
		want     string
		wantErr  bool
		wantFail bool
	}{
		{
			name: "Write record Success",
			args: args{
				fileName:    formatinput.FileName(dataDateInput),
				totalRecord: 1,
				totalAmount: 29563.14,
			},
			want:     "T|1|29563.14",
			wantErr:  false,
			wantFail: false,
		},
		{
			name: "Write record Fail",
			args: args{
				fileName:    formatinput.FileName(dataDateInput),
				totalRecord: 2,
				totalAmount: 29563.14,
			},
			want:     "T|1|29563.14",
			wantErr:  false,
			wantFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteFooter(tt.args.fileName, tt.args.totalRecord, tt.args.totalAmount)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFooter() error = %v, wantErr %v", err, tt.wantErr)
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
			os.Remove(formatinput.FileName(dataDateInput))
		})
	}
}
