package utils

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func GetWorldDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	switch runtime.GOOS {
	case "windows":
		return fmt.Sprintf("%s/AppData/Roaming/DungeonsNDungeons/worlds", dirname)
	case "linux":
		return fmt.Sprintf("%s/.local/share/DungeonsNDungeons/worlds", dirname)
	}
	return ""
}
