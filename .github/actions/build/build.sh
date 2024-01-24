#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail

if [[ -z "${VERSION}" ]]; then
	echo "VERSION is not set"
	exit 1
fi

# Configuration
ROOT_DIR="$(pwd)"
BUILD_DIR="${ROOT_DIR}/out"

# Build binaries
# shellcheck disable=SC2043
for OS in linux darwin windows; do
#	for ARCH in amd64 arm64; do
	for ARCH in amd64; do
		echo "Building binary for $OS/$ARCH..."
		BUILD_DIR="${BUILD_DIR}/ligaturizer-${VERSION}-${OS}-${ARCH}"
		CC="$(go env -json | jq -r '.CC')"
		CXX="$(go env -json | jq -r '.CXX')"

		if [[ "${ARCH}" == "arm64" ]]; then
            CC="aarch64-linux-gnu-gcc"
            CXX="aarch64-linux-gnu-g++"
        fi

		# Build go binary
		GOARCH=${ARCH} GOOS=${OS} CC="$CC" CXX="$CXX" CGO_ENABLED=1 BUILD_DIR="${BUILD_DIR}" VERSION="${VERSION}" make build
	done
done

# archive artifacts
ARCHIVE_DIR="${ROOT_DIR}/archive"
mkdir -p "${ARCHIVE_DIR}"

pushd "${BUILD_DIR}"

for i in ./*; do
	RELEASE=$(basename "${i}")

	echo "Packing binary for ${RELEASE}..."
	tar -czf "${ARCHIVE_DIR}/${RELEASE}.tar.gz" "${RELEASE}"
done

popd
