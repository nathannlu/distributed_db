package raft


import (
	"fmt"
)

type LogEntry struct {
	Index int64
	Term int64
	Data string
}

type Log struct {
	Entries []LogEntry

}

// Find log at index
func (l *Log) GetLogEntry(index int64) (*LogEntry) {

	// Check if index exist in Log.Entries
	fmt.Println(int64(len(l.Entries)), index)
	if int64(len(l.Entries)) < index {
		return nil;
	}


	fmt.Println(l.Entries)
	// If index exist in log, find index and return
	e := l.Entries[index]
	
	return &e
}

// Append log at index
func AppendLog(log []LogEntry) {
	
}

// Deletes range of log entries
func DeleteRange(min, max int64) {
	
}
