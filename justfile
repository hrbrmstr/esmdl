# this is a justfile <https://github.com/casey/just>

# Lists the justfile commands
@default:
  @just --list

# build the executable
@build:
  go build -trimpath -ldflags="-s -w" -buildmode=exe

# build & sign a macOS universal bniary
@macos-build:
  ./macos-build.sh

# remove build artifacts
@clean:
  rm -rf esmdl esmdl-macos esmdl-macos-arm esmdl-macos-x86
