// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssturlshortener

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/sst-url-shortener-go/internal/apijson"
	"github.com/stainless-sdks/sst-url-shortener-go/internal/apiquery"
	"github.com/stainless-sdks/sst-url-shortener-go/internal/param"
	"github.com/stainless-sdks/sst-url-shortener-go/internal/requestconfig"
	"github.com/stainless-sdks/sst-url-shortener-go/option"
)

// URLService contains methods and other services that help with interacting with
// the sst-url-shortener API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLService] method instead.
type URLService struct {
	Options []option.RequestOption
}

// NewURLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewURLService(opts ...option.RequestOption) (r *URLService) {
	r = &URLService{}
	r.Options = opts
	return
}

// Create a new short url
func (r *URLService) New(ctx context.Context, body URLNewParams, opts ...option.RequestOption) (res *URLNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "urls/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a short url by original url
func (r *URLService) DeleteByOriginalURL(ctx context.Context, body URLDeleteByOriginalURLParams, opts ...option.RequestOption) (res *URLDeleteByOriginalURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "urls/delete-by-original-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Delete a short url by short id
func (r *URLService) DeleteByShortID(ctx context.Context, body URLDeleteByShortIDParams, opts ...option.RequestOption) (res *URLDeleteByShortIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "urls/delete-by-short-id"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Get the short url from the original url
func (r *URLService) FromOriginalURL(ctx context.Context, query URLFromOriginalURLParams, opts ...option.RequestOption) (res *URLFromOriginalURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "urls/from-original-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the short url from the short id
func (r *URLService) FromShortID(ctx context.Context, query URLFromShortIDParams, opts ...option.RequestOption) (res *URLFromShortIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "urls/from-short-id"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get approximate count of short urls in the DB. Updated every 6 hours.
func (r *URLService) QuickCount(ctx context.Context, opts ...option.RequestOption) (res *URLQuickCountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "urls/quick-count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Paginated search of short urls
func (r *URLService) Search(ctx context.Context, query URLSearchParams, opts ...option.RequestOption) (res *URLSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "urls/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Scan through the entire table to get real-time count of items
func (r *URLService) SlowCount(ctx context.Context, opts ...option.RequestOption) (res *URLSlowCountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "urls/slow-count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ShortURL struct {
	CreatedAt   time.Time    `json:"createdAt,required" format:"date-time"`
	OriginalURL string       `json:"originalUrl,required" format:"uri"`
	ShortID     string       `json:"shortId,required"`
	ShortURL    string       `json:"shortUrl,required" format:"uri"`
	ExpiredAt   time.Time    `json:"expiredAt" format:"date-time"`
	JSON        shortURLJSON `json:"-"`
}

// shortURLJSON contains the JSON metadata for the struct [ShortURL]
type shortURLJSON struct {
	CreatedAt   apijson.Field
	OriginalURL apijson.Field
	ShortID     apijson.Field
	ShortURL    apijson.Field
	ExpiredAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShortURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shortURLJSON) RawJSON() string {
	return r.raw
}

type ShortURLCountResult struct {
	Count float64                 `json:"count,required"`
	JSON  shortURLCountResultJSON `json:"-"`
}

// shortURLCountResultJSON contains the JSON metadata for the struct
// [ShortURLCountResult]
type shortURLCountResultJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShortURLCountResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shortURLCountResultJSON) RawJSON() string {
	return r.raw
}

type ShortURLSearchResult struct {
	Cursor string                    `json:"cursor,required,nullable"`
	URLs   []ShortURLSearchResultURL `json:"urls,required"`
	JSON   shortURLSearchResultJSON  `json:"-"`
}

// shortURLSearchResultJSON contains the JSON metadata for the struct
// [ShortURLSearchResult]
type shortURLSearchResultJSON struct {
	Cursor      apijson.Field
	URLs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShortURLSearchResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shortURLSearchResultJSON) RawJSON() string {
	return r.raw
}

type ShortURLSearchResultURL struct {
	CreatedAt   time.Time                   `json:"createdAt,required" format:"date-time"`
	OriginalURL string                      `json:"originalUrl,required" format:"uri"`
	ShortID     string                      `json:"shortId,required"`
	ShortURL    string                      `json:"shortUrl,required" format:"uri"`
	ExpiredAt   time.Time                   `json:"expiredAt" format:"date-time"`
	JSON        shortURLSearchResultURLJSON `json:"-"`
}

// shortURLSearchResultURLJSON contains the JSON metadata for the struct
// [ShortURLSearchResultURL]
type shortURLSearchResultURLJSON struct {
	CreatedAt   apijson.Field
	OriginalURL apijson.Field
	ShortID     apijson.Field
	ShortURL    apijson.Field
	ExpiredAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShortURLSearchResultURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shortURLSearchResultURLJSON) RawJSON() string {
	return r.raw
}

type URLNewResponse struct {
	Result ShortURL           `json:"result,required"`
	JSON   urlNewResponseJSON `json:"-"`
}

// urlNewResponseJSON contains the JSON metadata for the struct [URLNewResponse]
type urlNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNewResponseJSON) RawJSON() string {
	return r.raw
}

type URLDeleteByOriginalURLResponse = interface{}

type URLDeleteByShortIDResponse = interface{}

type URLFromOriginalURLResponse struct {
	Result ShortURL                       `json:"result,required"`
	JSON   urlFromOriginalURLResponseJSON `json:"-"`
}

// urlFromOriginalURLResponseJSON contains the JSON metadata for the struct
// [URLFromOriginalURLResponse]
type urlFromOriginalURLResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLFromOriginalURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlFromOriginalURLResponseJSON) RawJSON() string {
	return r.raw
}

type URLFromShortIDResponse struct {
	Result ShortURL                   `json:"result,required"`
	JSON   urlFromShortIDResponseJSON `json:"-"`
}

// urlFromShortIDResponseJSON contains the JSON metadata for the struct
// [URLFromShortIDResponse]
type urlFromShortIDResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLFromShortIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlFromShortIDResponseJSON) RawJSON() string {
	return r.raw
}

type URLQuickCountResponse struct {
	Result ShortURLCountResult       `json:"result,required"`
	JSON   urlQuickCountResponseJSON `json:"-"`
}

// urlQuickCountResponseJSON contains the JSON metadata for the struct
// [URLQuickCountResponse]
type urlQuickCountResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLQuickCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlQuickCountResponseJSON) RawJSON() string {
	return r.raw
}

type URLSearchResponse struct {
	Result ShortURLSearchResult  `json:"result,required"`
	JSON   urlSearchResponseJSON `json:"-"`
}

// urlSearchResponseJSON contains the JSON metadata for the struct
// [URLSearchResponse]
type urlSearchResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlSearchResponseJSON) RawJSON() string {
	return r.raw
}

type URLSlowCountResponse struct {
	Result ShortURLCountResult      `json:"result,required"`
	JSON   urlSlowCountResponseJSON `json:"-"`
}

// urlSlowCountResponseJSON contains the JSON metadata for the struct
// [URLSlowCountResponse]
type urlSlowCountResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLSlowCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlSlowCountResponseJSON) RawJSON() string {
	return r.raw
}

