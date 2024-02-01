package enums

type PlayListAccessibility int32

const (
	PlayListAccessibility_PRIVATE = 0
	PlayListAccessibility_PUBLIC  = 1
)

type Genre int32

const (
	Genre_POP    Genre = 0
	Genre_HIPHOP Genre = 1
	Genre_ROCK   Genre = 2
	Genre_JAZZ   Genre = 3
	Genre_FOLK   Genre = 4
	Genre_BHAKTI Genre = 5
	Genre_LOVE   Genre = 6
)

// Enum value maps for AnalyserEntity.
var (
	Genre_name = map[int32]string{
		0: "POP",
		1: "HIPHOP",
		2: "ROCK",
		3: "JAZZ",
		4: "FOLK",
		5: "BHAKTI",
		6: "LOVE",
	}
	Genre_value = map[string]int32{
		"POP":    0,
		"HIPHOP": 1,
		"ROCK":   2,
		"JAZZ":   3,
		"FOLK":   4,
		"BHAKTI": 5,
		"LOVE":   5,
	}
)
