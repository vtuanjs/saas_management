package pkgapi

import (
	"encoding/json"

	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func codeMapping(code pkgerr.Code) connect.Code {
	switch code {
	case pkgerr.CodeValidationError:
		return connect.CodeInvalidArgument
	case pkgerr.CodeNotFoundError:
		return connect.CodeNotFound
	case pkgerr.CodeUnauthorizedError:
		return connect.CodeUnauthenticated
	case pkgerr.CodeForbiddenError:
		return connect.CodePermissionDenied
	case pkgerr.CodeDuplicateError:
		return connect.CodeAlreadyExists
	case pkgerr.CodeConflictError:
		return connect.CodeFailedPrecondition
	case pkgerr.CodeTimeoutError:
		return connect.CodeDeadlineExceeded
	case pkgerr.CodeInternalError:
		return connect.CodeInternal
	case pkgerr.CodeTransactionError:
		return connect.CodeInternal
	case pkgerr.CodeCancelledError:
		return connect.CodeCanceled
	case pkgerr.CodeRetryableError:
		return connect.CodeUnavailable
	case pkgerr.CodeUnknownError:
		return connect.CodeUnknown
	case pkgerr.CodeAbortedError:
		return connect.CodeAborted
	default:
		return connect.CodeUnknown
	}
}

// NewError converts a ServerError to a connect.Error.
// If the error is nil, it returns nil.
// This function is used to convert application-specific errors to connect errors.
func NewRPCError(err error) *connect.Error {
	if err == nil {
		return nil
	}

	if appErr, ok := pkgerr.Check(err); ok {
		return connect.NewError(codeMapping(appErr.GetCode()), appErr)
	}

	appErr := pkgerr.Error{
		Code:    pkgerr.CodeInternalError,
		Message: err.Error(),
		Root:    err,
	}
	return connect.NewError(connect.CodeInternal, appErr)
}

func ResOK[T any](message *T) (*connect.Response[T], error) {
	return connect.NewResponse(message), nil
}

func ResFailed[T any](err error) (*connect.Response[T], error) {
	return nil, NewRPCError(err)
}

// ConvertStructToProto converts a Go struct to a protobuf message.
// sourceStruct: the Go struct to convert
// destProto: a pointer to the destination protobuf message (must be initialized)
// Returns an error if the conversion fails
//
// IMPORTANT: You must define enum as type int, int8, int16... in the Go struct
// to ensure proper conversion to protobuf enum types.
// Example usage:
//
//	type MyStruct struct {
//	    Field1 string `json:"field1"`
//	    Field2 int32  `json:"field2"`
//	    Field3 MyEnum `json:"field3"` // MyEnum should be defined as an int, int8, int16...
//	}
//	type MyEnum int32 // Define your enum as an int, int8, int16...
//	const (
//	    MyEnumUnspecified MyEnum = 0
//	    MyEnumValue1 MyEnum = 1
//	    MyEnumValue2 MyEnum = 2
//	)
func ConvertStructToProto(sourceStruct interface{}, destProto proto.Message) error {
	// Convert Go struct to JSON
	jsonData, err := json.Marshal(sourceStruct)
	if err != nil {
		return pkgerr.NewInternalError("ConvertStructToProto", "error marshaling struct to JSON: ", err)
	}

	// Convert JSON to protobuf using protojson
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true, // Ignore JSON fields that don't exist in the proto
		AllowPartial:   true, // Allow partial messages
	}

	if err := unmarshaler.Unmarshal(jsonData, destProto); err != nil {
		return pkgerr.NewInternalError("ConvertStructToProto", "error unmarshaling JSON to proto: ", err)
	}

	return nil
}

// ConvertProtoToStruct converts a protobuf message to a Go struct.
// sourceProto: the protobuf message to convert
// destStruct: a pointer to the destination Go struct
// Returns an error if the conversion fails
//
// IMPORTANT: Enum of the Go struct must be defined as int, int8, int16... to ensure proper conversion
// to protobuf enum types.
func ConvertProtoToStruct(sourceProto proto.Message, destStruct interface{}) error {
	// Convert protobuf to JSON
	marshaler := protojson.MarshalOptions{
		UseProtoNames:  true, // Use proto field names instead of lowerCamelCase
		UseEnumNumbers: true,
	}

	jsonData, err := marshaler.Marshal(sourceProto)
	if err != nil {
		return pkgerr.NewInternalError("ConvertProtoToStruct", "error marshaling proto to JSON: ", err)
	}

	// Convert JSON to Go struct
	if err := json.Unmarshal(jsonData, destStruct); err != nil {
		return pkgerr.NewInternalError("ConvertProtoToStruct", "error unmarshaling JSON to struct: ", err)
	}

	return nil
}
