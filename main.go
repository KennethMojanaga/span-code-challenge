package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"span-code-challenge/algorithms"
	"span-code-challenge/datastructures"
	"strconv"
	"strings"
)

func main() {
	inputStr := flag.String("input-string", "", "If the input is a string. For multiple match scores, the match scores should be separated by the \"\\n\" character. e.g. \"Lions 3, Snakes 3\\nLions 1, FC Awesome 1\"")
	inputFilePath := flag.String("input-file", "", "If the input is a file. This should be the path to the file. e.g. \"/home/johndoe/input_file.txt\"")

	flag.Parse()

	if *inputStr == "" && *inputFilePath == "" {
		println("Error!!! \nEither the -input-string flag OR the -input-file must be provided. For more details, use the --help flag")
		return
	}

	var matchesString string
	if *inputStr != "" {
		// We're processing a string input
		temp := *inputStr
		matchesString = strings.Replace(temp, `\n`, "\n", -1)
	}

	if *inputFilePath != "" {
		// We're processing a file input
		matchesString = readFile(*inputFilePath)
	}

	// Points map
	pointsMap := make(map[string]int)
	leagueMatchesArray := strings.Split(matchesString+"", "\n")
	//
	for _, leagueMatchString := range leagueMatchesArray {
		// Build scores
		teamToScoreMap := createTeamToScoreMap(leagueMatchString)

		// Build points
		teamToPointsMap := createTeamToPointsMap(teamToScoreMap)

		teamNames := make([]string, 0, len(teamToPointsMap))
		for key := range teamToPointsMap {
			teamNames = append(teamNames, key)
		}

		teamA := teamNames[0]
		teamB := teamNames[1]

		// Update team points
		pointsMap[teamA] = pointsMap[teamA] + teamToPointsMap[teamA]
		pointsMap[teamB] = pointsMap[teamB] + teamToPointsMap[teamB]
	}

	// Last bit is to sanitize the results in points order, and alphabetic order in the event of a draw
	orderedResult, _ := orderPoints(pointsMap)
	fmt.Printf("%v", orderedResult)
}

func createTeamToScoreMap(matchString string) map[string]int {
	scoresMap := make(map[string]int)

	teamAndScoreArr := strings.Split(matchString, ",")

	for _, teamAndScore := range teamAndScoreArr {
		// Extract team and score
		score := teamAndScore[len(teamAndScore)-1:]
		team := teamAndScore[0 : len(teamAndScore)-1]
		team = strings.TrimSpace(team)

		var err error
		scoresMap[team], err = strconv.Atoi(score)
		if err != nil {
			fmt.Printf("Error!!! Invalid score for team %v\n", team)
			os.Exit(1)
		}
	}

	return scoresMap
}

func createTeamToPointsMap(scoreMap map[string]int) map[string]int {
	pointsMap := make(map[string]int)

	// First get teamNames
	teamNames := make([]string, 0, len(scoreMap))
	for key := range scoreMap {
		teamNames = append(teamNames, key)
	}

	teamA := teamNames[0]
	teamB := teamNames[1]

	teamAScore := scoreMap[teamA]
	teamBScore := scoreMap[teamB]

	// ---- PROCESS TEAM POINTS
	// DRAW
	if teamAScore == teamBScore {
		pointsMap[teamA] = 1
		pointsMap[teamB] = 1
	}

	// TEAM A WIN
	if teamAScore > teamBScore {
		pointsMap[teamA] = 3
		pointsMap[teamB] = 0
	}

	// TEAM B WIN
	if teamAScore < teamBScore {
		pointsMap[teamA] = 0
		pointsMap[teamB] = 3
	}

	return pointsMap
}

// This function creates ordered results as specified in the spec
func orderPoints(pointsMap map[string]int) (string, []string) {
	outputResult := ""
	rankingResultsArr := []string{}

	pointsArray := []int{}

	// Group teams by points
	teamGroupingByPoints := make(map[int][]string)
	for teamName, teamPoints := range pointsMap {
		teamsGroupedByPoints := teamGroupingByPoints[teamPoints]
		teamsGroupedByPoints = append(teamsGroupedByPoints, teamName)

		teamGroupingByPoints[teamPoints] = teamsGroupedByPoints

		pointsArray = append(pointsArray, teamPoints)
	}

	// Convert to Set (to eliminate duplicates/draws)
	set := datastructures.SetInt{}
	setArray := set.Init(pointsArray)

	// Sort
	setArray = algorithms.IntMergeSort(setArray)
	setArray = algorithms.ReverseIntArray(setArray)

	for _, point := range setArray {
		teamsGroupedByPoints := teamGroupingByPoints[point]

		if len(teamsGroupedByPoints) == 1 {
			// Then its the only team with that point
			team := teamsGroupedByPoints[0]

			ptsPluralForm := " pt"
			if point > 1 || point == 0 {
				ptsPluralForm = " pts"
			}

			rankingResultsArr = append(rankingResultsArr, team+", "+fmt.Sprint(point)+ptsPluralForm)
		}

		if len(teamsGroupedByPoints) > 1 {
			// Then its a draw situation. So order the teams by name alphabetical order
			sort.Strings(teamsGroupedByPoints)

			for _, team := range teamsGroupedByPoints {
				ptsPluralForm := " pt"
				if point > 1 || point == 0 {
					ptsPluralForm = " pts"
				}
				rankingResultsArr = append(rankingResultsArr, team+", "+fmt.Sprint(point)+ptsPluralForm)
			}
		}
	}

	// Prepare output results
	for index, result := range rankingResultsArr {
		result = fmt.Sprint(index+1) + ". " + result + "\n"
		outputResult += result
	}

	return outputResult, rankingResultsArr
}
