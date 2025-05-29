package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GA4abghigEJSDNwMiTn56VlputDKL2S1NyCnMSSVL3KxNyc0BBWBsbd19Uyeib45n1WEGj77nGH8eNaWR3XXyGFXXGJScdast4sPGORmRf/c3moeO7Vux9sPnw6ZF8eulbI+YSN2YwPR3TbjzS3WMc1JexOfM027/ik3RsFyrqCVQO2pE2/KPVcYWmb0sWfP6583xsvGT5n4bT0C9kldg66y8UfVhYXtc/ZWLsj+eXajc+muzGb88y0kG9doxh3PleUI9/jYqXCrfUXjnS7Lc89aP902enr5VGTHR4yNUVXHRIUu5Rb9u7UoouVjxgz5v7t094uOv/83//NOk26Vv/DGna7v9xyUOdpsXWO/be9cxIPyS+1e6W1vH3PpHp+Bob//wO82TmyxaZKvmJgYDjMyMAACzsGhhVVe1DCjg0eduBgeqtpngHSjKwkwJuRSYQZEfTIBoOCHgaWNIJIYiICYSJ2B0GAAMN/RxtGBkznsbKBpJkYmBj6GRgYKhlBPEAAAAD//9xPuSEeAgAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
