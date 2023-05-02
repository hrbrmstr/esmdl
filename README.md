# esmdl

Download ESM and all ESM deps from CDN

Ref:

- https://simonwillison.net/2023/May/2/download-esm/
- https://github.com/simonw/download-esm

## Usage

``` bash
NAME:
   download-esm - Download ESM modules from npm and jsdelivr

USAGE:
   download-esm [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --package value, -p value   Package to download
   --location value, -l value  Location to save files (default: .)
   --help, -h                  show help
```

``` bash
./esmdl --package "@observablehq/plot@latest" --location /tmp/mjs
```

    observablehq-plot-0-6-6.js
    interval-tree-1d-1-0-4.js
    d3-7-8-4.js
    isoformat-0-2-1.js
    d3-scale-4-0-2.js
    d3-random-3-0-1.js
    d3-force-3-0-0.js
    d3-format-3-1-0.js
    d3-shape-3-2-0.js
    d3-array-3-2-3.js
    d3-color-3-1-0.js
    d3-geo-3-1-0.js
    d3-time-3-1-0.js
    d3-quadtree-3-0-1.js
    d3-delaunay-6-0-4.js
    d3-hierarchy-3-1-2.js
    d3-dsv-3-0-1.js
    d3-polygon-3-0-1.js
    internmap-2-0-3.js
    d3-array-3-2-0.js
    d3-axis-3-0-0.js
    d3-timer-3-0-1.js
    d3-chord-3-0-1.js
    d3-dispatch-3-0-1.js
    d3-selection-3-0-0.js
    delaunator-5-0-0.js
    d3-path-3-1-0.js
    d3-zoom-3-0-0.js
    d3-contour-4-0-2.js
    d3-ease-3-0-1.js
    d3-transition-3-0-1.js
    d3-fetch-3-0-1.js
    d3-brush-3-0-0.js
    binary-search-bounds-2-0-5.js
    d3-interpolate-3-0-1.js
    d3-scale-chromatic-3-0-0.js
    d3-drag-3-0-0.js
    d3-time-format-4-1-0.js
    internmap-2-0-3.js
    d3-dsv-3-0-1.js
    d3-dispatch-3-0-1.js
    d3-selection-3-0-0.js
    d3-timer-3-0-1.js
    d3-ease-3-0-1.js
    d3-transition-3-0-1.js
    d3-interpolate-3-0-1.js
    d3-array-3-2-1.js
    robust-predicates-3-0-1.js
    d3-color-3-1-0.js
    d3-time-3-1-0.js
    internmap-2-0-3.js
    d3-array-3-2-0.js
    d3-timer-3-0-1.js
    d3-selection-3-0-0.js
    d3-dispatch-3-0-1.js
    d3-ease-3-0-1.js
    internmap-2-0-3.js
