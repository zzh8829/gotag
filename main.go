package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Pull the latest tags from Git
	cmd := exec.Command("git", "fetch", "--tags")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error fetching tags:", err)
		return
	}

	// Get the latest tag
	cmd = exec.Command("git", "tag")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting tags:", err)
		return
	}

	tags := strings.Split(string(output), "\n")
	sort.Strings(tags)

	if len(tags) < 2 {
		fmt.Println("Warn: No tags found: using 0.0.0")
		tags = append(tags, "v0.0.0")
	}
	latestTag := tags[len(tags)-1]

	fmt.Println("Latest tag:", latestTag)

	// Ask user for version bump type
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bump ma(j)or, mi(n)or, or (p)atch?")
	bumpType, _ := reader.ReadString('\n')
	bumpType = strings.TrimSpace(bumpType)

	// Parse the latest tag and increment version
	parts := strings.Split(strings.TrimPrefix(latestTag, "v"), ".")
	if len(parts) != 3 {
		fmt.Println("Error: Latest tag does not follow semantic versioning")
		return
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	switch bumpType {
	case "j":
		major++
		minor = 0
		patch = 0
	case "n":
		minor++
		patch = 0
	case "p":
		patch++
	default:
		fmt.Println("Invalid option")
		return
	}

	newTag := fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	fmt.Println("New tag:", newTag)

	fmt.Println("Create the tag? (y/n)")
	create, _ := reader.ReadString('\n')
	create = strings.TrimSpace(create)
	if create != "y" {
		fmt.Println("Tag creation aborted.")
		return
	}

	// Create tag with current commit
	cmd = exec.Command("git", "tag", newTag)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error creating new tag:", err)
		return
	}

	// Ask user if they want to push the commit with the tag
	fmt.Println("Push the commit with the tag? (y/n)")
	push, _ := reader.ReadString('\n')
	push = strings.TrimSpace(push)

	if push == "y" {
		cmd = exec.Command("git", "push", "--atomic", "origin", "HEAD", newTag)
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error pushing new tag:", err)
			return
		}
		fmt.Println("Tag pushed successfully.")
	} else {
		fmt.Println("Tag creation completed without pushing.")
	}
}
