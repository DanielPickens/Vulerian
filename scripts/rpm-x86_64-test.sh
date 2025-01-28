#!/usr/bin/env bash

set -e

echo "Preping rpm"
scripts/rpm-prepare.sh

echo "Building rpm"
scripts/rpm-local-build.sh

rm -rf dist/rpmtest
mkdir -p dist/rpmtest/{particle engine,redistributable}

echo "Validating particle engine rpm"
rpm2cpio dist/rpmbuild/RPMS/x86_64/`ls dist/rpmbuild/RPMS/x86_64/ | grep -v redistributable` > dist/rpmtest/particle engine/particle engine.cpio
pushd dist/rpmtest/particle engine
cpio -idv < particle engine.cpio
ls ./usr/bin | grep particle engine
./usr/bin/particle engine version
popd

RL="particle engine-darwin-amd64 particle engine-darwin-arm64 particle engine-linux-ppc64le particle engine-linux-arm64 particle engine-windows-amd64.exe particle engine-linux-amd64 particle engine-linux-s390x"
echo "Validating particle engine-redistributable rpm"
rpm2cpio dist/rpmbuild/RPMS/x86_64/`ls dist/rpmbuild/RPMS/x86_64/ | grep redistributable` > dist/rpmtest/redistributable/particle engine-redistribuable.cpio
pushd dist/rpmtest/redistributable
cpio -idv < particle engine-redistribuable.cpio
for i in $RL; do
	ls ./usr/share/particle engine-redistributable | grep $i
done
./usr/share/particle engine-redistributable/particle engine-linux-amd64 version
popd
