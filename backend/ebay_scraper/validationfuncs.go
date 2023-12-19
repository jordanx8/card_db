package main

import (
	"slices"
	"strings"
)

func insertValidation(l Listing, cs CardSearch) bool {
	return strings.Contains(strings.ToLower(l.Title), strings.ToLower(cs.Insert))
}

func playerNameAndYearValidation(l Listing, cs CardSearch) bool {
	title := strings.ToLower(l.Title)
	if strings.Contains(title, cs.Year) && strings.Contains(title, strings.ToLower(cs.Player.FirstName)) && strings.Contains(title, strings.ToLower(cs.Player.LastName)) {
		return true
	}
	return false
}

func setValidation(l Listing, cs CardSearch) bool {
	manufacturers := []string{"Panini"}

	setStringSplit := strings.Split(cs.Set, " ")

	setStringArray := filterOutOfArray(setStringSplit, manufacturers)

	title := strings.ToLower(l.Title)
	for _, val := range setStringArray {
		if !strings.Contains(title, strings.ToLower(val)) {
			return false
		}
	}

	return true
}

func baseValidation(l Listing, cs CardSearch) bool {
	if strings.Contains(strings.ToLower(l.Title), "base") {
		return true
	} else if checkDetails(l, "Parallel/Variety", "Base") {
		return true
	} else if checkDetails(l, "Features", "Base") {
		return true
	} else {
		filteredTitle := strings.Replace(l.Title, cs.Set, "", -1)
		filteredTitle = strings.Replace(filteredTitle, cs.Parallel, "", -1)

		parallelWords := []string{"Prizm", "Mosaic", "Parallel", "Reactive"}
		for _, val := range parallelWords {
			if strings.Contains(strings.ToLower(filteredTitle), strings.ToLower(val)) {
				return false
			}
		}
	}
	return true
}

func parallelValidation(l Listing, cs CardSearch) bool {
	wordsToFilterOut := []string{"Prizm", "Mosaic", "Mojo"}
	parallelStringSplit := strings.Split(cs.Parallel, " ")
	parallelStringArray := filterOutOfArray(parallelStringSplit, wordsToFilterOut)
	filteredTitle := strings.Replace(l.Title, cs.Set, "", -1)
	title := strings.ToLower(filteredTitle)

	for _, val := range parallelStringArray {
		// looking for color/other descriptors here
		if !strings.Contains(title, strings.ToLower(val)) {
			return false
		}
	}

	for _, val := range parallelStringSplit {
		title = strings.Replace(filteredTitle, val, "", -1)
	}

	title = strings.ToLower(title)
	for _, val := range wordsToFilterOut {
		if strings.Contains(title, strings.ToLower(val)) {
			return false
		}
	}

	return true
}

func checkDetails(l Listing, detail string, comparison string) bool {
	val, ok := l.Details[detail]
	if ok {
		if strings.Contains(val, comparison) {
			return true
		}
	}
	return false
}

func filterOutOfArray(arrayToBeFiltered []string, filterWords []string) []string {
	var returnArray []string
	for _, val := range arrayToBeFiltered {
		if !slices.Contains(filterWords, val) {
			returnArray = append(returnArray, val)
		}
	}
	return returnArray
}
