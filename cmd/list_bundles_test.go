package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"os"

	. "github.com/onsi/ginkgo/v2/dsl/core"
	. "github.com/onsi/gomega"
)

var _ = Describe("List Bundles Cmd", func() {
	When("Calling opcap list bundles", func() {
		It("should succeed", func() {
			_, err := executeCommand(listBundlesCmd(), createFakeRepo(os.Args[1]))
			Expect(err).ToNot(HaveOccurred())
		})
		It("should succeed", func() {
			_, err := executeCommand(listBundlesCmd(), []string{"--from-dir=/Users/mkong/Projects/opcap/internal/bundle/data.test/certified-operators"}...)
			Expect(err).ToNot(HaveOccurred())
		})
		deleteFakeRepo(os.Args[1])
	})
})

func createFakeRepo(target string) string {
	cmd, err := exec.Command("/bin/sh", "/Users/mkong/Projects/opcap/internal/bundle/test-data.sh", target).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}

	flag := string(cmd)

	return flag
}

func deleteFakeRepo(target string) string {
	_, err := exec.Command("/bin/sh", "/Users/mkong/Projects/opcap/internal/bundle/delete_repo.sh", target).Output()
	if err != nil {
		log.Fatal(err)
	}
	return "fake repo deleted"
}
