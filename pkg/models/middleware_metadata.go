package models

import (
	"context"
	"strconv"

	"github.com/ahmad-khatib0-org/megacommerce-shared-go/pkg/utils"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryMetadataInterceptor(defaultLangs string, availableLanguages []string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx = context.WithValue(ctx, ContextKeyMethodName, info.FullMethod)

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			ctx = ExtractMetadataToContext(ctx, md, defaultLangs, availableLanguages)
		}

		return handler(ctx, req)
	}
}

func StreamMetadataInterceptor(defaultLangs string, availableLanguages []string) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := context.WithValue(ss.Context(), ContextKeyMethodName, info.FullMethod)

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			ctx = ExtractMetadataToContext(ctx, md, defaultLangs, availableLanguages)
		}

		wrapped := middleware.WrapServerStream(ss)
		wrapped.WrappedContext = ctx

		return handler(srv, wrapped)
	}
}

func ExtractMetadataToContext(ctx context.Context, md metadata.MD, defaultLangs string, availableLanguages []string) context.Context {
	mc := &Context{}
	mc.Session = &Session{}

	if vals := md.Get(HeaderUserAgent); len(vals) > 0 {
		mc.UserAgent = vals[0]
	}
	if vals := md.Get(HeaderXRequestID); len(vals) > 0 {
		mc.RequestID = vals[0]
	}
	if vals := md.Get(HeaderAuthorization); len(vals) > 0 {
		mc.Session.Token = vals[0]
	}
	if vals := md.Get(HeaderXIPAddress); len(vals) > 0 {
		mc.IPAddress = vals[0]
	}
	if vals := md.Get(HeaderXForwardedFor); len(vals) > 0 {
		mc.XForwardedFor = vals[0]
	}
	if vals := md.Get(HeaderAcceptLanguage); len(vals) > 0 {
		mc.AcceptLanguage = utils.ProcessAcceptedLanguage(vals[0], availableLanguages, defaultLangs)
	} else {
		mc.AcceptLanguage = defaultLangs
	}
	if vals := md.Get(XTimezone); len(vals) > 0 {
		mc.Timezone = vals[0]
	}

	// Updated session headers with x- prefix
	if vals := md.Get(HeaderXSessionID); len(vals) > 0 {
		mc.Session.ID = vals[0]
	}
	// Removed HeaderToken - using Authorization instead
	if vals := md.Get(HeaderXSessionCreatedAt); len(vals) > 0 {
		if val, err := strconv.Atoi(vals[0]); err == nil {
			mc.Session.CreatedAt = int64(val)
		}
	}
	if vals := md.Get(HeaderXSessionExpiresAt); len(vals) > 0 {
		if val, err := strconv.Atoi(vals[0]); err == nil {
			mc.Session.ExpiresAt = int64(val)
		}
	}
	if vals := md.Get(HeaderXLastActivityAt); len(vals) > 0 {
		if val, err := strconv.Atoi(vals[0]); err == nil {
			mc.Session.LastActivityAt = int64(val)
		}
	}
	if vals := md.Get(HeaderXUserID); len(vals) > 0 {
		mc.Session.UserID = vals[0]
	}
	if vals := md.Get(HeaderXDeviceID); len(vals) > 0 {
		mc.Session.DeviceID = vals[0]
	}
	if vals := md.Get(HeaderXRoles); len(vals) > 0 {
		mc.Session.Roles = vals[0]
	}
	if vals := md.Get(HeaderXIsOAuth); len(vals) > 0 {
		mc.Session.IsOAuth = vals[0] == "true"
	}
	if vals := md.Get(HeaderXProps); len(vals) > 0 {
		mc.Session.Props = utils.GetMetadataValue(vals)
	}
	if vals := md.Get(HeaderServerName); len(vals) > 0 {
		mc.ServerName = vals[0]
	}

	// if vals := md.Get(HeaderXCorrelationID); len(vals) > 0 {
	// 	mc.CorrelationID = vals[0]
	// }
	// if vals := md.Get(HeaderXClientVersion); len(vals) > 0 {
	// 	mc.ClientVersion = vals[0]
	// }
	// if vals := md.Get(HeaderXTraceID); len(vals) > 0 {
	// 	mc.TraceID = vals[0]
	// }

	mc.Context = ctx
	return context.WithValue(ctx, ContextKeyMetadata, mc)
}
