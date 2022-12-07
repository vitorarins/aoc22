package day7

import (
	"fmt"
	"strconv"
	"strings"
)

// Directory represents a directory in the filesystem.
type Directory struct {
	Name     string
	Size     int
	Parent   *Directory
	Children map[string]*Directory
}

// NewDirectory creates a new directory with the given name.
func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		Name:     name,
		Parent:   parent,
		Children: make(map[string]*Directory),
	}
}

// AddFileSize adds a file with the given size to the directory and updates the parents
func (d *Directory) AddFileSize(size int) {
	d.Size += size
	if d.Parent != nil {
		d.Parent.AddFileSize(size)
	}
}

// AddDirectory adds a child directory to the directory.
func (d *Directory) AddDirectory(child *Directory) {
	d.Size += child.Size
	d.Children[child.Name] = child
}

// FindDirectory finds a child directory with the given name in the directory.
func (d *Directory) FindDirectory(name string) (*Directory, bool) {
	if dir, ok := d.Children[name]; ok {
		return dir, ok
	}

	return nil, false
}

// TotalSize returns the total size of the directory, which is the sum of the
// sizes of all the files it contains, either directly or indirectly.
func (d *Directory) TotalSize() int {
	return d.Size
}

func (d *Directory) String() string {
	parent := "{}"

	if d.Parent != nil {
		parent = d.Parent.String()
	}

	children := "["
	for _, c := range d.Children {
		children += `"` + c.Name + `",`
	}

	if len(d.Children) > 0 {
		children = strings.TrimSuffix(children, ",")
	}

	children += "]"

	return fmt.Sprintf(`{"Name":"%s", "Size":"%d", "Parent":%s, "Children":%s}`, d.Name, d.Size, parent, children)
}

func ParseCommand(root, current *Directory, line string) *Directory {

	line = strings.TrimSpace(line[1:])

	parts := strings.Split(line, " ")

	// ignore ls commands
	if len(parts) <= 1 {
		return current
	}

	arg := parts[1]

	switch arg {
	case "/":
		current = root
	case "..":
		current = current.Parent
	default:
		child, ok := current.FindDirectory(arg)
		if !ok {
			child = NewDirectory(arg, current)
			current.AddDirectory(child)
		}
		current = child
	}

	return current
}

func ParseListOutput(current *Directory, line string) (*Directory, error) {

	parts := strings.Split(line, " ")

	name := parts[1]

	if parts[0] == "dir" {

		if _, ok := current.FindDirectory(name); !ok {
			child := NewDirectory(name, current)
			current.AddDirectory(child)
		}

		return current, nil
	}

	size, err := strconv.Atoi(parts[0])
	if err != nil {
		return current, err
	}

	current.AddFileSize(size)

	return current, nil
}

func ParseCommandsAndOutput(input string) (*Directory, error) {

	root := NewDirectory("/", nil)
	current := root

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, "$"):
			current = ParseCommand(root, current, line)
		default:
			var err error

			current, err = ParseListOutput(current, line)
			if err != nil {
				return nil, fmt.Errorf("failed to parse list output %q: %w", line, err)
			}
		}
	}

	return root, nil
}

func GetSumUpTo(root *Directory, limit int) int {

	var sum int

	// Perform a depth-first search of the filesystem, checking the size of each directory as it is visited.
	var visit func(dir *Directory)
	visit = func(dir *Directory) {
		if dir.TotalSize() <= limit {
			// If the directory has a total size of at most limit, add its size to the sum.
			sum += dir.TotalSize()
		}

		// Visit all of the children of the directory.
		for _, child := range dir.Children {
			visit(child)
		}
	}
	visit(root)

	return sum
}

// SumTotalSizes calculates the sum of the total sizes of the directories with a total size of at most maxSize. It also returns the smallest directory with a total size of at least minSize.
func SumTotalSizes(dir *Directory, maxSize, minSize int) (int, *Directory) {
	if dir.TotalSize() > maxSize {
		return 0, nil
	}
	sum := dir.TotalSize()
	minDir := dir

	for _, child := range dir.Children {
		childSum, childMinDir := SumTotalSizes(child, maxSize, minSize)
		sum += childSum
		if minDir == dir || (childMinDir != nil && childMinDir.TotalSize() < minDir.TotalSize()) {
			minDir = childMinDir
		}
	}

	if minDir.TotalSize() >= minSize {
		return sum, minDir
	}
	return sum, nil
}

func GetFileSizes(input string) (int, error) {

	input = strings.TrimSpace(input)

	elfFS, err := ParseCommandsAndOutput(input)
	if err != nil {
		return 0, fmt.Errorf("failed to parse commands: %w", err)
	}

	sumFileSizes := GetSumUpTo(elfFS, 100000)

	return sumFileSizes, nil
}

func GetSumMinDir(root *Directory, minSize int) int {

	chosenOne := root

	// Perform a depth-first search of the filesystem, checking the size of each directory as it is visited.
	var visit func(dir *Directory)
	visit = func(dir *Directory) {
		if dir.TotalSize() >= minSize && dir.Size < chosenOne.Size {
			// If the directory has a total size of at most limit, add its size to the sum.
			chosenOne = dir
		}

		// Visit all of the children of the directory.
		for _, child := range dir.Children {
			visit(child)
		}
	}
	visit(root)

	return chosenOne.TotalSize()
}

func GetMinDir(input string) (int, error) {

	input = strings.TrimSpace(input)

	elfFS, err := ParseCommandsAndOutput(input)
	if err != nil {
		return 0, fmt.Errorf("failed to parse commands: %w", err)
	}

	totalSpace := 70000000
	required := 30000000

	toRemove := required - (totalSpace - elfFS.TotalSize())

	sumDir := GetSumMinDir(elfFS, toRemove)

	return sumDir, nil
}
