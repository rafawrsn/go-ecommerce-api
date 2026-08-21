package dto

type CreateProductRequest struct {
	Name       string `json:"name"`
	Price      int64    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID uint   `json:"category_id"`
}

type UpdateProductRequest struct {
    Name       string  `json:"name"`
    Price      int64 `json:"price"`
    Stock      int     `json:"stock"`
    CategoryID uint    `json:"category_id"`
}

type ProductResponse struct {
    ID         uint    `json:"id"`
    Name       string  `json:"name"`
    Price      int64 `json:"price"`
    Stock      int     `json:"stock"`
    CategoryID uint    `json:"category_id"`
	Category   CategoryResponse `json:"category"`
}

