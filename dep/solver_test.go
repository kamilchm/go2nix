package dep

import (
	"testing"

	"github.com/golang/dep/gps"
	"github.com/kamilchm/go2nix"
	"github.com/stretchr/testify/assert"
)

func TestConvertAllTargetFields(t *testing.T) {
	depPkgs := []gps.LockedProject{gps.NewLockedProject(
		gps.ProjectIdentifier{
			ProjectRoot: "github.com/golang/project",
			Source:      "",
		},
		gps.NewVersion("v0.1.3").Is(gps.Revision("123456")),
		[]string{},
	)}

	deps := convertDeps(depPkgs)

	assert.Len(t, deps, 1)

	assert.Equal(t, go2nix.ImportPath("github.com/golang/project"), deps[0].Name)
	assert.Equal(t, "v0.1.3", deps[0].Version)
	assert.Equal(t, "123456", deps[0].Revision)
}
