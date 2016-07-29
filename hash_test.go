package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHash(t *testing.T) {
	Convey("Given nix-prefetch-git output", t, func() {
		prefetchOut := []byte(`2015/10/27 16:44:07 git revision is d30f09973e19c1dfcd120b2d9c4f168e68d6b5d5
Commit date is 2015-09-16 13:57:42 -0700
3f8e0f809c43744023db06ce47a271dd521db5441c5a11a0551b86b077158035
`)
		Convey("When we parse it", func() {
			hash := hashFromNixPrefetch("notgit", prefetchOut)

			Convey("Then we should have extracted nix-hash", func() {
				So(hash, ShouldEqual,
					"3f8e0f809c43744023db06ce47a271dd521db5441c5a11a0551b86b077158035")
			})
		})
	})
}
