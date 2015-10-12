# Twilio #

Simple little code for my own use, but feel free to use it.

```
package main

import (
	"log"

	"github.com/shawnmilo/twilio"
)

func main() {
	err := twilio.Send("number goes here", "message goes here")
	if err != nil {
		log.Fatal(err)
	}
}
```

License: MIT
