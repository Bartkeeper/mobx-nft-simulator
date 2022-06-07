package main

import (
	"math"
)

type User struct {
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
type UserSet []User

func (us UserSet) DefineWeight() {

	// defines the weight based on the staked amount

	for i, mobxUser := range us {
		mobxUser.weight = math.Log(mobxUser.stakedMobx+1.1) / math.Log(1.25)
		us[i].weight = mobxUser.weight
	}
}

func (us UserSet) DefineMilesWeight() {
	totalMilesWeight = 0

	// calculates the miles weight based on the weight defined in the upper function, the collected miles and the NFT weight

	for i, mobxUser := range us {
		mobxUser.milesWeight = mobxUser.weight * mobxUser.collectedMiles * mobxUser.nftWeight
		us[i].milesWeight = mobxUser.milesWeight
		totalMilesWeight = totalMilesWeight + mobxUser.milesWeight
	}
}

func (us UserSet) DefineRatio() {

	// calculates the reward share based on the total amount of miles weight

	for i, mobxUser := range us {
		mobxUser.rewardShare = mobxUser.milesWeight / totalMilesWeight
		us[i].rewardShare = mobxUser.rewardShare
	}
}

func (us UserSet) DefineRewards() {

	// for each user, the real MOBX reward is calculated based on the reward share

	var rewardPool float64 = 10000

	for i, mobxUser := range us {
		mobxUser.mobxRewards = mobxUser.rewardShare * rewardPool
		us[i].mobxRewards = mobxUser.mobxRewards
	}
}

func (us UserSet) ResetNFT() {

	// sets the NFT weights of all users to 1
	// should only be called by a new userGroup

	for i, mobxUser := range us {
		mobxUser.nftWeight = 1
		us[i].nftWeight = mobxUser.nftWeight
	}
}

func (us UserSet) CalculateRewards() {
	us.DefineWeight()
	us.DefineMilesWeight()
	us.DefineRatio()
	us.DefineRewards()
}
