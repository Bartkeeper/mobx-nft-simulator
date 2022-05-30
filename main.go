package main

import (
	"fmt"
	"math"
)

type user struct {
	address        string
	stakedMobx     float64
	collectedMiles float64
	weight         float64
	milesWeight    float64
	rewardShare    float64
	mobxRewards    float64
	nftWeight      float64
}
type userSet []user

type nfTier struct {
	tierClass   float64
	address     string
	mobxRewards float64
	nftWeight   float64
}

type nftSet []nfTier

var totalMilesWeight float64

func main() {

	userGroup := newUser()
	userGroup2 := newUser()

	userGroup2.resetNFT()
	userGroup.calculateRewards()
	userGroup2.calculateRewards()
	userGroup.calculateNFTbonus(userGroup2)
}

func newUser() userSet {

	user1 := user{
		address:        "fetch1034pkj6fcm6te04vfq9d6qcm6493xa7dacswvh",
		stakedMobx:     100,
		collectedMiles: 100,
		nftWeight:      1.5,
	}

	user2 := user{
		address:        "fetch105zhyy3lyjqhxdtllmz4rmp57gqmzxrpd5qz0q",
		stakedMobx:     100,
		collectedMiles: 100,
		nftWeight:      1,
	}

	user3 := user{
		address:        "fetch106jc99nlh5jspd80q4xnv69d63qc9eg4m0sc2x",
		stakedMobx:     100,
		collectedMiles: 100,
		nftWeight:      2,
	}

	user4 := user{
		address:        "fetch1027maq7mdtaxa5wan00f0f5nmt70nz933z6vd5",
		stakedMobx:     100,
		collectedMiles: 100,
		nftWeight:      3,
	}

	user5 := user{
		address:        "fetch103ngv5cftngje4yyhe5qkmp9adgdpsvy4fnkwz",
		stakedMobx:     100,
		collectedMiles: 100,
		nftWeight:      1.5,
	}

	user6 := user{
		address:        "fetch102ntrhyxpfeyfc5kl0wam3ehmzz4atc52t0ddf",
		stakedMobx:     100,
		collectedMiles: 100,
		nftWeight:      1.5,
	}

	userGroup := userSet{user1, user2, user3, user4, user5, user6}

	return userGroup
}

func (us userSet) defineWeight() {

	for i, mobxUser := range us {
		mobxUser.weight = math.Log(mobxUser.stakedMobx+1.1) / math.Log(1.25)
		// fmt.Println("The weight of user", i, "is", mobxUser.weight)
		us[i].weight = mobxUser.weight
	}
}

func (us userSet) defineMilesWeight() {
	totalMilesWeight = 0

	for i, mobxUser := range us {
		mobxUser.milesWeight = mobxUser.weight * mobxUser.collectedMiles * mobxUser.nftWeight
		// fmt.Println("The miles weight of user", i, "is", mobxUser.milesWeight)
		us[i].milesWeight = mobxUser.milesWeight
		totalMilesWeight = totalMilesWeight + mobxUser.milesWeight
	}
}

func (us userSet) defineRatio() {

	for i, mobxUser := range us {
		mobxUser.rewardShare = mobxUser.milesWeight / totalMilesWeight
		// fmt.Println("The reward share of user", i, "is", mobxUser.rewardShare)
		us[i].rewardShare = mobxUser.rewardShare
	}
}

func (us userSet) defineRewards() {
	var rewardPool float64 = 10000

	for i, mobxUser := range us {
		mobxUser.mobxRewards = mobxUser.rewardShare * rewardPool
		// fmt.Println("The mobx rewards of user", i, "is", mobxUser.mobxRewards)
		us[i].mobxRewards = mobxUser.mobxRewards
	}
}

