package glog

import (
	"testing"
)

func TestColorLog(t *testing.T) {
	VerboseF("TestColorLog F", "go get -u %s", "gooim.me/dewdrop/glog")
	TraceF("TestColorLog F", "go get -u %s", "gooim.me/dewdrop/glog")
	ErrorF("TestColorLog F", "go get -u %s", "gooim.me/dewdrop/glog")
	WarnF("TestColorLog F", "go get -u %s", "gooim.me/dewdrop/glog")
	InfoF("TestColorLog F", "go get -u %s", "gooim.me/dewdrop/glog")
	DebugF("TestColorLog F", "go get -u %s", "gooim.me/dewdrop/glog")
	AssetF("TestColorLog F", "go get -u %s", "gooim.me/dewdrop/glog")
	Verbose("TestColorLog", "go get -u gooim.me/dewdrop/glog")
	Trace("TestColorLog", "go get -u gooim.me/dewdrop/glog")
	Error("TestColorLog", "go get -u gooim.me/dewdrop/glog")
	Warn("TestColorLog", "go get -u gooim.me/dewdrop/glog")
	Info("TestColorLog", "go get -u gooim.me/dewdrop/glog")
	Debug("TestColorLog", "go get -u gooim.me/dewdrop/glog")
	Asset("TestColorLog", "go get -u gooim.me/dewdrop/glog")
}
