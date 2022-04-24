package main

import (
   "log"
   "os"
   "io/ioutil"
   "regexp"
   "strconv"
   "math/big"
   "net/http"
)

func calculate(p1 int, p2 int) (float32, string) { /* calculate BMI */
	var bmi, l1, l2 float32
	bmi = float32(p2) / ((float32(p1) / 100) * (float32(p1) / 100)); l1 = 18.5; l2 = 24.9 /* calculation with constant 18 24 */
	label := "normal";r1 := big.NewFloat(float64(bmi)).Cmp(big.NewFloat(float64(l1)));r2 := big.NewFloat(float64(bmi)).Cmp(big.NewFloat(float64(l2))) /* comparison */
	if r1 > 0 {if r2 > 0 {label = "overweight"
	} else {label = "normal"
	}} else {label = "underweight"} /* populate variable for BMI's */
	return bmi, label
}
func bmi(w http.ResponseWriter, r *http.Request) { /* BMI func */
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":
      r.ParseForm()
      p1 := r.FormValue("height"); p2 := r.FormValue("weight"); 
      var rgx = regexp.MustCompile(`^[0-9]{1,3}$`)
      if auth(r.Header.Get("User-Agent")) && rgx.MatchString(p1) == true && rgx.MatchString(p2) == true { /* input checking */
        var pay = []byte(`{
        "bmi": "`)
        var load = []byte(`",
	"label": "`)
        p1c, _ := strconv.Atoi(p1); p2c, _ := strconv.Atoi(p2)
        bmi, label := calculate(p1c,p2c); bmic := strconv.FormatFloat(float64(bmi), 'f', 2, 32) /* calculate BMI */
        var result1 = append(pay, append([]byte(bmic), append(load, append([]byte(label), []byte(`"}`)...)...)...)...)
        w.WriteHeader(http.StatusOK) /* replying & logging */
        w.Write([]byte(result1)); log.Println("CMD:>bmi height="+p1+" weight="+p2+" bmi="+bmic)
      } else { w.WriteHeader(http.StatusBadRequest); w.Write([]byte(`{"text": "invalid input"}`))}; log.Println("CMD:>bmi invalid in")
    default:
      w.WriteHeader(http.StatusNotFound) /* invalid method */
      w.Write([]byte(`{"text": "invalid method"}`)); log.Println("CMD:>bmi invalid method")
    }
}
func hcx(w http.ResponseWriter, r *http.Request) { /* health check */
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":
      w.WriteHeader(http.StatusOK) /* replying & logging */
      w.Write([]byte(`pong`)); log.Println("CMD:>hcx")
    default:
      w.WriteHeader(http.StatusNotFound)
      w.Write([]byte(`{"text": "invalid method"}`)); log.Println("CMD:>hcx invalid method")
    }
}

func main() { /* listen each port & endpoint and serve the corresponding content */
    f, err := os.OpenFile("/var/log/a/access.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {panic(err)}; defer f.Close() /* log to text file */
    log.SetOutput(f)
    http.HandleFunc("/快鲜", bmi)
    http.HandleFunc("/建查", hcx)
    http.ListenAndServe(":10485", nil)
}