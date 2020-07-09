/*
 *
 * scheduler_test.go
 * grapeTimer
 *
 * Created by lintao on 2020/7/7 5:05 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package grapeTimer

import (
	"testing"
	"time"
)

func TestInitGrapeScheduler(t *testing.T) {
	type args struct {
		t   time.Duration
		ars bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				t:   1000,
				ars: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
