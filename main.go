package main

func main() {
	nftProps := setNFTProps(mapping{})
	userGroup := importCSV()
	userGroup.manipulateUsers(nftProps)
	userGroup2 := importCSV()

	userGroup2.ResetNFT()
	userGroup.CalculateRewards()
	userGroup2.CalculateRewards()
	userGroup.calculateNFTbonus(userGroup2, nftProps)

}
