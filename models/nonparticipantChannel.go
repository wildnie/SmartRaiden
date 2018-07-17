package models

import (
	"fmt"

	"math/big"

	"bytes"

	"github.com/SmartMeshFoundation/SmartRaiden/log"
	"github.com/SmartMeshFoundation/SmartRaiden/utils"
	"github.com/asdine/storm"
	"github.com/ethereum/go-ethereum/common"
)

type NonParticipantChannel struct {
	Participant1        common.Address
	Participant2        common.Address
	Participant1Balance *big.Int
	Participant2Balance *big.Int
}

func participant2bytes(p1, p2 common.Address) []byte {
	b := make([]byte, len(p1)*2)
	copy(b[0:len(p1)], p1[:])
	copy(b[len(p1):], p2[:])
	return b
}
func bytes2participant(data []byte) (p1, p2 common.Address) {
	if len(data) != len(p1)*2 {
		return
	}
	copy(p1[:], data[:len(p1)])
	copy(p2[:], data[len(p1):])
	return
}
func participantKey(p1, p2 common.Address) common.Address {
	t := utils.Sha3(p1[:], p2[:])
	return common.BytesToAddress(t[:])
}

/*
如果这个 map 很大,怎么办?存储效率肯定会很低.
否则怎么遍历呢?
*/
type ChannelParticipantMap map[common.Address][]byte

const bucketChannel = "bucketChannel"

func (model *ModelDB) NewNonParticipantChannel(token, participant1, participant2 common.Address) error {
	var m ChannelParticipantMap
	log.Trace(fmt.Sprintf("NewNonParticipantChannel token=%s,participant1=%s,participant2=%s",
		utils.APex2(token),
		utils.APex2(participant1),
		utils.APex2(participant2),
	))
	err := model.db.Get(bucketChannel, token, &m)
	if err != nil {
		if err == storm.ErrNotFound {
			m = make(ChannelParticipantMap)
		} else {
			return err
		}

	}
	if participant1 == participant2 {
		panic(fmt.Sprintf("channel error, p1 andf p2 is the same,token=%s,participant=%s", token.String(), participant1.String()))
	}
	if bytes.Compare(participant1[:], participant2[:]) > 0 {
		participant1, participant2 = participant2, participant1
	}
	key := participantKey(participant1, participant2)
	if m[key] != nil {
		//startup ...
		log.Warn(fmt.Sprintf("add channel ,but channel already exists, maybe duplicates channelnew events,participant1=%s,participant2=%s",
			utils.APex2(participant1), utils.APex2(participant2)))
		return nil
	}
	m[key] = participant2bytes(participant1, participant2)
	err = model.db.Set(bucketChannel, token, m)
	return err
}
func (model *ModelDB) RemoveNonParticipantChannel(token, participant1, participant2 common.Address) error {
	var m ChannelParticipantMap
	err := model.db.Get(bucketChannel, token, &m)
	if err != nil {
		if err == storm.ErrNotFound {
			m = make(ChannelParticipantMap)
		} else {
			return err
		}

	}
	if participant1 == participant2 {
		panic(fmt.Sprintf("channel error, p1 andf p2 is the same,token=%s,participant=%s", token.String(), participant1.String()))
	}
	if bytes.Compare(participant1[:], participant2[:]) > 0 {
		participant1, participant2 = participant2, participant1
	}
	key := participantKey(participant1, participant2)
	if m[key] == nil {
		//startup ...
		return fmt.Errorf("delete channel ,but channel don't exists")
	}
	delete(m, key)
	err = model.db.Set(bucketChannel, token, m)
	return err
}

//GetAllTokens returna all tokens on this registry contract
func (model *ModelDB) GetAllNonParticipantChannel(token common.Address) (edges []common.Address, err error) {
	var m ChannelParticipantMap
	err = model.db.Get(bucketChannel, token, &m)
	if err == storm.ErrNotFound {
		err = nil
		return
	}
	for _, data := range m {
		p1, p2 := bytes2participant(data)
		edges = append(edges, p1, p2)
	}
	return
}