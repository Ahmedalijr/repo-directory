func (s *KeeperTestHelper) SetupValidator(bondStatus stakingtypes.BondStatus) sdk.ValAddress {
	valPub := secp256k1.GenPrivKey().PubKey()
	valAddr := sdk.ValAddress(valPub.Address())
	bondDenom := s.App.StakingKeeper.GetParams(s.Ctx).BondDenom


	s.FundAcc(sdk.AccAddress(valAddr), selfBond)

	

	val, found := s.App.StakingKeeper.GetValidator(s.Ctx, valAddr)
