package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//Path : estructura
type Path struct {
	name        string
	directories int
	symlinks    int
	devices     int
	sockets     int
	other       int
}

func printPath(path Path) {
	fmt.Println("+-------------------------+------+")
	fmt.Printf("| Path                     | %s |\n", path.name)
	fmt.Println("+-------------------------+------+")
	fmt.Printf("| Directories              | %d |\n", path.directories)
	fmt.Printf("| Symbolic Links           | %d |\n", path.symlinks)
	fmt.Printf("| Devices                  | %d |\n", path.devices)
	fmt.Printf("| Sockets                  | %d |\n", path.sockets)
	fmt.Printf("| Other files              | %d |\n", path.other)
}

// scanDir stands for the directory scanning implementation
func scanDir(dir string) error {
	// fmt.Println(dir)
	path := Path{dir, 0, 0, 0, 0, 0}
	err := filepath.Walk(dir, walkFunc)
	if err != nil {
		fmt.Printf("Error en el path %q: %v\n", dir, err)
	}
	printPath(path)
	return nil
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error [%v] en directorio [%q]\n", err, path)
		return err
	}
	switch mode := info.Mode(); {
	case mode.IsDir():
		//Es directorio
		fmt.Printf("Directorio: %v\n", info.Name())

	case mode&os.ModeSymlink != 0:
		//Es directorio
		fmt.Printf("Es symbolic link: %v\n", info.Name())
	case mode&os.ModeDevice != 0:
		//Es Device
		fmt.Printf("Es Device: %v\n", info.Name())
	case mode&os.ModeSocket != 0:
		//Es Socket
		fmt.Printf("Es Socket: %v\n", info.Name())
	case mode&os.ModeSocket != 0:
		//Es Socket
		fmt.Printf("Es Socket: %v\n", info.Name())
	default:
		//Es otra cosa
		fmt.Printf("Es otra Cosa: %v\n", info.Name())

	}
	return nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./dir-scan <directory>")
		os.Exit(1)
	}

	scanDir(os.Args[1])
}
