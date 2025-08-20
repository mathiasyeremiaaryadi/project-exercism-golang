package main

import "fmt"

type ElectionResult struct {
	// Name of the candidate
	Name string
	// Number of votes the candidate had
	Votes int
}

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
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
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{Name: candidateName, Votes: votes}
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}

var (
	initialVotes   int
	counter        *int
	votes          int
	voteCounter    *int
	nilVoteCounter *int
	result         *ElectionResult
)

func main() {
	initialVotes = 2

	counter = NewVoteCounter(initialVotes)
	fmt.Println(*counter == initialVotes)

	votes = 3

	voteCounter = &votes

	fmt.Println(VoteCount(voteCounter))
	// => 3

	fmt.Println(VoteCount(nilVoteCounter))
	// => 0

	result = NewElectionResult("Peter", 3)

	fmt.Println(result.Name == "Peter") // true
	fmt.Println(result.Votes == 3)      // true

	result = &ElectionResult{
		Name:  "John",
		Votes: 32,
	}

	DisplayResult(result)
	// => John (32)

	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}

	DecrementVotesOfCandidate(finalResults, "Mary")

	fmt.Println(finalResults["Mary"])
	// => 9
}
