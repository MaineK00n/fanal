package binary

import (
	"os"

	"golang.org/x/xerrors"

	"github.com/aquasecurity/fanal/analyzer"
	"github.com/aquasecurity/fanal/analyzer/language"
	"github.com/aquasecurity/fanal/types"
	"github.com/aquasecurity/go-dep-parser/pkg/golang/binary"
)

func init() {
	analyzer.RegisterAnalyzer(&gobinaryLibraryAnalyzer{})
}

const version = 1

type gobinaryLibraryAnalyzer struct{}

func (a gobinaryLibraryAnalyzer) Analyze(target analyzer.AnalysisTarget) (*analyzer.AnalysisResult, error) {
	res, err := language.Analyze(types.GoBinary, target.FilePath, target.Content, binary.Parse)
	if err != nil {
		return nil, xerrors.Errorf("unable to parse %s: %w", target.FilePath, err)
	}
	return res, nil
}

func (a gobinaryLibraryAnalyzer) Required(_ string, fileInfo os.FileInfo) bool {
	// Check executable file
	if fileInfo.Mode()&0111 != 0 {
		return true
	}
	return false
}

func (a gobinaryLibraryAnalyzer) Type() analyzer.Type {
	return analyzer.TypeGoBinary
}

func (a gobinaryLibraryAnalyzer) Version() int {
	return version
}
