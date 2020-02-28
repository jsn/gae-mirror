package main

import (
    "log"
    "strings"
    "net"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
)

func getEnv(evar string, default_ string) string {
    v := os.Getenv(evar)
    if v == "" {
        return default_
    } else {
        return v
    }
}

func makeHandler(upstream string) func (w http.ResponseWriter, r *http.Request){
    url, _ := url.Parse(upstream)
    proxy := httputil.NewSingleHostReverseProxy(url)

    return func (w http.ResponseWriter, r *http.Request) {
	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
            r.Header.Set("X-Real-IP", clientIP)
            if prior, ok := r.Header["X-Forwarded-For"]; ok {
                clientIP = strings.Join(prior, ", ") + ", " + clientIP
            }
            r.Header.Set("X-Forwarded-For", clientIP)
	} else {
            r.Header.Set("X-Real-IP", r.RemoteAddr)
        }
        r.Header.Set("X-Proxy-Version", "gae-mirror/0.0.1")
        // Update the headers to allow for SSL redirection
        r.URL.Host = url.Host
        r.URL.Scheme = url.Scheme
        r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
        r.Host = url.Host
        proxy.ServeHTTP(w, r)
    }
}

func main() {
    port := getEnv("PORT", "8080")
    upstream := getEnv("UPSTREAM", "https://giganda.graniru.org/")
    http.HandleFunc("/", makeHandler(upstream))

    log.Printf("Proxying :%s -> %s", port, upstream)

    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
