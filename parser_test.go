/*
 *
 * timerParser_test.go
 * grapeTimer
 *
 * Created by lintao on 2020/7/7 2:45 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package grapeTimer

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestParserLoc(t *testing.T) {
	type args struct {
		DateFmt string
		Loc     *time.Location
	}
	type Tests struct {
		Name    string
		Args    args
		Want    *time.Time
		WantErr bool
	}

	test1 := func() Tests {
		tea := Tests{}
		local, _ := time.LoadLocation(LocationFormat)
		tt, err := AtTime("09:01:33", time.Now(), local)
		if err != nil {
			panic(err)
		}
		*tt = tt.AddDate(0, 0, 7)
		tea.Name = "解析周时间"
		tea.Args = args{
			DateFmt: fmt.Sprintf("Week %d 09:01:33", time.Now().Weekday()),
			Loc:     local,
		}
		tea.Want = tt
		tea.WantErr = false
		return tea
	}

	test2 := func() Tests {
		tea := Tests{}
		local, _ := time.LoadLocation(LocationFormat)
		tt, err := AtTime("15:15:33", time.Now(), local)
		if err != nil {
			panic(err)
		}
		tea.Name = "解析周时间"
		tea.Args = args{
			DateFmt: fmt.Sprintf("Week %d 15:15:33", time.Now().Weekday()),
			Loc:     local,
		}
		tea.Want = tt
		tea.WantErr = false
		return tea
	}

	tests := []Tests{
		test1(),
		test2(),
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := ParserLoc(tt.Args.DateFmt, time.Now(), tt.Args.Loc)
			if (err != nil) != tt.WantErr {
				t.Errorf("ParserLoc() error = %v, wantErr %v", err, tt.WantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("ParserLoc() got = %v, want %v", got, tt.Want)
			}
		})
	}
}

func TestParserStartAt(t *testing.T) {
	local, _ := time.LoadLocation(LocationFormat)
	tt, err := AtTime("09:01:33", time.Now().AddDate(0, 0, 7), local)
	startAt, err := AtTime("09:01:33", time.Now().AddDate(0, 0, 5), local)
	if err != nil {
		panic(err)
	}
	type args struct {
		startAt time.Time
		dateFmt string
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			name: "从起始时间解析",
			args: args{
				startAt: *startAt,
				dateFmt: fmt.Sprintf("Week %d 09:01:33", time.Now().Weekday()),
			},
			want:    tt,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParserStartAt(tt.args.startAt, tt.args.dateFmt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParserStartAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParserStartAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
