package piscine

func StrRev(s string) string {
	rev := ""
	for i := 0; i < len(s); i++ {
		rev = string(s[i]) + rev
	}
	return rev
}
