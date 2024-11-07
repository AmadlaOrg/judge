package plugin

// Plugins contains the information that are required
type Plugins struct {
	Path         string // The absolute path to the plugin of the file to execute
	IsBinary     bool   // If binary than it means that is has
	IsExecutable bool   // The code checks if
	Shell        string // (default is the as the terminal)
}
