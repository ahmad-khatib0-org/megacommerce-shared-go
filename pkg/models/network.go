package models

import (
	pb "github.com/ahmad-khatib0-org/megacommerce-proto/gen/go/shared/v1"
	"github.com/ahmad-khatib0-org/megacommerce-shared-go/pkg/utils"
	"github.com/oklog/ulid/v2"
	"google.golang.org/grpc/codes"
)

const (
	HeaderServerName          = "server-name"
	HeaderAuthorization       = "authorization"
	HeaderXRequestID          = "x-request-id"
	HeaderXCorrelationID      = "x-correlation-id"
	HeaderXIPAddress          = "x-ip-address"
	HeaderXForwardedFor       = "x-forwarded-for"
	HeaderXForwardedProto     = "x-forwarded-proto"
	HeaderXForwardedHost      = "x-forwarded-host"
	HeaderXClientVersion      = "x-client-version"
	HeaderXClientID           = "x-client-id"
	HeaderXDeviceID           = "x-device-id"
	HeaderXSessionID          = "x-session-id"
	HeaderXUserID             = "x-user-id"
	HeaderXTraceID            = "x-trace-id"
	HeaderXSpanID             = "x-span-id"
	HeaderXRoles              = "x-roles"
	HeaderXIsOAuth            = "x-is-oauth"
	HeaderXSessionCreatedAt   = "x-session-created-at"
	HeaderXSessionExpiresAt   = "x-session-expires-at"
	HeaderXLastActivityAt     = "x-last-activity-at"
	XTimezone                 = "x-timezone"
	HeaderXProps              = "x-props"
	HeaderXAPIKey             = "x-api-key"
	HeaderXCSRFToken          = "x-csrf-token"
	HeaderXRateLimitLimit     = "x-rate-limit-limit"
	HeaderXRateLimitRemaining = "x-rate-limit-remaining"
	HeaderXRateLimitReset     = "x-rate-limit-reset"
	// Standard headers
	HeaderContentType    = "content-type"
	HeaderUserAgent      = "user-agent"
	HeaderAccept         = "accept"
	HeaderAcceptLanguage = "accept-language"
	HeaderAcceptEncoding = "accept-encoding"
	HeaderCacheControl   = "cache-control"
	// gRPC specific
	HeaderGRPCWeb      = "x-grpc-web"
	HeaderGRPCEncoding = "grpc-encoding"
	HeaderGRPCMessage  = "grpc-message"
	HeaderGRPCStatus   = "grpc-status"
)

// BuildPaginationResponse calculates pagination metadata
// Equivalent to TypeScript's buildPaginationResponse
func BuildPaginationResponse(pagination *pb.PaginationRequest, resultsCount int) *pb.PaginationResponse {
	pageSize := uint32(10) // Default page size
	if pagination != nil && *pagination.PageSize > 0 {
		pageSize = *pagination.PageSize
	}

	hasNext := resultsCount >= int(pageSize)

	page := uint32(1)
	if pagination != nil && *pagination.Page > 0 {
		page = *pagination.Page
	}

	hasPrevious := (page > 1)
	return &pb.PaginationResponse{
		HasNext:           &hasNext,
		HasPrevious:       utils.NewPointer(hasPrevious),
		NextPageToken:     utils.NewPointer(""),
		PreviousPageToken: utils.NewPointer(""),
	}
}

// CheckLastID validates pagination parameters with ULID validation
// Equivalent to TypeScript's checkLastId
func CheckLastID(ctx *Context, where string, pagination *pb.PaginationRequest) *AppError {
	if pagination == nil {
		return NewAppError(
			ctx,
			where,
			"request.pagination.invalid",
			nil,
			"",
			int(codes.InvalidArgument),
			&AppErrorErrorsArgs{},
		)
	}

	lastID := pagination.LastId
	page := pagination.Page

	if *page > 1 && *lastID == "" {
		return NewAppError(
			ctx,
			where,
			"request.last_id.missing",
			nil,
			"",
			int(codes.InvalidArgument),
			&AppErrorErrorsArgs{},
		)
	}

	if *page > 1 && *lastID != "" && !IsValidUlid(*lastID) {
		return NewAppError(
			ctx,
			where,
			"request.last_id.invalid",
			nil,
			"",
			int(codes.InvalidArgument),
			&AppErrorErrorsArgs{},
		)
	}

	return nil
}

// IsValidUlid validates a string as a proper ULID
// Equivalent to TypeScript's isValidUlid
func IsValidUlid(id string) bool {
	if id == "" {
		return false
	}

	// ULID should be exactly 26 characters
	if len(id) != 26 {
		return false
	}

	// Use the ulid library to parse and validate
	_, err := ulid.Parse(id)
	return err == nil
}
