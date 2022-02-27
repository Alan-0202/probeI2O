package list

import (
	"I2Oprobe/internal/g"
	"fmt"
	"testing"
)

func TestSourFile_Read(t *testing.T) {
	g.ConfigParse()
	fmt.Println(*g.SourcePath)
	NewSourFile(*g.SourcePath).Read()


}
