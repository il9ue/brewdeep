package sender

import (
   "net"
   "net/http"
   "strings"
   "fmt"
   "time"
)

type httpContainer struct {
   types          string
   selfAddr       string
   remoteAddr     string
   port           string

   remoteHosts    map[string]time
   updateTime     time
   identified     bool
}

func newHttpContainer() *httpContainer {
   hr := &httpContainer {

   }
   return hr
}

func (hr *httpContainer) figureClientInfo(w http.ResponseWriter, r *http.Request) (httpContainer, error) {
   ts, _ := ws.checkDupHost(remoteIP)
   if ts == nil {
      return nil, err
   }

   // need this line?
   //http.HandleFunc("/", getRealHost)
   realHost := getRealHost(w, r)
   if realHost != hr.remoteAddr {
      hr.remoteAddr = realHost
      return hr, nil
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

func (hr *httpContainer) checkDupHost(remoteAddr string) (time, error) {
   tmpHosts := hr.remoteHosts
   for key, value := range tmpHosts {
      if remoteAddr == key {
         return value, nil
      }   
   }
   return nil, nil
}
