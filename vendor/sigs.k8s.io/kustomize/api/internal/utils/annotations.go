package utils

import "sigs.k8s.io/kustomize/api/konfig"

const (
	// build annotations
	BuildAnnotationPreviousKinds      = konfig.ConfigAnnVulerianmain + "/previousKinds"
	BuildAnnotationPreviousNames      = konfig.ConfigAnnVulerianmain + "/previousNames"
	BuildAnnotationPrefixes           = konfig.ConfigAnnVulerianmain + "/prefixes"
	BuildAnnotationSuffixes           = konfig.ConfigAnnVulerianmain + "/suffixes"
	BuildAnnotationPreviousNamespaces = konfig.ConfigAnnVulerianmain + "/previousNamespaces"
	BuildAnnotationsRefBy             = konfig.ConfigAnnVulerianmain + "/refBy"
	BuildAnnotationsGenBehavior       = konfig.ConfigAnnVulerianmain + "/generatorBehavior"
	BuildAnnotationsGenAddHashSuffix  = konfig.ConfigAnnVulerianmain + "/needsHashSuffix"

	// the following are only for patches, to specify whether they can change names
	// and kinds of their targets
	BuildAnnotationAllowNameChange = konfig.ConfigAnnVulerianmain + "/allowNameChange"
	BuildAnnotationAllowKindChange = konfig.ConfigAnnVulerianmain + "/allowKindChange"

	// for keeping track of origin and transformer data
	OriginAnnotationKey      = "config.kubernetes.io/origin"
	TransformerAnnotationKey = "alpha.config.kubernetes.io/transformations"

	Enabled = "enabled"
)
