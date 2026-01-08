package assignment

import (
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var shouldForciblyPrepare bool

var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare an assignment for work.",
	Long:  `This command`,

	Run: func(cmd *cobra.Command, args []string) {
		assignmentName := args[0]

		assignments, err := feature.GetAssignments()
		if err != nil {
			cmd.PrintErrf("Error retrieving assignments: %v\n", err)
			return
		}

		// Check if the assignment name exists.
		assignment := feature.FindAssignmentByName(assignments, assignmentName)
		if assignment == nil {
			cmd.PrintErrf("Assignment '%s' not found\n", assignmentName)
			return
		}

		// Check if the prototype assignment directory exists.
		prototypeAssignmentDirPath, err := feature.GetPrototypeAssignmentDirPath(assignmentName)
		if err != nil {
			cmd.PrintErrf("Error retrieving prototype assignment directory: %v\n", err)
			return
		}

		// Check if the user assignment directory already exists.
		userAssignmentDirPath, err := feature.GetUserAssignmentDirPath(assignmentName)
		if err != nil {
			cmd.PrintErrf("Error retrieving user assignment directory: %v\n", err)
			return
		}

		// If not forcing, check if the user assignment directory already exists.
		if !shouldForciblyPrepare {
			exists, err := feature.DirExists(userAssignmentDirPath)
			if err != nil {
				cmd.PrintErrf("Error checking user assignment directory: %v\n", err)
				return
			}

			if exists {
				cmd.PrintErrf(
					"Assignment '%s' already prepared. Use --force to overwrite.\n",
					assignmentName,
				)
			}
		}

		// Delete the existing user assignment directory if forcing.
		if shouldForciblyPrepare {
			err = feature.DeleteDir(userAssignmentDirPath)
			if err != nil {
				cmd.PrintErrf("Error deleting existing user assignment directory: %v\n", err)
				return
			}
		}

		// Ensure the parent directory of the user assignment directory exists.
		err = feature.MakeDirIfNotExists(filepath.Dir(userAssignmentDirPath))
		if err != nil {
			cmd.PrintErrf("Error creating user assignments base directory: %v\n", err)
			return
		}

		// Copy the prototype assignment to the user assignment directory.
		err = feature.CopyDir(prototypeAssignmentDirPath, userAssignmentDirPath)
		if err != nil {
			cmd.PrintErrf("Error preparing assignment: %v\n", err)
			return
		}
	},
}

func init() {
	prepareCmd.PersistentFlags().
		BoolVarP(
			&shouldForciblyPrepare,
			"force",
			"f",
			false,
			"Forcibly prepare the assignment, overwriting any existing work.",
		)
}
