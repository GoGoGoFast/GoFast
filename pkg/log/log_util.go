package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// LogLevel defines the log levels
type LogLevel int

const (
	DEBUG LogLevel = iota // DEBUG level log
	INFO                  // INFO level log
	WARN                  // WARN level log
	ERROR                 // ERROR level log
)

const logChannelBufferSize = 1000 // Buffer channel size

// LoggerConfig is the configuration structure for Logger
type LoggerConfig struct {
	Level      LogLevel // Log level
	FilePath   string   // Log file path
	MaxSize    int64    // Maximum size of a single log file (bytes)
	MaxBackups int      // Maximum number of old log files to retain
	MaxAge     int      // Maximum number of days to retain old log files
	Compress   bool     // Whether to compress old log files
}

// Logger is a custom logger
type Logger struct {
	mu      sync.Mutex   // Mutex for ensuring concurrency safety
	config  LoggerConfig // Log configuration
	logFile *os.File     // Log file handle
	log     *log.Logger  // Go standard library logger
	logChan chan string  // Buffer channel for asynchronous log writing
}

// Global singleton instance
var instance *Logger
var once sync.Once

// NewLogger creates a new Logger instance (singleton pattern)
func NewLogger(config LoggerConfig) (*Logger, error) {
	once.Do(func() {
		instance = &Logger{
			config:  config,
			logChan: make(chan string, logChannelBufferSize),
		}
		if err := instance.rotateLogFile(); err != nil {
			instance = nil
		} else {
			go instance.writeLog()
		}
	})
	if instance == nil {
		return nil, fmt.Errorf("failed to create logger")
	}
	return instance, nil
}

// rotateLogFile rotates the log file, creates a new file and sets multi-output to console and file
func (l *Logger) rotateLogFile() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.logFile != nil {
		if err := l.logFile.Close(); err != nil {
			return err
		}
	}

	dir := filepath.Dir(l.config.FilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	newFilePath := l.config.FilePath + "." + time.Now().Format("20060102-150405")
	logFile, err := os.OpenFile(newFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	l.logFile = logFile
	multiWriter := io.MultiWriter(os.Stdout, l.logFile)
	l.log = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)

	go l.cleanupOldLogs()

	return nil
}

// cleanupOldLogs cleans up old log files based on configuration for deletion and compression operations
func (l *Logger) cleanupOldLogs() {
	files, err := filepath.Glob(l.config.FilePath + ".*")
	if err != nil {
		return
	}

	var logFiles []os.FileInfo

	for _, file := range files {
		if info, err := os.Stat(file); err == nil && !info.IsDir() {
			logFiles = append(logFiles, info)
		}
	}

	l.removeExcessBackups(logFiles)
	l.removeExpiredLogs(logFiles)
}

// removeExcessBackups deletes excess backup files, retaining up to MaxBackups files
func (l *Logger) removeExcessBackups(logFiles []os.FileInfo) {
	if len(logFiles) > l.config.MaxBackups {
		for _, file := range logFiles[:len(logFiles)-l.config.MaxBackups] {
			if err := os.Remove(filepath.Join(filepath.Dir(l.config.FilePath), file.Name())); err != nil {
				return
			}
		}
	}
}

// removeExpiredLogs deletes expired log files and compresses them based on configuration
func (l *Logger) removeExpiredLogs(logFiles []os.FileInfo) {
	for _, file := range logFiles {
		if time.Since(file.ModTime()).Hours() > float64(24*l.config.MaxAge) {
			if err := os.Remove(filepath.Join(filepath.Dir(l.config.FilePath), file.Name())); err != nil {
				return
			}
		}

		if l.config.Compress && filepath.Ext(file.Name()) == "" {
			if err := compressLog(filepath.Join(filepath.Dir(l.config.FilePath), file.Name())); err != nil {
				return
			}
		}
	}
}

// compressLog compresses the log file (example code, actual compression logic needs to be implemented)
func compressLog(filePath string) error {
	fmt.Println("Compressing:", filePath)
	return nil
}

// writeLog asynchronously writes logs to file and console
func (l *Logger) writeLog() {
	for msg := range l.logChan {
		l.mu.Lock()
		err := l.log.Output(2, msg)
		if err != nil {
			fmt.Printf("failed to write output: %v\n", err)
		}
		l.mu.Unlock()

		if l.logFile != nil && l.getSize() > l.config.MaxSize {
			_ = l.rotateLogFile()
		}
	}
}

// getSize gets the current size of the log file (bytes)
func (l *Logger) getSize() int64 {
	info, _ := l.logFile.Stat()
	return info.Size()
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level LogLevel) {
	l.config.Level = level
}

// Debug logs a message at DEBUG level
func (l *Logger) Debug(v ...interface{}) {
	if l.config.Level <= DEBUG {
		msg := fmt.Sprintln(v...)
		select {
		case l.logChan <- "[DEBUG] " + msg:
		default:
			fmt.Printf("log channel is full, dropping message: %s", msg)
		}
	}
}

// Info logs a message at INFO level
func (l *Logger) Info(v ...interface{}) {
	if l.config.Level <= INFO {
		msg := fmt.Sprintln(v...)
		select {
		case l.logChan <- "[INFO] " + msg:
		default:
			fmt.Printf("log channel is full, dropping message: %s", msg)
		}
	}
}

// Warn logs a message at WARN level
func (l *Logger) Warn(v ...interface{}) {
	if l.config.Level <= WARN {
		msg := fmt.Sprintln(v...)
		select {
		case l.logChan <- "[WARN] " + msg:
		default:
			fmt.Printf("log channel is full, dropping message: %s", msg)
		}
	}
}

// Error logs a message at ERROR level
func (l *Logger) Error(v ...interface{}) {
	if l.config.Level <= ERROR {
		msg := fmt.Sprintln(v...)
		select {
		case l.logChan <- "[ERROR] " + msg:
		default:
			fmt.Printf("log channel is full, dropping message: %s", msg)
		}
	}
}

// Close closes the log file handle and releases resources, and closes the buffer channel
func (l *Logger) Close() error {
	close(l.logChan)
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}
