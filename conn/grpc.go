package conn;

import (
	"fmt"
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/nathannlu/distributed_db/pb";
	"github.com/nathannlu/distributed_db/raft";
)




type Server struct {
	Pool *Pool
	RaftState *raft.RaftState
}


func StartGrpcServer() {
	fmt.Println("Starting GRPC server")

	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	p := &Pool{}
	r := raft.RaftInit()
	s := Server{
		Pool: p,
		RaftState: &r,
	}



	grpcServer := grpc.NewServer()
	pb.RegisterRaftServer(grpcServer, &s)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}


// Raft membership management
func (s *Server) AddServer(ctx context.Context, req *pb.AddServerRequest) (*pb.AddServerResponse, error) {

	newNode := NodeConnection{
		Name: req.Name,
	}
	s.Pool.AddNode(newNode)

	return &pb.AddServerResponse{
		Message: "aasd",
	}, nil
}

func (s *Server) ListServers(ctx context.Context, req *pb.ListServersRequest) (*pb.ListServersResponse, error) {

	s.Pool.GetAllNodes()
	
	return &pb.ListServersResponse{
		Message: "aasd",
	}, nil
}

// Raft core implementation
func (s *Server) AppendEntries(ctx context.Context, req *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {

	resp := &pb.AppendEntriesResponse{
		Term: s.RaftState.CurrentTerm,
		Success: false,
	}

	// Get previous log
	receiverPrevLogEntry := s.RaftState.Log.GetLogEntry(req.PrevLogIndex)
	receiverLogLength := int64(len(s.RaftState.Log.Entries))
	if receiverLogLength < 0 {
		fmt.Println("calculated index is less than 0")
	}
	receiverLastLogEntry := s.RaftState.Log.GetLogEntry(
		receiverLogLength - 1,
	)
	if receiverPrevLogEntry == nil {
		// cannot find log entry
		fmt.Println("Previous log entry does not exist")
		return resp, nil
	}


	// Ignore past term
	if req.Term < s.RaftState.CurrentTerm {
		return resp, nil
	}
	// For all servers: if RPC request contains term T > currentTerm:
	// set currentTerm = T, convert to follower 
	if req.Term > s.RaftState.CurrentTerm {
		// convert to follower
		s.RaftState.CurrentTerm = req.Term
		s.RaftState.State = raft.Follower
	}

	// If Entries length is 0, request is just a heartbeat
	if len(req.Entries) == 0 {
		resp.Success = true
		return resp, nil
	}

	//
	// Process entries
	//

	// monkey patch - handle server's last log index
	if req.PrevLogIndex == 0 {
		toBeCommited := raft.LogEntry {
			Index: req.Entries[0].Index,
			Term: req.Entries[0].Term,
			Data: req.Entries[0].Data,
		}
		s.RaftState.Log.Entries = append(s.RaftState.Log.Entries, toBeCommited)
		return resp, nil
	}

	// Leader previous log index does not match RPC receiver's index
	if req.PrevLogIndex != receiverPrevLogEntry.Index {
		fmt.Println("No commit. Previous log index does not match received log index")

		return resp, nil
	}
	// Leader previous log term does not match RPC receiver's term
	if req.PrevLogTerm != receiverPrevLogEntry.Term {
		fmt.Println("No commit. Previous log term does not match received log term")
		return resp, nil
	}

	// ** everything after this means it matches **
	// if index is not at the end of list, cut everything after the list
	if req.PrevLogIndex != receiverLastLogEntry.Index {
		// cut commit log
		s.RaftState.Log.Entries = s.RaftState.Log.Entries[:req.PrevLogIndex + 1]
		
		fmt.Println("Log has excess entries, cutting log")
		fmt.Println("Log before cut:", s.RaftState.Log.Entries)
		fmt.Println("Log after cut:", s.RaftState.Log.Entries)

		// then append
	}

	// Passed all checks; append log
	toBeCommited := raft.LogEntry {
		Index: req.Entries[0].Index,
		Term: req.Entries[0].Term,
		Data: req.Entries[0].Data,
	}
	fmt.Println("---")
	fmt.Println("Before:", s.RaftState.Log.Entries)
	s.RaftState.Log.Entries = append(s.RaftState.Log.Entries, toBeCommited)

	fmt.Println("After:", s.RaftState.Log.Entries)

	// Everything went well, set success
	return resp, nil
}


func (s *Server) RequestVote(ctx context.Context, req *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {

	if(req.Term < s.RaftState.CurrentTerm) {
		return &pb.RequestVoteResponse{
			Term: s.RaftState.CurrentTerm,
			VoteGranted: false,
		}, nil
	}
	
	// Check if server has already voted

	s.RaftState.VotedFor = req.CandidateId

	return &pb.RequestVoteResponse{
		Term: s.RaftState.CurrentTerm,
		VoteGranted: true,
	}, nil
}


