#
# Makefile
#
PROG=appsrc
usage:
	@echo "usage: make [build|run]"

build b:
	clang -o $(PROG) $(PROG).c `pkg-config --cflags --libs gstreamer-1.0 gstreamer-app-1.0`

run r:
	./$(PROG) /Users/stoney/coding/media/open/bunny.mp4

clean:
	rm -f $(PROG)

kill k:
	pkill $(PROG)


git-update gu:
	cd ..; make gu


