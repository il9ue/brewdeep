package stack

import (
   "errors"
   "fmt"
   "strings"
   "time"
   "net"

   "database"
   "database/sql"
)

type StatEngine struct {

}

func newStatEngine() *StatEngine {
   sEng := &StatEngine{}
   sEng.makeDefaults()
   return sEng
}

func (s *StatEngine) accessEngine(ddName string, dsName string) (*DB, error) {
   db, err := sql.Open(ddName, dsName)
   if err != nil {
      log.Fatal(err)
   }
   defer db.Close()
}


