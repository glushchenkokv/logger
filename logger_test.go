package logger

import (
	"testing"
)

func TestDefaultConfigLogger(t *testing.T) {
	log := NewLogger(Config{})

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
}

func TestTextLogger(t *testing.T) {
	log := NewLogger(Config{
		Level:            "trace",
		Format:           "text",
		DisableTimestamp: false,
	})

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
}

func TestJsonLogger(t *testing.T) {
	log := NewLogger(Config{
		Level:            "trace",
		Format:           "json",
		DisableTimestamp: false,
	})

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
}

func TestFatalLogger(t *testing.T) {
	t.Skip()
	log := NewLogger(Config{
		Level:            "trace",
		Format:           "json",
		DisableTimestamp: false,
	})

	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
}

func TestPanicLogger(t *testing.T) {
	t.Skip()
	log := NewLogger(Config{
		Level:            "trace",
		Format:           "json",
		DisableTimestamp: false,
	})

	// Calls panic() after logging
	log.Panic("I'm bailing.")
}
