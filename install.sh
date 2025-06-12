#!/bin/sh
set -e

REPO="ChristoferBerruz/aplpdown"
VERSION=${1:-latest}
OS=$(uname | awk '{print toupper(substr($0,1,1)) tolower(substr($0,2))}')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
  ARCH="x86_64"
elif [ "$ARCH" = "amd64" ]; then
  ARCH="x86_64"
elif [ "$ARCH" = "aarch64" ]; then
  ARCH="arm64"
fi

BINARY="aplpdown"

if [ "$VERSION" = "latest" ]; then
  VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep tag_name | cut -d '"' -f 4)
fi

URL="https://github.com/$REPO/releases/download/$VERSION/${BINARY}_${OS}_${ARCH}.tar.gz"

echo "Installing $BINARY $VERSION for $OS/$ARCH..."
echo "Downloading from $URL"
curl -L $URL | tar -xz
chmod +x $BINARY
sudo mv $BINARY /usr/local/bin/

echo "$BINARY installed!"
