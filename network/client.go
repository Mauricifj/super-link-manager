package network

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/super-link-manager/models"
	"github.com/super-link-manager/utils"
	"os"
	"strings"
)

var baseUrl = os.Getenv("BASEURL")

func GetAccessToken() *models.OAuthResponse {
	var username = os.Getenv("USERNAME")
	var password = os.Getenv("PASSWORD")

	request := resty.New().R()

	request.SetBasicAuth(username, password)

	request.SetFormData(map[string]string{
		"grant_type": "client_credentials",
	})

	request.SetResult(models.OAuthResponse{})

	url := baseUrl + os.Getenv("OAUTH")

	response, err := request.Post(url)
	utils.CheckErr(err)

	if response.StatusCode() == 201 {
		OAuthResponse := response.Result().(*models.OAuthResponse)

		fmt.Println("AUTHENTICATION STEP")
		fmt.Println(" - ACCESS TOKEN:", OAuthResponse.AccessToken[0:60], "...")
		fmt.Println(" - TOKEN TYPE:", OAuthResponse.TokenType)
		fmt.Println(" - EXPIRES IN:", OAuthResponse.ExpiresIn)
		fmt.Println()

		return OAuthResponse
	}
	return nil
}

func CreateLink(linkType, name string, price int) models.Response {
	request := resty.New().R()

	request.SetAuthToken(GetAccessToken().AccessToken)

	request.SetHeader("Content-Type", "application/json")

	body := models.Link{
		LinkType:                linkType,
		Name:                    name,
		Description:             "My description of this sample order",
		Price:                   price,
		Weight:                  1000,
		ExpirationDate:          "2030-12-31",
		MaxNumberOfInstallments: 1,
		Quantity:                1,
		Sku:                     "SKU-TEST",
		SoftDescriptor:          "STOREMAURICI",
		Shipping: 				 models.Shipping{
			Type:  				 "WithoutShipping",
			Name:  				 "No Shipping",
			Price: 				 "0",
		},
	}

	request.SetBody(body)

	request.SetResult(models.Link{})
	request.SetError(models.ErrorResponse{})

	url := baseUrl + os.Getenv("NEW")

	response, err := request.Post(url)
	utils.CheckErr(err)

	if response.StatusCode() == 201 {
		link := response.Result().(*models.Link)

		fmt.Println("NEW LINK CREATED")
		fmt.Println(" - ID:", link.Id)
		fmt.Println(" - TYPE:", link.LinkType)
		fmt.Println(" - NAME:", link.Name)
		fmt.Println(" - PRICE:", link.Price)
		fmt.Println()

		return models.Response{ Link: link }
	} else if response.StatusCode() == 400 {
		errors := response.Error().(*models.ErrorResponse)

		fmt.Println("ERRORS ON LINK CREATION")
		for _, errorItem := range *errors {
			fmt.Println("FIELD: " , errorItem.Field)
			fmt.Println("MESSAGE: ", errorItem.Message)
		}

		return models.Response{Errors: errors}
	}
	return models.Response{}
}

func UpdateLink(id, linkType, name string, price int) models.Response {
	request := resty.New().R()

	request.SetAuthToken(GetAccessToken().AccessToken)

	request.SetHeader("Content-Type", "application/json")

	body := models.Link{
		LinkType:                linkType,
		Name:                    name,
		Description:             "My description of this sample order",
		Price:                   price,
		Weight:                  1000,
		ExpirationDate:          "2030-12-31",
		MaxNumberOfInstallments: 1,
		Quantity:                1,
		Sku:                     "SKU-TEST",
		SoftDescriptor:          "STOREMAURICI",
		Shipping: 				 models.Shipping{
			Type:  				 "WithoutShipping",
			Name:  				 "No Shipping",
			Price: 				 "0",
		},
	}

	request.SetBody(body)

	request.SetResult(models.Link{})

	url := baseUrl + strings.Replace(os.Getenv("UPDATE"), "ID", id, -1)

	response, err := request.Put(url)
	utils.CheckErr(err)

	if response.StatusCode() == 200 {
		link := response.Result().(*models.Link)

		fmt.Println("LINK UPDATED")
		fmt.Println(" - ID:", link.Id)
		fmt.Println(" - TYPE:", link.LinkType)
		fmt.Println(" - NAME:", link.Name)
		fmt.Println(" - PRICE:", link.Price)
		fmt.Println()

		return models.Response{Link: link}
	} else {
		if response.StatusCode() == 400 {
			errors := response.Error().(*models.ErrorResponse)

			fmt.Println("ERRORS ON LINK CREATION")
			for _, errorItem := range *errors {
				fmt.Println("FIELD: " , errorItem.Field)
				fmt.Println("MESSAGE: ", errorItem.Message)
			}

			return models.Response{Errors: errors}
		}
	}
	return models.Response{}
}

func LinkById(id string) *models.Link {
	request := resty.New().R()

	request.SetAuthToken(GetAccessToken().AccessToken)

	request.SetHeader("Content-Type", "application/json")

	request.SetResult(models.Link{})

	url := baseUrl + strings.Replace(os.Getenv("QUERY"), "ID", id, -1)

	response, err := request.Get(url)
	utils.CheckErr(err)

	if response.StatusCode() == 200 {
		Link := response.Result().(*models.Link)

		fmt.Println("LINK FOUND")
		fmt.Println(" - ID:", Link.Id)
		fmt.Println(" - TYPE:", Link.LinkType)
		fmt.Println(" - NAME:", Link.Name)
		fmt.Println(" - PRICE:", Link.Price)
		fmt.Println()

		return Link
	}
	return nil
}

func DeleteLink(id string) bool {
	request := resty.New().R()

	request.SetAuthToken(GetAccessToken().AccessToken)

	request.SetHeader("Content-Type", "application/json")

	url := baseUrl + strings.Replace(os.Getenv("DELETE"), "ID", id, -1)

	response, err := request.Delete(url)
	utils.CheckErr(err)

	if response.StatusCode() == 204 {
		fmt.Println("LINK DELETED")
		fmt.Println(" - ID:", id)
		fmt.Println()
		return true
	}

	return false
}