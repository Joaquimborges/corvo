package corvo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigErrorCases(t *testing.T) {
	for _, tc := range []struct {
		name                 string
		postCard             string
		authorizationCode    string
		urls                 map[EndpointURL]string
		expectedErrorMessage string
	}{
		{
			name:                 "should return authorization code error",
			postCard:             "",
			authorizationCode:    "foo",
			urls:                 make(map[EndpointURL]string),
			expectedErrorMessage: "cartão postagem e código de autorização são obrigatórios",
		},
		{
			name:                 "should return empty url mapper error",
			postCard:             "foo",
			authorizationCode:    "bar",
			urls:                 make(map[EndpointURL]string),
			expectedErrorMessage: "o mapper de urls não pode estar vazio",
		},
		{
			name:                 "should return empty additional services error",
			postCard:             "foo",
			authorizationCode:    "bar",
			urls:                 map[EndpointURL]string{CheckDeliveryProductPriceURL: "/test"},
			expectedErrorMessage: "se você pretende usar a api de preço, serviços adicionais é um parametro obrigatório",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server, err := NewConfig(tc.postCard, tc.authorizationCode, tc.urls)

			require.Nil(t, server)
			require.Error(t, err)
			require.Equal(t, tc.expectedErrorMessage, err.Error())
		})
	}

	t.Run("shoudl return missing product specifications error", func(t *testing.T) {
		urls := map[EndpointURL]string{CheckDeliveryProductPriceURL: "/test"}
		additionalServices := []string{"foo", "bar"}
		expectedErrorMessage := "adicione as espcificações do produto para usar a api de preço [peso, cumprimento, altura, largura]"

		config, err := NewConfig(
			"123",
			"this is a strong code",
			urls,
			ConfigWithCheckPriceAdditionalServices(additionalServices),
		)

		require.Nil(t, config)
		require.Error(t, err)
		require.Equal(t, expectedErrorMessage, err.Error())
	})
}

func TestBuildFullConfigSuccessfully(t *testing.T) {
	urls := map[EndpointURL]string{CheckDeliveryProductPriceURL: "/test"}
	additionalServices := []string{"foo", "bar"}

	config, err := NewConfig(
		"123",
		"this is a strong code",
		urls,
		ConfigWithCheckPriceAdditionalServices(additionalServices),
		ConfigWithFloatPriceEnabled(),
		ConfigWithDeliveryType(1),
		ConfigWithProductSpecification(20, 20, 20, 20),
		ConfigWithOriginZipCode("04376000"),
		ConfigWithDefaultDeclaredValue(300),
	)

	require.NoError(t, err)
	require.NotNil(t, config)
	require.True(t, config.productSpecificationsWasSet)
	require.True(t, config.shouldGenerateFloatPrice)
	require.Equal(t, additionalServices, config.AdditionalServices)

}
