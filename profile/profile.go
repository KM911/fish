package profile

import "github.com/pkg/profile"

/*
mode:
1: cpu
2: mem
3: block
4: mutex
*/
func Profile(mode uint) interface{ Stop() } {
	switch mode {
	case 1:
		return profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	case 2:
		return profile.Start(profile.MemProfile, profile.ProfilePath("."))
	case 3:
		return profile.Start(profile.BlockProfile, profile.ProfilePath("."))
	case 4:
		return profile.Start(profile.MutexProfile, profile.ProfilePath("."))
	default:
		return profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	}
}

func CPU() interface{ Stop() } {
	return profile.Start(profile.CPUProfile, profile.ProfilePath("."))
}

func Mem() interface{ Stop() } {
	return profile.Start(profile.MemProfile, profile.ProfilePath("."))
}

func Block() interface{ Stop() } {
	return profile.Start(profile.BlockProfile, profile.ProfilePath("."))
}

func Mutex() interface{ Stop() } {
	return profile.Start(profile.MutexProfile, profile.ProfilePath("."))
}
