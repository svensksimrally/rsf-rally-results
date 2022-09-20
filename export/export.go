package export

import (
	"os"
	"path/filepath"

	"github.com/caarlos0/log"
	"github.com/go-pkgz/fileutils"
	"github.com/spf13/viper"
)

// Do the export of the series standings and results into html.
func Do() error {
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	log.WithField("tempdir", tempDir).Debug("Created temporary folder")

	// TODO: fetch clubs and generate the templates
	if err := commit(viper.GetString("output"), tempDir); err != nil {
		return err
	}

	// TODO: for now just copy the template

	log.Info("All done!")
	return nil
}

func commit(outputDir, inputDir string) error {
	log.WithField("from", inputDir).WithField("to", outputDir).Debug("Committing files to output directory")

	log.WithField("from", inputDir).Debug("Cleaning up existing output directory")
	if err := os.RemoveAll(outputDir); err != nil {
		return err
	}
	// TODO: copy assets

	// TODO: temporary just copying the index file
	if err := fileutils.CopyFile("templates/index.html", filepath.Join(outputDir, "index.html")); err != nil {
		return nil
	}

	return fileutils.CopyDir(inputDir, outputDir)
}
