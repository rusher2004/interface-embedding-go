package stuff_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/rusher2004/interface-embedding-go/stuff"
)

type mockThis struct {
	stuff.Thinger

	exp stuff.Thing
	err error
}

func (m mockThis) DoThisThing() (stuff.Thing, error) {
	return m.exp, m.err
}

func TestDoSomeStuff(t *testing.T) {
	tests := []struct {
		name    string
		want    stuff.Thing
		wantErr error
		mock    mockThis
	}{
		// happy path
		{
			name: "happy path",
			want: stuff.Thing{ID: "123"},
			mock: mockThis{
				exp: stuff.Thing{ID: "123"},
				err: nil,
			},
		},
		// errors
		{
			name:    "errors",
			wantErr: errors.New("error doing some stuff: something went wrong"),
			mock: mockThis{
				err: errors.New("something went wrong"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stuff.DoSomeStuff(tt.mock)

			if tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("%s:\ngot error: %v\nwanted: %v", tt.name, err, tt.wantErr)
					return
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoSomeStuff() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockThat struct {
	stuff.Thinger

	err error
}

func (m mockThat) DoThatThing() error {
	return m.err
}

func TestDoOtherStuff(t *testing.T) {
	tests := []struct {
		name    string
		wantErr error
		mock    mockThat
	}{
		// happy path
		{
			name: "happy path",
			mock: mockThat{
				err: nil,
			},
		},
		// errors
		{
			name:    "errors",
			wantErr: errors.New("error doing that thing: something went wrong"),
			mock: mockThat{
				err: errors.New("something went wrong"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := stuff.DoOtherStuff(tt.mock)

			if tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("%s:\ngot error: %v\nwanted: %v", tt.name, err, tt.wantErr)
					return
				}
			}
		})
	}
}
