package router

import (
	"encoding/json"
	// "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tonychill/ifitu/lib/utils"
)

func (r *routerImpl) handleGetTask(c *fiber.Ctx) error {

	getTaskURI := "https://api.clickup.com/api/v2/task/86azfarup"
    request := fiber.Get(getTaskURI)
    request.Debug()

    // to set headers
    request.Set("Authorization", "pk_90117669_I71JDAZKKTTHQFJF4V48VRFWRWKDTIM9")
    request.Set("Content-Type", "application/json")


    _, getTaskResponse, err := request.Bytes()
    if err != nil {
        panic(err)
    }

	task:= &Task{}

    jsonErr := json.Unmarshal(getTaskResponse, task)
    if jsonErr != nil {
        panic(err)
    }

	//fmt.Printf("respose from fiber: %d\n%s\n%+v", statusCode, getTaskResponse, err)
	
	var customerEmail string;

	for idx := range task.CustomFields {

        if task.CustomFields[idx].Name == "customer_email" {
            customerEmail = task.CustomFields[idx].Value
        }
    }

	getCheckOut := "https://ifitu.fly.dev/checkout"
    getCheckOutRequest := fiber.Post(getCheckOut)
    request.Debug()

    _, getCheckOutResponse, err := getCheckOutRequest.Bytes()
    if err != nil {
        panic(err)
    }

	checkout:= &Checkout{}

    jsonErr2 := json.Unmarshal(getCheckOutResponse, checkout)
    if jsonErr2 != nil {
        panic(err)
    }

	utils.Send(customerEmail,"Invoice Pending", "Thanks for your purchase.\nTo proceed with the payment, go to the following link: "+checkout.SessionLink)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rules": "test",
		"customer_email": customerEmail,
	})
}


func (r *routerImpl) handleCreateTask(c *fiber.Ctx) error {
	// type something struct {
	// }

	// some := &something{}
	// if err := utils.DecodeFiberRequest(c, some); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"request_id": reqId,
	// 		"message":    fmt.Sprintf("Error decoding client's request: %s", err.Error()),
	// 	})
	// }
	// Your code here
	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"rules": "test",
	// })
	// Create a new Gmail API client
	// client, err := createGmailClient()
	// if err != nil {
	// 	log.Fatalf("Failed to create Gmail client: %v", err)
	// }

	// // Send email
	// err = sendEmail(client, "recipient@example.com", "Subject", "Body")
	// if err != nil {
	// 	log.Fatalf("Failed to send email: %v", err)
	// }

	// r.finImpl.CaptureFunds(c.Context(), &finSvc.CaptureFundsRequest{})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rules": "test",
	})
}

// func createGmailClient() (*gmail.Service, error) {
// 	// Load credentials from a file (replace with your own credentials file)
// 	credentialsFile := "path/to/credentials.json"
// 	config, err := google.ConfigFromJSON(credentialsFile, gmail.MailGoogleComScope)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load credentials: %v", err)
// 	}

// 	// Create a new OAuth2 token source
// 	tokenSource := config.TokenSource(context.Background(), &oauth2.Token{})

// 	// Create a new Gmail client
// 	client, err := gmail.NewService(context.Background(), oauth2.ReuseTokenSource(nil, tokenSource))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create Gmail client: %v", err)
// 	}

// 	return client, nil
// }

// func sendEmail(client *gmail.Service, recipient, subject, body string) error {
// 	// Create the email message
// 	message := gmail.Message{
// 		Raw: base64.URLEncoding.EncodeToString([]byte(
// 			fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", recipient, subject, body),
// 		)),
// 	}

// 	// Send the email
// 	_, err := client.Users.Messages.Send("me", &message).Do()
// 	if err != nil {
// 		return fmt.Errorf("failed to send email: %v", err)
// 	}

// 	return nil
// }
