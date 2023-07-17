module github.com/krmcbride/hexample-golang/boot

go 1.20

require (
	github.com/krmcbride/hexample-golang/core v0.0.0
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect

replace github.com/krmcbride/hexample-golang/core => ../core