func (us userSet) resetNFT() {
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

// func getNFTRewards(us userSet) (string, float64) {
// 	for _, mobxUser := range us {
// 		fmt.Println(mobxUser.address, mobxUser.nftWeight)
// 		if mobxUser.nftWeight == 1.5 {
// 			fmt.Println("MOBXuser with NFT found")
// 			return mobxUser.address, mobxUser.mobxRewards
// 		} else {
// 			fmt.Println("This user doesn't have an NFT")
// 		}
// 	}

// 	return "0", 0
// }

// ToDo: function shouldn't stop after the first NFT find

func getNFTRewards2(us userSet, nft nftSet) []nfTier {

	for _, mobxUser := range us {

		for _, nftUser := range nft {
			// fmt.Println("entered nft loop")
			switch mobxUser.nftWeight {
			case 1:
				if nftUser.address == "" {
					// fmt.Println("For Tier 1 - entering the if worked")
					tier1 := nftSet{{1, mobxUser.address, mobxUser.mobxRewards, mobxUser.nftWeight}}
					nft[0] = tier1[0]
					// fmt.Println(nft[0])
				}
			case 1.5:
				if nftUser.address == "" {
					// fmt.Println("For Tier 1.5 - entering the if worked")
					tier15 := nftSet{{1.5, mobxUser.address, mobxUser.mobxRewards, mobxUser.nftWeight}}
					nft[1] = tier15[0]
					// fmt.Println(nft[1])
				}
			case 2:
				if nftUser.address == "" {
					// fmt.Println("For Tier 2 - entering the if worked")
					tier2 := nftSet{{2, mobxUser.address, mobxUser.mobxRewards, mobxUser.nftWeight}}
					nft[2] = tier2[0]
					// fmt.Println(nft[2])
				}
			case 3:
				if nftUser.address == "" {
					// fmt.Println("For Tier 3 - entering the if worked")
					tier3 := nftSet{{3, mobxUser.address, mobxUser.mobxRewards, mobxUser.nftWeight}}
					nft[3] = tier3[0]
					// fmt.Println(nft[3])
				}
			default:
				fmt.Println("default")
			}

		}
	}

	// fmt.Println("Exited getNFTrewards")
	// fmt.Println(nft)
	return nft
}

// func getNFTRewards3(us userSet, ns nftSet) []string {

// 	for i, mobxUser := range us {
// 		switch mobxUser.nftWeight {
// 		case 1:
// 			for _, nftSet := range ns {
// 				if nftSet.nftWeight != 1 {
// 					ns[i].address = mobxUser.address
// 					ns[i].nftWeight = mobxUser.nftWeight
// 				}
// 			}
// 		case 1.5:
// 			for _, nftSet := range ns {
// 				if nftSet.nftWeight != 1.5 {
// 					ns[i].address = mobxUser.address
// 					ns[i].nftWeight = mobxUser.nftWeight
// 				}
// 			}
// 		case 2:
// 			for _, nftSet := range ns {
// 				if nftSet.nftWeight != 1 {
// 					ns[i].address = mobxUser.address
// 					ns[i].nftWeight = mobxUser.nftWeight
// 				}
// 			}
// 		case 3:
// 			for _, nftSet := range ns {
// 				if nftSet.nftWeight != 1 {
// 					ns[i].address = mobxUser.address
// 					ns[i].nftWeight = mobxUser.nftWeight
// 				}
// 			}
// 		}
// 	}

// 	nftTier := []string{tier0, tier1, tier2, tier3}
// 	fmt.Println(nftTier)
// 	return nftTier
// }

func (us userSet) calculateNFTbonus(us2 userSet) {
	nftGroup := nftSet{{1, "", 0, 0}, {1.5, "", 0, 0}, {2, "", 0, 0}, {3, "", 0, 0}}
	nft := getNFTRewards2(us, nftGroup)

	for _, nfTier := range nft {
		for _, mobxUser := range us2 {
			if mobxUser.address == nfTier.address {
				fmt.Println("+++++++++++++++++++++++++++++++++")
				fmt.Println("Reward with Tier", nfTier.nftWeight, "is :", nfTier.mobxRewards)
				fmt.Println("Reward without NFT", mobxUser.mobxRewards)
				nftBonus := ((nfTier.mobxRewards / mobxUser.mobxRewards) - 1) * 100

				fmt.Println("The NFT bonus with Tier", nfTier.nftWeight, " is:", nftBonus, "%")
			}
		}
	}

}
