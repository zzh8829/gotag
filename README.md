# Go Tag

This Go script provides a simple CLI tool to interactively bump semantic versioning (semver) tags in your Git repository. It automates the process of fetching the latest tag, determining the next version, and optionally pushing the new tag to the remote repository.

## Features

- Fetches the latest semantic versioning tags from the Git repository.
- Interactively asks the user to choose the version bump type: major, minor, or patch.
- Creates a new tag with the incremented version based on the user's choice.
- Optionally pushes the new tag to the remote repository.

## Prerequisites

- Git installed and accessible from your command line.
- Go installed on your system.

## Installation

To use the Semver Tag Bumper, follow these steps:

1. Clone this repository or copy the script into a new `.go` file in your Go workspace.
2. Ensure you have Go installed and set up correctly.

## Usage

1. Navigate to your Git repository where you want to bump the version tag.
2. Run the script using `go run` followed by the script's file name. For example:

```bash
go run semver_tag_bumper.go
```

3. Follow the interactive prompts to bump the tag version:
   - Choose `j` for a major version bump.
   - Choose `n` for a minor version bump.
   - Choose `p` for a patch version bump.
4. Decide whether to create the tag.
5. Decide whether to push the tag to the remote repository.

## Notes

- The script assumes your tags follow the `v<major>.<minor>.<patch>` format.
- It will default to `v0.0.0` if no tags are found.
- Ensure you have the necessary permissions to push tags to the remote repository.

## Contributions

Contributions are welcome! If you have improvements or bug fixes, please open an issue or pull request on this repository.

## License

This script is provided "as is", without warranty of any kind, express or implied. Feel free to use and modify it as needed for your projects.
