package raft

import (
	"fmt"
)

type RaftNodeRole int64;

const (
	Follower RaftNodeRole = iota
	Candidate
	Leader
)

type Raft struct {}

type RaftState struct {
	// Persistent state
	CurrentTerm int64
	VotedFor string
	Log *Log

	/*
	// Volatile state
	commitIndex int
	lastApplied int

	// Leader only states
	nextIndex []string
	matchIndex []string
	*/

	State RaftNodeRole
}

func RaftInit() (RaftState) {

	l := Log{
		Entries: make([]LogEntry, 0),
	}


	s := RaftState{
		// Persistent state
		CurrentTerm: 0,
		VotedFor: "",
		Log: &l,
		State: Follower,
	}

	// All nodes start in follower state
//	s.RaftNodeRole = Follower;

	// heartbeat := timer.random
	// electionTimeout := timer.random


	// if node is in follower state
	// send requestvote if electiontimeout passes
	
		// if follower node receives an appendEntries RPC (heartbeat),
		// electiontimeout resets
		

	// if node is in leader state
	// send appendEntries on heartbeat timer


	fmt.Println("Initialized Raft state")
	return s
}


/*
func (r *Raft) AppendEntries(term int64, leaderId int64, prevLogIndex int64, prevLogTerm int64, entries *Log) (int64, bool) {

	// 1. Reply false if term < currentTerm
	if req.Term < s.RaftState.currentTerm {
		return s.RaftState.currentTerm, false
	}
	// For all servers: if RPC request contains term T > currentTerm:
	// set currentTerm = T, convert to follower 
	if req.Term > s.RaftState.currentTerm {
		// convert to follower
		s.RaftState.currentTerm = req.Term
		s.RaftState.state = Follower
	}

	// Verify last log entry
	if req.prevLogEntry > 0 {
		// 2.1. Check if log doesn’t contain an entry at prevLogIndex
		if s.RaftState.log[req.PrevLogIndex] == nil {
			// return false
			return s.RaftState.currentTerm, false
		}

		// 2.2 Check if entry term matches prevLogTerm
		if s.RaftState.log[req.PrevLogIndex].term != req.PrevLogTerm {
			// return false
			return s.RaftState.currentTerm, false
		}
	}


	// Process entries
	// If Entries length is 0, request is just a heartbeat
	if len(req.Entries) > 0 {
		// 3. If an existing entry conflicts with a new one (same index
		// but different terms, delete the existing entry and all that
		// follow it
		

		// server's last log index
		lastLogIndex := len(s.RaftState.log) - 1
		lastLogTerm := s.RaftState.log[lastLogIndex].term

		if lastLogIndex != prevLogIndex {
			// mismatch
			
		}
		if lastLogTerm != prevLogTerm {
			// mismatch
			
		}

		if lastLogIndex == prevLogIndex && lastLogTerm == prevLogTerm {
			s.RaftState.log = append(s.RaftState.log, entries)
			
		}


		



		for i := 0; i < len(req.Entries); i++ {
			if req.Entries.index > lastLogIndex {
				toBeAppended = append(toBeAppended, req.Entries[i])	
			}
		}

		// 4. Append any new entries not already in the log
		s.RaftState.log = append(s.RaftState.log, toBeAppended)
	}

	
	// Everything went well, set success
	return s.RaftState.currentTerm, true
}


*/