type URLNewParams struct {
	OriginalURL param.Field[string]    `json:"originalUrl,required" format:"uri"`
	ExpiredAt   param.Field[time.Time] `json:"expiredAt" format:"date-time"`
}

func (r URLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type URLDeleteByOriginalURLParams struct {
	OriginalURL param.Field[string] `query:"originalUrl,required" format:"uri"`
}

// URLQuery serializes [URLDeleteByOriginalURLParams]'s query parameters as
// `url.Values`.
func (r URLDeleteByOriginalURLParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type URLDeleteByShortIDParams struct {
	ShortID param.Field[string] `query:"shortId,required"`
}

// URLQuery serializes [URLDeleteByShortIDParams]'s query parameters as
// `url.Values`.
func (r URLDeleteByShortIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type URLFromOriginalURLParams struct {
	OriginalURL param.Field[string] `query:"originalUrl,required" format:"uri"`
}

// URLQuery serializes [URLFromOriginalURLParams]'s query parameters as
// `url.Values`.
func (r URLFromOriginalURLParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type URLFromShortIDParams struct {
	ShortID param.Field[string] `query:"shortId,required"`
}

// URLQuery serializes [URLFromShortIDParams]'s query parameters as `url.Values`.
func (r URLFromShortIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type URLSearchParams struct {
	Cursor                param.Field[string]    `query:"cursor"`
	ExpiredAtLte          param.Field[time.Time] `query:"expiredAtLTE" format:"date-time"`
	Limit                 param.Field[float64]   `query:"limit"`
	OriginalURLBeginsWith param.Field[string]    `query:"originalUrlBeginsWith"`
}

// URLQuery serializes [URLSearchParams]'s query parameters as `url.Values`.
func (r URLSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
