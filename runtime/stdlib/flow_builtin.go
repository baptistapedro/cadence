package stdlib

import (
	"github.com/dapperlabs/flow-go/language/runtime/ast"
	"github.com/dapperlabs/flow-go/language/runtime/interpreter"
	"github.com/dapperlabs/flow-go/language/runtime/sema"
)

// This file defines functions built in to the Flow runtime.

var flowLocation = ast.StringLocation("flow")

// built-in function types

// TODO: improve types
var createAccountFunctionType = &sema.FunctionType{
	ParameterTypeAnnotations: sema.NewTypeAnnotations(
		// publicKeys
		&sema.VariableSizedType{
			Type: &sema.VariableSizedType{
				Type: &sema.IntType{},
			},
		},
		// code
		&sema.OptionalType{
			Type: &sema.VariableSizedType{
				Type: &sema.IntType{},
			},
		},
	),
	// address
	ReturnTypeAnnotation: sema.NewTypeAnnotation(
		&sema.AddressType{},
	),
	// additional arguments are passed to the contract initializer
	RequiredArgumentCount: (func() *int {
		var count = 2
		return &count
	})(),
}

var addAccountKeyFunctionType = &sema.FunctionType{
	ParameterTypeAnnotations: sema.NewTypeAnnotations(
		// address
		&sema.AddressType{},
		// key
		&sema.VariableSizedType{
			Type: &sema.IntType{},
		},
	),
	// nothing
	ReturnTypeAnnotation: sema.NewTypeAnnotation(
		&sema.VoidType{},
	),
}

var removeAccountKeyFunctionType = &sema.FunctionType{
	ParameterTypeAnnotations: sema.NewTypeAnnotations(
		// address
		&sema.AddressType{},
		// index
		&sema.IntType{},
	),
	// nothing
	ReturnTypeAnnotation: sema.NewTypeAnnotation(
		&sema.VoidType{},
	),
}

var updateAccountCodeFunctionType = &sema.FunctionType{
	ParameterTypeAnnotations: sema.NewTypeAnnotations(
		// address
		&sema.AddressType{},
		// code
		&sema.VariableSizedType{
			Type: &sema.IntType{},
		},
	),
	// nothing
	ReturnTypeAnnotation: sema.NewTypeAnnotation(
		&sema.VoidType{},
	),
	// additional arguments are passed to the contract initializer
	RequiredArgumentCount: (func() *int {
		var count = 2
		return &count
	})(),
}

var getAccountFunctionType = &sema.FunctionType{
	ParameterTypeAnnotations: sema.NewTypeAnnotations(
		// address
		&sema.AddressType{},
	),
	ReturnTypeAnnotation: sema.NewTypeAnnotation(
		&sema.PublicAccountType{},
	),
}

var logFunctionType = &sema.FunctionType{
	ParameterTypeAnnotations: sema.NewTypeAnnotations(
		&sema.AnyStructType{},
	),
	ReturnTypeAnnotation: sema.NewTypeAnnotation(
		&sema.VoidType{},
	),
}

// FlowBuiltinImpls defines the set of functions needed to implement the Flow
// built-in functions.
type FlowBuiltinImpls struct {
	CreateAccount     interpreter.HostFunction
	AddAccountKey     interpreter.HostFunction
	RemoveAccountKey  interpreter.HostFunction
	UpdateAccountCode interpreter.HostFunction
	GetAccount        interpreter.HostFunction
	Log               interpreter.HostFunction
}

// FlowBuiltInFunctions returns a list of standard library functions, bound to
// the provided implementation.
func FlowBuiltInFunctions(impls FlowBuiltinImpls) StandardLibraryFunctions {
	return StandardLibraryFunctions{
		NewStandardLibraryFunction(
			"createAccount",
			createAccountFunctionType,
			impls.CreateAccount,
			nil,
		),
		NewStandardLibraryFunction(
			"addAccountKey",
			addAccountKeyFunctionType,
			impls.AddAccountKey,
			nil,
		),
		NewStandardLibraryFunction(
			"removeAccountKey",
			removeAccountKeyFunctionType,
			impls.RemoveAccountKey,
			nil,
		),
		NewStandardLibraryFunction(
			"updateAccountCode",
			updateAccountCodeFunctionType,
			impls.UpdateAccountCode,
			nil,
		),
		NewStandardLibraryFunction(
			"getAccount",
			getAccountFunctionType,
			impls.GetAccount,
			nil,
		),
		NewStandardLibraryFunction(
			"log",
			logFunctionType,
			impls.Log,
			nil,
		),
	}
}

// built-in event types

var AccountCreatedEventType = sema.EventType{
	Location:   flowLocation,
	Identifier: "AccountCreated",
	Fields: []sema.EventFieldType{
		{
			Identifier: "address",
			Type:       &sema.StringType{},
		},
	},
	ConstructorParameterTypeAnnotations: []*sema.TypeAnnotation{
		{
			Move: false,
			Type: &sema.StringType{},
		},
	},
}

var AccountKeyAddedEventType = sema.EventType{
	Location:   flowLocation,
	Identifier: "AccountKeyAdded",
	Fields: []sema.EventFieldType{
		{
			Identifier: "address",
			Type:       &sema.StringType{},
		},
		{
			Identifier: "publicKey",
			Type: &sema.VariableSizedType{
				Type: &sema.IntType{},
			},
		},
	},
	ConstructorParameterTypeAnnotations: []*sema.TypeAnnotation{
		{
			Move: false,
			Type: &sema.StringType{},
		},
		{
			Move: false,
			Type: &sema.VariableSizedType{
				Type: &sema.IntType{},
			},
		},
	},
}

var AccountKeyRemovedEventType = sema.EventType{
	Location:   flowLocation,
	Identifier: "AccountKeyRemoved",
	Fields: []sema.EventFieldType{
		{
			Identifier: "address",
			Type:       &sema.StringType{},
		},
		{
			Identifier: "publicKey",
			Type: &sema.VariableSizedType{
				Type: &sema.IntType{},
			},
		},
	},
	ConstructorParameterTypeAnnotations: []*sema.TypeAnnotation{
		{
			Move: false,
			Type: &sema.StringType{},
		},
		{
			Move: false,
			Type: &sema.VariableSizedType{
				Type: &sema.IntType{},
			},
		},
	},
}

var AccountCodeUpdatedEventType = sema.EventType{
	Location:   flowLocation,
	Identifier: "AccountCodeUpdated",
	Fields: []sema.EventFieldType{
		{
			Identifier: "address",
			Type:       &sema.StringType{},
		},
		{
			Identifier: "codeHash",
			Type: &sema.VariableSizedType{
				Type: &sema.IntType{},
			},
		},
	},
	ConstructorParameterTypeAnnotations: []*sema.TypeAnnotation{
		{
			Move: false,
			Type: &sema.StringType{},
		},
		{
			Move: false,
			Type: &sema.VariableSizedType{
				Type: &sema.IntType{},
			},
		},
	},
}
