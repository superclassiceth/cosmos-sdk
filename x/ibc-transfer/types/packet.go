package types

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// DefaultRelativePacketTimeoutHeight is the default packet timeout height (in blocks) relative
	// to the current block height of the counterparty chain provided by the client state. The
	// timeout is disabled when set to 0.
	DefaultRelativePacketTimeoutHeight = uint64(1000)

	// DefaultRelativePacketTimeoutTimestamp is the default packet timeout timestamp (in nanoseconds)
	// relative to the current block timestamp of the counterparty chain provided by the client
	// state. The timeout is disabled when set to 0. The default is currently set to a 10 minute
	// timeout.
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

// NewFungibleTokenPacketData contructs a new FungibleTokenPacketData instance
func NewFungibleTokenPacketData(
	denom string, amount uint64,
	sender, receiver string,
) FungibleTokenPacketData {
	return FungibleTokenPacketData{
		Denom:    denom,
		Amount:   amount,
		Sender:   sender,
		Receiver: receiver,
	}
}

// ValidateBasic is used for validating the token transfer
func (ftpd FungibleTokenPacketData) ValidateBasic() error {
	if strings.TrimSpace(ftpd.Denom) == "" {
		return sdkerrors.Wrap(ErrInvalidDenomForTransfer, "denom cannot be empty")
	}
	if ftpd.Amount == 0 {
		return sdkerrors.Wrap(ErrInvalidAmount, "amount cannot be 0")
	}
	if strings.TrimSpace(ftpd.Sender) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be blank")
	}
	if strings.TrimSpace(ftpd.Receiver) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "receiver address cannot be blank")
	}
	return nil
}

// GetBytes is a helper for serialising
func (ftpd FungibleTokenPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(ftpd))
}

// GetBytes is a helper for serialising
func (ack FungibleTokenPacketAcknowledgement) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(ack))
}
