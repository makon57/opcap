package cmd

import (
	. "github.com/onsi/ginkgo/v2/dsl/core"
	. "github.com/onsi/gomega"
)

var _ = Describe("List Bundles Cmd", func() {

	// create testdata directory place 'acc-operator' data in and 'git init' directory to great git repo for testing

	When("Calling opcap list bundles", func() {
		It("should succeed", func() {
			_, err := executeCommand(listBundlesCmd(), []string{"--bundles-path=internal/bundle/.test/certified-operators", "--bundles-repo=https://github.com/redhat-openshift-ecosystem/certified-operators.git"}...)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should succeed", func() {
			_, err := executeCommand(listBundlesCmd(), []string{"--bundles-path=internal/bundle/.test/marketplace-operators", "--bundles-repo=https://github.com/redhat-openshift-ecosystem/redhat-marketplace-operators.git"}...)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
