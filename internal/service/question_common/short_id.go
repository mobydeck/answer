package questioncommon

import "context"

// isShortLinkEnabled returns true if the site is configured to use short IDs in URLs.
func (qs *QuestionCommon) isShortLinkEnabled(ctx context.Context) bool {
	if qs.siteInfoService == nil {
		return false
	}
	seo, err := qs.siteInfoService.GetSiteSeo(ctx)
	if err != nil {
		return false
	}
	return seo.IsShortLink()
}
