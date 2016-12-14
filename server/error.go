package main

type NDStoreArgsErrStruct struct {
	VOLUME string
}

var NDStoreArgsErr = NDStoreArgsErrStruct{
	VOLUME: "Error: Expected argument of the form /volume/<<server>>/<<token>>/<<channel>>/<<res>>/<<x0,x1>>/<<y0,y1>>/<<z0,z1>>/",
}
