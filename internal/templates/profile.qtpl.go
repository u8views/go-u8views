// Code generated by qtc from "profile.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamProfile(qw422016 *qt422016.Writer, profile ProfileView, currentUser CurrentUserView, stats ProfileViewsStats) {
	qw422016.N().S(`
<!DOCTYPE html>
<html class="page-root" lang="uk">

<head>
    <title>`)
	qw422016.E().S(profile.GetName())
	qw422016.N().S(` profile views statistic</title>
    <meta name="description" content="`)
	qw422016.E().S(profile.GetName())
	qw422016.N().S(` profile views statistic">
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="icon" sizes="192x192" href="/assets/files/icon192.png">
    <link rel="shortcut icon" sizes="16x16" href="/assets/files/favicon.ico">
    <link rel="shortcut icon" type="image/png" sizes="32x32" href="/assets/files/icon192.png">

    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500&display=swap" rel="stylesheet">
    <script src="https://cdn.jsdelivr.net/npm/apexcharts"></script>

    <style>html{box-sizing:border-box}*,:before,:after{box-sizing:inherit}html{-webkit-text-size-adjust:100%;-moz-text-size-adjust:100%;text-size-adjust:100%}body{-webkit-overflow-scrolling:touch}html,body,div,span,applet,object,iframe,h1,h2,h3,h4,h5,h6,p,blockquote,pre,a,abbr,acronym,address,big,cite,code,del,dfn,em,img,ins,kbd,q,s,samp,small,strike,strong,sub,sup,tt,var,b,u,i,center,dl,dt,dd,ol,ul,li,fieldset,form,label,legend,table,caption,tbody,tfoot,thead,tr,th,td,article,aside,canvas,details,embed,figure,figcaption,footer,header,menu,nav,output,ruby,section,summary,time,mark,audio,video{margin:0;padding:0;border:0}article,aside,details,figcaption,figure,footer,header,main,menu,nav,section,summary{display:block}audio,canvas,progress,video{display:inline-block}audio:not([controls]){display:none;height:0}a{background-color:transparent}abbr[title]{border-bottom:none;text-decoration:underline;-webkit-text-decoration:underline dotted;text-decoration:underline dotted}b,strong{font-weight:bolder}dfn{font-style:italic}mark{background-color:#ff0;color:#000}svg:not(:root){overflow:hidden}code,kbd,pre,samp{font-family:monospace,monospace;font-size:1em}hr{box-sizing:content-box;height:0;overflow:visible}button,input,select,textarea{font:inherit;margin:0}button,input{overflow:visible}button,select{text-transform:none}button,[type=button],[type=reset],[type=submit]{-webkit-appearance:button;-moz-appearance:button;appearance:button}input,textarea,button,select,a{-webkit-tap-highlight-color:transparent}address{font-style:normal}a:focus:not(:focus-visible),select:focus:not(:focus-visible),input:focus:not(:focus-visible),textarea:focus:not(:focus-visible){outline:0}button::-moz-focus-inner,[type=button]::-moz-focus-inner,[type=reset]::-moz-focus-inner,[type=submit]::-moz-focus-inner{border-style:none;padding:0}button,input[type=reset],input[type=button],input[type=submit]{cursor:pointer}button[disabled],input[disabled]{cursor:default}button{-webkit-appearance:none;-moz-appearance:none;appearance:none;background:0 0;padding:0;border:0;border-radius:0;line-height:1}button:focus:not(:focus-visible){outline:0}a,a:hover{text-decoration:none}[href="javascript:void();"],[href="javascript:"]{cursor:default}ul,ol{list-style:none}blockquote,q{quotes:none}blockquote:before,blockquote:after,q:before,q:after{content:none}table{border-collapse:collapse;border-spacing:0}input[type=text],input[type=password],input[type=date],input[type=datetime],input[type=datetime-local],input[type=email],input[type=month],input[type=number],input[type=search],input[type=tel],input[type=time],input[type=url],input[type=week],textarea{box-sizing:border-box}[type=checkbox],[type=radio]{box-sizing:border-box;margin:0;padding:0}input[type=search]{-webkit-appearance:textfield;-moz-appearance:textfield;appearance:textfield;outline-offset:-2px}input[type=search]::-webkit-search-decoration,input[type=search]::-webkit-search-cancel-button{-webkit-appearance:none;appearance:none}input[type=number]::-webkit-inner-spin-button,input[type=number]::-webkit-outer-spin-button{height:auto;-webkit-appearance:none;appearance:none}::-webkit-file-upload-button{-webkit-appearance:button;appearance:button;font:inherit}input[type=number]{-webkit-appearance:textfield;-moz-appearance:textfield;appearance:textfield}select{width:100%;height:20px;border:0;background:0 0}textarea{resize:none;border:0;overflow:auto}::-webkit-input-placeholder{color:#777;line-height:normal}::-moz-placeholder{color:#777;line-height:normal}::placeholder{color:#777;line-height:normal}[hidden]{display:none}.headline{font-family:"Inter",arial,helvetica neue,helvetica,sans-serif;font-weight:500;color:#111028;line-height:120%}.headline--lvl1{font-size:44px}.headline--lvl2{font-size:38px}.headline--lvl3{font-size:28px}.headline--lvl4{font-size:18px}.button{font-family:"Inter",arial,helvetica neue,helvetica,sans-serif;font-size:18px;font-weight:500;display:-webkit-inline-flex;display:inline-flex;-webkit-align-items:center;align-items:center;-webkit-justify-content:center;justify-content:center;height:48px;border-radius:4px;padding-left:32px;padding-right:32px;border-width:2px;border-style:solid;border-color:transparent;transition:background-color .225s ease,box-shadow .225s ease,border-color .225s ease,color .225s ease}.button--black{color:#fff;background-color:#111028;border-color:#111028}.button--black:hover{background-color:#3f3a92;border-color:#3f3a92}.button--black:active{background-color:#2f2c6d;border-color:#2f2c6d}.button--black:focus-visible{background-color:#3f3a92;border-color:#dcdbf0;box-shadow:none}.button--black:disabled{color:#79769c;background-color:#dcdbf0;border-color:#dcdbf0}.input{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column}.input.is-active .input__clear{display:-webkit-flex;display:flex}.input.is-error .input__error{display:-webkit-flex;display:flex}.input.is-error .input__element{border-color:#e9331a}.input.is-error .input__error-text{display:block}.input__label{font-size:16px;line-height:19px;color:#8c8ba3}.input__wrapper{position:relative;margin-top:8px}.input__element{font-size:18px;line-height:22px;color:#111028;height:48px;width:100%;padding-left:20px;padding-right:48px;border-radius:4px;border-width:1px;border-style:solid;border-color:transparent;background-color:#f3f3f6;transition:border-color .225s ease,box-shadow .225s ease}.input__element:hover{border-color:#918fc7;box-shadow:0 0 6px rgba(67,66,88,.16)}.input__element:focus{box-shadow:none;outline:0}.input__element::-webkit-input-placeholder{font-size:18px;line-height:22px;color:#adacc2}.input__element::-moz-placeholder{font-size:18px;line-height:22px;color:#adacc2}.input__element::placeholder{font-size:18px;line-height:22px;color:#adacc2}.input__element:disabled{pointer-events:none}.input__element:disabled::-webkit-input-placeholder{color:#d4d4d4}.input__element:disabled::-moz-placeholder{color:#d4d4d4}.input__element:disabled::placeholder{color:#d4d4d4}.input__clear{display:none;position:absolute;right:20px;top:12px;width:24px;height:24px;-webkit-align-items:center;align-items:center;-webkit-justify-content:center;justify-content:center}.input__error{display:none;position:absolute;right:20px;top:12px}.input__error-text{display:none;font-size:14px;line-height:17px;color:#e9331a;position:absolute;top:calc(100% + 8px);left:0}.select__label{font-size:16px;line-height:120%;color:#8c8ba3}.select__wrapper{height:40px;padding-top:8px;padding-bottom:8px;padding-left:15px;padding-right:15px;border-radius:4px;border:1px solid #dcdbf0;margin-top:12px;transition:border-color .225s ease}.select__wrapper:focus-within{border-color:#b6b5d9}.select__element{font-size:18px;line-height:25px;color:#201f3a;cursor:pointer}.select__element option{font-size:16px;color:#7c7b88}.select__element option:disabled{color:#cbcadb}.select__element:focus-visible{outline:0}.checkbox{font-size:18px;color:#3c3a59;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;height:24px;cursor:pointer}.checkbox__input{position:absolute;-webkit-appearance:none;-moz-appearance:none;appearance:none}.checkbox__input:focus-visible{width:18px;height:18px;outline:0;box-shadow:0 0 0 2px #000,0 0 0 3px #fff;border-radius:4px}.checkbox__input:checked+.checkbox__element{background-color:#3c3a59;background-size:17px 14px;background-repeat:no-repeat;background-position:3px center;background-image:url("data:image/svg+xml,%3Csvg width='17' height='14' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath fill-rule='evenodd' clip-rule='evenodd' d='M16.173.26a1 1 0 0 1 .067 1.413L5.534 13.449.293 8.207a1 1 0 1 1 1.414-1.414l3.759 3.758L14.76.327A1 1 0 0 1 16.173.26Z' fill='%23fff'/%3E%3C/svg%3E")}.checkbox__element{width:24px;height:24px;border:1px solid #3c3a59;border-radius:4px;background-color:#fff;margin-right:12px;transition:background-color .225s ease}.filters__group{padding-top:32px}.filters__sub-headline{font-size:18px;font-weight:600;color:#3c3a59;height:24px;line-height:24px}.filters__visibility{width:28px;height:28px;position:relative;cursor:pointer}.filters__visibility:focus-visible{outline:0;box-shadow:0 0 0 2px #000,0 0 0 3px #fff;border-radius:4px}.filters__visibility:before{content:"";position:absolute;top:50%;left:50%;width:14px;height:2px;margin-left:-7px;margin-top:-1px;background-color:#828282}.filters__elements{margin-top:16px;max-height:160px;overflow:auto}.filters__elements--no-scroll{overflow:initial}.filters__elements::-webkit-scrollbar{width:4px}.filters__elements::-webkit-scrollbar-track{border-radius:8px;background-color:#dddce4}.filters__elements::-webkit-scrollbar-thumb{border-radius:8px;background-color:#b8b6e2}.filters__filled-elements{display:-webkit-flex;display:flex;-webkit-flex-wrap:wrap;flex-wrap:wrap;gap:8px;margin-top:16px}.filters__element:not(:first-child){margin-top:16px}html{height:100%}body{font-family:"Inter",arial,helvetica neue,helvetica,sans-serif;font-size:16px;color:#21212b;display:grid;grid-template-rows:auto 1fr auto;min-height:100%;background-color:#f2f2f3}.wrapper{max-width:1248px;width:100%;padding-left:16px;padding-right:16px;margin-left:auto;margin-right:auto}.disabled-scroll{overflow-y:hidden}.visually-hidden:not(:focus):not(:active){position:absolute;width:1px;height:1px;margin:-1px;border:0;padding:0;white-space:nowrap;-webkit-clip-path:inset(100%);clip-path:inset(100%);clip:rect(0 0 0 0);overflow:hidden}.header__wrapper{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.header__search{width:100%;padding:40px;border-radius:24px;background-color:#fff;margin-top:80px}.header__action-group{display:-webkit-flex;display:flex;-webkit-align-items:flex-end;align-items:flex-end;-webkit-justify-content:space-between;justify-content:space-between;margin-top:32px}.header__input{margin-top:32px}.footer__group{color:rgba(0,0,0,.8);display:grid;grid-template-columns:1fr auto auto;-webkit-align-items:start;align-items:start;padding-top:40px;padding-bottom:60px;padding-left:0;padding-right:0;border-top:1px solid #d0cfdf}.footer__copyrights{font-size:14px;line-height:20px;color:#585764;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;-webkit-column-gap:14px;-moz-column-gap:14px;column-gap:14px}.footer__dev{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;margin-right:95px}.footer__dev:last-child{margin-right:45px}.footer__dev-label{font-size:12px;line-height:17px;color:#9f9ead}.footer__dev-name{font-size:14px;line-height:20px;color:#585764;margin-top:8px}.footer__dev-link{font-size:14px;line-height:20px;color:#3c2df9;-webkit-align-self:flex-start;align-self:flex-start;margin-top:4px}</style>
    <style>.footer{background:#020111}.footer__wrapper{padding:80px 0 32px;display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between}.footer__info{color:#fff}.footer__title{display:block;font-weight:600;font-size:24px;line-height:120%;margin-bottom:16px;color:#fff}.footer__subtitle{line-height:140%;margin-bottom:48px}.footer__map{margin-bottom:56px}.footer__copyrights{display:-webkit-flex;display:flex;gap:14px;color:#cfd2d7;font-size:14px}.footer__support{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;gap:16px;-webkit-align-items:flex-start;align-items:flex-start}.footer__link{padding:16.5px 32px 16.5px 40px;border:1px solid #636077;border-radius:4px;color:#fff;transition:border .3s}.footer__link:hover{border:1px solid #fff}.footer__link:active{border:1px solid #fff;background:#060423}.footer__figure{display:-webkit-flex;display:flex;gap:12px}.header{box-shadow:0 4px 12px rgba(134,132,177,.1);background:#fff;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;min-height:80px}.header__wrapper{display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between;-webkit-align-items:center;align-items:center;width:1216px;height:100%}.header__logo{line-height:0;display:block;margin-right:auto}.header__profile{position:relative;margin-left:40px}.modal.active{opacity:1;pointer-events:all}.modal{position:absolute;display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;width:400px;right:-57%;top:166%;border-radius:10px;background:#fff;opacity:0;pointer-events:none}.modal::after{content:"";position:absolute;right:40px;top:-12px;width:0;height:0;border-left:12px solid transparent;border-right:12px solid transparent;border-bottom:12px solid #fff;border-radius:1px}.modal__profile{display:-webkit-flex;display:flex;gap:16px;padding:24px;border-bottom:1px solid #dddce4}.modal__profile-name{font-weight:600;font-size:18px;line-height:140%;color:#020111}.modal__link{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:4px;font-size:14px;line-height:140%;color:#0d1dab}.modal__log-out{padding:24px}.modal__button{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:11px}.hero{background-image:url(/assets/files/bg.jpg);background-repeat:no-repeat;background-size:cover;background-position:center;display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;-webkit-justify-content:center;justify-content:center;-webkit-align-items:center;align-items:center;color:#fff;margin:64px 0 104px;height:600px;border-radius:16px;padding:0 300px;text-align:center}.hero__title{font-size:64px;line-height:120%;font-weight:600;margin-bottom:24px}.hero__subtitle{font-size:21px;line-height:140%;margin-bottom:48px}.hero__button{display:-webkit-flex;display:flex;padding:16px 56px;background:#fff;border-radius:4px;transition:.9s}.hero__figure{display:-webkit-flex;display:flex;gap:10px;-webkit-align-items:center;align-items:center;color:#24292f}.hero__button:hover figcaption{color:#3f3a92}.hero__button:hover path{fill:#3f3a92}.hero__button:active figcaption{color:#2f2c6d}.hero__button:active path{fill:#2f2c6d}.history{margin-bottom:104px}.history__title{font-size:48px;line-height:120%;margin-bottom:16px}.history__subtitle{line-height:140%;color:#7c7b88;width:590px;margin-bottom:40px}.history__main{background:#fff;border-radius:16px;padding:40px}.history__main-values{width:100%;display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between;margin-bottom:12px;color:#7c7b88}.history__list{height:717px;overflow:auto;width:100%}.history__card{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;padding:24px 0;border-top:1px solid #dddce4;margin-right:24px}.history__user{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:16px;width:480px;margin-right:16px}.history__user-name{font-weight:600;font-size:21px;line-height:140%;color:#24292f}.history__link{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:4px;color:#0d1dab}.history__user-github{width:400px;text-overflow:ellipsis;white-space:nowrap;overflow:hidden}.history__badges{display:-webkit-flex;display:flex;gap:2px;font-size:14px;line-height:140%}.history__badge{display:-webkit-flex;display:flex;color:#fff}.history__badge-title{display:block;padding:2px 6px 3px 8px;background:#000;border:1px solid #e4eaf1;border-radius:4px 0 0 4px}.history__badge-count{display:block;padding:2px 6px 3px 8px;background:#6d96ff;border-width:1px 1px 1px 0;border-style:solid;border-color:#e4eaf1;border-radius:0 4px 4px 0}.history__user-time{margin-left:auto;color:#7c7b88}.history ::-webkit-scrollbar{width:10px}.history ::-webkit-scrollbar-bottom{width:4px}.history ::-webkit-scrollbar-track{background-color:#f2f2f8}.history ::-webkit-scrollbar-thumb{background-color:#b8b6e2;border-radius:10px}.stat-reg{margin-top:64px}.stat-reg__title{font-weight:600;font-size:48px;line-height:120%;margin-bottom:16px}.stat-reg__subtitle{line-height:140%;color:#7c7b88;margin-bottom:40px;max-width:524px}.stat-reg .sorting__list{text-align:center;top:116%;right:3%;bottom:unset!important;left:unset!important}.stat-reg .sorting__current{margin-left:auto}.profile{margin-top:64px;color:#111028}.profile__card{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:32px;margin-bottom:40px}.profile__info{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;gap:8px}.profile__name{font-weight:600;font-size:48px;line-height:120%;color:#020111}.profile__link{display:block;font-size:18px;line-height:140%;color:#0d1dab;margin-bottom:12px}.profile__statistics{background:#fff;border-radius:16px;width:100%;height:488px;padding:40px;margin-bottom:104px;color:#111028}.profile__header{display:-webkit-flex;display:flex;-webkit-align-items:flex-end;align-items:flex-end;margin-bottom:40px}.profile__title{font-weight:600;font-size:28px;line-height:120%}.profile__view-count{font-weight:400;font-size:18px;line-height:140%;margin-left:auto;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.profile__circle{display:inline-block;width:16px;height:16px;border-radius:50%;background:#24292f;margin-right:12px}.profile__select{margin-left:64px;width:90px}.sorting__list{text-align:center;top:116%;right:3%;bottom:unset!important;left:unset!important}.sorting__current{margin-left:auto}.badge{margin-bottom:131px}.badge__title{font-weight:600;font-size:48px;line-height:120%;margin-bottom:16px}.badge__subtitle{line-height:140%;color:#7c7b88;margin-bottom:40px}.badge__example{background:#fff;border-radius:16px;width:100%;word-wrap:break-word;padding:32px}.badge__code{margin-bottom:16px;line-height:140%}.badge__copy{position:relative;margin-left:auto;overflow:hidden}.badge__copy-button{display:-webkit-flex;display:flex;-webkit-justify-content:center;justify-content:center;-webkit-align-items:center;align-items:center;gap:16px;width:200px;height:45px;margin-left:auto;background:#111028;color:#fff;font-weight:600;font-size:18px;line-height:140%;border-radius:4px;transition:background .3s}.badge__copy-button:hover{background:#232149}.badge__copy-button:active{background:#060423}.badge__copy-info{display:-webkit-flex;display:flex;-webkit-justify-content:center;justify-content:center;-webkit-align-items:center;align-items:center;gap:16px;position:absolute;right:-275px;top:0;width:275px;pointer-events:none;background:#3f3a92;color:#fff;height:45px;font-weight:600;font-size:18px;line-height:140%;border-radius:4px}@-webkit-keyframes anim{0%{right:-275px}10%{right:0}90%{right:0}}@keyframes anim{0%{right:-275px}10%{right:0}90%{right:0}}.wrapper-u8{max-width:1216px;margin:0 auto;color:#020111}body{font-size:18px}main{background:#f2f2f8}.sorting{margin:40px 0 24px;padding-bottom:16px;border-bottom:1px solid #dddce4;font-size:18px;display:-webkit-flex;display:flex}.sorting__item{display:-webkit-flex;display:flex;margin-right:32px;position:relative}.sorting__item h4{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.sorting__item:last-child{margin-right:0}.sorting__title{padding:8px;padding-right:12px}.sorting__select{position:relative;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.sorting__current{width:-webkit-max-content;width:-moz-max-content;width:max-content;border:0;font-weight:600;font-size:18px;line-height:140%;color:#111028;padding:0;padding-right:8px;text-align:left;background:0 0;position:relative}.sorting__icon{padding:7px 4px}.sorting__list{opacity:0;pointer-events:none;display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;padding:8px 0;border:1px solid #ededf8;box-shadow:0 6px 16px rgba(17,16,40,.15);border-radius:4px;background:#fff;color:#3c3a59;font-size:18px;line-height:140%;position:absolute;bottom:-156px;left:80px;z-index:1;width:-webkit-max-content;width:-moz-max-content;width:max-content}.sorting__option{padding:8px 16px;text-align:left;background:inherit;transition:background .3s;color:#7c7b88}.sorting__option:hover{background:#f2f2f8}.sorting__list.is-visible{transition:opacity .4s;opacity:1;pointer-events:visible}</style>
</head>

<body>
    <header class="header ">
	<div class="header__wrapper auth wrapper-u8">
		<a href="/" class="header__logo">
			<img class="header__logo-img" src="/assets/files/logo.svg" alt=""></a>
		<div class="header__stars">
			<iframe src="https://ghbtns.com/github-btn.html?user=u8views&repo=go-u8views&type=star&count=true&size=large"
				frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
		</div>
		<div class="header__profile">
			<button class="header__profile-button">
				<img src="https://avatars.githubusercontent.com/u/`)
	qw422016.E().S(currentUser.SocialProviderUserID)
	qw422016.N().S(`?v=4&s=48" alt="photo-profile">
			</button>
			<div class="header__modal modal">
				<div class="modal__profile">
					<img src="https://avatars.githubusercontent.com/u/`)
	qw422016.E().S(currentUser.SocialProviderUserID)
	qw422016.N().S(`?v=4&s=48" width="32" height="32" alt="photo-profile">
					<div class="modal__profile-info">
						<div class="modal__profile-name">`)
	qw422016.E().S(currentUser.GetName())
	qw422016.N().S(`</div>
						<a href="https://github.com/`)
	qw422016.E().S(currentUser.Username)
	qw422016.N().S(`" class="modal__link">
							<img src="/assets/files/link.svg" width="16" height="16" alt="link">
							<span class="modal__profile-github">https://github.com/`)
	qw422016.E().S(currentUser.Username)
	qw422016.N().S(`</span>
						</a>
					</div>

				</div>
				<div class="modal__log-out">
					<a href="/logout" class="modal__button">
						<img src="/assets/files/log-out.svg" alt="log-out">
						<span>Log out</span>
					</a>
				</div>
			</div>
		</div>
	</div>
</header>

<main class="main">
<div class="wrapper-u8">
    <section class="profile">
        <div class="profile__card">
            <img class="profile__photo" src="https://avatars.githubusercontent.com/u/`)
	qw422016.E().S(profile.SocialProviderUserID)
	qw422016.N().S(`?v=4&s=176" alt="profile">
            <div class="profile__info">
                <span class="profile__name">`)
	qw422016.E().S(profile.GetName())
	qw422016.N().S(`</span>
                <a href="" class="profile__link">
                    <img src="/assets/files/link.svg" width="16" height="16" alt="link">
                    <span class="profile__github">https://github.com/`)
	qw422016.E().S(profile.Username)
	qw422016.N().S(`</span>
                </a>
                <div class="history__badges">
                    <div class="history__badge">
                        <span class="history__badge-title">Views per day</span>
                        <span class="history__badge-count">`)
	qw422016.N().DL(stats.DayCount)
	qw422016.N().S(`</span>
                    </div>
                    <div class="history__badge">
                        <span class="history__badge-title">Views per week</span>
                        <span class="history__badge-count">`)
	qw422016.N().DL(stats.WeekCount)
	qw422016.N().S(`</span>
                    </div>
                    <div class="history__badge">
                        <span class="history__badge-title">Views per month</span>
                        <span class="history__badge-count">`)
	qw422016.N().DL(stats.MonthCount)
	qw422016.N().S(`</span>
                    </div>
                    <div class="history__badge">
                        <span class="history__badge-title">Total views</span>
                        <span class="history__badge-count">`)
	qw422016.N().DL(stats.TotalCount)
	qw422016.N().S(`</span>
                    </div>
                </div>
            </div>
        </div>
        <div class="profile__statistics">
            <div class="profile__header">
                <h3 class="profile__title">Your GitHub profile views statistic</h3>
                <div class="profile__view-count">
                    <span class="profile__circle"></span>View count
                </div>
                <div class="profile__select">
                    <div class="sorting__item">
                        <button class="sorting__select sorting-select-js">
                            <input class="sorting__current sorting-current-js" type="button" value="Day">
                            <img class="sorting__icon" src="/assets/files/arrow-down.svg" alt="arrow-down">
                        </button>
                        <div class="sorting__list sorting-list-js">
                            <button class="sorting__option sorting-option-js" value="Day">Day</button>
                            <button class="sorting__option sorting-option-js" value="Week">Week</button>
                            <button class="sorting__option sorting-option-js" value="Month">Month</button>
                            <button class="sorting__option sorting-option-js" value="All time">All time</button>
                        </div>
                    </div>
                </div>
            </div>
            <div class="profile__chart">
                <div class="profile__chart chart-js"></div>
            </div>
        </div>
    </section>
    `)
	if profile.ID == currentUser.ID {
		qw422016.N().S(`
    <section class="badge">
        <h2 class="badge__title">Profile badge</h2>
        <p class="badge__subtitle">Copy the badge and place it on your GitHub profile</p>
        <div class="badge__example">
            <p class="badge__code">
                [![Hits](https://u8views.com/api/v1/github/profiles/`)
		qw422016.E().S(profile.SocialProviderUserID)
		qw422016.N().S(`/views/day-week-month-total-count.svg)](https://u8views.com/github/`)
		qw422016.E().S(profile.Username)
		qw422016.N().S(`)
            </p>
            <div class="badge__copy">
                <button class="badge__copy-button button-copy-code-js">
                    <img class="badge__copy-icon" src="/assets/files/copy.svg" alt="copy">
                    <span>Copy</span>
                </button>
                <div class="badge__copy-info copy-notification-js">
                    <img src="/assets/files/check.svg" alt="check">
                    <span> Copied to clipboard </span>
                </div>
            </div>
        </div>
    </section>
    `)
	}
	qw422016.N().S(`
</div>

</main>
<footer class="footer">
	<div class="footer__wrapper wrapper-u8">
		<div class="footer__info">
			<a href="/" class="footer__title">u8views</a>
			<p class="footer__subtitle">Registration statistics on u8views</p>
			<div class="footer__map">
				<img class="footer__map-Ukraine" src="/assets/files/map-of-ukraine.png" alt="Map of Ukraine">
			</div>
			<div class="footer__copyrights">
				<span>© 2023 Yaroslav Podorvanov</span>
				<img class="footer__flag-of-ukraine" src="/assets/files/flag-of-ukraine.svg" alt="Flag Of Ukraine">
			</div>
		</div>
		<div class="footer__support">
			<a href="" class="footer__link">
				<figure class="footer__figure">
					<img class="footer__img" width="93" height="24" src="/assets/files/national-bank-of-ukraine.png" alt="support">
					<figcaption class="footer__caption">Support</figcaption>
					<img src="/assets/files/arrow.svg" alt="arrow">

				</figure>
			</a>
			<a href="" class="footer__link">
				<figure class="footer__figure">
					<img class="" width="60" height="24" src="/assets/files/come-back-alive-ukraine.png" alt="support">
					<figcaption class="footer__caption">Support</figcaption>
					<img src="/assets/files/arrow.svg" alt="arrow">

				</figure>
			</a>
			<a href="" class="footer__link">
				<figure class="footer__figure">
					<figcaption class="footer__caption">war.ukraine.ua</figcaption>
					<img src="/assets/files/arrow.svg" alt="arrow">
				</figure>
			</a>
		</div>
	</div>
</footer>
<script src="/assets/files/u8views_2a3df2711b2de3d1342dd49415ddd24f.js"></script>

</body>

</html>
`)
}

func WriteProfile(qq422016 qtio422016.Writer, profile ProfileView, currentUser CurrentUserView, stats ProfileViewsStats) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamProfile(qw422016, profile, currentUser, stats)
	qt422016.ReleaseWriter(qw422016)
}

func Profile(profile ProfileView, currentUser CurrentUserView, stats ProfileViewsStats) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteProfile(qb422016, profile, currentUser, stats)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}