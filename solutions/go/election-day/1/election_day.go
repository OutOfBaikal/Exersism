package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
    if initialVotes < 0 {
		panic("Initial votes cannot be negative")
	}
    return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    if counter == nil {
		return 0
	}
    return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
    if counter == nil {
		*counter = 0
	}
	if increment < 0 {
		panic("Increment value cannot be negative")
	}
    *counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
    if candidateName == "" {
		panic("Candidate name cannot be empty")
	}
	if votes < 0 {
		panic("Votes cannot be negative")
	}
	return &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
	panic("Please implement the NewElectionResult() function")
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
    if result == nil {
		panic("Election result is nil")
	}
    return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
	panic("Please implement the DisplayResult() function")
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
    if results == nil {
		panic("Results map is nil")
	}
	if candidate == "" {
		panic("Candidate name cannot be empty")
	}
	if _, exists := results[candidate]; !exists {
		panic("Candidate not found in results")
	}
    results[candidate]--
}
