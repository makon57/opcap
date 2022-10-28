package bundles

type bundle struct {
	packageName string
	channel     string
	version     string
	startingCsv string
	ocpVersions string
}

var listedBundles []bundle

func ListBundles() ([]bundle, error) {
	err := GitCloneOrPullBundles("https://github.com/redhat-openshift-ecosystem/certified-operators.git", "internal/bundles/test-data")
	if err != nil {
		return nil, err
	}

	bundle := bundle{}
	if err := bundle.List("internal/bundles/test-data/operators"); err != nil {
		return nil, err
	}
	return listedBundles, nil
}
