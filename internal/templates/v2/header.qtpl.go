// Code generated by qtc from "header.qtpl". DO NOT EDIT.
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

func streamheader(qw422016 *qt422016.Writer, currentPageProfile ProfileView, sessionProfile ProfileView) {
	qw422016.N().S(`
    `)
	if sessionProfile.ID == 0 {
		qw422016.N().S(`
        `)
		streamunauthorizedHeader(qw422016, currentPageProfile)
		qw422016.N().S(`
    `)
	} else {
		qw422016.N().S(`
        `)
		streamauthorizedHeader(qw422016, sessionProfile)
		qw422016.N().S(`
    `)
	}
	qw422016.N().S(`
`)
}

func writeheader(qq422016 qtio422016.Writer, currentPageProfile ProfileView, sessionProfile ProfileView) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamheader(qw422016, currentPageProfile, sessionProfile)
	qt422016.ReleaseWriter(qw422016)
}

func header(currentPageProfile ProfileView, sessionProfile ProfileView) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeheader(qb422016, currentPageProfile, sessionProfile)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func streamunauthorizedHeader(qw422016 *qt422016.Writer, currentPageProfile ProfileView) {
	qw422016.N().S(`
<header class="header">
    <div class="header__wrapper wrapper-u8">
        <a href="/" class="header__logo">
            <img class="header__logo-img" src="/assets/images/logo.svg" alt="logo">
        </a>
        <div class="header__stars">
            <iframe src="https://ghbtns.com/github-btn.html?user=u8views&repo=go-u8views&type=star&count=true&size=large"
                    frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
        </div>
        <button class="button__log-in">
            <a href="/login/github?referrer=`)
	qw422016.N().DL(currentPageProfile.ID)
	qw422016.N().S(`" class="button__link">
                <img
                        src="/assets/images/github.svg"
                        width="24"
                        height="24"
                        alt="github"
                        class="button__log-in-img"
                />
                <span class="button__text">Log in with GitHub</span>
            </a>
        </button>
    </div>
</header>
`)
}

func writeunauthorizedHeader(qq422016 qtio422016.Writer, currentPageProfile ProfileView) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamunauthorizedHeader(qw422016, currentPageProfile)
	qt422016.ReleaseWriter(qw422016)
}

func unauthorizedHeader(currentPageProfile ProfileView) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeunauthorizedHeader(qb422016, currentPageProfile)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func streamauthorizedHeader(qw422016 *qt422016.Writer, sessionProfile ProfileView) {
	qw422016.N().S(`
<header class="header">
    <div class="header__wrapper auth wrapper-u8">
        <a href="/" class="header__logo">
            <img class="header__logo-img" src="/assets/images/logo.svg" alt="logo">
        </a>
        <div class="header__stars">
            <iframe src="https://ghbtns.com/github-btn.html?user=u8views&repo=go-u8views&type=star&count=true&size=large"
                    frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
        </div>
        <div class="header__profile">
            <button class="header__profile-button">
                <img src="https://avatars.githubusercontent.com/u/`)
	qw422016.E().S(sessionProfile.SocialProviderUserID)
	qw422016.N().S(`?v=4&s=48"
                     width="48" height="48" alt="`)
	qw422016.E().S(sessionProfile.GetName())
	qw422016.N().S(` profile photo">
            </button>
            <div class="header__modal modal">
                <div class="modal__profile">
                    <img src="https://avatars.githubusercontent.com/u/`)
	qw422016.E().S(sessionProfile.SocialProviderUserID)
	qw422016.N().S(`?v=4&s=48"
                         class="modal__user-photo" width="48" height="48"
                         alt="`)
	qw422016.E().S(sessionProfile.GetName())
	qw422016.N().S(` profile photo">
                    <div class="modal__profile-info">
                        <div class="modal__profile-name">`)
	qw422016.E().S(sessionProfile.GetName())
	qw422016.N().S(`</div>
                        <a href="https://u8views.com/github/`)
	qw422016.E().S(sessionProfile.Username)
	qw422016.N().S(`" class="modal__link">
                            <img src="/assets/images/u8-icon.svg" width="24" height="24" alt="link">
                            <span class="modal__profile-u8views">u8views.com/github/`)
	qw422016.E().S(sessionProfile.Username)
	qw422016.N().S(`</span>
                        </a>
                        <a href="https://github.com/`)
	qw422016.E().S(sessionProfile.Username)
	qw422016.N().S(`" class="modal__link">
                            <img src="/assets/images/github.svg" width="24" height="24" alt="link">
                            <span class="modal__profile-github">github.com/`)
	qw422016.E().S(sessionProfile.Username)
	qw422016.N().S(`</span>
                        </a>
                    </div>
                </div>
                <div class="modal__log-out">
                    <a href="/logout" class="modal__button">
                        <img src="/assets/images/log-out.svg" alt="log-out">
                        <span>Log out</span>
                    </a>
                </div>
            </div>
        </div>
    </div>
</header>

<script>
    document.querySelector(".header__profile-button").addEventListener("click", () => {
        document.querySelector(".header__modal").classList.toggle("active");
    })
</script>
`)
}

func writeauthorizedHeader(qq422016 qtio422016.Writer, sessionProfile ProfileView) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamauthorizedHeader(qw422016, sessionProfile)
	qt422016.ReleaseWriter(qw422016)
}

func authorizedHeader(sessionProfile ProfileView) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeauthorizedHeader(qb422016, sessionProfile)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
