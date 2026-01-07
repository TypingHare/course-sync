package app

// Application name.
const NAME = "Course Sync"

// Application version.
const VERSION = "2026.1"

// Executable name.
const EXECUTABLE_NAME = "csync"

// Project root marker directory name. Folders containing this directory are considered project root
// directory (or simply project directory). The Git hidden directory is used as the project root
// marker because Course Sync is intended to be used in Git repositories.
const PROJECT_ROOT_MARKER_DIR_NAME = ".git"

// Configuration file name
const CONFIG_FILE_NAME = "course_sync.config.json"

// Application hidden directory name. This directory is created inside the project root directory.
// It contains temporary files used by Course Sync. It should be ignored by version control systems.
const HIDDEN_DIR_NAME = ".csync"
