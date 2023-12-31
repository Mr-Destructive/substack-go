package substack

import (
	"net/http"
)

type Api struct {
	session        *http.Client
	publicationURL string
	cookies        []*http.Cookie
}

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Handle         string `json:"handle"`
	PreviousName   string `json:"previous_name"`
	PhotoURL       string `json:"photo_url"`
	Bio            string `json:"bio"`
	ProfileSetUpAt string `json:"profile_set_up_at"`
}

type SubscribeContent struct {
	Type    string `json:"type"`
	Content []struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}
}

type Publication struct {
	Subdomain           string `json:"subdomain"`
	Name                string `json:"name"`
	CustomDomain        string `json:"custom_domain"`
	LogoURL             string `json:"logo_url"`
	LogoURLWide         string `json:"logo_url_wide"`
	CoverPhotoURL       string `json:"cover_photo_url"`
	Copyright           string `json:"copyright"`
	EmailFromName       string `json:"email_from_name"`
	SubscribeContent    string `json:"subscribe_content"`
	SubscribeFooter     string `json:"subscribe_footer"`
	WelcomeEmailContent string `json:"welcome_email_content"`
}

type Post struct {
	ID                      int            `json:"id"`
	UUID                    string         `json:"uuid"`
	EditorV2                bool           `json:"editor_v2"`
	PublicationID           int            `json:"publication_id"`
	Type                    string         `json:"type"`
	PostDate                string         `json:"post_date"`
	DraftCreatedAt          string         `json:"draft_created_at"`
	EmailSentAt             string         `json:"email_sent_at"`
	IsPublished             bool           `json:"is_published"`
	Title                   string         `json:"title"`
	DraftTitle              string         `json:"draft_title"`
	DraftUpdatedAt          string         `json:"draft_updated_at"`
	Audience                string         `json:"audience"`
	Slug                    string         `json:"slug"`
	ShouldSendEmail         string         `json:"should_send_email"`
	WriteCommentPermissions string         `json:"write_comment_permissions"`
	DefaultCommentSort      string         `json:"default_comment_sort"`
	SectionID               string         `json:"section_id"`
	ShouldSendFreePreview   bool           `json:"should_send_free_preview"`
	VideoUploadID           string         `json:"video_upload_id"`
	SectionSlug             string         `json:"section_slug"`
	SectionName             string         `json:"section_name"`
	DraftSectionName        string         `json:"draft_section_name"`
	IsSectionPinned         bool           `json:"is_section_pinned"`
	Reactions               map[string]int `json:"reactions"`
	Reaction                string         `json:"reaction"`
	TopExclusions           []interface{}  `json:"top_exclusions"`
	Pins                    []interface{}  `json:"pins"`
	DraftBylines            []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Handle         string `json:"handle"`
		PreviousName   string `json:"previous_name"`
		PhotoURL       string `json:"photo_url"`
		Bio            string `json:"bio"`
		ProfileSetUpAt string `json:"profile_set_up_at"`
	} `json:"draftBylines"`
	PublishedBylines []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Handle         string `json:"handle"`
		PreviousName   string `json:"previous_name"`
		PhotoURL       string `json:"photo_url"`
		Bio            string `json:"bio"`
		ProfileSetUpAt string `json:"profile_set_up_at"`
	} `json:"publishedBylines"`
	ReactionCount     int `json:"reaction_count"`
	CommentCount      int `json:"comment_count"`
	ChildCommentCount int `json:"child_comment_count"`
	Bylines           []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Handle         string `json:"handle"`
		PreviousName   string `json:"previous_name"`
		PhotoURL       string `json:"photo_url"`
		Bio            string `json:"bio"`
		ProfileSetUpAt string `json:"profile_set_up_at"`
	} `json:"bylines"`
	Stats struct {
		Views                        int `json:"views"`
		Opens                        int `json:"opens"`
		Opened                       int `json:"opened"`
		OpenRate                     int `json:"open_rate"`
		Clicked                      int `json:"clicked"`
		Clicks                       int `json:"clicks"`
		Sent                         int `json:"sent"`
		Delivered                    int `json:"delivered"`
		Downloads                    int `json:"downloads"`
		DownloadsDay7                int `json:"downloads_day7"`
		DownloadsDay30               int `json:"downloads_day30"`
		DownloadsDay90               int `json:"downloads_day90"`
		PodcastPreviewDownloads      int `json:"podcast_preview_downloads"`
		PodcastPreviewDownloadsDay30 int `json:"podcast_preview_downloads_day30"`
		VideoViewers                 int `json:"video_viewers"`
		VideoViews                   int `json:"video_views"`
		VideoMinutesWatched          int `json:"video_minutes_watched"`
		SignupsWithin1Day            int `json:"signups_within_1_day"`
		DisablesWithin1Day           int `json:"disables_within_1_day"`
		SubscriptionsWithin1Day      int `json:"subscriptions_within_1_day"`
		UnsubscribesWithin1Day       int `json:"unsubscribes_within_1_day"`
		Signups                      int `json:"signups"`
		Subscribes                   int `json:"subscribes"`
		Shares                       int `json:"shares"`
		EstimatedValue               int `json:"estimated_value"`
		ClickThroughRate             int `json:"click_through_rate"`
		EngagementRate               int `json:"engagement_rate"`
	} `json:"stats"`
}

type Posts struct {
	Posts  []Post `json:"posts"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Total  int    `json:"total"`
}

type Category struct {
	ID                     int    `json:"id"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
	Name                   string `json:"name"`
	CanonicalName          string `json:"canonical_name"`
	Active                 bool   `json:"active"`
	Rank                   int    `json:"rank"`
	ParentTagID            int    `json:"parent_tag_id"`
	Slug                   string `json:"slug"`
	Emoji                  string `json:"emoji"`
	LeaderboardDescription string `json:"leaderboard_description"`
}

type Subscription struct {
	ID                     int    `json:"id"`
	UserID                 int    `json:"user_id"`
	PublicationID          int    `json:"publication_id"`
	Expiry                 string `json:"expiry"`
	EmailDisabled          bool   `json:"email_disabled"`
	DigestEnabled          bool   `json:"digest_enabled"`
	MembershipState        string `json:"membership_state"`
	Type                   string `json:"type"`
	GiftUserID             int    `json:"gift_user_id"`
	CreatedAt              string `json:"created_at"`
	GiftedAt               string `json:"gifted_at"`
	Paused                 string `json:"paused"`
	IsGroupParent          bool   `json:"is_group_parent"`
	Visibility             string `json:"visibility"`
	IsFounding             bool   `json:"is_founding"`
	IsFavorite             bool   `json:"is_favorite"`
	PodcastRSSToken        string `json:"podcast_rss_token"`
	EmailSettings          string `json:"email_settings"`
	SectionPodcastsEnabled string `json:"section_podcasts_enabled"`
}

type SubscriptionsList struct {
	Subscriptions []Subscription `json:"subscriptions"`
}
