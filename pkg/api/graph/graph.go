// Package graph provides GraphQL functionality
package graph

import (
	"gorm.io/gorm"

	"github.com/example/graphql-mysql-api/pkg/api/graph/generated"
)

// NewExecutableSchema creates a new GraphQL executable schema
var NewExecutableSchema = generated.NewExecutableSchema

// Resolver holds database connection for GraphQL resolvers
type Resolver struct {
	DB *gorm.DB
}

// NewResolver creates a new resolver instance
func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{DB: db}
}

// Account returns AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Currency returns CurrencyResolver implementation.
func (r *Resolver) Currency() generated.CurrencyResolver { return &currencyResolver{r} }

// Exchange returns ExchangeResolver implementation.
func (r *Resolver) Exchange() generated.ExchangeResolver { return &exchangeResolver{r} }

// Expense returns ExpenseResolver implementation.
func (r *Resolver) Expense() generated.ExpenseResolver { return &expenseResolver{r} }

// Income returns IncomeResolver implementation.
func (r *Resolver) Income() generated.IncomeResolver { return &incomeResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Transaction returns TransactionResolver implementation.
func (r *Resolver) Transaction() generated.TransactionResolver { return &transactionResolver{r} }

// Transfer returns TransferResolver implementation.
func (r *Resolver) Transfer() generated.TransferResolver { return &transferResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// Resolver types
type accountResolver struct{ *Resolver }
type categoryResolver struct{ *Resolver }
type currencyResolver struct{ *Resolver }
type exchangeResolver struct{ *Resolver }
type expenseResolver struct{ *Resolver }
type incomeResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type transactionResolver struct{ *Resolver }
type transferResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// Config represents the GraphQL configuration
type Config = generated.Config
