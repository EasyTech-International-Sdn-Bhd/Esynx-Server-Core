package contracts

type ICmsMobileModule interface {
	Get(module string) (string, error)
	SalesmanGroup() (map[string][]string, error)
}
