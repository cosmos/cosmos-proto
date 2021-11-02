package customtype

type AccAddress string

func (a AccAddress) Validate(sdkConfig interface{}) error {
	return nil
}
