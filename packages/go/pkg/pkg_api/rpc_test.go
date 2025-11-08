package pkgapi_test

import (
	"errors"
	"testing"
	"time"

	"connectrpc.com/connect"
	pkgapi "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_api"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRPC_StructToProto(t *testing.T) {
	t.Parallel()
	// Go struct to convert
	type CustomerStruct struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	type LineItemStruct struct {
		ProductID string  `json:"product_id"`
		Name      string  `json:"name"`
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
	}

	type OrderStatus int8
	const (
		OrderStatusUnspecified OrderStatus = 0
		OrderStatusPending     OrderStatus = 1
		OrderStatusCompleted   OrderStatus = 2
		OrderStatusCancelled   OrderStatus = 3
	)

	type OrderStruct struct {
		OrderID     string           `json:"order_id"`
		Customer    CustomerStruct   `json:"customer"`
		LineItems   []LineItemStruct `json:"line_items"`
		TotalAmount float64          `json:"total_amount"`
		CreatedAt   *time.Time       `json:"created_at"`
		Status      OrderStatus      `json:"status"`
	}

	t.Run("Test conversion with Order struct", func(t *testing.T) {
		// Arrange
		now := time.Now()
		order := OrderStruct{
			OrderID: "order-123",
			Customer: CustomerStruct{
				ID:    "cust-456",
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			LineItems: []LineItemStruct{
				{
					ProductID: "prod-789",
					Name:      "Widget",
					Quantity:  2,
					Price:     19.99,
				},
			},
			TotalAmount: 39.98,
			CreatedAt:   &now,
			Status:      OrderStatusPending,
		}

		// expected protobuf message
		orderProto := &OrderProto{
			OrderId: "order-123",
			Customer: &CustomerProto{
				Id:    "cust-456",
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			LineItems: []*LineItemProto{
				{
					ProductId: "prod-789",
					Name:      "Widget",
					Quantity:  int32(2),
					Price:     19.99,
				},
			},
			TotalAmount: 39.98,
			CreatedAt:   timestamppb.New(now),
			Status:      OrderStatus_ORDER_STATUS_PENDING,
		}

		// Action
		result := &OrderProto{}
		err := pkgapi.ConvertStructToProto(&order, result)

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if orderProto.OrderId != result.OrderId {
			t.Errorf("expected OrderId %s, got %s", orderProto.OrderId, result.OrderId)
		}

		// Assert customer data
		if result.Customer == nil {
			t.Fatal("expected customer to not be nil")
		}
		if orderProto.GetCustomer().GetId() != result.GetCustomer().GetId() {
			t.Errorf("expected customer id %s, got %s", orderProto.GetCustomer().GetId(), result.GetCustomer().GetId())
		}
		if orderProto.GetCustomer().GetName() != result.GetCustomer().GetName() {
			t.Errorf("expected customer name %s, got %s", orderProto.GetCustomer().GetName(), result.GetCustomer().GetName())
		}
		if orderProto.GetCustomer().GetEmail() != result.GetCustomer().GetEmail() {
			t.Errorf("expected customer email %s, got %s", orderProto.GetCustomer().GetEmail(), result.GetCustomer().GetEmail())
		}

		// Assert line items
		if len(result.GetLineItems()) != 1 {
			t.Fatalf("expected 1 line item, got %d", len(result.GetLineItems()))
		}
		if orderProto.GetLineItems()[0].GetProductId() != result.GetLineItems()[0].GetProductId() {
			t.Errorf("expected product id %s, got %s", orderProto.GetLineItems()[0].GetProductId(), result.GetLineItems()[0].GetProductId())
		}
		if orderProto.GetLineItems()[0].GetName() != result.GetLineItems()[0].GetName() {
			t.Errorf("expected line item name %s, got %s", orderProto.GetLineItems()[0].GetName(), result.GetLineItems()[0].GetName())
		}
		if orderProto.GetLineItems()[0].GetQuantity() != result.GetLineItems()[0].GetQuantity() {
			t.Errorf("expected line item quantity %d, got %d", orderProto.GetLineItems()[0].GetQuantity(), result.GetLineItems()[0].GetQuantity())
		}
		if orderProto.GetLineItems()[0].GetPrice() != result.GetLineItems()[0].GetPrice() {
			t.Errorf("expected line item price %f, got %f", orderProto.GetLineItems()[0].GetPrice(), result.GetLineItems()[0].GetPrice())
		}

		// Assert other fields
		if orderProto.GetTotalAmount() != result.GetTotalAmount() {
			t.Errorf("expected total amount %f, got %f", orderProto.GetTotalAmount(), result.GetTotalAmount())
		}
		if result.GetCreatedAt() == nil {
			t.Fatal("expected created at to not be nil")
		}
		if orderProto.GetCreatedAt().AsTime().Unix() != result.GetCreatedAt().AsTime().Unix() {
			t.Errorf("expected created at unix %d, got %d", orderProto.GetCreatedAt().AsTime().Unix(), result.GetCreatedAt().AsTime().Unix())
		}
		if orderProto.GetStatus() != result.GetStatus() {
			t.Errorf("expected status %v, got %v", orderProto.GetStatus(), result.GetStatus())
		}
	})
}

func TestRPC_ProtoToStruct(t *testing.T) {
	// Go struct to convert
	type CustomerStruct struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	type LineItemStruct struct {
		ProductID string  `json:"product_id"`
		Name      string  `json:"name"`
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
	}

	type OrderStatus int8
	const (
		OrderStatusUnspecified OrderStatus = 0
		OrderStatusPending     OrderStatus = 1
		OrderStatusCompleted   OrderStatus = 2
		OrderStatusCancelled   OrderStatus = 3
	)

	type OrderStruct struct {
		OrderID     string           `json:"order_id"`
		Customer    CustomerStruct   `json:"customer"`
		LineItems   []LineItemStruct `json:"line_items"`
		TotalAmount float64          `json:"total_amount"`
		CreatedAt   *time.Time       `json:"created_at"`
		Status      OrderStatus      `json:"status"`
	}

	t.Run("Test conversion with Order proto", func(t *testing.T) {
		// Arrange
		now := time.Now()
		orderProto := &OrderProto{
			OrderId: "order-123",
			Customer: &CustomerProto{
				Id:    "cust-456",
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			LineItems: []*LineItemProto{
				{
					ProductId: "prod-789",
					Name:      "Widget",
					Quantity:  int32(2),
					Price:     19.99,
				},
			},
			TotalAmount: 39.98,
			CreatedAt:   timestamppb.New(now),
			Status:      OrderStatus_ORDER_STATUS_PENDING,
		}

		// expected protobuf message
		order := OrderStruct{
			OrderID: "order-123",
			Customer: CustomerStruct{
				ID:    "cust-456",
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			LineItems: []LineItemStruct{
				{
					ProductID: "prod-789",
					Name:      "Widget",
					Quantity:  2,
					Price:     19.99,
				},
			},
			TotalAmount: 39.98,
			CreatedAt:   &now,
			Status:      OrderStatusPending,
		}

		// Action
		result := &OrderStruct{}
		err := pkgapi.ConvertProtoToStruct(orderProto, result)

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if order.OrderID != result.OrderID {
			t.Errorf("expected OrderID %s, got %s", order.OrderID, result.OrderID)
		}

		// Assert customer data
		if result.Customer != (CustomerStruct{}) {
			if order.Customer.ID != result.Customer.ID {
				t.Errorf("expected customer id %s, got %s", order.Customer.ID, result.Customer.ID)
			}
			if order.Customer.Name != result.Customer.Name {
				t.Errorf("expected customer name %s, got %s", order.Customer.Name, result.Customer.Name)
			}
			if order.Customer.Email != result.Customer.Email {
				t.Errorf("expected customer email %s, got %s", order.Customer.Email, result.Customer.Email)
			}
		} else {
			t.Fatal("expected customer to not be zero value")
		}

		// Assert line items
		if len(result.LineItems) != 1 {
			t.Fatalf("expected 1 line item, got %d", len(result.LineItems))
		}
		if len(result.LineItems) > 0 {
			if order.LineItems[0].ProductID != result.LineItems[0].ProductID {
				t.Errorf("expected product id %s, got %s", order.LineItems[0].ProductID, result.LineItems[0].ProductID)
			}
			if order.LineItems[0].Name != result.LineItems[0].Name {
				t.Errorf("expected line item name %s, got %s", order.LineItems[0].Name, result.LineItems[0].Name)
			}
			if order.LineItems[0].Quantity != result.LineItems[0].Quantity {
				t.Errorf("expected line item quantity %d, got %d", order.LineItems[0].Quantity, result.LineItems[0].Quantity)
			}
			if order.LineItems[0].Price != result.LineItems[0].Price {
				t.Errorf("expected line item price %f, got %f", order.LineItems[0].Price, result.LineItems[0].Price)
			}
		}

		// Assert other fields
		if order.TotalAmount != result.TotalAmount {
			t.Errorf("expected total amount %f, got %f", order.TotalAmount, result.TotalAmount)
		}
		if result.CreatedAt == nil {
			t.Fatal("expected created at to not be nil")
		}
		if order.CreatedAt.Unix() != result.CreatedAt.Unix() {
			t.Errorf("expected created at unix %d, got %d", order.CreatedAt.Unix(), result.CreatedAt.Unix())
		}
		if order.Status != result.Status {
			t.Errorf("expected status %v, got %v", order.Status, result.Status)
		}
	})
}

func TestRPC_ResOK(t *testing.T) {
	type TestUser struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	t.Run("should return successful response with string message", func(t *testing.T) {
		message := "Hello, World!"
		resp, err := pkgapi.ResOK(&message)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if resp == nil {
			t.Fatal("Expected response to not be nil")
		}

		if resp.Msg == nil {
			t.Fatal("Expected response message to not be nil")
		}

		if *resp.Msg != message {
			t.Errorf("Expected message %s, got %s", message, *resp.Msg)
		}
	})

	t.Run("should return successful response with struct message", func(t *testing.T) {
		user := &TestUser{
			ID:   1,
			Name: "John Doe",
		}
		resp, err := pkgapi.ResOK(user)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if resp == nil {
			t.Fatal("Expected response to not be nil")
		}

		if resp.Msg == nil {
			t.Fatal("Expected response message to not be nil")
		}

		if resp.Msg.ID != user.ID {
			t.Errorf("Expected ID %d, got %d", user.ID, resp.Msg.ID)
		}

		if resp.Msg.Name != user.Name {
			t.Errorf("Expected name %s, got %s", user.Name, resp.Msg.Name)
		}
	})

	t.Run("should return successful response with nil message", func(t *testing.T) {
		var message *string = nil
		resp, err := pkgapi.ResOK(message)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if resp == nil {
			t.Fatal("Expected response to not be nil")
		}

		if resp.Msg != nil {
			t.Errorf("Expected message to be nil, got %v", resp.Msg)
		}
	})
}

func TestRPC_ResFailed(t *testing.T) {
	type TestUser struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	t.Run("should return error when given standard error", func(t *testing.T) {
		originalErr := errors.New("something went wrong")
		resp, err := pkgapi.ResFailed[string](originalErr)

		if resp != nil {
			t.Errorf("Expected response to be nil, got %v", resp)
		}

		if err == nil {
			t.Fatal("Expected error to not be nil")
		}

		// Check if the error is a connect error
		var connectErr *connect.Error
		if !errors.As(err, &connectErr) {
			t.Fatal("Expected error to be a connect.Error")
		}

		// Should be mapped to internal error by default
		if connectErr.Code() != connect.CodeInternal {
			t.Errorf("Expected error code %v, got %v", connect.CodeInternal, connectErr.Code())
		}
	})

	t.Run("should return error when given app error", func(t *testing.T) {
		appErr := pkgerr.Error{
			Code:    pkgerr.CodeNotFoundError,
			Message: "user not found",
			Key:     "USER_NOT_FOUND",
		}
		resp, err := pkgapi.ResFailed[TestUser](&appErr)

		if resp != nil {
			t.Errorf("Expected response to be nil, got %v", resp)
		}

		if err == nil {
			t.Fatal("Expected error to not be nil")
		}

		// Check if the error is a connect error
		var connectErr *connect.Error
		if !errors.As(err, &connectErr) {
			t.Fatal("Expected error to be a connect.Error")
		}

		// Should be mapped to not found error
		if connectErr.Code() != connect.CodeNotFound {
			t.Errorf("Expected error code %v, got %v", connect.CodeNotFound, connectErr.Code())
		}
	})

	t.Run("should return error when given validation app error", func(t *testing.T) {
		appErr := &pkgerr.Error{
			Code:    pkgerr.CodeValidationError,
			Message: "invalid input",
			Key:     "VALIDATION_ERROR",
		}
		resp, err := pkgapi.ResFailed[int](appErr)

		if resp != nil {
			t.Errorf("Expected response to be nil, got %v", resp)
		}

		if err == nil {
			t.Fatal("Expected error to not be nil")
		}

		// Check if the error is a connect error
		var connectErr *connect.Error
		if !errors.As(err, &connectErr) {
			t.Fatal("Expected error to be a connect.Error")
		}

		// Should be mapped to invalid argument error
		if connectErr.Code() != connect.CodeInvalidArgument {
			t.Errorf("Expected error code %v, got %v", connect.CodeInvalidArgument, connectErr.Code())
		}
	})

	t.Run("should handle different generic types", func(t *testing.T) {
		originalErr := errors.New("test error")

		// Test with string type
		respStr, errStr := pkgapi.ResFailed[string](originalErr)
		if respStr != nil {
			t.Error("Expected string response to be nil")
		}
		if errStr == nil {
			t.Error("Expected string error to not be nil")
		}

		// Test with int type
		respInt, errInt := pkgapi.ResFailed[int](originalErr)
		if respInt != nil {
			t.Error("Expected int response to be nil")
		}
		if errInt == nil {
			t.Error("Expected int error to not be nil")
		}

		// Test with struct type
		respStruct, errStruct := pkgapi.ResFailed[TestUser](originalErr)
		if respStruct != nil {
			t.Error("Expected struct response to be nil")
		}
		if errStruct == nil {
			t.Error("Expected struct error to not be nil")
		}
	})
}
