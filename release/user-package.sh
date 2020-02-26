#!/usr/bin/env bash
# Bash3 Boilerplate. Copyright (c) 2014, kvz.io

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

trap 'echo -e "Aborted, error $? in command: $BASH_COMMAND"; trap ERR; exit 1' ERR

# Set magic variables for current file & dir
__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
__file="${__dir}/$(basename "${BASH_SOURCE[0]}")"
# shellcheck disable=SC2034
__base="$(basename "${__file}" .sh)"
# shellcheck disable=SC2034
__root="$(cd "$(dirname "${__dir}")" && pwd)" # <-- change this as it depends on your app

NOW=$(date '+%Y%m%d')
TMP=$(mktemp -d)
SRCDIR="$__dir/.."

CODENAME=""
BUILDNAME=$NOW
VERSIONTAG=$(git describe --tags)

cleanup() { rm -rf "$TMP"; }
trap cleanup INT TERM ERR

build_v2() {
  pushd "$SRCDIR"
  LDFLAGS="-s -w -X v2ray.com/core.build=${BUILDNAME} -X v2ray.com/core.version=${VERSIONTAG}"
  if [[ -n "$CODENAME" ]]; then
    LDFLAGS+="-X v2ray.com/core.codename=${CODENAME}"
  fi

  echo ">>> Compile v2ray ..."
  env CGO_ENABLED=0 go build -o "${TMP}/v2ray${EXESUFFIX}" -ldflags "${LDFLAGS}" ./main
  if [[ $GOOS == "windows" ]]; then
    env CGO_ENABLED=0 go build -o "${TMP}/wv2ray${EXESUFFIX}" -ldflags "-H windowsgui ${LDFLAGS}" ./main
  fi

  echo ">>> Compile v2ctl ..."
  env CGO_ENABLED=0 go build -o "${TMP}/v2ctl${EXESUFFIX}" -tags confonly -ldflags "${LDFLAGS}" ./infra/control/main
  popd
}

download_dat() {
  echo ">>> Downloading newest geoip ..."
  wget -O "${SRCDIR}/release/config/geoip.dat" "https://github.com/v2ray/geoip/releases/latest/download/geoip.dat"

  echo ">>> Downloading newest geosite ..."
  wget -O "${SRCDIR}/release/config/geosite.dat" "https://github.com/v2ray/domain-list-community/releases/latest/download/dlc.dat"
}

copyconf() {
  echo ">>> Copying config..."
  pushd "${SRCDIR}/release/config"
  tar c --exclude "*.dat" . | tar x -C "$TMP"
}

copydat() {
  echo ">>> Copying dat..."
  cp "${SRCDIR}/release/config/geosite.dat" "$TMP"
  cp "${SRCDIR}/release/config/geoip.dat" "$TMP"
}

packzip() {
  echo ">>> Generating zip package"
  pushd "$TMP"
  local PKG=${__dir}/v2ray-${GOARCH}-${GOOS}-${PKGSUFFIX}${NOW}.zip
  zip -r "$PKG" .
  echo ">>> Generated: $(basename "$PKG")"
}

packtgz() {
  echo ">>> Generating tgz package"
  pushd "$TMP"
  local PKG=${__dir}/v2ray-${GOARCH}-${GOOS}-${PKGSUFFIX}${NOW}.tar.gz
  tar cvfz "$PKG" .
  echo ">>> Generated: $(basename "$PKG")"
}

packtgzAbPath() {
  local ABPATH="$1"
  echo ">>> Generating tgz package at $ABPATH"
  pushd "$TMP"
  tar cvfz "$ABPATH" .
  echo ">>> Generated: $ABPATH"
}

pkg=zip
nodat=0
noconf=0
GOOS=linux
GOARCH=amd64
EXESUFFIX=
PKGSUFFIX=

for arg in "$@"; do
  case $arg in
  arm*)
    GOARCH=$arg
    ;;
  mips*)
    GOARCH=$arg
    ;;
  386)
    GOARCH=386
    ;;
  windows)
    GOOS=windows
    EXESUFFIX=.exe
    ;;
  darwin)
    GOOS=$arg
    ;;
  nodat)
    nodat=1
    PKGSUFFIX=${PKGSUFFIX}nodat-
    ;;
  noconf)
    noconf=1
    ;;
  tgz)
    pkg=tgz
    ;;
  abpathtgz=*)
    pkg=${arg##abpathtgz=}
    ;;
  codename=*)
    CODENAME=${arg##codename=}
    ;;
  buildname=*)
    BUILDNAME=${arg##buildname=}
    ;;
  esac
done

export GOOS GOARCH
echo "Build ARGS: GOOS=${GOOS} GOARCH=${GOARCH} CODENAME=${CODENAME} BUILDNAME=${BUILDNAME}"
echo "PKG ARGS: pkg=${pkg}"
build_v2

if [[ $nodat != 1 ]]; then
  if [[ ! -f "${SRCDIR}/release/config/geosite.dat" ]] ||
    [[ ! -f "${SRCDIR}/release/config/geoip.dat" ]]; then
    download_dat
  fi
  copydat
fi

if [[ $noconf != 1 ]]; then
  copyconf
fi

if [[ $pkg == "zip" ]]; then
  packzip
elif [[ $pkg == "tgz" ]]; then
  packtgz
else
  packtgzAbPath "$pkg"
fi

cleanup
