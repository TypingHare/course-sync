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

// Application directory name. This directory is created inside the project root directory and is
// hidden. It contains application-specific data, such as assignments.
const APP_DIR_NAME = ".csync"

// Documentation directory name.
const DOCS_DIR_NAME = "docs"

// Source code source directory name.
const SRC_DIR_NAME = "src"

// Prototype directory name. This directory is under the code source directory.
const PROTOTYPE_DIR_NAME = "[prototype]"

// Assignments file name inside the application directory.
const ASSIGNMENTS_FILE_NAME = "assignments.json"

// Grade history file name inside the application directory.
const GRADE_HISTORY_FILE_NAME = "grade_history.json"
