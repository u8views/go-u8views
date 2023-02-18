// Code generated by qtc from "footer.qtpl". DO NOT EDIT.
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

func streamfooter(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<footer class="footer">
    <div class="footer__wrapper wrapper-u8">
        <div class="footer__group-u8">
            <div class="footer__info">
                <a href="/" class="footer__title">
                    <img src="/assets/images/logo.svg" alt="logo">
                </a>
                <p class="footer__subtitle">Views counter and views statistics for your GitHub profile</p>
                <div class="footer__map">
                    <img class="footer__map-Ukraine" src="/assets/images/map-of-ukraine.png" alt="Map of Ukraine">
                </div>
            </div>

            <div class="footer__middle-section">
                <a href="/stats" class="footer__statistics-link">Registration statistics on u8views</a>
                <a href="https://github.com/u8views/go-u8views" class="footer__github-link">
                    View the project on GitHub
                    <img src="/assets/images/arrow.svg" alt="arrow">
                </a>
                <iframe class="footer__stars"
                        src="https://ghbtns.com/github-btn.html?user=twbs&repo=bootstrap&type=star&count=true&size=large"
                        frameborder="0" scrolling="0" width="170" height="34" title="GitHub"></iframe>
            </div>

            <div class="footer__support">
                <a href="https://savelife.in.ua/en/" class="footer__link">
                    <figure class="footer__figure">
                        <img width="60" height="24" src="/assets/images/come-back-alive-ukraine.png" alt="support">
                        <figcaption class="footer__caption">Support</figcaption>
                        <img src="/assets/images/arrow.svg" alt="arrow">
                    </figure>
                </a>
                <a href="https://bank.gov.ua/en/about/support-the-armed-forces/" class="footer__link">
                    <figure class="footer__figure">
                        <img width="60" height="24" src="/assets/images/national-bank-of-ukraine.png" alt="support">
                        <figcaption class="footer__caption">Support</figcaption>
                        <img src="/assets/images/arrow.svg" alt="arrow">
                    </figure>
                </a>
                <a href="https://war.ukraine.ua/" class="footer__link">
                    <figure class="footer__figure">
                        <figcaption class="footer__caption">war.ukraine.ua</figcaption>
                        <img src="/assets/images/arrow.svg" alt="arrow">
                    </figure>
                </a>
            </div>
        </div>
        <div class="footer__copyrights">
            <span>© 2023 Yaroslav Podorvanov</span>
            <img class="footer__flag-UA" src="/assets/images/flag-of-ukraine.svg" alt="Flag of Ukraine">
        </div>
    </div>
</footer>
`)
}

func writefooter(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamfooter(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func footer() string {
	qb422016 := qt422016.AcquireByteBuffer()
	writefooter(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
