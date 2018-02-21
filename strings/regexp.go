// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.

// The matching should cover the entire input string (not partial).

// The function prototype should be:
// bool isMatch(const char *s, const char *p)

// Some examples:
// isMatch("aa","a") → false
// isMatch("aa","aa") → true
// isMatch("aaa","aa") → false
// isMatch("aa", "a*") → true
// isMatch("aa", ".*") → true
// isMatch("ab", ".*") → true
// isMatch("aab", "c*a*b") → true
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
	initial     *State
	transitions map[string]map[string]string
}

func (nfa *NFA) String() string {
	var output bytes.Buffer
	output.WriteString(fmt.Sprintf("Initial: %s\n", nfa.initial.key))
	output.WriteString("Transitions: \n")
	for state, _transitions := range nfa.transitions {
		for input, target_state := range _transitions {
			output.WriteString(fmt.Sprintf("%s ---(%s)---> %s\n", state, input, target_state))
		}
	}

	return output.String()
}

type State struct {
	key       string
	initial   bool
	accepting bool
}

// buildNFA builds a Nondeterministic Finite Automaton from a basic pattern.
func buildNFA(p string) *NFA {

	// map of the form:
	// { "1": { "a": "2", "E*": "2"} }
	nfa := &NFA{transitions: make(map[string]map[string]string)}

	for i := 0; i < len(p); i++ {
		char := string(p[i])

		// states and transitions associated with * are generated from previous chars.
		if char == zeroOrMoreOfPreviousToken {
			continue
		}

		stateKey := strconv.Itoa(i)
		s := &State{key: stateKey}

		if i == 0 {
			s.initial = true
			nfa.initial = s
		}

		var nextChar string
		if i+1 < len(p) {
			nextChar = string(p[i+1])
		}

		// if the current char is the last one OR if the next one is * and it's the last one then we mark this one as an ending state.
		if i == len(p)-1 || i+1 == len(p)-1 && nextChar == zeroOrMoreOfPreviousToken {
			s.accepting = true
		}

		// get ready for registering the transition.
		if nfa.transitions[stateKey] == nil {
			nfa.transitions[stateKey] = make(map[string]string)
		}
		nextStateKey := strconv.Itoa(i + 1)

		// if the next char is *, then we need to create a couple of Epsilon state transitions:
		// * One from the current state to the next one, and viceversa.
		if nextChar == zeroOrMoreOfPreviousToken {
			if nfa.transitions[nextStateKey] == nil {
				nfa.transitions[nextStateKey] = make(map[string]string)
			}

			nfa.transitions[stateKey][epsilon] = nextStateKey
			nfa.transitions[nextStateKey][epsilon] = stateKey
		}

		// finally register this char transition to the next state.
		nfa.transitions[stateKey][char] = nextStateKey
	}

	// TODO: get rid of State{} in favor of just NFA attributes.
	return nfa
}

// isMatch checks whether or not s is a full match of p.
//
// Internally, p is converted to a NFA and then to a DFA.
// Then we use each of the chars in the string to traverse the DFA and
// if the final state is an accepting state, then s matches p.
func isMatch(s string, p string) bool {
	verbose := true
	if verbose {
		fmt.Printf("String: %s | Pattern: %s\n", s, p)
	}

	return true
}

func main() {
	nfa := buildNFA("ab*c")
	fmt.Println(nfa)
	// fmt.Println(isMatch("aaa", "aaaa") == false)
	// fmt.Println(isMatch("a", "ab*") == true)
	// fmt.Println(isMatch("aaa", "ab*a") == false)
	// fmt.Println(isMatch("ab", ".*c") == false)
	// fmt.Println(isMatch("aa", "a") == false)
	// fmt.Println(isMatch("aa", "aa") == true)
	// fmt.Println(isMatch("aaa", "aa") == false)
	// fmt.Println(isMatch("aa", "a*") == true)
	// fmt.Println(isMatch("aa", ".*") == true)
	// fmt.Println(isMatch("ab", ".*") == true)
	// fmt.Println(isMatch("aab", "c*a*b") == true)
	// fmt.Println(isMatch("aba", ".*a") == true)
	// fmt.Println(isMatch("aaa", "a*a") == true)
	// fmt.Println(isMatch("bbbba", ".*a*a") == true)
}
