nextTag:= $(shell autotag --scheme=conventional -n)
currentTag:=$(shell git describe --tags)

changelog:
	git-chglog -o CHANGELOG.md --next-tag v$(nextTag)

release:
	git-chglog -o CHANGELOG.md --next-tag v$(nextTag)
	git add CHANGELOG.md
	git commit -m "CHANGELOG.md update v$(nextTag)"
	git tag v$(nextTag)