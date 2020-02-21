package gunzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Gunzip struct {
	Src, Dest string
	D, L      bool
}

func (g *Gunzip) Unzip() error {
	// check zip source path.
	if filepath.Ext(g.Src) != ".zip" {
		return fmt.Errorf("invalid argment: %s is not zip file.", g.Src)
	}

	// check destination option.
	// option false is decompress to current directory.
	dest := ""
	if g.D {
		dest = g.Dest
	} else {
		curdir, err := os.Getwd()
		if err != nil {
			return err
		}
		dest = curdir
	}

	fmt.Printf("filepath: %v", dest)
	// open zip.
	r, err := zip.OpenReader(g.Src)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", err)
	}
	defer r.Close()

	// check one file
	for _, f := range r.File {

		fmt.Printf("file name: %v", f.Name)
		//
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		fmt.Printf("filepath: %v", filepath.Dir(fpath))

		// if f is directory, go to next for roop.
		if f.FileInfo().IsDir() {
			continue
		}

		// check destination directory is exist.
		// not exist, make directory.
		_, err := os.Stat(filepath.Dir(fpath))
		if os.IsNotExist(err) {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}
		}

		// create and open file in destination directory.
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		// open file in zip.
		rc, err := f.Open()
		if err != nil {
			return err
		}

		// copy
		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}

		outFile.Close()
		rc.Close()
	}

	return nil
}
