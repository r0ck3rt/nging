package checkinstall

import (
	"os/exec"

	"github.com/admpub/once"
)

func New(name string, checker ...func(name string) bool) *CheckInstall {
	c := &CheckInstall{name: name}
	if len(checker) > 0 && checker[0] != nil {
		c.checker = checker[0]
	} else {
		c.checker = DefaultChecker
	}
	return c
}

var DefaultChecker = func(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

type CheckInstall struct {
	name      string
	supported bool
	checkonce once.Once
	checker   func(name string) bool
}

func (c *CheckInstall) check() {
	c.supported = c.checker(c.name)
}

func (c *CheckInstall) IsSupported() bool {
	c.checkonce.Do(c.check)
	return c.supported
}

func (c *CheckInstall) ResetCheck() {
	c.checkonce.Reset()
}
