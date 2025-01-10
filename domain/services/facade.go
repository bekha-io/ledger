package services

type ServiceFacade struct {
	Accounts AccountServiceI
	Entries  EntryServiceI
}
