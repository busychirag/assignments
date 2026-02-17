package routes

import (
	"github.com/busychirag/assignments/tree/main/Assignment2/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// Bank routes
		banks := api.Group("/banks")
		{
			banks.GET("", controllers.GetAllBanks)
			banks.POST("", controllers.CreateBank)
		}
		bank := api.Group("/bank")
		{
			bank.GET("/:id", controllers.GetBankByID)
			bank.PUT("/:id", controllers.UpdateBank)
			bank.DELETE("/:id", controllers.DeleteBank)
			bank.GET("/:id/branches", controllers.GetBranchesByBankID)
		}

		// Branch routes
		branches := api.Group("/branches")
		{
			branches.GET("", controllers.GetAllBranches)
			branches.POST("", controllers.CreateBranch)
		}
		branch := api.Group("/branch")
		{
			branch.GET("/:id", controllers.GetBranchByID)
			branch.PUT("/:id", controllers.UpdateBranch)
			branch.DELETE("/:id", controllers.DeleteBranch)
		}

		// Customer routes
		customers := api.Group("/customers")
		{
			customers.GET("", controllers.GetAllCustomers)
			customers.POST("", controllers.CreateCustomer)
		}
		customer := api.Group("/customer")
		{
			customer.GET("/:id", controllers.GetCustomerByID)
			customer.PUT("/:id", controllers.UpdateCustomer)
			customer.DELETE("/:id", controllers.DeleteCustomer)
			customer.GET("/:id/accounts", controllers.GetAccountsByCustomerID)
			customer.GET("/:id/loans", controllers.GetLoansByCustomerID)
		}

		// Account routes
		accounts := api.Group("/accounts")
		{
			accounts.GET("", controllers.GetAllAccounts)
			accounts.POST("", controllers.CreateAccount)
		}
		account := api.Group("/account")
		{
			account.GET("/:id", controllers.GetAccountByID)
			account.POST("/:id/deposit", controllers.Deposit)
			account.POST("/:id/withdraw", controllers.Withdraw)
			account.GET("/:id/transactions", controllers.GetTransactionsByAccountID)
			account.GET("/:id/holders", controllers.GetAccountHolders)
			account.POST("/:id/holders", controllers.AddJointHolder)
			account.DELETE("/:id/holders", controllers.RemoveJointHolder)
		}

		// Loan routes
		loans := api.Group("/loans")
		{
			loans.GET("", controllers.GetAllLoans)
			loans.POST("", controllers.CreateLoan)
		}
		loan := api.Group("/loan")
		{
			loan.GET("/:id", controllers.GetLoanByID)
			loan.POST("/:id/repay", controllers.RepayLoan)
		}
	}

	return r
}
