package main

import (
	"fmt"
	"strings"
)

type Rule struct {
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Id    string `json:"id"`
}

func (g *Rule) String() string {
	return fmt.Sprintf("value:%s, tag:%s id:%s\n", g.Value, g.Tag, g.Id)
}

type RuleList []Rule

func (g RuleList) String() string {
	slist := make([]string, 0)
	for _, v := range g {
		slist = append(slist, v.String())
	}
	return strings.Join(slist, "")
}

func main() {
	g1 := Rule{Value: "hello", Tag: "1", Id: "1"}
	g2 := Rule{Value: "hello", Tag: "1", Id: "1"}
	g3 := Rule{Value: "hello", Tag: "1", Id: "1"}

	glist := RuleList{g1, g2, g3}

	fmt.Println(glist)

}
