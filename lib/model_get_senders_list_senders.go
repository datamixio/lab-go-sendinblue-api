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

type GetSendersListSenders struct {
	// Id of the sender
	Id int64 `json:"id"`
	// From Name associated to the sender
	Name string `json:"name"`
	// From Email associated to the sender
	Email string `json:"email"`
	// Status of sender (true=activated, false=deactivated)
	Active bool `json:"active"`
	// List of dedicated IP(s) available in the account. This data is displayed only for dedicated IPs
	Ips []GetSendersListIps `json:"ips,omitempty"`
}
