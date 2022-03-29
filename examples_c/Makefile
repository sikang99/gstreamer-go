#
# Makefile
#
PROG=helloworld
usage:
	@echo "usage: make [build|run]"

build b:
	clang -o $(PROG) $(PROG).c `pkg-config --cflags --libs gstreamer-1.0`

run r:
	./$(PROG) /Users/stoney/coding/media/open/bunny.mp4

git-update gu:
	cd ..; make gu


