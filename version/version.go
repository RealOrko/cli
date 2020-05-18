package version

import "github.com/blang/semver"

const DefaultVersion = "6.51.0"

var (
	binaryVersion   string
	binarySHA       string
	binaryBuildDate string
)

func VersionString() string {
	versionString, err := semver.Make(binaryVersion)
	if err != nil {
		versionString = semver.MustParse(DefaultVersion)
	}

	metaData := []string{}
	if binarySHA != "" {
		metaData = append(metaData, binarySHA)
	}

	if binaryBuildDate != "" {
		metaData = append(metaData, binaryBuildDate)
	}

	versionString.Build = metaData

	return versionString.String()
}
