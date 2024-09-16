# Habbo Origins Launcher (HOL)

Habbo Origins Launcher (HOL) is a lightweight, efficient launcher for Habbo Origins. Unlike the default Habbo launcher, HOL is designed to be resource-efficient and easy to use, providing a seamless way to update and launch the Habbo client without leaving resource-intensive processes running in the background.

## Features
- **Multi-Instance Support:** Handles temporary folders for different instances, allowing multiple accounts to run with ease.
- **Lightweight Launcher:** Avoids the resource-intensive nature of the official Habbo launcher, freeing up system resources.
- **Automatic Updates:** Automatically checks for new versions of Habbo Origins and downloads updates as they become available.
- **Customizable Installation Path:** Allows for custom installation paths via `config.yml`, with a default fallback to the same directory as the Habbo launcher.
- **Old Version Cleanup:** Configurable option to delete old versions and their temporary subfolders to save disk space, similar to how habbo does it.

## Build
1. Download the repository and build the project using Go:
   ```bash
   git clone https://github.com/Edaurd/HOL.git
   cd HOL
   go build -o bin/HOL.exe -ldflags="-H=windowsgui" .
   ```
2. Place the executable in your desired directory.

## Configuration
HOL uses a `config.yml` file for configuration. If no `config.yml` is provided, it will use default settings.

Here’s the default `config.yml`:
```yaml
path: ""
country: "us"
xl: true
delete_old_version: true
```
if the config is absent or missing or has issues the values above will be used

### Configuration Options
- **`path`**: Specifies the folder where the launcher stores the respective clients/version. If left empty, HOL will use the default path (`%appdata%/Habbo Launcher/downloads/shockwave`). If a relative path is specified (e.g., `./HabboOrigins`), it will create a folder with all the version stored within the specified path.
- **`country`**: Specifies the country version of Habbo to launch (e.g., `"us"`, `es`, `"br"`, `d`). Defaults to `"us"` If invalid or missing.
- **`xl`**: A boolean indicating whether to launch the large client. If invalid or missing, it defaults to `true`.
- **`delete_old_version`**: A boolean that controls whether old versions should be deleted when a new update occurs:
  - `true`: Deletes the previous version along with its subfolders.
  - `false`: Deletes only the instance folders while keeping the main folder.
  - After an update, all instance folders (e.g., `35_1`, `37_1`, `37_2`, `37_3`) for every version will be deleted, regardless of the `delete_old_version` setting. These instance folders are merely a copy of their main version folders and do not hold any significance.

## Usage
1. [Download the latest version](https://github.com/Edaurd/HOL/releases/latest) of HOL. 
2. If necessary, create a `config.yml` file in the same folder as your exe and edit it to your preferences.<br>
  Optional: right click `HOL.exe` and select `Send to` > `Desktop (Create Shortcut)` to keep it looking clean.<br>
            You could also drag and drop it into your navigation bar.  
4. Run `HOL.exe`
5. HOL will automatically check for new updates, download them, and launch the Habbo client based on your configuration.

## How It Works
1. **Checks for Updates:** HOL checks for the latest version of Habbo Origins. If a new version is available, it downloads and extracts it to the specified path.
2. **Handles Instance Folders:** If you want to launch multiple clients at once, it creates instance folders (e.g., `37_1`, `37_2`) because only one instance of the client can be ran per installation.
3. **Deletes Old Versions:** If `delete_old_version` is set to `true`, it deletes the most recent version (e.g., new update `38` it will only delete `37` and won't touch `36`), any left over instance folders (e.g., `37_1`, `37_2` but also `36_1`). If set to `false`, only instance folders are deleted.

## Troubleshooting
- Drag and drop the exe into a command prompt (cmd.exe) and hit enter, if anything wrong it should show some sort of message there 

## Contributing
Feel free to fork the repository and submit pull requests if you have suggestions or improvements for the launcher.

## TLDR;
it can do everything habbo launcher can and more. See above for features 
