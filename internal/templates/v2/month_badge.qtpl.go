// Code generated by qtc from "month_badge.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line internal/templates/v2/month_badge.qtpl:1
package v2

//line internal/templates/v2/month_badge.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line internal/templates/v2/month_badge.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line internal/templates/v2/month_badge.qtpl:1
func StreamMonthBadge(qw422016 *qt422016.Writer, stats ProfileViewsStats) {
//line internal/templates/v2/month_badge.qtpl:1
	qw422016.N().S(`
<svg width="568" height="24" viewBox="0 0 568 24" fill="none" xmlns="http://www.w3.org/2000/svg">
 <g>
     <rect width="113" height="24" rx="4" fill="#24292F"></rect>
         <rect x="112" width="54" height="24" rx="4" fill="#6D96FF"></rect>
         <rect x="112" width="1" height="24" fill="#E4EAF1"></rect>
     <g
       fill="#fff"
       text-anchor="middle"
       font-family="DejaVu Sans,Verdana,Geneva,sans-serif"
       font-size="14"
     >
       <text x="57.5" y="16">Views per month</text>
       <text x="138" y="16">`)
//line internal/templates/v2/month_badge.qtpl:14
	qw422016.N().DL(stats.MonthCount)
//line internal/templates/v2/month_badge.qtpl:14
	qw422016.N().S(`</text>
     </g>
   </g>
</svg>
`)
//line internal/templates/v2/month_badge.qtpl:18
}

//line internal/templates/v2/month_badge.qtpl:18
func WriteMonthBadge(qq422016 qtio422016.Writer, stats ProfileViewsStats) {
//line internal/templates/v2/month_badge.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/templates/v2/month_badge.qtpl:18
	StreamMonthBadge(qw422016, stats)
//line internal/templates/v2/month_badge.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line internal/templates/v2/month_badge.qtpl:18
}

//line internal/templates/v2/month_badge.qtpl:18
func MonthBadge(stats ProfileViewsStats) string {
//line internal/templates/v2/month_badge.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/templates/v2/month_badge.qtpl:18
	WriteMonthBadge(qb422016, stats)
//line internal/templates/v2/month_badge.qtpl:18
	qs422016 := string(qb422016.B)
//line internal/templates/v2/month_badge.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/templates/v2/month_badge.qtpl:18
	return qs422016
//line internal/templates/v2/month_badge.qtpl:18
}
