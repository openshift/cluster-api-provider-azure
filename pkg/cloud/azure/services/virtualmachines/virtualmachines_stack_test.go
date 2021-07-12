package virtualmachines

import (
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
)

func TestGetTagListFromSpecStackHub(t *testing.T) {
	testCases := []struct {
		spec     *StackHubSpec
		expected map[string]*string
	}{
		{
			spec: &StackHubSpec{
				Name: "test",
				Tags: map[string]string{
					"foo": "bar",
				},
			},
			expected: map[string]*string{
				"foo": to.StringPtr("bar"),
			},
		},
		{
			spec: &StackHubSpec{
				Name: "test",
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tagList := getTagListFromSpecStackHub(tc.spec)
		if !reflect.DeepEqual(tagList, tc.expected) {
			t.Errorf("Expected %v, got: %v", tc.expected, tagList)
		}
	}
}
