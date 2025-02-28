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

import (
	"time"
)

// Note Details
type Note struct {
	// Unique note Id
	Id string `json:"id,omitempty"`
	// Text content of a note
	Text string `json:"text"`
	// Contact ids linked to a note
	ContactIds []int32 `json:"contactIds,omitempty"`
	// Deal ids linked to a note
	DealIds []string `json:"dealIds,omitempty"`
	// Account details of user which created the note
	AuthorId map[string]interface{} `json:"authorId,omitempty"`
	// Note created date/time
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Note updated date/time
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
