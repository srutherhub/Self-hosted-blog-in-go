# Go Markdown Blog Web Server

This project is a simple web server written in Go that serves markdown files. It's designed as a learning experience for myself to explore Go's capabilities, particularly in web development and file system interaction. I've also made it publicly available to showcase my work and provide a useful tool for others.

## Features

* **Automatic Markdown Rendering:** Converts markdown files to HTML on the fly.
* **Simple Setup:** Easy to use with minimal configuration.
* **Cross-Platform Executable:** Can be run on macOS, Linux, and Windows.

## Getting Started

You can use this project in two ways: by downloading a pre-built executable or by cloning the repository and building it yourself.

### Method 1: Downloading the Executable

1.  **Download the Executable:**
    * Called self-hosted-blog

2.  **Create a `home` Folder:**
    * Create a directory named `home` in the same location as the executable.
    * Place your markdown files (`.md`) inside the `home` folder. You can also create subdirectories within the `home` folder to organize your files.

3.  **Run the Executable:**
    * Open your terminal or command prompt.
    * Navigate to the directory where you downloaded the executable.
    * Run the executable:
        * macOS/Linux: `./self-hosted-blog`
        * Windows: `self-hosted-blog`

4.  **Access the Server:**
    * Open your web browser and go to `http://localhost:8080`.
    * You can access your markdown files by navigating to `http://localhost:8080/md/yourfile.md` (or the relative path inside the home folder). For example if you have a file home/about/me.md, you would access it using http://localhost:8080/md/about/me.md

### Method 2: Cloning the Repository

1.  **Clone the Repository:**
    * You have to have Go downloaded to compile the project
    * Once you do, add your markdown to 'home' directory

2. **Run Web Server:**
    * run 'go run .' to start the webserver and access if from `http://localhost:8080`

## Configuration

* By default, the server serves markdown files from the `home` directory.
* The server listens on port `8080`.

## Contributing

Contributions are welcome! If you find a bug or have an idea for a new feature, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Author

Sam Rutherford

## Contact
