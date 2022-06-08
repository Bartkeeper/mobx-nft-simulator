package main

import (
	"fmt"
	"math/rand"
)

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

func getNFTRewards2(us UserSet, nft nftSet) []nfTier {

	// Requires a set of NFT tiers and then only picks one account of each tier
	// returns a set of accounts with distinct NFT tiers

	for _, mobxUser := range us {
		if mobxUser.collectedMiles != 0 {

			for i, nftUser := range nft {
				if nftUser.address == "" && nftUser.nftWeight == mobxUser.nftWeight {

					tier := nftSet{{nftUser.tierClass, mobxUser.address, mobxUser.mobxRewards, mobxUser.nftWeight, mobxUser.nftBonus}}

					nft[i] = tier[0]
					break
				}

			}
		}
	}

	return nft
}

func (us UserSet) calculateNFTbonus(us2 UserSet, m mapping) {

	// (currently) calculates the difference of the rewards compared to a world without NFT bonus
	// prints out final results

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

	nft := getNFTRewards2(us, nftGroup)
	fmt.Println("++++++++++++++++++++++++++++++++ Results ++++++++++++++++++++++++++++++++")
	for i, nfTier := range nft {
		for _, mobxUser := range us2 {
			if mobxUser.address == nfTier.address {

				nft[i].nftBonus = (((nfTier.mobxRewards / mobxUser.mobxRewards) - 1) * 100)
				effectiveBonus := int(nft[i].nftBonus - nft[0].nftBonus)

				fmt.Println("The absolute change with NFT Tier", nfTier.nftWeight, " is:", nft[i].nftBonus, "%")
				fmt.Println("The relative change with NFT Tier", nfTier.nftWeight, "compared to a non-NFT holder is", effectiveBonus, "%")
			}
		}
	}

}

func setNFTProps(m mapping) mapping {
	fmt.Println("How many different tiers do you want? (excluding non-NFT holders)")

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
	return m
}

func (us UserSet) manipulateUsers(m mapping) {

	for _, nftTier := range m {

		ln := len(us)
		am := int(nftTier.amount)

		for l := 0; l <= am; l++ {

			for o := 0; o < 100; o++ {
				user := int(rand.Intn(ln))
				if us[user].stakedMobx >= 10 {
					us[user].nftWeight = nftTier.multiplier
					break
				}
			}
		}

	}

}
