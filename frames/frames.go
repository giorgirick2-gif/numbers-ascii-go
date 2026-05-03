package frames

import "time"

// FrameType defines the interface for an animation
type FrameType struct {
	GetFrame  func(int) string
	GetLength func() int
	GetSleep  func() time.Duration
}

// DefaultGetFrame calculates which frame to show based on the loop index
func DefaultGetFrame(frames []string) func(int) string {
	return func(i int) string {
		return frames[i%(len(frames))]
	}
}

// DefaultGetLength returns the total number of frames in the slice
func DefaultGetLength(frames []string) func() int {
	return func() int {
		return len(frames)
	}
}

// DefaultGetSleep sets the animation speed (70ms is the default)
func DefaultGetSleep() func() time.Duration {
	return func() time.Duration {
		return time.Millisecond * 70
	}
}

// DefaultFrameType is a constructor that builds the FrameType struct from a string slice
func DefaultFrameType(frames []string) FrameType {
	return FrameType{
		GetFrame:  DefaultGetFrame(frames),
		GetLength: DefaultGetLength(frames),
		GetSleep:  DefaultGetSleep(),
	}
}

// FrameMap is the "Directory" the web server uses to find your animations
var FrameMap = map[string]FrameType{
	"num":             Num,
	"numascii":        NumAscii,
	"coin":            Coin,
	"batman":          Batman,
	"batman-running":  BNR,
	"bnr":             BNR,
	"can-you-hear-me": Rick,
	"clock":           Clock,
	"donut":           Donut,
	"dvd":             Dvd,
	"forrest":         Forrest,
	"hes":             HES,
	"knot":            TorusKnot,
	"nyan":            Nyan,
	"parrot":          Parrot,
	"rick":            Rick,
	"spidyswing":      Spidy,
	"torus-knot":      TorusKnot,
	"purdue":          Purdue,
	"as":              AStrend,
	"bomb":            Bomb,
	"maxwell":         Maxwell,
	"earth":           Earth,
	"kitty":           Kitty,
	"india":           India,
	"brittany":        Brittany,
}