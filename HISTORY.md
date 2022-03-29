## Change History

- 2022/03/29
    - gstreamer sample source of [helloworld.c](https://github.com/GStreamer/gstreamer/tree/master/tests/examples/helloworld) does not also work.

- 2022/03/29
    - not work with gstreamer 2.0 after `brew upgrade gstreamer` from gstreamer 1.18.5, glib 2.70.4
    - resync to original source [notedit/gstreamer-go](https://github.com/notedit/gstreamer-go)

- 2021/03/30
    - change defrecated memdup() into [memdup2()](https://developer.gnome.org/glib/stable/glib-Memory-Allocation.html#g-memdup2) after glib 2.67.4 in gstreamer.c

- 2021/03/30
    - fork [notedit/gstreamer-go](https://github.com/notedit/gstreamer-go)



### Open Source
- libs.garden: [Go/gstreamer](https://libs.garden/go/search?q=gstreamer)