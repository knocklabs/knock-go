// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/knocklabs/knock-go/internal/apijson"
	"github.com/knocklabs/knock-go/internal/requestconfig"
	"github.com/knocklabs/knock-go/option"
)

type EntriesCursorPageInfo struct {
	After string                    `json:"after"`
	JSON  entriesCursorPageInfoJSON `json:"-"`
}

// entriesCursorPageInfoJSON contains the JSON metadata for the struct
// [EntriesCursorPageInfo]
type entriesCursorPageInfoJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntriesCursorPageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entriesCursorPageInfoJSON) RawJSON() string {
	return r.raw
}

type EntriesCursor[T any] struct {
	Entries  []T                   `json:"entries"`
	PageInfo EntriesCursorPageInfo `json:"page_info"`
	JSON     entriesCursorJSON     `json:"-"`
	cfg      *requestconfig.RequestConfig
	res      *http.Response
}

// entriesCursorJSON contains the JSON metadata for the struct [EntriesCursor[T]]
type entriesCursorJSON struct {
	Entries     apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntriesCursor[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entriesCursorJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EntriesCursor[T]) GetNextPage() (res *EntriesCursor[T], err error) {
	next := r.PageInfo.After
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("after", next))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *EntriesCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EntriesCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EntriesCursorAutoPager[T any] struct {
	page *EntriesCursor[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewEntriesCursorAutoPager[T any](page *EntriesCursor[T], err error) *EntriesCursorAutoPager[T] {
	return &EntriesCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EntriesCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Entries) == 0 {
		return false
	}
	if r.idx >= len(r.page.Entries) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Entries) == 0 {
			return false
		}
	}
	r.cur = r.page.Entries[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *EntriesCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *EntriesCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *EntriesCursorAutoPager[T]) Index() int {
	return r.run
}

type ItemsCursorPageInfo struct {
	After string                  `json:"after"`
	JSON  itemsCursorPageInfoJSON `json:"-"`
}

// itemsCursorPageInfoJSON contains the JSON metadata for the struct
// [ItemsCursorPageInfo]
type itemsCursorPageInfoJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ItemsCursorPageInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemsCursorPageInfoJSON) RawJSON() string {
	return r.raw
}

type ItemsCursor[T any] struct {
	Items    []T                 `json:"items"`
	PageInfo ItemsCursorPageInfo `json:"page_info"`
	JSON     itemsCursorJSON     `json:"-"`
	cfg      *requestconfig.RequestConfig
	res      *http.Response
}

// itemsCursorJSON contains the JSON metadata for the struct [ItemsCursor[T]]
type itemsCursorJSON struct {
	Items       apijson.Field
	PageInfo    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ItemsCursor[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemsCursorJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ItemsCursor[T]) GetNextPage() (res *ItemsCursor[T], err error) {
	next := r.PageInfo.After
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("after", next))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ItemsCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ItemsCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ItemsCursorAutoPager[T any] struct {
	page *ItemsCursor[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewItemsCursorAutoPager[T any](page *ItemsCursor[T], err error) *ItemsCursorAutoPager[T] {
	return &ItemsCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ItemsCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ItemsCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *ItemsCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *ItemsCursorAutoPager[T]) Index() int {
	return r.run
}

type SlackChannelsCursor[T any] struct {
	NextCursor    string                  `json:"next_cursor"`
	SlackChannels []T                     `json:"slack_channels"`
	JSON          slackChannelsCursorJSON `json:"-"`
	cfg           *requestconfig.RequestConfig
	res           *http.Response
}

// slackChannelsCursorJSON contains the JSON metadata for the struct
// [SlackChannelsCursor[T]]
type slackChannelsCursorJSON struct {
	NextCursor    apijson.Field
	SlackChannels apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SlackChannelsCursor[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slackChannelsCursorJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SlackChannelsCursor[T]) GetNextPage() (res *SlackChannelsCursor[T], err error) {
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("query_options.cursor", next))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SlackChannelsCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SlackChannelsCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SlackChannelsCursorAutoPager[T any] struct {
	page *SlackChannelsCursor[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewSlackChannelsCursorAutoPager[T any](page *SlackChannelsCursor[T], err error) *SlackChannelsCursorAutoPager[T] {
	return &SlackChannelsCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SlackChannelsCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.SlackChannels) == 0 {
		return false
	}
	if r.idx >= len(r.page.SlackChannels) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.SlackChannels) == 0 {
			return false
		}
	}
	r.cur = r.page.SlackChannels[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SlackChannelsCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *SlackChannelsCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *SlackChannelsCursorAutoPager[T]) Index() int {
	return r.run
}

type MsTeamsPagination[T any] struct {
	SkipToken    string                `json:"skip_token"`
	MsTeamsTeams []T                   `json:"ms_teams_teams"`
	JSON         msTeamsPaginationJSON `json:"-"`
	cfg          *requestconfig.RequestConfig
	res          *http.Response
}

// msTeamsPaginationJSON contains the JSON metadata for the struct
// [MsTeamsPagination[T]]
type msTeamsPaginationJSON struct {
	SkipToken    apijson.Field
	MsTeamsTeams apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MsTeamsPagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r msTeamsPaginationJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *MsTeamsPagination[T]) GetNextPage() (res *MsTeamsPagination[T], err error) {
	next := r.SkipToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("query_options.$skiptoken", next))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *MsTeamsPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &MsTeamsPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type MsTeamsPaginationAutoPager[T any] struct {
	page *MsTeamsPagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewMsTeamsPaginationAutoPager[T any](page *MsTeamsPagination[T], err error) *MsTeamsPaginationAutoPager[T] {
	return &MsTeamsPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *MsTeamsPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.MsTeamsTeams) == 0 {
		return false
	}
	if r.idx >= len(r.page.MsTeamsTeams) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.MsTeamsTeams) == 0 {
			return false
		}
	}
	r.cur = r.page.MsTeamsTeams[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *MsTeamsPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *MsTeamsPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *MsTeamsPaginationAutoPager[T]) Index() int {
	return r.run
}
