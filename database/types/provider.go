package types

import (
	"database/sql/driver"
	"fmt"

	providertypes "github.com/ovrclk/akash/x/provider/types/v1beta2"
)

// ProviderRow represents a single row inside the provider table
type ProviderRow struct {
	OwnerAddress string         `db:"owner_address"`
	HostURI      string         `db:"host_uri"`
	Attributes   string         `db:"attributes"`
	Info         DbProviderInfo `db:"info"`
	Height       int64          `db:"height"`
}

// NewAccountBalanceRow allows to build a new AccountBalanceRow instance
func NewProviderRow(ownerAddress string, hostURI string, attributes string, info DbProviderInfo, height int64) ProviderRow {
	return ProviderRow{
		OwnerAddress: ownerAddress,
		HostURI:      hostURI,
		Attributes:   attributes,
		Info:         info,
		Height:       height,
	}
}

// Equal tells whether a and b contain the same data
func (a ProviderRow) Equal(b ProviderRow) bool {
	return a.OwnerAddress == b.OwnerAddress &&
		a.HostURI == b.HostURI &&
		// TO-DO: check those stored as JSON, how to equal??
		a.Attributes == b.Attributes &&
		a.Info.Equal(b.Info) &&
		a.Height == b.Height
}

// DbProviderInfo represents the information stored inside the database about a single provider info
type DbProviderInfo struct {
	EMail   string
	Website string
}

// NewDbInfo builds a DbInfo starting from an akash provider info
func NewDbInfo(info providertypes.ProviderInfo) DbProviderInfo {
	return DbProviderInfo{
		EMail:   info.EMail,
		Website: info.Website,
	}
}

// Equal tells whether coin and d represent the same coin with the same amount
func (a DbProviderInfo) Equal(b DbProviderInfo) bool {
	return a.EMail == b.EMail && a.Website == b.Website
}

// Value implements driver.Valuer
func (info *DbProviderInfo) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s)", info.EMail, info.Website), nil
}