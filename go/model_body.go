/*
 * simple blog
 *
 * API version: 1.0.0
 */

package swagger

type Body struct {
	Author string `json:"author,omitempty"`
	Content string `json:"content,omitempty"`
}
