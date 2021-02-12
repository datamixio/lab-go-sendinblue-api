/*
 * SendinBlue API
 *
 * SendinBlue provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/sendinblue  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |   | 406  | Error. Not Acceptable  | 
 *
 * API version: 3.0.0
 * Contact: contact@sendinblue.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package lib

type GetTransacEmailContent struct {
	// Email address to which transactional email has been sent
	Email string `json:"email"`
	// Subject of the sent email
	Subject string `json:"subject"`
	// Id of the template
	TemplateId int64 `json:"templateId,omitempty"`
	// Date on which transactional email was sent
	Date string `json:"date"`
	// Series of events which occurred on the transactional email
	Events []GetTransacEmailContentEvents `json:"events"`
	// Actual content of the transactional email that has been sent
	Body string `json:"body"`
	// Count of the attachments that were sent in the email
	AttachmentCount int64 `json:"attachmentCount"`
}
