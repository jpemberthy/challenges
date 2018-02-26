package main

import (
	"bytes"
	"fmt"
	"strconv"
)

const anyToken = "."
const zeroOrMoreOfPreviousToken = "*"
const epsilon = "E*"

type NFA struct {
	initialState    string
	acceptingStates []string
	transitions     map[string]map[string][]string
}

func (nfa *NFA) String() string {
	var output bytes.Buffer
	output.WriteString(fmt.Sprintf("Initial: %s\n", nfa.initialState))
	output.WriteString("Transitions: \n")
	for state, _transitions := range nfa.transitions {
		for input, target_state := range _transitions {
			output.WriteString(fmt.Sprintf("%s ---(%s)---> %s\n", state, input, target_state))
		}
	}
	output.WriteString("Accepting states: ")
	for _, state := range nfa.acceptingStates {
		output.WriteString(fmt.Sprintf("%s - ", state))
	}
	output.WriteString("\n")
	return output.String()
}

// buildNFA builds a Nondeterministic Finite Automaton from a basic pattern.
func buildNFA(p string) *NFA {

	// map of the form:
	// { "1": { "a": ["2"], "E*": ["2"]} }
	nfa := &NFA{transitions: make(map[string]map[string][]string)}

	stateCount := 0
	for i := 0; i < len(p); i++ {
		char := string(p[i])

		// states and transitions associated with * are generated from previous chars.
		if char == zeroOrMoreOfPreviousToken {
			continue
		}

		stateKey := strconv.Itoa(stateCount)
		if i == 0 {
			nfa.initialState = stateKey
		}

		var nextChar string
		if i+1 < len(p) {
			nextChar = string(p[i+1])
		}

		// get ready for registering the transition.
		if nfa.transitions[stateKey] == nil {
			nfa.transitions[stateKey] = make(map[string][]string)
		}
		nextStateKey := strconv.Itoa(stateCount + 1)

		// if the next char is *, then we need to create a couple of Epsilon state transitions:
		// * One from the current state to the next one, and viceversa.
		if nextChar == zeroOrMoreOfPreviousToken {
			if nfa.transitions[nextStateKey] == nil {
				nfa.transitions[nextStateKey] = make(map[string][]string)
			}

			nfa.transitions[stateKey][epsilon] = []string{nextStateKey}
			nfa.transitions[nextStateKey][epsilon] = []string{stateKey}
		}

		// finally register this char transition to the next state.
		nfa.transitions[stateKey][char] = []string{nextStateKey}
		stateCount += 1

		// if the current char is the last one OR if the next one is * and it's the last one then we mark this one as an ending state.
		if i == len(p)-1 || i+1 == len(p)-1 && nextChar == zeroOrMoreOfPreviousToken {
			nfa.acceptingStates = append(nfa.acceptingStates, strconv.Itoa(stateCount))
		}
	}

	return nfa
}

func main() {
	pattern := "xy*"
	fmt.Printf("Pattern: %s\n", pattern)
	nfa := buildNFA(pattern)
	fmt.Println(nfa)
}
