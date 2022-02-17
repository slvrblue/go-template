package storage_test

import (
	"errors"
	"testing"

	"github.com/blattaria7/go-template/internal/storage"
	"github.com/stretchr/testify/assert"
)

func Test_Get(t *testing.T) {
	testCase := []struct {
		name     string
		items    map[string]string
		id       string
		wantResp string
		wantErr  error
	}{
		{
			name: "succes",
			items: map[string]string{
				"key": "value",
			},
			id:       "key",
			wantResp: "value",
			wantErr:  nil,
		},
		{
			name: "not found",
			items: map[string]string{
				"key": "value",
			},
			id:       "lol",
			wantResp: "",
			wantErr:  errors.New("value lol not found"),
		},
	}

	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			s := storage.NewStorage(test.items)
			gotResp, err := s.Get(test.id)
			assert.Equal(t, test.wantResp, gotResp)

			if gotResp == "" && assert.Error(t, test.wantErr, err.Error()) {
				assert.EqualError(t, err, test.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
