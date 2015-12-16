package stack

import (
   "errors"
   "fmt"
   "strings"
   "time"
   "net"
)

const (
   maximumEngineLimit   = 5
)

type StatMetaInfo struct {
   statID         string
   sequence       uint64
   updated        bool
   activeNodes    Addr[]
   statLapse      time.Duration
   usedTime       time.Duration
   engineCount    int
   engineID       string
   latestEngine   string
   engineList     Addr[]
}

var engine = statEngine.newStatEngine()

func NewStatMeta() *StatMetaInfo {
   sMeta := &StatMetaInfo{}
   sMeta.makeDefaults();
   return sMeta
}

func (s *StatMetaInfo) makeDefaults() {
   s.StatMetaInfo.makeDefaults()
   s.statID = ""
   s.sequence = 0
   s.updated = false
   s.statLapse = time.Now()
   s.engineCount = 0
}


func (s *StatMetaInfo) bootstrapStatEngine() error {
   
   if s.countEngines() >= maximumEngineLimit {
      fmt.Printf("Can't initiate db engine, alreay reached limit : %d\n", s.countEngines())

   }
   string drvName := "mysql" // by defaults
   string dsrc := "user:password@tcp(127.0.0.1:3307/hello" // example strings
   
   dbEng, err := s.accessEngine(drvName, dsrc)
   if err != nil {
      log.Fatal(err)
   }

}

// do priority queue here?
func (s *StatMetaInfo) redirectOther() (string, error) {
   if s.latestEngine != nil {
      return s.latestEngine, nil
   } else {
      zeroTime := time.Time{}
      candidate := ""
   }
}

func (s *StatMetaInfo) selectHotEngine() (string, error) {
   zeroTime := time.Time{}
   prevTime := zeroTime
   for i := 0; i < s.countEngines(); i++ {
      memoTime := (float)zeroTime
      if engineList[i].string() == s.engineID {
         curTime := time.Now()
         diffTime := curTime.Sub(s.usedTime)
         prevTime = memoTime
         memoTime = Min(diffTime, memoTime)
         if prevTime > memoTime {
            s.Swap(0, i)
         }
      }
   }

}

func (s *StatMetaInfo) Swap(i, j int) {
   s.engineList[i], s.engineList[j] = s.engineList[j], s.engineList[i]
}


func (s *StatMetaInfo) countEngines() (int, error) {
   return len(s.engineList) - 1
}


// statDataAvailable checks if there are new network stat data to fetch.
// To consider it as fetchable, stat data has more than one attributes
// that is updated from the very last stat data.
// [i.e.] http://go-database-sql.org/overview.html
func (ctx *Context) statDataAvailable() (bool, error) {

}