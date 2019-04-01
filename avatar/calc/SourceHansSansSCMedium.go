package calc

import (
	"regexp"

	"github.com/g-airport/go-infra/avatar"
)

type SourceHansSansSCMedium struct{}

func (f *SourceHansSansSCMedium) CalculateCenterLocation(s string, ab *avatarbuilder.AvatarBuilder) (x int, y int) {
	cr := regexp.MustCompile("[\u4e00-\u9FA5]")
	er := regexp.MustCompile("[a-zA-Z]")
	nr := regexp.MustCompile("[0-9]")

	cCount := len(cr.FindAllStringSubmatch(s, -1))
	eCount := len(er.FindAllStringSubmatch(s, -1))
	nCount := len(nr.FindAllStringSubmatch(s, -1))

	x = ab.W/2 - (cCount*ab.GetFontWidth()+eCount*ab.GetFontWidth()*3/5+nCount*ab.GetFontWidth()*3/5)/2
	y = ab.H/2 + ab.GetFontWidth()*4/11

	return
}
