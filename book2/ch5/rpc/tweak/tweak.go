package tweak

import "strings"

type StringTweak struct{}

type Args struct {
	String  string
	ToUpper bool
	Reverse bool
}

func (s *StringTweak) Tweak(args *Args, resp *string) error {
	result := string(args.String)

	if args.ToUpper {
		result = strings.ToUpper(result)
	}
	if args.Reverse {
		runes := []rune(result)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result = string(runes)
	}

	*resp = result
	return nil
}
