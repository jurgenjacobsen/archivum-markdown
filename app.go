package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	watcher     *fsnotify.Watcher
	initialFile string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// InitialFile represents the file passed via command line
type InitialFile struct {
	Path   string `json:"path"`
	Parent string `json:"parent"`
}

// GetInitialFile returns the file path passed via command line
func (a *App) GetInitialFile() *InitialFile {
	if a.initialFile == "" {
		return nil
	}
	return &InitialFile{
		Path:   a.initialFile,
		Parent: filepath.Dir(a.initialFile),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// startWatching begins monitoring a directory recursively
func (a *App) startWatching(path string) {
	if a.watcher != nil {
		a.watcher.Close()
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Failed to create watcher: %v\n", err)
		return
	}
	a.watcher = watcher

	// Add all subdirectories to the watcher
	filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			err = a.watcher.Add(p)
			if err != nil {
				fmt.Printf("Failed to add %s to watcher: %v\n", p, err)
			}
		}
		return nil
	})

	go func() {
		for {
			select {
			case event, ok := <-a.watcher.Events:
				if !ok {
					return
				}
				// Handle structural changes
				if event.Op&(fsnotify.Create|fsnotify.Remove|fsnotify.Rename) != 0 {
					// If a new directory is created, watch it
					if event.Op&fsnotify.Create != 0 {
						info, err := os.Stat(event.Name)
						if err == nil && info.IsDir() {
							a.watcher.Add(event.Name)
						}
					}
					// Notify frontend with the parent directory of the change
					parent := filepath.Dir(event.Name)
					runtime.EventsEmit(a.ctx, "workspace-update", parent)
				}
			case err, ok := <-a.watcher.Errors:
				if !ok {
					return
				}
				fmt.Printf("Watcher error: %v\n", err)
			}
		}
	}()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// FileNode represents a file or directory in the file system
type FileNode struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	IsDir bool   `json:"isDir"`
}

// OpenWorkspaceDialog opens a directory dialog and returns the selected path
func (a *App) OpenWorkspaceDialog() (string, error) {
	result, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Workspace Folder",
	})
	if err != nil {
		return "", err
	}
	if result != "" {
		a.startWatching(result)
	}
	return result, nil
}

// GetDirectoryLevel returns the files and folders within a given directory path
func (a *App) GetDirectoryLevel(path string) ([]FileNode, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var nodes []FileNode
	for _, entry := range entries {
		nodes = append(nodes, FileNode{
			Name:  entry.Name(),
			Path:  filepath.Join(path, entry.Name()),
			IsDir: entry.IsDir(),
		})
	}
	return nodes, nil
}

// CreateFile creates a new empty file at the given path
func (a *App) CreateFile(path string) error {
	return os.WriteFile(path, []byte(""), 0644)
}

// CreateDirectory creates a new directory at the given path
func (a *App) CreateDirectory(path string) error {
	return os.Mkdir(path, 0755)
}

// Delete removes a file or directory at the given path
func (a *App) Delete(path string) error {
	return os.RemoveAll(path)
}

// Rename renames a file or directory
func (a *App) Rename(oldPath string, newPath string) error {
	return os.Rename(oldPath, newPath)
}

// ReadFile reads the content of a file
func (a *App) ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// SaveFile saves the content to a file
func (a *App) SaveFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
