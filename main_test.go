package main

import "testing"

func TestCreateTeamToScoreMap(t *testing.T) {
	testString := "Team A 4, Team B 2"

	teamToScoreMap := createTeamToScoreMap(testString)

	teamAScore, teamAExists := teamToScoreMap["Team A"]
	teamBScore, teamBExists := teamToScoreMap["Team B"]

	// Test for key presence
	if !teamAExists {
		t.Errorf("TestCreateTeamToScoreMap failed. Expected \"Team A\" to be present as a key in the \"teamToScoreMap\"")
	}
	if !teamBExists {
		t.Errorf("TestCreateTeamToScoreMap failed. Expected \"Team B\" to be present as a key in the \"teamToScoreMap\"")
	}

	// Test for valid score values
	if teamAScore != 4 {
		t.Errorf("TestCreateTeamToScoreMap failed. Expected \"Team A\" score to be 4")
	}
	if teamBScore != 2 {
		t.Errorf("TestCreateTeamToScoreMap failed. Expected \"Team B\" score to be 2")
	}

	t.Logf("TestCreateTeamToScoreMap ran succefully")
}

func TestCreateTeamToPointsMap(t *testing.T) {
	testString := "Team A 4, Team B 2"

	teamToScoreMap := createTeamToScoreMap(testString)
	teamToPointsMap := createTeamToPointsMap(teamToScoreMap)

	teamAPoints, teamAExists := teamToPointsMap["Team A"]
	teamBPoints, teamBExists := teamToPointsMap["Team B"]

	// Test for key presence
	if !teamAExists {
		t.Errorf("TestCreateTeamToPointsMap failed. Expected \"Team A\" to be present as a key in the \"teamToPointsMap\"")
	}
	if !teamBExists {
		t.Errorf("TestCreateTeamToPointsMap failed. Expected \"Team B\" to be present as a key in the \"teamToPointsMap\"")
	}

	// Test for valid point values
	if teamAPoints != 3 {
		t.Errorf("TestCreateTeamToPointsMap failed. Expected \"Team A\" points to be 3")
	}
	if teamBPoints != 0 {
		t.Errorf("TestCreateTeamToPointsMap failed. Expected \"Team B\" points to be 0")
	}

	t.Logf("TestCreateTeamToPointsMap ran succefully")
}

func TestOrderTeamPoints(t *testing.T) {
	teamToPointsMap := make(map[string]int)
	teamToPointsMap["B Team"] = 4
	teamToPointsMap["D Team"] = 6
	teamToPointsMap["A Team"] = 1
	teamToPointsMap["C Team"] = 4

	_, orderedResults := orderPoints(teamToPointsMap)

	// Test for results length presence
	if len(orderedResults) != 4 {
		t.Errorf("TestOrderTeamPoints failed. Expected \"orderedResults\" length to be 4")
	}

	// Test for valid ranking
	if orderedResults[0] != "D Team, 6 pts" {
		t.Errorf("TestOrderTeamPoints failed. Expected \"orderedResults[0]\" to be \"D Team, 6 pts\"")
	}
	if orderedResults[1] != "B Team, 4 pts" {
		t.Errorf("TestOrderTeamPoints failed. Expected \"orderedResults[0]\" to be \"B Team, 4 pts\"")
	}
	if orderedResults[2] != "C Team, 4 pts" {
		t.Errorf("TestOrderTeamPoints failed. Expected \"orderedResults[0]\" to be \"C Team, 4 pts\"")
	}
	if orderedResults[3] != "A Team, 1 pt" {
		t.Errorf("TestOrderTeamPoints failed. Expected \"orderedResults[0]\" to be \"A Team, 1 pt\"")
	}

	t.Logf("TestOrderTeamPoints ran succefully")
}
