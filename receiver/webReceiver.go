package sender

import (
   "net"
   "net/http"
   "strings"
   "fmt"
   "time"
)

type webReceiver struct {
   types          string
   selfAddr       string
   remoteAddr     string
   port           string

   remoteHosts    map[string]time
   updateTime     time
   identified     bool
}

func newWebReceiver() *webReceiver {
   wr := &webReceiver {

   }
   return wr
}

func (wr *webReceiver) figureClientInfo(w http.ResponseWriter, r *http.Request) (webReceiver, error) {
   ts, _ := ws.checkDupHost(remoteIP)
   if ts == nil {
      return nil, err
   }

   // need this line?
   //http.HandleFunc("/", getRealHost)
   realHost := getRealHost(w, r)
   if realHost != wr.remoteAddr {
      wr.remoteAddr = realHost
      return wr, nil
   }
   return nil, nil
}

func getRealHost(w http.ResponseWriter, r *http.Request) string {
   ip, _, _ := net.SplitHostPort(r.RemoteAddr)
   w.RemoteAddr = r.Header.Get("X-FORWARDED-FOR")
   if w.RemoteAddr != nil {
      return w.RemoteAddr
   }
   return nil
}

func (wr *webReceiver) checkDupHost(remoteAddr string) (time, error) {
   tmpHosts := wr.remoteHosts
   for key, value := range tmpHosts {
      if remoteAddr == key {
         return value, nil
      }   
   }
   return nil, nil
}
