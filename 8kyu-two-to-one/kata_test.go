package kata

import "testing"

func TestTwoToOne(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected string
	}{
		{
			str1:     "xyaabbbccccdefww",
			str2:     "xxxxyyyyabklmopq",
			expected: "abcdefklmopqwxy",
		},
		{
			str1:     "abcdefghijklmnopqrstuvwxyz",
			str2:     "abcdefghijklmnopqrstuvwxyz",
			expected: "abcdefghijklmnopqrstuvwxyz",
		},
	}

	for _, test := range tests {
		actual := TwoToOne(test.str1, test.str2)

		if actual != test.expected {
			t.Errorf("TwoToOne(%s, %s) = %s, expected %s", test.str1, test.str2, actual, test.expected)
		}
	}
}
