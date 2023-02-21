package model

type PaymentInvolvment struct {
	UserId             int
	ContributionOfUser float32
	PayByUser          float32
}
type Payment struct {
	PaymentId          int
	PaymentDescription string
	PaymentAmount      float32
	PayByUsers         []PaymentInvolvment
}
