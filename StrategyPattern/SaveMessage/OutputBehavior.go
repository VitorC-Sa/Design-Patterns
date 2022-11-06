package savemessage

import (
	"fmt"
	"log"
	"os"
)

type OutputBehavior interface {
	Save(string)
}

type Log struct{}
type FmtPrintln struct{}
type TxtFile struct{}
type Postgres struct{}
type MySql struct{}

func (Log) Save(msg string)        { log.Println(msg) }
func (FmtPrintln) Save(msg string) { fmt.Println(msg) }
func (TxtFile) Save(msg string) {
	f, _ := os.Create("output.txt")
	defer f.Close()
	f.WriteString(msg)
}
func (Postgres) Save(msg string) {}
func (MySql) Save(msg string)    {}
