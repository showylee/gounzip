package gunzip

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUnzip(t *testing.T) {
	g := Gunzip{}

	dest, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	defer os.RemoveAll(dest)

	t.Run("unzip", func(t *testing.T) {
		g.Src = "./testdata/example1.zip"
		text := "example1.csv"
		g.D = true
		g.Dest = dest
		expected := ("2020/02/20 22:10:33, test, test, false")

		g.Unzip()

		bs, err := ioutil.ReadFile(filepath.Join(dest, text))
		if err != nil {
			t.Errorf("failed: %v", err)
		}

		result := string(bs)

		if strings.TrimSpace(result) != expected {
			t.Errorf("failed: %v != %v", result, expected)
		}
	})

	t.Run("unzip in dir", func(t *testing.T) {
		g.Src = "./testdata/example2.zip"
		text := "dir1/dir2/example2.csv"
		g.D = true
		dest, err := ioutil.TempDir("", "example")
		if err != nil {
			t.Errorf("failed: %v", err)
		}
		defer os.RemoveAll(dest)
		g.Dest = dest

		expected := ("2020/02/20 22:10:33, test, test, false")

		g.Unzip()

		bs, err := ioutil.ReadFile(filepath.Join(dest, text))
		if err != nil {
			t.Errorf("failed: %v", err)
		}

		result := string(bs)

		if strings.TrimSpace(result) != expected {
			t.Errorf("failed: %v != %v", result, expected)
		}
	})

	t.Run("no directory option", func(t *testing.T) {
		g.Src = "./testdata/example1.zip"
		text := "example1.csv"
		g.D = false

		dest, _ := os.Getwd()
		expected := ("2020/02/20 22:10:33, test, test, false")

		g.Unzip()

		bs, err := ioutil.ReadFile(filepath.Join(dest, text))
		if err != nil {
			t.Errorf("failed: %v", err)
		}

		result := string(bs)

		if strings.TrimSpace(result) != expected {
			t.Errorf("failed: %v != %v", result, expected)
		}

		if err := os.Remove(filepath.Join(dest, text)); err != nil {
			t.Errorf("failed: %v", err)
		}
	})
}
