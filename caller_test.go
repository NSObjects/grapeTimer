/*
 *
 * caller_test.go
 * grapeTimer
 *
 * Created by lintao on 2020/7/7 5:07 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package grapeTimer

import "testing"

func TestNewTimeDataLoop(t *testing.T) {
	type args struct {
		data  string
		count int
		fn    GrapeExecFn
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "循环日期",
			args: args{
				data:  "",
				count: 1,
				fn:    nil,
				args:  nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := NewTimeDataLoop(tt.args.data, tt.args.count, tt.args.fn, tt.args.args...); got != tt.want {
			//	t.Errorf("NewTimeDataLoop() = %v, want %v", got, tt.want)
			//}
		})
	}
}
