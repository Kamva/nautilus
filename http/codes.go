package http

type StatusCode int

const (
	Continue StatusCode = iota + 100
	SwitchingProtocols
	Processing
)

const (
	OK StatusCode = iota + 200
	Created
	Accepted
	NonAuthoritativeInformation
	NoContent
	ResetContent
	PartialContent
	MultiStatus
	AlreadyReported
	IMUsed StatusCode = 226
)

const (
	MultipleChoices StatusCode = iota + 300
	MovedPermanently
	Found
	SeeOther
	NotModified
	UseProxy
	TemporaryRedirect
	PermanentRedirect
)

const (
	BadRequest StatusCode = iota + 400
	Unauthorized
	PaymentRequired
	Forbidden
	NotFound
	MethodNotAllowed
	NotAcceptable
	ProxyAuthenticationRequired
	RequestTimeout
	Conflict
	Gone
	LengthRequired
	PreconditionFailed
	PayloadTooLarge
	RequestURITooLong
	UnsupportedMediaType
	RequestedRangeNotSatisfiable
	ExpectationFailed
	ImATeapot
	_
	_
	MisdirectedRequest
	UnprocessableEntity
	Locked
	FailedDependency
	_
	UpgradeRequired
	_
	PreconditionRequired
	TooManyRequests
	RequestHeaderFieldsTooLarge
	_
	ConnectionClosedWithoutResponse StatusCode = 444
	UnavailableForLegalReasons      StatusCode = 451
	ClientClosedRequest             StatusCode = 499
)

const (
	InternalServerError StatusCode = iota + 500
	NotImplemented
	BadGateway
	ServiceUnavailable
	GatewayTimeout
	HTTPVersionNotSupported
	VariantAlsoNegotiates
	InsufficientStorage
	LoopDetected
	_
	NotExtended
	NetworkAuthenticationRequired
	NetworkConnectTimeoutError StatusCode = 599
)
