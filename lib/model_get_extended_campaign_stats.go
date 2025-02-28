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

type GetExtendedCampaignStats struct {
	// Overall statistics of the campaign
	GlobalStats map[string]interface{} `json:"globalStats"`
	// List-wise statistics of the campaign.
	CampaignStats []interface{} `json:"campaignStats"`
	// Number of clicks on mirror link
	MirrorClick int64 `json:"mirrorClick"`
	// Number of remaning emails to send
	Remaining int64 `json:"remaining"`
	// Statistics about the number of clicks for the links
	LinksStats    map[string]interface{} `json:"linksStats"`
	StatsByDomain *GetStatsByDomain      `json:"statsByDomain"`
	// Statistics about the campaign on the basis of various devices
	StatsByDevice *GetStatsByDevice `json:"statsByDevice"`
	// Statistics about the campaign on the basis of various browsers
	StatsByBrowser *GetStatsByBrowser `json:"statsByBrowser"`
}
