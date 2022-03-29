#
# Makefile
#
PROG=mp3play
usage:
	@echo "usage: make [build|run]"

build b:
	clang -o $(PROG) $(PROG).c `pkg-config --cflags --libs gstreamer-1.0`

run r:
	./$(PROG) $(HOME)/coding/media/audio/snow-fight.mp3

play p:
	gst-play-1.0  $(HOME)/coding/media/audio/snow-fight.mp3

clean:
	rm -f $(PROG)

kill k:
	pkill $(PROG)

git-update gu:
	cd ..; make gu


