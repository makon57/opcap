package bundles

import (
	// "fmt"
	"os"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	// operatorv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	"gopkg.in/yaml.v2"
)

func GitCloneOrPullBundles(URL string, dir string) error {
	// Clone bundle repository, main branch only (should be the default for certified-operators)
	if _, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      URL,
		Progress: os.Stdout,
	}); err != nil && err != git.ErrRepositoryAlreadyExists {
		return err
	}
	// Pull data if repo already exists
	repo, err := git.PlainOpen("internal/bundles/test-data")
	if err != nil {
		return err
	}
	workTree, err := repo.Worktree()
	if err != nil {
		return err
	}
	err = workTree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}
	return nil
}

func ReadDirectory(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dir, err := file.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	var directories []string
	for _, entry := range dir {
		if entry != "ci.yaml" && entry != "README.md" {
			directories = append(directories, entry)
		}
	}

	return directories, nil
}

func (b *bundle) List(repoPath string) error { // TODO: return []bundle in list method

	operators, err := ReadDirectory(repoPath)
	if err != nil {
		return err
	}

	for _, operator := range operators {
		path := strings.Join([]string{repoPath, operator}, "/")
		versions, err := ReadDirectory(path)
		if err != nil {
			return err
		}

		for _, version := range versions {
			b.version = version
			versionDir := strings.Join([]string{path, version}, "/")
			data, err := ReadDirectory(versionDir)
			if err != nil {
				return err
			}

			for _, d := range data {
				if d == "manifests" {
					// get startingCSV name
					var csvFile string
					files, err := os.ReadDir(versionDir + "/manifests")
					if err != nil {
						return err
					}

					for _, file := range files {
						if strings.Contains(file.Name(), "clusterserviceversion") {
							csvFile = file.Name()
						}
					}

					csvPath := strings.Join([]string{versionDir, "manifests", csvFile}, "/")
					b.getStartingCsv(csvPath)
				}

				if d == "metadata" {
					// get package, channel and ocpVersions
					annotationsPath := strings.Join([]string{versionDir, "metadata", "annotations.yaml"}, "/")
					b.getAnnotations(annotationsPath)
				}

			}
			listedBundles = append(listedBundles, *b)
			// fmt.Printf("packageName: %s, channel: %s, version: %s, startingCsv: %s, ocpVersions: %s\n", b.packageName, b.channel, b.version, b.startingCsv, b.ocpVersions)
		}
	}
	return nil
}

func (b *bundle) getStartingCsv(filePath string) error {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// csv := &operatorv1alpha1.ClusterServiceVersion{}
	csv := make(map[string]map[string]string)
	_ = yaml.Unmarshal(f, csv)
	// TODO: handle error here, unmarshalls name but not the whole csv
	// if err != nil {
	// 	return err
	// }

	b.startingCsv = csv["metadata"]["name"]

	return nil
}

func (b *bundle) getAnnotations(filePath string) error {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	data := make(map[string]map[string]string)
	err = yaml.Unmarshal(f, &data)
	if err != nil {
		return err
	}

	// TODO: validate if annotations file has the correct map fields and is well formed.
	b.ocpVersions = data["annotations"]["com.redhat.openshift.versions"]
	b.channel = data["annotations"]["operators.operatorframework.io.bundle.channel.default.v1"]
	b.packageName = data["annotations"]["operators.operatorframework.io.bundle.package.v1"]

	return nil
}
