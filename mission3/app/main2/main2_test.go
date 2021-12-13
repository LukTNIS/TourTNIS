package main2

import (
	"os"
	"testing"

	"com.mission/mission3/entity"
)

func TestRecordInstallmentAmountListByMonth(t *testing.T) {
	type args struct {
		dbConfig entity.ConfigDBInput
	}
	tests := []struct {
		name    string
		args    args
		osArgs  []string
		wantErr bool
	}{
		{
			name: "RecordInstallmentAmountListByMonth Success",
			args: args{
				dbConfig: entity.ConfigDBInput{
					Host:     "localhost",
					Port:     5436,
					Username: "admin",
					Password: "admin@1234",
					Dbname:   "core_banking",
				},
			},
			osArgs:  []string{"go run main2.go", "2020-04-01"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.osArgs
			RecordInstallmentAmountListByMonth(tt.args.dbConfig)
			os.Remove("INSTALLMENT_20200401.txt")
		})
	}
}
