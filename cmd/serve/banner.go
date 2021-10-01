package serve

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

var textLogoWebsite = "https://responserms.com"
var textLogoDocsWebsite = "https://docs.responserms.com"
var textLogo = `   ___                                
  / _ \___ ___ ___  ___  ___  ___ ___ 
 / , _/ -_|_-</ _ \/ _ \/ _ \(_-</ -_)
/_/|_|\__/___/ .__/\___/_//_/___/\__/  %s
            /_/
Open source CAD, MDT, and RMS for gaming communities.
%s · %s
----------------------------------------------------------------------
`

func WriteBanner(w io.Writer, version string) {
	_, _ = w.Write([]byte(fmt.Sprintf(textLogo, color.RedString(version), color.YellowString(textLogoWebsite), color.BlueString(textLogoDocsWebsite))))
}
