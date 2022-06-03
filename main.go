package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

// ToDo - Split up structs and functions into seperate go files of the same package

type user struct {
	address        string
	stakedMobx     float64
	collectedMiles float64
	weight         float64
	milesWeight    float64
	rewardShare    float64
	mobxRewards    float64
	nftWeight      float64
	nftBonus       float64
}
type userSet []user

type nfTier struct {
	tierClass   float64
	address     string
	mobxRewards float64
	nftWeight   float64
	nftBonus    float64
}

type nftSet []nfTier

var totalMilesWeight float64

type nftMap struct {
	tier       float64
	multiplier float64
	amount     float64
}

type mapping []nftMap

func main() {
	nftProps := setNFTProps(mapping{})
	userGroup := importCSV()
	userGroup.manipulateUsers(nftProps)
	userGroup2 := importCSV()

	userGroup2.resetNFT()
	userGroup.calculateRewards()
	userGroup2.calculateRewards()
	userGroup.calculateNFTbonus(userGroup2, nftProps)

}

func (us userSet) defineWeight() {

	// defines the weight based on the staked amount

	for i, mobxUser := range us {
		mobxUser.weight = math.Log(mobxUser.stakedMobx+1.1) / math.Log(1.25)
		// fmt.Println("The weight of user", i, "is", mobxUser.weight)
		us[i].weight = mobxUser.weight
	}
}

func (us userSet) defineMilesWeight() {
	totalMilesWeight = 0

	// calculates the miles weight based on the weight defined in the upper function, the collected miles and the NFT weight

	for i, mobxUser := range us {
		mobxUser.milesWeight = mobxUser.weight * mobxUser.collectedMiles * mobxUser.nftWeight
		// fmt.Println("The miles weight of user", i, "is", mobxUser.milesWeight)
		us[i].milesWeight = mobxUser.milesWeight
		totalMilesWeight = totalMilesWeight + mobxUser.milesWeight
	}
}

func (us userSet) defineRatio() {

	// calculates the reward share based on the total amount of miles weight

	for i, mobxUser := range us {
		mobxUser.rewardShare = mobxUser.milesWeight / totalMilesWeight
		// fmt.Println("The reward share of user", i, "is", mobxUser.rewardShare)
		us[i].rewardShare = mobxUser.rewardShare
	}
}

func (us userSet) defineRewards() {

	// for each user, the real MOBX reward is calculated based on the reward share

	var rewardPool float64 = 10000

	for i, mobxUser := range us {
		mobxUser.mobxRewards = mobxUser.rewardShare * rewardPool
		// fmt.Println("The mobx rewards of user", i, "is", mobxUser.mobxRewards)
		us[i].mobxRewards = mobxUser.mobxRewards
	}
}

func (us userSet) resetNFT() {

	// sets the NFT weights of all users to 1
	// should only be called by a new userGroup

	for i, mobxUser := range us {
		mobxUser.nftWeight = 1
		// fmt.Println("The mobx rewards of user", i, "is", mobxUser.nftWeight)
		us[i].nftWeight = mobxUser.nftWeight
	}
}

func (us userSet) calculateRewards() {
	us.defineWeight()
	us.defineMilesWeight()
	us.defineRatio()
	us.defineRewards()
}

func getNFTRewards2(us userSet, nft nftSet) []nfTier {

	// Requires a set of NFT tiers and then only picks one account of each tier
	// returns a set of accounts with distinct NFT tiers

	// ToDo: Performance increase - only the first value should be entered and not loop through the entire userset
	// fmt.Println("When entering the nft set looks like this", nft)
	for _, mobxUser := range us {
		if mobxUser.collectedMiles != 0 {

			for i, nftUser := range nft {
				// fmt.Println(nftUser.nftWeight)
				if nftUser.address == "" && nftUser.nftWeight == mobxUser.nftWeight {

					tier := nftSet{{nftUser.tierClass, mobxUser.address, mobxUser.mobxRewards, mobxUser.nftWeight, mobxUser.nftBonus}}
					// fmt.Println("Ardians Tier: ", tier)

					nft[i] = tier[0]
					// fmt.Println(nft[i])
					break
				}

				// switch nftUser.tierClass {
				// case 1:
				// 	if nftUser.address == "" {
				// 		tier1 := nftSet{{1, mobxUser.address, mobxUser.mobxRewards, mobxUser.nftWeight, mobxUser.nftBonus}}
				// 		nft[0] = tier1[0]
				// 	}
				// case 1.5:
				// 	if nftUser.address == "" {
				// 		tier15 := nftSet{{1.5, mobxUser.address, mobxUser.mobxRewards, mobxUser.nftWeight, mobxUser.nftBonus}}
				// 		nft[1] = tier15[0]
				// 	}
				// case 2:
				// 	if nftUser.address == "" {
				// 		tier2 := nftSet{{2, mobxUser.address, mobxUser.mobxRewards, mobxUser.nftWeight, mobxUser.nftBonus}}
				// 		nft[2] = tier2[0]
				// 	}
				// case 3:
				// 	if nftUser.address == "" {
				// 		tier3 := nftSet{{3, mobxUser.address, mobxUser.mobxRewards, mobxUser.nftWeight, mobxUser.nftBonus}}
				// 		nft[3] = tier3[0]
				// 	}
				// default:
				// 	fmt.Println("default")
				// }
			}
		}
	}

	// fmt.Println(nft)

	return nft
}

