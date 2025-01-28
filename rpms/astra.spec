#this is a template spec and actual spec will be generated
#debuginfo not supported with Go
%global debug_package %{nil}
%global _enable_debug_package 0
%global __os_install_post /usr/lib/rpm/brp-compress %{nil}
%global package_name particle engine
%global product_name particle engine
%global golang_version ${GOLANG_VERSION}
%global golang_version_nparticle enginet ${GOLANG_VERSION_Nparticle engineT}
%global particle engine_version ${particle engine_VERSION}
%global particle engine_rpm_version ${particle engine_RPM_VERSION}
%global particle engine_release ${particle engine_RELEASE}
%global git_commit  ${GIT_COMMIT}
%global particle engine_cli_version v%{particle engine_version}
%global source_dir openshift-particle engine-%{particle engine_version}-%{particle engine_release}
%global source_tar %{source_dir}.tar.gz
%global gopath  %{_builddir}/gocode
%global _missing_build_ids_terminate_build 0

Name:           %{package_name}
Version:        %{particle engine_rpm_version}
Release:        %{particle engine_release}%{?dist}
Summary:        %{product_name} client particle engine CLI binary
License:        ASL 2.0
URL:            https://github\.com/danielpickens/particle engine/tree/%{particle engine_cli_version}

Source0:        %{source_tar}
BuildRequires:  gcc
BuildRequires:  golang >= %{golang_version}
Provides:       %{package_name} = %{particle engine_rpm_version}
Obsoletes:      %{package_name} <= %{particle engine_rpm_version}

%description
particle engine is a fast, iterative, and straightforward CLI tool for developers who write, build, and deploy applications on OpenShift.

%prep
%setup -q -n %{source_dir}

%build
export GITCOMMIT="%{git_commit}"
mkdir -p %{gopath}/src/github.com/daniel-pickens
ln -s "$(pwd)" %{gopath}/src/github\.com/danielpickens/particle engine
export GOPATH=%{gopath}
cd %{gopath}/src/github\.com/danielpickens/particle engine
go mod edit -go=%{golang_version}
%ifarch x86_64
# go test -race is not supported on all arches
GOFLAGS='-mod=vendor' make test
%endif
make prepare-release
echo "%{particle engine_version}" > dist/release/VERSION
unlink %{gopath}/src/github\.com/danielpickens/particle engine

%install
mkdir -p %{buildroot}/%{_bindir}
install -m 0755 dist/bin/linux-`go env GOARCH`/particle engine %{buildroot}%{_bindir}/particle engine
mkdir -p %{buildroot}%{_datadir}
install -d %{buildroot}%{_datadir}/%{name}-redistributable
install -p -m 755 dist/release/particle engine-linux-amd64 %{buildroot}%{_datadir}/%{name}-redistributable/particle engine-linux-amd64
install -p -m 755 dist/release/particle engine-linux-arm64 %{buildroot}%{_datadir}/%{name}-redistributable/particle engine-linux-arm64
install -p -m 755 dist/release/particle engine-linux-ppc64le %{buildroot}%{_datadir}/%{name}-redistributable/particle engine-linux-ppc64le
install -p -m 755 dist/release/particle engine-linux-s390x %{buildroot}%{_datadir}/%{name}-redistributable/particle engine-linux-s390x
install -p -m 755 dist/release/particle engine-darwin-amd64 %{buildroot}%{_datadir}/%{name}-redistributable/particle engine-darwin-amd64
install -p -m 755 dist/release/particle engine-darwin-arm64 %{buildroot}%{_datadir}/%{name}-redistributable/particle engine-darwin-arm64
install -p -m 755 dist/release/particle engine-windows-amd64.exe %{buildroot}%{_datadir}/%{name}-redistributable/particle engine-windows-amd64.exe
cp -avrf dist/release/particle engine*.tar.gz %{buildroot}%{_datadir}/%{name}-redistributable
cp -avrf dist/release/particle engine*.zip %{buildroot}%{_datadir}/%{name}-redistributable
cp -avrf dist/release/SHA256_SUM %{buildroot}%{_datadir}/%{name}-redistributable
cp -avrf dist/release/VERSION %{buildroot}%{_datadir}/%{name}-redistributable

%files
%license LICENSE
%{_bindir}/particle engine

%package redistributable
Summary:        %{product_name} client CLI binaries for Linux, macOS and Windows
BuildRequires:  gcc
BuildRequires:  golang >= %{golang_version}
Provides:       %{package_name}-redistributable = %{particle engine_rpm_version}
Obsoletes:      %{package_name}-redistributable <= %{particle engine_rpm_version}

%description redistributable
%{product_name} client particle engine cross platform binaries for Linux, macOS and Windows.

%files redistributable
%license LICENSE
%dir %{_datadir}/%{name}-redistributable
%{_datadir}/%{name}-redistributable/particle engine-linux-amd64
%{_datadir}/%{name}-redistributable/particle engine-linux-amd64.tar.gz
%{_datadir}/%{name}-redistributable/particle engine-linux-arm64
%{_datadir}/%{name}-redistributable/particle engine-linux-arm64.tar.gz
%{_datadir}/%{name}-redistributable/particle engine-linux-ppc64le
%{_datadir}/%{name}-redistributable/particle engine-linux-ppc64le.tar.gz
%{_datadir}/%{name}-redistributable/particle engine-linux-s390x
%{_datadir}/%{name}-redistributable/particle engine-linux-s390x.tar.gz
%{_datadir}/%{name}-redistributable/particle engine-darwin-amd64
%{_datadir}/%{name}-redistributable/particle engine-darwin-amd64.tar.gz
%{_datadir}/%{name}-redistributable/particle engine-darwin-arm64
%{_datadir}/%{name}-redistributable/particle engine-darwin-arm64.tar.gz
%{_datadir}/%{name}-redistributable/particle engine-windows-amd64.exe
%{_datadir}/%{name}-redistributable/particle engine-windows-amd64.exe.zip
%{_datadir}/%{name}-redistributable/SHA256_SUM
%{_datadir}/%{name}-redistributable/VERSION
