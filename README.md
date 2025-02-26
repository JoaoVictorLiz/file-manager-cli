# File Manager CLI

File Manager CLI is a command-line tool developed in Go to facilitate file and directory management.

## 🚀 Installation

1. Clone this repository:
   ```sh
   git clone https://github.com/your-username/file-manager-cli.git
   cd file-manager-cli
   ```

2. Build the project:
   ```sh
   go build -o file-manager-cli
   ```

3. (Optional) Move the executable to a globally accessible directory:
   ```sh
   mv file-manager-cli /usr/local/bin/
   ```
   On Windows, move it to a path that is in the `PATH`.

## 📌 Usage

File Manager CLI allows you to list, create, remove, move, and rename files/directories.

### 🔹 List files and directories
```sh
file-manager-cli list
```

### 🔹 Create a file
```sh
file-manager-cli create -f my-file.txt
```

### 🔹 Create a directory
```sh
file-manager-cli create -d my-folder
```

### 🔹 Remove a file
```sh
file-manager-cli remove -r my-file.txt
```

### 🔹 Move a file to a directory
```sh
file-manager-cli move -m my-file.txt my-folder/
```

### 🔹 Rename a file
```sh
file-manager-cli rename -n old.txt new.txt
```

## 🛠 Technologies Used
- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra) for CLI creation
- [Fatih Color](https://github.com/fatih/color) for coloring terminal output

## 🤝 Contribution
Feel free to contribute! To do so:
1. Fork this repository
2. Create a branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to the main branch (`git push origin my-new-feature`)
5. Open a Pull Request 🚀

