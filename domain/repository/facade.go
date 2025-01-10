package repository

type RepoFacade struct {
	Accounts AccountRepositoryI
	Entries  EntryRepositoryI
}
