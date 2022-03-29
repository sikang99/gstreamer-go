#
# Makefile
#
PROG=appsrc
usage:
	@echo "usage: make [build|run]"

build b:
	clang -o $(PROG) $(PROG).c `pkg-config --cflags --libs gstreamer-1.0 gstreamer-app-1.0`

b2:
	clang -o $(PROG) $(PROG).c -I/Library/Frameworks/GStreamer.framework/Headers -L/Library/Frameworks/GStreamer.framework/Libraries -lgstapp-1.0 -lgstbase-1.0 -lgstreamer-1.0 -lgobject-2.0 -lglib-2.0 -lintl

run r:
	./$(PROG) /Users/stoney/coding/media/open/bunny.mp4

clean:
	rm -f $(PROG)

kill k:
	pkill $(PROG)


git-update gu:
	cd ..; make gu


