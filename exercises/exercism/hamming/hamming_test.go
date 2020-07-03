package hamming

import "testing"

func TestHamming(t *testing.T) {
	for _, tc := range testCases {
		got, err := Distance(tc.s1, tc.s2)
		if tc.expectError {
			// check if err is of error type
			var _ error = err

			// we expect error
			if err == nil {
				t.Logf("Distance(%q, %q); expected error, got nil.",
					tc.s1, tc.s2)
				t.Fail()
			}
		} else {
			if got != tc.want {
				t.Logf("Distance(%q, %q) = %d, want %d.",
					tc.s1, tc.s2, got, tc.want)
				t.Fail()
			}

			// we do not expect error
			if err != nil {
				t.Logf("Distance(%q, %q) returned unexpected error: %v",
					tc.s1, tc.s2, err)
				t.Fail()
			}
		}
	}
}

func BenchmarkHamming(b *testing.B) {
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Distance(tc.s1, tc.s2)
		}
	}
}
