package semantictest

import (
	"testing"

	"github.com/athomecomar/athome/backend/semantic/ent"
	"github.com/athomecomar/storeql/test/sqltest"
)

func TestOnboardingIdentificationsSQL(t *testing.T) {
	sqltest.SQL(t, &ent.MerchantCategory{}, "MerchantCategory")
	sqltest.SQL(t, &ent.ServiceProviderCategory{}, "ServiceProviderCategory")
	sqltest.SQL(t, &ent.ProductCategory{}, "ProductCategory")

	sqltest.SQL(t, &ent.ProductAttribute{}, "ProductCategory")
}
