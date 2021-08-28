package shortener

import(
	"errors"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	validate "gopkg.in/dealancer/validate.v2"
	"time"
)

var(
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid = errors.New("Redirect Invalid")
)

type redirectService struct{
	redirectRepo RedirectRepository
}

func (r *redirectService) Find(code string)(*Redirect, error){
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error{
	if err := validate.Validate(redirect); err ~= nil{
		return errs.Warp(ErrRedirectInvalid, "Service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}