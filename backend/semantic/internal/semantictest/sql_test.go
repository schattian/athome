package semantictest

import (
	"testing"

	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/schema"
	"github.com/athomecomar/storeql/test/sqltest"
)

func TestSQL(t *testing.T) {
	sqltest.SQL(t, &schema.MerchantCategory{}, "MerchantCategory")
	sqltest.SQL(t, &schema.ServiceProviderCategory{}, "ServiceProviderCategory")
	sqltest.SQL(t, &schema.ProductCategory{}, "ProductCategory")

	sqltest.SQL(t, &schema.ProductAttributeSchema{}, "ProductAttributeSchema")
	sqltest.SQL(t, &data.ProductAttributeData{}, "ProductAttributeData")
}