func (us userSet) calculateNFTbonus(us2 userSet, m mapping) {

	// (currently) calculates the difference of the rewards compared to a world without NFT bonus
	// prints out final results

	// nftGroup := nftSet{{1, "", 0, 0, 0}, {1.5, "", 0, 0, 0}, {2, "", 0, 0, 0}, {3, "", 0, 0, 0}}

	nftGroup := nftSet{}
	class1 := nfTier{
		tierClass:   1,
		address:     "",
		mobxRewards: 0,
		nftWeight:   1,
		nftBonus:    0,
	}
	nftGroup = append(nftGroup, class1)

	for _, m := range m {
		newTier := nfTier{
			tierClass:   m.tier,
			address:     "",
			mobxRewards: 0,
			nftWeight:   m.multiplier,
			nftBonus:    0,
		}

		nftGroup = append(nftGroup, newTier)

	}

	// fmt.Println("The entire Tier set looks like this: ", nftGroup)

	nft := getNFTRewards2(us, nftGroup)
	fmt.Println("++++++++++++++++++++++++++++++++ Results ++++++++++++++++++++++++++++++++")
	for i, nfTier := range nft {
		for _, mobxUser := range us2 {
			if mobxUser.address == nfTier.address {

				nft[i].nftBonus = (((nfTier.mobxRewards / mobxUser.mobxRewards) - 1) * 100)
				effectiveBonus := int(nft[i].nftBonus - nft[0].nftBonus)

				fmt.Println("The NFT bonus with Multiplier", nfTier.nftWeight, " is:", nft[i].nftBonus, "%")
				fmt.Println("By buying an NFT with Multiplier", nfTier.nftWeight, ", the Rewards increased by", effectiveBonus, "%")
			}
		}
	}

}

func importCSV() userSet {

	csvUserGroup := userSet{}
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

		newUser := user{
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

func setNFTProps(m mapping) mapping {
	fmt.Println("How many different tiers do you want? (excluding non-NFT holders)")

	// reader := bufio.NewReader(os.Stdin)

	var f4 float64
	_, _ = fmt.Scanf("%f4", &f4)

	numberOfNFTs := int(f4)

	for j := 0; j < numberOfNFTs; j++ {

		var f float64
		var f2 float64
		var f3 float64

		fmt.Println("For Nft Tier", j, ", please enter the tier identifier as a number. DON'T USE 1")
		_, _ = fmt.Scanf("%f", &f)

		fmt.Println("For NFT Tier", j, ", please enter the multiplier as a number")
		_, _ = fmt.Scanf("%f2", &f2)

		fmt.Println("For Nft Tier", j, ", please enter the issued NFT amount")
		_, _ = fmt.Scanf("%f3", &f3)

		tierF := float64(f)
		multiplierF := float64(f2)
		amountF := float64(f3)

		m2 := nftMap{tierF, multiplierF, amountF}

		m = append(m, m2)
	}
	// fmt.Println("The Tier set looks like this: ", m)
	return m
}

func (us userSet) manipulateUsers(m mapping) {

	for _, nftTier := range m {

		ln := len(us)
		am := int(nftTier.amount)

		for l := 0; l <= am; l++ {

			// does print out the same addresses
			user := int(rand.Intn(ln))
			us[user].nftWeight = nftTier.multiplier
		}

	}

}
