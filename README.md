# tudp
test for udp socket. used on arm-linux

# usage

	Toby@GITER-TOBY-PC /D/prj/go/src/github.com/GiterLab/tudp (master)
	$ ./tudp.exe -h
	Usage of d:\prj\go\src\github.com\GiterLab\tudp\tudp.exe:
	  -m="Hello GiterLab!": Udp message.
	  -n=false: Send message end with \n.
	  -r="127.0.0.1:3333": Remote udp address
	  -t=0: Setup udp timeout.
	Toby@GITER-TOBY-PC /D/prj/go/src/github.com/GiterLab/tudp (master)
	$ ./tudp.exe -r echo.giter.org:8888 -m zhangxiaojie -n -t 5
	Send message: "zhangxiaojie" to echo.giter.org:8888
	Recv: 7A 68 61 6E 67 78 69 61 6F 6A 69 65 0A
	[ToStrings] --> zhangxiaojie

