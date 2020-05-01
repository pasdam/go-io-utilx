package ioutilx_test

import (
	"os"
	"strings"
	"testing"

	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/pasdam/go-io-utilx/pkg/ioutilx"
	"github.com/stretchr/testify/assert"
)

func Test_ReaderToFile_Fail_ShouldReturnErrorIfItCannotCreateParents(t *testing.T) {
	dstPath := "reader_to_file_test.go/parent_is_a_file"
	reader := strings.NewReader("reader-string")

	err := ioutilx.ReaderToFile(reader, dstPath)

	assert.NotNil(t, err)
}

func Test_ReaderToFile_Fail_ShouldReturnErrorIfItCannotCreateFile(t *testing.T) {
	tmpDir := filestest.TempDir(t)
	defer os.RemoveAll(tmpDir)
	dstPath := tmpDir + ""
	expectedContent := "reader-string"
	reader := strings.NewReader(expectedContent)

	err := ioutilx.ReaderToFile(reader, dstPath)

	assert.NotNil(t, err)
}

func Test_ReaderToFile_Success_ShouldCopyFile(t *testing.T) {
	tmpDir := filestest.TempDir(t)
	defer os.RemoveAll(tmpDir)
	dstPath := tmpDir + "out_file"
	expectedContent := "reader-string"
	reader := strings.NewReader(expectedContent)

	err := ioutilx.ReaderToFile(reader, dstPath)

	assert.Nil(t, err)
	filestest.FileExistsWithContent(t, dstPath, expectedContent)
}
