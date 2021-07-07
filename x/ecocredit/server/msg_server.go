package server

import (
	"fmt"

	"github.com/regen-network/regen-ledger/types"

	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/util"
)

// CreateClass creates a new class of ecocredit
//
// The designer is charged a fee for creating the class. This is controlled by
// the global parameter CreditClassFee, which can be updated through the
// governance process.
func (s serverImpl) CreateClass(ctx types.Context, req *ecocredit.MsgCreateClassRequest) (*ecocredit.MsgCreateClassResponse, error) {
	classID := s.idSeq.NextVal(ctx)
	classIDStr := util.Uint64ToBase58Check(classID)

	// Charge the designer a fee to create the credit class
	designerAddress, err := sdk.AccAddressFromBech32(req.Designer)
	if err != nil {
		return nil, err
	}

	err = s.chargeCreditClassFee(ctx.Context, designerAddress)
	if err != nil {
		return nil, err
	}

	err = s.classInfoTable.Create(ctx, &ecocredit.ClassInfo{
		ClassId:  classIDStr,
		Designer: req.Designer,
		Issuers:  req.Issuers,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCreateClass{
		ClassId:  classIDStr,
		Designer: req.Designer,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateClassResponse{ClassId: classIDStr}, nil
}

func (s serverImpl) CreateBatch(ctx types.Context, req *ecocredit.MsgCreateBatchRequest) (*ecocredit.MsgCreateBatchResponse, error) {
	classID := req.ClassId
	if err := s.assertClassIssuer(ctx, classID, req.Issuer); err != nil {
		return nil, err
	}

	batchID := s.idSeq.NextVal(ctx)
	batchDenom := batchDenomT(fmt.Sprintf("%s/%s", classID, util.Uint64ToBase58Check(batchID)))
	tradableSupply := apd.New(0, 0)
	retiredSupply := apd.New(0, 0)
	var maxDecimalPlaces uint32 = 0

	store := ctx.KVStore(s.storeKey)

	for _, issuance := range req.Issuance {
		tradable, err := math.ParseNonNegativeDecimal(issuance.TradableUnits)
		if err != nil {
			return nil, err
		}

		decPlaces := math.NumDecimalPlaces(tradable)
		if decPlaces > maxDecimalPlaces {
			maxDecimalPlaces = decPlaces
		}

		retired, err := math.ParseNonNegativeDecimal(issuance.RetiredUnits)
		if err != nil {
			return nil, err
		}

		decPlaces = math.NumDecimalPlaces(retired)
		if decPlaces > maxDecimalPlaces {
			maxDecimalPlaces = decPlaces
		}

		recipient := issuance.Recipient
		recipientAddr, err := sdk.AccAddressFromBech32(recipient)
		if err != nil {
			return nil, err
		}

		if !tradable.IsZero() {
			err = math.Add(tradableSupply, tradableSupply, tradable)
			if err != nil {
				return nil, err
			}

			err = getAddAndSetDecimal(store, TradableBalanceKey(recipientAddr, batchDenom), tradable)
			if err != nil {
				return nil, err
			}
		}

		if !retired.IsZero() {
			err = math.Add(retiredSupply, retiredSupply, retired)
			if err != nil {
				return nil, err
			}

			err = retire(ctx, store, recipientAddr, batchDenom, retired, issuance.RetirementLocation)
			if err != nil {
				return nil, err
			}
		}

		var sum apd.Decimal
		err = math.Add(&sum, tradable, retired)
		if err != nil {
			return nil, err
		}

		err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventReceive{
			Recipient:  recipient,
			BatchDenom: string(batchDenom),
			Units:      math.DecimalString(&sum),
		})
		if err != nil {
			return nil, err
		}
	}

	setDecimal(store, TradableSupplyKey(batchDenom), tradableSupply)
	setDecimal(store, RetiredSupplyKey(batchDenom), retiredSupply)

	var totalSupply apd.Decimal
	err := math.Add(&totalSupply, tradableSupply, retiredSupply)
	if err != nil {
		return nil, err
	}

	totalSupplyStr := math.DecimalString(&totalSupply)
	err = s.batchInfoTable.Create(ctx, &ecocredit.BatchInfo{
		ClassId:    classID,
		BatchDenom: string(batchDenom),
		Issuer:     req.Issuer,
		TotalUnits: totalSupplyStr,
		Metadata:   req.Metadata,
	})
	if err != nil {
		return nil, err
	}

	err = setUInt32(store, MaxDecimalPlacesKey(batchDenom), maxDecimalPlaces)
	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCreateBatch{
		ClassId:    classID,
		BatchDenom: string(batchDenom),
		Issuer:     req.Issuer,
		TotalUnits: totalSupplyStr,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateBatchResponse{BatchDenom: string(batchDenom)}, nil
}

func (s serverImpl) Send(ctx types.Context, req *ecocredit.MsgSendRequest) (*ecocredit.MsgSendResponse, error) {
	store := ctx.KVStore(s.storeKey)
	sender := req.Sender
	recipient := req.Recipient

	senderAddr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return nil, err
	}

	recipientAddr, err := sdk.AccAddressFromBech32(recipient)
	if err != nil {
		return nil, err
	}

	for _, credit := range req.Credits {
		denom := batchDenomT(credit.BatchDenom)

		maxDecimalPlaces, err := getUint32(store, MaxDecimalPlacesKey(denom))
		if err != nil {
			return nil, err
		}

		tradable, err := math.ParseNonNegativeFixedDecimal(credit.TradableUnits, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		retired, err := math.ParseNonNegativeFixedDecimal(credit.RetiredUnits, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		var sum apd.Decimal
		err = math.Add(&sum, tradable, retired)
		if err != nil {
			return nil, err
		}

		// subtract balance
		err = getSubAndSetDecimal(store, TradableBalanceKey(senderAddr, denom), &sum)
		if err != nil {
			return nil, err
		}

		// Add tradable balance
		err = getAddAndSetDecimal(store, TradableBalanceKey(recipientAddr, denom), tradable)
		if err != nil {
			return nil, err
		}

		if !retired.IsZero() {
			// subtract retired from tradable supply
			err = getSubAndSetDecimal(store, TradableSupplyKey(denom), retired)
			if err != nil {
				return nil, err
			}

			// Add retired balance
			err = retire(ctx, store, recipientAddr, denom, retired, credit.RetirementLocation)
			if err != nil {
				return nil, err
			}

			// Add retired supply
			err = getAddAndSetDecimal(store, RetiredSupplyKey(denom), retired)
			if err != nil {
				return nil, err
			}
		}

		err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventReceive{
			Sender:     sender,
			Recipient:  recipient,
			BatchDenom: string(denom),
			Units:      math.DecimalString(&sum),
		})
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgSendResponse{}, nil
}

func (s serverImpl) Retire(ctx types.Context, req *ecocredit.MsgRetireRequest) (*ecocredit.MsgRetireResponse, error) {
	store := ctx.KVStore(s.storeKey)
	holder := req.Holder

	for _, credit := range req.Credits {
		denom := batchDenomT(credit.BatchDenom)
		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not a valid credit denom", denom))
		}

		maxDecimalPlaces, err := getUint32(store, MaxDecimalPlacesKey(denom))
		if err != nil {
			return nil, err
		}

		toRetire, err := math.ParsePositiveFixedDecimal(credit.Units, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		holderAddr, err := sdk.AccAddressFromBech32(holder)
		if err != nil {
			return nil, err
		}

		// subtract tradable balance
		err = getSubAndSetDecimal(store, TradableBalanceKey(holderAddr, denom), toRetire)
		if err != nil {
			return nil, err
		}

		// subtract tradable supply
		err = getSubAndSetDecimal(store, TradableSupplyKey(denom), toRetire)
		if err != nil {
			return nil, err
		}

		//  Add retired balance
		err = retire(ctx, store, holderAddr, denom, toRetire, req.Location)
		if err != nil {
			return nil, err
		}

		//  Add retired supply
		err = getAddAndSetDecimal(store, RetiredSupplyKey(denom), toRetire)
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgRetireResponse{}, nil
}

func (s serverImpl) SetPrecision(ctx types.Context, req *ecocredit.MsgSetPrecisionRequest) (*ecocredit.MsgSetPrecisionResponse, error) {
	var batchInfo ecocredit.BatchInfo
	err := s.batchInfoTable.GetOne(ctx, orm.RowID(req.BatchDenom), &batchInfo)
	if err != nil {
		return nil, err
	}
	if req.Issuer != batchInfo.Issuer {
		return nil, sdkerrors.ErrUnauthorized
	}
	store := ctx.KVStore(s.storeKey)
	key := MaxDecimalPlacesKey(batchDenomT(req.BatchDenom))
	x, err := getUint32(store, key)
	if err != nil {
		return nil, err
	}

	if req.MaxDecimalPlaces <= x {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Maximum decimal can only be increased, it is currently %d, and %d was requested", x, req.MaxDecimalPlaces))
	}

	err = setUInt32(store, key, req.MaxDecimalPlaces)
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgSetPrecisionResponse{}, nil
}

// assertClassIssuer makes sure that the issuer is part of issuers of given classID.
// Returns ErrUnauthorized otherwise.
func (s serverImpl) assertClassIssuer(ctx types.Context, classID, issuer string) error {
	classInfo, err := s.getClassInfo(ctx, classID)
	if err != nil {
		return err
	}
	for _, i := range classInfo.Issuers {
		if issuer == i {
			return nil
		}
	}
	return sdkerrors.ErrUnauthorized
}

func retire(ctx types.Context, store sdk.KVStore, recipient sdk.AccAddress, batchDenom batchDenomT, retired *apd.Decimal, location string) error {
	err := getAddAndSetDecimal(store, RetiredBalanceKey(recipient, batchDenom), retired)
	if err != nil {
		return err
	}

	return ctx.EventManager().EmitTypedEvent(&ecocredit.EventRetire{
		Retirer:    recipient.String(),
		BatchDenom: string(batchDenom),
		Units:      math.DecimalString(retired),
		Location:   location,
	})
}
