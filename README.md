# aplpdown

A simple Go program to recursively crawl a given URL and download all files and folders found on that page, preserving the original website's structure.

`aplpdown` stands for Advanced Programming Languages Principles Downloader Tool. This tool was made to help me download the CS252 course files.

Initially, I made a Python script but decided to port the tool to Go.

## Features

- Recursively downloads files and folders from a given URL.
- Maintains the directory structure of the source website.
- Supports a dry-run mode to preview actions without downloading files.

## Downloading and installing
All binaries are located at: https://github.com/ChristoferBerruz/aplpdown/releases/latest. You can always manually download the tar file, unpack it, move the file to a different directory, and add it to the path.

🚀 Alternatively, can use the following command to use the install script.
```bash
curl -sSL https://raw.githubusercontent.com/ChristoferBerruz/aplpdown/refs/heads/main/install.sh | sh
```

The above command was tested on a linux machine, and it should also work on MacOS.
In general, it is not a good idea to download and execute scripts from the internet. You can double check what the `install.sh` script does!

## Building and Installing

```sh
go install github.com/ChristoferBerruz/aplpdown@latest
```

> **Note:** After running `go install`, make sure that your Go bin directory (usually `$HOME/go/bin`) is in your `PATH` environment variable.  
> You can check this with `echo $PATH` and add it to your shell profile if needed.

## Usage

```sh
Usage: aplpdown [--dry-run] [--only regex] [--exclude regex] [--max-depth N] <url> <destination>
```

- `<url>`: The starting URL to crawl.
- `<destination>`: The local directory where files and folders will be saved.
- `-dry-run/--dry-run`: (Optional) If set, prints the files and folders that would be downloaded/created without actually downloading them.
- `--only` <regex>: (Optional) Only download files matching this regex pattern.
- `--exclude` <regex>: (Optional) Exclude files matching this regex pattern.
- `--max-depth` <N>: (Optional) Maximum crawl depth. -1 means unlimited.

### Example

```sh
aplpdown -dry-run https://example.com/files ./downloaded-files
```

This will print the files and folders that would be downloaded from `https://example.com/files` to the `./downloaded-files` directory, without actually downloading them.

If you want to actually download the files, simply omit the `-dry-run` flag.

## Notes

- The program was created to help download course materials from the CS252 course website.
- It uses the [`golang.org/x/net/html`](https://pkg.go.dev/golang.org/x/net/html) package for HTML parsing.

## License

MIT License