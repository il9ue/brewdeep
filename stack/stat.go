package stack

import (
   "errors"
   "fmt"
   "strings"
   "time"
   "net"
   "sync"
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

   engineList     chan Addr[]
   nodeList       chan Addr[]
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
            break
         }
      }
   }
   return s.engineList[0], nil
}


func (s *StatMetaInfo) mapNodes() <-chan Addr {
   
   // if stat Track nodes are more than just local one,
   // start publish/subscribe model and map network nodes 
   // with stat db engine engine
   if len(s.nodeList) > 1 {
      
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
func (s *StatMetaInfo) statDataAvailable() (bool, error) {
   
}


// mergeStats do concurrency routine to compute updated node statistics value 
func (s *StatMetaInfo) mergeStats(done <-chan struct{}, cs ...<-chan int) <-chan int {
   var wg sync.WaitGroup
   out := make(chan int)

   // Start an output goroutine for each input channel in cs.  output
   // copies values from c to out until c is closed or it receives a value
   // from done, then output calls wg.Done.
   output := func(c <-chan int) {
      for m := range c {
         select {
         case out <- m:
         case <-done:
         }
      }
      wg.Done()
   }
   wr.Add(len(cs))

   // Start a goroutine to close out once all the output goroutines are
   // done.  This must start after the wg.Add call.
   // fetch some stat value by call related functions
   go func () {
      wg.Wait()
      close(out)
   }()
   return out
}

func (s *StatMetaInfo) generateEngine() error {
   out := make(chan int)
   go func() {
      
      close(out)
   }()
   return out
}

// scanEngines lookup active engine or/and nodes and fetch updated values if any..
// return string value for engineID
func (s* StatMetaInfo) scanEngines() (string, error) {
   
}


