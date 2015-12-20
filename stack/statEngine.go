package stack

import (
   "errors"
   "fmt"
   "strings"
   "time"
   "net"
   "sync"
   "hash/fnv"

   "database"
   "database/sql"
)

type StatEngine struct {
   engineID          uint32
   nodeAddr          Addr
   engineUptime      time.Duration
   engineActive      bool

   lastAccessed      time.Duration
   lastAppended      time.Duration
   timeCreated       time.Duration

   upObservers       chan int  
}

func newStatEngine() *StatEngine {
   s := &StatEngine{}
   s.makeDefaults()
   return s
}

func (s *StatEngine) makeDefaults() {
   s.StatEngine.makeDefaults()
   s.engineID = 0
   s.engineUptime = time.Now()
   s.engineActive = time.Time{}
   s.lastAccessed = time.Now()
   s.lastAppended = time.Time{}
   s.timeCreated = time.Time{}
}

func (s *StatEngine) bootstrapEngine(nodes []Addr) <-chan Addr {

   out := make(chan uint32)

   go func() {
      for _, n := range nodes {
         if s.setEnginebyNode(n) != nil {
            id := s.makeHashID(n.String())
            out <- id
         }
      }
      close(out)
   }()
   return out
}

func (s *StatEngine) setEnginebyNode(a Addr) Addr {
   if s.nodeAddr != nil {
      return s.nodeAddr
   } else {
      s.nodeAddr = a
      return s.nodeAddr
   }
}


func (s *StatEngine) makeHashID(name string) uint32 {
   h := fnv.New32a()
   h.Write([]byte(name))

   if s.engineID > 0 {
      return s.engineID
   }
   s.engineID = h.Sum32()
   return s.engineID
}


func (s *StatEngine) accessEngine(ddName string, dsName string) (*DB, error) {
   db, err := sql.Open(ddName, dsName)
   if err != nil {
      log.Fatal(err)
   }
   defer db.Close()
}

