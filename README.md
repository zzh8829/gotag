# GoTag - SemVer Tag Bumper

An interactive CLI tool for semantic version tagging in Git repositories, enabling users to easily bump version tags and push changes.

## Features

- Fetches latest semantic version tags from the Git repository.
- Interactive selection for version bump type: major, minor, or patch.
- Option to create and push new tags to the remote repository.

## Prerequisites

- Git accessible from the command line.
- Go installed on your system.

## Installation

You can install `gotag` using Go directly. Choose one of the following methods:

### Install the Latest Version

```bash
go install github.com/zzh8829/gotag@latest
```

### Install a Specific Version

```bash
go install github.com/zzh8829/gotag@v0.0.1
```

After installation, you can run `gotag` from anywhere in your command line.

## Usage

### Running Directly

If you've installed `gotag`, simply run:

```bash
gotag
```

### Running a Specific Version Directly Without Installing

```bash
go run github.com/zzh8829/gotag@v0.0.1
```

### Cloning and Running

Alternatively, you can clone the repository and run the tool directly:

```bash
git clone https://github.com/zzh8829/gotag.git
cd gotag
go run main.go
```

Follow the interactive prompts to select the version bump type and manage tags.

## Notes

- The script expects tags to follow the `v<major>.<minor>.<patch>` format and defaults to `v0.0.0` if no tags are found.
- Ensure you have the required permissions to push tags to the remote repository.

## Contributions

Contributions are welcome! If you have improvements or bug fixes, please open an issue or pull request on this repository.

## License

This script is provided "as is", without warranty of any kind, express or implied. Feel free to use and modify it as needed for your projects.
