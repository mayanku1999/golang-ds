package enums

type PublishedMode int

const (
	PublishMode_UNSPECIFIED PublishedMode = 0
	PublishedMode_DRAFT                   = 1
	PublishedMode_PUBLISHED               = 2
)

type AccessMode int

const (
	AccessMode_UNSPECIFIED = 0
	AccessMode_WRITE       = 1
	AccessMode_READ        = 2
	AccessMode_WRITE_READ  = 3
)

type AccessStatus int

const (
	AccessStatus_UNSPECIFIED = 0
	AccessStatus_ACTIVE      = 1
	AccessStatus_INACTIVE    = 2
)

type UserRole int

const (
	UserRole_UNSPECIFIED = 0
	UserRole_L1          = 1 // most privileged i.e. admin
	UserRole_L2          = 2
	UserRole_L3          = 3
)

var (
	UserRole_name = map[int32]string{
		0: "USER_ROLE_UNSPECIFIED",
		1: "L1",
		2: "L2",
		3: "L3",
	}
	UserRole_value = map[string]int32{
		"USER_ROLE_UNSPECIFIED": 0,
		"L1":                    1,
		"L2":                    2,
		"L3":                    3,
	}
)
