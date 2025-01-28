package version

/*
===================
=    IMPORTANT    =
===================

This package is solely for versioning information when releasing Vulerian..

Changing these values will change the versioning information when releasing Vulerian.
*/

var (
	// VERSION  is version number that will be displayed when running ./Vulerian version
	VERSION = "v3.16.1"

	// GITCOMMIT is hash of the commit that will be displayed when running ./Vulerian version
	// this will be overwritten when running  build like this: go build -ldflags="-X github\.com/danielpickens/Vulerian/cmd.GITCOMMIT=$(GITCOMMIT)"
	// HEAD is default indicating that this was not set during build
	GITCOMMIT = "HEAD"
)
