package model

type ProdukTerlaris struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

type SalesReport struct {
	TotalRevenue    int            `json:"total_revenue"`
	TotalTransaksi  int            `json:"total_transaksi"`
	ProdukTerlaris  ProdukTerlaris `json:"produk_terlaris"`
}
