package logger

import (
	"log"
	"os"
	"sync"
	"time"
)

/**
 * Nivel de log
 */
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
}

type SimpleLogger struct {
	mu       sync.Mutex
	minLevel LogLevel
	logger   *log.Logger
}

func RequestLogger(minLevel LogLevel) Logger {
	return &SimpleLogger{
		minLevel: minLevel,
		// logger: log.New(os.Stdout, "", 0)
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (s *SimpleLogger) log(level LogLevel, prefix string, msg string) {
	if level < s.minLevel {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	s.logger.Printf("%s [%s] %s\n", timestamp, prefix, msg)
}

func (s *SimpleLogger) Debug(msg string) {
	s.log(DEBUG, "DEBUG", msg)
}

func (s *SimpleLogger) Info(msg string) {
	s.log(INFO, "INFO", msg)
}

func (s *SimpleLogger) Warn(msg string) {
	s.log(WARN, "WARN", msg)
}

func (s *SimpleLogger) Error(msg string) {
	s.log(ERROR, "ERROR", msg)
}

func (s *SimpleLogger) SetLevel(level LogLevel) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.minLevel = level
}

/**
logger := RequestLogger(INFO) // nivel mínimo INFO

logger.Debug("Este mensaje no se mostrará porque nivel mínimo es INFO")
logger.Info("Mensaje informativo")
logger.Warn("Advertencia")
logger.Error("Error crítico")

// Cambiar nivel a DEBUG para ver más detalles
logger.SetLevel(DEBUG)
logger.Debug("Ahora este mensaje sí se mostrará")
*/
