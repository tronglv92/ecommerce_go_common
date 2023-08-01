package imgprocessing

import "github.com/tronglv92/ecommerce_go_common/sdkcm"

type Response struct {
	sdkcm.AppError
	Data *sdkcm.Image `json:"data"`
}
