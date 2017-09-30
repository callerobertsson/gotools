package human

import "testing"

func TestBytes(t *testing.T) {

	cases := []struct {
		in   int64
		dec  int
		long bool
		out  string
	}{
		{-1, 2, false, "-1.00 B"},
		{0, 2, false, "0.00 B"},
		{1, 2, false, "1.00 B"},
		{1023, 2, false, "1023.00 B"},
		{1024, 2, false, "1.00 KB"},
		{1025, 2, false, "1.00 KB"},
		{2048, 2, false, "2.00 KB"},
		{1000000, 2, false, "976.56 KB"},
		{1048575, 2, false, "1024.00 KB"},
		{1048576, 2, false, "1.00 MB"},
		{10000000, 2, false, "9.54 MB"},
		{100000000, 2, false, "95.37 MB"},
		{1073741824, 2, false, "1.00 GB"},
		{1234567890, 2, false, "1.15 GB"},
		{1099511627776, 2, false, "1.00 TB"},
		{1125899906842624, 2, false, "1.00 PB"},
		{1152921504606846976, 2, false, "1.00 EB"},
	}

	for _, c := range cases {
		res := Bytes(c.in, c.dec, c.long)
		if res != c.out {
			t.Errorf("Bytes(%d, %v) == %q, expected %q\n", c.in, c.long, res, c.out)
		}

	}
}

func TestKilos(t *testing.T) {

	cases := []struct {
		in   int64
		dec  int
		long bool
		out  string
	}{
		{-1, 2, false, "-1.00"},
		{0, 2, false, "0.00"},
		{1, 2, false, "1.00"},
		{999, 2, false, "999.00"},
		{1000, 2, false, "1.00 k"},
		{1001, 2, false, "1.00 k"},
		{2000, 2, false, "2.00 k"},
		{999999, 2, false, "1000.00 k"},
		{1000000, 2, false, "1.00 M"},
		{1000001, 2, false, "1.00 M"},
		{1234567, 2, false, "1.23 M"},
		{10000000, 2, false, "10.00 M"},
		{100000000, 2, false, "100.00 M"},
		{1073741824, 2, false, "1.07 G"},
		{1234567890, 2, false, "1.23 G"},
		{9876543210987, 2, false, "9.88 T"},
		{1099511627776, 2, false, "1.10 T"},
		{6546546546565465, 2, false, "6.55 P"},
		{7777777777777777777, 2, false, "7.78 E"},
	}

	for _, c := range cases {
		res := Kilos(c.in, c.dec, c.long)
		if res != c.out {
			t.Errorf("Kilos(%d, %v) == %q, expected %q\n", c.in, c.long, res, c.out)
		}

	}
}
