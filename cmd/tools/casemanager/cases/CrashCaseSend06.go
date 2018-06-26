package cases

import (
	"fmt"

	"time"

	"github.com/SmartMeshFoundation/SmartRaiden/cmd/tools/casemanager/models"
	"github.com/SmartMeshFoundation/SmartRaiden/cmd/tools/casemanager/utils"
	"github.com/SmartMeshFoundation/SmartRaiden/params"
)

// CrashCaseSend06 场景六：EventSendRefundTransferAfter
// 发送refundtransfer交易崩溃
// 节点2发送45token给节点6 ，发送refundtransfer后节点3崩，节点2锁定45，其余节点无锁定;
// 重启节点3后，交易失败，通道2-3中节点2，3各锁定 45，通道2-7中节点2锁定45
func (cm *CaseManager) CrashCaseSend06() (err error) {
	env, err := models.NewTestEnv("./cases/CrashCaseSend06.ENV")
	if err != nil {
		return
	}
	defer env.KillAllRaidenNodes()
	// 源数据
	var transAmount int32
	var msg string
	transAmount = 45
	tokenAddress := env.Tokens[0].Address
	N2, N3, N6, N7 := env.Nodes[0], env.Nodes[1], env.Nodes[2], env.Nodes[3]
	models.Logger.Println(env.CaseName + " BEGIN ====>")
	// 启动节点2，6，7
	N2.Start(env)
	N6.Start(env)
	N7.Start(env)
	// 启动节点3, EventSendRefundTransferAfter
	N3.StartWithConditionQuit(env, &params.ConditionQuit{
		QuitEvent: "EventSendRefundTransferAfter",
	})

	// 节点2向节点6转账20token
	N2.SendTrans(tokenAddress, transAmount, N6.Address, false)
	time.Sleep(time.Second * 30)
	//  崩溃判断
	if N3.IsRunning() {
		msg = "Node " + N3.Name + " should be exited,but it still running, FAILED !!!"
		models.Logger.Println(msg)
		return fmt.Errorf(msg)
	}

	// 查询cd23，锁定45
	cd23middle := utils.GetChannelBetween(N2, N3, tokenAddress).PrintDataAfterCrash()
	if cd23middle.LockedAmount != transAmount {
		msg = fmt.Sprintf("Expect locked amount = %d,but got %d ,FAILED!!!", transAmount, cd23middle.LockedAmount)
		models.Logger.Println(msg)
		return fmt.Errorf(msg)
	}
	// 查询cd63,cd27,cd73均无锁定
	cd63middle := utils.GetChannelBetween(N6, N3, tokenAddress).PrintDataAfterCrash()
	if cd63middle.LockedAmount != 0 || cd63middle.PartnerLockedAmount != 0 {
		msg = fmt.Sprintf("Expect locked amount = %d,but got %d ,FAILED!!!", transAmount, cd63middle.LockedAmount)
		models.Logger.Println(msg)
		return fmt.Errorf(msg)
	}
	cd27middle := utils.GetChannelBetween(N2, N7, tokenAddress).PrintDataAfterCrash()
	if cd63middle.LockedAmount != 0 || cd27middle.PartnerLockedAmount != 0 {
		msg = fmt.Sprintf("Expect locked amount = %d,but got %d ,FAILED!!!", transAmount, cd27middle.LockedAmount)
		models.Logger.Println(msg)
		return fmt.Errorf(msg)
	}
	cd73middle := utils.GetChannelBetween(N7, N3, tokenAddress).PrintDataAfterCrash()
	if cd73middle.LockedAmount != 0 || cd73middle.PartnerLockedAmount != 0 {
		msg = fmt.Sprintf("Expect locked amount = %d,but got %d ,FAILED!!!", transAmount, cd73middle.LockedAmount)
		models.Logger.Println(msg)
		return fmt.Errorf(msg)
	}

	// 重启节点3，自动发送之前中断的交易
	N3.ReStartWithoutConditionquit(env)

	// 查询cd23并校验
	cd23new := utils.GetChannelBetween(N2, N3, tokenAddress).PrintDataAfterRestart()
	if cd23new.PartnerLockedAmount != transAmount || cd23new.LockedAmount != transAmount {
		models.Logger.Println(env.CaseName + " END ====> FAILED")
		return fmt.Errorf("Case [%s] FAILED", env.CaseName)
	}
	// 查询cd63并校验
	cd63new := utils.GetChannelBetween(N6, N3, tokenAddress).PrintDataAfterRestart()
	if cd63new.Balance-cd63middle.Balance != 0 {
		models.Logger.Println(env.CaseName + " END ====> FAILED")
		return fmt.Errorf("Case [%s] FAILED", env.CaseName)
	}
	// 查询cd73并校验
	cd73new := utils.GetChannelBetween(N7, N3, tokenAddress).PrintDataAfterRestart()
	if cd73new.Balance-cd73middle.Balance != 0 {
		models.Logger.Println(env.CaseName + " END ====> FAILED")
		return fmt.Errorf("Case [%s] FAILED", env.CaseName)
	}
	// 查询cd27并校验
	cd27new := utils.GetChannelBetween(N2, N7, tokenAddress).PrintDataAfterRestart()
	if cd27new.LockedAmount != transAmount {
		models.Logger.Println(env.CaseName + " END ====> FAILED")
		return fmt.Errorf("Case [%s] FAILED", env.CaseName)
	}
	models.Logger.Println(env.CaseName + " END ====> SUCCESS")
	return
}
