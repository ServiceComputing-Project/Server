/*
 * simple blog
 *
 * API version: 1.0.0
 */

package swagger

type Body struct {
	Content string `json:"content,omitempty"`
	Author string `json:"author,omitempty"`
}
