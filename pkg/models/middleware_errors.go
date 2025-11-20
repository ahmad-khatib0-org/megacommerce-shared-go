package models

import (
	"google.golang.org/grpc/codes"
)

func GetGrpcErrMsg(lang string, code codes.Code) string {
	switch code {
	case codes.Canceled:
		return Tr(lang, "error.canceled", nil)
	case codes.Unknown:
		return Tr(lang, "error.unknown", nil)
	case codes.InvalidArgument:
		return Tr(lang, "error.invalid_argument", nil)
	case codes.DeadlineExceeded:
		return Tr(lang, "error.deadline_exceeded", nil)
	case codes.NotFound:
		return Tr(lang, "error.not_found", nil)
	case codes.AlreadyExists:
		return Tr(lang, "error.already_exists", nil)
	case codes.PermissionDenied:
		return Tr(lang, "error.permission_denied", nil)
	case codes.ResourceExhausted:
		return Tr(lang, "error.resource_exhausted", nil)
	case codes.FailedPrecondition:
		return Tr(lang, "error.failed_precondition", nil)
	case codes.Aborted:
		return Tr(lang, "error.aborted", nil)
	case codes.OutOfRange:
		return Tr(lang, "error.out_of_range", nil)
	case codes.Unimplemented:
		return Tr(lang, "error.unimplemented", nil)
	case codes.Internal:
		return Tr(lang, "error.internal", nil)
	case codes.Unavailable:
		return Tr(lang, "error.unavailable", nil)
	case codes.DataLoss:
		return Tr(lang, "error.data_loss", nil)
	case codes.Unauthenticated:
		return Tr(lang, "error.unauthenticated", nil)
	}

	return Tr(lang, "error.unknown", nil)
}
