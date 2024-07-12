// Code generated by qtc from "week_badge.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package v2

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamWeekBadge(qw422016 *qt422016.Writer, stats int64) {
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
      <text x="57.5" y="16">Views per week</text>
      <text x="138" y="16">`)
	qw422016.N().DL(stats)
	qw422016.N().S(`</text>
    </g>
  </g>
</svg>
`)
}

func WriteWeekBadge(qq422016 qtio422016.Writer, stats int64) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamWeekBadge(qw422016, stats)
	qt422016.ReleaseWriter(qw422016)
}

func WeekBadge(stats int64) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteWeekBadge(qb422016, stats)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
