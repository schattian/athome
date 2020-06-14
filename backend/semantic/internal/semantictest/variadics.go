package semantictest

import "github.com/athomecomar/athome/backend/semantic/schema"

type variadicServiceProviderCategories struct {
	Root     *schema.ServiceProviderCategory
	Branches *variationServiceProviderCategories
	Leaves   *embeddedServiceProviderCategories
}

type embeddedServiceProviderCategories struct {
	First  *variationServiceProviderCategories
	Second *variationServiceProviderCategories
}

type variationServiceProviderCategories struct {
	A *schema.ServiceProviderCategory
	B *schema.ServiceProviderCategory
}

type variadicProductCategories struct {
	Root     *schema.ProductCategory
	Branches *variationProductCategories
	Leaves   *embeddedProductCategories
}

type embeddedProductCategories struct {
	First  *variationProductCategories
	Second *variationProductCategories
}

type variationProductCategories struct {
	A *schema.ProductCategory
	B *schema.ProductCategory
}

type variadicMerchantCategories struct {
	Root     *schema.MerchantCategory
	Branches *variationMerchantCategories
	Leaves   *embeddedMerchantCategories
}

type embeddedMerchantCategories struct {
	First  *variationMerchantCategories
	Second *variationMerchantCategories
}

type variationMerchantCategories struct {
	A *schema.MerchantCategory
	B *schema.MerchantCategory
}
