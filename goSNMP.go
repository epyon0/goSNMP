package main

import (
	"fmt"
	"os"
	"strings"

	utils "github.com/epyon0/goUtils"
)

var debug bool = false
var mibs []string
var dirs []string

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		arg := args[i]

		switch arg {
		case "-h":
			fallthrough
		case "--help":
			fmt.Printf("-h | --help    Print this help message\n")
			fmt.Printf("-v | --verbose Display verbose output\n")
			fmt.Printf("-d | --dir     Specify directory of mib files to load, files must have .mib\n")
			fmt.Printf("               file extension")
			fmt.Printf("-m | --mib     Specify individual mob file to load, will accept any file\n")

			return
		case "--verbose":
			fallthrough
		case "-v":
			debug = true
		case "-d":
			fallthrough
		case "--dir":
			if i+1 < len(args) {
				dirs = append(dirs, args[i+1])
				i++
			}
		case "-m":
			fallthrough
		case "--mib":
			if i+1 < len(args) {
				fileInfo, err := os.Stat(args[i+1])
				utils.Er(err)

				if fileInfo.Mode().IsRegular() {
					utils.Debug(fmt.Sprintf("Loading MIB file \"%s\"", args[i+1]), debug)
					mibs = append(mibs, args[i+1])
				} else {
					utils.Debug(fmt.Sprintf("Ignoring file \"%s\", file is of type: %v", args[i+1], fileInfo.Mode()), debug)
				}
				i++
			}
		default:
			fmt.Printf("Unknown flag \"%s\", terminating...\n", arg)
			return
		}
	}

	utils.Debug("Initializing...", debug)
	for i := 0; i < len(dirs); i++ {
		dir := dirs[i]
		fileInfo, err := os.Stat(dir)
		utils.Er(err)
		if fileInfo.Mode().IsDir() {
			files, err := os.ReadDir(dir)
			utils.Er(err)

			for j := 0; j < len(files); j++ {
				file := fmt.Sprintf("%s/%s", dir, files[j].Name())
				tmpInfo, err := os.Stat(file)
				utils.Er(err)

				if tmpInfo.Mode().IsRegular() {
					if len(file) > 4 && strings.ToLower(file[len(file)-4:]) == ".mib" {
						utils.Debug(fmt.Sprintf("Loading MIB file \"%s\"", file), debug)
						mibs = append(mibs, file)
					} else {
						utils.Debug(fmt.Sprintf("Ignoring non-MIB file \"%s\"", file), debug)
					}
				} else {
					utils.Debug(fmt.Sprintf("Ignoring file \"%s\", file is of type: %v", file, tmpInfo.Mode()), debug)
				}
			}
		} else {
			utils.Er(fmt.Errorf("\"%s\" is not a valid directory", dirs[i]))
		}
	}
	utils.Debug("Configuration:", debug)
	utils.Debug(fmt.Sprintf("  debug = %t", debug), debug)
	for i := 0; i < len(dirs); i++ {
		utils.Debug(fmt.Sprintf("  dir = %s", dirs[i]), debug)
	}
	for i := 0; i < len(mibs); i++ {
		utils.Debug(fmt.Sprintf("  mib = %s", mibs[i]), debug)
	}

	/*
		dir := "/mnt/c/Users/jasong/mibs/IETF"
		files, err := os.ReadDir(dir)
		utils.Er(err)
		for _, file := range files {
			//mod, err := smi.ParseModule("/mnt/c/Users/jasong/mibs/IETF/IF-MIB.mib")
			mod, err := smi.ParseModule(fmt.Sprintf("%s/%s", dir, file.Name()))
			if err != nil {
				utils.Debug(fmt.Sprintf("%v", err))
				continue
			}

			fmt.Printf("Name:     %s\n", mod.Name)
			fmt.Printf("File:     %s\n", mod.File)
			fmt.Printf("IsLoaded: %t\n", mod.IsLoaded)
			//fmt.Printf("Imports: %v, %T\n", mod.Imports, mod.Imports)
			fmt.Printf("Imports:\n")
			for i := 0; i < len(mod.Imports); i++ {
				fmt.Printf("    From: %s\n", mod.Imports[i].From)
				for j := 0; j < len(mod.Imports[i].Symbols); j++ {
					fmt.Printf("        Symbols: %s\n", mod.Imports[i].Symbols[j])
				}
			}
			fmt.Printf("Nodes:\n")
			for i := 0; i < len(mod.Nodes); i++ {
				fmt.Printf("    Label: %s\n", mod.Nodes[i].Label)
				fmt.Printf("    Type:  %d\n", mod.Nodes[i].Type)
				fmt.Printf("    IDs:\n")
				for j := 0; j < len(mod.Nodes[i].IDs); j++ {
					fmt.Printf("        ID:    %d\n", mod.Nodes[i].IDs[j].ID)
					fmt.Printf("        Label: %s\n", mod.Nodes[i].IDs[j].Label)
				}
			}
			fmt.Printf("Symbols:\n")
			for k, v := range mod.Symbols {
				fmt.Printf("    %s:\n", k)
				fmt.Printf("        Name: %s\n", v.Name)
				fmt.Printf("        ID:   %d\n", v.ID)

			}
		}
	*/

	//mib := smi.NewMIB("/mnt/c/Users/jasong/mibs/IETF")
	//mib.Debug = true

	//err := mib.LoadModules("IF-MIB")
	//utils.Er(err)
}
