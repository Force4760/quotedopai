package suggest

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

const formURL = "https://forms.gle/UuoruSKGJpiMkqam8"

func OpenURL() {
	openbrowser(formURL)
}

func openbrowser(url string) {
	var err error

	// Different OS's have different ways of opening the default browser
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}

}
