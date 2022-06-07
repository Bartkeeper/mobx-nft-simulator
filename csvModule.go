package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func importCSV() UserSet {

	csvUserGroup := UserSet{}
	csvFile, err := os.Open("stakers05-30-1.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {

		address := line[0]
		stakedMobx, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			fmt.Println(err)
		}
		collectedMiles, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			fmt.Println(err)
		}
		weight, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			fmt.Println(err)
		}
		milesWeight, err := strconv.ParseFloat(line[4], 64)
		if err != nil {
			fmt.Println(err)
		}
		rewardShare, err := strconv.ParseFloat(line[5], 64)
		if err != nil {
			fmt.Println(err)
		}
		mobxRewards, err := strconv.ParseFloat(line[6], 64)
		if err != nil {
			fmt.Println(err)
		}
		nftWeight, err := strconv.ParseFloat(line[7], 64)
		if err != nil {
			fmt.Println(err)
		}
		nftBonus, err := strconv.ParseFloat(line[8], 64)
		if err != nil {
			fmt.Println(err)
		}

		newUser := User{
			address:        address,
			stakedMobx:     stakedMobx,
			collectedMiles: collectedMiles,
			weight:         weight,
			milesWeight:    milesWeight,
			rewardShare:    rewardShare,
			mobxRewards:    mobxRewards,
			nftWeight:      nftWeight,
			nftBonus:       nftBonus,
		}

		csvUserGroup = append(csvUserGroup, newUser)

	}
	return csvUserGroup
}
