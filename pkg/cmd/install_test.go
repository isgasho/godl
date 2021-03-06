package cmd

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

type testGzUnArchiver struct{}

func (testGzUnArchiver) Unarchive(source, target string) error { return nil }

type fakeRemover struct{}

func (fr fakeRemover) RemoveAll(path string) error { return nil }

func TestInstallRelease(t *testing.T) {
	testClient := newTestClient(func(req *http.Request) *http.Response {
		testData := bytes.NewBufferString("This is test data")

		return &http.Response{
			StatusCode:    http.StatusOK,
			Body:          ioutil.NopCloser(testData),
			ContentLength: int64(len(testData.Bytes())),
		}
	})

	failingTestClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(bytes.NewBufferString("")),
		}
	})

	tests := map[string]struct {
		c                 *http.Client
		downloadedVersion string
		installVersion    string
		success           bool
	}{
		"installRelease downloads from remote when version not found locally": {
			testClient, "1.10.1", "1.11.7", true,
		},
		"installRelease installs local downloaded version": {testClient, "1.10.6", "1.10.6", true},
		"installRelease handle error when fetching binary from remote": {
			failingTestClient, "1.10.1", "1.11.9", false,
		},
	}

	tmpDir, err := createTempGodlDownloadDir()
	if err != nil {
		t.Fatalf("TestInstallRelease failed: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tmpFile, err := createTempGoBinaryArchive(tmpDir, tc.downloadedVersion)
			defer tmpFile.Close()

			ua := testGzUnArchiver{}
			fr := fakeRemover{}
			ic := &inMemoryFileCreatorRenamer{}
			dl := &goBinaryDownloader{
				baseURL:     "https://storage.googleapis.com/golang/",
				client:      tc.c,
				downloadDir: ".",
				fCR:         ic,
				genHash:     genTestHash,
				verifyHash:  fakeVerifyHash,
			}

			err = installRelease(tc.installVersion, tmpDir, ua, fr, dl)
			var got bool
			if err != nil {
				got = false
			} else {
				got = true
			}

			if got != tc.success {
				t.Errorf("Error installing go binary: %v", err)
			}
		})
	}
}

func TestInstallCmdCalledWithNoArgs(t *testing.T) {
	_, err := executeCommand(rootCmd, "install")
	expected := "provide binary archive version to install"
	got := err.Error()
	if got != expected {
		t.Errorf("godl install Unknown error: %v", err)
	}
}

func TestInstallCommandHelp(t *testing.T) {
	_, err := executeCommand(rootCmd, "install", "-h")
	if err != nil {
		t.Errorf("godl install failed: %v", err)
	}
}
