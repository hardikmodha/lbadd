package raft

import "github.com/tomarrell/lbadd/internal/raft/message"

// AppendEntriesResponse function is called on a request from the leader to append log data
// to the follower node. This function generates the response to be sent to the leader node.
// This is the response to the contact by the leader to assert it's leadership.
func AppendEntriesResponse(node *Node, req *message.AppendEntriesRequest) *message.AppendEntriesResponse {
	success := true
	leaderTerm := req.GetTerm()
	nodePersistentState := node.PersistentState
	nodeTerm := nodePersistentState.CurrentTerm
	if nodeTerm > leaderTerm { // Reply false if term < currentTerm
		success = false
	} else if req.GetPrevLogIndex() > node.VolatileState.CommitIndex {
		success = false
	} else if nodePersistentState.Log[req.PrevLogIndex].Term != req.GetPrevLogTerm() {
		success = false
	}

	if !success {
		return &message.AppendEntriesResponse{
			Term:    nodeTerm,
			Success: success,
		}
	}

	entries := req.GetEntries()
	if len(entries) > 0 { // if heartbeat, skip adding entries
		nodePersistentState.mu.Lock()
		if req.GetPrevLogIndex() < node.VolatileState.CommitIndex {
			node.PersistentState.Log = node.PersistentState.Log[:req.GetPrevLogIndex()]
		}
		for _, entry := range entries {
			node.PersistentState.Log = append(node.PersistentState.Log, *entry)
		}
		node.PersistentState.mu.Unlock()
	}

	if req.GetLeaderCommit() > node.VolatileState.CommitIndex {
		nodeCommitIndex := req.GetLeaderCommit()
		if int(req.GetLeaderCommit()) > len(node.PersistentState.Log) {
			nodeCommitIndex = len(node.PersistentState.Log)
		}
		node.VolatileState.CommitIndex = nodeCommitIndex
		// apply the log command & update lastApplied
	}

	return &message.AppendEntriesResponse{
		Term:    nodeTerm,
		Success: success,
	}

}

