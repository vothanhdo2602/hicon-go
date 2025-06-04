export VERSION=v1.1.1

# make update-submodules branch=develop
update-submodules:
	git submodule update --init --recursive && \
	git submodule foreach git checkout $(branch) && \
	git submodule foreach git pull origin $(branch)

remove-tag:
	git tag -d ${VERSION} && \
	git push origin -d ${VERSION}

publish:
	git tag ${VERSION} && \
	git push origin ${VERSION} && \
	go list -m github.com/vothanhdo2602/hicon-go@${VERSION}
