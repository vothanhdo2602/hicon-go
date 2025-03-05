package orm

import (
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"github.com/vothanhdo2602/hicon/external/config"
)

type ModelBuilder struct {
	tableConfig *config.TableConfiguration
	definition  dynamicstruct.Builder
}
