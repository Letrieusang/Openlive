package common

type ResponsePaginate struct {
	HasError    int         `json:"hasError"`
	Message     string      `json:"message"`
	Limit       int         `json:"limit"`
	TotalRecord int64       `json:"totalRecord"`
	Page        int         `json:"page"`
	Data        interface{} `json:"data"`
}

type Response struct {
	HasError int         `json:"hasError"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

type ResponseShow struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	ExtraImage  string `json:"extra_image"`
	/*ExternalLink         string `json:"external_link"`
	SellerFeeBasisPoints string `json:"seller_fee_basis_points"`
	FeeRecipient         string `json:"fee_recipient"`*/
}

func NewResponseShow(id, name, description, image string) *ResponseShow {
	return &ResponseShow{
		Id:          id,
		Name:        name,
		Description: description,
		Image:       image,
		ExtraImage:  image,
	}
}

func NewResponse(hasError int, message string, data interface{}) *Response {
	return &Response{
		HasError: hasError,
		Message:  message,
		Data:     data,
	}
}

func NewResponseWithPaging(hasError int, message string, data interface{}, page Paging) *ResponsePaginate {
	return &ResponsePaginate{
		HasError:    hasError,
		Message:     message,
		Data:        data,
		Limit:       page.Limit,
		TotalRecord: page.Total,
		Page:        page.Page,
	}
}

func ErrInternal(err error) *Response {
	return &Response{
		HasError: 1,
		Message:  "something went wrong in the server",
		Data:     err.Error(),
	}
}
