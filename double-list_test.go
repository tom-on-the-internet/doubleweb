package main

import (
	"math"
	"reflect"
	"sync"
	"testing"
)

func Test_newDoubleList(t *testing.T) {
	tests := []struct {
		name string
		want *doubleList
	}{
		{
			name: "it makes a doubleList",
			want: &doubleList{
				dblArr: []int{},
				dblMap: map[int]int{},
				mutex:  &sync.Mutex{},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := newDoubleList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDoubleList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doubleList_canHandle(t *testing.T) {
	type fields struct {
		dblArr []int
		dblMap map[int]int
		mutex  *sync.Mutex
	}

	type args struct {
		num int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "can handle small numbers",
			fields: fields{
				dblArr: []int{},
				dblMap: map[int]int{},
				mutex:  &sync.Mutex{},
			},
			args: args{
				num: 999,
			},
			want: true,
		},
		{
			name: "can handle negative numbers",
			fields: fields{
				dblArr: []int{},
				dblMap: map[int]int{},
				mutex:  &sync.Mutex{},
			},
			args: args{
				num: -999,
			},
			want: true,
		},
		{
			name: "cannot handle huge numbers",
			fields: fields{
				dblArr: []int{},
				dblMap: map[int]int{},
				mutex:  &sync.Mutex{},
			},
			args: args{
				num: math.MaxInt - 1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := &doubleList{
				dblArr: tt.fields.dblArr,
				dblMap: tt.fields.dblMap,
				mutex:  tt.fields.mutex,
			}
			if got := d.canHandle(tt.args.num); got != tt.want {
				t.Errorf("doubleList.canHandle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doubleList_has(t *testing.T) {
	type fields struct {
		dblArr []int
		dblMap map[int]int
		mutex  *sync.Mutex
	}

	type args struct {
		num int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "true when has",
			fields: fields{
				dblArr: []int{3, 4, 5},
				dblMap: map[int]int{3: 6, 4: 8},
				mutex:  &sync.Mutex{},
			},
			args: args{
				num: 3,
			},
			want: true,
		},
		{
			name: "false when has not",
			fields: fields{
				dblArr: []int{3, 4, 5},
				dblMap: map[int]int{3: 6, 4: 8},
				mutex:  &sync.Mutex{},
			},
			args: args{
				num: 9,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := &doubleList{
				dblArr: tt.fields.dblArr,
				dblMap: tt.fields.dblMap,
				mutex:  tt.fields.mutex,
			}
			if got := d.has(tt.args.num); got != tt.want {
				t.Errorf("doubleList.has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doubleList_list(t *testing.T) {
	type fields struct {
		dblArr []int
		dblMap map[int]int
		mutex  *sync.Mutex
	}

	tests := []struct {
		name   string
		fields fields
		want   [][2]string
	}{
		{
			name: "lists elements",
			fields: fields{
				dblArr: []int{2, 8, 9, 11},
				dblMap: map[int]int{2: 4, 8: 16, 11: 22},
				mutex:  &sync.Mutex{},
			},
			want: [][2]string{{"2", "4"}, {"8", "16"}, {"9", "⏳"}, {"11", "22"}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := &doubleList{
				dblArr: tt.fields.dblArr,
				dblMap: tt.fields.dblMap,
				mutex:  tt.fields.mutex,
			}
			if got := d.list(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doubleList.list() = %v, want %v", got, tt.want)
			}
		})
	}
}
