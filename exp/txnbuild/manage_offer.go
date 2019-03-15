package txnbuild

import (
	"github.com/stellar/go/amount"
	"github.com/stellar/go/price"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/xdr"
)

// ManageOffer represents the Stellar manage offer operation. See
// https://www.stellar.org/developers/guides/concepts/list-of-operations.html
type ManageOffer struct {
	Selling *Asset
	Buying  *Asset
	Amount  string
	Price   string // TODO: Extend to include number, and n/d fraction. See package 'amount'
	OfferID uint64
	xdrOp   xdr.ManageOfferOp
}

// BuildXDR for ManageOffer returns a fully configured XDR Operation.
func (mo *ManageOffer) BuildXDR() (xdr.Operation, error) {
	xdrSelling, err := mo.Selling.ToXDR()
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "Failed to set XDR 'Selling' field")
	}
	mo.xdrOp.Selling = xdrSelling

	xdrBuying, err := mo.Buying.ToXDR()
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "Failed to set XDR 'Buying' field")
	}
	mo.xdrOp.Buying = xdrBuying

	xdrAmount, err := amount.Parse(mo.Amount)
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "Failed to parse 'Amount'")
	}
	mo.xdrOp.Amount = xdrAmount

	xdrPrice, err := price.Parse(mo.Price)
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "Failed to parse 'Price'")
	}
	mo.xdrOp.Price = xdrPrice

	mo.xdrOp.OfferId = xdr.Uint64(mo.OfferID)

	opType := xdr.OperationTypeManageOffer
	body, err := xdr.NewOperationBody(opType, mo.xdrOp)
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "Failed to build XDR OperationBody")
	}

	return xdr.Operation{Body: body}, nil
}
