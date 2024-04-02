package model

import "github.com/phonepeproj/proj/enums"

type Privileges struct {
	Privileges map[string]*PrivilegeInfo
}

type PrivilegeInfo struct {
	AccessMode   enums.AccessMode
	AccessStatus enums.AccessStatus
}
