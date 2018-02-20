package models

import (
	"os"
	"testing"

	"path"

	"reflect"

	"fmt"

	"bytes"
	"encoding/gob"

	"encoding/hex"

	"github.com/SmartMeshFoundation/raiden-network/channel"
	"github.com/SmartMeshFoundation/raiden-network/network"
	"github.com/SmartMeshFoundation/raiden-network/params"
	"github.com/SmartMeshFoundation/raiden-network/transfer"
	"github.com/SmartMeshFoundation/raiden-network/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func init() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlTrace, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))
}
func setupDb(t *testing.T) (model *ModelDB) {
	dbPath := path.Join(os.TempDir(), "testxxxx.db")
	//os.Remove(dbPath)
	os.Remove(dbPath + ".lock")
	model, err := OpenDb(dbPath)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(model.db)
	return
}

func TestNewStateChangeLog(t *testing.T) {
	model := setupDb(t)
	defer func() {
		model.CloseDB()
	}()
	st := &transfer.BlockStateChange{3}
	id, err := model.LogStateChange(st)
	if err != nil {
		t.Error(err)
	}
	//if id != 1 {
	//	t.Error("id not equal 1, ", id)
	//}
	st2, err := model.GetStateChangeById(id)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(st, st2) {
		t.Error("data not equal")
	}
	s := &snapshotToWrite{
		StateChangeId: 1,
		State:         st,
	}
	_, err = model.Snapshot(s.StateChangeId, s)
	if err != nil {
		t.Error(err)
	}
	s2, err := model.LoadSnapshot()
	if err != nil {
		t.Error(err)
	}
	if reflect.DeepEqual(st, s2) {
		t.Error("data3 not equal")
	}
	number := utils.RandSrc.Int63()
	err = model.LogEvents(id, []transfer.Event{st}, number)
	if err != nil {
		t.Error(err)
	}
	events2, err := model.GetEventsInBlockRange(number, number+1)
	if err != nil {
		t.Error(err)
	}
	if len(events2) != 1 {
		t.Error("events length error")
	}
	t.Log("events2=%#v", events2[0])
}

func TestToken(t *testing.T) {
	model := setupDb(t)
	defer func() {
		model.CloseDB()
	}()
	var cbtokens []common.Address
	funcb := func(token common.Address) bool {
		cbtokens = append(cbtokens, token)
		return true
	}
	ts, err := model.GetAllTokens()
	if err == nil {
		t.Error("should not found")
	}
	if len(ts) != 0 {
		t.Error("should be empty")
	}
	tokens := make(map[common.Address]*network.ChannelGraph)
	for i := 0; i < 3; i++ {
		tokens[utils.NewRandomAddress()] = nil
	}
	err = model.SetAllTokens(tokens)
	if err != nil {
		t.Error(err)
	}
	t1 := utils.NewRandomAddress()
	model.RegisterNewTokenCallback(funcb)
	err = model.AddToken(t1)
	if err != nil {
		t.Error(err)
	}
	t2 := utils.NewRandomAddress()
	err = model.AddToken(t2)
	if err != nil {
		t.Error(err)
	}
	if len(cbtokens) != 1 && cbtokens[0] != t1 {
		t.Error("add token error")
	}
}

func TestChannel(t *testing.T) {
	model := setupDb(t)
	defer func() {
		model.CloseDB()
	}()
	var newaddrs []common.Address
	var updateContractBalanceAddrs []common.Address
	var UpdateChannelStateAddrs []common.Address
	newchannelcb := func(c *channel.ChannelSerialization) bool {
		newaddrs = append(newaddrs, c.ChannelAddress)
		return true
	}
	updateContractBalancechannelcb := func(c *channel.ChannelSerialization) bool {
		updateContractBalanceAddrs = append(updateContractBalanceAddrs, c.ChannelAddress)
		return true
	}
	UpdateChannelStatecb := func(c *channel.ChannelSerialization) bool {
		UpdateChannelStateAddrs = append(UpdateChannelStateAddrs, c.ChannelAddress)
		return true
	}
	model.RegisterNewChannellCallback(newchannelcb)
	model.RegisterChannelDepositCallback(updateContractBalancechannelcb)
	model.RegisterChannelStateCallback(UpdateChannelStatecb)

	ch, _ := channel.MakeTestPairChannel()
	c := channel.NewChannelSerialization(ch)
	err := model.NewChannel(c)
	if err != nil {
		t.Error(err)
		return
	}
	chs, err := model.GetChannelList(utils.EmptyAddress, utils.EmptyAddress)
	if err != nil || len(chs) != 1 {
		t.Error(err)
		t.Log(fmt.Sprintf("chs=%v", utils.StringInterface(chs, 5)))
		return
	}
	err = model.UpdateChannelNoTx(c)
	if err != nil {
		t.Error(err)
		return
	}
	chs, err = model.GetChannelList(utils.EmptyAddress, utils.EmptyAddress)
	if err != nil || len(chs) != 1 {
		t.Error(err)
		return
	}
	err = model.UpdateChannelContractBalance(c)
	if err != nil {
		t.Error(err)
		return
	}
	err = model.UpdateChannelContractBalance(c)
	if err != nil {
		t.Error(err)
		return
	}
	err = model.UpdateChannelState(c)
	if err != nil {
		t.Error(err)
		return
	}
	err = model.UpdateChannelState(c)
	if err != nil {
		t.Error(err)
		return
	}
	//if len(newaddrs) != 1 || newaddrs[0] != c.ChannelAddress {
	//	t.Error("new channel error")
	//}
	//if len(updateContractBalanceAddrs) != 1 || updateContractBalanceAddrs[0] != c.ChannelAddress {
	//	t.Error("new channel error")
	//}
	//if len(UpdateChannelStateAddrs) != 1 || UpdateChannelStateAddrs[0] != c.ChannelAddress {
	//	t.Error("new channel error")
	//}
}
func TestChannelTwice(t *testing.T) {
	TestChannel(t)
	TestChannel(t)
}
func TestGob(t *testing.T) {
	s1 := params.ROPSTEN_REGISTRY_ADDRESS
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(s1)
	if err != nil {
		t.Error(err)
		return
	}
	encodedData := buf.Bytes()
	fmt.Printf("first\n%s", hex.Dump(encodedData))
	dec := gob.NewDecoder(bytes.NewBuffer(encodedData))
	var sb common.Address
	err = dec.Decode(&sb)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(s1, sb) {
		t.Error("not equal")
	}
	var buf2 bytes.Buffer
	enc2 := gob.NewEncoder(&buf2)
	enc2.Encode(&sb)
	encodedData2 := buf2.Bytes()
	fmt.Printf("second\n%s", hex.Dump(encodedData2))
	if !reflect.DeepEqual(encodedData, encodedData2) {
		t.Error("not equal")
	}

}