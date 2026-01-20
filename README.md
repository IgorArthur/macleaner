### MaCleaner

MaCleaner is a command-line tool for cleaning up Docker-related files on macOS, Linux and Windows systems.

## Installation

`go install github.com/igorarthur/macleaner@latest`

## Usage

### macleaner doctor

Explains system data growth with Docker and what this tool can/cannot clean.

    example:
        `macleaner doctor`

### macleaner scan

Scans system for Docker-related files and echoes the total size of them.

    example:
        `macleaner scan`

### macleaner clean

Cleans up Docker-related files.

    flags:
        `--dry-run`   => shows files that will be delete, but does not conclude the operation
        `--yes, -y`   => skip user confirmation
    
    example:
        `macleaner clean --dry-run -y`