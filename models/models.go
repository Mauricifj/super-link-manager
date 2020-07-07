package models

type Response struct {
	Link   *Link
	Errors *ErrorResponse
}

type Link struct {
	Id                      string   `json:"id,omitempty"`
	LinkType                string   `json:"type"`
	Name                    string   `json:"name"`
	Description             string   `json:"description"`
	Price                   int      `json:"price"`
	Weight                  int      `json:"weight"`
	ExpirationDate          string   `json:"expirationDate"`
	MaxNumberOfInstallments int      `json:"maxNumberOfInstallments"`
	Quantity                int      `json:"quantity"`
	Sku                     string   `json:"sku"`
	Shipping                Shipping `json:"shipping"`
	SoftDescriptor          string   `json:"softDescriptor"`
}

type Shipping struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type ErrorResponse []ErrorResponseItem

type ErrorResponseItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
}