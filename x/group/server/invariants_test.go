package server

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
)

func TestTallyVotesInvariant(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	key := sdk.NewKVStoreKey(group.ModuleName)
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}
	curCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	curCtx = curCtx.WithBlockHeight(10)
	prevCtx := curCtx.WithBlockHeight(curCtx.BlockHeight() - 1)

	// Proposal Table
	proposalTableBuilder := orm.NewAutoUInt64TableBuilder(ProposalTablePrefix, ProposalTableSeqPrefix, key, &group.Proposal{}, cdc)
	proposalTable := proposalTableBuilder.Build()

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	curBlockTime, err := gogotypes.TimestampProto(curCtx.BlockTime())
	if err != nil {
		fmt.Println("block time conversion")
		panic(err)
	}
	prevBlockTime, err := gogotypes.TimestampProto(prevCtx.BlockTime())
	if err != nil {
		fmt.Println("block time conversion")
		panic(err)
	}

	specs := map[string]struct {
		prevReq []*group.Proposal
		curReq  []*group.Proposal
		expErr  bool
	}{
		"invariant not broken": {
			prevReq: []*group.Proposal{
				{
					ProposalId:          0,
					Address:             addr1.String(),
					Proposers:           []string{addr1.String()},
					SubmittedAt:         *prevBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "1", NoCount: "0", AbstainCount: "0", VetoCount: "0"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},

			curReq: []*group.Proposal{
				{
					ProposalId:          0,
					Address:             addr2.String(),
					Proposers:           []string{addr2.String()},
					SubmittedAt:         *curBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "2", NoCount: "0", AbstainCount: "0", VetoCount: "0"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},
		},
		"current block yes vote count must be greater than previous block yes vote count": {
			prevReq: []*group.Proposal{
				{
					ProposalId:          0,
					Address:             addr1.String(),
					Proposers:           []string{addr1.String()},
					SubmittedAt:         *prevBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "2", NoCount: "0", AbstainCount: "0", VetoCount: "0"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},
			curReq: []*group.Proposal{
				{
					ProposalId:          0,
					Address:             addr2.String(),
					Proposers:           []string{addr2.String()},
					SubmittedAt:         *curBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "1", NoCount: "0", AbstainCount: "0", VetoCount: "0"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},
			expErr: true,
		},
		"current block no vote count must be greater than previous block no vote count": {
			prevReq: []*group.Proposal{
				{
					ProposalId:          0,
					Address:             addr1.String(),
					Proposers:           []string{addr1.String()},
					SubmittedAt:         *prevBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "0", NoCount: "2", AbstainCount: "0", VetoCount: "0"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},
			curReq: []*group.Proposal{
				{
					ProposalId:          0,
					Address:             addr2.String(),
					Proposers:           []string{addr2.String()},
					SubmittedAt:         *curBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "0", NoCount: "1", AbstainCount: "0", VetoCount: "0"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},
			expErr: true,
		},
		"current block abstain vote count must be greater than previous block abstain vote count": {
			prevReq: []*group.Proposal{
				{
					ProposalId:          0,
					Address:             addr1.String(),
					Proposers:           []string{addr1.String()},
					SubmittedAt:         *prevBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "0", NoCount: "0", AbstainCount: "2", VetoCount: "0"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},
			curReq: []*group.Proposal{
				{
					ProposalId:          0,
					Address:             addr2.String(),
					Proposers:           []string{addr2.String()},
					SubmittedAt:         *curBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "0", NoCount: "0", AbstainCount: "1", VetoCount: "0"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},
			expErr: true,
		},
		"current block veto vote count must be greater than previous block veto vote count": {
			prevReq: []*group.Proposal{
				{
					ProposalId:          0,
					Address:             addr1.String(),
					Proposers:           []string{addr1.String()},
					SubmittedAt:         *prevBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "0", NoCount: "0", AbstainCount: "0", VetoCount: "2"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},
			curReq: []*group.Proposal{
				{
					ProposalId:          0,
					Address:             addr2.String(),
					Proposers:           []string{addr2.String()},
					SubmittedAt:         *curBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "0", NoCount: "0", AbstainCount: "0", VetoCount: "1"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},
			expErr: true,
		},
	}

	for _, spec := range specs {

		prevProposals := spec.prevReq
		curProposals := spec.curReq

		for i := 0; i < len(prevProposals) && i < len(curProposals); i++ {
			_, err = proposalTable.Create(prevCtx, prevProposals[i])
			if err != nil {
				fmt.Println(err)
				panic("create proposal")
			}
			_, err = proposalTable.Create(curCtx, curProposals[i])
			if err != nil {
				fmt.Println(err)
				panic("create proposal")
			}
		}

		var test require.TestingT
		_, broken := tallyVotesInvariant(curCtx, proposalTable)
		require.Equal(test, spec.expErr, broken)
	}
}