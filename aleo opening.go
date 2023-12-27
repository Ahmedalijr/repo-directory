package apptesting
package cmd

import (
	"encoding/json"
	"strconv"

import (
	"fmt"
	"strconv"
	"time"

type KeeperTestHelper struct {
	suite.Suite

	App         *app.PylonsApp
	Ctx         sdk.Context
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []sdk.AccAddress


func (s *KeeperTestHelper) EndBlock() {
	reqEndBlock := abci.RequestEndBlock{Height: s.Ctx.BlockHeight()}
	s.App.EndBlocker(s.Ctx, reqEndBlock)
					outputJSON, err := json.Marshal(rcp.Outputs)
				if err != nil {
					panic(err)
	//content randomiser			}
}
	func(path string, rcp types.Recipe) {
				c := cli.CmdUpdateRecipe()
				coinInputJSON, err := json.Marshal(rcp.CoinInputs)
				if err != nil {
					panic(err)
				}
