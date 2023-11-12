package customerimporter

// EmailDomainCustomerCount stores email domain, nd how many customers it has
type EmailDomainCustomerCount struct {
	EmailDomain   string
	CustomerCount int
}

// EmailDomainCustomerCountList created for later sorting EmailDomainCustomerCount by domain
type EmailDomainCustomerCountList []EmailDomainCustomerCount

// Below methods o satisfy sort.Interface
func (edccList EmailDomainCustomerCountList) Len() int {
	return len(edccList)
}
func (edccList EmailDomainCustomerCountList) Less(i, j int) bool {
	return edccList[i].EmailDomain < edccList[j].EmailDomain
}
func (edccList EmailDomainCustomerCountList) Swap(i, j int) {
	edccList[i], edccList[j] = edccList[j], edccList[i]
}

// Customer stores only email, because we only care about extracting email's domain and count occurences
type Customer struct {
	Email string
}
