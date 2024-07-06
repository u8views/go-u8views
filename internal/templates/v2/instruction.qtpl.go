// Code generated by qtc from "instruction.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line internal/templates/v2/instruction.qtpl:1
package v2

//line internal/templates/v2/instruction.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line internal/templates/v2/instruction.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line internal/templates/v2/instruction.qtpl:1
func streaminstruction(qw422016 *qt422016.Writer, currentPageProfile ProfileView, sessionProfile ProfileView, exampleProfile ProfileView, done bool) {
//line internal/templates/v2/instruction.qtpl:1
	qw422016.N().S(`
<section class="instruction">
    <div class="instruction__header">
        <div class="instruction__step">
            <img
                    src="/assets/images/arrow-right.svg"
                    width="24"
                    height="24"
                    alt="arrow"
            />
            INSTRUCTION
        </div>
        <div class="instruction__header-group">
            <div class="instruction__title">
                Follow these steps to get started:
            </div>
            <button class="instruction__show-button js-instruction-toggle-button">
                <img
                        src="/assets/images/arrow-down.svg"
                        width="15"
                        height="10"
                        alt="arrow"
                        class="instruction__show-img js-instruction-visibility-state"
                />
            </button>
        </div>
    </div>

    <div class="instruction__steps js-instruction">
        <div class="instruction__step-1 `)
//line internal/templates/v2/instruction.qtpl:30
	if sessionProfile.ID > 0 {
//line internal/templates/v2/instruction.qtpl:30
		qw422016.N().S(`active`)
//line internal/templates/v2/instruction.qtpl:30
	}
//line internal/templates/v2/instruction.qtpl:30
	qw422016.N().S(` step">
            <div class="step__progress">
                <div class="step__done">
                    <img
                            src="/assets/images/done.svg"
                            class="step__done-img"
                            alt="done"
                    />
                </div>
                <div class="step__progress-bar"></div>
            </div>
            <div class="step__view">
                <h2 class="step__title">Sign up with GitHub on u8views</h2>
                <div class="step__main">
                    <div class="wrapper-u8 step-wrapper">
                        <img
                                class="step__browser"
                                src="/assets/images/browser.png"
                                alt="browser"
                        />
                        <header class="header step-block">
                            <div class="header__wrapper wrapper-u8">
                                <a href="/design" class="header__logo">
                                    <img
                                            class="header__logo-img"
                                            src="/assets/images/logo.svg"
                                            alt=""
                                    /></a>
                                <div class="header__stars">
                                    <iframe
                                            src="https://ghbtns.com/github-btn.html?user=u8views&repo=go-u8views&type=star&count=true"
                                            frameborder="0"
                                            scrolling="0"
                                            width="110"
                                            height="20"
                                            title="GitHub"
                                    ></iframe>
                                </div>
                                <button class="button__log-in step-block">
                                    <a href="/login/github?referrer=`)
//line internal/templates/v2/instruction.qtpl:69
	qw422016.N().DL(currentPageProfile.ID)
//line internal/templates/v2/instruction.qtpl:69
	qw422016.N().S(`"
                                       class="button__log-in-link">
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
                        <section class="hero step-block">
                            <div class="hero__main">
                                <h1 class="hero__title">
                                    Track your GitHub profile views
                                </h1>
                                <p class="hero__subtitle">
                                    Receive, view and analyze your profile views and profile performance statistics
                                </p>
                                <a href="/login/github?referrer=`)
//line internal/templates/v2/instruction.qtpl:91
	qw422016.N().DL(currentPageProfile.ID)
//line internal/templates/v2/instruction.qtpl:91
	qw422016.N().S(`" class="hero__button">
                                    <button class="button__log-in hero-button step-block">
                                        <img
                                                src="/assets/images/github.svg"
                                                width="12"
                                                height="12"
                                                alt="github"
                                                class="button__log-in-img"
                                        />
                                        <span class="button__text">Log in with GitHub</span>
                                    </button>
                                </a>
                            </div>
                        </section>

                        <div class="example__badges step-block">
                            <h3 class="example__title">That badge looks like this:</h3>
                            <div class="example__badges-group">
                                <div class="history__badge">
                                    <span class="history__badge-title">Views per day</span>
                                    <span class="history__badge-count">17</span>
                                </div>
                                <div class="history__badge">
                                    <span class="history__badge-title">Views per week</span>
                                    <span class="history__badge-count">328</span>
                                </div>
                                <div class="history__badge">
                                    <span class="history__badge-title">Views per month</span>
                                    <span class="history__badge-count">986</span>
                                </div>
                                <div class="history__badge">
                                    <span class="history__badge-title">Total views</span>
                                    <span class="history__badge-count">48656</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="instruction__step-2 `)
//line internal/templates/v2/instruction.qtpl:132
	if done {
//line internal/templates/v2/instruction.qtpl:132
		qw422016.N().S(`active`)
//line internal/templates/v2/instruction.qtpl:132
	}
//line internal/templates/v2/instruction.qtpl:132
	qw422016.N().S(` step js-instruction-repository-exists">
            <div class="step__progress">
                <div class="step__done">
                    <img
                            src="/assets/images/done.svg"
                            class="step__done-img"
                            alt="done"
                    />
                </div>
                <div class="step__progress-bar"></div>
            </div>
            <div class="step__view">
                <h2 class="step__title step-link">
                    Create a public repository with the same name as your username
                </h2>
                <div class="step__subtitle">
                    <a href="https://github.com/new" class="step__link">Go to create repository
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path d="M7 17L17 7" stroke="#0057FF" stroke-width="2" stroke-linecap="round"
                                  stroke-linejoin="round"/>
                            <path d="M7 7H17V17" stroke="#0057FF" stroke-width="2" stroke-linecap="round"
                                  stroke-linejoin="round"/>
                        </svg>
                    </a>
                </div>
                <div class="step__main step-2">
                    <div class="step-2__item-group">
                        <div class="step-2__item">
                            <label for="repository-owner" class="step-2__label">Owner <span
                                    class="step-2__label-star">*</span></label>
                            <div class="step-2__select-block">
                                <img
                                        class="step-2__profile-photo"
                                        src="https://avatars.githubusercontent.com/u/`)
//line internal/templates/v2/instruction.qtpl:165
	qw422016.E().S(exampleProfile.SocialProviderUserID)
//line internal/templates/v2/instruction.qtpl:165
	qw422016.N().S(`?v=4&s=48"
                                        alt="photo"
                                />
                                <select
                                        name="repository-owner"
                                        id="repository-owner"
                                        class="step-2__select"
                                        type="text"
                                >
                                    <option value="`)
//line internal/templates/v2/instruction.qtpl:174
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:174
	qw422016.N().S(`">`)
//line internal/templates/v2/instruction.qtpl:174
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:174
	qw422016.N().S(`
                                    </option>
                                </select>
                            </div>
                        </div>
                        <span class="step-2__slash">/</span>
                        <div class="step-2__item">
                            <label for="repository_name" class="step-2__label">Repository name <span
                                    class="step-2__label-star">*</span></label>
                            <div class="step-2__input-block">
                                <input id="repository_name" class="step-2__input" type="text"
                                       value="`)
//line internal/templates/v2/instruction.qtpl:185
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:185
	qw422016.N().S(`"/>
                                <svg width="24" height="25" viewBox="0 0 24 25" fill="none"
                                     xmlns="http://www.w3.org/2000/svg">
                                    <path d="M20 6.5L9 17.5L4 12.5" stroke="#61B33D" stroke-width="2"
                                          stroke-linecap="round" stroke-linejoin="round"/>
                                </svg>
                            </div>
                        </div>
                    </div>
                    <div class="step-2__readme">
                        <img
                                src="https://github.githubassets.com/images/mona-whisper.gif"
                                class="step-2__readme-img"
                                alt="Whispering..."
                        />
                        <div class="step-2__readme-text">
                            <span class="step-2__readme-strong">`)
//line internal/templates/v2/instruction.qtpl:201
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:201
	qw422016.N().S(`/`)
//line internal/templates/v2/instruction.qtpl:201
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:201
	qw422016.N().S(`</span>
                            is a ✨<span class="step-2__readme-italic">special</span>✨ repository that you can use to add
                            a README.md to your GitHub profile. Make sure it's public and initialize it with a <span
                                class="step-2__readme-strong">README</span> to get started.
                        </div>
                    </div>
                    <div class="step-2__readme-file">
                        <input type="checkbox" id="checkbox" class="step-2__checkbox" checked/>
                        <span class="step-2__readme-file-text">Add a README file</span>
                    </div>

                    <div class="step-2__footer">
                        <button class="step-2__create-repo">
                            <a href="https://github.com/new" class="step-2__create-link">Create repository</a>
                        </button>
                    </div>
                </div>
            </div>
        </div>

        <div class="instruction__step-3 `)
//line internal/templates/v2/instruction.qtpl:221
	if done {
//line internal/templates/v2/instruction.qtpl:221
		qw422016.N().S(`active`)
//line internal/templates/v2/instruction.qtpl:221
	}
//line internal/templates/v2/instruction.qtpl:221
	qw422016.N().S(` step">
            <div class="step__progress">
                <div class="step__done">
                    <img
                            src="/assets/images/done.svg"
                            class="step__done-img"
                            alt="done"
                    />
                </div>
                <div class="step__progress-bar"></div>
            </div>
            <div class="step__view">
                <h2 class="step__title">Copy a badge code from your profile page</h2>
                <div class="step__main step-3">
                    <div class="step-3__item-group">
                        `)
//line internal/templates/v2/instruction.qtpl:236
	if sessionProfile.ID == 0 {
//line internal/templates/v2/instruction.qtpl:236
		qw422016.N().S(`
                        <div class="step-3__item">
                            <div class="step-3__title">Markdown:</div>
                            <img
                                    src="/assets/images/markdown-blur.jpg"
                                    class="step-3__img"
                                    alt="markdown link"
                            />
                        </div>
                        <div class="step-3__item step-3__item-log">
                            <div class="step-3__title">HTML link:</div>
                            <img
                                    src="/assets/images/html-link-blur.jpg"
                                    class="step-3__img"
                                    alt="markdown link"
                            />
                            <div class="step-3__button">
                                <button class="button__log-in black hero-button">
                                    <a href="/login/github?referrer=`)
//line internal/templates/v2/instruction.qtpl:254
		qw422016.N().DL(currentPageProfile.ID)
//line internal/templates/v2/instruction.qtpl:254
		qw422016.N().S(`" class="button__link">
                                        <img
                                                src="/assets/images/github-white.svg"
                                                width="24"
                                                height="24"
                                                alt="github"
                                                class="button__log-in-img"
                                        />
                                        <span class="button__text">Log in with GitHub</span>
                                    </a>
                                </button>

                            </div>
                        </div>
                        `)
//line internal/templates/v2/instruction.qtpl:268
	} else {
//line internal/templates/v2/instruction.qtpl:268
		qw422016.N().S(`
                        <div class="step-3__item">
                            <div class="step-3__title">Markdown:</div>
                            <div class="step-3__item-content">
                                <p class="step-3__item-text js-code-for-copy">
                                    [![`)
//line internal/templates/v2/instruction.qtpl:273
		qw422016.E().S(exampleProfile.GetName())
//line internal/templates/v2/instruction.qtpl:273
		qw422016.N().S(` profile views](https://u8views.com/api/v1/github/profiles/`)
//line internal/templates/v2/instruction.qtpl:273
		qw422016.E().S(exampleProfile.SocialProviderUserID)
//line internal/templates/v2/instruction.qtpl:273
		qw422016.N().S(`/views/day-week-month-total-count.svg)](https://u8views.com/github/`)
//line internal/templates/v2/instruction.qtpl:273
		qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:273
		qw422016.N().S(`)
                                </p>
                                <button class="step-3__copy-text js-copy-code-button">
                                    <img
                                            src="/assets/images/copy-black.svg"
                                            width="24"
                                            height="24"
                                            alt="copy"
                                            class="step-3__copy-img js-copy-code-check"
                                    />
                                    <img
                                            src="/assets/images/check-green.svg"
                                            width="24"
                                            height="24"
                                            alt="copy"
                                            class="step-3__copy-img step-3__copy-done js-copy-code-done"
                                    />
                                </button>
                            </div>
                        </div>
                        <div class="step-3__item">
                            <div class="step-3__title">HTML:</div>
                            <div class="step-3__item-content">
                                <p class="step-3__item-text js-code-for-copy">
                                    &lt;a href=&quot;https://u8views.com/github/`)
//line internal/templates/v2/instruction.qtpl:297
		qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:297
		qw422016.N().S(`&quot;&gt;&lt;img src=&quot;https://u8views.com/api/v1/github/profiles/`)
//line internal/templates/v2/instruction.qtpl:297
		qw422016.E().S(exampleProfile.SocialProviderUserID)
//line internal/templates/v2/instruction.qtpl:297
		qw422016.N().S(`/views/day-week-month-total-count.svg&quot;&gt;&lt;/a&gt;
                                </p>
                                <button class="step-3__copy-text js-copy-code-button">
                                    <img
                                            src="/assets/images/copy-black.svg"
                                            width="24"
                                            height="24"
                                            alt="copy"
                                            class="step-3__copy-img js-copy-code-check"
                                    />
                                    <img
                                            src="/assets/images/check-green.svg"
                                            width="24"
                                            height="24"
                                            alt="copy"
                                            class="step-3__copy-img step-3__copy-done js-copy-code-done"
                                    />
                                </button>
                            </div>
                        </div>
                        `)
//line internal/templates/v2/instruction.qtpl:317
	}
//line internal/templates/v2/instruction.qtpl:317
	qw422016.N().S(`
                    </div>
                </div>
            </div>
        </div>

        <div class="instruction__step-4 `)
//line internal/templates/v2/instruction.qtpl:323
	if done {
//line internal/templates/v2/instruction.qtpl:323
		qw422016.N().S(`active`)
//line internal/templates/v2/instruction.qtpl:323
	}
//line internal/templates/v2/instruction.qtpl:323
	qw422016.N().S(` step">
            <div class="step__progress">
                <div class="step__done">
                    <img
                            src="/assets/images/done.svg"
                            class="step__done-img"
                            alt="done"
                    />
                </div>
                <div class="step__progress-bar"></div>
            </div>
            <div class="step__view-group">
                <div class="step__view">
                    <h2 class="step__title step-link">
                        Add a badge to GitHub profile
                    </h2>
                    <a href="https://github.com/`)
//line internal/templates/v2/instruction.qtpl:339
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:339
	qw422016.N().S(`/`)
//line internal/templates/v2/instruction.qtpl:339
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:339
	qw422016.N().S(`/edit/main/README.md"
                       class="step__subtitle step__readme-file-link">Go to edit the README.md file</a>
                    <div class="step__main step-4">
                        <nav class="step-4__nav">
                            <ul class="step-4__nav-list">
                                <li class="step-4__nav-item">
                                    <a href="https://github.com/`)
//line internal/templates/v2/instruction.qtpl:345
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:345
	qw422016.N().S(`"
                                       class="step-4__nav-link">
                                        <img
                                                width="10"
                                                height="10"
                                                src="/assets/images/nav-book.svg"
                                                alt="book"
                                        />
                                        Overview
                                    </a>
                                </li>
                                <li class="step-4__nav-item">
                                    <a href="https://github.com/`)
//line internal/templates/v2/instruction.qtpl:357
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:357
	qw422016.N().S(`?tab=repositories"
                                       class="step-4__nav-link">
                                        <img
                                                width="10"
                                                height="10"
                                                src="/assets/images/nav-repo.svg"
                                                alt="Repositories"
                                        />
                                        Repositories
                                    </a>
                                </li>
                                <li class="step-4__nav-item">
                                    <a href="https://github.com/`)
//line internal/templates/v2/instruction.qtpl:369
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:369
	qw422016.N().S(`?tab=projects"
                                       class="step-4__nav-link">
                                        <img
                                                width="10"
                                                height="10"
                                                src="/assets/images/nav-project.svg"
                                                alt="Projects"
                                        />
                                        Projects
                                    </a>
                                </li>
                                <li class="step-4__nav-item">
                                    <a href="https://github.com/`)
//line internal/templates/v2/instruction.qtpl:381
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:381
	qw422016.N().S(`?tab=packages"
                                       class="step-4__nav-link">
                                        <img
                                                width="10"
                                                height="10"
                                                src="/assets/images/nav-packages.svg"
                                                alt="Packages"
                                        />
                                        Packages
                                    </a>
                                </li>
                                <li class="step-4__nav-item">
                                    <a href="https://github.com/`)
//line internal/templates/v2/instruction.qtpl:393
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:393
	qw422016.N().S(`?tab=stars"
                                       class="step-4__nav-link">
                                        <img
                                                width="10"
                                                height="10"
                                                src="/assets/images/nav-star.svg"
                                                alt="Stars"
                                        />
                                        Stars
                                    </a>
                                </li>
                            </ul>
                        </nav>
                        <div class="step-4__profile">
                            <div class="step-4__left-section">
                                <div class="step-4__profile-photo">
                                    <img
                                            src="https://avatars.githubusercontent.com/u/`)
//line internal/templates/v2/instruction.qtpl:410
	qw422016.E().S(exampleProfile.SocialProviderUserID)
//line internal/templates/v2/instruction.qtpl:410
	qw422016.N().S(`?v=4&s=192"
                                            alt="`)
//line internal/templates/v2/instruction.qtpl:411
	qw422016.E().S(exampleProfile.GetName())
//line internal/templates/v2/instruction.qtpl:411
	qw422016.N().S(` profile photo"
                                            class="step-4__photo"
                                    />
                                </div>
                                <h3 class="step-4__profile-name">`)
//line internal/templates/v2/instruction.qtpl:415
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:415
	qw422016.N().S(`</h3>
                                <button class="step-4__edit-profile">
                                    <a href="https://github.com/settings/profile" class="step-4__edit-profile-link">Edit
                                        profile</a>
                                </button>
                            </div>
                            <div class="step-4__right-section">
                                <div class="step-4__readme-file">
                                    <p class="step-4__readme-user">
                                        <span class="step-4__readme-name">`)
//line internal/templates/v2/instruction.qtpl:424
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:424
	qw422016.N().S(`</span> /
                                        README.md
                                    </p>
                                    <button class="step-4__edit">
                                        <a href="https://github.com/`)
//line internal/templates/v2/instruction.qtpl:428
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:428
	qw422016.N().S(`/`)
//line internal/templates/v2/instruction.qtpl:428
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:428
	qw422016.N().S(`/edit/main/README.md"
                                           class="step-4__edit-link">
                                            <img src="/assets/images/edit.svg" width="12" height="12" alt="Edit"/>
                                        </a>
                                    </button>
                                    <p class="step-4__readme-code">Hi there 👋</p>
                                </div>
                                <div class="step-4__popular-repo">
                                    <h4 class="step-4__popular-title">
                                        Popular repositories
                                    </h4>
                                    <div class="step-4__popular-item">
                                        <a href="https://github.com/`)
//line internal/templates/v2/instruction.qtpl:440
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:440
	qw422016.N().S(`/`)
//line internal/templates/v2/instruction.qtpl:440
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:440
	qw422016.N().S(`"
                                           class="step-4__repo-link">`)
//line internal/templates/v2/instruction.qtpl:441
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:441
	qw422016.N().S(`</a>
                                        <span class="step-4__repo-setings">Public</span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="step__view">
                    <div class="step__subtitle">
                        and add the counter to the README.md file using the Markdown syntax
                    </div>
                    <div class="step__main step__main-padding step-4">
                        <div class="step-4__readme-description">
                            <span class="step-4__file-user-name">`)
//line internal/templates/v2/instruction.qtpl:455
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:455
	qw422016.N().S(`/`)
//line internal/templates/v2/instruction.qtpl:455
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:455
	qw422016.N().S(`</span>
                            is a special repository: its README.md will appear on your profile!
                        </div>
                        <div class="step-4__file">
                            <div class="step-4__file-buttons">
                                <button class="step-4__file-edit">
                                    <img
                                            src="/assets/images/tag-code.svg"
                                            alt="code"
                                            width="10"
                                            height="10"
                                    />Edit file
                                </button>
                                <button class="step-4__file-preview">
                                    <img
                                            src="/assets/images/eye.svg"
                                            alt="code"
                                            width="10"
                                            height="10"
                                    />Preview
                                </button>
                            </div>
                            <div class="step-4__code">
                                <ol class="step-4__code-list">
                                    <li class="step-4__code-item">### Hi there 👋</li>
                                    <li class="step-4__code-item"></li>
                                    <li class="step-4__code-item">
                                        <span class="step-4__markdown">
                                            [![Hits](https://u8views.com/api/v1/github/profiles/`)
//line internal/templates/v2/instruction.qtpl:483
	qw422016.E().S(exampleProfile.SocialProviderUserID)
//line internal/templates/v2/instruction.qtpl:483
	qw422016.N().S(`/views/day-week-month-total-count.svg)](https://u8views.com/github/`)
//line internal/templates/v2/instruction.qtpl:483
	qw422016.E().S(exampleProfile.Username)
//line internal/templates/v2/instruction.qtpl:483
	qw422016.N().S(`)
                                        </span>
                                    </li>
                                </ol>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="instruction__step-5 `)
//line internal/templates/v2/instruction.qtpl:494
	if done {
//line internal/templates/v2/instruction.qtpl:494
		qw422016.N().S(`active`)
//line internal/templates/v2/instruction.qtpl:494
	}
//line internal/templates/v2/instruction.qtpl:494
	qw422016.N().S(` step">
            <div class="step__progress">
                <div class="step__done">
                    <img
                            src="/assets/images/done.svg"
                            class="step__done-img"
                            alt="done"
                    />
                </div>
                <div class="step__progress-bar"></div>
            </div>
            <div class="step__view">
                <h2 class="step__title">
                    View and analyze views statistics for your GitHub profile
                </h2>
                <div class="step__subtitle">
                    You can also see referral registration statistics from your profile
                </div>
                <div class="step__main step-5">
                    <img src="/assets/images/graph.png" alt="Views chart example"/>
                    <img src="/assets/images/graph-rows.png" alt="Registration chart example"/>
                </div>
            </div>
        </div>
    </div>
</section>
`)
//line internal/templates/v2/instruction.qtpl:520
}

//line internal/templates/v2/instruction.qtpl:520
func writeinstruction(qq422016 qtio422016.Writer, currentPageProfile ProfileView, sessionProfile ProfileView, exampleProfile ProfileView, done bool) {
//line internal/templates/v2/instruction.qtpl:520
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/templates/v2/instruction.qtpl:520
	streaminstruction(qw422016, currentPageProfile, sessionProfile, exampleProfile, done)
//line internal/templates/v2/instruction.qtpl:520
	qt422016.ReleaseWriter(qw422016)
//line internal/templates/v2/instruction.qtpl:520
}

//line internal/templates/v2/instruction.qtpl:520
func instruction(currentPageProfile ProfileView, sessionProfile ProfileView, exampleProfile ProfileView, done bool) string {
//line internal/templates/v2/instruction.qtpl:520
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/templates/v2/instruction.qtpl:520
	writeinstruction(qb422016, currentPageProfile, sessionProfile, exampleProfile, done)
//line internal/templates/v2/instruction.qtpl:520
	qs422016 := string(qb422016.B)
//line internal/templates/v2/instruction.qtpl:520
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/templates/v2/instruction.qtpl:520
	return qs422016
//line internal/templates/v2/instruction.qtpl:520
}
