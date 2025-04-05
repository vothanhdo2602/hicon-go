# make update-submodules branch=develop
update-submodules:
	git submodule update --init --recursive && \
	git submodule foreach git checkout $(branch) && \
	git submodule foreach git pull origin $(branch)
