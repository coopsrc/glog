package glog

import (
	"testing"
)

func TestColorLog(t *testing.T) {
	Verbose("TestColorLog", "go get -u %s", "http://gooim.me/dewdrop/glog")
	Trace("TestColorLog", "go get -u %s", "http://gooim.me/dewdrop/glog")
	Error("TestColorLog", "go get -u %s", "http://gooim.me/dewdrop/glog")
	Warn("TestColorLog", "go get -u %s", "http://gooim.me/dewdrop/glog")
	Info("TestColorLog", "go get -u %s", "http://gooim.me/dewdrop/glog")
	Debug("TestColorLog", "go get -u %s", "http://gooim.me/dewdrop/glog")
	Asset("TestColorLog", "go get -u %s", "http://gooim.me/dewdrop/glog")
}
