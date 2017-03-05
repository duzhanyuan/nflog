package nflog

// See linux/netfilter/nfnetlink.h.

const NFNETLINK_V0 = 0

const (
	NFNL_SUBSYS_NONE = iota
	NFNL_SUBSYS_CTNETLINK
	NFNL_SUBSYS_CTNETLINK_EXP
	NFNL_SUBSYS_QUEUE
	NFNL_SUBSYS_ULOG
	NFNL_SUBSYS_OSF
	NFNL_SUBSYS_IPSET
	NFNL_SUBSYS_ACCT
	NFNL_SUBSYS_CTNETLINK_TIMEOUT
	NFNL_SUBSYS_CTHELPER
	NFNL_SUBSYS_NFTABLES
	NFNL_SUBSYS_NFT_COMPAT
	NFNL_SUBSYS_COUNT
)
