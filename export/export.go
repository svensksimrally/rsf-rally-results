package export

import (
	"os"
	"path/filepath"

	"github.com/caarlos0/log"
	"github.com/go-pkgz/fileutils"
	"github.com/spf13/viper"
	"github.com/tobier/rsf-rally-results/model"
	"github.com/tobier/rsf-rally-results/site"
)

// Do the export of the series standings and results into html.
func Do() (err error) {
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		return
	}
	defer os.RemoveAll(tempDir)

	log.WithField("tempdir", tempDir).Debug("Created temporary folder")

	m := model.Model{}

	// TODO: fetch series and generate the templates
	if err = site.Render(&m, tempDir); err != nil {
		return
	}

	// TODO: render the front page; what should be on it ?

	if err = commit(viper.GetString("output"), tempDir); err != nil {
		return
	}

	log.Info("All done!")
	return
}

func commit(outputDir, inputDir string) error {
	log.WithField("from", inputDir).WithField("to", outputDir).Debug("Committing files to output directory")

	log.WithField("from", inputDir).Debug("Cleaning up existing output directory")
	if err := os.RemoveAll(outputDir); err != nil {
		return err
	}

	log.Debug("Copying generated files")
	if err := fileutils.CopyDir(inputDir, filepath.Join(outputDir)); err != nil {
		return err
	}

	log.Debug("Copying static assets")
	if err := fileutils.CopyDir("assets", filepath.Join(outputDir, "assets")); err != nil {
		return err
	}

	return fileutils.CopyDir(inputDir, outputDir)
}
