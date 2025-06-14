# Welcome to CobraClip CLI

CobraClip is a powerful command-line tool designed to streamline your GitHub repository management with a wide range of operations. With just a single login using your GitHub Personal Access Token (PAT), you can execute CobraClip commands from any terminal to interact with GitHub effortlessly.

## Installation
1. **Install Go**:
   - Download and install Go from [golang.org](https://golang.org/dl/). Follow the instructions for your operating system (Windows, macOS, or Linux).
   - Verify installation:
     ```bash
     go version
     ```
2. **Install CobraClip**:
   - Clone the repository:
     ```bash
     git clone https://github.com/DipanshuOjha/cobraclip.git
     cd cobraclip
     ```
   - Install the CLI:
     ```bash
     go install
     ```
3. **Build CobraClip**:
   - Build the executable:
     ```bash
     go build -o cobraclip
     ```
   - On Windows, this creates `cobraclip.exe`. Move it to a directory in your PATH (e.g., `C:\Users\<your-username>\go\bin`) for global access.
4. **Enjoy CobraClip**:
   - Run `cobraclip --help` to explore commands.

## Features
- **Create a Repository**: Initialize new GitHub repositories with ease.
- **List Personal Repositories**: View all repositories owned by your account, including public and private ones.
- **Search Organizations Globally**: Find GitHub organizations worldwide by name or keyword.
- **Search Repositories Globally**: Discover repositories across GitHub by name, author, or other criteria.
- **List Issues and Pull Requests**: Retrieve issues, pull requests, and more from any repository.
- **Fork Repositories**: Create forks of any public repository to contribute or experiment.
- **Delete and Update Personal Repositories**: Modify repository details like title (name) and description, or delete repositories you own.
- **Clone Repositories**: Clone repositories to your local machine while exploring global repos.

## Getting Started
1. **Generate a GitHub PAT**:
   - Go to [GitHub Settings > Developer settings > Personal access tokens](https://github.com/settings/tokens).
   - Create a token with `repo` scope (and others as needed, e.g., `delete_repo` for deletion).
2. **Log In with CobraClip**:
   - Run:
     ```bash
     cobraclip login --token <your-token>
     ```
   - Your token is securely stored in Windows Credential Manager (or equivalent for other platforms).
3. **Use CobraClip Commands**:
   - Example commands:
     ```bash
     cobraclip repo create --name my-new-repo --private
     cobraclip repo listmyrepo
     cobraclip repo SearchRepo --repo cobra --author spf13
     cobraclip repo forkRepo --org spf13 --repo cobra
     cobraclip repo update-title --owner <username> --repo my-repo --name new-name
     ```

## Why CobraClip?
- **Simple Authentication**: Log in once, and access all features securely.
- **Global Search**: Explore repositories and organizations worldwide.
- **Comprehensive Management**: From creation to deletion, manage your GitHub workflow in one tool.
- **User-Friendly**: Built with Cobra for intuitive commands and flags.

## Support
If you encounter any issues or have feature requests, feel free to [open an issue](https://github.com/DipanshuOjha/cobraclip/issues) on our GitHub repository.

## Thank You for Using CobraClip!
Crafted with ❤️ by Dipanshu Ojha