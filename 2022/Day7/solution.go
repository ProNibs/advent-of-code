package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Define Directory data structure
type Directory struct {
	name string
	// Parent and child names should match DirectoryList
	parent    string
	children  map[string]string
	files     map[string]File
	totalSize int
}

// Define File data structure
type File struct {
	name string
	size int
}

func initDirectory(name string) Directory {
	new_directory := new(Directory)
	new_directory.name = strings.Clone(name)
	new_directory.files = map[string]File{}
	new_directory.children = make(map[string]string)
	new_directory.totalSize = 0
	return *new_directory
}

func main() {
	lines, _ := readLines("data.txt")
	// Root is always there at least
	root_directory := initDirectory("/")
	var running_ls_command = false
	var directory_list = make(map[string]Directory)
	directory_list["/"] = root_directory

	var current_directory = root_directory
	var pwd = "/"
	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			//fmt.Println("Got the following command:", line)
			if line == "$ cd /" {
				current_directory = root_directory
				pwd = "/"
			} else if strings.HasPrefix(line, "$ cd") {
				var cd_to_this = strings.Split(line, " ")[2]
				if cd_to_this != ".." {
					// We are cd'ing to a subdir, so let's go in
					pwd = pwd + cd_to_this + "/"
					if _, ok := current_directory.children[pwd]; ok {
						// Good, it exists already
						current_directory = directory_list[pwd]
					} else {
						fmt.Println("Directory doesn't exist already, bad assumption")
						fmt.Println("Pwd attempted:", pwd)
						fmt.Println("Current dir:", current_directory)
						break
					}
				} else {
					var full_cur_dir_name = strings.Split(current_directory.name, "/")
					var cur_dir_name = full_cur_dir_name[len(full_cur_dir_name)-2]
					//fmt.Println("Trim me", pwd, full_cur_dir_name, cur_dir_name)
					pwd = strings.TrimSuffix(pwd, cur_dir_name+"/")
					//fmt.Println("Trimmed pwd", pwd)
					current_directory = directory_list[current_directory.parent]
				}
			}
			if line == "$ ls" {
				running_ls_command = true
			} else {
				running_ls_command = false
			}
		} else if running_ls_command {
			// Must be files if doesn't start with $, but helps to be sure
			if strings.HasPrefix(line, "dir") {
				var dir_name = strings.Split(line, " ")[1]
				// Check if directory already exists
				if _, ok := current_directory.children[dir_name]; ok {
					fmt.Println("ls on the same directory??? Bad assumptions dude...")
					fmt.Println("Current directory:", current_directory)
					fmt.Println("Directory list:", directory_list)
					break
				} else {
					// Found a directory that doesn't exist, let's make a new one
					//fmt.Println("Created directory", dir_name)
					var new_directory = initDirectory(pwd + dir_name + "/")
					new_directory.parent = current_directory.name
					directory_list[new_directory.name] = new_directory
					directory_list[current_directory.name].children[new_directory.name] = new_directory.name
				}
			} else {
				// Not a directory? Must be a file!
				var split_array = strings.Split(line, " ")
				var file_name = split_array[1]
				var file_size, _ = strconv.Atoi(split_array[0])
				current_directory.files[file_name] = File{name: file_name, size: file_size}
			}
		}
	}
	// Make an array of arrays
	// Top level is to denote how many child directories a directory has
	// Insides are the names of the dictionaries
	child_nodes := make([][]string, len(directory_list)-1)
	for _, v := range directory_list {
		child_nodes[len(v.children)] = append(child_nodes[len(v.children)], v.name)
	}
	// Now from 0->"/" directory, let's get those file sizes
	for _, dir_list := range child_nodes {
		for _, dir_name := range dir_list {
			var local_size int = 0
			for _, v := range directory_list[dir_name].files {
				local_size += v.size
			}
			// We can assume any childs already had their size computed
			for _, j := range directory_list[dir_name].children {
				local_size += directory_list[j].totalSize
			}
			// Have to make a copy of struct to update it
			var dir_update = directory_list[dir_name]
			dir_update.totalSize = local_size
			// Then replace it
			directory_list[dir_name] = dir_update
		}
	}
	// Check total sizes now
	var solution_one = 0
	for _, v := range directory_list {
		if v.totalSize <= 100000 {
			solution_one += v.totalSize
		}
		fmt.Println(v)
	}
	fmt.Println(solution_one)
}
