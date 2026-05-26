package parser

import (
	"io/fs"
	"regexp"

	"github.com/cucumber/godog/internal/flags"
	"github.com/cucumber/godog/internal/models"
)

var pathLineRe = regexp.MustCompile(`:([\d]+)$`)

// ExtractFeaturePathLine ...
func ExtractFeaturePathLine(p string) (string, int) { _ = "STUB: not implemented"; return "", 0 }

func parseFeatureFile(fsys fs.FS, path, dialect string, newIDFunc func() string) (*models.Feature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBytes(path string, feature []byte, dialect string, newIDFunc func() string) (*models.Feature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFeatureDir(fsys fs.FS, dir, dialect string, newIDFunc func() string) ([]*models.Feature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parsePath(fsys fs.FS, path, dialect string, newIDFunc func() string) ([]*models.Feature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter scenario by line number

// ParseFeatures ...
func ParseFeatures(fsys fs.FS, filter, dialect string, paths []string) ([]*models.Feature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FeatureContent = flags.Feature

func ParseFromBytes(filter, dialect string, featuresInputs []FeatureContent) ([]*models.Feature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterFeatures(filter string, features []*models.Feature) (result []*models.Feature) {
	_ = "STUB: not implemented"
	return nil
}
